package greeter

import "fmt"

type Language string

var phraseBook = map[Language]string{
	"fr": "Bonjour",
	"en": "Hello",
	"es": "Hola",
}

func greet(l Language) string {
	greeting, ok := phraseBook[l]
	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}
	return greeting
}
