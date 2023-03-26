package main //every statement in go written in packages

import "fmt" //In go if we need do use built in packages we need to import explicityly

func main() { // program starts form main

	airlineName := "Emirates" //This is same as var airlineName = "Emirates"
	//But it simplifies that and it doesn't add new functionality
	fmt.Print(airlineName)

}
