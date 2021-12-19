package structs

type Vanity struct {
	Success     bool   `json:"success"`
	Cause       string `json:"cause"`
	LastUpdated int64  `json:"lastUpdated"`
	Types       []struct {
		Key     string `json:"key"`
		Name    string `json:"name"`
		Rarity  string `json:"rarity"`
		Package string `json:"package"`
	} `json:"types"`
	Rarities []struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"rarities"`
}
