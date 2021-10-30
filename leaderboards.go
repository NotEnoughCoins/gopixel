package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) Leaderboards() (structs.Leaderboards, error) {
	var leaderboards structs.Leaderboards

	data, err := get("api.hypixel.net/leaderboards?key=" + client.key)
	if err != nil {
		return leaderboards, err
	}

	err = json.Unmarshal(data, &leaderboards)

	return leaderboards, err
}
