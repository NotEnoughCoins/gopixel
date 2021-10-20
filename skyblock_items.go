package gopixel

import (
	"encoding/json"

	structs "gopixel/structs"
)

func (client *Client) SkyblockItems() (structs.SkyblockItems, error) {
	var skyblockItems structs.SkyblockItems

	data, err := get("api.hypixel.net/resources/skyblock/items?key=" + client.key)
	if err != nil {
		return skyblockItems, err
	}

	err = json.Unmarshal(data, &skyblockItems)

	return skyblockItems, err
}