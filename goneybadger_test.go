package goneybadger

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	assert "github.com/pilu/miniassert"
)

type MockClient struct {
	req *http.Request

	resp *http.Response
	err  error
}

func (c *MockClient) Do(req *http.Request) (*http.Response, error) {
	c.req = req
	return c.resp, c.err
}

func NewMockClient(statusCode int, err error) *MockClient {
	return &MockClient{
		resp: &http.Response{
			StatusCode: statusCode,
			Body:       &MockBody{},
		},
		err: err,
	}
}

type MockBody struct{}

func (b *MockBody) Read(p []byte) (n int, err error) { return 0, nil }
func (b *MockBody) Close() error                     { return nil }

// TestNotify201 will check that the payload is generated and POSTed properly.
func TestNotify201(t *testing.T) {
	expectedMessage := "Hey dude!"

	c := NewMockClient(201, nil)
	gb := Honeybadger{httpClient: c}
	err := gb.Notify(expectedMessage)

	receivedPayload := Payload{}
	err = json.NewDecoder(c.req.Body).Decode(&receivedPayload)
	assert.Nil(t, err)

	assert.Nil(t, err)
	assert.Equal(t, receivedPayload.Error.Message, expectedMessage)
}

// TestNotify201 will asure that if the Honeybadger API returns us something
// different from a 201 status code it will be considered as an error.
func TestNotifyNon201(t *testing.T) {
	expectedStatus := 444

	gb := Honeybadger{httpClient: NewMockClient(444, nil)}
	err := gb.Notify("")

	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf("the API returned a %d instead an expected 201", expectedStatus),
	)
}

// TestNotifyError will test that in case of error DO'ing the request we will
// get it back.
func TestNotifyError(t *testing.T) {
	gb := Honeybadger{
		httpClient: NewMockClient(666, errors.New("It's failing, awesome!")),
	}
	err := gb.Notify("")

	assert.NotNil(t, err)
}

// TestCreationOfClient will check that the attributes of the struct are
// initialized properly.
func TestCreationOfClient(t *testing.T) {
	expectedApiKey := "apiKey"
	expectedEnv := "env"

	gb := New(expectedApiKey, expectedEnv)

	assert.Equal(t, gb.apiKey, expectedApiKey)
	assert.Equal(t, gb.env, expectedEnv)
	assert.NotEqual(t, gb.hostname, "")
	assert.NotNil(t, gb.httpClient)
}
