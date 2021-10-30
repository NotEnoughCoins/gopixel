package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) SkyblockItems() (structs.SkyblockItems, error) {
	var skyblockItems structs.SkyblockItems

	data, err := get("api.hypixel.net/resources/skyblock/items?key=" + client.Key)
	if err != nil {
		return skyblockItems, err
	}

	err = json.Unmarshal(data, &skyblockItems)

	return skyblockItems, err
}
