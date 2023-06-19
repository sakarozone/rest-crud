package usercontrollers

import "strings"

// Helper functions to check if the password is valid
func containsSpecialChar(s string) bool {
	specialChars := "~`!@#$%^&*()-_=+\\|[{]};:'\",<.>/?"
	for _, char := range specialChars {
		if strings.ContainsRune(s, char) {
			return true
		}
	}
	return false
}

func containsUpperCase(s string) bool {
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func containsNumber(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}
