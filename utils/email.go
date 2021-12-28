package utils

import "net/mail"

func CheckValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
