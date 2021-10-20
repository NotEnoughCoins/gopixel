package gopixel

import (
	"encoding/json"

	"errors"

	"gopixel/structs"
)

func (client *Client) Achievements() (structs.Achievements, error) {
	var achievements structs.Achievements

	data, err := get("api.hypixel.net/resources/achievements?key=" + client.key)

	if err != nil {
		return achievements, err
	}

	err = json.Unmarshal(data, &achievements)

	if !achievements.Success {
		err = errors.New(achievements.Cause)
	}

	return achievements, err
}
