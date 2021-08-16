package main

import "fmt"

func Hello(name string) string {
	return "Hello " + name + "! How are you doing?"
}

func main() {
	fmt.Println(Hello("Jordi"))
}
