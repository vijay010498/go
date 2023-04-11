package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	vijay := Passenger{"vijay", 1, false}
	fmt.Println(vijay)

	var (
		pablo = Passenger{Name: "pablo", TicketNumber: 2}
		neha  = Passenger{Name: "neha", TicketNumber: 3}
	)

	fmt.Println(pablo, neha)

	var heidi Passenger
	heidi.Name = "heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	vijay.Boarded = true
	pablo.Boarded = true
	heidi.Boarded = true

	bus := Bus{heidi}
	fmt.Println(bus.FrontSeat.Name)
}
