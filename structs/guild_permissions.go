package structs

type GuildPermissions struct {
	Success     bool  `json:"success"`
	LastUpdated int64 `json:"lastUpdated"`
	Permissions []struct {
		EnUs struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Item        struct {
				Name string `json:"name"`
			} `json:"item"`
		} `json:"en_us"`
	} `json:"permissions"`
}
