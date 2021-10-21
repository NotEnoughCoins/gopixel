package structs

type RankedSkywars struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause"`
	Result  struct {
		Key      string `json:"key"`
		Position int    `json:"position"`
		Score    int    `json:"score"`
	} `json:"result"`
}
