package gopixel

import (
	"encoding/json"

	"errors"

	"github.com/comblock/gopixel/structs"
)

func (c *Client) GuildAchievements() (structs.GuildAchievements, error) {
	var guildAchievements structs.GuildAchievements

	data, err := get("api.hypixel.net/resources/guilds/achievements?key=" + c.Key)
	if err != nil {
		return guildAchievements, err
	}

	err = json.Unmarshal(data, &guildAchievements)

	return guildAchievements, err
}

func (client *Client) GuildPermissions() (structs.GuildPermissions, error) {
	var guildPermissions structs.GuildPermissions

	data, err := get("api.hypixel.net/resources/guilds/permissions?key=" + client.Key)
	if err != nil {
		return guildPermissions, err
	}

	err = json.Unmarshal(data, &guildPermissions)

	return guildPermissions, err
}

func (client *Client) GuildById(id string) (structs.Guild, error) {
	var guild structs.Guild

	data, err := get("api.hypixel.net/guild?id=" + id + "&key=" + client.Key)
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

	data, err := get("api.hypixel.net/guild?name=" + name + "&key=" + client.Key)
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

	data, err := get("api.hypixel.net/guild?player=" + uuid + "&key=" + client.Key)
	if err != nil {
		return guild, err
	}
	err = json.Unmarshal(data, &guild)

	if !guild.Success {
		err = errors.New(guild.Cause)
	}

	return guild, err
}
