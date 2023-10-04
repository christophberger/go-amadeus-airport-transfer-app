package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"
)

func (a *Amadeus) SearchHandler(w http.ResponseWriter, r *http.Request) {

	// Parse the query parameters from the request URL
	queryParams := r.URL.Query()

	// Retrieve the query parameters and save them to local variables
	searchParams := SearchParameters{
		StartAddressLine: queryParams.Get("streetAddress") + " " + queryParams.Get("houseNumber"),
		StartCityName:    queryParams.Get("city"),
		StartZipCode:     queryParams.Get("zipCode"),
		StartCountryCode: queryParams.Get("countryCode"),
		StartGeoCode:     queryParams.Get("latitude") + "," + queryParams.Get("longitude"),
		EndLocationCode:  queryParams.Get("endLocationCode"),
	}

	// Check if any parameter (except houseNumber) is empty
	if searchParams.EndLocationCode == "" ||
		searchParams.StartAddressLine == " " ||
		searchParams.StartCityName == "" ||
		searchParams.StartZipCode == "" ||
		searchParams.StartCountryCode == "" ||
		searchParams.StartGeoCode == "" {
		http.Error(w, "Address data is incomplete", http.StatusBadRequest)
		return
	}

	response, err := a.Search(searchParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the template
	tmpl, err := template.New("offerList").Parse(offerListTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template to the ResponseWriter
	err = tmpl.Execute(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *Amadeus) Search(p SearchParameters) (SearchResponse, error) {
	url := a.baseURL + "/shopping/transfer-offers"
	method := "POST"

	params, err := json.Marshal(p)
	if err != nil {
		return SearchResponse{}, err
	}

	payload := bytes.NewReader(params)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return SearchResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+a.token)

	res, err := client.Do(req)
	if err != nil {
		// See if the error is due to missing authentication
		if res.Body != nil {
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return SearchResponse{}, err
			}
			errResp := SearchErrorResponse{}
			err = json.Unmarshal(body, &errResp)
			if err != nil {
				return SearchResponse{}, err
			}
			if errResp.Errors[0].Code == 38192 {
				// Access token expired: re-authenticate
				err = a.authenticate()
				if err != nil {
					return SearchResponse{}, err
				}
				res, err = client.Do(req)
				if err != nil {
					return SearchResponse{}, err
				}
			}
		}
		return SearchResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return SearchResponse{}, err
	}

	result := SearchResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return SearchResponse{}, err
	}
	return result, nil
}

const offerListTemplate = `
    <!DOCTYPE html>
    <html>
        <head>
            <title>My Page</title>
        </head>
        <body>
			<table>
				{{range .Data}}
				<tr>
					<td>Transfer Type</td>
					<td>{{.TransferType}}</td>
				</tr>
				<tr>
					<td>Start Time</td>
					<td>{{.Start.DateTime}}</td>
				</tr>
				<tr>
					<td>Arrival Time</td>
					<td>{{.End.DateTime}}</td>
				</tr>
				<tr>
					<td>Service Provider</td>
					<td>{{.ServiceProvider.Name}}</td>
				</tr>
				<tr>
					<td>Estimated Cost</td>
					<td>{{.Quotation.CurrencyCode}} {{.Quotation.MonetaryAmount}}</td>
				</tr>
					<td><button id="book">Book this transfer</button></td>
					<td></td>
				{{end}}
			</table>
        </body>
    </html>

`

func generateOfferList(response SearchResponse) (string, error) {
	// Parse the template
	tmpl, err := template.New("offerList").Parse(offerListTemplate)
	if err != nil {
		return "", fmt.Errorf("generateOfferList: parse template: %w", err)
	}

	// Render the template to a buffer
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, response)
	if err != nil {
		return "", fmt.Errorf("generateOfferList: execute template: %w", err)
	}

	return buf.String(), nil
}
