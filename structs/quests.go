package structs

type Quests struct {
	Success     bool  `json:"success"`
	LastUpdated int64 `json:"lastUpdated"`
	Quests      struct {
		Quake []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"quake"`
		Walls []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"walls"`
		Paintball []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"paintball"`
		Hungergames []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"hungergames"`
		Tntgames []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"tntgames"`
		Vampirez []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"vampirez"`
		Walls3 []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"walls3"`
		Arcade []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"arcade"`
		Arena []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"arena"`
		Uhc []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"uhc"`
		Mcgo []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"mcgo"`
		Battleground []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"battleground"`
		Supersmash []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"supersmash"`
		Gingerbread []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"gingerbread"`
		Skywars []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"skywars"`
		Truecombat []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"truecombat"`
		Skyclash []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"skyclash"`
		Prototype []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"prototype"`
		Bedwars []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"bedwars"`
		Murdermystery []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"murdermystery"`
		Buildbattle []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"buildbattle"`
		Duels []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Rewards []struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"rewards"`
			Objectives []struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Integer int    `json:"integer"`
			} `json:"objectives"`
			Requirements []struct {
				Type string `json:"type"`
			} `json:"requirements"`
			Description string `json:"description"`
		} `json:"duels"`
	} `json:"quests"`
}
