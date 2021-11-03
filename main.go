package gopixel

import (
	"io/ioutil"

	"net/http"
)

type Client struct {
	Key string
}

// Returns a client object
func NewClient(key string) *Client {
	return &Client{Key: key}
}

// Internal function to handle http GET requests
func get(url string) ([]byte, error) {
	resp, err := http.Get("https://" + url)

	if err != nil {
		return nil, nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
