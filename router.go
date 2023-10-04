package main

import (
	"log"
	"net/http"
)

func startServer(a Amadeus) {
	mux := http.NewServeMux()

	// Route for the search form page
	mux.HandleFunc("/", a.HomeHandler)

	// Route for submitting the search
	mux.HandleFunc("/search", a.SearchHandler)

	// Route for the booking handler
	mux.HandleFunc("/booking", a.BookingHandler)

	// Route for the booking confirmation page
	mux.HandleFunc("/bookingconfirmation", a.BookingConfirmationHandler)

	// Start the server
	go func() {
		log.Fatal(http.ListenAndServe("localhost:8020", mux))
	}()
}
