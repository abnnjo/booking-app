package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	for {
		var firstName string
		var lastName string
		var userTickets uint
		var emailID string

		fmt.Println("Enter your First Name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email ID")
		fmt.Scan(&emailID)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailID)
			fmt.Printf("%v tickets were remaining for %v\n", remainingTickets, conferenceName)

			var firstNames []string
			for _, name := range bookings {
				var names = strings.Fields(name)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These first names of bookings %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break

			}

		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}
}
