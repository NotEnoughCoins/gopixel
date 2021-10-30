package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) RecentGames(name string) (structs.RecentGames, error) {
	var recentGames structs.RecentGames

	uuid, err := Uuid(name)
	if err != nil {
		return recentGames, err
	}

	data, err := get("api.hypixel.net/recentgames?uuid=" + uuid + "&key=" + client.key)
	if err != nil {
		return recentGames, err
	}

	err = json.Unmarshal(data, &recentGames)

	return recentGames, err
}
