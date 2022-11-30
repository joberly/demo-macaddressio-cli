package api

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

// Client is a macaddress.io API client.
type Client struct {
	client http.Client
	apiKey string
}

var (
	// Errors

	ErrParameters      = errors.New("invalid parameter(s)")
	ErrInvalidAPIKey   = errors.New("access restricted, invalid API key")
	ErrCredits         = errors.New("access restricted, check credits")
	ErrInvalidMAC      = errors.New("invalid MAC or OUI received")
	ErrTooManyRequests = errors.New("too many requests")
	ErrInternal        = errors.New("internal server error")
)

const baseURL = "https://api.macaddress.io/v1?output=json&%s"

// New creates a new macaddress.io API client.
func New(options ...func(*Client)) *Client {
	c := &Client{}
	for _, opt := range options {
		opt(c)
	}
	return c
}

// WithAPIKey sets the API key for a new Client.
func WithAPIKey(apiKey string) func(*Client) {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// Request makes a generic API query request and return the response body or an error.
func (c *Client) request(query string) (*bytes.Buffer, error) {
	// Create request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(baseURL, query), http.NoBody)

	// Return if request is invalid
	if err != nil {
		return nil, err
	}

	// Add auth header with API key
	req.Header.Add("X-Authentication-Token", c.apiKey)

	// Make API request
	resp, err := c.client.Do(req)

	// Return error if request fails to connect, etc.
	if err != nil {
		return nil, err
	}

	// Check response status, return error if unsuccessful
	if resp.StatusCode != http.StatusOK {
		var err error
		switch resp.StatusCode {
		case http.StatusBadRequest:
			err = ErrParameters
		case http.StatusUnauthorized:
			err = ErrInvalidAPIKey
		case http.StatusPaymentRequired:
			err = ErrCredits
		case http.StatusUnprocessableEntity:
			err = ErrInvalidMAC
		case http.StatusTooManyRequests:
			err = ErrTooManyRequests
		case http.StatusInternalServerError:
			err = ErrInternal
		default:
			err = fmt.Errorf("unknown API status code %d", resp.StatusCode)
		}
		return nil, err
	}

	// Successful API call, read the response body into a buffer
	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	return buf, nil
}
