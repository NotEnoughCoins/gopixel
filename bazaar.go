package hypixel

import (
	"encoding/json"
	"errors"

	structs "hypixel.go/structs"
)

func (client *Client) BazaarData() (structs.Bazaar, error) {
	var bazaar structs.Bazaar

	data, err := get("api.hypixel.net/skyblock/bazaar?key=" + client.key)
	if err != nil {
		return bazaar, err
	}

	err = json.Unmarshal(data, &bazaar)

	if !bazaar.Success {
		err = errors.New(bazaar.Cause)
	}

	return bazaar, err
}
