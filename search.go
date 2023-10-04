package main

import "net/http"

func searchFormHandler(w http.ResponseWriter, r *http.Request) {
	// Render the search form template
	// ...
}

type Search struct {
	StartLocationCode string `json:"startLocationCode,omitempty"`
	EndAddressLine    string `json:"endAddressLine,omitempty"`
	EndCityName       string `json:"endCityName,omitempty"`
	EndZipCode        string `json:"endZipCode,omitempty"`
	EndCountryCode    string `json:"endCountryCode,omitempty"`
	EndName           string `json:"endName,omitempty"`
	EndGeoCode        string `json:"endGeoCode,omitempty"`
	TransferType      string `json:"transferType,omitempty"`
	StartDateTime     string `json:"startDateTime,omitempty"`
	ProviderCodes     string `json:"providerCodes,omitempty"`
	Passengers        int    `json:"passengers,omitempty"`
	StopOvers         []struct {
		Duration       string `json:"duration,omitempty"`
		SequenceNumber int    `json:"sequenceNumber,omitempty"`
		AddressLine    string `json:"addressLine,omitempty"`
		CountryCode    string `json:"countryCode,omitempty"`
		CityName       string `json:"cityName,omitempty"`
		ZipCode        string `json:"zipCode,omitempty"`
		Name           string `json:"name,omitempty"`
		GeoCode        string `json:"geoCode,omitempty"`
		StateCode      string `json:"stateCode,omitempty"`
	} `json:"stopOvers,omitempty"`
	StartConnectedSegment struct {
		TransportationType   string `json:"transportationType,omitempty"`
		TransportationNumber string `json:"transportationNumber,omitempty"`
		Departure            struct {
			LocalDateTime string `json:"localDateTime,omitempty"`
			IataCode      string `json:"iataCode,omitempty"`
		} `json:"departure,omitempty"`
		Arrival struct {
			LocalDateTime string `json:"localDateTime,omitempty"`
			IataCode      string `json:"iataCode,omitempty"`
		} `json:"arrival,omitempty"`
	} `json:"startConnectedSegment,omitempty"`
	PassengerCharacteristics []struct {
		PassengerTypeCode string `json:"passengerTypeCode,omitempty"`
		Age               int    `json:"age,omitempty"`
	} `json:"passengerCharacteristics,omitempty"`
}

type SearchResponse struct {
	Data []struct {
		ID           string `json:"id"`
		Type         string `json:"type"`
		TransferType string `json:"transferType"`
		Start        struct {
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
			ImageURL    string `json:"imageURL"`
			Baggages    []struct {
				Count int    `json:"count"`
				Size  string `json:"size"`
			} `json:"baggages"`
			Seats []struct {
				Count int `json:"count"`
			} `json:"seats"`
		} `json:"vehicle"`
		ServiceProvider struct {
			Code     string   `json:"code"`
			Name     string   `json:"name"`
			TermsURL string   `json:"termsUrl"`
			LogoURL  string   `json:"logoUrl"`
			Settings []string `json:"settings"`
		} `json:"serviceProvider"`
		Quotation struct {
			MonetaryAmount string `json:"monetaryAmount"`
			CurrencyCode   string `json:"currencyCode"`
			Taxes          []struct {
				MonetaryAmount string `json:"monetaryAmount"`
			} `json:"taxes"`
			TotalTaxes struct {
				MonetaryAmount string `json:"monetaryAmount"`
			} `json:"totalTaxes"`
			Base struct {
				MonetaryAmount string `json:"monetaryAmount"`
			} `json:"base"`
			Discount struct {
				MonetaryAmount string `json:"monetaryAmount"`
			} `json:"discount"`
			TotalFees struct {
				MonetaryAmount string `json:"monetaryAmount"`
			} `json:"totalFees"`
		} `json:"quotation"`
		CancellationRules []struct {
			FeeType         string `json:"feeType"`
			FeeValue        string `json:"feeValue"`
			CurrencyCode    string `json:"currencyCode"`
			MetricType      string `json:"metricType"`
			MetricMin       string `json:"metricMin"`
			MetricMax       string `json:"metricMax"`
			RuleDescription string `json:"ruleDescription"`
		} `json:"cancellationRules"`
		MethodsOfPaymentAccepted []string `json:"methodsOfPaymentAccepted"`
		PassengerCharacteristics []struct {
			PassengerTypeCode string `json:"passengerTypeCode"`
			Age               int    `json:"age"`
		} `json:"passengerCharacteristics"`
		Converted struct {
			MonetaryAmount string `json:"monetaryAmount"`
			CurrencyCode   string `json:"currencyCode"`
			Taxes          []struct {
				MonetaryAmount string `json:"monetaryAmount"`
			} `json:"taxes"`
			TotalTaxes struct {
				MonetaryAmount string `json:"monetaryAmount"`
			} `json:"totalTaxes"`
			Base struct {
				MonetaryAmount string `json:"monetaryAmount"`
			} `json:"base"`
			Discount struct {
				MonetaryAmount string `json:"monetaryAmount"`
			} `json:"discount"`
			TotalFees struct {
				MonetaryAmount string `json:"monetaryAmount"`
			} `json:"totalFees"`
		} `json:"converted"`
	} `json:"data"`
}

type SearchErrorResponse struct {
	Errors []struct {
		Status int    `json:"status"`
		Code   int    `json:"code"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Source struct {
			Parameter string `json:"parameter"`
		} `json:"source"`
	} `json:"errors"`
}
