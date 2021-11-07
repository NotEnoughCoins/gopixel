package gopixel

import (
	"encoding/json"

	"errors"

	"github.com/comblock/gopixel/structs"
)

// Method to get the global achievements
func (client *Client) Achievements() (structs.Achievements, error) {
	var achievements structs.Achievements

	data, err := client.get("api.hypixel.net/resources/achievements?key=" + client.Key)

	if err != nil {
		return achievements, err
	}

	err = json.Unmarshal(data, &achievements)

	if !achievements.Success {
		err = errors.New(achievements.Cause)
	}

	return achievements, err
}

// Method to get a list of the games
func (client *Client) Games() (structs.Games, error) {
	var games structs.Games

	data, err := client.get("api.hypixel.net/resources/games?key=" + client.Key)
	if err != nil {
		return games, err
	}

	err = json.Unmarshal(data, &games)

	return games, err
}

// Method to get the active boosters
func (client *Client) Boosters() (structs.Boosters, error) {
	var boosters structs.Boosters

	data, err := client.get("api.hypixel.net/boosters?key=" + client.Key)
	if err != nil {
		return boosters, err
	}

	err = json.Unmarshal(data, &boosters)

	return boosters, err
}

// Method to get a list of all challenges
func (client *Client) Challenges() (structs.Challenges, error) {
	var challenges structs.Challenges

	data, err := client.get("api.hypixel.net/resources/challenges?challenges=" + client.Key)
	if err != nil {
		return challenges, err
	}

	err = json.Unmarshal(data, &challenges)

	return challenges, err
}

// Method to check if a key is valid and
func (client *Client) KeyData() (structs.Key, error) {
	var key structs.Key

	data, err := client.get("api.hypixel.net/key?key=" + client.Key)
	if err != nil {
		return key, err
	}

	err = json.Unmarshal(data, &key)

	if !key.Success {
		err = errors.New(key.Cause)
	}

	return key, err
}

// Method to get the current player counts
func (client *Client) PlayerCounts() (structs.PlayerCounts, error) {
	var counts structs.PlayerCounts
	data, err := client.get("api.hypixel.net/counts?key=" + client.Key)
	if err != nil {
		return counts, err
	}

	err = json.Unmarshal(data, &counts)

	return counts, err
}

// Method to get the current leaderboards
func (client *Client) Leaderboards() (structs.Leaderboards, error) {
	var leaderboards structs.Leaderboards

	data, err := client.get("api.hypixel.net/leaderboards?key=" + client.Key)
	if err != nil {
		return leaderboards, err
	}

	err = json.Unmarshal(data, &leaderboards)

	return leaderboards, err
}

// Method to get the punishment statistics
func (client *Client) PunishmentStats() (structs.PunishmentStats, error) {
	var stats structs.PunishmentStats

	data, err := client.get("api.hypixel.net/punishmentstats?key=" + client.Key)
	if err != nil {
		return stats, err
	}

	err = json.Unmarshal(data, &stats)

	return stats, err
}

// Method to get the quests
func (client *Client) Quests() (structs.Quests, error) {
	var quests structs.Quests

	data, err := client.get("api.hypixel.net/resources/quests?key=" + client.Key)
	if err != nil {
		return quests, err
	}

	err = json.Unmarshal(data, &quests)

	return quests, err
}

// Method to get the ranked skywars data of a player. Will return an error if no data is found
func (client *Client) RankedSkywars(name string) (structs.RankedSkywars, error) {
	var rankedSkywars structs.RankedSkywars

	uuid, err := client.Uuid(name)

	if err != nil {
		return rankedSkywars, err
	}

	data, err := client.get("api.hypixel.net/player/ranked/skywars?uuid=" + uuid + "&key=" + client.Key)

	if err != nil {
		return rankedSkywars, err
	}

	err = json.Unmarshal(data, &rankedSkywars)

	if !rankedSkywars.Success {
		err = errors.New(rankedSkywars.Cause)
	}

	return rankedSkywars, err
}
