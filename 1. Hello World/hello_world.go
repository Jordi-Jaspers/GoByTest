package main

//imports
import "fmt"

//Constants
const french = "French"
const english = "English"

const englishPrefix = "Hello "
const frenchPrefix = "Bonjour "

const englishAffix = "! How are you doing?"
const frenchAffix = "! Cava?"

//Global variables

//functions

func Hello(name string, language string) string {
	if len(name) == 0 {
		return Hello("world", language)
	}

	switch language {
	case french:
		return frenchPrefix + name + frenchAffix

	case english:
		return englishPrefix + name + englishAffix

	default:
		return Hello(name, english)
	}
}

func main() {
	fmt.Println(Hello("Jordi", ""))
}
