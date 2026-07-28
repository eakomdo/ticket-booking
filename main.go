package main

import "fmt"

func main(){

	var conferenceName string = "Go Conference"
	const conferenceTickets int= 50
	var remainingTickets int = 50
	var bookings 

	fmt.Printf("Hello, you are welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")


	var firstName string
	var lastName string
	var email string
	var userTickets int
	

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want to purchase: ")
	fmt.Scan(&userTickets)

	remainingTickets = conferenceTickets - userTickets


	fmt.Printf("Hi %v %v. You're welcome to %v. A confirmation email will be sent to %v. We have %v tickets available now", firstName, lastName, conferenceName, email, remainingTickets)

}