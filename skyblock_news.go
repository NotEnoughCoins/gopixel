package gopixel

import (
	"encoding/json"
	"gopixel/structs"
)

func (c *Client) SkyblockNews() (structs.SkyblockNews, error) {
	var news structs.SkyblockNews

	data, err := get("api.hypixel.net/resources/guilds/permissions?key=" + c.key)
	if err != nil {
		return news, err
	}

	err = json.Unmarshal(data, &news)

	return news, err
}
