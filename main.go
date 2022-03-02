package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Lightning Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Weleome to the %v booking\n", conferenceName)
	fmt.Printf("We hope to see you in attendance %v tickets were created\n", conferenceTickets)
	fmt.Printf("%v tickets remain\n", remainingTickets)

	bookings := []string{}

	for {
		var firstName string
		var ticketNumber uint
		var lastName string
		var email string
		// Tom's tocket order

		greetUsers()

		fmt.Println("Enter your firstname:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastname:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want?")
		fmt.Scan(&ticketNumber)

		if ticketNumber >= remainingTickets {

			fmt.Printf("You're requesting too many tickets, add a number less than %v\n", remainingTickets)
			continue

		}

		remainingTickets = remainingTickets - ticketNumber
		bookings = append(bookings, firstName+" "+lastName+",")

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names are: %v\n", firstNames)

		//fmt.Printf("the slice: %v\n", bookings)

		//fmt.Printf("%v %v got %v tickets sent to %v\n", firstName, lastName, ticketNumber, email)
		//fmt.Printf("Thre are %v left\n", remainingTickets)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {

			//end program
			fmt.Println("Our conference is now sold out, congrats you got the last tickets!")
			break

		} else {

			fmt.Printf("We have %v remaining tickets\n", remainingTickets)

		}
	}

}
