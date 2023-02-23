package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50                   //Can't change this value

var conferenceName = "Go Conference"       
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

// var bookings = make([]map[string]string, 0)    //Syntax to declare booking map
// var book = []string{}                         // Syntax to declare booking slices

type UserData struct {
	firstName string
	lastName string
	emailId string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {					
	greetUsers()	
	for {			
		firstName, lastName, emailId, userTickets := userInputs()
		isValidName, isValidEmail, isValidTicketNumber:= helper.IsValidInput(firstName, lastName, emailId, userTickets, remainingTickets)
		
		// fmt.Println("Valid name:", isValidName)
		// fmt.Println("Valid Email:", isValidEmail)
		// fmt.Println("Valid TicketNumber:", isValidTicketNumber)
		
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicketsStruct(userTickets, firstName, lastName, emailId)
			
			wg.Add(1)
			go sendTicket(firstName, lastName, userTickets, emailId)

		} else {
			invalidError(isValidEmail, isValidName, isValidTicketNumber)
			continue
		} 

		// getSlices()

		firstNames:= getFirstNames(bookings)
		fmt.Println("First names are as follows", firstNames)

		noTicketsRemaining:= remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("All of the tickets are booked! ")
			break
		}
	}
	wg.Wait()	 
} 

func greetUsers() {
	fmt.Println("Welcome to our", conferenceName, "booking application")	
	// fmt.Println("We have total of", conferenceTickets, "tickets out of which", remainingTickets, "tickets are remaining" )
	fmt.Printf("We have total of %v tickets out of which %v tickets are remaining \n", conferenceTickets, remainingTickets )
	fmt.Println("Book your Tickets here")
}

func getFirstNames(bookings []UserData) []string{
	
	firstNames := []string{}

	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func userInputs() (string, string, string, uint) {
	
	var firstName string
	var lastName string
	var emailId string
	var userTickets uint
	
	fmt.Println("Enter the First name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email ID")
	fmt.Scan(&emailId)

	fmt.Println("Enter the Number of tickets you want to book")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailId, userTickets
}

func elementNumber() {
	fmt.Println("Enter the element number you want to see")		
	var elementNumber uint
	fmt.Scan(&elementNumber)
	fmt.Println(bookings[elementNumber])
}

func invalidError(isValidEmail bool, isValidName bool, isValidTicketNumber bool) {
	if !isValidEmail {
		fmt.Println("You have entered wrong emailID, please check again")		
	}
	if !isValidName{
		fmt.Println("You have entered invalid First Name or Last Name. Please Enter more than 2 letters")
	}
	if !isValidTicketNumber{
		fmt.Printf("You have entered wrong number of tickets. Please Enter the the tickets more less than %v tickets \n", remainingTickets )
	}
}


func bookTicketsStruct(userTickets uint, firstName string, lastName string, emailId string) {
	remainingTickets = remainingTickets - userTickets
	
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		emailId: emailId,
		userTickets: userTickets, 
	}		
	bookings = append(bookings, userData)

	// fmt.Println("List of bookings are:", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive an email confirmation on your email ID %v soon :) \n", firstName, lastName, userTickets, emailId)
	fmt.Printf("Remaining tickets are: %v \n", remainingTickets)
}

func sendTicket(firstName string, lastName string,  userTickets uint, emailId string) {
	time.Sleep(10 * time.Second)
	fmt.Printf("!.........................##################################.........................! \n")
	fmt.Printf("Hi %v %v you have booked %v number of tickets which you will receive on your email id %v \n", firstName, lastName, userTickets, emailId)
	fmt.Printf("!.........................##################################.........................! \n")
	wg.Done()
}


// func bookTicketsMap(userTickets uint, firstName string, lastName string, emailId string) {
// 	remainingTickets = remainingTickets - userTickets
	
// 	var userData = make(map[string]string)
// 	userData["FirstName"] = firstName
// 	userData["LastName"] = lastName
// 	userData["emailId"] = emailId
// 	userData["NumberOfTicketsBooked"] = strconv.FormatUint(uint64(userTickets), 10)
	
// 	bookings = append(bookings, userData)

// 	fmt.Println("List of bookings are:", bookings)

// 	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive an email confirmation on your email ID %v soon :) \n", firstName, lastName, userTickets, emailId)
// 	fmt.Printf("Remaining tickets are: %v \n", remainingTickets)
// }

// func getSlices() {
// 	fmt.Printf("The Size of the slice is: %v \n", len(bookings))
// 	fmt.Printf("Type of slice is: %T \n", bookings)
// 	fmt.Printf("Here are the element of the slice: %v \n", bookings)							
// 	fmt.Printf("First element is: %v \n", bookings[0])
// }