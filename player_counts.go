package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) PlayerCounts() (structs.PlayerCounts, error) {
	var counts structs.PlayerCounts
	data, err := get("api.hypixel.net/counts?key=" + client.key)
	if err != nil {
		return counts, err
	}

	err = json.Unmarshal(data, &counts)

	return counts, err
}
