package main

import (
	"log"
	"net/http"
)

func startServer() {
	mux := http.DefaultServeMux

	// Route for the search form page
	mux.HandleFunc("/", searchFormHandler)

	// Route for submitting the search form
	mux.HandleFunc("/search", searchResultsHandler)

	// Route for the booking confirmation page
	mux.HandleFunc("/booking", bookingConfirmationHandler)

	// Route for the booking receipt page
	mux.HandleFunc("/receipt", bookingReceiptHandler)

	// Start the server
	go func() {
		log.Fatal(http.ListenAndServe(":8080", mux))
	}()
}
