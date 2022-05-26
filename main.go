package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T, remainingTickets is %T and remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	//fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Printf("We have total of %v tickets and  %v are still available\n", conferenceTickets, remainingTickets)
	//fmt.Println("Get your tickets here to attend")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email : ")
	fmt.Scan(&email)

	fmt.Println("Enter tickets you want to buy : ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	userTickets = 2
	fmt.Printf("Thank you  %v %v for booking %v tickets, you will receive confirmation email at: %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets remaining for %v\n", remainingTickets, conferenceName)
}
