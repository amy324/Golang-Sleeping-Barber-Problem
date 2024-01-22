package main

import (
	"github.com/fatih/color"
	"time"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isAsleep := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			//barber goes to sleep if there are no clients
			if len(shop.ClientsChan) == 0 {
				color.Yellow("%s does not have to cut anyone's hair right now, so he has a nap.", barber)
				isAsleep = true
			}

			client, shopOpen := <- shop.ClientsChan
			if shopOpen {
				if isAsleep {
					color.Yellow("%s wakes up %s.", client, barber)
					isAsleep = false
				}
				//barber cuts client's hair
				shop.cutHair(barber, client) 
			} else {
				//shop is closed, so barber needs to go home and go routine can close
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Cyan("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Cyan("%s has finished cutting %'s hair", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Magenta("%s is going home.", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForToday() {
	color.Magenta("Closing shop for today")
	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<- shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)

	color.Cyan("--------------------------------------------------------------")
	color.Cyan("The barbershop has closed for today and everyone has gone home")
}

func (shop *BarberShop) addClient(client string) {
	color.Cyan("*** %s has arrived.", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client: 
		color.Yellow("%s takes a seat in the wait")
		default:
			color.Magenta("The waiting room is full, so %s leaves.", client)
		}
	} else {
		color.Magenta("The shop has already closed, so %s leaves.", client)
	}
}
