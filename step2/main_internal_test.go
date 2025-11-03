package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello
}

// pointeur t
func TestGreet(t *testing.T) {
	// phase de préparation : définition des valeurs attendues
	expectedGreeting := "Hello World"

	// phase d'execution : appel de la fonction qu'on veut tester
	greeting := greet()

	// phase de décision : on vérifie les valeurs obtenues
	if greeting != expectedGreeting {
		t.Errorf("Greeting = %q; want: %q", greeting, expectedGreeting)
	}
}
