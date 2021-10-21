package structs

type PlayerStatus struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause"`
	UUID    string `json:"uuid"`
	Session struct {
		GameType string `json:"gameType"`
		Map      string `json:"map"`
		Mode     string `json:"mode"`
		Online   bool   `json:"online"`
	} `json:"session"`
}
