package gopixel

import (
	"encoding/json"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) Challenges() (structs.Challenges, error) {
	var challenges structs.Challenges

	data, err := get("api.hypixel.net/resources/challenges?challenges=" + client.Key)
	if err != nil {
		return challenges, err
	}

	err = json.Unmarshal(data, &challenges)

	return challenges, err
}
