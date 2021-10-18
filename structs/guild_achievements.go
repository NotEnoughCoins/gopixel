package structs

type GuildAchievements struct {
	Success     bool  `json:"success"`
	LastUpdated int64 `json:"lastUpdated"`
	OneTime     struct {
	} `json:"one_time"`
	Tiered struct {
		Prestige struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Tiers       []struct {
				Tier   int `json:"tier"`
				Amount int `json:"amount"`
			} `json:"tiers"`
		} `json:"PRESTIGE"`
		ExperienceKings struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Tiers       []struct {
				Tier   int `json:"tier"`
				Amount int `json:"amount"`
			} `json:"tiers"`
		} `json:"EXPERIENCE_KINGS"`
		Winners struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Tiers       []struct {
				Tier   int `json:"tier"`
				Amount int `json:"amount"`
			} `json:"tiers"`
		} `json:"WINNERS"`
		OnlinePlayers struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Tiers       []struct {
				Tier   int `json:"tier"`
				Amount int `json:"amount"`
			} `json:"tiers"`
		} `json:"ONLINE_PLAYERS"`
	} `json:"tiered"`
}
