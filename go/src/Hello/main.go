package main
import "fmt"

// func main(){
// 	fmt.Println("Hello 90 days of devops")
// }

// func main(){
// 	var challenge = "#90daysofdevops!"
// 	const totaldays = 90
// 	var dayscomplete = 11
// 	fmt.Println("Welcome to", challenge, "")
// 	fmt.Println("This is a", totaldays, "day challenge")
// 	fmt.Println("You have completed", dayscomplete, "days")
// }

func main(){
	const DaysTotal int = 90
	var remainingDays uint = 90
	challenge := "#90daysofdevops"
	
	fmt.Printf("Welcome to the %v challenge. \nThis challenege consists of %v days\n", challenge, DaysTotal)

	var TwitterName string
	var DaysCompleted uint

	fmt.Println("Enter your twitter handle:")
	fmt.Scanln(&TwitterName)

	fmt.Println("How many days have you completed?")
	fmt.Scanln(&DaysCompleted)


	remainingDays = remainingDays - DaysCompleted

	fmt.Printf("You have %v days more to go! Lets go strong %v YOU GOT THIS!!!\n" , remainingDays, TwitterName)


	// the below line gives us the memory address of the pointer
	fmt.Println(&challenge)
}