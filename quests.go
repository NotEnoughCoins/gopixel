package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) Quests() (structs.Quests, error) {
	var quests structs.Quests

	data, err := get("api.hypixel.net/resources/quests?key=" + client.key)
	if err != nil {
		return quests, err
	}

	err = json.Unmarshal(data, &quests)

	return quests, err
}
