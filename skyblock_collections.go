package gopixel

import (
	"encoding/json"

	structs "gopixel/structs"
)

func (client *Client) SkyblockCollections() (structs.SkyblockCollections, error) {
	var skyblockCollections structs.SkyblockCollections

	data, err := get("api.hypixel.net/resources/skyblock/collections?key=" + client.key)
	if err != nil {
		return skyblockCollections, err
	}

	err = json.Unmarshal(data, &skyblockCollections)

	return skyblockCollections, err
}
