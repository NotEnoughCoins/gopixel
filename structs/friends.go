package structs

type Friends struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause"`
	UUID    string `json:"uuid"`
	Records []struct {
		ID           string `json:"_id"`
		UUIDSender   string `json:"uuidSender"`
		UUIDReceiver string `json:"uuidReceiver"`
		Started      int64  `json:"started"`
	} `json:"records"`
}
