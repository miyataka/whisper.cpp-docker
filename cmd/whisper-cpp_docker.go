package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const preparedImage = "miyataka/whisper.cpp-docker:v0.9.1"

func main() {
	// TODO flag
	// TODO windows support
	// TODO log

	pwd := os.Getenv("PWD")
	volumeOption := pwd + ":/usr/local/src/whisper.cpp/mnt"

	cmd := exec.Command("docker", "run", "-v", volumeOption, preparedImage, "entrypoint.sh", "./mnt/speaking_sample.mp4", "en")
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
	}
}
