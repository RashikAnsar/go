package main

import "net/http"

func main() {
	http.HandleFunc("/", sayHello)
	// http.HandleFunc("/mars", sayHelloMars)
	err := http.ListenAndServe(":3200", nil)
	if err != nil {
		panic(nil)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// http://localhost:3200/?planet=Jupiter
	planet := r.URL.Query().Get("planet")
	if planet == "" {
		planet = "World"
	}
	w.Write([]byte("Hello, " + planet))
}

// func sayHelloMars(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello Mars..."))
// }
