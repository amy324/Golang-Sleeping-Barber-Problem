package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables
var seatCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// random number generator
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	// print welcome message
	color.Green("The Sleeping Barber Problem")
	color.Green("---------------------------")

	// create channels if needed
	clientChan := make(chan string, seatCapacity)
	doneChan := make(chan bool)

	// create data structure to represent barbershop
	shop := BarberShop{
		ShopCapacity:    seatCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Cyan("The shop is open for today")

	// add barbers
	shop.addBarber("Frank")
	shop.addBarber("Gerard")
	shop.addBarber("Milton")
	shop.addBarber("Patrick")
	shop.addBarber("Jim")

	// start the barbershop (go routine)
	shopClosing := make(chan bool)
	closed := make(chan bool)
	
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForToday()
		closed <- true
	}()

	// add clients
	i := 1
	
	go func(){
		for {
			//get a random number with average arrival rate
			randomMilliseconds := rng.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing: 
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}

		}
	}()
	// run application until shop closed
	<-closed
}
