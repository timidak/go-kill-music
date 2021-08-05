package main

import (
	"fmt"
	"strings"
	"time"

	color "github.com/fatih/color"
	"github.com/timidak/go-kill-music/helpers"
)

const process string = "Music.app/Contents/XPCServices/VisualizerService.xpc/Contents/MacOS/VisualizerService"

func main() {
	fmt.Println(color.GreenString("🚀 Starting..."))
	fmt.Println(color.HiMagentaString("🔫 Waiting for Apple Music to start..."))
	for {
		output := helpers.GetProcesses()
		if strings.Contains(output, process) {
			helpers.KillMusic()
		}
		time.Sleep(time.Second)
	}
}
