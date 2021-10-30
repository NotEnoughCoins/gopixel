package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) GuildPermissions() (structs.GuildPermissions, error) {
	var guildPermissions structs.GuildPermissions

	data, err := get("api.hypixel.net/resources/guilds/permissions?key=" + client.Key)
	if err != nil {
		return guildPermissions, err
	}

	err = json.Unmarshal(data, &guildPermissions)

	return guildPermissions, err
}
