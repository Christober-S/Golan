package main //every statement in go written in packages

import "fmt" //In go if we need do use built in packages we need to import explicityly

func main() { // program starts form main

	var airlineName = "Emirates Airline" // Variable declaration
	const airlineTickets = 200           // constant values cannot be changed
	var remainingTickets = 200

	fmt.Println("Welcome to the", airlineName, "ticket booking application")
	fmt.Println("We have total of", airlineTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Book your tickets here!")

}
