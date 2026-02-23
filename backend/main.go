package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"os"
	"time"
	"slices"
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

var cache Cache = Cache{
	containers: make([]SentInfo, 0),
	time: time.Now(),
}

func getContainerInfo(ignoreContainers []string) ([]SentInfo, error) {
	
	cmd := exec.Command("docker", "ps", "--format", "{{json .}}")
	out_bytes, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("Error running docker command: ", err)
	}
	
	out := string(out_bytes)
	
	if len(out) < 10 {
		return nil, fmt.Errorf("Command output too short")
	}

	out = strings.Replace(out, "\\\"", "", 1)
	containers := strings.Split(out, "\n")
	
	var parsedContainers []SentInfo = make([]SentInfo, 0)

	// fmt.Println(containers)

	for _, containerString := range containers {
		if containerString == "" {
			continue
		}
		// fmt.Println("Container:", containerString)

		var container ContainerInfo
		err := json.Unmarshal([]byte(containerString), &container)
		
		if slices.Index(ignoreContainers, container.Names) != -1 {
			continue
		}

		if err != nil {
			fmt.Println("getContainerInfo: Error parsing container (", containerString, ") string to object: ", err)
			continue
		}

		split := strings.Split(container.Status, "(")

		var health string = ""
		var status string = split[0]

		if len(split) > 1 {
			health = split[1][:len(split[1])-1]
		}

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
	now := time.Now()

	var containerInfo []SentInfo
	var err error

	// fmt.Println("Time diff: ", now.Sub(cache.time), now.Sub(cache.time) > 10 * time.Second)
	if true {
	// if now.Sub(cache.time) > 10 * time.Second {
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
	
	str, err := json.Marshal(containerInfo)
	
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
	http.HandleFunc("/containerInfo", httpGetContainerInfo)
	
	fmt.Println("Running")
	_ = http.ListenAndServe(":3000", nil)
	fmt.Println("exiting")

}
