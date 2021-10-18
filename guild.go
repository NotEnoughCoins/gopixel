package gopixel

import (
	"encoding/json"
	"errors"

	structs "gopixel/structs"
)

func (client *Client) GuildById(id string) (structs.Guild, error) {
	var guild structs.Guild

	data, err := get("api.hypixel.net/guild?id=" + id + "&key=" + client.key)
	if err != nil {
		return guild, err
	}

	err = json.Unmarshal(data, &guild)

	if !guild.Success {
		err = errors.New(guild.Cause)
	}

	return guild, err
}

func (client *Client) GuildByName(name string) (structs.Guild, error) {
	var guild structs.Guild

	data, err := get("api.hypixel.net/guild?name=" + name + "&key=" + client.key)
	if err != nil {
		return guild, err
	}

	err = json.Unmarshal(data, &guild)

	if !guild.Success {
		err = errors.New(guild.Cause)
	}
	return guild, err
}

func (client *Client) GuildByPlayer(player string) (structs.Guild, error) {
	var guild structs.Guild

	uuid, err := Uuid(player)
	if err != nil {
		return guild, err
	}

	data, err := get("api.hypixel.net/guild?player=" + uuid + "&key=" + client.key)
	if err != nil {
		return guild, err
	}
	err = json.Unmarshal(data, &guild)

	if !guild.Success {
		err = errors.New(guild.Cause)
	}

	return guild, err
}
