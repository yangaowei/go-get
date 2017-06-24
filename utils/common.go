package utils

import (
	"bytes"
	"log"
	"os/exec"
)

func Cmd(cmds string) (result string) {
	log.Println("run cmd:", cmds)
	var cmd *exec.Cmd
	cmd = exec.Command("/bin/sh", "-c", cmds)
	var domifstat bytes.Buffer
	cmd.Stdout = &domifstat
	err := cmd.Run()
	if err != nil {
		log.Printf("Error while exec cmd %a", err)
		return ""
	}
	result = domifstat.String()
	return
}
