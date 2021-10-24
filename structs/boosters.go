package structs

type Boosters struct {
	Success  bool `json:"success"`
	Boosters []struct {
		ID             string      `json:"_id"`
		PurchaserUUID  string      `json:"purchaserUuid"`
		Amount         float64     `json:"amount"`
		OriginalLength int         `json:"originalLength"`
		Length         int         `json:"length"`
		GameType       int         `json:"gameType"`
		Stacked        interface{} `json:"stacked,omitempty"`
		DateActivated  int64       `json:"dateActivated"`
	} `json:"boosters"`
	BoosterState struct {
		Decrementing bool `json:"decrementing"`
	} `json:"boosterState"`
}
