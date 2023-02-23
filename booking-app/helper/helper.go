package helper

import (
	"strings"
)


func IsValidInput(firstName string, lastName string, emailId string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(emailId, "@")
	isValidTicketNumber := remainingTickets >= userTickets && userTickets !=0

	return isValidName, isValidEmail, isValidTicketNumber
}
