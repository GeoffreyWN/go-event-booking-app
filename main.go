package main

import (
	"booking-app/helpers"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

const eventTickets = 50

var eventName = "Go Event" // create a variable and assign it a value
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

func greetPeople() {
	fmt.Printf("Welcome to the Biggest %v of the year \n", eventName)
	fmt.Printf("Total number of tickets %v and remaining tickets up for grabs: %v \n", eventTickets, remainingTickets)
	fmt.Println("Let's get you a ticket Champ!")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your First name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email has been to your email at %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for the %v \n", remainingTickets, eventName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// simulate an async operation
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("**********")
	fmt.Printf("Sending ticket: \n %v \n to email address %v \n", ticket, email)
	fmt.Println("**********")

	// Decrements the waitgroup counter by 1 thus indicating that the goroutine is finished
	wg.Done()

}

func main() {

	greetPeople()

	// for {

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidNoOfTickets := helpers.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidNoOfTickets && isValidName && isValidEmail {

		bookTicket(userTickets, firstName, lastName, email)

		// wait for the launched goroutine to finish
		wg.Add(1) // set number of goroutines to wait for

		// assign this task to another goroutine (green thread: abstraction of a actual thread) allowing for concurrency. This way the main goroutine continues to run
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()

		fmt.Println("Current bookings summary: ", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Event is fully booked!")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name is too short")
		}

		if !isValidEmail {
			fmt.Println("Email address does not contain @ sign")
		}

		if !isValidNoOfTickets {
			fmt.Println("Provided number of tickets is invalid")
		}
	}
	// }
	// wait for all the goroutines that were added to  i.e until wait group is 0
	wg.Wait()
}
