package structs

type PlayerCounts struct {
	Success bool `json:"success"`
	Games   struct {
		MAINLOBBY struct {
			Players int `json:"players"`
		} `json:"MAIN_LOBBY"`
		TOURNAMENTLOBBY struct {
			Players int `json:"players"`
		} `json:"TOURNAMENT_LOBBY"`
		SMP struct {
			Players int `json:"players"`
		} `json:"SMP"`
		MURDERMYSTERY struct {
			Players int `json:"players"`
			Modes   struct {
				MURDERDOUBLEUP  int `json:"MURDER_DOUBLE_UP"`
				MURDERINFECTION int `json:"MURDER_INFECTION"`
				MURDERASSASSINS int `json:"MURDER_ASSASSINS"`
				MURDERCLASSIC   int `json:"MURDER_CLASSIC"`
			} `json:"modes"`
		} `json:"MURDER_MYSTERY"`
		SPEEDUHC struct {
			Players int `json:"players"`
			Modes   struct {
				TeamNormal int `json:"team_normal"`
				SoloNormal int `json:"solo_normal"`
			} `json:"modes"`
		} `json:"SPEED_UHC"`
		UHC struct {
			Players int `json:"players"`
			Modes   struct {
				TEAMS int `json:"TEAMS"`
				SOLO  int `json:"SOLO"`
			} `json:"modes"`
		} `json:"UHC"`
		WALLS3 struct {
			Players int `json:"players"`
			Modes   struct {
				Standard int `json:"standard"`
			} `json:"modes"`
		} `json:"WALLS3"`
		LEGACY struct {
			Players int `json:"players"`
			Modes   struct {
				VAMPIREZ    int `json:"VAMPIREZ"`
				QUAKECRAFT  int `json:"QUAKECRAFT"`
				PAINTBALL   int `json:"PAINTBALL"`
				WALLS       int `json:"WALLS"`
				ARENA       int `json:"ARENA"`
				GINGERBREAD int `json:"GINGERBREAD"`
			} `json:"modes"`
		} `json:"LEGACY"`
		ARCADE struct {
			Players int `json:"players"`
			Modes   struct {
				PARTY                  int `json:"PARTY"`
				HOLEINTHEWALL          int `json:"HOLE_IN_THE_WALL"`
				DEFENDER               int `json:"DEFENDER"`
				MINIWALLS              int `json:"MINI_WALLS"`
				SIMONSAYS              int `json:"SIMON_SAYS"`
				ZOMBIESBADBLOOD        int `json:"ZOMBIES_BAD_BLOOD"`
				HIDEANDSEEKPARTYPOOPER int `json:"HIDE_AND_SEEK_PARTY_POOPER"`
				DAYONE                 int `json:"DAYONE"`
				DRAWTHEIRTHING         int `json:"DRAW_THEIR_THING"`
				ZOMBIESALIENARCADIUM   int `json:"ZOMBIES_ALIEN_ARCADIUM"`
				SOCCER                 int `json:"SOCCER"`
				PVPCTW                 int `json:"PVP_CTW"`
				THROWOUT               int `json:"THROW_OUT"`
				ENDER                  int `json:"ENDER"`
				STARWARS               int `json:"STARWARS"`
				HALLOWEENSIMULATOR     int `json:"HALLOWEEN_SIMULATOR"`
				DRAGONWARS2            int `json:"DRAGONWARS2"`
				ZOMBIESDEADEND         int `json:"ZOMBIES_DEAD_END"`
				HIDEANDSEEKPROPHUNT    int `json:"HIDE_AND_SEEK_PROP_HUNT"`
				FARMHUNT               int `json:"FARM_HUNT"`
			} `json:"modes"`
		} `json:"ARCADE"`
		TNTGAMES struct {
			Players int `json:"players"`
			Modes   struct {
				PVPRUN    int `json:"PVPRUN"`
				TNTAG     int `json:"TNTAG"`
				TNTRUN    int `json:"TNTRUN"`
				BOWSPLEEF int `json:"BOWSPLEEF"`
				CAPTURE   int `json:"CAPTURE"`
			} `json:"modes"`
		} `json:"TNTGAMES"`
		SURVIVALGAMES struct {
			Players int `json:"players"`
			Modes   struct {
				SoloNormal  int `json:"solo_normal"`
				TeamsNormal int `json:"teams_normal"`
			} `json:"modes"`
		} `json:"SURVIVAL_GAMES"`
		REPLAY struct {
			Players int `json:"players"`
			Modes   struct {
				BASE int `json:"BASE"`
			} `json:"modes"`
		} `json:"REPLAY"`
		SUPERSMASH struct {
			Players int `json:"players"`
			Modes   struct {
				OneV1Normal int `json:"1v1_normal"`
				SoloNormal  int `json:"solo_normal"`
				TwoV2Normal int `json:"2v2_normal"`
				TeamsNormal int `json:"teams_normal"`
			} `json:"modes"`
		} `json:"SUPER_SMASH"`
		BEDWARS struct {
			Players int `json:"players"`
			Modes   struct {
				BEDWARSTWOFOUR            int `json:"BEDWARS_TWO_FOUR"`
				BEDWARSEIGHTONE           int `json:"BEDWARS_EIGHT_ONE"`
				BEDWARSEIGHTTWOVOIDLESS   int `json:"BEDWARS_EIGHT_TWO_VOIDLESS"`
				BEDWARSEIGHTTWORUSH       int `json:"BEDWARS_EIGHT_TWO_RUSH"`
				BEDWARSEIGHTTWOSWAP       int `json:"BEDWARS_EIGHT_TWO_SWAP"`
				BEDWARSEIGHTTWOUNDERWORLD int `json:"BEDWARS_EIGHT_TWO_UNDERWORLD"`
				BEDWARSEIGHTTWO           int `json:"BEDWARS_EIGHT_TWO"`
				BEDWARSFOURFOUR           int `json:"BEDWARS_FOUR_FOUR"`
				BEDWARSFOURFOURUNDERWORLD int `json:"BEDWARS_FOUR_FOUR_UNDERWORLD"`
				BEDWARSEIGHTTWOLUCKY      int `json:"BEDWARS_EIGHT_TWO_LUCKY"`
				BEDWARSPRACTICE           int `json:"BEDWARS_PRACTICE"`
				BEDWARSFOURTHREE          int `json:"BEDWARS_FOUR_THREE"`
				BEDWARSFOURFOURLUCKY      int `json:"BEDWARS_FOUR_FOUR_LUCKY"`
			} `json:"modes"`
		} `json:"BEDWARS"`
		BUILDBATTLE struct {
			Players int `json:"players"`
			Modes   struct {
				BUILDBATTLEHALLOWEEN        int `json:"BUILD_BATTLE_HALLOWEEN"`
				BUILDBATTLESOLONORMALLATEST int `json:"BUILD_BATTLE_SOLO_NORMAL_LATEST"`
				BUILDBATTLEGUESSTHEBUILD    int `json:"BUILD_BATTLE_GUESS_THE_BUILD"`
				BUILDBATTLETEAMSNORMAL      int `json:"BUILD_BATTLE_TEAMS_NORMAL"`
				BUILDBATTLESOLOPRO          int `json:"BUILD_BATTLE_SOLO_PRO"`
				BUILDBATTLESOLONORMAL       int `json:"BUILD_BATTLE_SOLO_NORMAL"`
			} `json:"modes"`
		} `json:"BUILD_BATTLE"`
		PROTOTYPE struct {
			Players int `json:"players"`
			Modes   struct {
				PIXELPARTY         int `json:"PIXEL_PARTY"`
				TOWERWARSSOLO      int `json:"TOWERWARS_SOLO"`
				TOWERWARSTEAMOFTWO int `json:"TOWERWARS_TEAM_OF_TWO"`
			} `json:"modes"`
		} `json:"PROTOTYPE"`
		DUELS struct {
			Players int `json:"players"`
			Modes   struct {
				DUELSBOWSPLEEFDUEL int `json:"DUELS_BOWSPLEEF_DUEL"`
				DUELSBOWDUEL       int `json:"DUELS_BOW_DUEL"`
				DUELSMWDUEL        int `json:"DUELS_MW_DUEL"`
				DUELSUHCFOUR       int `json:"DUELS_UHC_FOUR"`
				DUELSUHCMEETUP     int `json:"DUELS_UHC_MEETUP"`
				DUELSSWDOUBLES     int `json:"DUELS_SW_DOUBLES"`
				DUELSUHCDOUBLES    int `json:"DUELS_UHC_DOUBLES"`
				DUELSBRIDGEFOUR    int `json:"DUELS_BRIDGE_FOUR"`
				DUELSSUMODUEL      int `json:"DUELS_SUMO_DUEL"`
				DUELSBRIDGETHREES  int `json:"DUELS_BRIDGE_THREES"`
				DUELSOPDUEL        int `json:"DUELS_OP_DUEL"`
				DUELSDUELARENA     int `json:"DUELS_DUEL_ARENA"`
				DUELSCOMBODUEL     int `json:"DUELS_COMBO_DUEL"`
				DUELSBRIDGEDOUBLES int `json:"DUELS_BRIDGE_DOUBLES"`
				DUELSBOXINGDUEL    int `json:"DUELS_BOXING_DUEL"`
				DUELSBRIDGE3V3V3V3 int `json:"DUELS_BRIDGE_3V3V3V3"`
				DUELSUHCDUEL       int `json:"DUELS_UHC_DUEL"`
				DUELSOPDOUBLES     int `json:"DUELS_OP_DOUBLES"`
				DUELSPARKOUREIGHT  int `json:"DUELS_PARKOUR_EIGHT"`
				DUELSCAPTURETHREES int `json:"DUELS_CAPTURE_THREES"`
				DUELSBLITZDUEL     int `json:"DUELS_BLITZ_DUEL"`
				DUELSPOTIONDUEL    int `json:"DUELS_POTION_DUEL"`
				DUELSCLASSICDUEL   int `json:"DUELS_CLASSIC_DUEL"`
				DUELSBRIDGEDUEL    int `json:"DUELS_BRIDGE_DUEL"`
				DUELSSWDUEL        int `json:"DUELS_SW_DUEL"`
			} `json:"modes"`
		} `json:"DUELS"`
		BATTLEGROUND struct {
			Players int `json:"players"`
			Modes   struct {
				CtfMini        int `json:"ctf_mini"`
				Domination     int `json:"domination"`
				TeamDeathmatch int `json:"team_deathmatch"`
			} `json:"modes"`
		} `json:"BATTLEGROUND"`
		PIT struct {
			Players int `json:"players"`
			Modes   struct {
				PIT int `json:"PIT"`
			} `json:"modes"`
		} `json:"PIT"`
		MCGO struct {
			Players int `json:"players"`
			Modes   struct {
				Normal     int `json:"normal"`
				Deathmatch int `json:"deathmatch"`
			} `json:"modes"`
		} `json:"MCGO"`
		SKYBLOCK struct {
			Players int `json:"players"`
			Modes   struct {
				DungeonHub     int `json:"dungeon_hub"`
				Farming1       int `json:"farming_1"`
				Mining2        int `json:"mining_2"`
				Mining3        int `json:"mining_3"`
				CrystalHollows int `json:"crystal_hollows"`
				Combat3        int `json:"combat_3"`
				Dynamic        int `json:"dynamic"`
				Combat1        int `json:"combat_1"`
				Foraging1      int `json:"foraging_1"`
				Hub            int `json:"hub"`
				Combat2        int `json:"combat_2"`
				Dungeon        int `json:"dungeon"`
				Mining1        int `json:"mining_1"`
			} `json:"modes"`
		} `json:"SKYBLOCK"`
		SKYWARS struct {
			Players int `json:"players"`
			Modes   struct {
				RankedNormal    int `json:"ranked_normal"`
				SoloInsaneSlime int `json:"solo_insane_slime"`
				TeamsInsane     int `json:"teams_insane"`
				SoloNormal      int `json:"solo_normal"`
				SoloInsane      int `json:"solo_insane"`
				TeamsNormal     int `json:"teams_normal"`
				MegaNormal      int `json:"mega_normal"`
			} `json:"modes"`
		} `json:"SKYWARS"`
		HOUSING struct {
			Players int `json:"players"`
		} `json:"HOUSING"`
		LIMBO struct {
			Players int `json:"players"`
		} `json:"LIMBO"`
		IDLE struct {
			Players int `json:"players"`
		} `json:"IDLE"`
		QUEUE struct {
			Players int `json:"players"`
		} `json:"QUEUE"`
	} `json:"games"`
	PlayerCount int `json:"playerCount"`
}
