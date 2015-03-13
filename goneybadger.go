package goneybadger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const apiURL = "https://api.honeybadger.io/v1/notices"

type client interface {
	Do(*http.Request) (*http.Response, error)
}

// Honeybadger stores your Honebadger client information.
type Honeybadger struct {
	apiKey        string
	hostname, env string
	httpClient    client
}

// New creates a new Honeybadger client
func New(apiKey, env string) *Honeybadger {
	return NewWithTimeout(apiKey, env, 5*time.Second)
}

// NewWithTimeout creates a new Honeybadger client specifying the timeout for
// the http calls.
func NewWithTimeout(apiKey, env string, timeout time.Duration) *Honeybadger {
	hostname, _ := os.Hostname()

	return &Honeybadger{
		apiKey:     apiKey,
		hostname:   hostname,
		env:        env,
		httpClient: &http.Client{Timeout: timeout},
	}
}

// Notify POSTs a message to Honeybadger with the needed payload to store the
// `message` error.
func (h *Honeybadger) Notify(message string) error {
	payload, err := json.Marshal(NewPayload(h.hostname, h.env, message))
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	for k, v := range map[string]string{
		"X-API-Key":    h.apiKey,
		"Content-Type": "application/json",
		"Accept":       "application/json",
	} {
		req.Header.Set(k, v)
	}

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf(
			"the API returned a %d instead an expected %d",
			resp.StatusCode,
			http.StatusCreated)
	}

	return nil
}
