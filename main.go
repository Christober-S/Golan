package main //every statement in go written in packages
import "fmt" //In go if we need do use built in packages we need to import explicityly
func main() { // program starts form main
	var airlineName = "Emirates Airline" // Variable declaration
	const airlineTickets = 200           // constant values cannot be changed
	var remainingTickets = 200
	var userName string // To specifiy the type of data
	var userTicket int  // To specifiy the int type of dT

	fmt.Printf("Welcome to the %v ticket booking application", airlineName)

	fmt.Printf("Welcome to the %v ticket booking application", airlineName) // %v is  the placehoder for value
	fmt.Println("We have total of", airlineTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Book your tickets here!")
	fmt.Print("Enter the user name: ")
	fmt.Scan(&userName)
	fmt.Print("Enter the no of tickets: ")
	fmt.Scan(&userTicket)
	fmt.Printf("The user %v has booked %v number of ticktets", userName, userTicket)

}
