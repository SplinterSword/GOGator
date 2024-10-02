package main

import (
	"log"
	"os"
	"os/exec"
)

func WriteInTerminal(text string) {
	cmd := exec.Command("echo", text)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
