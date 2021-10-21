package gopixel

import (
	"encoding/json"

	structs "gopixel/structs"
)

func (client *Client) SkyblockSkills() (structs.SkyblockSkills, error) {
	var skyblockSkills structs.SkyblockSkills

	data, err := get("api.hypixel.net/resources/skyblock/skills?key=" + client.key)
	if err != nil {
		return skyblockSkills, err
	}

	err = json.Unmarshal(data, &skyblockSkills)

	return skyblockSkills, err
}
