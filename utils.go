package main

import (
	log "github.com/alecthomas/log4go"
	"os/exec"
)

func executeCommand(command string) string {
	cmd := exec.Command("/bin/bash", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("run command failed with %s\n", err)
	}
	return string(out)
}
