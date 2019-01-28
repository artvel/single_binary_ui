package cmd

import (
	"log"
	"os"
	"os/exec"
)

func Run() {
	cmd := exec.Command("./ui_service", "-h", "0.0.0.0", "-p", "58082")
	cmd.Stdout = os.Stdout
	err := cmd.Start()//non blocking
	if err != nil {
		log.Fatal(err)
	}
	cmd2 := exec.Command("java", "-jar", "document-service.jar")
	cmd2.Stdout = os.Stdout
	err = cmd2.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd2.Wait()//block
	if err != nil {
		log.Fatal(err)
	}
}
