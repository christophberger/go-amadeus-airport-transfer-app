package amadeus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// token returns the current access token. If none exists yet, or if the existing one has expired, it fetches a new one from the Amadeus authentication API.
func (c *Client) token() string {
	if c.accessToken == "" || time.Now().After(c.expiration) {
		// fetch new token from the API
		err := c.authenticate()
		if err != nil {
			return err.Error()
		}
	}
	return c.accessToken
}

// authenticate reads client ID and secret from the environment variables and updates the access tokend and expiration time from the Amadeus authentication API.
func (c *Client) authenticate() error {

	url := c.baseURL + "/security/oauth2/token"
	method := "POST"

	id := os.Getenv("AMADEUS_CLIENT_ID")
	secret := os.Getenv("AMADEUS_CLIENT_SECRET")

	if id == "" || secret == "" {
		return fmt.Errorf("authenticate: missing client ID or secret (check the environment variables AMADEUS_CLIENT_ID and AMADEUS_CLIENT_SECRET)")
	}

	payload := strings.NewReader("client_id=" + id +
		"&client_secret=" + secret +
		"&grant_type=client_credentials")

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return fmt.Errorf("authenticate: http.NewRequest: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("authenticate: client.Do: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("authenticate: io.ReadAll: %w", err)
	}

	// Unmarshal the response. AuthResponse is a struct that combines
	// the responses for the successful case and for the error case.
	// Unmarshal() does not complain if the JSON does not fill the
	// struct completely, which we use here to simplify the unmarshalling.
	var authResponse AuthResponse
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return fmt.Errorf("authenticate: json.Unmarshal: %w", err)
	}
	if authResponse.Error != "" {
		return fmt.Errorf("authenticate: %w (%s: %s (error: %s, code: %d)",
			err,
			authResponse.Title,
			authResponse.ErrorDescription,
			authResponse.Error,
			authResponse.Code)
	}

	c.accessToken = authResponse.AccessToken
	c.expiration = time.Now().Add(time.Duration(authResponse.ExpiresIn) * time.Second)

	return nil
}

type AuthResponse struct {
	AuthSuccessResponse
	AuthErrorResponse
}

type AuthSuccessResponse struct {
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
