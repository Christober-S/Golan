package main //every statement in go written in packages
import "fmt" //In go if we need do use built in packages we need to import explicityly
func main() { // program starts form main
	var airlineName = "Emirates Airline" // Variable declaration
	const airlineTickets = 200           // constant values cannot be changed
	var remainingTickets = 200

	fmt.Printf("Welcome to the %v ticket booking application\n", airlineName) // %v is  the placehoder for value
	fmt.Println("We have total of", airlineTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Book your tickets here!")
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter your no of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
