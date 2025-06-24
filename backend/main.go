package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	// cmd := exec.Command("docker", "ps", "--format", "'{{json .}}'")
	// out_bytes, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// out := string(out_bytes)

	out := `{"Command":"\"/usr/bin/entrypoint…\"","CreatedAt":"2025-06-22 12:47:04 -0400 EDT","ID":"b42fe0c1868c","Image":"barumel/docker-astroneer-server:latest","Labels":"com.docker.compose.config-hash=6eed513227598dcec0ffd4af868478cad4581bd57fbdf1e847e84cc95be63f37,com.docker.compose.version=2.36.0,org.opencontainers.image.created=2024-11-24T21:21:31.131Z,org.opencontainers.image.licenses=MIT,org.opencontainers.image.version=2.2.0,com.docker.compose.oneoff=False,org.opencontainers.image.description=Docker Astroneer Dedicated Server,org.opencontainers.image.title=docker-astroneer-server,org.opencontainers.image.url=https://github.com/barumel/docker-astroneer-server,com.docker.compose.container-number=1,com.docker.compose.depends_on=,com.docker.compose.image=sha256:d639a33083359901190abc57df3c9a951b0c26c1857e68ac0aad50024b158a3b,com.docker.compose.project=astroneer,org.opencontainers.image.revision=4ba8173dd11031b8764aa70c94437ac16e903f15,org.opencontainers.image.source=https://github.com/barumel/docker-astroneer-server,com.docker.compose.project.config_files=/home/prushton/containers/astroneer/compose.yaml,com.docker.compose.project.working_dir=/home/prushton/containers/astroneer,com.docker.compose.service=server","LocalVolumes":"3","Mounts":"astroneer_astr…,astroneer_back…,astroneer_stea…","Names":"astroneer-server-1","Networks":"astroneer_default","Ports":"0.0.0.0:8777-\u003e8777/tcp, 0.0.0.0:8777-\u003e8777/udp, :::8777-\u003e8777/tcp, :::8777-\u003e8777/udp","RunningFor":"34 hours ago","Size":"2.79MB (virtual 3.59GB)","State":"running","Status":"Up 2 minutes (unhealthy)"}\n{"Command":"\"/docker-entrypoint.…\"","CreatedAt":"2025-06-22 10:47:16 -0400 EDT","ID":"5a757bad3e01","Image":"nginx","Labels":"com.docker.compose.service=nginx,com.docker.compose.version=2.36.0,com.docker.compose.container-number=1,com.docker.compose.depends_on=,com.docker.compose.oneoff=False,com.docker.compose.project.config_files=/home/prushton/containers/srv/compose.yaml,com.docker.compose.project.working_dir=/home/prushton/containers/srv,maintainer=NGINX Docker Maintainers \u003cdocker-maint@nginx.com\u003e,com.docker.compose.config-hash=c956227cc07afa942ad331d6261711fe8c9744c2d646d3b196c3c0dbf0aa25ee,com.docker.compose.image=sha256:1e5f3c5b981a9f91ca91cf13ce87c2eedfc7a083f4f279552084dd08fc477512,com.docker.compose.project=srv","LocalVolumes":"0","Mounts":"/home/prushton…,/home/prushton…,/home/prushton…,/home/prushton…","Names":"nginx","Networks":"srv_default","Ports":"0.0.0.0:80-\u003e80/tcp, :::80-\u003e80/tcp, 0.0.0.0:443-\u003e443/tcp, :::443-\u003e443/tcp","RunningFor":"36 hours ago","Size":"1.09kB (virtual 192MB)","State":"running","Status":"Up 35 hours"}\n{"Command":"\"/usr/bin/frpc -c /e…\"","CreatedAt":"2025-06-22 10:47:16 -0400 EDT","ID":"603dabfd0ad0","Image":"snowdreamtech/frpc:alpine","Labels":"com.docker.compose.project.config_files=/home/prushton/containers/srv/compose.yaml,com.docker.compose.service=frp_client,org.opencontainers.image.description=Docker Images for Frp.,org.opencontainers.image.title=frp,com.docker.compose.depends_on=,com.docker.compose.version=2.36.0,org.opencontainers.image.base.name=snowdreamtech/frpc:alpine,org.opencontainers.image.documentation=https://hub.docker.com/r/snowdreamtech/frpc,org.opencontainers.image.source=https://github.com/snowdreamtech/frp,com.docker.compose.container-number=1,com.docker.compose.image=sha256:2fc9e95dd5badda46bdbbe0698721fd9afedaebdd206c77d2a6a54eefa58aba8,com.docker.compose.project=srv,com.docker.compose.project.working_dir=/home/prushton/containers/srv,org.opencontainers.image.revision=6b5d2b88a595b93beed0fdfafdd29a39c7738b5c,org.opencontainers.image.url=https://github.com/snowdreamtech/frp,org.opencontainers.image.vendor=Snowdream Tech,org.opencontainers.image.version=latest,com.docker.compose.config-hash=03788a4e6d909608625af44208dc8fe787c7086b0f82b8ffba8fbf575d411b59,com.docker.compose.oneoff=False,org.opencontainers.image.authors=Snowdream Tech,org.opencontainers.image.created=2025-06-18T13:13:41.418Z,org.opencontainers.image.licenses=MIT","LocalVolumes":"0","Mounts":"/home/prushton…","Names":"frp_client","Networks":"host","Ports":"","RunningFor":"36 hours ago","Size":"0B (virtual 40MB)","State":"running","Status":"Up 35 hours (healthy)"}`

	if len(out) < 10 {
		io.WriteString(w, "[]")
		return
	}

	// out = strings.Replace(string(out), "\\\"", "", -1)
	containers := strings.Split(out, "\\n")

	var parsedContainers []SentInfo

	for _, containerString := range containers {
		var container ContainerInfo

		err := json.Unmarshal([]byte(containerString), &container)

		if err != nil {
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
		fmt.Println(err)
		io.WriteString(w, "err")
		return
	}

	io.Writer.Write(w, str)
}

func main() {
	http.HandleFunc("/containerInfo", getContainerInfo)

	_ = http.ListenAndServe(":3000", nil)
}
