package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	var conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Println(conferenceName)
	fmt.Printf("We have a total of %v tickets and %v available\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Last Name: ")
	fmt.Scan(&lastName)

	fmt.Print("E-mail: ")
	fmt.Scan(&email)

	fmt.Print("Tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("%v %v <%v> Tickets: %v\n", firstName, lastName, email, userTickets)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
}
