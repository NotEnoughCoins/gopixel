package gopixel

import (
	"encoding/json"

	"errors"

	"gopixel/structs"
)

func (client *Client) SkyblockProfile(profile string) (structs.SkyblockProfile, error) {
	var skyblockProfile structs.SkyblockProfile

	data, err := get("api.hypixel.net/skyblock/profile?profile=" + profile + "&key=" + client.key)

	if err != nil {
		return skyblockProfile, err
	}

	err = json.Unmarshal(data, &skyblockProfile)

	if !skyblockProfile.Success {
		err = errors.New(skyblockProfile.Cause)
	}

	return skyblockProfile, err
}
