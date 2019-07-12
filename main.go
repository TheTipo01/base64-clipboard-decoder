package main

import (
	"encoding/base64"
	"fmt"
	"github.com/atotto/clipboard"
	"log"
	"net/url"
	"os/exec"
	"runtime"
	"time"
)

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
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

func main() {
	fmt.Println("Running. To stop press CTRL+C")
	for {
		clip, _ := clipboard.ReadAll()
		decoded, _ := base64.StdEncoding.DecodeString(clip)
		if isValidUrl(string(decoded)) {
			log.Print("URL detected")
			openbrowser(string(decoded))
			_ = clipboard.WriteAll(string(decoded))
		}
		time.Sleep(time.Second / 2)
	}

}
