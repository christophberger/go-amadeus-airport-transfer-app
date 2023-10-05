package main

import (
	"html/template"
	"net/http"

	"appliedgo.net/what"
)

func (a *app) BookingHandler(w http.ResponseWriter, r *http.Request) {
	// Get the offer ID from the query string
	offerID := r.URL.Query().Get("offerId")

	// Call the Amadeus Transfer Booking API
	// (see internal/amadeus/book.go)
	response, err := a.ac.Book(offerID)
	if err != nil {
		// Render the erorr nicely
		template.Must(template.New("bookingError").Parse(bookingErrorTemplate)).Execute(w, err)
		return
	}

	what.Happens("Response: %s", response)

	// Render the booking receipt template
	tmpl, err := template.New("bookingReceipt").Parse(bookingConfirmationTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// booking confirmation template
// include detail from the BookingResponse
var bookingConfirmationTemplate = `<html>
<body>
	<h1>Booking Confirmation</h1>
	<p>Reference: {{.Data.Reference}}</p>
	<p>Booking ID: {{.Data.ID}}</p>
	<p>Thank you for your booking!</p>
	<p><a href="/">New search</a></p>
</body>
</html>`

var bookingErrorTemplate = `<html>
<body>
	<h1>Booking Error</h1>
	<p>{{.}}</p>
	<p><a href="/">New search</a></p>
</body>`
