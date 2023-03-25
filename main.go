package main //every statement in go written in packages

import "fmt" //In go if we need do use built in packages we need to import explicityly

func main() { // program starts form main
	fmt.Println("Welcome to the Emirates Airline")
	fmt.Println("Book your tickets here!")

	var airlineName = "Emirates Airline" // Variable declaration
	fmt.Print(airlineName)
}
