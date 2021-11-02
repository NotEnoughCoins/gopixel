package gopixel

import (
	"encoding/json"

	"errors"

	"github.com/comblock/gopixel/structs"
)

func (client *Client) Bazaar() (structs.Bazaar, error) {
	var bazaar structs.Bazaar

	data, err := get("api.hypixel.net/skyblock/bazaar?key=" + client.Key)
	if err != nil {
		return bazaar, err
	}

	err = json.Unmarshal(data, &bazaar)

	if !bazaar.Success {
		err = errors.New(bazaar.Cause)
	}

	return bazaar, err
}

func (client *Client) SkyblockActiveAuctions() (structs.SkyblockActiveAuctions, error) {
	var auctions structs.SkyblockActiveAuctions

	data, err := get("api.hypixel.net/skyblock/auctions?key=" + client.Key)
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, &auctions)

	return auctions, err
}

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

func (client *Client) SkyblockCollections() (structs.SkyblockCollections, error) {
	var skyblockCollections structs.SkyblockCollections

	data, err := get("api.hypixel.net/resources/skyblock/collections?key=" + client.Key)
	if err != nil {
		return skyblockCollections, err
	}

	err = json.Unmarshal(data, &skyblockCollections)

	return skyblockCollections, err
}

func (client *Client) SkyblockEndedAuctions() (structs.SkyblockEndedAuctions, error) {
	var auctions structs.SkyblockEndedAuctions

	data, err := get("api.hypixel.net/skyblock/auctions_ended?key=" + client.Key)
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, &auctions)

	return auctions, err
}

func (client *Client) SkyblockItems() (structs.SkyblockItems, error) {
	var skyblockItems structs.SkyblockItems

	data, err := get("api.hypixel.net/resources/skyblock/items?key=" + client.Key)
	if err != nil {
		return skyblockItems, err
	}

	err = json.Unmarshal(data, &skyblockItems)

	return skyblockItems, err
}

func (c *Client) SkyblockNews() (structs.SkyblockNews, error) {
	var news structs.SkyblockNews

	data, err := get("api.hypixel.net/resources/guilds/permissions?key=" + c.Key)
	if err != nil {
		return news, err
	}

	err = json.Unmarshal(data, &news)

	return news, err
}

func (client *Client) SkyblockProfile(profile string) (structs.SkyblockProfile, error) {
	var skyblockProfile structs.SkyblockProfile

	data, err := get("api.hypixel.net/skyblock/profile?profile=" + profile + "&key=" + client.Key)

	if err != nil {
		return skyblockProfile, err
	}

	err = json.Unmarshal(data, &skyblockProfile)

	if !skyblockProfile.Success {
		err = errors.New(skyblockProfile.Cause)
	}

	return skyblockProfile, err
}

func (client *Client) SkyblockProfiles(name string) (structs.SkyblockProfiles, error) {
	var skyblockProfile structs.SkyblockProfiles

	uuid, err := Uuid(name)

	if err != nil {
		return skyblockProfile, err
	}

	data, err := get("api.hypixel.net/skyblock/profiles?uuid=" + uuid + "&key=" + client.Key)

	if err != nil {
		return skyblockProfile, err
	}

	err = json.Unmarshal(data, &skyblockProfile)

	if !skyblockProfile.Success {
		err = errors.New(skyblockProfile.Cause)
	}

	return skyblockProfile, err
}

func (client *Client) SkyblockSkills() (structs.SkyblockSkills, error) {
	var skyblockSkills structs.SkyblockSkills

	data, err := get("api.hypixel.net/resources/skyblock/skills?key=" + client.Key)
	if err != nil {
		return skyblockSkills, err
	}

	err = json.Unmarshal(data, &skyblockSkills)

	return skyblockSkills, err
}
