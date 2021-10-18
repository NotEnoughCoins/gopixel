package structs

type Guild struct {
	Success bool `json:"success"`
	Data    struct {
		ID        string `json:"_id"`
		Name      string `json:"name"`
		Coins     int    `json:"coins"`
		CoinsEver int    `json:"coinsEver"`
		Created   int    `json:"created"`
		Members   []struct {
			UUID               string `json:"uuid"`
			Rank               string `json:"rank"`
			Joined             int    `json:"joined"`
			QuestParticipation int    `json:"questParticipation"`
			ExpHistory         struct {
			} `json:"expHistory"`
		} `json:"members"`
		Tag          string `json:"tag"`
		Joinable     bool   `json:"joinable"`
		Achievements struct {
			Winners         int `json:"WINNERS"`
			ExperienceKings int `json:"EXPERIENCE_KINGS"`
			OnlinePlayers   int `json:"ONLINE_PLAYERS"`
		} `json:"achievements"`
		Exp           int `json:"exp"`
		ChatMute      int `json:"chatMute"`
		LegacyRanking int `json:"legacyRanking"`
		Ranks         []struct {
			Name     string `json:"name"`
			Default  bool   `json:"default"`
			Tag      string `json:"tag"`
			Created  int    `json:"created"`
			Priority int    `json:"priority"`
		} `json:"ranks"`
		PubliclyListed     bool   `json:"publiclyListed"`
		NameLower          string `json:"name_lower"`
		TagColor           string `json:"tagColor"`
		Description        string `json:"description"`
		GuildExpByGameType struct {
			Quakecraft    int `json:"QUAKECRAFT"`
			Walls         int `json:"WALLS"`
			Paintball     int `json:"PAINTBALL"`
			SurvivalGames int `json:"SURVIVAL_GAMES"`
			Tntgames      int `json:"TNTGAMES"`
			Vampirez      int `json:"VAMPIREZ"`
			Walls3        int `json:"WALLS3"`
			Arcade        int `json:"ARCADE"`
			Arena         int `json:"ARENA"`
			Mcgo          int `json:"MCGO"`
			Uhc           int `json:"UHC"`
			Battleground  int `json:"BATTLEGROUND"`
			SuperSmash    int `json:"SUPER_SMASH"`
			Gingerbread   int `json:"GINGERBREAD"`
			Housing       int `json:"HOUSING"`
			Skywars       int `json:"SKYWARS"`
			SpeedUhc      int `json:"SPEED_UHC"`
			Prototype     int `json:"PROTOTYPE"`
			Bedwars       int `json:"BEDWARS"`
			MurderMystery int `json:"MURDER_MYSTERY"`
			BuildBattle   int `json:"BUILD_BATTLE"`
			Duels         int `json:"DUELS"`
			Pit           int `json:"PIT"`
			Legacy        int `json:"LEGACY"`
			Replay        int `json:"REPLAY"`
			Skyblock      int `json:"SKYBLOCK"`
			Smp           int `json:"SMP"`
		} `json:"guildExpByGameType"`
		Banner struct {
			Base     string `json:"Base"`
			Patterns []struct {
				Pattern string `json:"Pattern"`
				Color   string `json:"Color"`
			} `json:"Patterns"`
		} `json:"banner"`
		PreferredGames  []interface{} `json:"preferredGames"`
		HideGmTag       bool          `json:"hideGmTag"`
		MemberSizeLevel int           `json:"memberSizeLevel"`
		BankSizeLevel   int           `json:"bankSizeLevel"`
		CanTag          bool          `json:"canTag"`
		VipCount        int           `json:"vipCount"`
		MvpCount        int           `json:"mvpCount"`
		CanMotd         bool          `json:"canMotd"`
		CanParty        bool          `json:"canParty"`
		Motd            []interface{} `json:"motd"`
		Discord         string        `json:"discord"`
		ChatThrottle    int           `json:"chatThrottle"`
	} `json:"guild"`
	Cause    string `json:"cause"`
	Throttle bool   `json:"throttle"`
}
