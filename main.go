package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	var bookings = []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isvalidName, isvalidEmail, isvalidTicketNumber := validateUser(firstName, lastName, email, userTickets, remainingTickets)

		if isvalidName && isvalidEmail && isvalidTicketNumber {
			remainingTickets, bookings = bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			firstNames := getUserFirstNames(bookings)
			fmt.Printf("These are the bookings %v for the conference\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("The tickets for %v have been sold out. See you next year!\n", conferenceName)
				break
			}
		} else {
			if !isvalidName {
				fmt.Println("Your first name or last name is too short")
			}

			if !isvalidEmail {
				fmt.Println("Your email does not have @ symbol")
			}

			if !isvalidTicketNumber {
				fmt.Println("The ticket number you entered is invalid")
			}
		}
	}
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets int) {
	fmt.Printf("Hello, you are welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, int) {

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	// collect user data
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want to purchase: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

// function to validate user input dta
func validateUser(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {

	var isvalidName = len(firstName) >= 2 && len(lastName) >= 2
	var isvalidEmail = strings.Contains(email, "@")
	var isvalidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

	return isvalidName, isvalidEmail, isvalidTicketNumber

}

// 
func getUserFirstNames(bookings []string) []string {
	var firstNames = []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func bookTicket(remainingTickets int, userTickets int, bookings []string, firstName string, lastName string, email string, conferenceName string) (int, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Hi %v %v. You're welcome to %v. A confirmation email will be sent to %v. We have %v tickets available now\n", firstName, lastName, conferenceName, email, remainingTickets)

	return remainingTickets, bookings

}
