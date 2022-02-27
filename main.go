package main

import "fmt"

func main() {
	var conferenceName = "Lightning Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Weleome to the %v booking\n", conferenceName)
	fmt.Printf("We hope to see you in attendance %v tickets were created\n", conferenceTickets)
	fmt.Printf("%v tickets remain\n", remainingTickets)

	var userName string
	var tomTix int
	// Tom's tocket order

	fmt.Println("Enter your firstname:")
	fmt.Scan(&userName)
	tomTix = 2

	fmt.Printf("%v got %v tickets", userName, tomTix)

}
