package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) Games() (structs.Games, error) {
	var games structs.Games

	data, err := get("api.hypixel.net/resources/games?key=" + client.key)
	if err != nil {
		return games, err
	}

	err = json.Unmarshal(data, &games)

	return games, err
}
