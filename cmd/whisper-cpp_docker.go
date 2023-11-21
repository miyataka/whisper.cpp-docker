package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const preparedImage = "miyataka/whisper.cpp-docker:v1.4.3"

func main() {
	// flag
	target := flag.String("target", "", "target file name")
	lang := flag.String("lang", "en", "transcription language. e.g. ja, en, etc...")
	debug := flag.Bool("debug", false, "debug log flag")

	flag.Parse()

	if *debug {
		fmt.Printf("target: %s\n", *target)
		fmt.Printf("lang: %s\n", *lang)
	}

	if *target == "" {
		fmt.Println("target parameter is required.")
		os.Exit(1)
	}

	basename := filepath.Base(*target)
	if !exists(basename) {
		fmt.Println("please place target file your current directory.")
		os.Exit(1)
	}

	path := "./mnt/" + basename

	pwd := os.Getenv("PWD")
	volumeOption := pwd + ":/usr/local/src/whisper.cpp/mnt"

	cmd := exec.Command("docker", "run", "-v", volumeOption, preparedImage, "entrypoint.sh", path, *lang)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	// execute
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
