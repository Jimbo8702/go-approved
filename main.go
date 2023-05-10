package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func seedBooking(store Storage, startDate string, length int) *Booking {
	var dates []string
	userId := uuid.NewString()
	stripe := uuid.NewString()
	product := uuid.NewString()
	booking, err := NewBooking(startDate, userId, stripe, product, length, dates)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateBooking(booking); err != nil {
		log.Fatal(err)
	}

	fmt.Println("New Booking =>", booking.ID )

	return booking
} 

func seedBookings (s Storage) {
	seedBooking(s, "20-15-22", 4)
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	
	if *seed {
		fmt.Println("Seeding the DB")
		seedBookings(store)
	}
	
	server := NewAPIServer(":3000", store)
	server.Run()
}