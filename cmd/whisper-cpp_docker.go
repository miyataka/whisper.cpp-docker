package main

import (
	"fmt"
	"os"
	"os/exec"
)

// main runs
func main() {
	// TODO flag
	// TODO error output
	// TODO windows support
	// TODO log

	pwd := os.Getenv("PWD")
	volumeOption := pwd + ":/usr/local/src/whisper.cpp/mnt"

	cmd := exec.Command("docker", "run", "-v", volumeOption, "-it", "miyataka/whisper.cpp-docker:v0.9.0", "entrypoint.sh", "./mnt/speaking_sample.mp4", "ja")
	fmt.Printf("%+v\n", cmd)
	if err := cmd.Run(); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println("end")
}
