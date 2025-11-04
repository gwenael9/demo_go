package main

import (
	"fmt"
	"strings"
)

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	RecipientEmail string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Printf("Envoi de l'email à %s : %s\n", e.RecipientEmail, message)
	return nil
}

type SmsNotifier struct {
	PhoneNumber string
}

func (s SmsNotifier) Send(message string) error {
	if !strings.HasPrefix(s.PhoneNumber, "06") {
		return fmt.Errorf("numéro invalide : %s", s.PhoneNumber)
	}
	fmt.Printf("Envoi du SMS au %s : %s\n", s.PhoneNumber, message)
	return nil
}

func main() {
	notifToSend := []Notifier{
		EmailNotifier{RecipientEmail: "toto@mail.com"},
		SmsNotifier{PhoneNumber: "0622334455"},
		EmailNotifier{RecipientEmail: "toto2@mail.com"},
		SmsNotifier{PhoneNumber: "0522334455"},
	}

	message := "Mise à jour"

	for _, notifier := range notifToSend {
		err := notifier.Send(message)
		if err != nil {
			fmt.Println("Erreur :", err)
		}
	}
}
