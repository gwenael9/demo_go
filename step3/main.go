package main

import "fmt"

func main() {
	var lang language = "fr"
	greeting := greet(lang)
	fmt.Println(greeting)
}

type language string

func greet(l language) string {
	switch l {
	case "en":
		return "Hello World"
	case "fr":
		return "Bonjour le monde"
	case "es":
		return "Hola el mundo"
	default:
		return ""
	}
}
