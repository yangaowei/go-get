package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
)

func Cmd(cmds string) (result string) {
	//log.Println("run cmd:", cmds)
	var cmd *exec.Cmd
	cmd = exec.Command("/bin/sh", "-c", cmds)
	var domifstat bytes.Buffer
	cmd.Stdout = &domifstat
	err := cmd.Run()
	if err != nil {
		log.Printf("Error while exec cmd %v", err)
		return ""
	}
	result = domifstat.String()
	return
}

func RandInt(min, max int64) int64 {
	//rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

func MD5(s string) (result string) {
	h := md5.New()
	h.Write([]byte(s))
	result = fmt.Sprintf("%x", h.Sum(nil))
	return
}
