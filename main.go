package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference 2022"
	conferenceTickets := 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Book your tickets here to attend")

	fmt.Printf("conferenceName: %T, conferenceTickets: %T, remainingTickets: %T\n", conferenceName, conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets int
	bookings := []string{}

	for {
		fmt.Println("Your First Name?")
		fmt.Scan(&firstName)
		fmt.Println("Your Last Name?")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email")
		fmt.Scan(&email)
		fmt.Println("How many tickets do you want to book?")
		fmt.Scan(&userTickets)
		fmt.Printf("Thanks for registering, %v %v. We have received your order from %v for %v tickets\n", firstName, lastName, email, userTickets)
		bookings = append(bookings, firstName+" "+lastName)

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("Remaining tickets is now %v\n", remainingTickets)
		fmt.Printf("List of users that already booked: %v\n", bookings)

		firstNames := []string{}
		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("First Names: %v\n", firstNames)

	}

}
