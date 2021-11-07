package gopixel

import (
	"errors"
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
	for i := 0; i < 10; i++ {
		resp, err := http.Get("https://" + url)

		if err != nil {
			return nil, nil
		}

		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)

		return body, err
	}
	return make([]byte, 0), errors.New("server side error")
}
