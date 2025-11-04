package greeter

import "fmt"

type Language string

var phraseBook = map[Language]string{
	"fr": "Bonjour",
	"en": "Hello",
	"es": "Hola",
}

func Greet(l Language) (string, error) {
	greeting, ok := phraseBook[l]
	if !ok {
		return "", fmt.Errorf("unsupported language: %q", l)
	}
	return greeting, nil
}
