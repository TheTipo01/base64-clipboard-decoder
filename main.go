package main

import (
	"encoding/base64"
	"github.com/atotto/clipboard"
	"os/exec"
	"runtime"
)

func main() {
	url, _ := clipboard.ReadAll()
	decoded, _ := base64.StdEncoding.DecodeString(url)
	openbrowser(string(decoded))
	_ = clipboard.WriteAll(string(decoded))
}

func openbrowser(url string) {

	switch runtime.GOOS {
	case "linux":
		_ = exec.Command("xdg-open", url).Start()
	case "windows":
		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		_ = exec.Command("open", url).Start()
	}
}
