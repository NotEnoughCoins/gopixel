package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) PunishmentStats() (structs.PunishmentStats, error) {
	var stats structs.PunishmentStats

	data, err := get("api.hypixel.net/punishmentstats?key=" + client.Key)
	if err != nil {
		return stats, err
	}

	err = json.Unmarshal(data, &stats)

	return stats, err
}
