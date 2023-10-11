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

// token returns the current access token. If none exists yet, or if the existing one has expired, it fetches a new one from the Amadeus authorization API. If fetching fails, token returns an error.
func (c *Client) token() (string, error) {
	if len(c.tokenErrorCh) > 0 { // works, because tokenErrorCh is buffered
		e := <-c.tokenErrorCh
		return "", e
	}
	t := <-c.accessTokenCh
	return t, nil
}

// startTokenFetcher starts a goroutine that fetches a new access token from the Amadeus authorization API if there is none yet, or if the current one expires. It returns channels for returning the current token, or an error if the token could not be fetched.
func (c *Client) startTokenFetcher() {
	go func() {
		var token string
		var expiration int
		var err error

		c.accessTokenCh = make(chan string)
		c.tokenErrorCh = make(chan error, 1)

		expired := time.After(0) // the timer fires immediately

		for {
			select {
			case <-expired:
				// fetch a new token from the API
				token, expiration, err = authorize(c.baseURL)
				if err != nil {
					c.tokenErrorCh <- err
					return
				}
				// set the timer to fire 60 seconds before the token expires
				expired = time.After(time.Duration(expiration-60) * time.Second)

			case c.accessTokenCh <- token:
				// someone has read the token, nothing to do
				// the next iteration will send the token to the channel again
			}
		}
	}()
}

// authorize reads client ID and secret from the environment variables and updates the access token and its lifespan (in seconds) from the Amadeus authorization API.
func authorize(baseURL string) (token string, lifespan int, err error) {

	url := baseURL + "/security/oauth2/token"
	method := "POST"

	id := os.Getenv("AMADEUS_CLIENT_ID")
	secret := os.Getenv("AMADEUS_CLIENT_SECRET")

	if id == "" || secret == "" {
		return "", 0, fmt.Errorf("authorize: missing client ID or secret (check the environment variables AMADEUS_CLIENT_ID and AMADEUS_CLIENT_SECRET)")
	}

	payload := strings.NewReader("client_id=" + id +
		"&client_secret=" + secret +
		"&grant_type=client_credentials")

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", 0, fmt.Errorf("authorize: http.NewRequest: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("authorize: client.Do: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", 0, fmt.Errorf("authorize: io.ReadAll: %w", err)
	}

	// Unmarshal the response. AuthResponse is a struct that combines
	// the responses for the successful case and for the error case.
	// Unmarshal() does not complain if the JSON does not fill the
	// struct completely, which we use here to simplify the unmarshalling.
	var authResponse AuthResponse
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return "", 0, fmt.Errorf("authorize: json.Unmarshal: %w", err)
	}
	if authResponse.Error != "" {
		return "", 0, fmt.Errorf("authorize: %w (%s: %s (error: %s, code: %d)",
			err,
			authResponse.Title,
			authResponse.ErrorDescription,
			authResponse.Error,
			authResponse.Code)
	}

	return authResponse.AccessToken,
		authResponse.ExpiresIn,
		nil

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
