package gopixel

import (
	"errors"

	"github.com/NotEnoughCoins/gopixel/structs"
)

// Method to get a list of guild achievements
func (client *Client) GuildAchievements() (structs.GuildAchievements, error) {
	var guildAchievements structs.GuildAchievements

	data, err := client.get("api.hypixel.net/resources/guilds/achievements?key=" + client.Key)
	if err != nil {
		return guildAchievements, err
	}

	err = json.Unmarshal(data, &guildAchievements)

	return guildAchievements, err
}

// Method to get a guild by its id
func (client *Client) GuildById(id string) (structs.Guild, error) {
	var guild structs.Guild

	data, err := client.get("api.hypixel.net/guild?id=" + id + "&key=" + client.Key)
	if err != nil {
		return guild, err
	}

	err = json.Unmarshal(data, &guild)

	if !guild.Success {
		err = errors.New(guild.Cause)
	}

	return guild, err
}

// Method to get a guild by its name
func (client *Client) GuildByName(name string) (structs.Guild, error) {
	var guild structs.Guild

	data, err := client.get("api.hypixel.net/guild?name=" + name + "&key=" + client.Key)
	if err != nil {
		return guild, err
	}

	err = json.Unmarshal(data, &guild)

	if !guild.Success {
		err = errors.New(guild.Cause)
	}
	return guild, err
}

// Method to get a guild by a player
func (client *Client) GuildByPlayer(player string) (structs.Guild, error) {
	var guild structs.Guild

	uuid, err := client.Uuid(player)
	if err != nil {
		return guild, err
	}

	data, err := client.get("api.hypixel.net/guild?player=" + uuid + "&key=" + client.Key)
	if err != nil {
		return guild, err
	}
	err = json.Unmarshal(data, &guild)

	if !guild.Success {
		err = errors.New(guild.Cause)
	}

	return guild, err
}
