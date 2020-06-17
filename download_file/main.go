package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	imgURL := "https://golang.org/doc/gopher/doc.png"
	response, err := http.Get(imgURL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	file, err2 := os.Create("gopher.png")
	if err2 != nil {
		panic(err2)
	}

	_, err3 := io.Copy(file, response.Body)
	if err3 != nil {
		panic(err3)
	}

	file.Close()
	fmt.Println("Image Download is successful")
}
