package gopixel

import (
	"encoding/json"

	"errors"

	"github.com/comblock/gopixel/structs"
)

func (client *Client) RankedSkywars(name string) (structs.RankedSkywars, error) {
	var rankedSkywars structs.RankedSkywars

	uuid, err := Uuid(name)

	if err != nil {
		return rankedSkywars, err
	}

	data, err := get("api.hypixel.net/player/ranked/skywars?uuid=" + uuid + "&key=" + client.key)

	if err != nil {
		return rankedSkywars, err
	}

	err = json.Unmarshal(data, &rankedSkywars)

	if !rankedSkywars.Success {
		err = errors.New(rankedSkywars.Cause)
	}

	return rankedSkywars, err
}
