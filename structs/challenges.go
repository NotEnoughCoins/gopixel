package structs

type Challenges struct {
	Success     bool  `json:"success"`
	LastUpdated int64 `json:"lastUpdated"`
	Data        struct {
		Arcade []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"arcade"`
		Arena []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"arena"`
		Bedwars []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"bedwars"`
		Hungergames []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"hungergames"`
		Buildbattle []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"buildbattle"`
		Truecombat []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"truecombat"`
		Duels []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"duels"`
		Mcgo []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"mcgo"`
		Murdermystery []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"murdermystery"`
		Paintball []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"paintball"`
		Quake []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"quake"`
		Skyclash []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"skyclash"`
		Skywars []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"skywars"`
		Supersmash []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"supersmash"`
		Speeduhc []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"speeduhc"`
		Gingerbread []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"gingerbread"`
		Tntgames []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"tntgames"`
		Uhc []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"uhc"`
		Vampirez []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"vampirez"`
		Walls3 []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"walls3"`
		Walls []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"walls"`
		Battleground []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
		} `json:"battleground"`
	} `json:"challenges"`
}
