package main //every statement in go written in packages
import "fmt" //In go if we need do use built in packages we need to import explicityly
func main() { // program starts form main
	var airlineName = "Emirates Airline" // Variable declaration
	const airlineTickets = 200           // constant values cannot be changed
	var remainingTickets uint = 200

	fmt.Printf("Welcome to the %v ticket booking application\n", airlineName) // %v is  the placehoder for value
	fmt.Println("We have total of", airlineTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Book your tickets here!")
	for { //infinte for loop

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		var bookings []string //Array Slice using this we don't need to specifiy the data type

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("Enter your no of tickets: ")
		fmt.Scan(&userTickets)

		bookings = append(bookings, firstName+" "+lastName) //In slice we can add element using append() method
		remainingTickets = remainingTickets - userTickets   //In go lang we can't do calculation's if both are not same datatypes

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets available for %v\n", remainingTickets, airlineName)
		fmt.Printf("%v\n", bookings[0])
		fmt.Printf("%T\n", bookings)
		fmt.Printf("Length of array:%v\n", len(bookings)) // length of array

	}
}
