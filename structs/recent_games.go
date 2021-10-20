package structs

type RecentGames struct {
	Games []struct {
		Date     int64  `json:"date"`
		Ended    int64  `json:"ended"`
		GameType string `json:"gameType"`
		Map      string `json:"map"`
		Mode     string `json:"mode"`
	} `json:"games"`
	Success bool   `json:"success"`
	UUID    string `json:"uuid"`
}
