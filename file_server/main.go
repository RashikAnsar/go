package main

import "net/http"

func main() {
	// http.Handle("/", http.FileServer(http.Dir("../")))
	// http.ListenAndServe(":3300", nil)
	// // or
	http.ListenAndServe(":3300", http.FileServer(http.Dir("../")))
}
