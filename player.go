package gopixel

import (
	"encoding/json"

	"errors"

	structs "github.com/comblock/gopixel/structs"
)

func Uuid(name string) (string, error) {
	type MojangPlayer struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}

	data, err := get("api.mojang.com/users/profiles/minecraft/" + name)
	bytes := []byte(data)

	var mojang_player MojangPlayer

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(bytes, &mojang_player)

	return mojang_player.ID, err

}

func (client *Client) PlayerData(name string) (structs.Player, error) {
	var player structs.Player

	uuid, err := Uuid(name)

	if err != nil {
		return player, err
	}
	data, err := get("api.hypixel.net/player?uuid=" + uuid + "&key=" + client.Key)

	if err != nil {
		return player, err
	}

	err = json.Unmarshal(data, &player)

	if !player.Success {
		err = errors.New(player.Cause)
	}

	return player, err
}
