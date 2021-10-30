package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) SkyblockEndedAuctions() (structs.SkyblockEndedAuctions, error) {
	var auctions structs.SkyblockEndedAuctions

	data, err := get("api.hypixel.net/skyblock/auctions_ended?key=" + client.Key)
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, &auctions)

	return auctions, err
}
