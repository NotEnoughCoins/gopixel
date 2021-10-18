package hypixel

import (
	"encoding/json"

	structs "hypixel.go/structs"
)

func (client *Client) PlayerFriends(name string) (structs.Friends, error) {
	var PlayerFriends structs.Friends

	uuid, err := Uuid(name)

	if err != nil {
		return PlayerFriends, err
	}
	data, err := get("api.hypixel.net/friends?uuid=" + uuid + "&key=" + client.key)

	if err != nil {
		return PlayerFriends, err
	}

	err = json.Unmarshal(data, &PlayerFriends)

	return PlayerFriends, err
}
