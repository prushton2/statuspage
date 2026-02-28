package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"os"
	"time"
	"slices"

	"github.com/joho/godotenv"
)

type Cache struct {
	containers []SentInfo
	time time.Time
}

type ContainerInfo struct {
	Command      string `json:"Command"`
	CreatedAt    string `json:"CreatedAt"`
	ID           string `json:"ID"`
	Image        string `json:"Image"`
	Labels       string `json:"Labels"`
	LocalVolumes string `json:"LocalVolumes"`
	Mounts       string `json:"Mounts"`
	Names        string `json:"Names"`
	Networks     string `json:"Networks"`
	Ports        string `json:"Ports"`
	RunningFor   string `json:"RunningFor"`
	Size         string `json:"Size"`
	State        string `json:"State"`
	Status       string `json:"Status"`
}

type SentInfo struct {
	Name    string `json:"Name"`
	Network string `json:"Network"`
	Status  string `json:"Status"`
	Size    string `json:"Size"`
	Health  string `json:"Health"`
}

type HTTPResponse struct {
	Containers     []SentInfo `json:"containers"`
	TopNetworks    []string   `json:"topNetworks"`
	BottomNetworks []string   `json:"bottomNetworks"`
}

var cache Cache = Cache{
	containers: make([]SentInfo, 0),
	time: time.Unix(0, 0),
}

func getComposeOutput() (string, error) {
	cmd := exec.Command("docker", "ps", "--format", "{{json .}}")
	out_bytes, err := cmd.Output()
	
	if err != nil {
		return "", fmt.Errorf("Error running docker command: ", err)
	}
	
	return string(out_bytes), nil
}

func getTestOutput() string {
	b, err := os.ReadFile("test_ps_output.txt") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    return string(b) // convert content to a 'string'
}

func getContainerInfo(ignoreContainers []string) ([]SentInfo, error) {
	
	var out string
	var err error

	if os.Getenv("DEBUG") == "true" {
		out= getTestOutput()
	} else {
		out, err = getComposeOutput()
		if err != nil {
			return nil, err
		}
	}

	if len(out) < 10 {
		return nil, fmt.Errorf("Command output too short")
	}

	out = strings.Replace(out, "\\\"", "", 1)
	containers := strings.Split(out, "\n")
	
	var parsedContainers []SentInfo = make([]SentInfo, 0)

	for _, containerString := range containers {
		if containerString == "" {
			continue
		}

		var container ContainerInfo
		err := json.Unmarshal([]byte(containerString), &container)
		
		if slices.Index(ignoreContainers, container.Names) != -1 {
			continue
		}

		if err != nil {
			fmt.Println("getContainerInfo: Error parsing container (", containerString, ") string to object: ", err)
			continue
		}

    	r, _ := regexp.Compile("\\((health\\: starting|healthy|unhealthy)\\)")
		match := r.FindString(container.Status)

		var status string = strings.ReplaceAll(container.Status, match, "")
		status = strings.TrimSpace(status)

		match = strings.ReplaceAll(match, "(", "")
		var health string = strings.ReplaceAll(match, ")", "")

		var info SentInfo = SentInfo{
			Name:    container.Names,
			Network: container.Networks,
			Size:    container.Size,
			Status:  status,
			Health:  health,
		}

		parsedContainers = append(parsedContainers, info)
	}

	return parsedContainers, nil
}

func httpGetContainerInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	var ignoreContainers []string = strings.Split(os.Getenv("IGNORE_CONTAINERS"), "\n")
	var topNetworks []string = strings.Split(os.Getenv("TOP_NETWORKS"), "\n")
	var bottomNetworks []string = strings.Split(os.Getenv("BOTTOM_NETWORKS"), "\n")
	now := time.Now()

	var containerInfo []SentInfo
	var err error

	if now.Sub(cache.time) > 10 * time.Second {
		containerInfo, err = getContainerInfo(ignoreContainers)

		if err != nil {
			fmt.Println("Error during getContainerInfo:", err)
			http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
			return
		}

		cache.containers = containerInfo
		cache.time = now

	} else {
		containerInfo = cache.containers
	}

	var httpResponse HTTPResponse = HTTPResponse{
		Containers:     containerInfo,
		TopNetworks:    topNetworks,
		BottomNetworks: bottomNetworks,
	}
	
	str, err := json.Marshal(httpResponse)
	
	if err != nil {
		fmt.Println("Final Marshal: Error parsing objects to string")
		fmt.Println(err)
		http.Error(w, "Final Marshal: Error parsing objects to string", http.StatusInternalServerError)
		io.WriteString(w, "")
		return
	}

	io.Writer.Write(w, str)
}

func main() {

	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(err)
		}
	}

	http.HandleFunc("/containerInfo", httpGetContainerInfo)
	
	fmt.Println("Running")
	_ = http.ListenAndServe(":3000", nil)
	fmt.Println("exiting")

}
