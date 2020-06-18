package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://golang.org"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	html, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		panic(err2)
	}

	// fmt.Println(string(html))
	err3 := ioutil.WriteFile("golang.html", html, 0644)
	if err3 != nil {
		panic(err3)
	}
}
