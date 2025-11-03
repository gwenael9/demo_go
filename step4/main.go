package main

import "fmt"

func main() {
	var lang language = "fr"
	greeting := greet(lang)
	fmt.Println(greeting)
}

type language string

var phraseBook = map[language]string{
	"fr": "Bonjour",
	"en": "Hello",
	"es": "Hola",
}

func greet(l language) string {
	greeting, ok := phraseBook[l]
	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}
	return greeting
}
