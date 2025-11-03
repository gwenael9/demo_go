package main

import "testing"

// pointeur t
func TestGreet(t *testing.T) {
	// phase de préparation : définition des valeurs attendues
	type testCase struct {
		lang language
		want string
	}

	// ici, string sera le nom des tests
	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello",
		},
		"Français": {
			lang: "fr",
			want: "Bonjour",
		},
		"Espagnol": {
			lang: "es",
			want: "Hola",
		},
		"Not supported, portuguese": {
			lang: "pt",
			want: `Unsupported language: "pt"`,
		},
		"Empty": {
			lang: "",
			want: `Unsupported language: ""`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// phase d'execution : appel de la fonction qu'on veut tester
			got := greet(test.lang)

			// phase de décision : on vérifie les valeurs obtenues
			if test.want != got {
				t.Errorf("Greeting = %q; want: %q", got, test.want)
			}
		})
	}
}
