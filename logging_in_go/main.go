package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.Create("log_file")

	if err != nil {
		fmt.Println("ERROR:: ", err.Error())
	}

	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Writing log file")
	log.Fatalln("ERROR: Application crashed! check log file..")
}
