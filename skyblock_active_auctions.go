package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) SkyblockActiveAuctions() (structs.SkyblockActiveAuctions, error) {
	var auctions structs.SkyblockActiveAuctions

	data, err := get("api.hypixel.net/skyblock/auctions?key=" + client.Key)
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, &auctions)

	return auctions, err
}
