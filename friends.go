package hypixel

import (
	"encoding/json"
	"errors"

	structs "hypixel.go/structs"
)

func (client *Client) Friends(name string) (structs.Friends, error) {
	var Friends structs.Friends

	uuid, err := Uuid(name)

	if err != nil {
		return Friends, err
	}
	data, err := get("api.hypixel.net/friends?uuid=" + uuid + "&key=" + client.key)

	if err != nil {
		return Friends, err
	}

	err = json.Unmarshal(data, &Friends)

	if !Friends.Success {
		err = errors.New(Friends.Cause)
	}

	return Friends, err
}
