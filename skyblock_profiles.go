package gopixel

import (
	"encoding/json"

	"errors"

	"github.com/comblock/gopixel/structs"
)

func (client *Client) SkyblockProfiles(name string) (structs.SkyblockProfiles, error) {
	var skyblockProfile structs.SkyblockProfiles

	uuid, err := Uuid(name)

	if err != nil {
		return skyblockProfile, err
	}

	data, err := get("api.hypixel.net/skyblock/profiles?uuid=" + uuid + "&key=" + client.key)

	if err != nil {
		return skyblockProfile, err
	}

	err = json.Unmarshal(data, &skyblockProfile)

	if !skyblockProfile.Success {
		err = errors.New(skyblockProfile.Cause)
	}

	return skyblockProfile, err
}
