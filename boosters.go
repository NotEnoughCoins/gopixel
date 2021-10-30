package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) Boosters() (structs.Boosters, error) {
	var boosters structs.Boosters

	data, err := get("api.hypixel.net/boosters?key=" + client.key)
	if err != nil {
		return boosters, err
	}

	err = json.Unmarshal(data, &boosters)

	return boosters, err
}
