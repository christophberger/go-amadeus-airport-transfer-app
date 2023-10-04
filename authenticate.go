package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

func (a *Amadeus) authenticate() error {

	url := a.baseURL + "/security/oauth2/token"
	method := "POST"

	payload := strings.NewReader("client_id=" + a.clientID +
		"&client_secret=" + a.secret +
		"&grant_type=client_credentials")

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	authResponse := AuthResponse{}
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return err
	}

	a.token = authResponse.AccessToken

	return nil
}

type AuthResponse struct {
	Type            string `json:"type"`
	Username        string `json:"username"`
	ApplicationName string `json:"application_name"`
	ClientID        string `json:"client_id"`
	TokenType       string `json:"token_type"`
	AccessToken     string `json:"access_token"`
	ExpiresIn       int    `json:"expires_in"`
	State           string `json:"state"`
	Scope           string `json:"scope"`
}

type AuthErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Code             int    `json:"code"`
	Title            string `json:"title"`
}
