package main

import "fmt"

func main() {
	var conferenceName = "Lightning Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Weleome to the %v booking\n", conferenceName)
	fmt.Println("We hope to see you in attendance", conferenceTickets, "tickets were created")
	fmt.Println(remainingTickets, "tickets remain")

}
