package structs

type Games struct {
	Success     bool  `json:"success"`
	LastUpdated int64 `json:"lastUpdated"`
	Data        struct {
		Quakecraft struct {
			ModeNames struct {
				Teams       string `json:"teams"`
				Solo        string `json:"solo"`
				SoloTourney string `json:"solo_tourney"`
			} `json:"modeNames"`
			Legacy       bool   `json:"legacy"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"QUAKECRAFT"`
		Skyclash struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			Retired      bool   `json:"retired"`
			ID           int    `json:"id"`
		} `json:"SKYCLASH"`
		BuildBattle struct {
			ModeNames struct {
				BuildBattleHalloween         string `json:"BUILD_BATTLE_HALLOWEEN"`
				BuildBattleChristmasNewTeams string `json:"BUILD_BATTLE_CHRISTMAS_NEW_TEAMS"`
				BuildBattleSoloNormalLatest  string `json:"BUILD_BATTLE_SOLO_NORMAL_LATEST"`
				BuildBattleGuessTheBuild     string `json:"BUILD_BATTLE_GUESS_THE_BUILD"`
				BuildBattleTeamsNormal       string `json:"BUILD_BATTLE_TEAMS_NORMAL"`
				BuildBattleSoloNormal        string `json:"BUILD_BATTLE_SOLO_NORMAL"`
				BuildBattleSoloPro           string `json:"BUILD_BATTLE_SOLO_PRO"`
				BuildBattleChristmasNewSolo  string `json:"BUILD_BATTLE_CHRISTMAS_NEW_SOLO"`
				BuildBattleChristmas         string `json:"BUILD_BATTLE_CHRISTMAS"`
			} `json:"modeNames"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"BUILD_BATTLE"`
		Uhc struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"UHC"`
		Legacy struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"LEGACY"`
		Skyblock struct {
			ModeNames struct {
				Farming1       string `json:"farming_1"`
				CrystalHollows string `json:"crystal_hollows"`
				Foraging1      string `json:"foraging_1"`
				DarkAuction    string `json:"dark_auction"`
				Dungeon        string `json:"dungeon"`
				Combat3        string `json:"combat_3"`
				Hub            string `json:"hub"`
				Dynamic        string `json:"dynamic"`
				Mining3        string `json:"mining_3"`
				Mining1        string `json:"mining_1"`
				Combat2        string `json:"combat_2"`
				Mining2        string `json:"mining_2"`
				Combat1        string `json:"combat_1"`
			} `json:"modeNames"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"SKYBLOCK"`
		Housing struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"HOUSING"`
		Mcgo struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"MCGO"`
		SurvivalGames struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"SURVIVAL_GAMES"`
		Battleground struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"BATTLEGROUND"`
		MurderMystery struct {
			ModeNames struct {
				MurderDoubleUp  string `json:"MURDER_DOUBLE_UP"`
				MurderInfection string `json:"MURDER_INFECTION"`
				MurderAssassins string `json:"MURDER_ASSASSINS"`
				MurderClassic   string `json:"MURDER_CLASSIC"`
			} `json:"modeNames"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"MURDER_MYSTERY"`
		Arcade struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"ARCADE"`
		Arena struct {
			Legacy       bool   `json:"legacy"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"ARENA"`
		Tntgames struct {
			ModeNames struct {
				TeamsNormal   string `json:"TEAMS_NORMAL"`
				Pvprun        string `json:"PVPRUN"`
				Tntag         string `json:"TNTAG"`
				TntrunTourney string `json:"TNTRUN_TOURNEY"`
				Tntrun        string `json:"TNTRUN"`
				Bowspleef     string `json:"BOWSPLEEF"`
				Capture       string `json:"CAPTURE"`
			} `json:"modeNames"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"TNTGAMES"`
		Walls struct {
			Legacy       bool   `json:"legacy"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"WALLS"`
		Skywars struct {
			ModeNames struct {
				SoloInsaneLucky           string `json:"solo_insane_lucky"`
				SoloInsaneSlime           string `json:"solo_insane_slime"`
				TeamsInsaneSlime          string `json:"teams_insane_slime"`
				TeamsInsaneRush           string `json:"teams_insane_rush"`
				TeamsInsaneLucky          string `json:"teams_insane_lucky"`
				SoloNormal                string `json:"solo_normal"`
				TeamsInsane               string `json:"teams_insane"`
				SoloInsaneHuntersVsBeasts string `json:"solo_insane_hunters_vs_beasts"`
				RankedNormal              string `json:"ranked_normal"`
				SoloInsaneTntMadness      string `json:"solo_insane_tnt_madness"`
				MegaDoubles               string `json:"mega_doubles"`
				SoloInsaneRush            string `json:"solo_insane_rush"`
				SoloInsane                string `json:"solo_insane"`
				TeamsInsaneTntMadness     string `json:"teams_insane_tnt_madness"`
				TeamsNormal               string `json:"teams_normal"`
				MegaNormal                string `json:"mega_normal"`
			} `json:"modeNames"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"SKYWARS"`
		Vampirez struct {
			Legacy       bool   `json:"legacy"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"VAMPIREZ"`
		Prototype struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"PROTOTYPE"`
		Walls3 struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"WALLS3"`
		Bedwars struct {
			ModeNames struct {
				BedwarsTwoFour            string `json:"BEDWARS_TWO_FOUR"`
				BedwarsEightOne           string `json:"BEDWARS_EIGHT_ONE"`
				BedwarsFourFourRush       string `json:"BEDWARS_FOUR_FOUR_RUSH"`
				BedwarsEightTwoRush       string `json:"BEDWARS_EIGHT_TWO_RUSH"`
				BedwarsEightTwoVoidless   string `json:"BEDWARS_EIGHT_TWO_VOIDLESS"`
				BedwarsFourFourArmed      string `json:"BEDWARS_FOUR_FOUR_ARMED"`
				BedwarsEightTwoArmed      string `json:"BEDWARS_EIGHT_TWO_ARMED"`
				BedwarsEightTwoUnderworld string `json:"BEDWARS_EIGHT_TWO_UNDERWORLD"`
				BedwarsEightTwoSwap       string `json:"BEDWARS_EIGHT_TWO_SWAP"`
				BedwarsEightTwo           string `json:"BEDWARS_EIGHT_TWO"`
				BedwarsFourFour           string `json:"BEDWARS_FOUR_FOUR"`
				BedwarsFourFourUnderworld string `json:"BEDWARS_FOUR_FOUR_UNDERWORLD"`
				BedwarsEightOneRush       string `json:"BEDWARS_EIGHT_ONE_RUSH"`
				BedwarsFourFourUltimate   string `json:"BEDWARS_FOUR_FOUR_ULTIMATE"`
				BedwarsEightTwoLucky      string `json:"BEDWARS_EIGHT_TWO_LUCKY"`
				BedwarsFourFourSwap       string `json:"BEDWARS_FOUR_FOUR_SWAP"`
				BedwarsFourThree          string `json:"BEDWARS_FOUR_THREE"`
				BedwarsFourFourVoidless   string `json:"BEDWARS_FOUR_FOUR_VOIDLESS"`
				BedwarsFourFourLucky      string `json:"BEDWARS_FOUR_FOUR_LUCKY"`
				BedwarsCastle             string `json:"BEDWARS_CASTLE"`
				BedwarsEightOneUltimate   string `json:"BEDWARS_EIGHT_ONE_ULTIMATE"`
				BedwarsEightTwoUltimate   string `json:"BEDWARS_EIGHT_TWO_ULTIMATE"`
			} `json:"modeNames"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"BEDWARS"`
		Paintball struct {
			Legacy       bool   `json:"legacy"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"PAINTBALL"`
		SuperSmash struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"SUPER_SMASH"`
		Smp struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"SMP"`
		Replay struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"REPLAY"`
		TrueCombat struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			Retired      bool   `json:"retired"`
			ID           int    `json:"id"`
		} `json:"TRUE_COMBAT"`
		Pit struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"PIT"`
		SpeedUhc struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"SPEED_UHC"`
		Duels struct {
			ModeNames struct {
				DuelsBowspleefDuel    string `json:"DUELS_BOWSPLEEF_DUEL"`
				DuelsBowDuel          string `json:"DUELS_BOW_DUEL"`
				DuelsMwDuel           string `json:"DUELS_MW_DUEL"`
				DuelsUhcFour          string `json:"DUELS_UHC_FOUR"`
				DuelsUhcMeetup        string `json:"DUELS_UHC_MEETUP"`
				DuelsSwDoubles        string `json:"DUELS_SW_DOUBLES"`
				DuelsUhcDoubles       string `json:"DUELS_UHC_DOUBLES"`
				DuelsSwFour           string `json:"DUELS_SW_FOUR"`
				DuelsBridgeFour       string `json:"DUELS_BRIDGE_FOUR"`
				DuelsSumoDuel         string `json:"DUELS_SUMO_DUEL"`
				DuelsOpDuel           string `json:"DUELS_OP_DUEL"`
				DuelsMwDoubles        string `json:"DUELS_MW_DOUBLES"`
				DuelsSumoTournament   string `json:"DUELS_SUMO_TOURNAMENT"`
				DuelsComboDuel        string `json:"DUELS_COMBO_DUEL"`
				DuelsBridge2V2V2V2    string `json:"DUELS_BRIDGE_2V2V2V2"`
				DuelsBridgeDoubles    string `json:"DUELS_BRIDGE_DOUBLES"`
				DuelsBridgeTournament string `json:"DUELS_BRIDGE_TOURNAMENT"`
				DuelsUhcDuel          string `json:"DUELS_UHC_DUEL"`
				DuelsBridge3V3V3V3    string `json:"DUELS_BRIDGE_3V3V3V3"`
				DuelsUhcTournament    string `json:"DUELS_UHC_TOURNAMENT"`
				DuelsOpDoubles        string `json:"DUELS_OP_DOUBLES"`
				DuelsBlitzDuel        string `json:"DUELS_BLITZ_DUEL"`
				DuelsMwFour           string `json:"DUELS_MW_FOUR"`
				DuelsClassicDuel      string `json:"DUELS_CLASSIC_DUEL"`
				DuelsPotionDuel       string `json:"DUELS_POTION_DUEL"`
				DuelsBridgeDuel       string `json:"DUELS_BRIDGE_DUEL"`
				DuelsSwTournament     string `json:"DUELS_SW_TOURNAMENT"`
				DuelsSwDuel           string `json:"DUELS_SW_DUEL"`
			} `json:"modeNames"`
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"DUELS"`
		Gingerbread struct {
			DatabaseName string `json:"databaseName"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"GINGERBREAD"`
	} `json:"games"`
}
