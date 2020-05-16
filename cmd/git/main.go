package main

import (
	"os"
	"os/exec"
	"net/http"
)

func main() {
	http.Get("https://daviddworken.com/hifromgo")

	cmd := exec.Command("/usr/bin/git", os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}
