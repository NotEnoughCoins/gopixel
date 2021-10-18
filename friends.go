package hypixel

import (
	"encoding/json"
	"errors"

	structs "hypixel.go/structs"
)

func (client *Client) Friends(name string) (structs.Friends, error) {
	var friends structs.Friends

	uuid, err := Uuid(name)

	if err != nil {
		return friends, err
	}
	data, err := get("api.hypixel.net/friends?uuid=" + uuid + "&key=" + client.key)

	if err != nil {
		return friends, err
	}

	err = json.Unmarshal(data, &friends)

	if !friends.Success {
		err = errors.New(friends.Cause)
	}

	return friends, err
}
