package main

import (
	"fmt"
	"os/exec"
)

func main() {
	lsCommand := exec.Command("ls")
	// lsCommand.Start()
	output, _ := lsCommand.Output()
	lsCommand.Run()
	fmt.Println(lsCommand.Process.Pid)
	fmt.Println(string(output))
}
