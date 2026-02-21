package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

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

func getContainerInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// return

	cmd := exec.Command("docker", "ps", "--format", "{{json .}}")
	out_bytes, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		fmt.Println(string(out_bytes))
		http.Error(w, "Error running docker command", http.StatusInternalServerError)
		io.WriteString(w, "")
		return
	}
	out := string(out_bytes)

	if len(out) < 10 {
		fmt.Println("len < 10")
		http.Error(w, "Command output too short", http.StatusInternalServerError)
		io.WriteString(w, "")
		return
	}

	out = strings.Replace(out, "\\\"", "", 1)

	containers := strings.Split(out, "\n")

	var parsedContainers []SentInfo

	for _, containerString := range containers {
		var container ContainerInfo

		err := json.Unmarshal([]byte(containerString), &container)

		if err != nil {
			fmt.Println("Error parsing container string to object")
			fmt.Println(err)
			fmt.Println(containerString)
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

	str, err := json.Marshal(parsedContainers)

	if err != nil {
		fmt.Println("Error parsing objects to string")
		fmt.Println(err)
		http.Error(w, "Error parsing objects to string", http.StatusInternalServerError)
		io.WriteString(w, "")
		return
	}

	io.Writer.Write(w, str)
}

func main() {
	http.HandleFunc("/containerInfo", getContainerInfo)
	
	fmt.Println("Running")
	_ = http.ListenAndServe(":3000", nil)
	fmt.Println("exiting")

}
