package gopixel

import (
	"errors"
	"github.com/NotEnoughCoins/gopixel/structs"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"

	"net/http"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type Client struct {
	Key          string
	Retries      int
	AuctionCache structs.SkyblockActiveAuctions
	HttpClient   http.Client
}

// NewClient Returns a client object
func NewClient(key string, retries int) *Client {
	return &Client{Key: key, Retries: retries, HttpClient: http.Client{}}
}

// Internal function to handle http GET requests
func (client *Client) get(url string) ([]byte, error) {
	for i := 0; i < client.Retries; i++ {
		resp, err := client.HttpClient.Get("https://" + url)

		if err != nil {
			return nil, nil
		}

		defer resp.Body.Close()

		if resp.StatusCode != 200 && resp.StatusCode != 404 && resp.StatusCode != 403 {
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)

		return body, err
	}
	return make([]byte, 0), errors.New("server side error")
}
