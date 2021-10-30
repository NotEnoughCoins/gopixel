package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) SkyblockAuctionByPlayer(player string) (structs.SkyblockAuctions, error) {
	var auctions structs.SkyblockAuctions

	uuid, err := Uuid(player)
	if err != nil {
		return auctions, err
	}

	data, err := get("api.hypixel.net/skyblock/auction?player=" + uuid + "&key=" + client.Key)
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, &auctions)

	return auctions, err
}

func (client *Client) SkyblockAuctionByUuid(uuid string) (structs.SkyblockAuctions, error) {
	var auctions structs.SkyblockAuctions

	data, err := get("api.hypixel.net/skyblock/auction?uuid=" + uuid + "&key=" + client.Key)
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, &auctions)

	return auctions, err
}

func (client *Client) SkyblockAuctionByProfileUuid(uuid string) (structs.SkyblockAuctions, error) {
	var auctions structs.SkyblockAuctions

	data, err := get("api.hypixel.net/skyblock/auction?profile=" + uuid + "&key=" + client.Key)
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, &auctions)

	return auctions, err
}
