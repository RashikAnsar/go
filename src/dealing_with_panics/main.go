package main

import "fmt"

func main() {
	sayHello()
	fmt.Println("After the panic was recovered")
}

func sayHello() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	writeSomething()
}

func writeSomething() {
	panic("Write operation error")
}
