package gopixel

import (
	"encoding/json"

	"errors"

	"github.com/comblock/gopixel/structs"
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

func (client *Client) Friends(name string) (structs.Friends, error) {
	var friends structs.Friends

	uuid, err := Uuid(name)

	if err != nil {
		return friends, err
	}
	data, err := get("api.hypixel.net/friends?uuid=" + uuid + "&key=" + client.Key)

	if err != nil {
		return friends, err
	}

	err = json.Unmarshal(data, &friends)

	if !friends.Success {
		err = errors.New(friends.Cause)
	}

	return friends, err
}

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

func (client *Client) RecentGames(name string) (structs.RecentGames, error) {
	var recentGames structs.RecentGames

	uuid, err := Uuid(name)
	if err != nil {
		return recentGames, err
	}

	data, err := get("api.hypixel.net/recentgames?uuid=" + uuid + "&key=" + client.Key)
	if err != nil {
		return recentGames, err
	}

	err = json.Unmarshal(data, &recentGames)

	return recentGames, err
}
