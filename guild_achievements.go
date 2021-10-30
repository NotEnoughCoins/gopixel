package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
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
