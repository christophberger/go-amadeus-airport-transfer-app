package main

import "net/http"

func bookingConfirmationHandler(w http.ResponseWriter, r *http.Request) {
	// Make the POST call to the API endpoint
	// Process the API response
	// Render the booking confirmation template
	// ...
}

type Booking struct {
	Data struct {
		Note       string `json:"note"`
		Passengers []struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Title     string `json:"title"`
			Contacts  struct {
				PhoneNumber string `json:"phoneNumber"`
				Email       string `json:"email"`
			} `json:"contacts"`
			BillingAddress struct {
				Line        string `json:"line"`
				Zip         string `json:"zip"`
				CountryCode string `json:"countryCode"`
				CityName    string `json:"cityName"`
			} `json:"billingAddress"`
		} `json:"passengers"`
		Agency struct {
			Contacts []struct {
				Email struct {
					Address string `json:"address"`
				} `json:"email"`
			} `json:"contacts"`
		} `json:"agency"`
		Payment struct {
			MethodOfPayment string `json:"methodOfPayment"`
			CreditCard      struct {
				Number     string `json:"number"`
				HolderName string `json:"holderName"`
				VendorCode string `json:"vendorCode"`
				ExpiryDate string `json:"expiryDate"`
				Cvv        string `json:"cvv"`
			} `json:"creditCard"`
		} `json:"payment"`
		ExtraServices []struct {
			Code   string `json:"code"`
			ItemID string `json:"itemId"`
		} `json:"extraServices"`
		Equipment []struct {
			Code string `json:"code"`
		} `json:"equipment"`
		Corporation struct {
			Address struct {
				Line        string `json:"line"`
				Zip         string `json:"zip"`
				CountryCode string `json:"countryCode"`
				CityName    string `json:"cityName"`
			} `json:"address"`
			Info struct {
				Au string `json:"AU"`
				Ce string `json:"CE"`
			} `json:"info"`
		} `json:"corporation"`
		StartConnectedSegment struct {
			TransportationType   string `json:"transportationType"`
			TransportationNumber string `json:"transportationNumber"`
			Departure            struct {
				UicCode       string `json:"uicCode"`
				IataCode      string `json:"iataCode"`
				LocalDateTime string `json:"localDateTime"`
			} `json:"departure"`
			Arrival struct {
				UicCode       string `json:"uicCode"`
				IataCode      string `json:"iataCode"`
				LocalDateTime string `json:"localDateTime"`
			} `json:"arrival"`
		} `json:"startConnectedSegment"`
		EndConnectedSegment struct {
			TransportationType   string `json:"transportationType"`
			TransportationNumber string `json:"transportationNumber"`
			Departure            struct {
				UicCode       string `json:"uicCode"`
				IataCode      string `json:"iataCode"`
				LocalDateTime string `json:"localDateTime"`
			} `json:"departure"`
			Arrival struct {
				UicCode       string `json:"uicCode"`
				IataCode      string `json:"iataCode"`
				LocalDateTime string `json:"localDateTime"`
			} `json:"arrival"`
		} `json:"endConnectedSegment"`
	} `json:"data"`
}

type BookingResponse struct {
	Data struct {
		Type       string `json:"type"`
		Reference  string `json:"reference"`
		ID         string `json:"id"`
		Passengers []struct {
			Type      string `json:"type"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Title     string `json:"title"`
			Contacts  struct {
				Email       string `json:"email"`
				PhoneNumber string `json:"phoneNumber"`
			} `json:"contacts"`
			BillingAddress struct {
				Line        string `json:"line"`
				Zip         string `json:"zip"`
				CountryCode string `json:"countryCode"`
				CityName    string `json:"cityName"`
			} `json:"billingAddress"`
		} `json:"passengers"`
		Transfers []struct {
			Status            string `json:"status"`
			ConfirmNbr        string `json:"confirmNbr"`
			Note              string `json:"note"`
			MethodOfPayment   string `json:"methodOfPayment"`
			OfferID           string `json:"offerId"`
			TransferType      string `json:"transferType"`
			CancellationRules []struct {
				FeeType         string `json:"feeType"`
				FeeValue        string `json:"feeValue"`
				CurrencyCode    string `json:"currencyCode"`
				MetricMax       string `json:"metricMax"`
				MetricType      string `json:"metricType"`
				MetricMin       string `json:"metricMin"`
				RuleDescription string `json:"ruleDescription"`
			} `json:"cancellationRules"`
			Start struct {
				DateTime     string `json:"dateTime"`
				LocationCode string `json:"locationCode"`
			} `json:"start"`
			End struct {
				DateTime string `json:"dateTime"`
				Address  struct {
					Line        string  `json:"line"`
					Zip         string  `json:"zip"`
					CountryCode string  `json:"countryCode"`
					CityName    string  `json:"cityName"`
					Latitude    float64 `json:"latitude"`
					Longitude   float64 `json:"longitude"`
				} `json:"address"`
				Name string `json:"name"`
			} `json:"end"`
			Vehicle struct {
				Code        string `json:"code"`
				Category    string `json:"category"`
				Description string `json:"description"`
				Baggages    []struct {
					Count int    `json:"count"`
					Size  string `json:"size"`
				} `json:"baggages"`
				Seats []struct {
					Count int `json:"count"`
				} `json:"seats"`
				ImageURL string `json:"imageURL"`
			} `json:"vehicle"`
			ServiceProvider struct {
				Code     string `json:"code"`
				Name     string `json:"name"`
				TermsURL string `json:"termsUrl"`
				LogoURL  string `json:"logoUrl"`
			} `json:"serviceProvider"`
			Quotation struct {
				MonetaryAmount string `json:"monetaryAmount"`
				CurrencyCode   string `json:"currencyCode"`
				Taxes          []struct {
					MonetaryAmount string `json:"monetaryAmount"`
				} `json:"taxes"`
				IsEstimated bool `json:"isEstimated"`
				TotalFees   struct {
					MonetaryAmount string `json:"monetaryAmount"`
				} `json:"totalFees"`
				TotalTaxes struct {
					MonetaryAmount string `json:"monetaryAmount"`
				} `json:"totalTaxes"`
			} `json:"quotation"`
			Converted struct {
				MonetaryAmount string `json:"monetaryAmount"`
				CurrencyCode   string `json:"currencyCode"`
				Taxes          []struct {
					MonetaryAmount string `json:"monetaryAmount"`
				} `json:"taxes"`
				IsEstimated bool `json:"isEstimated"`
				TotalFees   struct {
					MonetaryAmount string `json:"monetaryAmount"`
				} `json:"totalFees"`
				TotalTaxes struct {
					MonetaryAmount string `json:"monetaryAmount"`
				} `json:"totalTaxes"`
			} `json:"converted"`
		} `json:"transfers"`
	} `json:"data"`
}

type BookingErrorResponse struct {
	Errors []struct {
		Code   int    `json:"code"`
		Detail string `json:"detail"`
	} `json:"errors"`
}
