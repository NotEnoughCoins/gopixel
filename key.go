package gopixel

import (
	"encoding/json"
	"errors"

	structs "gopixel/structs"
)

func (client *Client) KeyInfo() (structs.Key, error) {
	var key structs.Key

	data, err := get("api.hypixel.net/key?key=" + client.key)
	if err != nil {
		return key, err
	}

	err = json.Unmarshal(data, &key)

	if !key.Success {
		err = errors.New(key.Cause)
	}

	return key, err
}
