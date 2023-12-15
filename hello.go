package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Println("hello, to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("how many tickets do you want to order:")
	fmt.Scan(&userTickets)

	//userName = "Tom"
	//userTickets = 2
	remainingTickets = remainingTickets - userTickets

	fmt.Println("User", firstName, lastName, "booked", userTickets, "tickets", "You will recieve your tickets on mail box", email)
	fmt.Println("The remaining tickets are:", remainingTickets)
}
