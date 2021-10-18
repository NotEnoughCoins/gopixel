package hypixel

import (
	"encoding/json"
	"errors"

	"hypixel.go/structs"
)

func (client *Client) Achievements(name string) (structs.Achievements, error) {
	var achievements structs.Achievements

	data, err := get("api.hypixel.net/resources/achievements")

	if err != nil {
		return achievements, err
	}

	err = json.Unmarshal(data, &achievements)

	if !achievements.Success {
		err = errors.New(achievements.Cause)
	}

	return achievements, err
}
