package main

import "testing"

// pointeur t
func TestGreet_English(t *testing.T) {
	// phase de préparation : définition des valeurs attendues
	lang := "en"
	want := "Hello World"

	// phase d'execution : appel de la fonction qu'on veut tester
	got := greet(language(lang))

	// phase de décision : on vérifie les valeurs obtenues
	if got != want {
		t.Errorf("Greeting = %q; want: %q", got, want)
	}
}

func TestGreet_French(t *testing.T) {
	lang := "fr"
	want := "Bonjour le monde"

	got := greet(language(lang))

	if got != want {
		t.Errorf("Greeting = %q; want: %q", got, want)
	}
}
