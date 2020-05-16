package main

import (
	"net/http"
	"os/exec"
)

func main() {
	http.Get("https://daviddworken.com/hifromgo_asm")
	exec.Command("touch", "/tmp/hifromgo").Run()
}
