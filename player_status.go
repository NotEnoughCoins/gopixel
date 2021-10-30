package gopixel

import (
	"encoding/json"

	"errors"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) PlayerStatus(name string) (structs.PlayerStatus, error) {
	var playerStatus structs.PlayerStatus

	uuid, err := Uuid(name)

	if err != nil {
		return playerStatus, err
	}
	data, err := get("api.hypixel.net/status?uuid=" + uuid + "&key=" + client.Key)

	if err != nil {
		return playerStatus, err
	}

	err = json.Unmarshal(data, &playerStatus)

	if !playerStatus.Success {
		err = errors.New(playerStatus.Cause)
	}

	return playerStatus, err
}
