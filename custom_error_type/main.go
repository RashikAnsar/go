package main

import "fmt"

// MyError has short and detailed message fields
type MyError struct {
	ShortMessage    string
	DetailedMessage string
}

func (e *MyError) Error() string {
	return e.ShortMessage + "\n" + e.DetailedMessage
}

func main() {
	err := doSomething()
	fmt.Println(err)
}

func doSomething() error {
	return &MyError{ShortMessage: "ERROR:: Custom error short message", DetailedMessage: "ERROR DETAILS:: Custom error detailed message"}
}
