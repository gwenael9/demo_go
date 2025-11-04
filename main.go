package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	var lang string
// 	// par défaut : en
// 	flag.StringVar(&lang, "lang", "en", "The required language, es, en, fr...")
// 	flag.Parse()

// 	greeting, err := greeter.Greet(greeter.Language(lang))
// 	if err != nil {
// 		log.Fatalf("Error greeting : %v", err)
// 	}
// 	fmt.Println(greeting)
// }

func main() {
	reader := bufio.NewReader((os.Stdin))
	fmt.Print("Votre choix : ")

	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
