package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool, bool) {
	isValidFirstName := len(firstName) >= 2
	isValidLastName := len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidNumberOfTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidFirstName, isValidLastName, isValidEmail, isValidNumberOfTickets
}
