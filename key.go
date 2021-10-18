package hypixel

import (
	"encoding/json"

	structs "hypixel.go/structs"
)

func (client *Client) KeyInfo() (structs.Key, error) {
	var key structs.Key

	data, err := get("api.hypixel.net/key?key=" + client.key)
	if err != nil {
		return key, err
	}

	err = json.Unmarshal(data, &key)

	return key, err
}
