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

	var userName string
	var userTickets int
	// ask user for their name
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
