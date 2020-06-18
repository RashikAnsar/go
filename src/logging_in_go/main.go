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
	// Fatalln() is similar to Println but calls os.Exit(1)
	log.Fatalln("ERROR: Application crashed! check log file..")
}
