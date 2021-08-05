package helpers

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func GetProcesses() string {
	first := exec.Command("ps", "aux")
	second := exec.Command("grep", "-E", "Music")

	// http://golang.org/pkg/io/#Pipe
	reader, writer := io.Pipe()

	// push first command output to writer
	first.Stdout = writer

	// read from first command output
	second.Stdin = reader

	// prepare a buffer to capture the output
	// after second command finished executing
	var buff bytes.Buffer
	second.Stdout = &buff

	first.Start()
	second.Start()
	first.Wait()
	writer.Close()
	second.Wait()

	result := buff.String()
	return result
}

func KillMusic() {
	fmt.Println("💥 Killing it...")
	cmd := exec.Command("killall", "Music")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
