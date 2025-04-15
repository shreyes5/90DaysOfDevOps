package main

import "fmt"

func main() {
	challenge := "#90DaysOfDevOps"
	const daystotal int = 90
	var remainingDays uint = 90

	// fmt.Println("Welcome to", challenge, "")
	// fmt.Println("This is a", daystotal, "challenge and you have completed", dayscompleted, "days")
	// fmt.Println("Keep going, you are doing great!")

	fmt.Printf("Welcome to %v days, this challenge consists of %v days\n", challenge, daystotal)

	var name string
	var daysDone uint

	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	fmt.Println("How many days have you completed?:")
	fmt.Scanln(&daysDone)

	remainingDays = remainingDays - daysDone

	fmt.Printf("This is a %v days challenge and you have completed %v days\n", daystotal, daysDone)
	fmt.Printf("You have %v days remaining\n", remainingDays)
	fmt.Println("Lessgo")

	fmt.Println(challenge)
	fmt.Println(&challenge) //& points to the memory address of the variable challenge
}
