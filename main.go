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

	fmt.Printf("Hello, you are welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	isvalidName = len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail = strings.Contains(email, "@")
	isvalidTicketNumber = userTickets > 0 || userTickets <= remainingTickets


	if isvalidName && isvalidEmail && isvalidTicketNumber{
		// ask user for number of tickets 
		remainingTickets = conferenceTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Hi %v %v. You're welcome to %v. A confirmation email will be sent to %v. We have %v tickets available now\n", firstName, lastName, conferenceName, email, remainingTickets)

		// add the user to bookings after collecting their input
		fmt.Println("These are the bookings we have now:", bookings)
		
	}

    
	

}
