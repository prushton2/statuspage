package main

import (
	"fmt"
	"net/http"
	"io"
	"os/exec"
	"strings"
)

func getContainerInfo(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("docker", "ps","--format","'{{json .}}'")
	out, err := cmd.Output()
	if(err != nil) {
		fmt.Println(err)
	}
	
	if(len(string(out)) < 10) {
		io.WriteString(w, "[]")
		return;
	}

	str := strings.Replace(string(out), "\\\"", "", -1)
	str = strings.Replace(string(out), "'\n'", ",", -1)

	str = str[1 : len(str)-2]
	str = "[" + str + "]"

	io.WriteString(w, str)	
}

func main() {
	http.HandleFunc("/containerInfo", getContainerInfo)

	_ = http.ListenAndServe(":3000", nil)
}