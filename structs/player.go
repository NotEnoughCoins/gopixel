package structs

type Parkour []struct {
	TimeStart int64 `json:"timeStart"`
	TimeTook  int   `json:"timeTook"`
}

type Player struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause"`
	Data    struct {
		ID           string `json:"_id"`
		Achievements struct {
			ArenaBossed                     int `json:"arena_bossed"`
			ArenaClimbTheRanks              int `json:"arena_climb_the_ranks"`
			ArenaGladiator                  int `json:"arena_gladiator"`
			ArenaGottaWearEmAll             int `json:"arena_gotta_wear_em_all"`
			BlitzCoins                      int `json:"blitz_coins"`
			BlitzKills                      int `json:"blitz_kills"`
			BlitzWins                       int `json:"blitz_wins"`
			GeneralCoins                    int `json:"general_coins"`
			GeneralWins                     int `json:"general_wins"`
			PaintballCoins                  int `json:"paintball_coins"`
			PaintballKills                  int `json:"paintball_kills"`
			PaintballWins                   int `json:"paintball_wins"`
			QuakeKillingSprees              int `json:"quake_killing_sprees"`
			QuakeKills                      int `json:"quake_kills"`
			QuakeWins                       int `json:"quake_wins"`
			TntgamesBowSpleefWins           int `json:"tntgames_bow_spleef_wins"`
			TntgamesTntRunWins              int `json:"tntgames_tnt_run_wins"`
			TntgamesWizardsWins             int `json:"tntgames_wizards_wins"`
			VampirezCoins                   int `json:"vampirez_coins"`
			VampirezKillSurvivors           int `json:"vampirez_kill_survivors"`
			VampirezKillVampires            int `json:"vampirez_kill_vampires"`
			VampirezSurvivorWins            int `json:"vampirez_survivor_wins"`
			Walls3Coins                     int `json:"walls3_coins"`
			Walls3Kills                     int `json:"walls3_kills"`
			Walls3Wins                      int `json:"walls3_wins"`
			WallsCoins                      int `json:"walls_coins"`
			WallsKills                      int `json:"walls_kills"`
			WallsWins                       int `json:"walls_wins"`
			WarlordsWarriorLevel            int `json:"warlords_warrior_level"`
			WarlordsMageLevel               int `json:"warlords_mage_level"`
			WarlordsPaladinLevel            int `json:"warlords_paladin_level"`
			WarlordsKills                   int `json:"warlords_kills"`
			WarlordsShamanLevel             int `json:"warlords_shaman_level"`
			WarlordsAssist                  int `json:"warlords_assist"`
			WarlordsRepairWeapons           int `json:"warlords_repair_weapons"`
			WarlordsCoins                   int `json:"warlords_coins"`
			SkywarsKitsTeam                 int `json:"skywars_kits_team"`
			SkywarsWinsTeam                 int `json:"skywars_wins_team"`
			SkywarsKillsSolo                int `json:"skywars_kills_solo"`
			SkywarsKillsTeam                int `json:"skywars_kills_team"`
			SkywarsCages                    int `json:"skywars_cages"`
			SkywarsWinsSolo                 int `json:"skywars_wins_solo"`
			SkywarsKitsSolo                 int `json:"skywars_kits_solo"`
			UhcChampion                     int `json:"uhc_champion"`
			UhcHunter                       int `json:"uhc_hunter"`
			UhcMovingUp                     int `json:"uhc_moving_up"`
			UhcBounty                       int `json:"uhc_bounty"`
			WarlordsCtfWins                 int `json:"warlords_ctf_wins"`
			WarlordsDomWins                 int `json:"warlords_dom_wins"`
			WarlordsTdmWins                 int `json:"warlords_tdm_wins"`
			Walls3Guardian                  int `json:"walls3_guardian"`
			TruecombatSoloWinner            int `json:"truecombat_solo_winner"`
			TruecombatSoloKiller            int `json:"truecombat_solo_killer"`
			TruecombatKitHoarderSolo        int `json:"truecombat_kit_hoarder_solo"`
			TruecombatKitHoarderTeam        int `json:"truecombat_kit_hoarder_team"`
			Walls3Rusher                    int `json:"walls3_rusher"`
			SupersmashSmashWinner           int `json:"supersmash_smash_winner"`
			SupersmashHeroSlayer            int `json:"supersmash_hero_slayer"`
			SupersmashSmashChampion         int `json:"supersmash_smash_champion"`
			CopsandcrimsSerialKiller        int `json:"copsandcrims_serial_killer"`
			CopsandcrimsHeroTerrorist       int `json:"copsandcrims_hero_terrorist"`
			CopsandcrimsBombSpecialist      int `json:"copsandcrims_bomb_specialist"`
			CopsandcrimsCacBanker           int `json:"copsandcrims_cac_banker"`
			SkywarsKitsMega                 int `json:"skywars_kits_mega"`
			SkywarsWinsMega                 int `json:"skywars_wins_mega"`
			SkywarsKillsMega                int `json:"skywars_kills_mega"`
			GeneralQuestMaster              int `json:"general_quest_master"`
			GeneralChallenger               int `json:"general_challenger"`
			SpeeduhcHunter                  int `json:"speeduhc_hunter"`
			SpeeduhcUhcMaster               int `json:"speeduhc_uhc_master"`
			SpeeduhcCollector               int `json:"speeduhc_collector"`
			SpeeduhcSalty                   int `json:"speeduhc_salty"`
			SkyclashCardsUnlocked           int `json:"skyclash_cards_unlocked"`
			SkyclashPacksOpened             int `json:"skyclash_packs_opened"`
			SkyclashKills                   int `json:"skyclash_kills"`
			SkyclashWins                    int `json:"skyclash_wins"`
			ArcadeArcadeWinner              int `json:"arcade_arcade_winner"`
			ArcadeArcadeBanker              int `json:"arcade_arcade_banker"`
			ArcadeFootballPro               int `json:"arcade_football_pro"`
			ArcadeFarmhuntDominator         int `json:"arcade_farmhunt_dominator"`
			ArcadeBountyHunter              int `json:"arcade_bounty_hunter"`
			ArcadeTeamWork                  int `json:"arcade_team_work"`
			ArcadeZombieKiller              int `json:"arcade_zombie_killer"`
			ArcadeMiniwallsWinner           int `json:"arcade_miniwalls_winner"`
			TruecombatTeamWinner            int `json:"truecombat_team_winner"`
			TruecombatTeamKiller            int `json:"truecombat_team_killer"`
			QuakeCoins                      int `json:"quake_coins"`
			GingerbreadRacer                int `json:"gingerbread_racer"`
			GingerbreadBanker               int `json:"gingerbread_banker"`
			GingerbreadWinner               int `json:"gingerbread_winner"`
			VampirezZombieKiller            int `json:"vampirez_zombie_killer"`
			TntgamesPvpRunWins              int `json:"tntgames_pvp_run_wins"`
			TntgamesPvpRunKiller            int `json:"tntgames_pvp_run_killer"`
			TntgamesTntTagWins              int `json:"tntgames_tnt_tag_wins"`
			TntgamesTntWizardsKills         int `json:"tntgames_tnt_wizards_kills"`
			TntgamesTntTriathlon            int `json:"tntgames_tnt_triathlon"`
			TntgamesTntBanker               int `json:"tntgames_tnt_banker"`
			TntgamesTntWizardsCaps          int `json:"tntgames_tnt_wizards_caps"`
			QuakeHeadshots                  int `json:"quake_headshots"`
			QuakeWeaponArsenal              int `json:"quake_weapon_arsenal"`
			BlitzWinsTeams                  int `json:"blitz_wins_teams"`
			BlitzWarVeteran                 int `json:"blitz_war_veteran"`
			BlitzLooter                     int `json:"blitz_looter"`
			BedwarsWins                     int `json:"bedwars_wins"`
			BedwarsBeds                     int `json:"bedwars_beds"`
			BedwarsLevel                    int `json:"bedwars_level"`
			BedwarsLootBox                  int `json:"bedwars_loot_box"`
			MurdermysteryWinsAsSurvivor     int `json:"murdermystery_wins_as_survivor"`
			MurdermysteryKillsAsMurderer    int `json:"murdermystery_kills_as_murderer"`
			MurdermysteryWinsAsMurderer     int `json:"murdermystery_wins_as_murderer"`
			TruecombatKing                  int `json:"truecombat_king"`
			Halloween2017Pumpkinator        int `json:"halloween2017_pumpkinator"`
			Halloween2017PumpkinSmasher     int `json:"halloween2017_pumpkin_smasher"`
			Christmas2017NoChristmas        int `json:"christmas2017_no_christmas"`
			Christmas2017SantaSaysRounds    int `json:"christmas2017_santa_says_rounds"`
			Christmas2017Advent             int `json:"christmas2017_advent"`
			BuildbattleBuildbattlePoints    int `json:"buildbattle_buildbattle_points"`
			DuelsDuelsWinner                int `json:"duels_duels_winner"`
			BuildbattleGuessTheBuildGuesses int `json:"buildbattle_guess_the_build_guesses"`
			BuildbattleGuessTheBuildWinner  int `json:"buildbattle_guess_the_build_winner"`
			BuildbattleBuildBattleScore     int `json:"buildbattle_build_battle_score"`
			BuildbattleBuildBattleVoter     int `json:"buildbattle_build_battle_voter"`
			BuildbattleBuildBattlePoints    int `json:"buildbattle_build_battle_points"`
			DuelsDuelsWinStreak             int `json:"duels_duels_win_streak"`
			QuakeGodlikes                   int `json:"quake_godlikes"`
			BedwarsBedwarsKiller            int `json:"bedwars_bedwars_killer"`
			BedwarsCollectorsEdition        int `json:"bedwars_collectors_edition"`
			WarlordsCtfObjective            int `json:"warlords_ctf_objective"`
			SkywarsWinsLab                  int `json:"skywars_wins_lab"`
			MurdermysteryHoarder            int `json:"murdermystery_hoarder"`
			PaintballKillStreaks            int `json:"paintball_kill_streaks"`
			GingerbreadMystery              int `json:"gingerbread_mystery"`
			PaintballHatCollector           int `json:"paintball_hat_collector"`
			ArenaMagicalBox                 int `json:"arena_magical_box"`
			Christmas2017Advent2018         int `json:"christmas2017_advent_2018"`
			CopsandcrimsHeadshotKills       int `json:"copsandcrims_headshot_kills"`
			DuelsGoals                      int `json:"duels_goals"`
			DuelsBridgeDuelsWins            int `json:"duels_bridge_duels_wins"`
			DuelsBridgeWins                 int `json:"duels_bridge_wins"`
			DuelsBridgeDoublesWins          int `json:"duels_bridge_doubles_wins"`
			DuelsDuelsTraveller             int `json:"duels_duels_traveller"`
			DuelsBridgeWinStreak            int `json:"duels_bridge_win_streak"`
			DuelsUniqueMapWins              int `json:"duels_unique_map_wins"`
			UhcConsumer                     int `json:"uhc_consumer"`
			BlitzKitCollector               int `json:"blitz_kit_collector"`
			SkywarsHeads                    int `json:"skywars_heads"`
			ArcadeZombiesNiceShot           int `json:"arcade_zombies_nice_shot"`
			ArcadeZombiesRoundProgression   int `json:"arcade_zombies_round_progression"`
			ArcadeZombiesHighRound          int `json:"arcade_zombies_high_round"`
			EasterThrowEggs                 int `json:"easter_throw_eggs"`
			SkyclashTreasureHunter          int `json:"skyclash_treasure_hunter"`
			SkyclashMobBeheading            int `json:"skyclash_mob_beheading"`
			WarlordsDomObjective            int `json:"warlords_dom_objective"`
			BlitzKitExpert                  int `json:"blitz_kit_expert"`
			BlitzMobMaster                  int `json:"blitz_mob_master"`
			BlitzFightingExpert             int `json:"blitz_fighting_expert"`
			BlitzTreasureSeeker             int `json:"blitz_treasure_seeker"`
			BlitzKitExperienceCollector     int `json:"blitz_kit_experience_collector"`
			TntgamesBlockRunner             int `json:"tntgames_block_runner"`
			SkyblockMinionLover             int `json:"skyblock_minion_lover"`
			SkyblockExcavator               int `json:"skyblock_excavator"`
			SkyblockCombat                  int `json:"skyblock_combat"`
			SkyblockGatherer                int `json:"skyblock_gatherer"`
			SkyblockAngler                  int `json:"skyblock_angler"`
			SkyblockHarvester               int `json:"skyblock_harvester"`
			SkyblockTreasury                int `json:"skyblock_treasury"`
			SkyblockConcoctor               int `json:"skyblock_concoctor"`
			SkyblockAugmentation            int `json:"skyblock_augmentation"`
			TntgamesClinic                  int `json:"tntgames_clinic"`
			SpeeduhcPromotion               int `json:"speeduhc_promotion"`
			TruecombatFeelsLucky            int `json:"truecombat_feels_lucky"`
			Christmas2017Advent2019         int `json:"christmas2017_advent_2019"`
			Christmas2017PresentCollector   int `json:"christmas2017_present_collector"`
			SkyblockUniqueGifts             int `json:"skyblock_unique_gifts"`
			Christmas2017SecretSanta        int `json:"christmas2017_secret_santa"`
			EasterEggFinder                 int `json:"easter_egg_finder"`
			EasterMasterTracker             int `json:"easter_master_tracker"`
			PitContracts                    int `json:"pit_contracts"`
			PitGold                         int `json:"pit_gold"`
			PitKills                        int `json:"pit_kills"`
			PitPrestiges                    int `json:"pit_prestiges"`
			PitRenown                       int `json:"pit_renown"`
			SkyblockDomesticator            int `json:"skyblock_domesticator"`
			SkyblockSlayer                  int `json:"skyblock_slayer"`
			SummerTreasureHoarder           int `json:"summer_treasure_hoarder"`
			SummerTreasureMaster            int `json:"summer_treasure_master"`
			SkywarsNewDayNewChallenge       int `json:"skywars_new_day_new_challenge"`
			SkywarsYouReAStar               int `json:"skywars_you_re_a_star"`
			DuelsDuelsDivision              int `json:"duels_duels_division"`
			ArcadeCtwOhSheep                int `json:"arcade_ctw_oh_sheep"`
			ArcadeCtwSlayer                 int `json:"arcade_ctw_slayer"`
			SkyblockDungeoneer              int `json:"skyblock_dungeoneer"`
			SkyblockTreasureHunter          int `json:"skyblock_treasure_hunter"`
			SkyblockHardWorkingMiner        int `json:"skyblock_hard_working_miner"`
			SkyblockGoblinKiller            int `json:"skyblock_goblin_killer"`
		} `json:"achievements"`
		AchievementsOneTime []string `json:"achievementsOneTime"`
		AutoSpawnPet        bool     `json:"auto_spawn_pet"`
		Channel             string   `json:"channel"`
		Chat                bool     `json:"chat"`
		Disguise            string   `json:"disguise"`
		Displayname         string   `json:"displayname"`
		EulaCoins           bool     `json:"eulaCoins"`
		FireworkStorage     []struct {
			FlightDuration int    `json:"flight_duration"`
			Shape          string `json:"shape"`
			Trail          bool   `json:"trail"`
			Twinkle        bool   `json:"twinkle"`
			Colors         string `json:"colors"`
			FadeColors     string `json:"fade_colors"`
			Selected       bool   `json:"selected"`
		} `json:"fireworkStorage"`
		FirstLogin                 int64         `json:"firstLogin"`
		FriendRequests             []interface{} `json:"friendRequests"`
		FriendRequestsUUID         []interface{} `json:"friendRequestsUuid"`
		Gadget                     string        `json:"gadget"`
		Karma                      int           `json:"karma"`
		KnownAliases               []string      `json:"knownAliases"`
		KnownAliasesLower          []string      `json:"knownAliasesLower"`
		LastLogin                  int64         `json:"lastLogin"`
		MostRecentMinecraftVersion int           `json:"mostRecentMinecraftVersion"`
		MostRecentlyThanked        string        `json:"mostRecentlyThanked"`
		MostRecentlyTipped         string        `json:"mostRecentlyTipped"`
		MostRecentlyTippedUUID     string        `json:"mostRecentlyTippedUuid"`
		NetworkExp                 int           `json:"networkExp"`
		Notifications              bool          `json:"notifications"`
		PackageRank                string        `json:"packageRank"`
		ParkourCompletions         struct {
			ArcadeGames    Parkour `json:"ArcadeGames"`
			BlitzLobby     Parkour `json:"BlitzLobby"`
			MainLobby      Parkour `json:"MainLobby"`
			QuakeCraft     Parkour `json:"QuakeCraft"`
			Tnt            Parkour `json:"TNT"`
			Christmas      Parkour `json:"christmas"`
			Main           Parkour `json:"main"`
			Turbo          Parkour `json:"Turbo"`
			Skywars        Parkour `json:"Skywars"`
			Paintball      Parkour `json:"Paintball"`
			CopsnCrims     Parkour `json:"CopsnCrims"`
			Vampirez       Parkour `json:"vampirez"`
			TruePVPParkour Parkour `json:"TruePVPParkour"`
			TheWallsLobby  Parkour `json:"TheWallsLobby"`
			MegaWalls      Parkour `json:"MegaWalls"`
			Uhc            Parkour `json:"uhc"`
			NewMainLobby   Parkour `json:"NewMainLobby"`
			Prototype      Parkour `json:"Prototype"`
			Bedwars        Parkour `json:"Bedwars"`
			Warlords       Parkour `json:"Warlords"`
			SpeedUHC       Parkour `json:"SpeedUHC"`
			SkywarsAug2017 Parkour `json:"SkywarsAug2017"`
		} `json:"ParkourCompletions"`
		Playername string `json:"playername"`
		Quests     struct {
			Blitzerk struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"blitzerk"`
			ExplosiveGames struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"explosive_games"`
			Gladiator struct {
				Active struct {
					Objectives struct {
						FourV4 int `json:"4v4"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"gladiator"`
			Megawaller struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"megawaller"`
			NuggetWarriors struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Vampirez int `json:"vampirez"`
						Quake    int `json:"quake"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"nugget_warriors"`
			PaintballExpert struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Kill int `json:"kill"`
						Play int `json:"play"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"paintball_expert"`
			SerialKiller struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Quake     int `json:"quake"`
						Megawalls int `json:"megawalls"`
						Blitz     int `json:"blitz"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"serial_killer"`
			SpaceMission struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"space_mission"`
			TntAddict struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"tnt_addict"`
			Waller struct {
				Active struct {
					Objectives struct {
						Win int `json:"win"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"waller"`
			WarlordsCtf struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warlords_ctf"`
			WarlordsDedication struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						WarlordsWeeklyDedi int `json:"warlords_weekly_dedi"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warlords_dedication"`
			WarlordsDomination struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_domination"`
			WarlordsWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warlords_win"`
			WarriorsJourney struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Vampirezkillvamps int `json:"vampirezkillvamps"`
						Quake25Kill       int `json:"quake25kill"`
						Megawallswin      int `json:"megawallswin"`
						Blitzkill         int `json:"blitzkill"`
						Vampirezkillhuman int `json:"vampirezkillhuman"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warriors_journey"`
			WelcomeToHell struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Megawalls int `json:"megawalls"`
						Blitz     int `json:"blitz"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"welcome_to_hell"`
			GingerbreadMastery struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_mastery"`
			GingerbreadRacer struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_racer"`
			GingerbreadMaps struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_maps"`
			GingerbreadBlingBling struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_bling_bling"`
			SkywarsWeeklyKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						SkywarsWeeklyKills int `json:"skywars_weekly_kills"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_weekly_kills"`
			SkywarsTeamKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_team_kills"`
			SkywarsTeamWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_team_win"`
			SkywarsSoloKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						SkywarsSoloKills int `json:"skywars_solo_kills"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_solo_kills"`
			SkywarsSoloWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_solo_win"`
			UhcWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						UhcKills int `json:"uhc_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_weekly"`
			UhcDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"uhc_daily"`
			BlitzWeeklyMaster struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Killblitz10      int `json:"killblitz10"`
						BlitzGamesPlayed int `json:"blitz_games_played"`
						Winblitz         int `json:"winblitz"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_weekly_master"`
			BlitzKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Killblitz10 int `json:"killblitz10"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_kills"`
			BlitzWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_win"`
			BlitzGameOfTheDay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"blitz_game_of_the_day"`
			WarlordsTdm struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warlords_tdm"`
			MegaWallsWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						MegaWallsKillWeekly int `json:"mega_walls_kill_weekly"`
						MegaWallsPlayWeekly int `json:"mega_walls_play_weekly"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_weekly"`
			MegaWallsKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						MegaWallsKillDaily int `json:"mega_walls_kill_daily"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_kill"`
			MegaWallsWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_win"`
			MegaWallsPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_play"`
			QuakeWeeklyPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"quake_weekly_play"`
			QuakeDailyKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"quake_daily_kill"`
			QuakeDailyPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						QuakeDailyPlay int `json:"quake_daily_play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"quake_daily_play"`
			CrazyWallsDailyPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						CrazyWallsDailyPlay int `json:"crazy_walls_daily_play"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"crazy_walls_daily_play"`
			CrazyWallsDailyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"crazy_walls_daily_win"`
			CrazyWallsDailyKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						CrazyWallsDailyKill int `json:"crazy_walls_daily_kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"crazy_walls_daily_kill"`
			CrazyWallsWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"crazy_walls_weekly"`
			VampirezDailyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_win"`
			VampirezWeeklyKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						VampirezWeeklyKillVampire int `json:"vampirez_weekly_kill_vampire"`
						VampirezWeeklyKillZombie  int `json:"vampirez_weekly_kill_zombie"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_weekly_kill"`
			VampirezWeeklyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						VampirezWeeklyWinSurvivor int `json:"vampirez_weekly_win_survivor"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"vampirez_weekly_win"`
			VampirezDailyKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						VampirezDailyKillVampire int `json:"vampirez_daily_kill_vampire"`
						VampirezDailyKillZombie  int `json:"vampirez_daily_kill_zombie"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_kill"`
			VampirezDailyPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_play"`
			TntWeeklyPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntWeeklyPlay int `json:"tnt_weekly_play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_weekly_play"`
			TntDailyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_daily_win"`
			TntDailyPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"tnt_daily_play"`
			SupersmashSoloWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"supersmash_solo_win"`
			SupersmashWeeklyKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						SupersmashWeeklyKills int `json:"supersmash_weekly_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"supersmash_weekly_kills"`
			SupersmashTeamKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"supersmash_team_kills"`
			SupersmashSoloKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"supersmash_solo_kills"`
			SupersmashTeamWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"supersmash_team_win"`
			WallsWeekly struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						WallsWeeklyPlay  int `json:"walls_weekly_play"`
						WallsWeeklyKills int `json:"walls_weekly_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_weekly"`
			WallsDailyWin struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_daily_win"`
			WallsDailyKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_daily_kill"`
			WallsDailyPlay struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_daily_play"`
			ArcadeSpecialist struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Play int `json:"play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"arcade_specialist"`
			ArcadeWinner struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"arcade_winner"`
			ArcadeGamer struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Play int `json:"play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"arcade_gamer"`
			Paintballer struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"paintballer"`
			PaintballKiller struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Kill int `json:"kill"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"paintball_killer"`
			ArenaDailyPlay struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"arena_daily_play"`
			ArenaDailyWins struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"arena_daily_wins"`
			ArenaWeeklyPlay struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"arena_weekly_play"`
			ArenaDailyKills struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"arena_daily_kills"`
			NormalBrawler struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"normal_brawler"`
			UhcMadness struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Kill int `json:"kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_madness"`
			UhcAddict struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_addict"`
			InsaneBrawler struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"insane_brawler"`
			HuntingSeason struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"hunting_season"`
			CvcKillDailyNormal struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"cvc_kill_daily_normal"`
			CvcKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"cvc_kill"`
			CvcWinDailyDeathmatch struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"cvc_win_daily_deathmatch"`
			CvcWinDailyNormal struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"cvc_win_daily_normal"`
			CvcKillWeekly struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						CvcPlayWeekly  int `json:"cvc_play_weekly"`
						CvcPlayWeekly2 int `json:"cvc_play_weekly_2"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"cvc_kill_weekly"`
			SkyclashVoid struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skyclash_void"`
			SkyclashKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Kill int `json:"kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skyclash_kills"`
			SkyclashWeeklyKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Kill int `json:"kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skyclash_weekly_kills"`
			SkyclashPlayGames struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skyclash_play_games"`
			SkyclashPlayPoints struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skyclash_play_points"`
			UhcTeam struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						UhcGames int `json:"uhc_games"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_team"`
			UhcDm struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_dm"`
			UhcSolo struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_solo"`
			QuakeDailyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"quake_daily_win"`
			TntPvprunWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntPvprunWeekly int `json:"tnt_pvprun_weekly"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_pvprun_weekly"`
			TntPvprunDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntPvprunDaily int `json:"tnt_pvprun_daily"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_pvprun_daily"`
			TntTntrunDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_tntrun_daily"`
			TntBowspleefWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntBowspleefWeekly int `json:"tnt_bowspleef_weekly"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_bowspleef_weekly"`
			TntWizardsDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_wizards_daily"`
			TntTnttagDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_tnttag_daily"`
			TntBowspleefDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntBowspleefDaily int `json:"tnt_bowspleef_daily"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_bowspleef_daily"`
			TntTntrunWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntTntrunWeekly int `json:"tnt_tntrun_weekly"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_tntrun_weekly"`
			TntWizardsWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntWizardsWeeklyKills int `json:"tnt_wizards_weekly_kills"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_wizards_weekly"`
			TntTnttagWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						TntTnttagWeekly int `json:"tnt_tnttag_weekly"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"tnt_tnttag_weekly"`
			BedwarsDailyOneMore struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						BedwarsDailyPlayed int `json:"bedwars_daily_played"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_daily_one_more"`
			BedwarsDailyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_daily_win"`
			BedwarsWeeklyBedElims struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						BedwarsBedElims int `json:"bedwars_bed_elims"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"bedwars_weekly_bed_elims"`
			SkywarsWeeklyArcadeWinAll struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						SkywarsArcadeWeeklyWin int `json:"skywars_arcade_weekly_win"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_weekly_arcade_win_all"`
			SkywarsArcadeWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_arcade_win"`
			MmWeeklyWins struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						MmWeeklyWin int `json:"mm_weekly_win"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"mm_weekly_wins"`
			MmDailyTargetKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mm_daily_target_kill"`
			MmDailyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"mm_daily_win"`
			MmWeeklyMurdererKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"mm_weekly_murderer_kills"`
			MmDailyPowerPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"mm_daily_power_play"`
			BedwarsWeeklyPumpkinator struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"bedwars_weekly_pumpkinator"`
			BedwarsWeeklySanta struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						BedwarsSpecialWeeklySanta int `json:"bedwars_special_weekly_santa"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"bedwars_weekly_santa"`
			SkywarsSpecialNorthPole struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_special_north_pole"`
			BlitzSpecialDailyNorthPole struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_special_daily_north_pole"`
			UhcWeeklySpecialCookie struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_weekly_special_cookie"`
			MmSpecialWeeklySanta struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mm_special_weekly_santa"`
			SkywarsWeeklyFreeLootChest struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"skywars_weekly_free_loot_chest"`
			SkywarsDailyTokens struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_daily_tokens"`
			SkywarsWeeklyHardChest struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						SkywarsWeeklyHardLootChest int `json:"skywars_weekly_hard_loot_chest"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_weekly_hard_chest"`
			DuelsPlayer struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Play int `json:"play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"duels_player"`
			DuelsWinner struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"duels_winner"`
			DuelsKiller struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"duels_killer"`
			DuelsWeeklyWins struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"duels_weekly_wins"`
			DuelsWeeklyKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Kill int `json:"kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"duels_weekly_kills"`
			BuildBattleWeekly struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Play int `json:"play"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"build_battle_weekly"`
			BuildBattleWinner struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_winner"`
			BuildBattlePlayer struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Play int `json:"play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_player"`
			PrototypePitWeeklyGold struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						PrototypePitWeeklyGold int `json:"prototype_pit_weekly_gold"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"prototype_pit_weekly_gold"`
			PrototypePitDailyContract struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"prototype_pit_daily_contract"`
			PrototypePitDailyKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Kill int `json:"kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"prototype_pit_daily_kills"`
			SkywarsMegaDoublesWins struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_mega_doubles_wins"`
			BedwarsWeeklyDreamWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						BedwarsDreamWins int `json:"bedwars_dream_wins"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_weekly_dream_win"`
			MegaWallsFaithful struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_faithful"`
			WarlordsAllStar struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						WarlordsWeeklyDamage int `json:"warlords_weekly_damage"`
						WarlordsWeeklyHeal   int `json:"warlords_weekly_heal"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warlords_all_star"`
			WarlordsObjectives struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_objectives"`
			WarlordsVictorious struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_victorious"`
			BedwarsDailyNightmares struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"bedwars_daily_nightmares"`
			SkywarsHalloweenHarvest struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						SkywarsHalloweenKills int `json:"skywars_halloween_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_halloween_harvest"`
			BedwarsDailyGifts struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						BedwarsDailySpecialChristmasGifts int `json:"bedwars_daily_special_christmas_gifts"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_daily_gifts"`
			SkywarsCorruptWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_corrupt_win"`
			BuildBattleChristmas struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"build_battle_christmas"`
			BlitzWeeklyChaosMaster struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_weekly_chaos_master"`
			BlitzWinChaos struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_win_chaos"`
			VampirezWeeklyHumanKill struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						VampirezWeeklyKillSurvivor int `json:"vampirez_weekly_kill_survivor"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"vampirez_weekly_human_kill"`
			VampirezDailyHumanKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						VampirezDailyKillHuman int `json:"vampirez_daily_kill_human"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_human_kill"`
			BlitzLootChestDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Lootchestblitz int `json:"lootchestblitz"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_loot_chest_daily"`
			BlitzLootChestWeekly struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Lootchestblitz  int `json:"lootchestblitz"`
						Dealdamageblitz int `json:"dealdamageblitz"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_loot_chest_weekly"`
			SoloBrawler struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"solo_brawler"`
			TeamBrawler struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"team_brawler"`
			SkywarsHalloweenHarvest2019 struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_halloween_harvest_2019"`
			BuildBattleHalloween struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_halloween"`
			MmSpecialWeeklyKillerInstinct2019 struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"mm_special_weekly_killer_instinct_2019"`
			TntWeeklySpecial struct {
				Active struct {
					Objectives struct {
						TntWeeklyCandyCanes int `json:"tnt_weekly_candy_canes"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_weekly_special"`
			BuildBattleChristmasWeekly struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_christmas_weekly"`
			TntLunarQuest struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_lunar_quest"`
			SkywarsLunarQuest struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_lunar_quest"`
		} `json:"quests"`
		SeeRequests bool `json:"seeRequests"`
		Stats       struct {
			Arcade struct {
				BountyKillsOneinthequiver                     int    `json:"bounty_kills_oneinthequiver"`
				Coins                                         int    `json:"coins"`
				DeathsOneinthequiver                          int    `json:"deaths_oneinthequiver"`
				DeathsThrowOut                                int    `json:"deaths_throw_out"`
				Flash                                         bool   `json:"flash"`
				HeadshotsDayone                               int    `json:"headshots_dayone"`
				KillsDayone                                   int    `json:"kills_dayone"`
				KillsDragonwars2                              int    `json:"kills_dragonwars2"`
				KillsOneinthequiver                           int    `json:"kills_oneinthequiver"`
				KillsThrowOut                                 int    `json:"kills_throw_out"`
				MaxWave                                       int    `json:"max_wave"`
				PoopCollected                                 int    `json:"poop_collected"`
				StampLevel                                    int    `json:"stamp_level"`
				TimeStamp                                     int64  `json:"time_stamp"`
				WinsDayone                                    int    `json:"wins_dayone"`
				WinsDragonwars2                               int    `json:"wins_dragonwars2"`
				WinsFarmHunt                                  int    `json:"wins_farm_hunt"`
				WinsOneinthequiver                            int    `json:"wins_oneinthequiver"`
				WinsParty                                     int    `json:"wins_party"`
				WinsParty2                                    int    `json:"wins_party_2"`
				SwKills                                       int    `json:"sw_kills"`
				SwShotsFired                                  int    `json:"sw_shots_fired"`
				SwWeeklyKillsB                                int    `json:"sw_weekly_kills_b"`
				SwRebelKills                                  int    `json:"sw_rebel_kills"`
				SwDeaths                                      int    `json:"sw_deaths"`
				SwMonthlyKillsA                               int    `json:"sw_monthly_kills_a"`
				WinsThrowOut                                  int    `json:"wins_throw_out"`
				WinsEnder                                     int    `json:"wins_ender"`
				WinsParty3                                    int    `json:"wins_party_3"`
				HitwRecordQ                                   int    `json:"hitw_record_q"`
				HitwRecordF                                   int    `json:"hitw_record_f"`
				RoundsHoleInTheWall                           int    `json:"rounds_hole_in_the_wall"`
				WinsHoleInTheWall                             int    `json:"wins_hole_in_the_wall"`
				SwWeeklyKillsA                                int    `json:"sw_weekly_kills_a"`
				SwEmpireKills                                 int    `json:"sw_empire_kills"`
				SwGameWins                                    int    `json:"sw_game_wins"`
				WinsSantaSays                                 int    `json:"wins_santa_says"`
				RoundsSantaSays                               int    `json:"rounds_santa_says"`
				RoundsSimonSays                               int    `json:"rounds_simon_says"`
				WeeklyCoinsB                                  int    `json:"weekly_coins_b"`
				WinsSimonSays                                 int    `json:"wins_simon_says"`
				MonthlyCoinsA                                 int    `json:"monthly_coins_a"`
				WeeklyCoinsA                                  int    `json:"weekly_coins_a"`
				MonthlyCoinsB                                 int    `json:"monthly_coins_b"`
				KillsMiniWalls                                int    `json:"kills_mini_walls"`
				DeathsMiniWalls                               int    `json:"deaths_mini_walls"`
				WitherKillsMiniWalls                          int    `json:"wither_kills_mini_walls"`
				FinalKillsMiniWalls                           int    `json:"final_kills_mini_walls"`
				PowerkicksSoccer                              int    `json:"powerkicks_soccer"`
				GoalsSoccer                                   int    `json:"goals_soccer"`
				WinsSoccer                                    int    `json:"wins_soccer"`
				MiniwallsActiveKit                            string `json:"miniwalls_activeKit"`
				ArrowsHitMiniWalls                            int    `json:"arrows_hit_mini_walls"`
				WinsMiniWalls                                 int    `json:"wins_mini_walls"`
				ArrowsShotMiniWalls                           int    `json:"arrows_shot_mini_walls"`
				WitherDamageMiniWalls                         int    `json:"wither_damage_mini_walls"`
				Dec2016Achievements                           bool   `json:"dec2016_achievements"`
				Dec2016Achievements2                          bool   `json:"dec2016_achievements2"`
				WinsDrawTheirThing                            int    `json:"wins_draw_their_thing"`
				Hints                                         bool   `json:"hints"`
				KicksSoccer                                   int    `json:"kicks_soccer"`
				WinsHypixelSports                             int    `json:"wins_hypixel_sports"`
				FastestTime20ZombiesAlienarcadiumNormal       int    `json:"fastest_time_20_zombies_alienarcadium_normal"`
				BestRoundZombiesAlienarcadium                 int    `json:"best_round_zombies_alienarcadium"`
				BestRoundZombiesAlienarcadiumNormal           int    `json:"best_round_zombies_alienarcadium_normal"`
				FastestTime20Zombies                          int    `json:"fastest_time_20_zombies"`
				BestRoundZombies                              int    `json:"best_round_zombies"`
				FastestTime10ZombiesAlienarcadiumNormal       int    `json:"fastest_time_10_zombies_alienarcadium_normal"`
				FastestTime10Zombies                          int    `json:"fastest_time_10_zombies"`
				RainbowZombieKillsZombies                     int    `json:"rainbow_zombie_kills_zombies"`
				GiantZombieKillsZombies                       int    `json:"giant_zombie_kills_zombies"`
				SpaceBlasterZombieKillsZombies                int    `json:"space_blaster_zombie_kills_zombies"`
				ChglugluZombieKillsZombies                    int    `json:"chgluglu_zombie_kills_zombies"`
				TimesKnockedDownZombiesAlienarcadium          int    `json:"times_knocked_down_zombies_alienarcadium"`
				WindowsRepairedZombiesAlienarcadium           int    `json:"windows_repaired_zombies_alienarcadium"`
				TimesKnockedDownZombies                       int    `json:"times_knocked_down_zombies"`
				ZombieKillsZombiesAlienarcadiumNormal         int    `json:"zombie_kills_zombies_alienarcadium_normal"`
				ZombieKillsZombies                            int    `json:"zombie_kills_zombies"`
				GhastZombieKillsZombies                       int    `json:"ghast_zombie_kills_zombies"`
				WormSmallZombieKillsZombies                   int    `json:"worm_small_zombie_kills_zombies"`
				DeathsZombiesAlienarcadium                    int    `json:"deaths_zombies_alienarcadium"`
				BulletsShotZombies                            int    `json:"bullets_shot_zombies"`
				BasicZombieKillsZombies                       int    `json:"basic_zombie_kills_zombies"`
				MegaMagmaZombieKillsZombies                   int    `json:"mega_magma_zombie_kills_zombies"`
				WindowsRepairedZombiesAlienarcadiumNormal     int    `json:"windows_repaired_zombies_alienarcadium_normal"`
				BulletsHitZombies                             int    `json:"bullets_hit_zombies"`
				WormZombieKillsZombies                        int    `json:"worm_zombie_kills_zombies"`
				IronGolemZombieKillsZombies                   int    `json:"iron_golem_zombie_kills_zombies"`
				HeadshotsZombies                              int    `json:"headshots_zombies"`
				PlayersRevivedZombies                         int    `json:"players_revived_zombies"`
				TotalRoundsSurvivedZombiesAlienarcadiumNormal int    `json:"total_rounds_survived_zombies_alienarcadium_normal"`
				DeathsZombies                                 int    `json:"deaths_zombies"`
				TimesKnockedDownZombiesAlienarcadiumNormal    int    `json:"times_knocked_down_zombies_alienarcadium_normal"`
				TotalRoundsSurvivedZombiesAlienarcadium       int    `json:"total_rounds_survived_zombies_alienarcadium"`
				WindowsRepairedZombies                        int    `json:"windows_repaired_zombies"`
				TotalRoundsSurvivedZombies                    int    `json:"total_rounds_survived_zombies"`
				PlayersRevivedZombiesAlienarcadium            int    `json:"players_revived_zombies_alienarcadium"`
				SentinelZombieKillsZombies                    int    `json:"sentinel_zombie_kills_zombies"`
				ZombieKillsZombiesAlienarcadium               int    `json:"zombie_kills_zombies_alienarcadium"`
				PigZombieZombieKillsZombies                   int    `json:"pig_zombie_zombie_kills_zombies"`
				DeathsZombiesAlienarcadiumNormal              int    `json:"deaths_zombies_alienarcadium_normal"`
				SpaceGruntZombieKillsZombies                  int    `json:"space_grunt_zombie_kills_zombies"`
				SkeletonZombieKillsZombies                    int    `json:"skeleton_zombie_kills_zombies"`
				BlobZombieKillsZombies                        int    `json:"blob_zombie_kills_zombies"`
				ClownZombieKillsZombies                       int    `json:"clown_zombie_kills_zombies"`
				PlayersRevivedZombiesAlienarcadiumNormal      int    `json:"players_revived_zombies_alienarcadium_normal"`
				MegaBlobZombieKillsZombies                    int    `json:"mega_blob_zombie_kills_zombies"`
				DeliveredSantaSimulator                       int    `json:"delivered_santa_simulator"`
				SpottedSantaSimulator                         int    `json:"spotted_santa_simulator"`
				EggsFoundEasterSimulator                      int    `json:"eggs_found_easter_simulator"`
				WinsEasterSimulator                           int    `json:"wins_easter_simulator"`
				ItemsFoundScubaSimulator                      int    `json:"items_found_scuba_simulator"`
				TotalPointsScubaSimulator                     int    `json:"total_points_scuba_simulator"`
				WinsScubaSimulator                            int    `json:"wins_scuba_simulator"`
				WinsGrinchSimulatorV2                         int    `json:"wins_grinch_simulator_v2"`
				GiftsGrinchSimulatorV2                        int    `json:"gifts_grinch_simulator_v2"`
				LastTourneyAd                                 int64  `json:"lastTourneyAd"`
			} `json:"Arcade"`
			Arena struct {
				ActiveRune      string   `json:"active_rune"`
				Coins           int      `json:"coins"`
				CoinsSpent      int      `json:"coins_spent"`
				Damage2V2       int      `json:"damage_2v2"`
				Damage4V4       int      `json:"damage_4v4"`
				DamageFfa       int      `json:"damage_ffa"`
				Deaths2V2       int      `json:"deaths_2v2"`
				Deaths4V4       int      `json:"deaths_4v4"`
				Games2V2        int      `json:"games_2v2"`
				Games4V4        int      `json:"games_4v4"`
				GamesFfa        int      `json:"games_ffa"`
				Hat             string   `json:"hat"`
				Healed2V2       int      `json:"healed_2v2"`
				Healed4V4       int      `json:"healed_4v4"`
				HealedFfa       int      `json:"healed_ffa"`
				Keys            int      `json:"keys"`
				Kills2V2        int      `json:"kills_2v2"`
				Kills4V4        int      `json:"kills_4v4"`
				KillsFfa        int      `json:"kills_ffa"`
				Losses2V2       int      `json:"losses_2v2"`
				Losses4V4       int      `json:"losses_4v4"`
				LvlCooldown     int      `json:"lvl_cooldown"`
				LvlDamage       int      `json:"lvl_damage"`
				LvlEnergy       int      `json:"lvl_energy"`
				LvlHealth       int      `json:"lvl_health"`
				MagicalChest    int      `json:"magical_chest"`
				Offensive       string   `json:"offensive"`
				Packages        []string `json:"packages"`
				Rating          float64  `json:"rating"`
				Support         string   `json:"support"`
				Utility         string   `json:"utility"`
				WinStreaks2V2   int      `json:"win_streaks_2v2"`
				WinStreaks4V4   int      `json:"win_streaks_4v4"`
				WinStreaksFfa   int      `json:"win_streaks_ffa"`
				Wins2V2         int      `json:"wins_2v2"`
				Wins4V4         int      `json:"wins_4v4"`
				WinsFfa         int      `json:"wins_ffa"`
				RuneLevelEnergy int      `json:"rune_level_energy"`
				SelectedSword   string   `json:"selected_sword"`
				Ultimate        string   `json:"ultimate"`
				RuneLevelDamage int      `json:"rune_level_damage"`
				Wins            int      `json:"wins"`
			} `json:"Arena"`
			Battleground struct {
				Assists                  int      `json:"assists"`
				BerserkerPlays           int      `json:"berserker_plays"`
				BrokenInventory          int      `json:"broken_inventory"`
				ChosenClass              string   `json:"chosen_class"`
				Coins                    int      `json:"coins"`
				CurrentWeapon            int64    `json:"current_weapon"`
				Damage                   int      `json:"damage"`
				DamageBerserker          int      `json:"damage_berserker"`
				DamageDefender           int      `json:"damage_defender"`
				DamagePrevented          int      `json:"damage_prevented"`
				DamagePreventedBerserker int      `json:"damage_prevented_berserker"`
				DamagePreventedDefender  int      `json:"damage_prevented_defender"`
				DamagePreventedWarrior   int      `json:"damage_prevented_warrior"`
				DamageTaken              int      `json:"damage_taken"`
				DamageWarrior            int      `json:"damage_warrior"`
				Deaths                   int      `json:"deaths"`
				DefenderPlays            int      `json:"defender_plays"`
				FlagConquerSelf          int      `json:"flag_conquer_self"`
				FlagConquerTeam          int      `json:"flag_conquer_team"`
				Heal                     int      `json:"heal"`
				HealBerserker            int      `json:"heal_berserker"`
				HealDefender             int      `json:"heal_defender"`
				HealWarrior              int      `json:"heal_warrior"`
				Kills                    int      `json:"kills"`
				LifeLeeched              int      `json:"life_leeched"`
				LifeLeechedBerserker     int      `json:"life_leeched_berserker"`
				LifeLeechedWarrior       int      `json:"life_leeched_warrior"`
				Losses                   int      `json:"losses"`
				LossesDefender           int      `json:"losses_defender"`
				LossesWarrior            int      `json:"losses_warrior"`
				MageArmorSelection       int      `json:"mage_armor_selection"`
				MageSpec                 string   `json:"mage_spec"`
				MagicDust                int      `json:"magic_dust"`
				Packages                 []string `json:"packages"`
				PaladinSpec              string   `json:"paladin_spec"`
				Repaired                 int      `json:"repaired"`
				RepairedCommon           int      `json:"repaired_common"`
				RepairedEpic             int      `json:"repaired_epic"`
				RepairedRare             int      `json:"repaired_rare"`
				SelectedMount            string   `json:"selected_mount"`
				UpgradeRandom            int      `json:"upgrade_random"`
				UpgradeRandomEpic        int      `json:"upgrade_random_epic"`
				VoidShards               int      `json:"void_shards"`
				WarriorArmorSelection    int      `json:"warrior_armor_selection"`
				WarriorCooldown          int      `json:"warrior_cooldown"`
				WarriorCritchance        int      `json:"warrior_critchance"`
				WarriorCritmultiplier    int      `json:"warrior_critmultiplier"`
				WarriorEnergy            int      `json:"warrior_energy"`
				WarriorHealth            int      `json:"warrior_health"`
				WarriorHelmetSelection   int      `json:"warrior_helmet_selection"`
				WarriorPlays             int      `json:"warrior_plays"`
				WarriorSkill1            int      `json:"warrior_skill1"`
				WarriorSkill2            int      `json:"warrior_skill2"`
				WarriorSkill3            int      `json:"warrior_skill3"`
				WarriorSkill4            int      `json:"warrior_skill4"`
				WarriorSkill5            int      `json:"warrior_skill5"`
				WarriorSpec              string   `json:"warrior_spec"`
				WeaponInventory          []struct {
					Spec struct {
						Spec        int `json:"spec"`
						PlayerClass int `json:"playerClass"`
					} `json:"spec"`
					Ability      int    `json:"ability"`
					AbilityBoost int    `json:"abilityBoost"`
					Damage       int    `json:"damage"`
					Energy       int    `json:"energy"`
					Chance       int    `json:"chance"`
					Multiplier   int    `json:"multiplier"`
					Health       int    `json:"health"`
					Cooldown     int    `json:"cooldown"`
					Movement     int    `json:"movement"`
					Material     string `json:"material"`
					ID           int64  `json:"id"`
					Category     string `json:"category"`
					Crafted      bool   `json:"crafted"`
					UpgradeMax   int    `json:"upgradeMax"`
					UpgradeTimes int    `json:"upgradeTimes"`
					PlayStreak   bool   `json:"playStreak,omitempty"`
				} `json:"weapon_inventory"`
				WinStreak                  int    `json:"win_streak"`
				Wins                       int    `json:"wins"`
				WinsBerserker              int    `json:"wins_berserker"`
				WinsBlu                    int    `json:"wins_blu"`
				WinsCapturetheflag         int    `json:"wins_capturetheflag"`
				WinsCapturetheflagBlu      int    `json:"wins_capturetheflag_blu"`
				WinsCapturetheflagRed      int    `json:"wins_capturetheflag_red"`
				WinsDefender               int    `json:"wins_defender"`
				WinsDomination             int    `json:"wins_domination"`
				WinsDominationBlu          int    `json:"wins_domination_blu"`
				WinsDominationRed          int    `json:"wins_domination_red"`
				WinsRed                    int    `json:"wins_red"`
				WinsWarrior                int    `json:"wins_warrior"`
				SalvagedDustReward         int    `json:"salvaged_dust_reward"`
				SalvagedWeaponsRare        int    `json:"salvaged_weapons_rare"`
				SalvagedWeapons            int    `json:"salvaged_weapons"`
				SalvagedShardsReward       int    `json:"salvaged_shards_reward"`
				SalvagedWeaponsCommon      int    `json:"salvaged_weapons_common"`
				HealMage                   int    `json:"heal_mage"`
				DamagePyromancer           int    `json:"damage_pyromancer"`
				DamagePreventedMage        int    `json:"damage_prevented_mage"`
				HealPyromancer             int    `json:"heal_pyromancer"`
				DamagePreventedPyromancer  int    `json:"damage_prevented_pyromancer"`
				PyromancerPlays            int    `json:"pyromancer_plays"`
				DamageMage                 int    `json:"damage_mage"`
				MagePlays                  int    `json:"mage_plays"`
				SalvagedWeaponsEpic        int    `json:"salvaged_weapons_epic"`
				LossesBerserker            int    `json:"losses_berserker"`
				Crafted                    int    `json:"crafted"`
				CraftedEpic                int    `json:"crafted_epic"`
				ShamanSpec                 string `json:"shaman_spec"`
				PlayStreak                 int    `json:"play_streak"`
				DamagePreventedCryomancer  int    `json:"damage_prevented_cryomancer"`
				HealCryomancer             int    `json:"heal_cryomancer"`
				WinsMage                   int    `json:"wins_mage"`
				DamageCryomancer           int    `json:"damage_cryomancer"`
				CryomancerPlays            int    `json:"cryomancer_plays"`
				WinsCryomancer             int    `json:"wins_cryomancer"`
				WinsTeamdeathmatch         int    `json:"wins_teamdeathmatch"`
				WinsTeamdeathmatchRed      int    `json:"wins_teamdeathmatch_red"`
				WinsTeamdeathmatchA        int    `json:"wins_teamdeathmatch_a"`
				LossesCryomancer           int    `json:"losses_cryomancer"`
				LossesMage                 int    `json:"losses_mage"`
				WinsTeamdeathmatchBlu      int    `json:"wins_teamdeathmatch_blu"`
				WinsTeamdeathmatchB        int    `json:"wins_teamdeathmatch_b"`
				RepairedLegendary          int    `json:"repaired_legendary"`
				SalvagedWeaponsLegendary   int    `json:"salvaged_weapons_legendary"`
				WinsPyromancer             int    `json:"wins_pyromancer"`
				LossesPyromancer           int    `json:"losses_pyromancer"`
				EarthwardenPlays           int    `json:"earthwarden_plays"`
				HealEarthwarden            int    `json:"heal_earthwarden"`
				HealShaman                 int    `json:"heal_shaman"`
				WinsShaman                 int    `json:"wins_shaman"`
				DamagePreventedEarthwarden int    `json:"damage_prevented_earthwarden"`
				WinsEarthwarden            int    `json:"wins_earthwarden"`
				DamageShaman               int    `json:"damage_shaman"`
				DamageEarthwarden          int    `json:"damage_earthwarden"`
				DamagePreventedShaman      int    `json:"damage_prevented_shaman"`
				ShamanPlays                int    `json:"shaman_plays"`
				Penalty                    int    `json:"penalty"`
				AfkWarned                  int    `json:"afk_warned"`
				MageHelmetSelection        int    `json:"mage_helmet_selection"`
				PaladinArmorSelection      int    `json:"paladin_armor_selection"`
				PaladinHelmetSelection     int    `json:"paladin_helmet_selection"`
				ShamanArmorSelection       int    `json:"shaman_armor_selection"`
				ShamanHelmetSelection      int    `json:"shaman_helmet_selection"`
				DamageAquamancer           int    `json:"damage_aquamancer"`
				HealAquamancer             int    `json:"heal_aquamancer"`
				AquamancerPlays            int    `json:"aquamancer_plays"`
				DamagePreventedAquamancer  int    `json:"damage_prevented_aquamancer"`
				LossesAquamancer           int    `json:"losses_aquamancer"`
				WinsAquamancer             int    `json:"wins_aquamancer"`
				DamagePreventedProtector   int    `json:"damage_prevented_protector"`
				ProtectorPlays             int    `json:"protector_plays"`
				DamagePreventedPaladin     int    `json:"damage_prevented_paladin"`
				WinsProtector              int    `json:"wins_protector"`
				HealProtector              int    `json:"heal_protector"`
				WinsPaladin                int    `json:"wins_paladin"`
				PaladinPlays               int    `json:"paladin_plays"`
				HealPaladin                int    `json:"heal_paladin"`
				DamageProtector            int    `json:"damage_protector"`
				DamagePaladin              int    `json:"damage_paladin"`
				AvengerPlays               int    `json:"avenger_plays"`
				WinsAvenger                int    `json:"wins_avenger"`
				HealAvenger                int    `json:"heal_avenger"`
				DamageAvenger              int    `json:"damage_avenger"`
				DamagePreventedAvenger     int    `json:"damage_prevented_avenger"`
				LossesAvenger              int    `json:"losses_avenger"`
				LossesPaladin              int    `json:"losses_paladin"`
				CraftedRare                int    `json:"crafted_rare"`
				CrusaderPlays              int    `json:"crusader_plays"`
				HealCrusader               int    `json:"heal_crusader"`
				WinsCrusader               int    `json:"wins_crusader"`
				DamageCrusader             int    `json:"damage_crusader"`
				DamagePreventedCrusader    int    `json:"damage_prevented_crusader"`
				WinsDominationB            int    `json:"wins_domination_b"`
				WinsCapturetheflagB        int    `json:"wins_capturetheflag_b"`
				MageSkill5                 int    `json:"mage_skill5"`
				MageSkill4                 int    `json:"mage_skill4"`
				MageSkill3                 int    `json:"mage_skill3"`
				MageSkill2                 int    `json:"mage_skill2"`
				MageSkill1                 int    `json:"mage_skill1"`
				MageCritmultiplier         int    `json:"mage_critmultiplier"`
				MageCritchance             int    `json:"mage_critchance"`
				MageCooldown               int    `json:"mage_cooldown"`
				MageEnergy                 int    `json:"mage_energy"`
				MageHealth                 int    `json:"mage_health"`
				CraftedLegendary           int    `json:"crafted_legendary"`
				RerollLegendary            int    `json:"reroll_legendary"`
				Reroll                     int    `json:"reroll"`
				LossesCrusader             int    `json:"losses_crusader"`
				DamagePreventedThunderlord int    `json:"damage_prevented_thunderlord"`
				HealThunderlord            int    `json:"heal_thunderlord"`
				WinsThunderlord            int    `json:"wins_thunderlord"`
				ThunderlordPlays           int    `json:"thunderlord_plays"`
				DamageThunderlord          int    `json:"damage_thunderlord"`
				UpgradeCraftedLegendary    int    `json:"upgrade_crafted_legendary"`
				UpgradeCrafted             int    `json:"upgrade_crafted"`
				WinsCapturetheflagA        int    `json:"wins_capturetheflag_a"`
				WinsDominationA            int    `json:"wins_domination_a"`
				RewardInventory            int    `json:"reward_inventory"`
				ShamanSkill5               int    `json:"shaman_skill5"`
				ShamanSkill4               int    `json:"shaman_skill4"`
				ShamanSkill3               int    `json:"shaman_skill3"`
				ShamanSkill2               int    `json:"shaman_skill2"`
				ShamanSkill1               int    `json:"shaman_skill1"`
				ShamanCritmultiplier       int    `json:"shaman_critmultiplier"`
				ShamanCooldown             int    `json:"shaman_cooldown"`
				ShamanEnergy               int    `json:"shaman_energy"`
				ShamanHealth               int    `json:"shaman_health"`
				ShamanCritchance           int    `json:"shaman_critchance"`
				PaladinSkill5              int    `json:"paladin_skill5"`
				PaladinSkill4              int    `json:"paladin_skill4"`
				PaladinSkill3              int    `json:"paladin_skill3"`
				PaladinSkill2              int    `json:"paladin_skill2"`
				PaladinSkill1              int    `json:"paladin_skill1"`
				PaladinCritmultiplier      int    `json:"paladin_critmultiplier"`
				PaladinCritchance          int    `json:"paladin_critchance"`
				PaladinCooldown            int    `json:"paladin_cooldown"`
				PaladinEnergy              int    `json:"paladin_energy"`
				PaladinHealth              int    `json:"paladin_health"`
				KillsDomination            int    `json:"kills_domination"`
				DomPointCaptures           int    `json:"dom_point_captures"`
				TotalDominationScore       int    `json:"total_domination_score"`
				PowerupsCollected          int    `json:"powerups_collected"`
				DomPointDefends            int    `json:"dom_point_defends"`
				KillsTeamdeathmatch        int    `json:"kills_teamdeathmatch"`
				KillsCapturetheflag        int    `json:"kills_capturetheflag"`
				FlagReturns                int    `json:"flag_returns"`
				MvpCount                   int    `json:"mvp_count"`
				LossesThunderlord          int    `json:"losses_thunderlord"`
				LossesShaman               int    `json:"losses_shaman"`
				Autostrikemode             bool   `json:"autostrikemode"`
				EnergyPowerups             bool   `json:"energyPowerups"`
				Hints                      bool   `json:"hints"`
				BoundWeapon                struct {
					Warrior struct {
					} `json:"warrior"`
				} `json:"bound_weapon"`
				Simplifiedresourcepack bool `json:"simplifiedresourcepack"`
			} `json:"Battleground"`
			HungerGames struct {
				Arachnologist      int      `json:"arachnologist"`
				Archer             int      `json:"archer"`
				Armorer            int      `json:"armorer"`
				Astronaut          int      `json:"astronaut"`
				Aura               string   `json:"aura"`
				Baker              int      `json:"baker"`
				Blaze              int      `json:"blaze"`
				Blood              bool     `json:"blood"`
				ChosenTaunt        string   `json:"chosen_taunt"`
				Coins              int      `json:"coins"`
				Creepertamer       int      `json:"creepertamer"`
				Deaths             int      `json:"deaths"`
				Fisherman          int      `json:"fisherman"`
				Horsetamer         int      `json:"horsetamer"`
				Kills              int      `json:"kills"`
				Knight             int      `json:"knight"`
				Meatmaster         int      `json:"meatmaster"`
				Necromancer        int      `json:"necromancer"`
				Packages           []string `json:"packages"`
				Reddragon          int      `json:"reddragon"`
				Rogue              int      `json:"rogue"`
				Scout              int      `json:"scout"`
				Slimeyslime        int      `json:"slimeyslime"`
				Snowman            int      `json:"snowman"`
				Speleologist       int      `json:"speleologist"`
				Tim                int      `json:"tim"`
				Toxicologist       int      `json:"toxicologist"`
				Troll              int      `json:"troll"`
				Wins               int      `json:"wins"`
				Wolftamer          int      `json:"wolftamer"`
				WinsTeams          int      `json:"wins_teams"`
				WeeklyKillsA       int      `json:"weekly_kills_a"`
				MonthlyKillsB      int      `json:"monthly_kills_b"`
				RamboWin           int      `json:"rambo_win"`
				Paladin            int      `json:"paladin"`
				WeeklyKillsB       int      `json:"weekly_kills_b"`
				MonthlyKillsA      int      `json:"monthly_kills_a"`
				VotesMiradorBasin  int      `json:"votes_Mirador Basin"`
				Hunter             int      `json:"hunter"`
				RamboWins          int      `json:"rambo_wins"`
				ChosenVictorydance string   `json:"chosen_victorydance"`
				HunterInventory    struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
					Num2 string `json:"2"`
					Num3 string `json:"3"`
					Num4 string `json:"4"`
					Num8 string `json:"8"`
					Num9 string `json:"9"`
				} `json:"HunterInventory"`
				ArmorerInventory struct {
					Num8 string `json:"8"`
				} `json:"ArmorerInventory"`
				AstronautInventory struct {
					Num0 string `json:"0"`
					Num2 string `json:"2"`
					Num8 string `json:"8"`
				} `json:"AstronautInventory"`
				Defaultkit      string `json:"defaultkit"`
				ArcherInventory struct {
					Num1 string `json:"1"`
					Num8 string `json:"8"`
					Num9 string `json:"9"`
				} `json:"ArcherInventory"`
				HorsetamerInventory struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
					Num2 string `json:"2"`
					Num4 string `json:"4"`
					Num8 string `json:"8"`
				} `json:"HorsetamerInventory"`
				ChosenFinisher     string `json:"chosen_finisher"`
				HypeTrain          int    `json:"hype train"`
				HypeTrainInventory struct {
					Num0  string `json:"0"`
					Num2  string `json:"2"`
					Num3  string `json:"3"`
					Num4  string `json:"4"`
					Num5  string `json:"5"`
					Num6  string `json:"6"`
					Num7  string `json:"7"`
					Num8  string `json:"8"`
					Num9  string `json:"9"`
					Num13 string `json:"13"`
					Num14 string `json:"14"`
					Num15 string `json:"15"`
					Num16 string `json:"16"`
				} `json:"Hype TrainInventory"`
				Afterkill              string `json:"afterkill"`
				KitItemRenameHunterBow string `json:"kit_item_rename_Hunter_bow"`
				LastTourneyAd          int64  `json:"lastTourneyAd"`
				WinsBackup             int    `json:"wins_backup"`
				WinsTeamsNormal        int    `json:"wins_teams_normal"`
				WinsSoloNormal         int    `json:"wins_solo_normal"`
				Autoarmor              bool   `json:"autoarmor"`
				NecromancerInventory   struct {
					Num0  string `json:"0"`
					Num1  string `json:"1"`
					Num2  string `json:"2"`
					Num3  string `json:"3"`
					Num7  string `json:"7"`
					Num8  string `json:"8"`
					Num9  string `json:"9"`
					Num10 string `json:"10"`
					Num11 string `json:"11"`
					Num12 string `json:"12"`
				} `json:"NecromancerInventory"`
				ArrowsFiredNecromancer  int `json:"arrows_fired_necromancer"`
				ChestsOpenedNecromancer int `json:"chests_opened_necromancer"`
				Damage                  int `json:"damage"`
				ExpNecromancer          int `json:"exp_necromancer"`
				MobsSpawnedNecromancer  int `json:"mobs_spawned_necromancer"`
				ArrowsHit               int `json:"arrows_hit"`
				DamageNecromancer       int `json:"damage_necromancer"`
				GamesPlayedNecromancer  int `json:"games_played_necromancer"`
				TimePlayedNecromancer   int `json:"time_played_necromancer"`
				ArrowsHitNecromancer    int `json:"arrows_hit_necromancer"`
				KillsSoloNormal         int `json:"kills_solo_normal"`
				PotionsDrunk            int `json:"potions_drunk"`
				ArrowsFired             int `json:"arrows_fired"`
				BlitzUses               int `json:"blitz_uses"`
				MobsSpawned             int `json:"mobs_spawned"`
				TimePlayed              int `json:"time_played"`
				ChestsOpened            int `json:"chests_opened"`
				DamageTakenNecromancer  int `json:"damage_taken_necromancer"`
				DamageTaken             int `json:"damage_taken"`
				PotionsDrunkNecromancer int `json:"potions_drunk_necromancer"`
				KillsNecromancer        int `json:"kills_necromancer"`
				GamesPlayed             int `json:"games_played"`
				DamageTakenHunter       int `json:"damage_taken_hunter"`
				PotionsThrown           int `json:"potions_thrown"`
				KillsHunter             int `json:"kills_hunter"`
				DamageHunter            int `json:"damage_hunter"`
				GamesPlayedHunter       int `json:"games_played_hunter"`
				TimePlayedHunter        int `json:"time_played_hunter"`
				ChestsOpenedHunter      int `json:"chests_opened_hunter"`
				ExpHunter               int `json:"exp_hunter"`
				PotionsThrownHunter     int `json:"potions_thrown_hunter"`
				PotionsDrunkHunter      int `json:"potions_drunk_hunter"`
				ArrowsFiredHunter       int `json:"arrows_fired_hunter"`
				MobsSpawnedHunter       int `json:"mobs_spawned_hunter"`
				KillsTeamsNormal        int `json:"kills_teams_normal"`
			} `json:"HungerGames"`
			Mcgo struct {
				BombsDefused                int      `json:"bombs_defused"`
				BombsPlanted                int      `json:"bombs_planted"`
				CarbineCostReduction        int      `json:"carbine_cost_reduction"`
				CarbineDamageIncrease       int      `json:"carbine_damage_increase"`
				CarbineRecoilReduction      int      `json:"carbine_recoil_reduction"`
				CarbineReloadSpeedReduction int      `json:"carbine_reload_speed_reduction"`
				Coins                       int      `json:"coins"`
				CopKills                    int      `json:"cop_kills"`
				CriminalKills               int      `json:"criminal_kills"`
				GameWins                    int      `json:"game_wins"`
				HeadshotKills               int      `json:"headshot_kills"`
				Kills                       int      `json:"kills"`
				Kills102014                 int      `json:"kills_10_2014"`
				Kills3102014                int      `json:"kills_3_10_2014"`
				KnifeDamageIncrease         int      `json:"knife_damage_increase"`
				MagnumCostReduction         int      `json:"magnum_cost_reduction"`
				MagnumReloadSpeedReduction  int      `json:"magnum_reload_speed_reduction"`
				Packages                    []string `json:"packages"`
				PistolDamageIncrease        int      `json:"pistol_damage_increase"`
				RifleCostReduction          int      `json:"rifle_cost_reduction"`
				RifleDamageIncrease         int      `json:"rifle_damage_increase"`
				RifleRecoilReduction        int      `json:"rifle_recoil_reduction"`
				RifleReloadSpeedReduction   int      `json:"rifle_reload_speed_reduction"`
				RoundWins                   int      `json:"round_wins"`
				SelectedCarbineDev          string   `json:"selectedCarbineDev"`
				SelectedMagnumDev           string   `json:"selectedMagnumDev"`
				SelectedOcelotChestplateDev string   `json:"selectedOcelotChestplateDev"`
				SelectedOcelotHelmetDev     string   `json:"selectedOcelotHelmetDev"`
				SelectedSmgDev              string   `json:"selectedSmgDev"`
				ShotgunDamageIncrease       int      `json:"shotgun_damage_increase"`
				ShotgunRecoilReduction      int      `json:"shotgun_recoil_reduction"`
				ShotsFired                  int      `json:"shots_fired"`
				Deaths                      int      `json:"deaths"`
				BountyHunter                int      `json:"bounty_hunter"`
				SelectedRifleDev            string   `json:"selectedRifleDev"`
				MagnumRecoilReduction       int      `json:"magnum_recoil_reduction"`
				PistolRecoilReduction       int      `json:"pistol_recoil_reduction"`
				Mcgo                        struct {
					Points int `json:"points"`
				} `json:"mcgo"`
				WeeklyKillsB                 int    `json:"weekly_kills_b"`
				MonthlyKillsA                int    `json:"monthly_kills_a"`
				KnifeAttackDelay             int    `json:"knife_attack_delay"`
				SelectedKnifeDev             string `json:"selectedKnifeDev"`
				StrengthTraining             int    `json:"strength_training"`
				PistolReloadSpeedReduction   int    `json:"pistol_reload_speed_reduction"`
				BodyArmorCost                int    `json:"body_armor_cost"`
				PocketChange                 int    `json:"pocket_change"`
				GameWinsAlleyway             int    `json:"game_wins_alleyway"`
				MagnumDamageIncrease         int    `json:"magnum_damage_increase"`
				GameWinsDeathmatch           int    `json:"game_wins_deathmatch"`
				KillsDeathmatch              int    `json:"kills_deathmatch"`
				DeathsDeathmatch             int    `json:"deaths_deathmatch"`
				SmgCostReduction             int    `json:"smg_cost_reduction"`
				SmgReloadSpeedReduction      int    `json:"smg_reload_speed_reduction"`
				ShotgunCostReduction         int    `json:"shotgun_cost_reduction"`
				ShotgunReloadSpeedReduction  int    `json:"shotgun_reload_speed_reduction"`
				SelectedCreeperHelmetDev     string `json:"selectedCreeperHelmetDev"`
				SelectedCreeperChestplateDev string `json:"selectedCreeperChestplateDev"`
				GameWinsTemple               int    `json:"game_wins_temple"`
				GameWinsReserve              int    `json:"game_wins_reserve"`
				GameWinsCarrier              int    `json:"game_wins_carrier"`
				CriminalKillsDeathmatch      int    `json:"criminal_kills_deathmatch"`
				CopKillsDeathmatch           int    `json:"cop_kills_deathmatch"`
				GrenadeKills                 int    `json:"grenade_kills"`
				LastTourneyAd                int64  `json:"lastTourneyAd"`
			} `json:"MCGO"`
			Paintball struct {
				Coins          int      `json:"coins"`
				Deaths         int      `json:"deaths"`
				Fortune        int      `json:"fortune"`
				Godfather      int      `json:"godfather"`
				Hat            string   `json:"hat"`
				Kills          int      `json:"kills"`
				Packages       []string `json:"packages"`
				ShotsFired     int      `json:"shots_fired"`
				Superluck      int      `json:"superluck"`
				Wins           int      `json:"wins"`
				Killstreaks    int      `json:"killstreaks"`
				Adrenaline     int      `json:"adrenaline"`
				Endurance      int      `json:"endurance"`
				Transfusion    int      `json:"transfusion"`
				WeeklyKillsB   int      `json:"weekly_kills_b"`
				MonthlyKillsB  int      `json:"monthly_kills_b"`
				WeeklyKillsA   int      `json:"weekly_kills_a"`
				MonthlyKillsA  int      `json:"monthly_kills_a"`
				ShowKillPrefix bool     `json:"showKillPrefix"`
			} `json:"Paintball"`
			Quake struct {
				Barrel                             string   `json:"barrel"`
				Coins                              int      `json:"coins"`
				Deaths                             int      `json:"deaths"`
				Hat                                string   `json:"hat"`
				Kills                              int      `json:"kills"`
				Killstreaks                        int      `json:"killstreaks"`
				Packages                           []string `json:"packages"`
				Trigger                            string   `json:"trigger"`
				Wins                               int      `json:"wins"`
				Case                               string   `json:"case"`
				KillsTeams                         int      `json:"kills_teams"`
				KillstreaksTeams                   int      `json:"killstreaks_teams"`
				DeathsTeams                        int      `json:"deaths_teams"`
				WinsTeams                          int      `json:"wins_teams"`
				Muzzle                             string   `json:"muzzle"`
				Sight                              string   `json:"sight"`
				InstantRespawn                     bool     `json:"instantRespawn"`
				Armor                              string   `json:"armor"`
				Null                               string   `json:"null"`
				Killsound                          string   `json:"killsound"`
				MonthlyKillsA                      int      `json:"monthly_kills_a"`
				WeeklyKillsA                       int      `json:"weekly_kills_a"`
				WeeklyKillsB                       int      `json:"weekly_kills_b"`
				MonthlyKillsB                      int      `json:"monthly_kills_b"`
				DashPower                          string   `json:"dash_power"`
				DashCooldown                       string   `json:"dash_cooldown"`
				Leggings                           string   `json:"leggings"`
				Boots                              string   `json:"boots"`
				AlternativeGunCooldownIndicator    bool     `json:"alternative_gun_cooldown_indicator"`
				CompassSelected                    bool     `json:"compass_selected"`
				EnableSound                        bool     `json:"enable_sound"`
				ShowDashCooldown                   bool     `json:"showDashCooldown"`
				HighestKillstreak                  int      `json:"highest_killstreak"`
				DistanceTravelledTeams             int      `json:"distance_travelled_teams"`
				KillsSinceUpdateFeb2017Teams       int      `json:"kills_since_update_feb_2017_teams"`
				ShotsFiredTeams                    int      `json:"shots_fired_teams"`
				HeadshotsTeams                     int      `json:"headshots_teams"`
				ShowKillPrefix                     bool     `json:"showKillPrefix"`
				KillsDm                            int      `json:"kills_dm"`
				KillsDmTeams                       int      `json:"kills_dm_teams"`
				KillsTimeattack                    int      `json:"kills_timeattack"`
				MessageOthersKillsDeaths           bool     `json:"messageOthers' Kills/deaths"`
				MessageYourKills                   bool     `json:"messageYour Kills"`
				MessageCoinMessages                bool     `json:"messageCoin Messages"`
				MessageMultiKills                  bool     `json:"messageMulti-kills"`
				MessageKillstreaks                 bool     `json:"messageKillstreaks"`
				MessagePowerupCollections          bool     `json:"messagePowerup Collections"`
				MessageYourDeaths                  bool     `json:"messageYour Deaths"`
				Headshots                          int      `json:"headshots"`
				DistanceTravelled                  int      `json:"distance_travelled"`
				ShotsFired                         int      `json:"shots_fired"`
				KillsSinceUpdateFeb2017            int      `json:"kills_since_update_feb_2017"`
				KillstreaksSoloTourney             int      `json:"killstreaks_solo_tourney"`
				HeadshotsSoloTourney               int      `json:"headshots_solo_tourney"`
				WinsSoloTourney                    int      `json:"wins_solo_tourney"`
				KillsSoloTourney                   int      `json:"kills_solo_tourney"`
				ShotsFiredSoloTourney              int      `json:"shots_fired_solo_tourney"`
				KillsSinceUpdateFeb2017SoloTourney int      `json:"kills_since_update_feb_2017_solo_tourney"`
				DistanceTravelledSoloTourney       int      `json:"distance_travelled_solo_tourney"`
				DeathsSoloTourney                  int      `json:"deaths_solo_tourney"`
				MessageCoin                        bool     `json:"messageCoin"`
			} `json:"Quake"`
			TNTGames struct {
				BloodwizardExplode            int      `json:"bloodwizard_explode"`
				BloodwizardRegen              int      `json:"bloodwizard_regen"`
				Coins                         int      `json:"coins"`
				DeathsBowspleef               int      `json:"deaths_bowspleef"`
				DeathsCapture                 int      `json:"deaths_capture"`
				DoublejumpTntrun              int      `json:"doublejump_tntrun"`
				FirewizardExplode             int      `json:"firewizard_explode"`
				FirewizardRegen               int      `json:"firewizard_regen"`
				KillsCapture                  int      `json:"kills_capture"`
				KineticwizardExplode          int      `json:"kineticwizard_explode"`
				KineticwizardRegen            int      `json:"kineticwizard_regen"`
				Packages                      []string `json:"packages"`
				SpleefDoublejump              int      `json:"spleef_doublejump"`
				SpleefRepulse                 int      `json:"spleef_repulse"`
				SpleefTriple                  int      `json:"spleef_triple"`
				TagsBowspleef                 int      `json:"tags_bowspleef"`
				WinsBowspleef                 int      `json:"wins_bowspleef"`
				WinsCapture                   int      `json:"wins_capture"`
				WinsTntag                     int      `json:"wins_tntag"`
				WinsTntrun                    int      `json:"wins_tntrun"`
				WitherwizardExplode           int      `json:"witherwizard_explode"`
				WitherwizardRegen             int      `json:"witherwizard_regen"`
				SelectedHat                   string   `json:"selected_hat"`
				RecordPvprun                  int      `json:"record_pvprun"`
				KillsPvprun                   int      `json:"kills_pvprun"`
				RecordTntrun                  int      `json:"record_tntrun"`
				TagSpeed                      int      `json:"tag_speed"`
				CaptureClass                  string   `json:"capture_class"`
				AssistsCapture                int      `json:"assists_capture"`
				WinsPvprun                    int      `json:"wins_pvprun"`
				VotesTNTRunA                  int      `json:"votes_TNT Run (A)"`
				ActiveParticleEffect          string   `json:"active_particle_effect"`
				IcewizardRegen                int      `json:"icewizard_regen"`
				IcewizardExplode              int      `json:"icewizard_explode"`
				ActiveDeathEffect             string   `json:"active_death_effect"`
				KillsTntag                    int      `json:"kills_tntag"`
				NewWitherwizardRegen          int      `json:"new_witherwizard_regen"`
				NewSpleefTripleshot           int      `json:"new_spleef_tripleshot"`
				NewIcewizardExplode           int      `json:"new_icewizard_explode"`
				NewSpleefRepulsor             int      `json:"new_spleef_repulsor"`
				NewKineticwizardRegen         int      `json:"new_kineticwizard_regen"`
				NewSpleefDoubleJumps          int      `json:"new_spleef_double_jumps"`
				NewFirewizardExplode          int      `json:"new_firewizard_explode"`
				NewTntagSpeedy                int      `json:"new_tntag_speedy"`
				NewIcewizardRegen             int      `json:"new_icewizard_regen"`
				NewWitherwizardExplode        int      `json:"new_witherwizard_explode"`
				NewBloodwizardRegen           int      `json:"new_bloodwizard_regen"`
				NewPvprunDoubleJumps          int      `json:"new_pvprun_double_jumps"`
				NewBloodwizardExplode         int      `json:"new_bloodwizard_explode"`
				NewTntrunDoubleJumps          int      `json:"new_tntrun_double_jumps"`
				NewFirewizardRegen            int      `json:"new_firewizard_regen"`
				NewKineticwizardExplode       int      `json:"new_kineticwizard_explode"`
				NewToxicwizardRegen           int      `json:"new_toxicwizard_regen"`
				NewToxicwizardExplode         int      `json:"new_toxicwizard_explode"`
				NewTntrunSlownessPotions      int      `json:"new_tntrun_slowness_potions"`
				NewTntrunSpeedPotions         int      `json:"new_tntrun_speed_potions"`
				TagSpeeditup                  int      `json:"tag_speeditup"`
				TagBlastprotection            int      `json:"tag_blastprotection"`
				TagSlowitdown                 int      `json:"tag_slowitdown"`
				NewPvprunFortitude            int      `json:"new_pvprun_fortitude"`
				NewPvprunRegeneration         int      `json:"new_pvprun_regeneration"`
				NewPvprunNotoriety            int      `json:"new_pvprun_notoriety"`
				RunPotionsSplashedOnPlayers   int      `json:"run_potions_splashed_on_players"`
				Winstreak                     int      `json:"winstreak"`
				WizardsSelectedClass          string   `json:"wizards_selected_class"`
				NewWitherwizardAssists        int      `json:"new_witherwizard_assists"`
				NewWitherwizardKills          int      `json:"new_witherwizard_kills"`
				NewWitherwizardDeaths         int      `json:"new_witherwizard_deaths"`
				DeathsPvprun                  int      `json:"deaths_pvprun"`
				DeathsTntrun                  int      `json:"deaths_tntrun"`
				NewKineticwizardKills         int      `json:"new_kineticwizard_kills"`
				NewKineticwizardAssists       int      `json:"new_kineticwizard_assists"`
				NewToxicwizardDeaths          int      `json:"new_toxicwizard_deaths"`
				NewToxicwizardKills           int      `json:"new_toxicwizard_kills"`
				NewKineticwizardDeaths        int      `json:"new_kineticwizard_deaths"`
				NewToxicwizardAssists         int      `json:"new_toxicwizard_assists"`
				NewActiveParticleEffect       string   `json:"new_active_particle_effect"`
				NewFirewizardAssists          int      `json:"new_firewizard_assists"`
				NewFirewizardKills            int      `json:"new_firewizard_kills"`
				NewFirewizardDeaths           int      `json:"new_firewizard_deaths"`
				Wins                          int      `json:"wins"`
				NewBloodwizardAssists         int      `json:"new_bloodwizard_assists"`
				NewBloodwizardKills           int      `json:"new_bloodwizard_kills"`
				NewBloodwizardDeaths          int      `json:"new_bloodwizard_deaths"`
				NewToxicwizardRegenLegacy     int      `json:"new_toxicwizard_regen_legacy"`
				NewIcewizardExplodeLegacy     int      `json:"new_icewizard_explode_legacy"`
				NewFirewizardRegenLegacy      int      `json:"new_firewizard_regen_legacy"`
				NewSpleefTripleshotLegacy     int      `json:"new_spleef_tripleshot_legacy"`
				NewSpleefDoubleJumpsLegacy    int      `json:"new_spleef_double_jumps_legacy"`
				NewKineticwizardExplodeLegacy int      `json:"new_kineticwizard_explode_legacy"`
				NewIcewizardRegenLegacy       int      `json:"new_icewizard_regen_legacy"`
				NewBloodwizardExplodeLegacy   int      `json:"new_bloodwizard_explode_legacy"`
				NewWitherwizardRegenLegacy    int      `json:"new_witherwizard_regen_legacy"`
				NewTntrunDoubleJumpsLegacy    int      `json:"new_tntrun_double_jumps_legacy"`
				NewKineticwizardRegenLegacy   int      `json:"new_kineticwizard_regen_legacy"`
				NewSpleefRepulsorLegacy       int      `json:"new_spleef_repulsor_legacy"`
				NewBloodwizardRegenLegacy     int      `json:"new_bloodwizard_regen_legacy"`
				NewWitherwizardExplodeLegacy  int      `json:"new_witherwizard_explode_legacy"`
				NewToxicwizardExplodeLegacy   int      `json:"new_toxicwizard_explode_legacy"`
				NewFirewizardExplodeLegacy    int      `json:"new_firewizard_explode_legacy"`
				NewPvprunDoubleJumpsLegacy    int      `json:"new_pvprun_double_jumps_legacy"`
				PointsCapture                 int      `json:"points_capture"`
				NewKineticwizardDamageTaken   int      `json:"new_kineticwizard_damage_taken"`
				NewFirewizardDamageTaken      int      `json:"new_firewizard_damage_taken"`
				NewKineticwizardHealing       int      `json:"new_kineticwizard_healing"`
				KineticHealingCapture         int      `json:"kinetic_healing_capture"`
				NewFirewizardHealing          int      `json:"new_firewizard_healing"`
				AirTimeCapture                int      `json:"air_time_capture"`
				Privategames                  struct {
					LowGravity       bool   `json:"low_gravity"`
					Speed            string `json:"speed"`
					TntRunSnowballs  bool   `json:"tnt_run_snowballs"`
					TntTagDeathmatch bool   `json:"tnt_tag_deathmatch"`
					PvpRunSwordType  string `json:"pvp_run_sword_type"`
					PvpRunArmorType  string `json:"pvp_run_armor_type"`
					MaxedPerks       bool   `json:"maxed_perks"`
					TntTagNoPowerups bool   `json:"tnt_tag_no_powerups"`
					PvpRunKnockback  bool   `json:"pvp_run_knockback"`
				} `json:"privategames"`
				Flags struct {
					EnableExplosiveDash              bool `json:"enable_explosive_dash"`
					GiveDjFeather                    bool `json:"give_dj_feather"`
					ShowTipHolograms                 bool `json:"show_tip_holograms"`
					ShowTntrunActionbarInfo          bool `json:"show_tntrun_actionbar_info"`
					ShowTnttagActionbarInfo          bool `json:"show_tnttag_actionbar_info"`
					ShowWizardsActionbarInfo         bool `json:"show_wizards_actionbar_info"`
					ShowWizardsCooldownNotifications bool `json:"show_wizards_cooldown_notifications"`
				} `json:"flags"`
				LastTourneyAd         int64 `json:"lastTourneyAd"`
				NewStormwizardExplode int   `json:"new_stormwizard_explode"`
				NewStormwizardRegen   int   `json:"new_stormwizard_regen"`
			} `json:"TNTGames"`
			Uhc struct {
				Coins                   int      `json:"coins"`
				Deaths                  int      `json:"deaths"`
				EquippedKit             string   `json:"equippedKit"`
				HeadsEaten              int      `json:"heads_eaten"`
				Kills                   int      `json:"kills"`
				Packages                []string `json:"packages"`
				PerkAlchemyLineA        int      `json:"perk_alchemy_line_a"`
				PerkAlchemyLineB        int      `json:"perk_alchemy_line_b"`
				PerkArmorsmithLineA     int      `json:"perk_armorsmith_line_a"`
				PerkArmorsmithLineB     int      `json:"perk_armorsmith_line_b"`
				PerkArmorsmithLineC     int      `json:"perk_armorsmith_line_c"`
				PerkEnchantingLineA     int      `json:"perk_enchanting_line_a"`
				PerkEnchantingLineB     int      `json:"perk_enchanting_line_b"`
				PerkEnchantingLineC     int      `json:"perk_enchanting_line_c"`
				PerkEngineeringLineA    int      `json:"perk_engineering_line_a"`
				PerkEngineeringLineB    int      `json:"perk_engineering_line_b"`
				PerkEngineeringLineC    int      `json:"perk_engineering_line_c"`
				PerkWeaponsmithLineA    int      `json:"perk_weaponsmith_line_a"`
				PerkWeaponsmithLineB    int      `json:"perk_weaponsmith_line_b"`
				PerkWeaponsmithLineC    int      `json:"perk_weaponsmith_line_c"`
				Score                   int      `json:"score"`
				Wins                    int      `json:"wins"`
				PerkBloodcraftLineA     int      `json:"perk_bloodcraft_line_a"`
				PerkBloodcraftLineB     int      `json:"perk_bloodcraft_line_b"`
				PerkHunterLineA         int      `json:"perk_hunter_line_a"`
				PerkHunterLineB         int      `json:"perk_hunter_line_b"`
				PerkHunterLineC         int      `json:"perk_hunter_line_c"`
				KitARCHERYTOOLS         int      `json:"kit_ARCHERY_TOOLS"`
				PerkCookingLineA        int      `json:"perk_cooking_line_a"`
				PerkCookingLineC        int      `json:"perk_cooking_line_c"`
				PerkCookingLineB        int      `json:"perk_cooking_line_b"`
				MonthlyKillsA           int      `json:"monthly_kills_a"`
				MonthlyWinsA            int      `json:"monthly_wins_a"`
				PerkCookingPrestige     int      `json:"perk_cooking_prestige"`
				PerkBloodcraftLineC     int      `json:"perk_bloodcraft_line_c"`
				PerkBloodcraftPrestige  int      `json:"perk_bloodcraft_prestige"`
				DeathsSolo              int      `json:"deaths_solo"`
				Cache3                  bool     `json:"cache3"`
				MonthlyKillsB           int      `json:"monthly_kills__b"`
				ClearupAchievement      bool     `json:"clearup_achievement"`
				KillsSolo               int      `json:"kills_solo"`
				HeadsEatenSolo          int      `json:"heads_eaten_solo"`
				PerkSurvivalismLineA    int      `json:"perk_survivalism_line_a"`
				PerkSurvivalismLineC    int      `json:"perk_survivalism_line_c"`
				PerkSurvivalismLineB    int      `json:"perk_survivalism_line_b"`
				KillsSolo2              int      `json:"kills_solo2"`
				SavedStats              bool     `json:"saved_stats"`
				Kills2                  int      `json:"kills2"`
				Wins2                   int      `json:"wins2"`
				WinsSolo                int      `json:"wins_solo"`
				PerkWeaponsmithPrestige int      `json:"perk_weaponsmith_prestige"`
			} `json:"UHC"`
			VampireZ struct {
				BabyHater            int      `json:"baby_hater"`
				Blood                bool     `json:"blood"`
				BloodBooster         int      `json:"blood_booster"`
				Coins                int      `json:"coins"`
				DrainPunch           int      `json:"drain_punch"`
				ExpertSwag           int      `json:"expert_swag"`
				ExplosiveKiller      int      `json:"explosive_killer"`
				FinalBreath          int      `json:"final_breath"`
				Fireproofing         int      `json:"fireproofing"`
				GoldBooster          int      `json:"gold_booster"`
				GoldStarter          int      `json:"gold_starter"`
				HumanDeaths          int      `json:"human_deaths"`
				HumanKills           int      `json:"human_kills"`
				HumanWins            int      `json:"human_wins"`
				LootDrops            int      `json:"loot_drops"`
				MostVampireKills     int      `json:"most_vampire_kills"`
				Packages             []string `json:"packages"`
				VampireDeaths        int      `json:"vampire_deaths"`
				VampireDoubler       int      `json:"vampire_doubler"`
				VampireKills         int      `json:"vampire_kills"`
				VampiricScream       int      `json:"vampiric_scream"`
				WaveBooster          int      `json:"wave_booster"`
				ZombieDoubler        int      `json:"zombie_doubler"`
				ZombieKills          int      `json:"zombie_kills"`
				Hellborn             int      `json:"hellborn"`
				Renfield             int      `json:"renfield"`
				BloodDrinker         int      `json:"blood_drinker"`
				TerrorLevel          int      `json:"terror_level"`
				KillBooster          int      `json:"kill_booster"`
				Transfusion          int      `json:"transfusion"`
				BasicSwag            int      `json:"basic_swag"`
				AdvancedSwag         int      `json:"advanced_swag"`
				Constitution         int      `json:"constitution"`
				Foresight            int      `json:"foresight"`
				VotesDarkValley      int      `json:"votes_Dark Valley"`
				GoldBought           int      `json:"gold_bought"`
				MonthlyHumanWinsB    int      `json:"monthly_human_wins_b"`
				WeeklyHumanWinsA     int      `json:"weekly_human_wins_a"`
				MonthlyVampireWinsB  int      `json:"monthly_vampire_wins_b"`
				WeeklyVampireWinsA   int      `json:"weekly_vampire_wins_a"`
				WeeklyHumanWinsB     int      `json:"weekly_human_wins_b"`
				MonthlyHumanWinsA    int      `json:"monthly_human_wins_a"`
				WeeklyVampireWinsB   int      `json:"weekly_vampire_wins_b"`
				MonthlyVampireWinsA  int      `json:"monthly_vampire_wins_a"`
				VampiricMinion       int      `json:"vampiric_minion"`
				FrankensteinsMonster int      `json:"frankensteins_monster"`
				UpdatedStats         bool     `json:"updated_stats"`
				VampireWins          int      `json:"vampire_wins"`
				MostVampireKillsNew  int      `json:"most_vampire_kills_new"`
			} `json:"VampireZ"`
			Walls struct {
				Berserk       int      `json:"berserk"`
				Blood         bool     `json:"blood"`
				BossDigger    int      `json:"boss_digger"`
				BossSkills    int      `json:"boss_skills"`
				Coins         int      `json:"coins"`
				Deaths        int      `json:"deaths"`
				Fortune       int      `json:"fortune"`
				InsaneFarmer  int      `json:"insane_farmer"`
				Kills         int      `json:"kills"`
				Losses        int      `json:"losses"`
				Packages      []string `json:"packages"`
				Scotsman      int      `json:"scotsman"`
				Swift         int      `json:"swift"`
				Wins          int      `json:"wins"`
				Chef          int      `json:"chef"`
				Hunter        int      `json:"hunter"`
				Sage          int      `json:"sage"`
				SmartBoy      int      `json:"smart_boy"`
				Bacon         int      `json:"bacon"`
				Guitarist     int      `json:"guitarist"`
				SnackLover    int      `json:"snack_lover"`
				Blacksmith    int      `json:"blacksmith"`
				FinalForm     int      `json:"final_form"`
				SkybaseKing   int      `json:"skybase_king"`
				Einstein      int      `json:"einstein"`
				Canadian      int      `json:"canadian"`
				Tenacity      int      `json:"tenacity"`
				Vampirism     int      `json:"vampirism"`
				ExpertMiner   int      `json:"expert_miner"`
				Adrenaline    int      `json:"adrenaline"`
				Lazyman       int      `json:"lazyman"`
				Step          int      `json:"step"`
				Bomberman     int      `json:"bomberman"`
				Ready         int      `json:"ready"`
				Artisan       int      `json:"artisan"`
				Fisherman     int      `json:"fisherman"`
				GoldRush      int      `json:"gold_rush"`
				LeatherWorker int      `json:"leather_worker"`
				SoupDrinker   int      `json:"soup_drinker"`
				BossGuardian  int      `json:"boss_guardian"`
				Opportunity   int      `json:"opportunity"`
				Haste         int      `json:"haste"`
				MonthlyKillsB int      `json:"monthly_kills_b"`
				WeeklyKillsA  int      `json:"weekly_kills_a"`
				Inventory     struct {
					Num0  string `json:"0"`
					Num1  string `json:"1"`
					Num2  string `json:"2"`
					Num3  string `json:"3"`
					Num4  string `json:"4"`
					Num5  string `json:"5"`
					Num6  string `json:"6"`
					Num8  string `json:"8"`
					Num9  string `json:"9"`
					Num10 string `json:"10"`
					Num11 string `json:"11"`
					Num12 string `json:"12"`
					Num13 string `json:"13"`
					Num14 string `json:"14"`
					Num15 string `json:"15"`
					Num16 string `json:"16"`
					Num18 string `json:"18"`
					Num19 string `json:"19"`
					Num20 string `json:"20"`
					Num21 string `json:"21"`
					Num22 string `json:"22"`
					Num23 string `json:"23"`
				} `json:"Inventory"`
				BurnBabyBurn     int `json:"burn_baby_burn"`
				Vitality         int `json:"vitality"`
				WeeklyWinsA      int `json:"weekly_wins_a"`
				MonthlyWinsB     int `json:"monthly_wins_b"`
				CreeperEgg       int `json:"creeper_egg"`
				Chainkiller      int `json:"chainkiller"`
				MasterTroll      int `json:"master_troll"`
				Fireproof        int `json:"fireproof"`
				InventoryLayout2 struct {
					Num0  string `json:"0"`
					Num1  string `json:"1"`
					Num2  string `json:"2"`
					Num3  string `json:"3"`
					Num4  string `json:"4"`
					Num5  string `json:"5"`
					Num6  string `json:"6"`
					Num7  string `json:"7"`
					Num8  string `json:"8"`
					Num9  string `json:"9"`
					Num10 string `json:"10"`
					Num11 string `json:"11"`
					Num12 string `json:"12"`
					Num13 string `json:"13"`
					Num14 string `json:"14"`
					Num15 string `json:"15"`
					Num16 string `json:"16"`
					Num17 string `json:"17"`
					Num18 string `json:"18"`
					Num19 string `json:"19"`
					Num20 string `json:"20"`
					Num21 string `json:"21"`
					Num22 string `json:"22"`
					Num23 string `json:"23"`
					Num24 string `json:"24"`
					Num25 string `json:"25"`
					Num26 string `json:"26"`
					Num27 string `json:"27"`
				} `json:"InventoryLayout2"`
				BlacksmithStarter int `json:"blacksmith_starter"`
			} `json:"Walls"`
			Walls3 struct {
				Assists                                  int      `json:"assists"`
				AssistsSkeleton                          int      `json:"assists_Skeleton"`
				BlazeA                                   int      `json:"blaze_a"`
				BlazeB                                   int      `json:"blaze_b"`
				BlazeC                                   int      `json:"blaze_c"`
				BlazeD                                   int      `json:"blaze_d"`
				BlazeG                                   int      `json:"blaze_g"`
				Blood                                    bool     `json:"blood"`
				ChosenClass                              string   `json:"chosen_class"`
				Coins                                    int      `json:"coins"`
				CreeperA                                 int      `json:"creeper_a"`
				CreeperB                                 int      `json:"creeper_b"`
				CreeperC                                 int      `json:"creeper_c"`
				CreeperD                                 int      `json:"creeper_d"`
				CreeperG                                 int      `json:"creeper_g"`
				Deaths                                   int      `json:"deaths"`
				DeathsBlaze                              int      `json:"deaths_Blaze"`
				DeathsCreeper                            int      `json:"deaths_Creeper"`
				DeathsEnderman                           int      `json:"deaths_Enderman"`
				DeathsGolem                              int      `json:"deaths_Golem"`
				DeathsHerobrine                          int      `json:"deaths_Herobrine"`
				DeathsPigman                             int      `json:"deaths_Pigman"`
				DeathsShaman                             int      `json:"deaths_Shaman"`
				DeathsSkeleton                           int      `json:"deaths_Skeleton"`
				DeathsSpider                             int      `json:"deaths_Spider"`
				DeathsZombie                             int      `json:"deaths_Zombie"`
				EndermanA                                int      `json:"enderman_a"`
				EndermanB                                int      `json:"enderman_b"`
				EndermanC                                int      `json:"enderman_c"`
				EndermanD                                int      `json:"enderman_d"`
				FinalKillsBlaze                          int      `json:"finalKills_Blaze"`
				FinalKillsCreeper                        int      `json:"finalKills_Creeper"`
				FinalKillsEnderman                       int      `json:"finalKills_Enderman"`
				FinalKillsGolem                          int      `json:"finalKills_Golem"`
				FinalKillsHerobrine                      int      `json:"finalKills_Herobrine"`
				FinalKillsPigman                         int      `json:"finalKills_Pigman"`
				FinalKillsShaman                         int      `json:"finalKills_Shaman"`
				FinalKillsSkeleton                       int      `json:"finalKills_Skeleton"`
				FinalKillsSpider                         int      `json:"finalKills_Spider"`
				FinalKillsZombie                         int      `json:"finalKills_Zombie"`
				GolemD                                   int      `json:"golem_d"`
				HerobrineA                               int      `json:"herobrine_a"`
				HerobrineB                               int      `json:"herobrine_b"`
				HerobrineC                               int      `json:"herobrine_c"`
				HerobrineD                               int      `json:"herobrine_d"`
				HerobrineG                               int      `json:"herobrine_g"`
				Kills                                    int      `json:"kills"`
				KillsBlaze                               int      `json:"kills_Blaze"`
				KillsCreeper                             int      `json:"kills_Creeper"`
				KillsEnderman                            int      `json:"kills_Enderman"`
				KillsGolem                               int      `json:"kills_Golem"`
				KillsHerobrine                           int      `json:"kills_Herobrine"`
				KillsPigman                              int      `json:"kills_Pigman"`
				KillsShaman                              int      `json:"kills_Shaman"`
				KillsSkeleton                            int      `json:"kills_Skeleton"`
				KillsSpider                              int      `json:"kills_Spider"`
				KillsZombie                              int      `json:"kills_Zombie"`
				Losses                                   int      `json:"losses"`
				LossesBlaze                              int      `json:"losses_Blaze"`
				LossesCreeper                            int      `json:"losses_Creeper"`
				LossesEnderman                           int      `json:"losses_Enderman"`
				LossesGolem                              int      `json:"losses_Golem"`
				LossesHerobrine                          int      `json:"losses_Herobrine"`
				LossesPigman                             int      `json:"losses_Pigman"`
				LossesShaman                             int      `json:"losses_Shaman"`
				LossesSkeleton                           int      `json:"losses_Skeleton"`
				LossesSpider                             int      `json:"losses_Spider"`
				LossesZombie                             int      `json:"losses_Zombie"`
				MonthlyFinalKillsHerobrineA              int      `json:"monthly_finalKills_Herobrine_a"`
				MonthlyFinalKillsPigmanA                 int      `json:"monthly_finalKills_Pigman_a"`
				MonthlyFinalKillsPigmanB                 int      `json:"monthly_finalKills_Pigman_b"`
				MonthlyFinalKillsSkeletonA               int      `json:"monthly_finalKills_Skeleton_a"`
				MonthlyFinalKillsSkeletonB               int      `json:"monthly_finalKills_Skeleton_b"`
				MonthlyFinalKillsSpiderA                 int      `json:"monthly_finalKills_Spider_a"`
				MonthlyFinalKillsSpiderB                 int      `json:"monthly_finalKills_Spider_b"`
				MonthlyFinalKillsZombieA                 int      `json:"monthly_finalKills_Zombie_a"`
				MutationsVisibility                      bool     `json:"mutations_visibility"`
				Packages                                 []string `json:"packages"`
				PigmanA                                  int      `json:"pigman_a"`
				PigmanB                                  int      `json:"pigman_b"`
				PigmanC                                  int      `json:"pigman_c"`
				PigmanD                                  int      `json:"pigman_d"`
				ShamanA                                  int      `json:"shaman_a"`
				ShamanB                                  int      `json:"shaman_b"`
				ShamanC                                  int      `json:"shaman_c"`
				ShamanD                                  int      `json:"shaman_d"`
				ShamanG                                  int      `json:"shaman_g"`
				SkeletonA                                int      `json:"skeleton_a"`
				SkeletonB                                int      `json:"skeleton_b"`
				SkeletonC                                int      `json:"skeleton_c"`
				SkeletonD                                int      `json:"skeleton_d"`
				SkeletonG                                int      `json:"skeleton_g"`
				SpiderA                                  int      `json:"spider_a"`
				SpiderB                                  int      `json:"spider_b"`
				SpiderD                                  int      `json:"spider_d"`
				WarCry                                   string   `json:"war_cry"`
				WeeklyDeaths                             int      `json:"weeklyDeaths"`
				WeeklyDeathsBlaze                        int      `json:"weeklyDeaths_Blaze"`
				WeeklyDeathsGolem                        int      `json:"weeklyDeaths_Golem"`
				WeeklyDeathsHerobrine                    int      `json:"weeklyDeaths_Herobrine"`
				WeeklyDeathsPigman                       int      `json:"weeklyDeaths_Pigman"`
				WeeklyDeathsSkeleton                     int      `json:"weeklyDeaths_Skeleton"`
				WeeklyDeathsSpider                       int      `json:"weeklyDeaths_Spider"`
				WeeklyDeathsZombie                       int      `json:"weeklyDeaths_Zombie"`
				WeeklyFinalKillsBlaze                    int      `json:"weeklyFinalKills_Blaze"`
				WeeklyFinalKillsGolem                    int      `json:"weeklyFinalKills_Golem"`
				WeeklyFinalKillsPigman                   int      `json:"weeklyFinalKills_Pigman"`
				WeeklyKills                              int      `json:"weeklyKills"`
				WeeklyKillsBlaze                         int      `json:"weeklyKills_Blaze"`
				WeeklyKillsGolem                         int      `json:"weeklyKills_Golem"`
				WeeklyKillsHerobrine                     int      `json:"weeklyKills_Herobrine"`
				WeeklyKillsPigman                        int      `json:"weeklyKills_Pigman"`
				WeeklyKillsSkeleton                      int      `json:"weeklyKills_Skeleton"`
				WeeklyKillsSpider                        int      `json:"weeklyKills_Spider"`
				WeeklyKillsZombie                        int      `json:"weeklyKills_Zombie"`
				WeeklyLosses                             int      `json:"weeklyLosses"`
				WeeklyLossesBlaze                        int      `json:"weeklyLosses_Blaze"`
				WeeklyLossesGolem                        int      `json:"weeklyLosses_Golem"`
				WeeklyLossesHerobrine                    int      `json:"weeklyLosses_Herobrine"`
				WeeklyLossesPigman                       int      `json:"weeklyLosses_Pigman"`
				WeeklyLossesSkeleton                     int      `json:"weeklyLosses_Skeleton"`
				WeeklyLossesSpider                       int      `json:"weeklyLosses_Spider"`
				WeeklyLossesZombie                       int      `json:"weeklyLosses_Zombie"`
				WeeklyWins                               int      `json:"weeklyWins"`
				WeeklyWinsBlaze                          int      `json:"weeklyWins_Blaze"`
				WeeklyWinsGolem                          int      `json:"weeklyWins_Golem"`
				WeeklyWinsPigman                         int      `json:"weeklyWins_Pigman"`
				WeeklyWinsSkeleton                       int      `json:"weeklyWins_Skeleton"`
				WeeklyWinsSpider                         int      `json:"weeklyWins_Spider"`
				WeeklyWinsZombie                         int      `json:"weeklyWins_Zombie"`
				WeeklyFinalKillsHerobrineB               int      `json:"weekly_finalKills_Herobrine_b"`
				WeeklyFinalKillsPigmanA                  int      `json:"weekly_finalKills_Pigman_a"`
				WeeklyFinalKillsPigmanB                  int      `json:"weekly_finalKills_Pigman_b"`
				WeeklyFinalKillsSkeletonA                int      `json:"weekly_finalKills_Skeleton_a"`
				WeeklyFinalKillsSkeletonB                int      `json:"weekly_finalKills_Skeleton_b"`
				WeeklyFinalKillsSpiderA                  int      `json:"weekly_finalKills_Spider_a"`
				WeeklyFinalKillsSpiderB                  int      `json:"weekly_finalKills_Spider_b"`
				WeeklyFinalKillsZombieB                  int      `json:"weekly_finalKills_Zombie_b"`
				Wins                                     int      `json:"wins"`
				WinsBlaze                                int      `json:"wins_Blaze"`
				WinsCreeper                              int      `json:"wins_Creeper"`
				WinsEnderman                             int      `json:"wins_Enderman"`
				WinsGolem                                int      `json:"wins_Golem"`
				WinsHerobrine                            int      `json:"wins_Herobrine"`
				WinsPigman                               int      `json:"wins_Pigman"`
				WinsShaman                               int      `json:"wins_Shaman"`
				WinsSkeleton                             int      `json:"wins_Skeleton"`
				WinsSpider                               int      `json:"wins_Spider"`
				WinsZombie                               int      `json:"wins_Zombie"`
				ZombieA                                  int      `json:"zombie_a"`
				ZombieB                                  int      `json:"zombie_b"`
				ZombieC                                  int      `json:"zombie_c"`
				ZombieD                                  int      `json:"zombie_d"`
				ZombieG                                  int      `json:"zombie_g"`
				LossesArcanist                           int      `json:"losses_Arcanist"`
				WeeklyLossesArcanist                     int      `json:"weeklyLosses_Arcanist"`
				MonthlyFinalKillsGolemA                  int      `json:"monthly_finalKills_Golem_a"`
				WeeklyFinalKillsGolemA                   int      `json:"weekly_finalKills_Golem_a"`
				AssistsGolem                             int      `json:"assists_Golem"`
				FinalAssistsGolem                        int      `json:"finalAssists_Golem"`
				AssistsHerobrine                         int      `json:"assists_Herobrine"`
				WeeklyWinsHerobrine                      int      `json:"weeklyWins_Herobrine"`
				FinalAssistsHerobrine                    int      `json:"finalAssists_Herobrine"`
				FinalAssistsSpider                       int      `json:"finalAssists_Spider"`
				PlaysStandard                            int      `json:"plays_standard"`
				MonthlyFinalKillsHerobrineB              int      `json:"monthly_finalKills_Herobrine_b"`
				WeeklyFinalKillsHerobrineA               int      `json:"weekly_finalKills_Herobrine_a"`
				MonthlyFinalKillsGolemB                  int      `json:"monthly_finalKills_Golem_b"`
				FinalAssistsSkeleton                     int      `json:"finalAssists_Skeleton"`
				PlaysPractice                            int      `json:"plays_practice"`
				WeeklyWinsPracticeSkeleton               int      `json:"weeklyWins_practice_Skeleton"`
				WinsPractice                             int      `json:"wins_practice"`
				WeeklyWinsPractice                       int      `json:"weeklyWins_practice"`
				WinsPracticeSkeleton                     int      `json:"wins_practice_Skeleton"`
				FinalKillsPirate                         int      `json:"finalKills_Pirate"`
				FinalAssistsPirate                       int      `json:"finalAssists_Pirate"`
				WeeklyLossesPirate                       int      `json:"weeklyLosses_Pirate"`
				AssistsPirate                            int      `json:"assists_Pirate"`
				DeathsPirate                             int      `json:"deaths_Pirate"`
				WeeklyKillsPirate                        int      `json:"weeklyKills_Pirate"`
				WeeklyFinalKillsPirateA                  int      `json:"weekly_finalKills_Pirate_a"`
				LossesPirate                             int      `json:"losses_Pirate"`
				WeeklyDeathsPirate                       int      `json:"weeklyDeaths_Pirate"`
				KillsPirate                              int      `json:"kills_Pirate"`
				MonthlyFinalKillsPirateB                 int      `json:"monthly_finalKills_Pirate_b"`
				FinalKillsHunter                         int      `json:"finalKills_Hunter"`
				WinsHunter                               int      `json:"wins_Hunter"`
				WeeklyWinsHunter                         int      `json:"weeklyWins_Hunter"`
				WeeklyKillsHunter                        int      `json:"weeklyKills_Hunter"`
				AssistsHunter                            int      `json:"assists_Hunter"`
				KillsHunter                              int      `json:"kills_Hunter"`
				FinalAssistsHunter                       int      `json:"finalAssists_Hunter"`
				WeeklyFinalKillsHunterA                  int      `json:"weekly_finalKills_Hunter_a"`
				MonthlyFinalKillsHunterB                 int      `json:"monthly_finalKills_Hunter_b"`
				LossesHunter                             int      `json:"losses_Hunter"`
				WeeklyLossesHunter                       int      `json:"weeklyLosses_Hunter"`
				WeeklyLossesFaceOffHunter                int      `json:"weeklyLosses_face_off_Hunter"`
				PlaysFaceOff                             int      `json:"plays_face_off"`
				LossesFaceOff                            int      `json:"losses_face_off"`
				LossesFaceOffHunter                      int      `json:"losses_face_off_Hunter"`
				WeeklyLossesFaceOff                      int      `json:"weeklyLosses_face_off"`
				WinsFaceOffHunter                        int      `json:"wins_face_off_Hunter"`
				WinsFaceOff                              int      `json:"wins_face_off"`
				WeeklyWinsFaceOffHunter                  int      `json:"weeklyWins_face_off_Hunter"`
				WeeklyDeathsHunter                       int      `json:"weeklyDeaths_Hunter"`
				WeeklyWinsFaceOff                        int      `json:"weeklyWins_face_off"`
				DeathsHunter                             int      `json:"deaths_Hunter"`
				WeeklyFinalKillsHunterB                  int      `json:"weekly_finalKills_Hunter_b"`
				MonthlyFinalKillsHunterA                 int      `json:"monthly_finalKills_Hunter_a"`
				WeeklyLossesFaceOffBlaze                 int      `json:"weeklyLosses_face_off_Blaze"`
				LossesFaceOffBlaze                       int      `json:"losses_face_off_Blaze"`
				AssistsSpider                            int      `json:"assists_Spider"`
				WitherCoins                              int      `json:"witherCoins"`
				SmileyKills                              string   `json:"smiley_kills"`
				MonthlyFinalKillsB                       int      `json:"monthly_finalKills_b"`
				WeeklyWinsFaceOffHerobrine               int      `json:"weeklyWins_face_off_Herobrine"`
				WinsFaceOffHerobrine                     int      `json:"wins_face_off_Herobrine"`
				WeeklyFinalKillsA                        int      `json:"weekly_finalKills_a"`
				MonthlyFinalKillsA                       int      `json:"monthly_finalKills_a"`
				WeeklyFinalKillsB                        int      `json:"weekly_finalKills_b"`
				WeeklyWinsArcanist                       int      `json:"weeklyWins_Arcanist"`
				FinalKillsArcanist                       int      `json:"finalKills_Arcanist"`
				WinsFaceOffArcanist                      int      `json:"wins_face_off_Arcanist"`
				WeeklyFinalKillsArcanistA                int      `json:"weekly_finalKills_Arcanist_a"`
				WeeklyDeathsArcanist                     int      `json:"weeklyDeaths_Arcanist"`
				WeeklyKillsArcanist                      int      `json:"weeklyKills_Arcanist"`
				MonthlyFinalKillsArcanistB               int      `json:"monthly_finalKills_Arcanist_b"`
				KillsArcanist                            int      `json:"kills_Arcanist"`
				DeathsArcanist                           int      `json:"deaths_Arcanist"`
				WinsArcanist                             int      `json:"wins_Arcanist"`
				FinalAssistsArcanist                     int      `json:"finalAssists_Arcanist"`
				WeeklyWinsFaceOffArcanist                int      `json:"weeklyWins_face_off_Arcanist"`
				AssistsArcanist                          int      `json:"assists_Arcanist"`
				WeeklyFinalKillsBlazeB                   int      `json:"weekly_finalKills_Blaze_b"`
				MonthlyFinalKillsBlazeB                  int      `json:"monthly_finalKills_Blaze_b"`
				FinalAssistsBlaze                        int      `json:"finalAssists_Blaze"`
				AssistsBlaze                             int      `json:"assists_Blaze"`
				WeeklyWinsFaceOffBlaze                   int      `json:"weeklyWins_face_off_Blaze"`
				WinsFaceOffBlaze                         int      `json:"wins_face_off_Blaze"`
				SpiderEnderchestLevel                    int      `json:"spider_enderchest_level"`
				HunterFinalKills                         int      `json:"hunter_final_kills"`
				HunterTotalFinalKills                    int      `json:"hunter_total_final_kills"`
				ZombieKills                              int      `json:"zombie_kills"`
				PigmanTotalFinalKills                    int      `json:"pigman_total_final_kills"`
				ZombieTotalFinalKillsStandard            int      `json:"zombie_total_final_kills_standard"`
				EndermanFinalKillsStandard               int      `json:"enderman_final_kills_standard"`
				ArcanistTotalFinalKills                  int      `json:"arcanist_total_final_kills"`
				SkeletonFinalKills                       int      `json:"skeleton_final_kills"`
				HerobrineEnderchestLevel                 int      `json:"herobrine_enderchest_level"`
				BlazeFinalKillsStandard                  int      `json:"blaze_final_kills_standard"`
				BlazeTotalFinalKillsStandard             int      `json:"blaze_total_final_kills_standard"`
				HunterTotalFinalKillsStandard            int      `json:"hunter_total_final_kills_standard"`
				EndermanTotalFinalKillsStandard          int      `json:"enderman_total_final_kills_standard"`
				PirateTotalFinalKills                    int      `json:"pirate_total_final_kills"`
				SkeletonFinalKillsStandard               int      `json:"skeleton_final_kills_standard"`
				CreeperTotalFinalKills                   int      `json:"creeper_total_final_kills"`
				ZombieWins                               int      `json:"zombie_wins"`
				ArcanistWins                             int      `json:"arcanist_wins"`
				CreeperFinalKillsStandard                int      `json:"creeper_final_kills_standard"`
				SkeletonTotalFinalKillsStandard          int      `json:"skeleton_total_final_kills_standard"`
				CreeperLosses                            int      `json:"creeper_losses"`
				HunterDeaths                             int      `json:"hunter_deaths"`
				ArcanistKills                            int      `json:"arcanist_kills"`
				HerobrineTotalFinalKillsStandard         int      `json:"herobrine_total_final_kills_standard"`
				BlazeWins                                int      `json:"blaze_wins"`
				CreeperFinalKills                        int      `json:"creeper_final_kills"`
				ArcanistDeaths                           int      `json:"arcanist_deaths"`
				ArcanistLosses                           int      `json:"arcanist_losses"`
				GolemTotalFinalKillsStandard             int      `json:"golem_total_final_kills_standard"`
				HerobrineKills                           int      `json:"herobrine_kills"`
				ZombieDeaths                             int      `json:"zombie_deaths"`
				PirateTotalFinalKillsStandard            int      `json:"pirate_total_final_kills_standard"`
				ShamanWins                               int      `json:"shaman_wins"`
				SpiderTotalFinalKills                    int      `json:"spider_total_final_kills"`
				ArcanistFinalKills                       int      `json:"arcanist_final_kills"`
				ShamanTotalFinalKillsStandard            int      `json:"shaman_total_final_kills_standard"`
				HunterWins                               int      `json:"hunter_wins"`
				ShamanFinalKills                         int      `json:"shaman_final_kills"`
				BlazeFinalKills                          int      `json:"blaze_final_kills"`
				GolemKills                               int      `json:"golem_kills"`
				GolemWins                                int      `json:"golem_wins"`
				CreeperKills                             int      `json:"creeper_kills"`
				EndermanWins                             int      `json:"enderman_wins"`
				GolemDeaths                              int      `json:"golem_deaths"`
				ZombieTotalFinalKills                    int      `json:"zombie_total_final_kills"`
				ShamanKills                              int      `json:"shaman_kills"`
				GolemTotalFinalKills                     int      `json:"golem_total_final_kills"`
				ArcanistTotalFinalKillsStandard          int      `json:"arcanist_total_final_kills_standard"`
				PirateDeaths                             int      `json:"pirate_deaths"`
				PigmanTotalFinalKillsStandard            int      `json:"pigman_total_final_kills_standard"`
				CreeperWins                              int      `json:"creeper_wins"`
				SpiderFinalKillsStandard                 int      `json:"spider_final_kills_standard"`
				BlazeKills                               int      `json:"blaze_kills"`
				HerobrineTotalFinalKills                 int      `json:"herobrine_total_final_kills"`
				ArcanistFinalKillsStandard               int      `json:"arcanist_final_kills_standard"`
				HerobrineFinalKillsStandard              int      `json:"herobrine_final_kills_standard"`
				PigmanWins                               int      `json:"pigman_wins"`
				HunterKills                              int      `json:"hunter_kills"`
				EndermanKills                            int      `json:"enderman_kills"`
				PigmanDeaths                             int      `json:"pigman_deaths"`
				PigmanLosses                             int      `json:"pigman_losses"`
				ZombieFinalKills                         int      `json:"zombie_final_kills"`
				PigmanFinalKills                         int      `json:"pigman_final_kills"`
				SpiderKills                              int      `json:"spider_kills"`
				ShamanLosses                             int      `json:"shaman_losses"`
				EndermanFinalKills                       int      `json:"enderman_final_kills"`
				PirateLosses                             int      `json:"pirate_losses"`
				SpiderWins                               int      `json:"spider_wins"`
				PirateKills                              int      `json:"pirate_kills"`
				SkeletonWins                             int      `json:"skeleton_wins"`
				PirateFinalKills                         int      `json:"pirate_final_kills"`
				SpiderLosses                             int      `json:"spider_losses"`
				CreeperTotalFinalKillsStandard           int      `json:"creeper_total_final_kills_standard"`
				SkeletonDeaths                           int      `json:"skeleton_deaths"`
				ZombieLosses                             int      `json:"zombie_losses"`
				BlazeTotalFinalKills                     int      `json:"blaze_total_final_kills"`
				SpiderFinalKills                         int      `json:"spider_final_kills"`
				SkeletonKills                            int      `json:"skeleton_kills"`
				HunterLosses                             int      `json:"hunter_losses"`
				HerobrineLosses                          int      `json:"herobrine_losses"`
				ShamanTotalFinalKills                    int      `json:"shaman_total_final_kills"`
				CreeperDeaths                            int      `json:"creeper_deaths"`
				EndermanTotalFinalKills                  int      `json:"enderman_total_final_kills"`
				GolemFinalKills                          int      `json:"golem_final_kills"`
				PirateFinalKillsStandard                 int      `json:"pirate_final_kills_standard"`
				EndermanLosses                           int      `json:"enderman_losses"`
				ShamanFinalKillsStandard                 int      `json:"shaman_final_kills_standard"`
				PigmanFinalKillsStandard                 int      `json:"pigman_final_kills_standard"`
				SkeletonLosses                           int      `json:"skeleton_losses"`
				ShamanDeaths                             int      `json:"shaman_deaths"`
				HunterFinalKillsStandard                 int      `json:"hunter_final_kills_standard"`
				PigmanKills                              int      `json:"pigman_kills"`
				SkeletonTotalFinalKills                  int      `json:"skeleton_total_final_kills"`
				GolemEnderchestLevel                     int      `json:"golem_enderchest_level"`
				HerobrineWins                            int      `json:"herobrine_wins"`
				ZombieFinalKillsStandard                 int      `json:"zombie_final_kills_standard"`
				BlazeLosses                              int      `json:"blaze_losses"`
				BlazeDeaths                              int      `json:"blaze_deaths"`
				SpiderDeaths                             int      `json:"spider_deaths"`
				GolemLosses                              int      `json:"golem_losses"`
				HerobrineDeaths                          int      `json:"herobrine_deaths"`
				SpiderTotalFinalKillsStandard            int      `json:"spider_total_final_kills_standard"`
				GolemFinalKillsStandard                  int      `json:"golem_final_kills_standard"`
				HerobrineFinalKills                      int      `json:"herobrine_final_kills"`
				EndermanDeaths                           int      `json:"enderman_deaths"`
				HerobrineAKillsStandard                  int      `json:"herobrine_a_kills_standard"`
				HerobrineADefenderAssists                int      `json:"herobrine_a_defender_assists"`
				HerobrineATotalKillsStandard             int      `json:"herobrine_a_total_kills_standard"`
				HerobrineAssists                         int      `json:"herobrine_assists"`
				HerobrineAActivationsStandard            int      `json:"herobrine_a_activations_standard"`
				HerobrineKillsMeleeBehindStandard        int      `json:"herobrine_kills_melee_behind_standard"`
				DefenderAssistsStandard                  int      `json:"defender_assists_standard"`
				HerobrineActivationsDeathmatchStandard   int      `json:"herobrine_activations_deathmatch_standard"`
				HerobrineBlocksBrokenStandard            int      `json:"herobrine_blocks_broken_standard"`
				HerobrineMetersFallenStandard            int      `json:"herobrine_meters_fallen_standard"`
				HerobrineGamesPlayed                     int      `json:"herobrine_games_played"`
				HerobrineIronOreBroken                   int      `json:"herobrine_iron_ore_broken"`
				TotalDeaths                              int      `json:"total_deaths"`
				HerobrineKillsStandard                   int      `json:"herobrine_kills_standard"`
				HerobrineBlocksBroken                    int      `json:"herobrine_blocks_broken"`
				LossesStandard                           int      `json:"losses_standard"`
				HerobrineKillsMeleeBehind                int      `json:"herobrine_kills_melee_behind"`
				HerobrineFinalDeathsStandard             int      `json:"herobrine_final_deaths_standard"`
				TotalKills                               int      `json:"total_kills"`
				ActivationsStandard                      int      `json:"activations_standard"`
				HerobrineADefenderKillsStandard          int      `json:"herobrine_a_defender_kills_standard"`
				HerobrineArrowsFiredStandard             int      `json:"herobrine_arrows_fired_standard"`
				TreasuresFound                           int      `json:"treasures_found"`
				HerobrineAKills                          int      `json:"herobrine_a_kills"`
				ActivationsDeathmatchStandard            int      `json:"activations_deathmatch_standard"`
				ArrowsFiredStandard                      int      `json:"arrows_fired_standard"`
				HerobrineADamageDealt                    int      `json:"herobrine_a_damage_dealt"`
				HerobrineDamageDealtStandard             int      `json:"herobrine_damage_dealt_standard"`
				HerobrineDefenderKills                   int      `json:"herobrine_defender_kills"`
				HerobrineBlocksPlaced                    int      `json:"herobrine_blocks_placed"`
				TimePlayed                               int      `json:"time_played"`
				DefenderKillsStandard                    int      `json:"defender_kills_standard"`
				HerobrineBlocksPlacedPreparation         int      `json:"herobrine_blocks_placed_preparation"`
				HerobrineLossesStandard                  int      `json:"herobrine_losses_standard"`
				HerobrineATotalKills                     int      `json:"herobrine_a_total_kills"`
				HerobrineDefenderAssists                 int      `json:"herobrine_defender_assists"`
				FinalDeaths                              int      `json:"final_deaths"`
				GamesPlayed                              int      `json:"games_played"`
				HerobrineKillsMelee                      int      `json:"herobrine_kills_melee"`
				BlocksPlacedStandard                     int      `json:"blocks_placed_standard"`
				HerobrineTotalDeaths                     int      `json:"herobrine_total_deaths"`
				BlocksPlacedPreparation                  int      `json:"blocks_placed_preparation"`
				HerobrineFinalDeaths                     int      `json:"herobrine_final_deaths"`
				HerobrineADefenderKills                  int      `json:"herobrine_a_defender_kills"`
				HerobrineADefenderAssistsStandard        int      `json:"herobrine_a_defender_assists_standard"`
				BlocksPlacedPreparationStandard          int      `json:"blocks_placed_preparation_standard"`
				HerobrineBlocksPlacedStandard            int      `json:"herobrine_blocks_placed_standard"`
				DamageDealt                              int      `json:"damage_dealt"`
				HerobrineActivations                     int      `json:"herobrine_activations"`
				TotalKillsStandard                       int      `json:"total_kills_standard"`
				HerobrineAActivationsDeathmatchStandard  int      `json:"herobrine_a_activations_deathmatch_standard"`
				TimePlayedStandard                       int      `json:"time_played_standard"`
				HerobrineAAssists                        int      `json:"herobrine_a_assists"`
				KillsMeleeStandard                       int      `json:"kills_melee_standard"`
				HerobrineTreasuresFoundStandard          int      `json:"herobrine_treasures_found_standard"`
				HerobrineIronOreBrokenStandard           int      `json:"herobrine_iron_ore_broken_standard"`
				KillsStandard                            int      `json:"kills_standard"`
				IronOreBrokenStandard                    int      `json:"iron_ore_broken_standard"`
				TreasuresFoundStandard                   int      `json:"treasures_found_standard"`
				HerobrineDefenderAssistsStandard         int      `json:"herobrine_defender_assists_standard"`
				BlocksBrokenStandard                     int      `json:"blocks_broken_standard"`
				HerobrineMetersWalkedSpeed               int      `json:"herobrine_meters_walked_speed"`
				DamageDealtStandard                      int      `json:"damage_dealt_standard"`
				MetersWalkedSpeedStandard                int      `json:"meters_walked_speed_standard"`
				HerobrineAAssistsStandard                int      `json:"herobrine_a_assists_standard"`
				HerobrineArrowsFired                     int      `json:"herobrine_arrows_fired"`
				HerobrineMetersWalked                    int      `json:"herobrine_meters_walked"`
				FinalDeathsStandard                      int      `json:"final_deaths_standard"`
				KillsMelee                               int      `json:"kills_melee"`
				HerobrineTreasuresFound                  int      `json:"herobrine_treasures_found"`
				DefenderKills                            int      `json:"defender_kills"`
				HerobrineDamageDealt                     int      `json:"herobrine_damage_dealt"`
				HerobrineBlocksPlacedPreparationStandard int      `json:"herobrine_blocks_placed_preparation_standard"`
				HerobrineTimePlayed                      int      `json:"herobrine_time_played"`
				MetersWalked                             int      `json:"meters_walked"`
				HerobrineActivationsStandard             int      `json:"herobrine_activations_standard"`
				MetersFallenStandard                     int      `json:"meters_fallen_standard"`
				KillsMeleeBehind                         int      `json:"kills_melee_behind"`
				HerobrineMetersWalkedSpeedStandard       int      `json:"herobrine_meters_walked_speed_standard"`
				BlocksBroken                             int      `json:"blocks_broken"`
				Activations                              int      `json:"activations"`
				AssistsStandard                          int      `json:"assists_standard"`
				BlocksPlaced                             int      `json:"blocks_placed"`
				HerobrineTotalDeathsStandard             int      `json:"herobrine_total_deaths_standard"`
				HerobrineAActivationsDeathmatch          int      `json:"herobrine_a_activations_deathmatch"`
				MetersWalkedSpeed                        int      `json:"meters_walked_speed"`
				ArrowsFired                              int      `json:"arrows_fired"`
				IronOreBroken                            int      `json:"iron_ore_broken"`
				TotalDeathsStandard                      int      `json:"total_deaths_standard"`
				HerobrineDefenderKillsStandard           int      `json:"herobrine_defender_kills_standard"`
				HerobrineTotalKillsStandard              int      `json:"herobrine_total_kills_standard"`
				HerobrineKillsMeleeStandard              int      `json:"herobrine_kills_melee_standard"`
				KillsMeleeBehindStandard                 int      `json:"kills_melee_behind_standard"`
				HerobrineGamesPlayedStandard             int      `json:"herobrine_games_played_standard"`
				HerobrineMetersFallen                    int      `json:"herobrine_meters_fallen"`
				HerobrineMetersWalkedStandard            int      `json:"herobrine_meters_walked_standard"`
				DefenderAssists                          int      `json:"defender_assists"`
				MetersFallen                             int      `json:"meters_fallen"`
				HerobrineAKillsMeleeStandard             int      `json:"herobrine_a_kills_melee_standard"`
				HerobrineADamageDealtStandard            int      `json:"herobrine_a_damage_dealt_standard"`
				HerobrineAssistsStandard                 int      `json:"herobrine_assists_standard"`
				HerobrineTimePlayedStandard              int      `json:"herobrine_time_played_standard"`
				MetersWalkedStandard                     int      `json:"meters_walked_standard"`
				HerobrineActivationsDeathmatch           int      `json:"herobrine_activations_deathmatch"`
				HerobrineAKillsMelee                     int      `json:"herobrine_a_kills_melee"`
				ActivationsDeathmatch                    int      `json:"activations_deathmatch"`
				GamesPlayedStandard                      int      `json:"games_played_standard"`
				HerobrineAActivations                    int      `json:"herobrine_a_activations"`
				HerobrineTotalKills                      int      `json:"herobrine_total_kills"`
				PhoenixInventory                         struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
					Num3 string `json:"3"`
					Num4 string `json:"4"`
					Num5 string `json:"5"`
					Num6 string `json:"6"`
					Num7 string `json:"7"`
					Num8 string `json:"8"`
					Num9 string `json:"9"`
				} `json:"phoenixInventory"`
				PhoenixMetersWalked                     int    `json:"phoenix_meters_walked"`
				PhoenixBlocksPlacedPreparation          int    `json:"phoenix_blocks_placed_preparation"`
				BlocksPlacedFaceOff                     int    `json:"blocks_placed_face_off"`
				GamesPlayedFaceOff                      int    `json:"games_played_face_off"`
				PhoenixMetersFallen                     int    `json:"phoenix_meters_fallen"`
				PhoenixGamesPlayedFaceOff               int    `json:"phoenix_games_played_face_off"`
				PhoenixBlocksPlacedPreparationFaceOff   int    `json:"phoenix_blocks_placed_preparation_face_off"`
				PhoenixBlocksPlacedFaceOff              int    `json:"phoenix_blocks_placed_face_off"`
				MetersFallenFaceOff                     int    `json:"meters_fallen_face_off"`
				PhoenixWinsFaceOff                      int    `json:"phoenix_wins_face_off"`
				PhoenixMetersWalkedFaceOff              int    `json:"phoenix_meters_walked_face_off"`
				BlocksPlacedPreparationFaceOff          int    `json:"blocks_placed_preparation_face_off"`
				PhoenixGamesPlayed                      int    `json:"phoenix_games_played"`
				PhoenixMetersFallenFaceOff              int    `json:"phoenix_meters_fallen_face_off"`
				PhoenixWins                             int    `json:"phoenix_wins"`
				PhoenixBlocksPlaced                     int    `json:"phoenix_blocks_placed"`
				MetersWalkedFaceOff                     int    `json:"meters_walked_face_off"`
				PhoenixLossesStandard                   int    `json:"phoenix_losses_standard"`
				PhoenixMetersWalkedStandard             int    `json:"phoenix_meters_walked_standard"`
				PhoenixArrowsFired                      int    `json:"phoenix_arrows_fired"`
				PhoenixMetersFallenStandard             int    `json:"phoenix_meters_fallen_standard"`
				PhoenixLosses                           int    `json:"phoenix_losses"`
				PhoenixArrowsFiredStandard              int    `json:"phoenix_arrows_fired_standard"`
				PhoenixGamesPlayedStandard              int    `json:"phoenix_games_played_standard"`
				PhoenixWinsStandard                     int    `json:"phoenix_wins_standard"`
				FinalKills                              int    `json:"final_kills"`
				SkeletonWinsStandard                    int    `json:"skeleton_wins_standard"`
				EndermanWinsStandard                    int    `json:"enderman_wins_standard"`
				ZombieWinsStandard                      int    `json:"zombie_wins_standard"`
				BlazeWinsStandard                       int    `json:"blaze_wins_standard"`
				ArcanistWinsStandard                    int    `json:"arcanist_wins_standard"`
				HunterWinsStandard                      int    `json:"hunter_wins_standard"`
				HerobrineWinsStandard                   int    `json:"herobrine_wins_standard"`
				ShamanWinsStandard                      int    `json:"shaman_wins_standard"`
				FinalKillsStandard                      int    `json:"final_kills_standard"`
				CreeperWinsStandard                     int    `json:"creeper_wins_standard"`
				TotalFinalKillsStandard                 int    `json:"total_final_kills_standard"`
				SpiderWinsStandard                      int    `json:"spider_wins_standard"`
				PigmanWinsStandard                      int    `json:"pigman_wins_standard"`
				GolemWinsStandard                       int    `json:"golem_wins_standard"`
				TotalFinalKills                         int    `json:"total_final_kills"`
				WerewolfBlocksPlacedPreparationStandard int    `json:"werewolf_blocks_placed_preparation_standard"`
				WerewolfPotionsDrunkStandard            int    `json:"werewolf_potions_drunk_standard"`
				WerewolfMetersWalkedSpeed               int    `json:"werewolf_meters_walked_speed"`
				WerewolfGamesPlayed                     int    `json:"werewolf_games_played"`
				PotionsDrunkStandard                    int    `json:"potions_drunk_standard"`
				WerewolfBlocksPlacedStandard            int    `json:"werewolf_blocks_placed_standard"`
				WerewolfLossesStandard                  int    `json:"werewolf_losses_standard"`
				WerewolfLosses                          int    `json:"werewolf_losses"`
				WerewolfMetersWalkedStandard            int    `json:"werewolf_meters_walked_standard"`
				WerewolfTreasuresFound                  int    `json:"werewolf_treasures_found"`
				WerewolfTimePlayed                      int    `json:"werewolf_time_played"`
				WerewolfIronOreBrokenStandard           int    `json:"werewolf_iron_ore_broken_standard"`
				WerewolfMetersFallenStandard            int    `json:"werewolf_meters_fallen_standard"`
				WerewolfBlocksBrokenStandard            int    `json:"werewolf_blocks_broken_standard"`
				WerewolfTreasuresFoundStandard          int    `json:"werewolf_treasures_found_standard"`
				WerewolfMetersWalkedSpeedStandard       int    `json:"werewolf_meters_walked_speed_standard"`
				WerewolfTimePlayedStandard              int    `json:"werewolf_time_played_standard"`
				WerewolfPotionsDrunk                    int    `json:"werewolf_potions_drunk"`
				WerewolfGamesPlayedStandard             int    `json:"werewolf_games_played_standard"`
				WerewolfMetersFallen                    int    `json:"werewolf_meters_fallen"`
				WerewolfMetersWalked                    int    `json:"werewolf_meters_walked"`
				WerewolfBlocksBroken                    int    `json:"werewolf_blocks_broken"`
				PotionsDrunk                            int    `json:"potions_drunk"`
				WerewolfIronOreBroken                   int    `json:"werewolf_iron_ore_broken"`
				WerewolfBlocksPlaced                    int    `json:"werewolf_blocks_placed"`
				WerewolfBlocksPlacedPreparation         int    `json:"werewolf_blocks_placed_preparation"`
				WerewolfDefenderKillsStandard           int    `json:"werewolf_defender_kills_standard"`
				WerewolfTotalDeathsStandard             int    `json:"werewolf_total_deaths_standard"`
				DiamondOreBrokenStandard                int    `json:"diamond_ore_broken_standard"`
				WerewolfTotalDeaths                     int    `json:"werewolf_total_deaths"`
				WerewolfAAssistsStandard                int    `json:"werewolf_a_assists_standard"`
				WerewolfKills                           int    `json:"werewolf_kills"`
				WerewolfBAssists                        int    `json:"werewolf_b_assists"`
				WoodChopped                             int    `json:"wood_chopped"`
				WerewolfDefenderAssistsStandard         int    `json:"werewolf_defender_assists_standard"`
				WerewolfAActivationsStandard            int    `json:"werewolf_a_activations_standard"`
				WoodChoppedStandard                     int    `json:"wood_chopped_standard"`
				WerewolfAssists                         int    `json:"werewolf_assists"`
				WerewolfADamageDealt                    int    `json:"werewolf_a_damage_dealt"`
				WerewolfWoodChopped                     int    `json:"werewolf_wood_chopped"`
				WerewolfDeathsStandard                  int    `json:"werewolf_deaths_standard"`
				WerewolfDamageDealtStandard             int    `json:"werewolf_damage_dealt_standard"`
				WerewolfDiamondOreBrokenStandard        int    `json:"werewolf_diamond_ore_broken_standard"`
				WerewolfWoodChoppedStandard             int    `json:"werewolf_wood_chopped_standard"`
				WerewolfDefenderAssists                 int    `json:"werewolf_defender_assists"`
				WerewolfDamageDealt                     int    `json:"werewolf_damage_dealt"`
				WerewolfGActivationsStandard            int    `json:"werewolf_g_activations_standard"`
				WerewolfTotalKills                      int    `json:"werewolf_total_kills"`
				DiamondOreBroken                        int    `json:"diamond_ore_broken"`
				WerewolfATotalKills                     int    `json:"werewolf_a_total_kills"`
				WerewolfBTotalKills                     int    `json:"werewolf_b_total_kills"`
				WerewolfADefenderAssists                int    `json:"werewolf_a_defender_assists"`
				WerewolfBAssistsStandard                int    `json:"werewolf_b_assists_standard"`
				WerewolfKillsMeleeStandard              int    `json:"werewolf_kills_melee_standard"`
				WerewolfADamageDealtStandard            int    `json:"werewolf_a_damage_dealt_standard"`
				WerewolfDeaths                          int    `json:"werewolf_deaths"`
				WerewolfKillsStandard                   int    `json:"werewolf_kills_standard"`
				WerewolfAssistsStandard                 int    `json:"werewolf_assists_standard"`
				WerewolfDiamondOreBroken                int    `json:"werewolf_diamond_ore_broken"`
				DeathsStandard                          int    `json:"deaths_standard"`
				WerewolfATotalKillsStandard             int    `json:"werewolf_a_total_kills_standard"`
				WerewolfDefenderKills                   int    `json:"werewolf_defender_kills"`
				WerewolfActivationsStandard             int    `json:"werewolf_activations_standard"`
				WerewolfGActivations                    int    `json:"werewolf_g_activations"`
				WerewolfTotalKillsStandard              int    `json:"werewolf_total_kills_standard"`
				WerewolfBTotalKillsStandard             int    `json:"werewolf_b_total_kills_standard"`
				WerewolfKillsMelee                      int    `json:"werewolf_kills_melee"`
				WerewolfAAssists                        int    `json:"werewolf_a_assists"`
				WerewolfAActivations                    int    `json:"werewolf_a_activations"`
				WerewolfADefenderAssistsStandard        int    `json:"werewolf_a_defender_assists_standard"`
				WerewolfActivations                     int    `json:"werewolf_activations"`
				ArcanistFinalAssists                    int    `json:"arcanist_final_assists"`
				SkeletonFinalAssists                    int    `json:"skeleton_final_assists"`
				HunterFinalAssists                      int    `json:"hunter_final_assists"`
				PirateFinalAssists                      int    `json:"pirate_final_assists"`
				GolemFinalAssists                       int    `json:"golem_final_assists"`
				BlazeFinalAssists                       int    `json:"blaze_final_assists"`
				SpiderFinalAssists                      int    `json:"spider_final_assists"`
				HerobrineFinalAssists                   int    `json:"herobrine_final_assists"`
				ActiveChallengeMap                      string `json:"activeChallengeMap"`
				WinsStandard                            int    `json:"wins_standard"`
				WerewolfWinsStandard                    int    `json:"werewolf_wins_standard"`
				WerewolfWins                            int    `json:"werewolf_wins"`
				HunterDeathsStandard                    int    `json:"hunter_deaths_standard"`
				HunterBlocksPlacedPreparation           int    `json:"hunter_blocks_placed_preparation"`
				HunterArrowsFiredStandard               int    `json:"hunter_arrows_fired_standard"`
				HunterTimePlayed                        int    `json:"hunter_time_played"`
				HunterMetersFallenStandard              int    `json:"hunter_meters_fallen_standard"`
				HunterMetersWalkedStandard              int    `json:"hunter_meters_walked_standard"`
				HunterBlocksPlaced                      int    `json:"hunter_blocks_placed"`
				HunterBlocksPlacedPreparationStandard   int    `json:"hunter_blocks_placed_preparation_standard"`
				HunterGamesPlayedStandard               int    `json:"hunter_games_played_standard"`
				HunterTimePlayedStandard                int    `json:"hunter_time_played_standard"`
				HunterMetersWalked                      int    `json:"hunter_meters_walked"`
				HunterLossesStandard                    int    `json:"hunter_losses_standard"`
				HunterTotalDeathsStandard               int    `json:"hunter_total_deaths_standard"`
				HunterArrowsFired                       int    `json:"hunter_arrows_fired"`
				HunterMetersFallen                      int    `json:"hunter_meters_fallen"`
				HunterBlocksPlacedStandard              int    `json:"hunter_blocks_placed_standard"`
				HunterGamesPlayed                       int    `json:"hunter_games_played"`
				HunterTotalDeaths                       int    `json:"hunter_total_deaths"`
				HunterActivationsDeathmatch             int    `json:"hunter_activations_deathmatch"`
				HunterKillsMeleeBehind                  int    `json:"hunter_kills_melee_behind"`
				HunterFinalKillsMeleeBehindStandard     int    `json:"hunter_final_kills_melee_behind_standard"`
				HunterTotalKills                        int    `json:"hunter_total_kills"`
				HunterBlocksBroken                      int    `json:"hunter_blocks_broken"`
				HunterAActivationsDeathmatchStandard    int    `json:"hunter_a_activations_deathmatch_standard"`
				FinalKillsMelee                         int    `json:"final_kills_melee"`
				BreadEatenStandard                      int    `json:"bread_eaten_standard"`
				HunterAssistsStandard                   int    `json:"hunter_assists_standard"`
				HunterBlocksBrokenStandard              int    `json:"hunter_blocks_broken_standard"`
				HunterDefenderAssistsStandard           int    `json:"hunter_defender_assists_standard"`
				FoodEatenStandard                       int    `json:"food_eaten_standard"`
				GoldenApplesEaten                       int    `json:"golden_apples_eaten"`
				HunterAActivationsDeathmatch            int    `json:"hunter_a_activations_deathmatch"`
				HunterBreadEaten                        int    `json:"hunter_bread_eaten"`
				HunterKillsStandard                     int    `json:"hunter_kills_standard"`
				FinalAssists                            int    `json:"final_assists"`
				HunterArrowsHitStandard                 int    `json:"hunter_arrows_hit_standard"`
				FinalKillsMeleeBehind                   int    `json:"final_kills_melee_behind"`
				HunterAActivationsStandard              int    `json:"hunter_a_activations_standard"`
				HunterDefenderKills                     int    `json:"hunter_defender_kills"`
				FinalKillsMeleeBehindStandard           int    `json:"final_kills_melee_behind_standard"`
				HunterBreadEatenStandard                int    `json:"hunter_bread_eaten_standard"`
				FoodEaten                               int    `json:"food_eaten"`
				HunterMetersWalkedSpeedStandard         int    `json:"hunter_meters_walked_speed_standard"`
				HunterFoodEaten                         int    `json:"hunter_food_eaten"`
				HunterSteaksEaten                       int    `json:"hunter_steaks_eaten"`
				HunterGoldenApplesEatenStandard         int    `json:"hunter_golden_apples_eaten_standard"`
				HunterKillsMeleeStandard                int    `json:"hunter_kills_melee_standard"`
				HunterTotalKillsStandard                int    `json:"hunter_total_kills_standard"`
				HunterActivationsStandard               int    `json:"hunter_activations_standard"`
				HunterFinalAssistsStandard              int    `json:"hunter_final_assists_standard"`
				FinalAssistsStandard                    int    `json:"final_assists_standard"`
				HunterActivations                       int    `json:"hunter_activations"`
				HunterArrowsHit                         int    `json:"hunter_arrows_hit"`
				HunterMetersWalkedSpeed                 int    `json:"hunter_meters_walked_speed"`
				BreadEaten                              int    `json:"bread_eaten"`
				SteaksEaten                             int    `json:"steaks_eaten"`
				SteaksEatenStandard                     int    `json:"steaks_eaten_standard"`
				HunterKillsMeleeBehindStandard          int    `json:"hunter_kills_melee_behind_standard"`
				HunterAActivations                      int    `json:"hunter_a_activations"`
				HunterActivationsDeathmatchStandard     int    `json:"hunter_activations_deathmatch_standard"`
				HunterDefenderKillsStandard             int    `json:"hunter_defender_kills_standard"`
				FinalKillsMeleeStandard                 int    `json:"final_kills_melee_standard"`
				HunterKillsMelee                        int    `json:"hunter_kills_melee"`
				HunterFinalKillsMelee                   int    `json:"hunter_final_kills_melee"`
				ArrowsHitStandard                       int    `json:"arrows_hit_standard"`
				HunterDefenderAssists                   int    `json:"hunter_defender_assists"`
				HunterFinalKillsMeleeStandard           int    `json:"hunter_final_kills_melee_standard"`
				ArrowsHit                               int    `json:"arrows_hit"`
				HunterGoldenApplesEaten                 int    `json:"hunter_golden_apples_eaten"`
				HunterAssists                           int    `json:"hunter_assists"`
				HunterFinalKillsMeleeBehind             int    `json:"hunter_final_kills_melee_behind"`
				GoldenApplesEatenStandard               int    `json:"golden_apples_eaten_standard"`
				HunterFoodEatenStandard                 int    `json:"hunter_food_eaten_standard"`
				HunterSteaksEatenStandard               int    `json:"hunter_steaks_eaten_standard"`
				DreadlordInventory                      struct {
					Num0 string `json:"0"`
					Num2 string `json:"2"`
					Num3 string `json:"3"`
					Num4 string `json:"4"`
					Num5 string `json:"5"`
					Num7 string `json:"7"`
					Num8 string `json:"8"`
					Num9 string `json:"9"`
				} `json:"dreadlordInventory"`
				DreadlordDamageDealtStandard             int `json:"dreadlord_damage_dealt_standard"`
				DreadlordTimePlayed                      int `json:"dreadlord_time_played"`
				DreadlordAssistsStandard                 int `json:"dreadlord_assists_standard"`
				DreadlordActivationsDeathmatchStandard   int `json:"dreadlord_activations_deathmatch_standard"`
				DreadlordTreasuresFound                  int `json:"dreadlord_treasures_found"`
				DreadlordAActivationsDeathmatchStandard  int `json:"dreadlord_a_activations_deathmatch_standard"`
				DreadlordMetersWalkedStandard            int `json:"dreadlord_meters_walked_standard"`
				DreadlordAFinalAssistsStandard           int `json:"dreadlord_a_final_assists_standard"`
				DreadlordATotalFinalKillsStandard        int `json:"dreadlord_a_total_final_kills_standard"`
				DreadlordTotalFinalKillsStandard         int `json:"dreadlord_total_final_kills_standard"`
				DreadlordAFinalKillsMelee                int `json:"dreadlord_a_final_kills_melee"`
				DreadlordActivationsDeathmatch           int `json:"dreadlord_activations_deathmatch"`
				DreadlordFinalDeaths                     int `json:"dreadlord_final_deaths"`
				DreadlordTotalKills                      int `json:"dreadlord_total_kills"`
				DreadlordKills                           int `json:"dreadlord_kills"`
				DreadlordDefenderKillsStandard           int `json:"dreadlord_defender_kills_standard"`
				DreadlordFinalDeathsStandard             int `json:"dreadlord_final_deaths_standard"`
				DreadlordADamageDealtStandard            int `json:"dreadlord_a_damage_dealt_standard"`
				DreadlordTimePlayedStandard              int `json:"dreadlord_time_played_standard"`
				DreadlordKillsMeleeStandard              int `json:"dreadlord_kills_melee_standard"`
				DreadlordATotalFinalKills                int `json:"dreadlord_a_total_final_kills"`
				DreadlordFinalKillsMelee                 int `json:"dreadlord_final_kills_melee"`
				DreadlordTotalKillsStandard              int `json:"dreadlord_total_kills_standard"`
				DreadlordMetersWalked                    int `json:"dreadlord_meters_walked"`
				WitherDamageStandard                     int `json:"wither_damage_standard"`
				DreadlordAAssistsStandard                int `json:"dreadlord_a_assists_standard"`
				DreadlordMetersFallenStandard            int `json:"dreadlord_meters_fallen_standard"`
				DreadlordADamageDealt                    int `json:"dreadlord_a_damage_dealt"`
				DreadlordKillsMelee                      int `json:"dreadlord_kills_melee"`
				DreadlordAFinalKillsStandard             int `json:"dreadlord_a_final_kills_standard"`
				DreadlordGamesPlayedStandard             int `json:"dreadlord_games_played_standard"`
				DreadlordDeaths                          int `json:"dreadlord_deaths"`
				DreadlordAFinalAssists                   int `json:"dreadlord_a_final_assists"`
				DreadlordFinalAssists                    int `json:"dreadlord_final_assists"`
				DreadlordBlocksPlacedPreparation         int `json:"dreadlord_blocks_placed_preparation"`
				DreadlordActivationsStandard             int `json:"dreadlord_activations_standard"`
				DreadlordATotalKills                     int `json:"dreadlord_a_total_kills"`
				DreadlordTreasuresFoundStandard          int `json:"dreadlord_treasures_found_standard"`
				DreadlordWitherDamage                    int `json:"dreadlord_wither_damage"`
				DreadlordBlocksBroken                    int `json:"dreadlord_blocks_broken"`
				DreadlordWitherDamageStandard            int `json:"dreadlord_wither_damage_standard"`
				DreadlordKillsMeleeBehind                int `json:"dreadlord_kills_melee_behind"`
				DreadlordMetersWalkedSpeedStandard       int `json:"dreadlord_meters_walked_speed_standard"`
				DreadlordKillsStandard                   int `json:"dreadlord_kills_standard"`
				DreadlordMetersFallen                    int `json:"dreadlord_meters_fallen"`
				DreadlordTotalDeaths                     int `json:"dreadlord_total_deaths"`
				DreadlordAActivationsDeathmatch          int `json:"dreadlord_a_activations_deathmatch"`
				DreadlordFinalKillsMeleeBehind           int `json:"dreadlord_final_kills_melee_behind"`
				DreadlordTotalFinalKills                 int `json:"dreadlord_total_final_kills"`
				DreadlordAFinalKills                     int `json:"dreadlord_a_final_kills"`
				DreadlordFinalKillsMeleeStandard         int `json:"dreadlord_final_kills_melee_standard"`
				DreadlordFinalAssistsStandard            int `json:"dreadlord_final_assists_standard"`
				DreadlordDeathsStandard                  int `json:"dreadlord_deaths_standard"`
				DreadlordADefenderAssistsStandard        int `json:"dreadlord_a_defender_assists_standard"`
				DreadlordBlocksPlaced                    int `json:"dreadlord_blocks_placed"`
				DreadlordFinalKills                      int `json:"dreadlord_final_kills"`
				DreadlordActivations                     int `json:"dreadlord_activations"`
				DreadlordBlocksPlacedStandard            int `json:"dreadlord_blocks_placed_standard"`
				DreadlordDefenderAssists                 int `json:"dreadlord_defender_assists"`
				DreadlordDefenderKills                   int `json:"dreadlord_defender_kills"`
				DreadlordATotalKillsStandard             int `json:"dreadlord_a_total_kills_standard"`
				DreadlordAssists                         int `json:"dreadlord_assists"`
				DreadlordIronOreBroken                   int `json:"dreadlord_iron_ore_broken"`
				DreadlordLosses                          int `json:"dreadlord_losses"`
				DreadlordBlocksPlacedPreparationStandard int `json:"dreadlord_blocks_placed_preparation_standard"`
				DreadlordTotalDeathsStandard             int `json:"dreadlord_total_deaths_standard"`
				DreadlordADefenderAssists                int `json:"dreadlord_a_defender_assists"`
				DreadlordPotionsDrunk                    int `json:"dreadlord_potions_drunk"`
				DreadlordMetersWalkedSpeed               int `json:"dreadlord_meters_walked_speed"`
				DreadlordBlocksBrokenStandard            int `json:"dreadlord_blocks_broken_standard"`
				DreadlordGamesPlayed                     int `json:"dreadlord_games_played"`
				DreadlordAFinalKillsMeleeBehindStandard  int `json:"dreadlord_a_final_kills_melee_behind_standard"`
				DreadlordPotionsDrunkStandard            int `json:"dreadlord_potions_drunk_standard"`
				WitherDamage                             int `json:"wither_damage"`
				DreadlordAFinalKillsMeleeStandard        int `json:"dreadlord_a_final_kills_melee_standard"`
				DreadlordFinalKillsStandard              int `json:"dreadlord_final_kills_standard"`
				DreadlordDamageDealt                     int `json:"dreadlord_damage_dealt"`
				DreadlordAFinalKillsMeleeBehind          int `json:"dreadlord_a_final_kills_melee_behind"`
				DreadlordLossesStandard                  int `json:"dreadlord_losses_standard"`
				DreadlordAActivations                    int `json:"dreadlord_a_activations"`
				DreadlordKillsMeleeBehindStandard        int `json:"dreadlord_kills_melee_behind_standard"`
				DreadlordAAssists                        int `json:"dreadlord_a_assists"`
				DreadlordAActivationsStandard            int `json:"dreadlord_a_activations_standard"`
				DreadlordFinalKillsMeleeBehindStandard   int `json:"dreadlord_final_kills_melee_behind_standard"`
				DreadlordDefenderAssistsStandard         int `json:"dreadlord_defender_assists_standard"`
				DreadlordIronOreBrokenStandard           int `json:"dreadlord_iron_ore_broken_standard"`
				DreadlordFoodEatenStandard               int `json:"dreadlord_food_eaten_standard"`
				IronArmorGiftedStandard                  int `json:"iron_armor_gifted_standard"`
				DreadlordBreadCraftedStandard            int `json:"dreadlord_bread_crafted_standard"`
				DreadlordIronArmorGifted                 int `json:"dreadlord_iron_armor_gifted"`
				DreadlordSteaksEaten                     int `json:"dreadlord_steaks_eaten"`
				BreadCraftedStandard                     int `json:"bread_crafted_standard"`
				DreadlordBreadCrafted                    int `json:"dreadlord_bread_crafted"`
				DreadlordWitherKillsStandard             int `json:"dreadlord_wither_kills_standard"`
				DreadlordWitherKills                     int `json:"dreadlord_wither_kills"`
				DreadlordSteaksEatenStandard             int `json:"dreadlord_steaks_eaten_standard"`
				IronArmorGifted                          int `json:"iron_armor_gifted"`
				BreadCrafted                             int `json:"bread_crafted"`
				DreadlordAKillsMeleeStandard             int `json:"dreadlord_a_kills_melee_standard"`
				WitherKillsStandard                      int `json:"wither_kills_standard"`
				DreadlordAKillsMelee                     int `json:"dreadlord_a_kills_melee"`
				WitherKills                              int `json:"wither_kills"`
				DreadlordIronArmorGiftedStandard         int `json:"dreadlord_iron_armor_gifted_standard"`
				DreadlordFoodEaten                       int `json:"dreadlord_food_eaten"`
				DreadlordAKillsStandard                  int `json:"dreadlord_a_kills_standard"`
				DreadlordAKills                          int `json:"dreadlord_a_kills"`
				DreadlordArrowsHitStandard               int `json:"dreadlord_arrows_hit_standard"`
				ApplesEaten                              int `json:"apples_eaten"`
				DreadlordArrowsFired                     int `json:"dreadlord_arrows_fired"`
				ApplesEatenStandard                      int `json:"apples_eaten_standard"`
				DreadlordBreadEaten                      int `json:"dreadlord_bread_eaten"`
				DreadlordADefenderKillsStandard          int `json:"dreadlord_a_defender_kills_standard"`
				DreadlordArrowsHit                       int `json:"dreadlord_arrows_hit"`
				DreadlordArrowsFiredStandard             int `json:"dreadlord_arrows_fired_standard"`
				DreadlordADefenderKills                  int `json:"dreadlord_a_defender_kills"`
				DreadlordGoldenApplesEaten               int `json:"dreadlord_golden_apples_eaten"`
				DreadlordApplesEatenStandard             int `json:"dreadlord_apples_eaten_standard"`
				DreadlordBreadEatenStandard              int `json:"dreadlord_bread_eaten_standard"`
				DreadlordApplesEaten                     int `json:"dreadlord_apples_eaten"`
				DreadlordGoldenApplesEatenStandard       int `json:"dreadlord_golden_apples_eaten_standard"`
				MolemanWitherDamageStandard              int `json:"moleman_wither_damage_standard"`
				MolemanDefenderKills                     int `json:"moleman_defender_kills"`
				MolemanAAssists                          int `json:"moleman_a_assists"`
				MolemanAActivationsDeathmatchStandard    int `json:"moleman_a_activations_deathmatch_standard"`
				MolemanMetersWalked                      int `json:"moleman_meters_walked"`
				MolemanBlocksPlacedPreparationStandard   int `json:"moleman_blocks_placed_preparation_standard"`
				MolemanAKillsMeleeStandard               int `json:"moleman_a_kills_melee_standard"`
				MolemanATotalFinalKillsStandard          int `json:"moleman_a_total_final_kills_standard"`
				MolemanActivationsDeathmatchStandard     int `json:"moleman_activations_deathmatch_standard"`
				MolemanBlocksPlacedStandard              int `json:"moleman_blocks_placed_standard"`
				MolemanFinalAssists                      int `json:"moleman_final_assists"`
				MolemanAssists                           int `json:"moleman_assists"`
				MolemanTotalDeathsStandard               int `json:"moleman_total_deaths_standard"`
				MolemanAActivations                      int `json:"moleman_a_activations"`
				MolemanTotalFinalKillsStandard           int `json:"moleman_total_final_kills_standard"`
				MolemanBActivations                      int `json:"moleman_b_activations"`
				MolemanFinalKills                        int `json:"moleman_final_kills"`
				MolemanAKillsStandard                    int `json:"moleman_a_kills_standard"`
				MolemanFoodEaten                         int `json:"moleman_food_eaten"`
				MolemanPotionsDrunkStandard              int `json:"moleman_potions_drunk_standard"`
				MolemanATotalFinalKills                  int `json:"moleman_a_total_final_kills"`
				MolemanFinalAssistsStandard              int `json:"moleman_final_assists_standard"`
				MolemanAKillsMelee                       int `json:"moleman_a_kills_melee"`
				MolemanBActivationsStandard              int `json:"moleman_b_activations_standard"`
				DefenderFinalAssists                     int `json:"defender_final_assists"`
				MolemanBlocksBroken                      int `json:"moleman_blocks_broken"`
				MolemanFinalKillsMelee                   int `json:"moleman_final_kills_melee"`
				MolemanGamesPlayed                       int `json:"moleman_games_played"`
				MolemanWitherDamage                      int `json:"moleman_wither_damage"`
				MolemanMetersFallen                      int `json:"moleman_meters_fallen"`
				MolemanBreadEaten                        int `json:"moleman_bread_eaten"`
				MolemanFinalKillsMeleeBehindStandard     int `json:"moleman_final_kills_melee_behind_standard"`
				MolemanATotalKillsStandard               int `json:"moleman_a_total_kills_standard"`
				MolemanIronOreBroken                     int `json:"moleman_iron_ore_broken"`
				MolemanFinalKillsMeleeBehind             int `json:"moleman_final_kills_melee_behind"`
				MolemanAFinalAssists                     int `json:"moleman_a_final_assists"`
				MolemanGamesPlayedStandard               int `json:"moleman_games_played_standard"`
				MolemanDefenderKillsStandard             int `json:"moleman_defender_kills_standard"`
				MolemanArrowsHit                         int `json:"moleman_arrows_hit"`
				MolemanDeaths                            int `json:"moleman_deaths"`
				MolemanAActivationsDeathmatch            int `json:"moleman_a_activations_deathmatch"`
				MolemanArrowsFired                       int `json:"moleman_arrows_fired"`
				MolemanLosses                            int `json:"moleman_losses"`
				MolemanBreadCrafted                      int `json:"moleman_bread_crafted"`
				MolemanActivationsDeathmatch             int `json:"moleman_activations_deathmatch"`
				MolemanBlocksPlaced                      int `json:"moleman_blocks_placed"`
				MolemanADamageDealtStandard              int `json:"moleman_a_damage_dealt_standard"`
				MolemanTotalKills                        int `json:"moleman_total_kills"`
				MolemanFoodEatenStandard                 int `json:"moleman_food_eaten_standard"`
				MolemanSteaksEatenStandard               int `json:"moleman_steaks_eaten_standard"`
				MolemanAssistsStandard                   int `json:"moleman_assists_standard"`
				MolemanTreasuresFoundStandard            int `json:"moleman_treasures_found_standard"`
				MolemanMetersWalkedStandard              int `json:"moleman_meters_walked_standard"`
				MolemanTotalFinalKills                   int `json:"moleman_total_final_kills"`
				MolemanTotalDeaths                       int `json:"moleman_total_deaths"`
				MolemanFinalDeathsStandard               int `json:"moleman_final_deaths_standard"`
				MolemanBreadCraftedStandard              int `json:"moleman_bread_crafted_standard"`
				MolemanFinalKillsMeleeStandard           int `json:"moleman_final_kills_melee_standard"`
				MolemanDefenderFinalAssists              int `json:"moleman_defender_final_assists"`
				MolemanDefenderFinalAssistsStandard      int `json:"moleman_defender_final_assists_standard"`
				MolemanABlocksBroken                     int `json:"moleman_a_blocks_broken"`
				MolemanDamageDealtStandard               int `json:"moleman_damage_dealt_standard"`
				MolemanADamageDealt                      int `json:"moleman_a_damage_dealt"`
				MolemanTimePlayedStandard                int `json:"moleman_time_played_standard"`
				MolemanActivationsStandard               int `json:"moleman_activations_standard"`
				MolemanTotalKillsStandard                int `json:"moleman_total_kills_standard"`
				MolemanFinalKillsStandard                int `json:"moleman_final_kills_standard"`
				MolemanSteaksEaten                       int `json:"moleman_steaks_eaten"`
				MolemanADefenderFinalAssists             int `json:"moleman_a_defender_final_assists"`
				MolemanADefenderFinalAssistsStandard     int `json:"moleman_a_defender_final_assists_standard"`
				MolemanMetersWalkedSpeed                 int `json:"moleman_meters_walked_speed"`
				MolemanMetersWalkedSpeedStandard         int `json:"moleman_meters_walked_speed_standard"`
				MolemanMetersFallenStandard              int `json:"moleman_meters_fallen_standard"`
				MolemanFinalDeaths                       int `json:"moleman_final_deaths"`
				MolemanBlocksPlacedPreparation           int `json:"moleman_blocks_placed_preparation"`
				MolemanPotionsDrunk                      int `json:"moleman_potions_drunk"`
				MolemanBlocksBrokenStandard              int `json:"moleman_blocks_broken_standard"`
				MolemanTimePlayed                        int `json:"moleman_time_played"`
				DefenderFinalAssistsStandard             int `json:"defender_final_assists_standard"`
				MolemanAKills                            int `json:"moleman_a_kills"`
				MolemanTreasuresFound                    int `json:"moleman_treasures_found"`
				MolemanDamageDealt                       int `json:"moleman_damage_dealt"`
				MolemanKills                             int `json:"moleman_kills"`
				MolemanKillsMeleeStandard                int `json:"moleman_kills_melee_standard"`
				MolemanATotalKills                       int `json:"moleman_a_total_kills"`
				MolemanAActivationsStandard              int `json:"moleman_a_activations_standard"`
				MolemanArrowsFiredStandard               int `json:"moleman_arrows_fired_standard"`
				MolemanAFinalAssistsStandard             int `json:"moleman_a_final_assists_standard"`
				MolemanAAssistsStandard                  int `json:"moleman_a_assists_standard"`
				MolemanLossesStandard                    int `json:"moleman_losses_standard"`
				MolemanBreadEatenStandard                int `json:"moleman_bread_eaten_standard"`
				MolemanABlocksBrokenStandard             int `json:"moleman_a_blocks_broken_standard"`
				MolemanKillsMelee                        int `json:"moleman_kills_melee"`
				MolemanKillsStandard                     int `json:"moleman_kills_standard"`
				MolemanIronOreBrokenStandard             int `json:"moleman_iron_ore_broken_standard"`
				MolemanArrowsHitStandard                 int `json:"moleman_arrows_hit_standard"`
				MolemanActivations                       int `json:"moleman_activations"`
				MolemanDeathsStandard                    int `json:"moleman_deaths_standard"`
				MolemanBTotalFinalKills                  int `json:"moleman_b_total_final_kills"`
				MolemanAbsorptionPotionsDrunkStandard    int `json:"moleman_absorption_potions_drunk_standard"`
				MolemanBFinalAssists                     int `json:"moleman_b_final_assists"`
				AbsorptionPotionsDrunkStandard           int `json:"absorption_potions_drunk_standard"`
				MolemanAbsorptionPotionsDrunk            int `json:"moleman_absorption_potions_drunk"`
				AbsorptionPotionsDrunk                   int `json:"absorption_potions_drunk"`
				MolemanBTotalFinalKillsStandard          int `json:"moleman_b_total_final_kills_standard"`
				MolemanBFinalAssistsStandard             int `json:"moleman_b_final_assists_standard"`
				SpiderGamesPlayedStandard                int `json:"spider_games_played_standard"`
				SpiderLossesStandard                     int `json:"spider_losses_standard"`
				SpiderGamesPlayed                        int `json:"spider_games_played"`
				SpiderTotalDeathsStandard                int `json:"spider_total_deaths_standard"`
				SpiderTotalDeaths                        int `json:"spider_total_deaths"`
				SpiderDeathsStandard                     int `json:"spider_deaths_standard"`
				SpiderTimePlayed                         int `json:"spider_time_played"`
				SpiderTimePlayedStandard                 int `json:"spider_time_played_standard"`
				SpiderMetersWalkedStandard               int `json:"spider_meters_walked_standard"`
				SpiderFinalDeathsStandard                int `json:"spider_final_deaths_standard"`
				SpiderMetersWalked                       int `json:"spider_meters_walked"`
				SpiderFinalDeaths                        int `json:"spider_final_deaths"`
				SpiderMetersFallenStandard               int `json:"spider_meters_fallen_standard"`
				SpiderBlocksPlaced                       int `json:"spider_blocks_placed"`
				SpiderBlocksPlacedPreparationStandard    int `json:"spider_blocks_placed_preparation_standard"`
				SpiderBlocksPlacedStandard               int `json:"spider_blocks_placed_standard"`
				SpiderIronOreBroken                      int `json:"spider_iron_ore_broken"`
				SpiderMetersFallen                       int `json:"spider_meters_fallen"`
				SpiderBlocksBrokenStandard               int `json:"spider_blocks_broken_standard"`
				SpiderBlocksPlacedPreparation            int `json:"spider_blocks_placed_preparation"`
				SpiderTreasuresFound                     int `json:"spider_treasures_found"`
				SpiderIronOreBrokenStandard              int `json:"spider_iron_ore_broken_standard"`
				SpiderTreasuresFoundStandard             int `json:"spider_treasures_found_standard"`
				SpiderBlocksBroken                       int `json:"spider_blocks_broken"`
				SpiderAActivations                       int `json:"spider_a_activations"`
				SpiderActivationsStandard                int `json:"spider_activations_standard"`
				SpiderActivations                        int `json:"spider_activations"`
				SpiderAActivationsStandard               int `json:"spider_a_activations_standard"`
				PigmanArrowsHitFaceOff                   int `json:"pigman_arrows_hit_face_off"`
				PigmanDamageDealtFaceOff                 int `json:"pigman_damage_dealt_face_off"`
				PigmanFoodEatenFaceOff                   int `json:"pigman_food_eaten_face_off"`
				PigmanFinalAssistsFaceOff                int `json:"pigman_final_assists_face_off"`
				PigmanPotionsDrunkFaceOff                int `json:"pigman_potions_drunk_face_off"`
				PigmanIronOreBrokenFaceOff               int `json:"pigman_iron_ore_broken_face_off"`
				PigmanTreasuresFoundFaceOff              int `json:"pigman_treasures_found_face_off"`
				PigmanAAssists                           int `json:"pigman_a_assists"`
				PigmanActivations                        int `json:"pigman_activations"`
				PigmanBActivationsFaceOff                int `json:"pigman_b_activations_face_off"`
				PigmanBlocksPlacedPreparationFaceOff     int `json:"pigman_blocks_placed_preparation_face_off"`
				PigmanAKillsMeleeFaceOff                 int `json:"pigman_a_kills_melee_face_off"`
				KillsFaceOff                             int `json:"kills_face_off"`
				PigmanActivationsDeathmatch              int `json:"pigman_activations_deathmatch"`
				ArrowsHitFaceOff                         int `json:"arrows_hit_face_off"`
				PigmanGamesPlayed                        int `json:"pigman_games_played"`
				PigmanAssistsFaceOff                     int `json:"pigman_assists_face_off"`
				PigmanATotalKillsFaceOff                 int `json:"pigman_a_total_kills_face_off"`
				FoodEatenFaceOff                         int `json:"food_eaten_face_off"`
				PigmanAActivationsFaceOff                int `json:"pigman_a_activations_face_off"`
				PigmanATotalKills                        int `json:"pigman_a_total_kills"`
				BlocksBrokenFaceOff                      int `json:"blocks_broken_face_off"`
				PigmanIronArmorGiftedFaceOff             int `json:"pigman_iron_armor_gifted_face_off"`
				SteaksEatenFaceOff                       int `json:"steaks_eaten_face_off"`
				PigmanBlocksBrokenFaceOff                int `json:"pigman_blocks_broken_face_off"`
				PigmanDefenderKills                      int `json:"pigman_defender_kills"`
				PigmanADefenderAssistsFaceOff            int `json:"pigman_a_defender_assists_face_off"`
				PigmanKillsMeleeFaceOff                  int `json:"pigman_kills_melee_face_off"`
				PigmanKillsMeleeBehindFaceOff            int `json:"pigman_kills_melee_behind_face_off"`
				PigmanGamesPlayedFaceOff                 int `json:"pigman_games_played_face_off"`
				PigmanMetersWalkedSpeedFaceOff           int `json:"pigman_meters_walked_speed_face_off"`
				PigmanIronOreBroken                      int `json:"pigman_iron_ore_broken"`
				PigmanDefenderAssists                    int `json:"pigman_defender_assists"`
				PigmanADefenderKillsFaceOff              int `json:"pigman_a_defender_kills_face_off"`
				PigmanTotalKillsFaceOff                  int `json:"pigman_total_kills_face_off"`
				PigmanADamageDealt                       int `json:"pigman_a_damage_dealt"`
				PigmanWinsFaceOff                        int `json:"pigman_wins_face_off"`
				KillsMeleeFaceOff                        int `json:"kills_melee_face_off"`
				PigmanTimePlayed                         int `json:"pigman_time_played"`
				PigmanBActivationsDeathmatchFaceOff      int `json:"pigman_b_activations_deathmatch_face_off"`
				PigmanActivationsFaceOff                 int `json:"pigman_activations_face_off"`
				PigmanTotalDeaths                        int `json:"pigman_total_deaths"`
				TotalFinalKillsFaceOff                   int `json:"total_final_kills_face_off"`
				PigmanADefenderKills                     int `json:"pigman_a_defender_kills"`
				KillsMeleeBehindFaceOff                  int `json:"kills_melee_behind_face_off"`
				MetersWalkedSpeedFaceOff                 int `json:"meters_walked_speed_face_off"`
				PigmanDeathsFaceOff                      int `json:"pigman_deaths_face_off"`
				PigmanTreasuresFound                     int `json:"pigman_treasures_found"`
				PigmanPotionsDrunk                       int `json:"pigman_potions_drunk"`
				WoodChoppedFaceOff                       int `json:"wood_chopped_face_off"`
				PigmanBreadCraftedFaceOff                int `json:"pigman_bread_crafted_face_off"`
				PigmanSteaksEatenFaceOff                 int `json:"pigman_steaks_eaten_face_off"`
				PigmanWoodChopped                        int `json:"pigman_wood_chopped"`
				TreasuresFoundFaceOff                    int `json:"treasures_found_face_off"`
				PigmanAActivationsDeathmatchFaceOff      int `json:"pigman_a_activations_deathmatch_face_off"`
				ActivationsDeathmatchFaceOff             int `json:"activations_deathmatch_face_off"`
				PigmanAKillsFaceOff                      int `json:"pigman_a_kills_face_off"`
				ActivationsFaceOff                       int `json:"activations_face_off"`
				DefenderKillsFaceOff                     int `json:"defender_kills_face_off"`
				PigmanKillsMelee                         int `json:"pigman_kills_melee"`
				PigmanATotalFinalKills                   int `json:"pigman_a_total_final_kills"`
				AssistsFaceOff                           int `json:"assists_face_off"`
				PigmanDefenderAssistsFaceOff             int `json:"pigman_defender_assists_face_off"`
				DamageDealtFaceOff                       int `json:"damage_dealt_face_off"`
				PigmanFinalKillsMeleeBehind              int `json:"pigman_final_kills_melee_behind"`
				DefenderAssistsFaceOff                   int `json:"defender_assists_face_off"`
				PigmanFinalKillsMelee                    int `json:"pigman_final_kills_melee"`
				PigmanGActivationsFaceOff                int `json:"pigman_g_activations_face_off"`
				PigmanBActivationsDeathmatch             int `json:"pigman_b_activations_deathmatch"`
				PigmanMetersWalkedFaceOff                int `json:"pigman_meters_walked_face_off"`
				PigmanAKills                             int `json:"pigman_a_kills"`
				TimePlayedFaceOff                        int `json:"time_played_face_off"`
				PigmanFinalKillsFaceOff                  int `json:"pigman_final_kills_face_off"`
				PigmanTotalKills                         int `json:"pigman_total_kills"`
				PigmanADefenderAssists                   int `json:"pigman_a_defender_assists"`
				PigmanBlocksPlaced                       int `json:"pigman_blocks_placed"`
				PigmanATotalFinalKillsFaceOff            int `json:"pigman_a_total_final_kills_face_off"`
				PigmanArrowsFired                        int `json:"pigman_arrows_fired"`
				PigmanAKillsMelee                        int `json:"pigman_a_kills_melee"`
				PigmanAFinalAssistsFaceOff               int `json:"pigman_a_final_assists_face_off"`
				PigmanTimePlayedFaceOff                  int `json:"pigman_time_played_face_off"`
				TotalKillsFaceOff                        int `json:"total_kills_face_off"`
				PigmanKillsFaceOff                       int `json:"pigman_kills_face_off"`
				PigmanBActivations                       int `json:"pigman_b_activations"`
				PigmanAActivations                       int `json:"pigman_a_activations"`
				PigmanMetersWalked                       int `json:"pigman_meters_walked"`
				PigmanFinalAssists                       int `json:"pigman_final_assists"`
				PigmanAActivationsDeathmatch             int `json:"pigman_a_activations_deathmatch"`
				PigmanArrowsHit                          int `json:"pigman_arrows_hit"`
				PigmanDamageDealt                        int `json:"pigman_damage_dealt"`
				PigmanActivationsDeathmatchFaceOff       int `json:"pigman_activations_deathmatch_face_off"`
				PigmanFinalKillsMeleeBehindFaceOff       int `json:"pigman_final_kills_melee_behind_face_off"`
				PigmanKillsMeleeBehind                   int `json:"pigman_kills_melee_behind"`
				IronArmorGiftedFaceOff                   int `json:"iron_armor_gifted_face_off"`
				PigmanTotalFinalKillsFaceOff             int `json:"pigman_total_final_kills_face_off"`
				PigmanWoodChoppedFaceOff                 int `json:"pigman_wood_chopped_face_off"`
				PigmanGActivations                       int `json:"pigman_g_activations"`
				PigmanBlocksPlacedPreparation            int `json:"pigman_blocks_placed_preparation"`
				ArrowsFiredFaceOff                       int `json:"arrows_fired_face_off"`
				PigmanIronArmorGifted                    int `json:"pigman_iron_armor_gifted"`
				PigmanBlocksPlacedFaceOff                int `json:"pigman_blocks_placed_face_off"`
				PigmanAAssistsFaceOff                    int `json:"pigman_a_assists_face_off"`
				PigmanTotalDeathsFaceOff                 int `json:"pigman_total_deaths_face_off"`
				PigmanDefenderKillsFaceOff               int `json:"pigman_defender_kills_face_off"`
				PigmanFoodEaten                          int `json:"pigman_food_eaten"`
				PigmanAFinalAssists                      int `json:"pigman_a_final_assists"`
				FinalKillsMeleeFaceOff                   int `json:"final_kills_melee_face_off"`
				PigmanMetersFallenFaceOff                int `json:"pigman_meters_fallen_face_off"`
				PigmanADamageDealtFaceOff                int `json:"pigman_a_damage_dealt_face_off"`
				PigmanBreadEatenFaceOff                  int `json:"pigman_bread_eaten_face_off"`
				PigmanBreadCrafted                       int `json:"pigman_bread_crafted"`
				BreadEatenFaceOff                        int `json:"bread_eaten_face_off"`
				PigmanArrowsFiredFaceOff                 int `json:"pigman_arrows_fired_face_off"`
				PigmanFinalKillsMeleeFaceOff             int `json:"pigman_final_kills_melee_face_off"`
				IronOreBrokenFaceOff                     int `json:"iron_ore_broken_face_off"`
				PigmanAssists                            int `json:"pigman_assists"`
				BreadCraftedFaceOff                      int `json:"bread_crafted_face_off"`
				PigmanMetersFallen                       int `json:"pigman_meters_fallen"`
				FinalKillsMeleeBehindFaceOff             int `json:"final_kills_melee_behind_face_off"`
				PigmanBlocksBroken                       int `json:"pigman_blocks_broken"`
				PigmanMetersWalkedSpeed                  int `json:"pigman_meters_walked_speed"`
				TotalDeathsFaceOff                       int `json:"total_deaths_face_off"`
				FinalKillsFaceOff                        int `json:"final_kills_face_off"`
				PotionsDrunkFaceOff                      int `json:"potions_drunk_face_off"`
				FinalAssistsFaceOff                      int `json:"final_assists_face_off"`
				PigmanBreadEaten                         int `json:"pigman_bread_eaten"`
				DeathsFaceOff                            int `json:"deaths_face_off"`
				PigmanSteaksEaten                        int `json:"pigman_steaks_eaten"`
				SkeletonKillsMeleeBehind                 int `json:"skeleton_kills_melee_behind"`
				SkeletonGamesPlayedFaceOff               int `json:"skeleton_games_played_face_off"`
				SkeletonAssists                          int `json:"skeleton_assists"`
				SkeletonTreasuresFound                   int `json:"skeleton_treasures_found"`
				SkeletonDefenderAssists                  int `json:"skeleton_defender_assists"`
				SkeletonBlocksPlacedPreparation          int `json:"skeleton_blocks_placed_preparation"`
				SkeletonBlocksPlacedFaceOff              int `json:"skeleton_blocks_placed_face_off"`
				SkeletonMetersWalkedSpeed                int `json:"skeleton_meters_walked_speed"`
				SkeletonTotalKills                       int `json:"skeleton_total_kills"`
				SkeletonLossesFaceOff                    int `json:"skeleton_losses_face_off"`
				SkeletonDefenderKills                    int `json:"skeleton_defender_kills"`
				SkeletonBlocksPlaced                     int `json:"skeleton_blocks_placed"`
				SkeletonBlocksBrokenFaceOff              int `json:"skeleton_blocks_broken_face_off"`
				SkeletonActivationsFaceOff               int `json:"skeleton_activations_face_off"`
				SkeletonIronOreBrokenFaceOff             int `json:"skeleton_iron_ore_broken_face_off"`
				SkeletonTreasuresFoundFaceOff            int `json:"skeleton_treasures_found_face_off"`
				SkeletonKillsFaceOff                     int `json:"skeleton_kills_face_off"`
				SkeletonBlocksPlacedPreparationFaceOff   int `json:"skeleton_blocks_placed_preparation_face_off"`
				SkeletonPotionsDrunk                     int `json:"skeleton_potions_drunk"`
				SkeletonMetersWalkedSpeedFaceOff         int `json:"skeleton_meters_walked_speed_face_off"`
				SkeletonMetersWalked                     int `json:"skeleton_meters_walked"`
				SkeletonTimePlayed                       int `json:"skeleton_time_played"`
				SkeletonMetersFallen                     int `json:"skeleton_meters_fallen"`
				SkeletonKillsMeleeFaceOff                int `json:"skeleton_kills_melee_face_off"`
				SkeletonMetersWalkedFaceOff              int `json:"skeleton_meters_walked_face_off"`
				SkeletonAActivations                     int `json:"skeleton_a_activations"`
				SkeletonPotionsDrunkFaceOff              int `json:"skeleton_potions_drunk_face_off"`
				SkeletonArrowsFired                      int `json:"skeleton_arrows_fired"`
				SkeletonDefenderAssistsFaceOff           int `json:"skeleton_defender_assists_face_off"`
				SkeletonAActivationsFaceOff              int `json:"skeleton_a_activations_face_off"`
				SkeletonArrowsHitFaceOff                 int `json:"skeleton_arrows_hit_face_off"`
				SkeletonBlocksBroken                     int `json:"skeleton_blocks_broken"`
				SkeletonAssistsFaceOff                   int `json:"skeleton_assists_face_off"`
				SkeletonArrowsHit                        int `json:"skeleton_arrows_hit"`
				SkeletonIronOreBroken                    int `json:"skeleton_iron_ore_broken"`
				SkeletonActivations                      int `json:"skeleton_activations"`
				SkeletonKillsMelee                       int `json:"skeleton_kills_melee"`
				SkeletonGamesPlayed                      int `json:"skeleton_games_played"`
				SkeletonKillsMeleeBehindFaceOff          int `json:"skeleton_kills_melee_behind_face_off"`
				SkeletonDefenderKillsFaceOff             int `json:"skeleton_defender_kills_face_off"`
				SkeletonMetersFallenFaceOff              int `json:"skeleton_meters_fallen_face_off"`
				SkeletonTotalKillsFaceOff                int `json:"skeleton_total_kills_face_off"`
				SkeletonTimePlayedFaceOff                int `json:"skeleton_time_played_face_off"`
				SkeletonArrowsFiredFaceOff               int `json:"skeleton_arrows_fired_face_off"`
			} `json:"Walls3"`
			GingerBread struct {
				EngineActive                                  string   `json:"engine_active"`
				FrameActive                                   string   `json:"frame_active"`
				BoosterActive                                 string   `json:"booster_active"`
				HelmetActive                                  string   `json:"helmet_active"`
				JacketActive                                  string   `json:"jacket_active"`
				PantsActive                                   string   `json:"pants_active"`
				SkinActive                                    string   `json:"skin_active"`
				ShoesActive                                   string   `json:"shoes_active"`
				Packages                                      []string `json:"packages"`
				Horn                                          string   `json:"horn"`
				Coins                                         int      `json:"coins"`
				GoldTrophy                                    int      `json:"gold_trophy"`
				LapsCompleted                                 int      `json:"laps_completed"`
				CoinsPickedUp                                 int      `json:"coins_picked_up"`
				BlueTorpedoHit                                int      `json:"blue_torpedo_hit"`
				BoxPickups                                    int      `json:"box_pickups"`
				BananaHitsSent                                int      `json:"banana_hits_sent"`
				Wins                                          int      `json:"wins"`
				Parts                                         string   `json:"parts"`
				BananaHitsReceived                            int      `json:"banana_hits_received"`
				SilverTrophy                                  int      `json:"silver_trophy"`
				BronzeTrophy                                  int      `json:"bronze_trophy"`
				ParticleTrail                                 string   `json:"particle_trail"`
				GoldTrophyOlympus                             int      `json:"gold_trophy_olympus"`
				GoldTrophyJunglerush                          int      `json:"gold_trophy_junglerush"`
				GoldTrophyHypixelgp                           int      `json:"gold_trophy_hypixelgp"`
				GoldTrophyRetro                               int      `json:"gold_trophy_retro"`
				GoldTrophyCanyon                              int      `json:"gold_trophy_canyon"`
				SilverTrophyWeeklyB                           int      `json:"silver_trophy_weekly_b"`
				SilverTrophyMonthlyB                          int      `json:"silver_trophy_monthly_b"`
				SilverTrophyOlympus                           int      `json:"silver_trophy_olympus"`
				SilverTrophyJunglerush                        int      `json:"silver_trophy_junglerush"`
				SilverTrophyRetro                             int      `json:"silver_trophy_retro"`
				SilverTrophyCanyon                            int      `json:"silver_trophy_canyon"`
				BronzeTrophyWeeklyB                           int      `json:"bronze_trophy_weekly_b"`
				BronzeTrophyMonthlyB                          int      `json:"bronze_trophy_monthly_b"`
				BronzeTrophyOlympus                           int      `json:"bronze_trophy_olympus"`
				BronzeTrophyJunglerush                        int      `json:"bronze_trophy_junglerush"`
				BronzeTrophyHypixelgp                         int      `json:"bronze_trophy_hypixelgp"`
				BronzeTrophyRetro                             int      `json:"bronze_trophy_retro"`
				BronzeTrophyCanyon                            int      `json:"bronze_trophy_canyon"`
				BoxPickupsWeeklyB                             int      `json:"box_pickups_weekly_b"`
				BoxPickupsMonthlyB                            int      `json:"box_pickups_monthly_b"`
				BoxPickupsOlympus                             int      `json:"box_pickups_olympus"`
				BoxPickupsJunglerush                          int      `json:"box_pickups_junglerush"`
				BoxPickupsHypixelgp                           int      `json:"box_pickups_hypixelgp"`
				BoxPickupsRetro                               int      `json:"box_pickups_retro"`
				BoxPickupsCanyon                              int      `json:"box_pickups_canyon"`
				OlympusPlays                                  int      `json:"olympus_plays"`
				JunglerushPlays                               int      `json:"junglerush_plays"`
				HypixelgpPlays                                int      `json:"hypixelgp_plays"`
				RetroPlays                                    int      `json:"retro_plays"`
				CanyonPlays                                   int      `json:"canyon_plays"`
				BoxPickupsWeeklyA                             int      `json:"box_pickups_weekly_a"`
				BoxPickupsMonthlyA                            int      `json:"box_pickups_monthly_a"`
				GoldTrophyMonthlyB                            int      `json:"gold_trophy_monthly_b"`
				GoldTrophyWeeklyB                             int      `json:"gold_trophy_weekly_b"`
				GoldTrophyMonthlyA                            int      `json:"gold_trophy_monthly_a"`
				SilverTrophyWeeklyA                           int      `json:"silver_trophy_weekly_a"`
				SilverTrophyMonthlyA                          int      `json:"silver_trophy_monthly_a"`
				GoldTrophyWeeklyA                             int      `json:"gold_trophy_weekly_a"`
				BronzeTrophyMonthlyA                          int      `json:"bronze_trophy_monthly_a"`
				BronzeTrophyWeeklyA                           int      `json:"bronze_trophy_weekly_a"`
				LastTourneyAd                                 int64    `json:"lastTourneyAd"`
				TourneyGingerbreadSolo0BananaHitsSent         int      `json:"tourney_gingerbread_solo_0_banana_hits_sent"`
				TourneyGingerbreadSolo0BlueTorpedoHit         int      `json:"tourney_gingerbread_solo_0_blue_torpedo_hit"`
				TourneyGingerbreadSolo0BoxPickups             int      `json:"tourney_gingerbread_solo_0_box_pickups"`
				TourneyGingerbreadSolo0BoxPickupsJunglerush   int      `json:"tourney_gingerbread_solo_0_box_pickups_junglerush"`
				TourneyGingerbreadSolo0BoxPickupsMonthlyA     int      `json:"tourney_gingerbread_solo_0_box_pickups_monthly_a"`
				TourneyGingerbreadSolo0BoxPickupsWeeklyA      int      `json:"tourney_gingerbread_solo_0_box_pickups_weekly_a"`
				TourneyGingerbreadSolo0CoinsPickedUp          int      `json:"tourney_gingerbread_solo_0_coins_picked_up"`
				TourneyGingerbreadSolo0JunglerushPlays        int      `json:"tourney_gingerbread_solo_0_junglerush_plays"`
				TourneyGingerbreadSolo0LapsCompleted          int      `json:"tourney_gingerbread_solo_0_laps_completed"`
				TourneyGingerbreadSolo0SilverTrophy           int      `json:"tourney_gingerbread_solo_0_silver_trophy"`
				TourneyGingerbreadSolo0SilverTrophyJunglerush int      `json:"tourney_gingerbread_solo_0_silver_trophy_junglerush"`
				TourneyGingerbreadSolo0SilverTrophyMonthlyA   int      `json:"tourney_gingerbread_solo_0_silver_trophy_monthly_a"`
				TourneyGingerbreadSolo0SilverTrophyWeeklyA    int      `json:"tourney_gingerbread_solo_0_silver_trophy_weekly_a"`
				TourneyGingerbreadSolo0Wins                   int      `json:"tourney_gingerbread_solo_0_wins"`
				TourneyGingerbreadSolo0BoxPickupsOlympus      int      `json:"tourney_gingerbread_solo_0_box_pickups_olympus"`
				TourneyGingerbreadSolo0GoldTrophy             int      `json:"tourney_gingerbread_solo_0_gold_trophy"`
				TourneyGingerbreadSolo0GoldTrophyMonthlyA     int      `json:"tourney_gingerbread_solo_0_gold_trophy_monthly_a"`
				TourneyGingerbreadSolo0GoldTrophyOlympus      int      `json:"tourney_gingerbread_solo_0_gold_trophy_olympus"`
				TourneyGingerbreadSolo0GoldTrophyWeeklyA      int      `json:"tourney_gingerbread_solo_0_gold_trophy_weekly_a"`
				TourneyGingerbreadSolo0OlympusPlays           int      `json:"tourney_gingerbread_solo_0_olympus_plays"`
				TourneyGingerbreadSolo0BoxPickupsWeeklyB      int      `json:"tourney_gingerbread_solo_0_box_pickups_weekly_b"`
				TourneyGingerbreadSolo0GoldTrophyWeeklyB      int      `json:"tourney_gingerbread_solo_0_gold_trophy_weekly_b"`
				TourneyGingerbreadSolo0BananaHitsReceived     int      `json:"tourney_gingerbread_solo_0_banana_hits_received"`
				TourneyGingerbreadSolo0BoxPickupsRetro        int      `json:"tourney_gingerbread_solo_0_box_pickups_retro"`
				TourneyGingerbreadSolo0RetroPlays             int      `json:"tourney_gingerbread_solo_0_retro_plays"`
				TourneyGingerbreadSolo0SilverTrophyOlympus    int      `json:"tourney_gingerbread_solo_0_silver_trophy_olympus"`
				TourneyGingerbreadSolo0SilverTrophyWeeklyB    int      `json:"tourney_gingerbread_solo_0_silver_trophy_weekly_b"`
				TourneyGingerbreadSolo0BoxPickupsHypixelgp    int      `json:"tourney_gingerbread_solo_0_box_pickups_hypixelgp"`
				TourneyGingerbreadSolo0HypixelgpPlays         int      `json:"tourney_gingerbread_solo_0_hypixelgp_plays"`
				TourneyGingerbreadSolo0SilverTrophyHypixelgp  int      `json:"tourney_gingerbread_solo_0_silver_trophy_hypixelgp"`
				TourneyGingerbreadSolo0BoxPickupsCanyon       int      `json:"tourney_gingerbread_solo_0_box_pickups_canyon"`
				TourneyGingerbreadSolo0BronzeTrophy           int      `json:"tourney_gingerbread_solo_0_bronze_trophy"`
				TourneyGingerbreadSolo0BronzeTrophyCanyon     int      `json:"tourney_gingerbread_solo_0_bronze_trophy_canyon"`
				TourneyGingerbreadSolo0BronzeTrophyMonthlyA   int      `json:"tourney_gingerbread_solo_0_bronze_trophy_monthly_a"`
				TourneyGingerbreadSolo0BronzeTrophyWeeklyB    int      `json:"tourney_gingerbread_solo_0_bronze_trophy_weekly_b"`
				TourneyGingerbreadSolo0CanyonPlays            int      `json:"tourney_gingerbread_solo_0_canyon_plays"`
				TourneyGingerbreadSolo0SilverTrophyRetro      int      `json:"tourney_gingerbread_solo_0_silver_trophy_retro"`
				TourneyGingerbreadSolo0SilverTrophyCanyon     int      `json:"tourney_gingerbread_solo_0_silver_trophy_canyon"`
				TourneyGingerbreadSolo0BronzeTrophyJunglerush int      `json:"tourney_gingerbread_solo_0_bronze_trophy_junglerush"`
				TourneyGingerbreadSolo0GoldTrophyJunglerush   int      `json:"tourney_gingerbread_solo_0_gold_trophy_junglerush"`
				TourneyGingerbreadSolo0BronzeTrophyHypixelgp  int      `json:"tourney_gingerbread_solo_0_bronze_trophy_hypixelgp"`
				TourneyGingerbreadSolo0GoldTrophyCanyon       int      `json:"tourney_gingerbread_solo_0_gold_trophy_canyon"`
				TourneyGingerbreadSolo0BronzeTrophyOlympus    int      `json:"tourney_gingerbread_solo_0_bronze_trophy_olympus"`
				TourneyGingerbreadSolo0GoldTrophyRetro        int      `json:"tourney_gingerbread_solo_0_gold_trophy_retro"`
				TourneyGingerbreadSolo0GoldTrophyHypixelgp    int      `json:"tourney_gingerbread_solo_0_gold_trophy_hypixelgp"`
			} `json:"GingerBread"`
			SkyWars struct {
				WinStreak                                     int      `json:"win_streak"`
				SurvivedPlayers                               int      `json:"survived_players"`
				Losses                                        int      `json:"losses"`
				LossesSolo                                    int      `json:"losses_solo"`
				SurvivedPlayersKitBasicSoloDefault            int      `json:"survived_players_kit_basic_solo_default"`
				BlocksBroken                                  int      `json:"blocks_broken"`
				DeathsKitBasicSoloDefault                     int      `json:"deaths_kit_basic_solo_default"`
				LossesSoloNormal                              int      `json:"losses_solo_normal"`
				DeathsSolo                                    int      `json:"deaths_solo"`
				Deaths                                        int      `json:"deaths"`
				Quits                                         int      `json:"quits"`
				DeathsSoloNormal                              int      `json:"deaths_solo_normal"`
				SurvivedPlayersSolo                           int      `json:"survived_players_solo"`
				LossesKitBasicSoloDefault                     int      `json:"losses_kit_basic_solo_default"`
				Coins                                         int      `json:"coins"`
				Souls                                         int      `json:"souls"`
				ArrowsHit                                     int      `json:"arrows_hit"`
				BlocksPlaced                                  int      `json:"blocks_placed"`
				KillsKitBasicSoloDefault                      int      `json:"kills_kit_basic_solo_default"`
				Games                                         int      `json:"games"`
				Kills                                         int      `json:"kills"`
				ArrowsShot                                    int      `json:"arrows_shot"`
				WinsSoloInsane                                int      `json:"wins_solo_insane"`
				GamesSolo                                     int      `json:"games_solo"`
				KillsSoloInsane                               int      `json:"kills_solo_insane"`
				WinsSolo                                      int      `json:"wins_solo"`
				Wins                                          int      `json:"wins"`
				WinsKitBasicSoloDefault                       int      `json:"wins_kit_basic_solo_default"`
				SoulsGathered                                 int      `json:"souls_gathered"`
				ItemsEnchanted                                int      `json:"items_enchanted"`
				KillsSolo                                     int      `json:"kills_solo"`
				GamesKitBasicSoloDefault                      int      `json:"games_kit_basic_solo_default"`
				SoulWell                                      int      `json:"soul_well"`
				UsedSoulWell                                  bool     `json:"usedSoulWell"`
				Packages                                      []string `json:"packages"`
				DeathsSoloInsane                              int      `json:"deaths_solo_insane"`
				LossesSoloInsane                              int      `json:"losses_solo_insane"`
				XezbethLuck                                   int      `json:"xezbeth_luck"`
				AssistsSolo                                   int      `json:"assists_solo"`
				AssistsKitBasicSoloDefault                    int      `json:"assists_kit_basic_solo_default"`
				EnderpearlsThrown                             int      `json:"enderpearls_thrown"`
				Assists                                       int      `json:"assists"`
				EggThrown                                     int      `json:"egg_thrown"`
				ActiveKitSOLO                                 string   `json:"activeKit_SOLO"`
				LossesKitBasicSoloEcologist                   int      `json:"losses_kit_basic_solo_ecologist"`
				DeathsKitBasicSoloEcologist                   int      `json:"deaths_kit_basic_solo_ecologist"`
				KillsKitBasicSoloEcologist                    int      `json:"kills_kit_basic_solo_ecologist"`
				SurvivedPlayersKitBasicSoloEcologist          int      `json:"survived_players_kit_basic_solo_ecologist"`
				GamesKitBasicSoloEcologist                    int      `json:"games_kit_basic_solo_ecologist"`
				WinsKitBasicSoloEcologist                     int      `json:"wins_kit_basic_solo_ecologist"`
				WinsSoloNormal                                int      `json:"wins_solo_normal"`
				KillsSoloNormal                               int      `json:"kills_solo_normal"`
				KillsKitMiningTeamDefault                     int      `json:"kills_kit_mining_team_default"`
				SurvivedPlayersTeam                           int      `json:"survived_players_team"`
				LossesKitMiningTeamDefault                    int      `json:"losses_kit_mining_team_default"`
				LossesTeamInsane                              int      `json:"losses_team_insane"`
				DeathsKitMiningTeamDefault                    int      `json:"deaths_kit_mining_team_default"`
				DeathsTeam                                    int      `json:"deaths_team"`
				KillsTeam                                     int      `json:"kills_team"`
				DeathsTeamInsane                              int      `json:"deaths_team_insane"`
				KillsTeamInsane                               int      `json:"kills_team_insane"`
				SurvivedPlayersKitMiningTeamDefault           int      `json:"survived_players_kit_mining_team_default"`
				LossesTeam                                    int      `json:"losses_team"`
				TeamResistanceBoost                           int      `json:"team_resistance_boost"`
				SoulWellLegendaries                           int      `json:"soul_well_legendaries"`
				RefillChestDestroy                            int      `json:"refill_chest_destroy"`
				PaidSouls                                     int      `json:"paid_souls"`
				SoulWellRares                                 int      `json:"soul_well_rares"`
				SoloResistanceBoost                           int      `json:"solo_resistance_boost"`
				ActiveKitTEAM                                 string   `json:"activeKit_TEAM"`
				LossesKitSupportingTeamHealer                 int      `json:"losses_kit_supporting_team_healer"`
				SurvivedPlayersKitSupportingTeamHealer        int      `json:"survived_players_kit_supporting_team_healer"`
				DeathsKitSupportingTeamHealer                 int      `json:"deaths_kit_supporting_team_healer"`
				WinsTeamInsane                                int      `json:"wins_team_insane"`
				GamesKitSupportingTeamHealer                  int      `json:"games_kit_supporting_team_healer"`
				KillsKitSupportingTeamHealer                  int      `json:"kills_kit_supporting_team_healer"`
				WinsKitSupportingTeamHealer                   int      `json:"wins_kit_supporting_team_healer"`
				WinsTeam                                      int      `json:"wins_team"`
				GamesTeam                                     int      `json:"games_team"`
				VotesTemple                                   int      `json:"votes_Temple"`
				AssistsKitSupportingTeamHealer                int      `json:"assists_kit_supporting_team_healer"`
				AssistsTeam                                   int      `json:"assists_team"`
				AssistsKitBasicSoloEcologist                  int      `json:"assists_kit_basic_solo_ecologist"`
				VotesShattered                                int      `json:"votes_Shattered"`
				VotesShire                                    int      `json:"votes_Shire"`
				TeamEnderMastery                              int      `json:"team_ender_mastery"`
				TeamMiningExpertise                           int      `json:"team_mining_expertise"`
				SoloMiningExpertise                           int      `json:"solo_mining_expertise"`
				VotesDwarven                                  int      `json:"votes_Dwarven"`
				TeamInstantSmelting                           int      `json:"team_instant_smelting"`
				GamesKitAttackingTeamHunter                   int      `json:"games_kit_attacking_team_hunter"`
				KillsKitAttackingTeamHunter                   int      `json:"kills_kit_attacking_team_hunter"`
				SurvivedPlayersKitAttackingTeamHunter         int      `json:"survived_players_kit_attacking_team_hunter"`
				WinsKitAttackingTeamHunter                    int      `json:"wins_kit_attacking_team_hunter"`
				DeathsKitAttackingTeamHunter                  int      `json:"deaths_kit_attacking_team_hunter"`
				LossesKitAttackingTeamHunter                  int      `json:"losses_kit_attacking_team_hunter"`
				VotesAtuin                                    int      `json:"votes_Atuin"`
				SurvivedPlayersKitAttackingTeamKnight         int      `json:"survived_players_kit_attacking_team_knight"`
				GamesKitAttackingTeamKnight                   int      `json:"games_kit_attacking_team_knight"`
				WinsKitAttackingTeamKnight                    int      `json:"wins_kit_attacking_team_knight"`
				KillsKitAttackingTeamKnight                   int      `json:"kills_kit_attacking_team_knight"`
				KillsTeamNormal                               int      `json:"kills_team_normal"`
				WinsTeamNormal                                int      `json:"wins_team_normal"`
				DeathsTeamNormal                              int      `json:"deaths_team_normal"`
				LossesTeamNormal                              int      `json:"losses_team_normal"`
				LossesKitAttackingTeamKnight                  int      `json:"losses_kit_attacking_team_knight"`
				DeathsKitAttackingTeamKnight                  int      `json:"deaths_kit_attacking_team_knight"`
				AssistsKitAttackingTeamKnight                 int      `json:"assists_kit_attacking_team_knight"`
				SoloBlazingArrows                             int      `json:"solo_blazing_arrows"`
				VotesCongo                                    int      `json:"votes_Congo"`
				DeathsKitAttackingTeamScout                   int      `json:"deaths_kit_attacking_team_scout"`
				WinsKitAttackingTeamScout                     int      `json:"wins_kit_attacking_team_scout"`
				SurvivedPlayersKitAttackingTeamScout          int      `json:"survived_players_kit_attacking_team_scout"`
				KillsKitAttackingTeamScout                    int      `json:"kills_kit_attacking_team_scout"`
				GamesKitAttackingTeamScout                    int      `json:"games_kit_attacking_team_scout"`
				AssistsKitAttackingTeamScout                  int      `json:"assists_kit_attacking_team_scout"`
				LossesKitAttackingTeamScout                   int      `json:"losses_kit_attacking_team_scout"`
				VotesLongIsland                               int      `json:"votes_LongIsland"`
				SurvivedPlayersKitSupportingTeamArmorsmith    int      `json:"survived_players_kit_supporting_team_armorsmith"`
				GamesKitSupportingTeamArmorsmith              int      `json:"games_kit_supporting_team_armorsmith"`
				WinsKitSupportingTeamArmorsmith               int      `json:"wins_kit_supporting_team_armorsmith"`
				KillsKitSupportingTeamArmorsmith              int      `json:"kills_kit_supporting_team_armorsmith"`
				VotesToadstool                                int      `json:"votes_Toadstool"`
				DeathsKitSupportingTeamArmorsmith             int      `json:"deaths_kit_supporting_team_armorsmith"`
				LossesKitSupportingTeamArmorsmith             int      `json:"losses_kit_supporting_team_armorsmith"`
				AssistsKitSupportingTeamArmorsmith            int      `json:"assists_kit_supporting_team_armorsmith"`
				VotesFrostbite                                int      `json:"votes_Frostbite"`
				WinsKitMiningTeamSpeleologist                 int      `json:"wins_kit_mining_team_speleologist"`
				GamesKitMiningTeamSpeleologist                int      `json:"games_kit_mining_team_speleologist"`
				KillsKitMiningTeamSpeleologist                int      `json:"kills_kit_mining_team_speleologist"`
				AssistsKitMiningTeamSpeleologist              int      `json:"assists_kit_mining_team_speleologist"`
				SurvivedPlayersKitMiningTeamSpeleologist      int      `json:"survived_players_kit_mining_team_speleologist"`
				DeathsKitMiningTeamSpeleologist               int      `json:"deaths_kit_mining_team_speleologist"`
				LossesKitMiningTeamSpeleologist               int      `json:"losses_kit_mining_team_speleologist"`
				VotesDragonice                                int      `json:"votes_Dragonice"`
				SoloInstantSmelting                           int      `json:"solo_instant_smelting"`
				GamesKitDefendingTeamArmorer                  int      `json:"games_kit_defending_team_armorer"`
				SurvivedPlayersKitDefendingTeamArmorer        int      `json:"survived_players_kit_defending_team_armorer"`
				WinsKitDefendingTeamArmorer                   int      `json:"wins_kit_defending_team_armorer"`
				KillsKitDefendingTeamArmorer                  int      `json:"kills_kit_defending_team_armorer"`
				AssistsKitDefendingTeamArmorer                int      `json:"assists_kit_defending_team_armorer"`
				VotesSiege                                    int      `json:"votes_Siege"`
				WinsKitBasicSoloArmorsmith                    int      `json:"wins_kit_basic_solo_armorsmith"`
				SurvivedPlayersKitBasicSoloArmorsmith         int      `json:"survived_players_kit_basic_solo_armorsmith"`
				KillsKitBasicSoloArmorsmith                   int      `json:"kills_kit_basic_solo_armorsmith"`
				GamesKitBasicSoloArmorsmith                   int      `json:"games_kit_basic_solo_armorsmith"`
				VotesElven                                    int      `json:"votes_Elven"`
				VotesOnionring                                int      `json:"votes_Onionring"`
				VotesOverfall                                 int      `json:"votes_Overfall"`
				TeamArrowRecovery                             int      `json:"team_arrow_recovery"`
				SoloMarksmanship                              int      `json:"solo_marksmanship"`
				SoloEnderMastery                              int      `json:"solo_ender_mastery"`
				TeamSpeedBoost                                int      `json:"team_speed_boost"`
				TeamBlazingArrows                             int      `json:"team_blazing_arrows"`
				DeathsKitAdvancedSoloCannoneer                int      `json:"deaths_kit_advanced_solo_cannoneer"`
				LossesKitAdvancedSoloCannoneer                int      `json:"losses_kit_advanced_solo_cannoneer"`
				SurvivedPlayersKitAdvancedSoloCannoneer       int      `json:"survived_players_kit_advanced_solo_cannoneer"`
				KillsKitAdvancedSoloCannoneer                 int      `json:"kills_kit_advanced_solo_cannoneer"`
				WinsKitAdvancedSoloCannoneer                  int      `json:"wins_kit_advanced_solo_cannoneer"`
				GamesKitAdvancedSoloCannoneer                 int      `json:"games_kit_advanced_solo_cannoneer"`
				LossesKitBasicSoloArmorsmith                  int      `json:"losses_kit_basic_solo_armorsmith"`
				DeathsKitBasicSoloArmorsmith                  int      `json:"deaths_kit_basic_solo_armorsmith"`
				SurvivedPlayersKitAdvancedSoloPyro            int      `json:"survived_players_kit_advanced_solo_pyro"`
				DeathsKitAdvancedSoloPyro                     int      `json:"deaths_kit_advanced_solo_pyro"`
				LossesKitAdvancedSoloPyro                     int      `json:"losses_kit_advanced_solo_pyro"`
				GamesKitAdvancedSoloPyro                      int      `json:"games_kit_advanced_solo_pyro"`
				WinsKitAdvancedSoloPyro                       int      `json:"wins_kit_advanced_solo_pyro"`
				KillsKitAdvancedSoloPyro                      int      `json:"kills_kit_advanced_solo_pyro"`
				AssistsKitAdvancedSoloPyro                    int      `json:"assists_kit_advanced_solo_pyro"`
				DeathsKitAttackingTeamSnowman                 int      `json:"deaths_kit_attacking_team_snowman"`
				SurvivedPlayersKitAttackingTeamSnowman        int      `json:"survived_players_kit_attacking_team_snowman"`
				LossesKitAttackingTeamSnowman                 int      `json:"losses_kit_attacking_team_snowman"`
				GamesKitAttackingTeamSnowman                  int      `json:"games_kit_attacking_team_snowman"`
				KillsKitAttackingTeamSnowman                  int      `json:"kills_kit_attacking_team_snowman"`
				WinsKitDefendingTeamGuardian                  int      `json:"wins_kit_defending_team_guardian"`
				SurvivedPlayersKitDefendingTeamGuardian       int      `json:"survived_players_kit_defending_team_guardian"`
				GamesKitDefendingTeamGuardian                 int      `json:"games_kit_defending_team_guardian"`
				KillsKitDefendingTeamGuardian                 int      `json:"kills_kit_defending_team_guardian"`
				LossesKitDefendingTeamGuardian                int      `json:"losses_kit_defending_team_guardian"`
				DeathsKitDefendingTeamGuardian                int      `json:"deaths_kit_defending_team_guardian"`
				AssistsKitDefendingTeamGuardian               int      `json:"assists_kit_defending_team_guardian"`
				AssistsKitAttackingTeamSnowman                int      `json:"assists_kit_attacking_team_snowman"`
				WinsKitAttackingTeamSnowman                   int      `json:"wins_kit_attacking_team_snowman"`
				AssistsKitAdvancedSoloCannoneer               int      `json:"assists_kit_advanced_solo_cannoneer"`
				ActiveKitMEGA                                 string   `json:"activeKit_MEGA"`
				DeathsKitMegaMegaDefault                      int      `json:"deaths_kit_mega_mega_default"`
				SurvivedPlayersMega                           int      `json:"survived_players_mega"`
				DeathsMegaNormal                              int      `json:"deaths_mega_normal"`
				LossesMegaNormal                              int      `json:"losses_mega_normal"`
				DeathsMega                                    int      `json:"deaths_mega"`
				SurvivedPlayersKitMegaMegaDefault             int      `json:"survived_players_kit_mega_mega_default"`
				LossesMega                                    int      `json:"losses_mega"`
				LossesKitMegaMegaDefault                      int      `json:"losses_kit_mega_mega_default"`
				KillsMega                                     int      `json:"kills_mega"`
				KillsKitMegaMegaDefault                       int      `json:"kills_kit_mega_mega_default"`
				WinsMega                                      int      `json:"wins_mega"`
				GamesKitMegaMegaDefault                       int      `json:"games_kit_mega_mega_default"`
				GamesMega                                     int      `json:"games_mega"`
				WinsKitMegaMegaDefault                        int      `json:"wins_kit_mega_mega_default"`
				KillsMegaNormal                               int      `json:"kills_mega_normal"`
				WinsMegaNormal                                int      `json:"wins_mega_normal"`
				AssistsKitMegaMegaDefault                     int      `json:"assists_kit_mega_mega_default"`
				AssistsMega                                   int      `json:"assists_mega"`
				SurvivedPlayersKitAdvancedSoloArmorer         int      `json:"survived_players_kit_advanced_solo_armorer"`
				GamesKitAdvancedSoloArmorer                   int      `json:"games_kit_advanced_solo_armorer"`
				WinsKitAdvancedSoloArmorer                    int      `json:"wins_kit_advanced_solo_armorer"`
				KillsKitAdvancedSoloArmorer                   int      `json:"kills_kit_advanced_solo_armorer"`
				LossesKitDefendingTeamArmorer                 int      `json:"losses_kit_defending_team_armorer"`
				DeathsKitDefendingTeamArmorer                 int      `json:"deaths_kit_defending_team_armorer"`
				KitMegaMegaPaladin                            int      `json:"kit_mega_mega_paladin"`
				KitMegaMegaArmorer                            int      `json:"kit_mega_mega_armorer"`
				KillsKitMegaMegaPaladin                       int      `json:"kills_kit_mega_mega_paladin"`
				SurvivedPlayersKitMegaMegaPaladin             int      `json:"survived_players_kit_mega_mega_paladin"`
				WinsKitMegaMegaPaladin                        int      `json:"wins_kit_mega_mega_paladin"`
				GamesKitMegaMegaPaladin                       int      `json:"games_kit_mega_mega_paladin"`
				DeathsKitMegaMegaPaladin                      int      `json:"deaths_kit_mega_mega_paladin"`
				LossesKitMegaMegaPaladin                      int      `json:"losses_kit_mega_mega_paladin"`
				MegaRusher                                    int      `json:"mega_rusher"`
				AssistsKitMegaMegaPaladin                     int      `json:"assists_kit_mega_mega_paladin"`
				MegaEnderMastery                              int      `json:"mega_ender_mastery"`
				LossesKitMegaMegaArmorer                      int      `json:"losses_kit_mega_mega_armorer"`
				AssistsKitMegaMegaArmorer                     int      `json:"assists_kit_mega_mega_armorer"`
				SurvivedPlayersKitMegaMegaArmorer             int      `json:"survived_players_kit_mega_mega_armorer"`
				DeathsKitMegaMegaArmorer                      int      `json:"deaths_kit_mega_mega_armorer"`
				KillsKitMegaMegaArmorer                       int      `json:"kills_kit_mega_mega_armorer"`
				WinsKitMegaMegaArmorer                        int      `json:"wins_kit_mega_mega_armorer"`
				GamesKitMegaMegaArmorer                       int      `json:"games_kit_mega_mega_armorer"`
				VotesDwarfFortress                            int      `json:"votes_Dwarf Fortress"`
				MegaTank                                      int      `json:"mega_tank"`
				AssistsKitBasicSoloArmorsmith                 int      `json:"assists_kit_basic_solo_armorsmith"`
				KitMegaMegaBaseballPlayer                     int      `json:"kit_mega_mega_baseball-player"`
				KillsKitMegaMegaBaseballPlayer                int      `json:"kills_kit_mega_mega_baseball-player"`
				AssistsKitMegaMegaBaseballPlayer              int      `json:"assists_kit_mega_mega_baseball-player"`
				WinsKitMegaMegaBaseballPlayer                 int      `json:"wins_kit_mega_mega_baseball-player"`
				SurvivedPlayersKitMegaMegaBaseballPlayer      int      `json:"survived_players_kit_mega_mega_baseball-player"`
				GamesKitMegaMegaBaseballPlayer                int      `json:"games_kit_mega_mega_baseball-player"`
				LossesKitMegaMegaBaseballPlayer               int      `json:"losses_kit_mega_mega_baseball-player"`
				DeathsKitMegaMegaBaseballPlayer               int      `json:"deaths_kit_mega_mega_baseball-player"`
				MegaMiningExpertise                           int      `json:"mega_mining_expertise"`
				VotesSteampunk                                int      `json:"votes_Steampunk"`
				MegaJuggernaut                                int      `json:"mega_juggernaut"`
				MegaArrowRecovery                             int      `json:"mega_arrow_recovery"`
				TeamJuggernaut                                int      `json:"team_juggernaut"`
				SoloJuggernaut                                int      `json:"solo_juggernaut"`
				KitMegaMegaCannoneer                          int      `json:"kit_mega_mega_cannoneer"`
				SurvivedPlayersKitMegaMegaCannoneer           int      `json:"survived_players_kit_mega_mega_cannoneer"`
				AssistsKitMegaMegaCannoneer                   int      `json:"assists_kit_mega_mega_cannoneer"`
				WinsKitMegaMegaCannoneer                      int      `json:"wins_kit_mega_mega_cannoneer"`
				GamesKitMegaMegaCannoneer                     int      `json:"games_kit_mega_mega_cannoneer"`
				KillsKitMegaMegaCannoneer                     int      `json:"kills_kit_mega_mega_cannoneer"`
				LossesKitMegaMegaCannoneer                    int      `json:"losses_kit_mega_mega_cannoneer"`
				DeathsKitMegaMegaCannoneer                    int      `json:"deaths_kit_mega_mega_cannoneer"`
				SoloArrowRecovery                             int      `json:"solo_arrow_recovery"`
				SoloBulldozer                                 int      `json:"solo_bulldozer"`
				LossesKitAdvancedSoloArmorer                  int      `json:"losses_kit_advanced_solo_armorer"`
				DeathsKitAdvancedSoloArmorer                  int      `json:"deaths_kit_advanced_solo_armorer"`
				AssistsKitAdvancedSoloArmorer                 int      `json:"assists_kit_advanced_solo_armorer"`
				SoloSpeedBoost                                int      `json:"solo_speed_boost"`
				VotesStrata                                   int      `json:"votes_Strata"`
				VotesJinzhou                                  int      `json:"votes_Jinzhou"`
				VotesCanopy                                   int      `json:"votes_Canopy"`
				VotesTribal                                   int      `json:"votes_Tribal"`
				VotesEntangled                                int      `json:"votes_Entangled"`
				VotesSkychurch                                int      `json:"votes_Skychurch"`
				KitMegaMegaWitch                              int      `json:"kit_mega_mega_witch"`
				LossesKitMegaMegaWitch                        int      `json:"losses_kit_mega_mega_witch"`
				AssistsKitMegaMegaWitch                       int      `json:"assists_kit_mega_mega_witch"`
				DeathsKitMegaMegaWitch                        int      `json:"deaths_kit_mega_mega_witch"`
				KillsKitMegaMegaWitch                         int      `json:"kills_kit_mega_mega_witch"`
				SurvivedPlayersKitMegaMegaWitch               int      `json:"survived_players_kit_mega_mega_witch"`
				WinsKitAdvancedSoloEnchanter                  int      `json:"wins_kit_advanced_solo_enchanter"`
				GamesKitAdvancedSoloEnchanter                 int      `json:"games_kit_advanced_solo_enchanter"`
				KillsKitAdvancedSoloEnchanter                 int      `json:"kills_kit_advanced_solo_enchanter"`
				SurvivedPlayersKitAdvancedSoloEnchanter       int      `json:"survived_players_kit_advanced_solo_enchanter"`
				LossesKitAdvancedSoloEnchanter                int      `json:"losses_kit_advanced_solo_enchanter"`
				AssistsKitAdvancedSoloEnchanter               int      `json:"assists_kit_advanced_solo_enchanter"`
				DeathsKitAdvancedSoloEnchanter                int      `json:"deaths_kit_advanced_solo_enchanter"`
				GamesKitMegaMegaWitch                         int      `json:"games_kit_mega_mega_witch"`
				WinsKitMegaMegaWitch                          int      `json:"wins_kit_mega_mega_witch"`
				KitMegaMegaArmorsmith                         int      `json:"kit_mega_mega_armorsmith"`
				KillsKitMegaMegaArmorsmith                    int      `json:"kills_kit_mega_mega_armorsmith"`
				AssistsKitMegaMegaArmorsmith                  int      `json:"assists_kit_mega_mega_armorsmith"`
				WinsKitMegaMegaArmorsmith                     int      `json:"wins_kit_mega_mega_armorsmith"`
				GamesKitMegaMegaArmorsmith                    int      `json:"games_kit_mega_mega_armorsmith"`
				SurvivedPlayersKitMegaMegaArmorsmith          int      `json:"survived_players_kit_mega_mega_armorsmith"`
				GamesKitAdvancedSoloFarmer                    int      `json:"games_kit_advanced_solo_farmer"`
				WinsKitAdvancedSoloFarmer                     int      `json:"wins_kit_advanced_solo_farmer"`
				SurvivedPlayersKitAdvancedSoloFarmer          int      `json:"survived_players_kit_advanced_solo_farmer"`
				KillsKitAdvancedSoloFarmer                    int      `json:"kills_kit_advanced_solo_farmer"`
				MegaBlazingArrows                             int      `json:"mega_blazing_arrows"`
				LossesKitMegaMegaArmorsmith                   int      `json:"losses_kit_mega_mega_armorsmith"`
				DeathsKitMegaMegaArmorsmith                   int      `json:"deaths_kit_mega_mega_armorsmith"`
				VotesTowers                                   int      `json:"votes_Towers"`
				KitMegaMegaKnight                             int      `json:"kit_mega_mega_knight"`
				AssistsKitMegaMegaKnight                      int      `json:"assists_kit_mega_mega_knight"`
				KillsKitMegaMegaKnight                        int      `json:"kills_kit_mega_mega_knight"`
				WinsKitMegaMegaKnight                         int      `json:"wins_kit_mega_mega_knight"`
				DeathsKitMegaMegaKnight                       int      `json:"deaths_kit_mega_mega_knight"`
				SurvivedPlayersKitMegaMegaKnight              int      `json:"survived_players_kit_mega_mega_knight"`
				KitMegaMegaScout                              int      `json:"kit_mega_mega_scout"`
				KitMegaMegaSkeletor                           int      `json:"kit_mega_mega_skeletor"`
				KitMegaMegaHunter                             int      `json:"kit_mega_mega_hunter"`
				KitMegaMegaHealer                             int      `json:"kit_mega_mega_healer"`
				WinsKitDefendingTeamBaseballPlayer            int      `json:"wins_kit_defending_team_baseball-player"`
				SurvivedPlayersKitDefendingTeamBaseballPlayer int      `json:"survived_players_kit_defending_team_baseball-player"`
				KillsKitDefendingTeamBaseballPlayer           int      `json:"kills_kit_defending_team_baseball-player"`
				GamesKitDefendingTeamBaseballPlayer           int      `json:"games_kit_defending_team_baseball-player"`
				VotesAegis                                    int      `json:"votes_Aegis"`
				KillsWeeklyB                                  int      `json:"kills_weekly_b"`
				KillsMonthlyB                                 int      `json:"kills_monthly_b"`
				GamesKitMegaMegaKnight                        int      `json:"games_kit_mega_mega_knight"`
				LossesKitMegaMegaKnight                       int      `json:"losses_kit_mega_mega_knight"`
				KillsWeeklyA                                  int      `json:"kills_weekly_a"`
				HarvestingSeason                              int      `json:"harvesting_season"`
				GamesKitMiningTeamCannoneer                   int      `json:"games_kit_mining_team_cannoneer"`
				SurvivedPlayersKitMiningTeamCannoneer         int      `json:"survived_players_kit_mining_team_cannoneer"`
				WinsKitMiningTeamCannoneer                    int      `json:"wins_kit_mining_team_cannoneer"`
				KillsKitMiningTeamCannoneer                   int      `json:"kills_kit_mining_team_cannoneer"`
				AssistsKitMiningTeamCannoneer                 int      `json:"assists_kit_mining_team_cannoneer"`
				AssistsKitDefendingTeamBaseballPlayer         int      `json:"assists_kit_defending_team_baseball-player"`
				LossesKitDefendingTeamBaseballPlayer          int      `json:"losses_kit_defending_team_baseball-player"`
				DeathsKitDefendingTeamBaseballPlayer          int      `json:"deaths_kit_defending_team_baseball-player"`
				LossesKitMiningTeamCannoneer                  int      `json:"losses_kit_mining_team_cannoneer"`
				DeathsKitMiningTeamCannoneer                  int      `json:"deaths_kit_mining_team_cannoneer"`
				KillsMonthlyA                                 int      `json:"kills_monthly_a"`
				AssistsKitMegaMegaScout                       int      `json:"assists_kit_mega_mega_scout"`
				KillsKitMegaMegaScout                         int      `json:"kills_kit_mega_mega_scout"`
				DeathsKitMegaMegaScout                        int      `json:"deaths_kit_mega_mega_scout"`
				SurvivedPlayersKitMegaMegaScout               int      `json:"survived_players_kit_mega_mega_scout"`
				LossesKitMegaMegaScout                        int      `json:"losses_kit_mega_mega_scout"`
				WinsKitMegaMegaScout                          int      `json:"wins_kit_mega_mega_scout"`
				GamesKitMegaMegaScout                         int      `json:"games_kit_mega_mega_scout"`
				LossesKitAdvancedSoloFarmer                   int      `json:"losses_kit_advanced_solo_farmer"`
				DeathsKitAdvancedSoloFarmer                   int      `json:"deaths_kit_advanced_solo_farmer"`
				KillsKitMegaMegaSkeletor                      int      `json:"kills_kit_mega_mega_skeletor"`
				SurvivedPlayersKitMegaMegaSkeletor            int      `json:"survived_players_kit_mega_mega_skeletor"`
				DeathsKitMegaMegaSkeletor                     int      `json:"deaths_kit_mega_mega_skeletor"`
				LossesKitMegaMegaSkeletor                     int      `json:"losses_kit_mega_mega_skeletor"`
				ActiveKitRANKED                               string   `json:"activeKit_RANKED"`
				WinstreakRanked                               int      `json:"winstreak_ranked"`
				AssistsRanked                                 int      `json:"assists_ranked"`
				Killstreak                                    int      `json:"killstreak"`
				GamesRanked                                   int      `json:"games_ranked"`
				WinsRankedNormal                              int      `json:"wins_ranked_normal"`
				KillsKitRankedRankedDefault                   int      `json:"kills_kit_ranked_ranked_default"`
				KillstreakRanked                              int      `json:"killstreak_ranked"`
				KillsRankedNormal                             int      `json:"kills_ranked_normal"`
				GamesKitRankedRankedDefault                   int      `json:"games_kit_ranked_ranked_default"`
				SurvivedPlayersRanked                         int      `json:"survived_players_ranked"`
				KillsRanked                                   int      `json:"kills_ranked"`
				KillstreakKitRankedRankedDefault              int      `json:"killstreak_kit_ranked_ranked_default"`
				Winstreak                                     int      `json:"winstreak"`
				WinstreakKitRankedRankedDefault               int      `json:"winstreak_kit_ranked_ranked_default"`
				AssistsKitRankedRankedDefault                 int      `json:"assists_kit_ranked_ranked_default"`
				SurvivedPlayersKitRankedRankedDefault         int      `json:"survived_players_kit_ranked_ranked_default"`
				WinsKitRankedRankedDefault                    int      `json:"wins_kit_ranked_ranked_default"`
				WinsRanked                                    int      `json:"wins_ranked"`
				HighestKillstreak                             int      `json:"highestKillstreak"`
				HighestWinstreak                              int      `json:"highestWinstreak"`
				LossesKitRankedRankedChampion                 int      `json:"losses_kit_ranked_ranked_champion"`
				DeathsKitRankedRankedChampion                 int      `json:"deaths_kit_ranked_ranked_champion"`
				AssistsKitRankedRankedChampion                int      `json:"assists_kit_ranked_ranked_champion"`
				LossesRankedNormal                            int      `json:"losses_ranked_normal"`
				DeathsRanked                                  int      `json:"deaths_ranked"`
				LossesRanked                                  int      `json:"losses_ranked"`
				DeathsRankedNormal                            int      `json:"deaths_ranked_normal"`
				GamesKitRankedRankedChampion                  int      `json:"games_kit_ranked_ranked_champion"`
				SurvivedPlayersKitRankedRankedChampion        int      `json:"survived_players_kit_ranked_ranked_champion"`
				KillstreakKitRankedRankedChampion             int      `json:"killstreak_kit_ranked_ranked_champion"`
				WinsKitRankedRankedChampion                   int      `json:"wins_kit_ranked_ranked_champion"`
				KillsKitRankedRankedChampion                  int      `json:"kills_kit_ranked_ranked_champion"`
				WinstreakKitRankedRankedChampion              int      `json:"winstreak_kit_ranked_ranked_champion"`
				MegaInstantSmelting                           int      `json:"mega_instant_smelting"`
				MegaNourishment                               int      `json:"mega_nourishment"`
				TeamNourishment                               int      `json:"team_nourishment"`
				SoloNourishment                               int      `json:"solo_nourishment"`
				MegaNotoriety                                 int      `json:"mega_notoriety"`
				TeamKnowledge                                 int      `json:"team_knowledge"`
				RankedInstantSmelting                         int      `json:"ranked_instant_smelting"`
				SoloKnowledge                                 int      `json:"solo_knowledge"`
				TeamSavior                                    int      `json:"team_savior"`
				TeamMarksmanship                              int      `json:"team_marksmanship"`
				KillstreakKitBasicSoloArmorsmith              int      `json:"killstreak_kit_basic_solo_armorsmith"`
				WinstreakKitBasicSoloArmorsmith               int      `json:"winstreak_kit_basic_solo_armorsmith"`
				KillstreakSolo                                int      `json:"killstreak_solo"`
				WinstreakSolo                                 int      `json:"winstreak_solo"`
				WinstreakKitEnderchestSoloEnderchest          int      `json:"winstreak_kit_enderchest_solo_enderchest"`
				GamesKitEnderchestSoloEnderchest              int      `json:"games_kit_enderchest_solo_enderchest"`
				WinsKitEnderchestSoloEnderchest               int      `json:"wins_kit_enderchest_solo_enderchest"`
				KillsKitEnderchestSoloEnderchest              int      `json:"kills_kit_enderchest_solo_enderchest"`
				SurvivedPlayersKitEnderchestSoloEnderchest    int      `json:"survived_players_kit_enderchest_solo_enderchest"`
				KillstreakKitEnderchestSoloEnderchest         int      `json:"killstreak_kit_enderchest_solo_enderchest"`
				LossesKitEnderchestSoloEnderchest             int      `json:"losses_kit_enderchest_solo_enderchest"`
				DeathsKitEnderchestSoloEnderchest             int      `json:"deaths_kit_enderchest_solo_enderchest"`
				KillstreakKitMegaMegaKnight                   int      `json:"killstreak_kit_mega_mega_knight"`
				WinstreakKitMegaMegaKnight                    int      `json:"winstreak_kit_mega_mega_knight"`
				WinstreakMega                                 int      `json:"winstreak_mega"`
				KillstreakMega                                int      `json:"killstreak_mega"`
				KillstreakKitAdvancedSoloFarmer               int      `json:"killstreak_kit_advanced_solo_farmer"`
				WinstreakKitAdvancedSoloFarmer                int      `json:"winstreak_kit_advanced_solo_farmer"`
				WinstreakKitAdvancedSoloPyro                  int      `json:"winstreak_kit_advanced_solo_pyro"`
				KillstreakKitAdvancedSoloPyro                 int      `json:"killstreak_kit_advanced_solo_pyro"`
				SurvivedPlayersKitAdvancedSoloHunter          int      `json:"survived_players_kit_advanced_solo_hunter"`
				DeathsKitAdvancedSoloHunter                   int      `json:"deaths_kit_advanced_solo_hunter"`
				LossesKitAdvancedSoloHunter                   int      `json:"losses_kit_advanced_solo_hunter"`
				SoloRevenge                                   int      `json:"solo_revenge"`
				GamesKitBasicSoloFisherman                    int      `json:"games_kit_basic_solo_fisherman"`
				SurvivedPlayersKitBasicSoloFisherman          int      `json:"survived_players_kit_basic_solo_fisherman"`
				WinsKitBasicSoloFisherman                     int      `json:"wins_kit_basic_solo_fisherman"`
				KillstreakKitBasicSoloFisherman               int      `json:"killstreak_kit_basic_solo_fisherman"`
				WinstreakKitBasicSoloFisherman                int      `json:"winstreak_kit_basic_solo_fisherman"`
				KillsKitBasicSoloFisherman                    int      `json:"kills_kit_basic_solo_fisherman"`
				WinstreakKitAdvancedSoloCannoneer             int      `json:"winstreak_kit_advanced_solo_cannoneer"`
				KillstreakKitAdvancedSoloCannoneer            int      `json:"killstreak_kit_advanced_solo_cannoneer"`
				KillstreakKitMegaMegaSkeletor                 int      `json:"killstreak_kit_mega_mega_skeletor"`
				WinsKitMegaMegaSkeletor                       int      `json:"wins_kit_mega_mega_skeletor"`
				GamesKitMegaMegaSkeletor                      int      `json:"games_kit_mega_mega_skeletor"`
				WinstreakKitMegaMegaSkeletor                  int      `json:"winstreak_kit_mega_mega_skeletor"`
				AssistsKitMegaMegaSkeletor                    int      `json:"assists_kit_mega_mega_skeletor"`
				TeamAnnoyOMite                                int      `json:"team_annoy-o-mite"`
				TeamFat                                       int      `json:"team_fat"`
				WinstreakKitAdvancedSoloEnchanter             int      `json:"winstreak_kit_advanced_solo_enchanter"`
				KillstreakKitAdvancedSoloEnchanter            int      `json:"killstreak_kit_advanced_solo_enchanter"`
				SoloAnnoyOMite                                int      `json:"solo_annoy-o-mite"`
				WinsKitRankedRankedArmorer                    int      `json:"wins_kit_ranked_ranked_armorer"`
				WinstreakKitRankedRankedArmorer               int      `json:"winstreak_kit_ranked_ranked_armorer"`
				GamesKitRankedRankedArmorer                   int      `json:"games_kit_ranked_ranked_armorer"`
				SurvivedPlayersKitRankedRankedArmorer         int      `json:"survived_players_kit_ranked_ranked_armorer"`
				KillsKitRankedRankedArmorer                   int      `json:"kills_kit_ranked_ranked_armorer"`
				KillstreakKitRankedRankedArmorer              int      `json:"killstreak_kit_ranked_ranked_armorer"`
				AssistsKitRankedRankedArmorer                 int      `json:"assists_kit_ranked_ranked_armorer"`
				LossesKitRankedRankedArmorer                  int      `json:"losses_kit_ranked_ranked_armorer"`
				DeathsKitRankedRankedArmorer                  int      `json:"deaths_kit_ranked_ranked_armorer"`
				WinstreakKitAdvancedSoloEnderman              int      `json:"winstreak_kit_advanced_solo_enderman"`
				WinsKitAdvancedSoloEnderman                   int      `json:"wins_kit_advanced_solo_enderman"`
				KillsKitAdvancedSoloEnderman                  int      `json:"kills_kit_advanced_solo_enderman"`
				SurvivedPlayersKitAdvancedSoloEnderman        int      `json:"survived_players_kit_advanced_solo_enderman"`
				GamesKitAdvancedSoloEnderman                  int      `json:"games_kit_advanced_solo_enderman"`
				KillstreakKitAdvancedSoloEnderman             int      `json:"killstreak_kit_advanced_solo_enderman"`
				AssistsKitAdvancedSoloFarmer                  int      `json:"assists_kit_advanced_solo_farmer"`
				WinsKitBasicSoloSnowman                       int      `json:"wins_kit_basic_solo_snowman"`
				WinstreakKitBasicSoloSnowman                  int      `json:"winstreak_kit_basic_solo_snowman"`
				GamesKitBasicSoloSnowman                      int      `json:"games_kit_basic_solo_snowman"`
				SurvivedPlayersKitBasicSoloSnowman            int      `json:"survived_players_kit_basic_solo_snowman"`
				KillsKitBasicSoloSnowman                      int      `json:"kills_kit_basic_solo_snowman"`
				KillstreakKitBasicSoloSnowman                 int      `json:"killstreak_kit_basic_solo_snowman"`
				WinstreakKitBasicSoloSpeleologist             int      `json:"winstreak_kit_basic_solo_speleologist"`
				WinsKitBasicSoloSpeleologist                  int      `json:"wins_kit_basic_solo_speleologist"`
				KillsKitBasicSoloSpeleologist                 int      `json:"kills_kit_basic_solo_speleologist"`
				GamesKitBasicSoloSpeleologist                 int      `json:"games_kit_basic_solo_speleologist"`
				KillstreakKitBasicSoloSpeleologist            int      `json:"killstreak_kit_basic_solo_speleologist"`
				SurvivedPlayersKitBasicSoloSpeleologist       int      `json:"survived_players_kit_basic_solo_speleologist"`
				DeathsKitBasicSoloFisherman                   int      `json:"deaths_kit_basic_solo_fisherman"`
				LossesKitBasicSoloFisherman                   int      `json:"losses_kit_basic_solo_fisherman"`
				MegaMarksmanship                              int      `json:"mega_marksmanship"`
				DeathsKitBasicSoloSnowman                     int      `json:"deaths_kit_basic_solo_snowman"`
				LossesKitBasicSoloSnowman                     int      `json:"losses_kit_basic_solo_snowman"`
				ActiveKillEffect                              string   `json:"activeKillEffect"`
				ActiveProjectileTrail                         string   `json:"activeProjectileTrail"`
				AssistsKitBasicSoloTroll                      int      `json:"assists_kit_basic_solo_troll"`
				WinstreakKitBasicSoloTroll                    int      `json:"winstreak_kit_basic_solo_troll"`
				KillsKitBasicSoloTroll                        int      `json:"kills_kit_basic_solo_troll"`
				KillstreakKitBasicSoloTroll                   int      `json:"killstreak_kit_basic_solo_troll"`
				SurvivedPlayersKitBasicSoloTroll              int      `json:"survived_players_kit_basic_solo_troll"`
				WinsKitBasicSoloTroll                         int      `json:"wins_kit_basic_solo_troll"`
				GamesKitBasicSoloTroll                        int      `json:"games_kit_basic_solo_troll"`
				SoloFat                                       int      `json:"solo_fat"`
				DeathsKitBasicSoloTroll                       int      `json:"deaths_kit_basic_solo_troll"`
				LossesKitBasicSoloTroll                       int      `json:"losses_kit_basic_solo_troll"`
				KillstreakKitSupportingTeamPharaoh            int      `json:"killstreak_kit_supporting_team_pharaoh"`
				SurvivedPlayersKitSupportingTeamPharaoh       int      `json:"survived_players_kit_supporting_team_pharaoh"`
				WinstreakKitSupportingTeamPharaoh             int      `json:"winstreak_kit_supporting_team_pharaoh"`
				WinsKitSupportingTeamPharaoh                  int      `json:"wins_kit_supporting_team_pharaoh"`
				KillstreakTeam                                int      `json:"killstreak_team"`
				WinstreakTeam                                 int      `json:"winstreak_team"`
				KillsKitSupportingTeamPharaoh                 int      `json:"kills_kit_supporting_team_pharaoh"`
				GamesKitSupportingTeamPharaoh                 int      `json:"games_kit_supporting_team_pharaoh"`
				AssistsKitBasicSoloFisherman                  int      `json:"assists_kit_basic_solo_fisherman"`
				DeathsKitMegaMegaHunter                       int      `json:"deaths_kit_mega_mega_hunter"`
				KillsKitMegaMegaHunter                        int      `json:"kills_kit_mega_mega_hunter"`
				SurvivedPlayersKitMegaMegaHunter              int      `json:"survived_players_kit_mega_mega_hunter"`
				LossesKitMegaMegaHunter                       int      `json:"losses_kit_mega_mega_hunter"`
				ActiveVictoryDance                            string   `json:"activeVictoryDance"`
				KillstreakKitRankedRankedPyromancer           int      `json:"killstreak_kit_ranked_ranked_pyromancer"`
				GamesKitRankedRankedPyromancer                int      `json:"games_kit_ranked_ranked_pyromancer"`
				WinstreakKitRankedRankedPyromancer            int      `json:"winstreak_kit_ranked_ranked_pyromancer"`
				KillsKitRankedRankedPyromancer                int      `json:"kills_kit_ranked_ranked_pyromancer"`
				SurvivedPlayersKitRankedRankedPyromancer      int      `json:"survived_players_kit_ranked_ranked_pyromancer"`
				WinsKitRankedRankedPyromancer                 int      `json:"wins_kit_ranked_ranked_pyromancer"`
				DeathsKitRankedRankedPyromancer               int      `json:"deaths_kit_ranked_ranked_pyromancer"`
				LossesKitRankedRankedPyromancer               int      `json:"losses_kit_ranked_ranked_pyromancer"`
				AssistsKitRankedRankedPyromancer              int      `json:"assists_kit_ranked_ranked_pyromancer"`
				KillstreakKitMegaMegaPaladin                  int      `json:"killstreak_kit_mega_mega_paladin"`
				WinstreakKitMegaMegaPaladin                   int      `json:"winstreak_kit_mega_mega_paladin"`
				KillstreakKitDefendingTeamArmorer             int      `json:"killstreak_kit_defending_team_armorer"`
				WinstreakKitDefendingTeamArmorer              int      `json:"winstreak_kit_defending_team_armorer"`
				KillsKitBasicSoloPharaoh                      int      `json:"kills_kit_basic_solo_pharaoh"`
				WinsKitBasicSoloPharaoh                       int      `json:"wins_kit_basic_solo_pharaoh"`
				GamesKitBasicSoloPharaoh                      int      `json:"games_kit_basic_solo_pharaoh"`
				WinstreakKitBasicSoloPharaoh                  int      `json:"winstreak_kit_basic_solo_pharaoh"`
				SurvivedPlayersKitBasicSoloPharaoh            int      `json:"survived_players_kit_basic_solo_pharaoh"`
				KillstreakKitBasicSoloPharaoh                 int      `json:"killstreak_kit_basic_solo_pharaoh"`
				LossesKitBasicSoloPharaoh                     int      `json:"losses_kit_basic_solo_pharaoh"`
				DeathsKitBasicSoloPharaoh                     int      `json:"deaths_kit_basic_solo_pharaoh"`
				AssistsKitBasicSoloPharaoh                    int      `json:"assists_kit_basic_solo_pharaoh"`
				KitMegaMegaHellhound                          int      `json:"kit_mega_mega_hellhound"`
				KillstreakKitMegaMegaArmorer                  int      `json:"killstreak_kit_mega_mega_armorer"`
				WinstreakKitMegaMegaArmorer                   int      `json:"winstreak_kit_mega_mega_armorer"`
				WinstreakKitAdvancedSoloArmorer               int      `json:"winstreak_kit_advanced_solo_armorer"`
				KillstreakKitAdvancedSoloArmorer              int      `json:"killstreak_kit_advanced_solo_armorer"`
				WinstreakKitAdvancedSoloHunter                int      `json:"winstreak_kit_advanced_solo_hunter"`
				KillsKitAdvancedSoloHunter                    int      `json:"kills_kit_advanced_solo_hunter"`
				GamesKitAdvancedSoloHunter                    int      `json:"games_kit_advanced_solo_hunter"`
				WinsKitAdvancedSoloHunter                     int      `json:"wins_kit_advanced_solo_hunter"`
				KillstreakKitAdvancedSoloHunter               int      `json:"killstreak_kit_advanced_solo_hunter"`
				LossesKitAdvancedSoloEnderman                 int      `json:"losses_kit_advanced_solo_enderman"`
				DeathsKitAdvancedSoloEnderman                 int      `json:"deaths_kit_advanced_solo_enderman"`
				WinstreakKitMiningTeamCannoneer               int      `json:"winstreak_kit_mining_team_cannoneer"`
				KillstreakKitMiningTeamCannoneer              int      `json:"killstreak_kit_mining_team_cannoneer"`
				SoloBridger                                   int      `json:"solo_bridger"`
				SoloEnvironmentalExpert                       int      `json:"solo_environmental_expert"`
				SoloLuckyCharm                                int      `json:"solo_lucky_charm"`
				MeleeKillsSolo                                int      `json:"melee_kills_solo"`
				MeleeKillsKitBasicSoloGrenade                 int      `json:"melee_kills_kit_basic_solo_grenade"`
				TimePlayedKitBasicSoloGrenade                 int      `json:"time_played_kit_basic_solo_grenade"`
				ChestsOpened                                  int      `json:"chests_opened"`
				MostKillsGameKitBasicSoloGrenade              int      `json:"most_kills_game_kit_basic_solo_grenade"`
				MeleeKills                                    int      `json:"melee_kills"`
				KillsKitBasicSoloGrenade                      int      `json:"kills_kit_basic_solo_grenade"`
				GamesKitBasicSoloGrenade                      int      `json:"games_kit_basic_solo_grenade"`
				SurvivedPlayersKitBasicSoloGrenade            int      `json:"survived_players_kit_basic_solo_grenade"`
				MostKillsGameSolo                             int      `json:"most_kills_game_solo"`
				DeathsKitBasicSoloGrenade                     int      `json:"deaths_kit_basic_solo_grenade"`
				ChestsOpenedSolo                              int      `json:"chests_opened_solo"`
				LossesKitBasicSoloGrenade                     int      `json:"losses_kit_basic_solo_grenade"`
				TimePlayedSolo                                int      `json:"time_played_solo"`
				TimePlayed                                    int      `json:"time_played"`
				MostKillsGame                                 int      `json:"most_kills_game"`
				ChestsOpenedKitBasicSoloGrenade               int      `json:"chests_opened_kit_basic_solo_grenade"`
				LongestBowShotKitBasicSoloGrenade             int      `json:"longest_bow_shot_kit_basic_solo_grenade"`
				LongestBowShot                                int      `json:"longest_bow_shot"`
				FastestWin                                    int      `json:"fastest_win"`
				FastestWinSolo                                int      `json:"fastest_win_solo"`
				FastestWinKitBasicSoloGrenade                 int      `json:"fastest_win_kit_basic_solo_grenade"`
				LongestBowShotSolo                            int      `json:"longest_bow_shot_solo"`
				VoidKills                                     int      `json:"void_kills"`
				ArrowsShotKitBasicSoloGrenade                 int      `json:"arrows_shot_kit_basic_solo_grenade"`
				WinstreakKitBasicSoloGrenade                  int      `json:"winstreak_kit_basic_solo_grenade"`
				VoidKillsKitBasicSoloGrenade                  int      `json:"void_kills_kit_basic_solo_grenade"`
				ArrowsHitSolo                                 int      `json:"arrows_hit_solo"`
				WinsKitBasicSoloGrenade                       int      `json:"wins_kit_basic_solo_grenade"`
				KillstreakKitBasicSoloGrenade                 int      `json:"killstreak_kit_basic_solo_grenade"`
				ArrowsHitKitBasicSoloGrenade                  int      `json:"arrows_hit_kit_basic_solo_grenade"`
				VoidKillsSolo                                 int      `json:"void_kills_solo"`
				ArrowsShotSolo                                int      `json:"arrows_shot_solo"`
				AssistsKitMegaMegaHellhound                   int      `json:"assists_kit_mega_mega_hellhound"`
				LossesKitMegaMegaHellhound                    int      `json:"losses_kit_mega_mega_hellhound"`
				DeathsKitMegaMegaHellhound                    int      `json:"deaths_kit_mega_mega_hellhound"`
				ChestsOpenedKitMegaMegaHellhound              int      `json:"chests_opened_kit_mega_mega_hellhound"`
				ChestsOpenedMega                              int      `json:"chests_opened_mega"`
				MostKillsGameMega                             int      `json:"most_kills_game_mega"`
				MeleeKillsMega                                int      `json:"melee_kills_mega"`
				MostKillsGameKitMegaMegaHellhound             int      `json:"most_kills_game_kit_mega_mega_hellhound"`
				MeleeKillsKitMegaMegaHellhound                int      `json:"melee_kills_kit_mega_mega_hellhound"`
				KillsKitMegaMegaHellhound                     int      `json:"kills_kit_mega_mega_hellhound"`
				TimePlayedKitMegaMegaHellhound                int      `json:"time_played_kit_mega_mega_hellhound"`
				SurvivedPlayersKitMegaMegaHellhound           int      `json:"survived_players_kit_mega_mega_hellhound"`
				TimePlayedMega                                int      `json:"time_played_mega"`
				AssistsKitBasicSoloGrenade                    int      `json:"assists_kit_basic_solo_grenade"`
				MegaLuckyCharm                                int      `json:"mega_lucky_charm"`
				MegaBridger                                   int      `json:"mega_bridger"`
				DeathsKitBasicSoloFrog                        int      `json:"deaths_kit_basic_solo_frog"`
				ChestsOpenedKitBasicSoloFrog                  int      `json:"chests_opened_kit_basic_solo_frog"`
				TimePlayedKitBasicSoloFrog                    int      `json:"time_played_kit_basic_solo_frog"`
				LossesKitBasicSoloFrog                        int      `json:"losses_kit_basic_solo_frog"`
				FastestWinKitBasicSoloFrog                    int      `json:"fastest_win_kit_basic_solo_frog"`
				SurvivedPlayersKitBasicSoloFrog               int      `json:"survived_players_kit_basic_solo_frog"`
				VoidKillsKitBasicSoloFrog                     int      `json:"void_kills_kit_basic_solo_frog"`
				WinstreakKitBasicSoloFrog                     int      `json:"winstreak_kit_basic_solo_frog"`
				GamesKitBasicSoloFrog                         int      `json:"games_kit_basic_solo_frog"`
				MostKillsGameKitBasicSoloFrog                 int      `json:"most_kills_game_kit_basic_solo_frog"`
				KillstreakKitBasicSoloFrog                    int      `json:"killstreak_kit_basic_solo_frog"`
				WinsKitBasicSoloFrog                          int      `json:"wins_kit_basic_solo_frog"`
				MeleeKillsKitBasicSoloFrog                    int      `json:"melee_kills_kit_basic_solo_frog"`
				KillsKitBasicSoloFrog                         int      `json:"kills_kit_basic_solo_frog"`
				LongestBowShotKitBasicSoloFrog                int      `json:"longest_bow_shot_kit_basic_solo_frog"`
				ArrowsShotKitBasicSoloFrog                    int      `json:"arrows_shot_kit_basic_solo_frog"`
				ArrowsHitKitBasicSoloFrog                     int      `json:"arrows_hit_kit_basic_solo_frog"`
				AssistsKitBasicSoloFrog                       int      `json:"assists_kit_basic_solo_frog"`
				LongestBowShotKitBasicSoloPrincess            int      `json:"longest_bow_shot_kit_basic_solo_princess"`
				FastestWinKitBasicSoloPrincess                int      `json:"fastest_win_kit_basic_solo_princess"`
				VoidKillsKitBasicSoloPrincess                 int      `json:"void_kills_kit_basic_solo_princess"`
				KillsKitBasicSoloPrincess                     int      `json:"kills_kit_basic_solo_princess"`
				WinstreakKitBasicSoloPrincess                 int      `json:"winstreak_kit_basic_solo_princess"`
				ArrowsShotKitBasicSoloPrincess                int      `json:"arrows_shot_kit_basic_solo_princess"`
				GamesKitBasicSoloPrincess                     int      `json:"games_kit_basic_solo_princess"`
				TimePlayedKitBasicSoloPrincess                int      `json:"time_played_kit_basic_solo_princess"`
				SurvivedPlayersKitBasicSoloPrincess           int      `json:"survived_players_kit_basic_solo_princess"`
				WinsKitBasicSoloPrincess                      int      `json:"wins_kit_basic_solo_princess"`
				ChestsOpenedKitBasicSoloPrincess              int      `json:"chests_opened_kit_basic_solo_princess"`
				ArrowsHitKitBasicSoloPrincess                 int      `json:"arrows_hit_kit_basic_solo_princess"`
				KillstreakKitBasicSoloPrincess                int      `json:"killstreak_kit_basic_solo_princess"`
				MostKillsGameKitBasicSoloPrincess             int      `json:"most_kills_game_kit_basic_solo_princess"`
				MeleeKillsKitBasicSoloPrincess                int      `json:"melee_kills_kit_basic_solo_princess"`
				SurvivedPlayersKitBasicSoloDisco              int      `json:"survived_players_kit_basic_solo_disco"`
				VoidKillsKitBasicSoloDisco                    int      `json:"void_kills_kit_basic_solo_disco"`
				MostKillsGameKitBasicSoloDisco                int      `json:"most_kills_game_kit_basic_solo_disco"`
				DeathsKitBasicSoloDisco                       int      `json:"deaths_kit_basic_solo_disco"`
				KillsKitBasicSoloDisco                        int      `json:"kills_kit_basic_solo_disco"`
				LossesKitBasicSoloDisco                       int      `json:"losses_kit_basic_solo_disco"`
				AssistsKitBasicSoloDisco                      int      `json:"assists_kit_basic_solo_disco"`
				TimePlayedKitBasicSoloDisco                   int      `json:"time_played_kit_basic_solo_disco"`
				GamesKitBasicSoloDisco                        int      `json:"games_kit_basic_solo_disco"`
				FastestWinKitBasicSoloDisco                   int      `json:"fastest_win_kit_basic_solo_disco"`
				WinsKitBasicSoloDisco                         int      `json:"wins_kit_basic_solo_disco"`
				WinstreakKitBasicSoloDisco                    int      `json:"winstreak_kit_basic_solo_disco"`
				KillstreakKitBasicSoloDisco                   int      `json:"killstreak_kit_basic_solo_disco"`
				ChestsOpenedKitBasicSoloDisco                 int      `json:"chests_opened_kit_basic_solo_disco"`
				MeleeKillsKitBasicSoloDisco                   int      `json:"melee_kills_kit_basic_solo_disco"`
				LongestBowShotKitBasicSoloDisco               int      `json:"longest_bow_shot_kit_basic_solo_disco"`
				ArrowsHitKitBasicSoloDisco                    int      `json:"arrows_hit_kit_basic_solo_disco"`
				ArrowsShotKitBasicSoloDisco                   int      `json:"arrows_shot_kit_basic_solo_disco"`
				ChestsOpenedKitBasicSoloBatguy                int      `json:"chests_opened_kit_basic_solo_batguy"`
				DeathsKitBasicSoloBatguy                      int      `json:"deaths_kit_basic_solo_batguy"`
				TimePlayedKitBasicSoloBatguy                  int      `json:"time_played_kit_basic_solo_batguy"`
				LossesKitBasicSoloBatguy                      int      `json:"losses_kit_basic_solo_batguy"`
				SurvivedPlayersKitBasicSoloBatguy             int      `json:"survived_players_kit_basic_solo_batguy"`
				LongestBowShotKitBasicSoloBatguy              int      `json:"longest_bow_shot_kit_basic_solo_batguy"`
				FastestWinKitBasicSoloBatguy                  int      `json:"fastest_win_kit_basic_solo_batguy"`
				ArrowsShotKitBasicSoloBatguy                  int      `json:"arrows_shot_kit_basic_solo_batguy"`
				VoidKillsKitBasicSoloBatguy                   int      `json:"void_kills_kit_basic_solo_batguy"`
				MostKillsGameKitBasicSoloBatguy               int      `json:"most_kills_game_kit_basic_solo_batguy"`
				GamesKitBasicSoloBatguy                       int      `json:"games_kit_basic_solo_batguy"`
				KillsKitBasicSoloBatguy                       int      `json:"kills_kit_basic_solo_batguy"`
				ArrowsHitKitBasicSoloBatguy                   int      `json:"arrows_hit_kit_basic_solo_batguy"`
				MeleeKillsKitBasicSoloBatguy                  int      `json:"melee_kills_kit_basic_solo_batguy"`
				WinsKitBasicSoloBatguy                        int      `json:"wins_kit_basic_solo_batguy"`
				WinstreakKitBasicSoloBatguy                   int      `json:"winstreak_kit_basic_solo_batguy"`
				KillstreakKitBasicSoloBatguy                  int      `json:"killstreak_kit_basic_solo_batguy"`
				LossesKitBasicSoloPrincess                    int      `json:"losses_kit_basic_solo_princess"`
				DeathsKitBasicSoloPrincess                    int      `json:"deaths_kit_basic_solo_princess"`
				AssistsKitBasicSoloPrincess                   int      `json:"assists_kit_basic_solo_princess"`
				MobKills                                      int      `json:"mob_kills"`
				MobKillsKitBasicSoloPrincess                  int      `json:"mob_kills_kit_basic_solo_princess"`
				MobKillsSolo                                  int      `json:"mob_kills_solo"`
				VoidKillsKitMegaMegaArmorsmith                int      `json:"void_kills_kit_mega_mega_armorsmith"`
				ChestsOpenedKitMegaMegaArmorsmith             int      `json:"chests_opened_kit_mega_mega_armorsmith"`
				MostKillsGameKitMegaMegaArmorsmith            int      `json:"most_kills_game_kit_mega_mega_armorsmith"`
				VoidKillsMega                                 int      `json:"void_kills_mega"`
				MeleeKillsKitMegaMegaArmorsmith               int      `json:"melee_kills_kit_mega_mega_armorsmith"`
				TimePlayedKitMegaMegaArmorsmith               int      `json:"time_played_kit_mega_mega_armorsmith"`
				FastestWinKitAdvancedSoloCannoneer            int      `json:"fastest_win_kit_advanced_solo_cannoneer"`
				TimePlayedKitAdvancedSoloCannoneer            int      `json:"time_played_kit_advanced_solo_cannoneer"`
				VoidKillsKitAdvancedSoloCannoneer             int      `json:"void_kills_kit_advanced_solo_cannoneer"`
				ChestsOpenedKitAdvancedSoloCannoneer          int      `json:"chests_opened_kit_advanced_solo_cannoneer"`
				MostKillsGameKitAdvancedSoloCannoneer         int      `json:"most_kills_game_kit_advanced_solo_cannoneer"`
				FastestWinKitBasicSoloArmorsmith              int      `json:"fastest_win_kit_basic_solo_armorsmith"`
				LongestBowShotKitBasicSoloArmorsmith          int      `json:"longest_bow_shot_kit_basic_solo_armorsmith"`
				ArrowsHitKitBasicSoloArmorsmith               int      `json:"arrows_hit_kit_basic_solo_armorsmith"`
				ChestsOpenedKitBasicSoloArmorsmith            int      `json:"chests_opened_kit_basic_solo_armorsmith"`
				MostKillsGameKitBasicSoloArmorsmith           int      `json:"most_kills_game_kit_basic_solo_armorsmith"`
				ArrowsShotKitBasicSoloArmorsmith              int      `json:"arrows_shot_kit_basic_solo_armorsmith"`
				VoidKillsKitBasicSoloArmorsmith               int      `json:"void_kills_kit_basic_solo_armorsmith"`
				TimePlayedKitBasicSoloArmorsmith              int      `json:"time_played_kit_basic_solo_armorsmith"`
				MeleeKillsKitBasicSoloArmorsmith              int      `json:"melee_kills_kit_basic_solo_armorsmith"`
				ChestsOpenedKitAdvancedSoloPyro               int      `json:"chests_opened_kit_advanced_solo_pyro"`
				MostKillsGameKitAdvancedSoloPyro              int      `json:"most_kills_game_kit_advanced_solo_pyro"`
				MeleeKillsKitAdvancedSoloPyro                 int      `json:"melee_kills_kit_advanced_solo_pyro"`
				TimePlayedKitAdvancedSoloPyro                 int      `json:"time_played_kit_advanced_solo_pyro"`
				LongestBowShotKitAdvancedSoloPyro             int      `json:"longest_bow_shot_kit_advanced_solo_pyro"`
				FastestWinKitAdvancedSoloPyro                 int      `json:"fastest_win_kit_advanced_solo_pyro"`
				ArrowsShotKitAdvancedSoloPyro                 int      `json:"arrows_shot_kit_advanced_solo_pyro"`
				VoidKillsKitAdvancedSoloPyro                  int      `json:"void_kills_kit_advanced_solo_pyro"`
				ArrowsHitKitAdvancedSoloPyro                  int      `json:"arrows_hit_kit_advanced_solo_pyro"`
				VotesFestiveTribute                           int      `json:"votes_Festive Tribute"`
				TimePlayedKitEnderchestSoloEnderchest         int      `json:"time_played_kit_enderchest_solo_enderchest"`
				ChestsOpenedKitEnderchestSoloEnderchest       int      `json:"chests_opened_kit_enderchest_solo_enderchest"`
				FastestWinTeam                                int      `json:"fastest_win_team"`
				FastestWinKitDefendingTeamArmorer             int      `json:"fastest_win_kit_defending_team_armorer"`
				MeleeKillsKitDefendingTeamArmorer             int      `json:"melee_kills_kit_defending_team_armorer"`
				MeleeKillsTeam                                int      `json:"melee_kills_team"`
				TimePlayedKitDefendingTeamArmorer             int      `json:"time_played_kit_defending_team_armorer"`
				MostKillsGameTeam                             int      `json:"most_kills_game_team"`
				ChestsOpenedTeam                              int      `json:"chests_opened_team"`
				TimePlayedTeam                                int      `json:"time_played_team"`
				MostKillsGameKitDefendingTeamArmorer          int      `json:"most_kills_game_kit_defending_team_armorer"`
				ChestsOpenedKitDefendingTeamArmorer           int      `json:"chests_opened_kit_defending_team_armorer"`
				VoidKillsTeam                                 int      `json:"void_kills_team"`
				VoidKillsKitDefendingTeamArmorer              int      `json:"void_kills_kit_defending_team_armorer"`
				LongestBowShotTeam                            int      `json:"longest_bow_shot_team"`
				LongestBowShotKitDefendingTeamArmorer         int      `json:"longest_bow_shot_kit_defending_team_armorer"`
				ArrowsShotTeam                                int      `json:"arrows_shot_team"`
				ArrowsHitTeam                                 int      `json:"arrows_hit_team"`
				ArrowsHitKitDefendingTeamArmorer              int      `json:"arrows_hit_kit_defending_team_armorer"`
				ArrowsShotKitDefendingTeamArmorer             int      `json:"arrows_shot_kit_defending_team_armorer"`
				FastestWinKitBasicSoloSnowman                 int      `json:"fastest_win_kit_basic_solo_snowman"`
				VoidKillsKitBasicSoloSnowman                  int      `json:"void_kills_kit_basic_solo_snowman"`
				MeleeKillsKitBasicSoloSnowman                 int      `json:"melee_kills_kit_basic_solo_snowman"`
				ChestsOpenedKitBasicSoloSnowman               int      `json:"chests_opened_kit_basic_solo_snowman"`
				MostKillsGameKitBasicSoloSnowman              int      `json:"most_kills_game_kit_basic_solo_snowman"`
				TimePlayedKitBasicSoloSnowman                 int      `json:"time_played_kit_basic_solo_snowman"`
				FastestWinKitAdvancedSoloEnderman             int      `json:"fastest_win_kit_advanced_solo_enderman"`
				MostKillsGameKitAdvancedSoloEnderman          int      `json:"most_kills_game_kit_advanced_solo_enderman"`
				TimePlayedKitAdvancedSoloEnderman             int      `json:"time_played_kit_advanced_solo_enderman"`
				VoidKillsKitAdvancedSoloEnderman              int      `json:"void_kills_kit_advanced_solo_enderman"`
				ChestsOpenedKitAdvancedSoloEnderman           int      `json:"chests_opened_kit_advanced_solo_enderman"`
				MeleeKillsKitAdvancedSoloEnderman             int      `json:"melee_kills_kit_advanced_solo_enderman"`
				SurvivedPlayersKitBasicSoloScout              int      `json:"survived_players_kit_basic_solo_scout"`
				LossesKitBasicSoloScout                       int      `json:"losses_kit_basic_solo_scout"`
				DeathsKitBasicSoloScout                       int      `json:"deaths_kit_basic_solo_scout"`
				TimePlayedKitBasicSoloScout                   int      `json:"time_played_kit_basic_solo_scout"`
				ChestsOpenedKitBasicSoloScout                 int      `json:"chests_opened_kit_basic_solo_scout"`
				AssistsKitAdvancedSoloEnderman                int      `json:"assists_kit_advanced_solo_enderman"`
				MeleeKillsKitAdvancedSoloCannoneer            int      `json:"melee_kills_kit_advanced_solo_cannoneer"`
				LongestBowShotKitAdvancedSoloCannoneer        int      `json:"longest_bow_shot_kit_advanced_solo_cannoneer"`
				ArrowsShotKitAdvancedSoloCannoneer            int      `json:"arrows_shot_kit_advanced_solo_cannoneer"`
				ArrowsHitKitAdvancedSoloCannoneer             int      `json:"arrows_hit_kit_advanced_solo_cannoneer"`
				FastestWinKitRankedRankedPyromancer           int      `json:"fastest_win_kit_ranked_ranked_pyromancer"`
				FastestWinRanked                              int      `json:"fastest_win_ranked"`
				ChestsOpenedRanked                            int      `json:"chests_opened_ranked"`
				TimePlayedRanked                              int      `json:"time_played_ranked"`
				MostKillsGameKitRankedRankedPyromancer        int      `json:"most_kills_game_kit_ranked_ranked_pyromancer"`
				VoidKillsRanked                               int      `json:"void_kills_ranked"`
				TimePlayedKitRankedRankedPyromancer           int      `json:"time_played_kit_ranked_ranked_pyromancer"`
				MostKillsGameRanked                           int      `json:"most_kills_game_ranked"`
				ChestsOpenedKitRankedRankedPyromancer         int      `json:"chests_opened_kit_ranked_ranked_pyromancer"`
				VoidKillsKitRankedRankedPyromancer            int      `json:"void_kills_kit_ranked_ranked_pyromancer"`
				LongestBowShotKitBasicSoloEnergix             int      `json:"longest_bow_shot_kit_basic_solo_energix"`
				FastestWinKitBasicSoloEnergix                 int      `json:"fastest_win_kit_basic_solo_energix"`
				KillsKitBasicSoloEnergix                      int      `json:"kills_kit_basic_solo_energix"`
				VoidKillsKitBasicSoloEnergix                  int      `json:"void_kills_kit_basic_solo_energix"`
				MostKillsGameKitBasicSoloEnergix              int      `json:"most_kills_game_kit_basic_solo_energix"`
				ArrowsShotKitBasicSoloEnergix                 int      `json:"arrows_shot_kit_basic_solo_energix"`
				ArrowsHitKitBasicSoloEnergix                  int      `json:"arrows_hit_kit_basic_solo_energix"`
				ChestsOpenedKitBasicSoloEnergix               int      `json:"chests_opened_kit_basic_solo_energix"`
				KillstreakKitBasicSoloEnergix                 int      `json:"killstreak_kit_basic_solo_energix"`
				TimePlayedKitBasicSoloEnergix                 int      `json:"time_played_kit_basic_solo_energix"`
				SurvivedPlayersKitBasicSoloEnergix            int      `json:"survived_players_kit_basic_solo_energix"`
				MeleeKillsKitBasicSoloEnergix                 int      `json:"melee_kills_kit_basic_solo_energix"`
				WinsKitBasicSoloEnergix                       int      `json:"wins_kit_basic_solo_energix"`
				GamesKitBasicSoloEnergix                      int      `json:"games_kit_basic_solo_energix"`
				WinstreakKitBasicSoloEnergix                  int      `json:"winstreak_kit_basic_solo_energix"`
				LossesKitBasicSoloEnergix                     int      `json:"losses_kit_basic_solo_energix"`
				DeathsKitBasicSoloEnergix                     int      `json:"deaths_kit_basic_solo_energix"`
				FastestWinKitRankedRankedArmorer              int      `json:"fastest_win_kit_ranked_ranked_armorer"`
				MeleeKillsKitRankedRankedArmorer              int      `json:"melee_kills_kit_ranked_ranked_armorer"`
				TimePlayedKitRankedRankedArmorer              int      `json:"time_played_kit_ranked_ranked_armorer"`
				MeleeKillsRanked                              int      `json:"melee_kills_ranked"`
				VoidKillsKitRankedRankedArmorer               int      `json:"void_kills_kit_ranked_ranked_armorer"`
				ChestsOpenedKitRankedRankedArmorer            int      `json:"chests_opened_kit_ranked_ranked_armorer"`
				MostKillsGameKitRankedRankedArmorer           int      `json:"most_kills_game_kit_ranked_ranked_armorer"`
				LongestBowShotRanked                          int      `json:"longest_bow_shot_ranked"`
				LongestBowShotKitRankedRankedArmorer          int      `json:"longest_bow_shot_kit_ranked_ranked_armorer"`
				ArrowsHitRanked                               int      `json:"arrows_hit_ranked"`
				ArrowsHitKitRankedRankedArmorer               int      `json:"arrows_hit_kit_ranked_ranked_armorer"`
				ArrowsShotRanked                              int      `json:"arrows_shot_ranked"`
				ArrowsShotKitRankedRankedArmorer              int      `json:"arrows_shot_kit_ranked_ranked_armorer"`
				AssistsKitBasicSoloEnergix                    int      `json:"assists_kit_basic_solo_energix"`
				ArrowsShotKitAdvancedSoloEnderman             int      `json:"arrows_shot_kit_advanced_solo_enderman"`
				LongestBowShotKitMegaMegaArmorsmith           int      `json:"longest_bow_shot_kit_mega_mega_armorsmith"`
				LongestBowShotMega                            int      `json:"longest_bow_shot_mega"`
				ArrowsShotKitMegaMegaArmorsmith               int      `json:"arrows_shot_kit_mega_mega_armorsmith"`
				ArrowsHitKitMegaMegaArmorsmith                int      `json:"arrows_hit_kit_mega_mega_armorsmith"`
				ArrowsShotMega                                int      `json:"arrows_shot_mega"`
				ArrowsHitMega                                 int      `json:"arrows_hit_mega"`
				LongestBowShotKitAdvancedSoloEnderman         int      `json:"longest_bow_shot_kit_advanced_solo_enderman"`
				ArrowsHitKitAdvancedSoloEnderman              int      `json:"arrows_hit_kit_advanced_solo_enderman"`
				MostKillsGameKitBasicSoloRookie               int      `json:"most_kills_game_kit_basic_solo_rookie"`
				SurvivedPlayersKitBasicSoloRookie             int      `json:"survived_players_kit_basic_solo_rookie"`
				GamesKitBasicSoloRookie                       int      `json:"games_kit_basic_solo_rookie"`
				KillsKitBasicSoloRookie                       int      `json:"kills_kit_basic_solo_rookie"`
				TimePlayedKitBasicSoloRookie                  int      `json:"time_played_kit_basic_solo_rookie"`
				LossesKitBasicSoloRookie                      int      `json:"losses_kit_basic_solo_rookie"`
				ChestsOpenedKitBasicSoloRookie                int      `json:"chests_opened_kit_basic_solo_rookie"`
				DeathsKitBasicSoloRookie                      int      `json:"deaths_kit_basic_solo_rookie"`
				MeleeKillsKitBasicSoloRookie                  int      `json:"melee_kills_kit_basic_solo_rookie"`
				MobKillsKitBasicSoloEnergix                   int      `json:"mob_kills_kit_basic_solo_energix"`
				GamesPlayedSkywars                            int      `json:"games_played_skywars"`
				LastMode                                      string   `json:"lastMode"`
				MeleeKillsKitMegaMegaKnight                   int      `json:"melee_kills_kit_mega_mega_knight"`
				ChestsOpenedKitMegaMegaKnight                 int      `json:"chests_opened_kit_mega_mega_knight"`
				MostKillsGameKitMegaMegaKnight                int      `json:"most_kills_game_kit_mega_mega_knight"`
				TimePlayedKitMegaMegaKnight                   int      `json:"time_played_kit_mega_mega_knight"`
				VoidKillsKitMegaMegaKnight                    int      `json:"void_kills_kit_mega_mega_knight"`
				LongestBowShotKitMegaMegaKnight               int      `json:"longest_bow_shot_kit_mega_mega_knight"`
				ArrowsShotKitMegaMegaKnight                   int      `json:"arrows_shot_kit_mega_mega_knight"`
				ArrowsHitKitMegaMegaKnight                    int      `json:"arrows_hit_kit_mega_mega_knight"`
				LongestBowShotKitAttackingTeamEnderman        int      `json:"longest_bow_shot_kit_attacking_team_enderman"`
				KillsKitAttackingTeamEnderman                 int      `json:"kills_kit_attacking_team_enderman"`
				LossesKitAttackingTeamEnderman                int      `json:"losses_kit_attacking_team_enderman"`
				MeleeKillsKitAttackingTeamEnderman            int      `json:"melee_kills_kit_attacking_team_enderman"`
				SurvivedPlayersKitAttackingTeamEnderman       int      `json:"survived_players_kit_attacking_team_enderman"`
				ArrowsShotKitAttackingTeamEnderman            int      `json:"arrows_shot_kit_attacking_team_enderman"`
				DeathsKitAttackingTeamEnderman                int      `json:"deaths_kit_attacking_team_enderman"`
				TimePlayedKitAttackingTeamEnderman            int      `json:"time_played_kit_attacking_team_enderman"`
				ArrowsHitKitAttackingTeamEnderman             int      `json:"arrows_hit_kit_attacking_team_enderman"`
				ChestsOpenedKitAttackingTeamEnderman          int      `json:"chests_opened_kit_attacking_team_enderman"`
				MostKillsGameKitAttackingTeamEnderman         int      `json:"most_kills_game_kit_attacking_team_enderman"`
				FastestWinKitMegaMegaArmorsmith               int      `json:"fastest_win_kit_mega_mega_armorsmith"`
				FastestWinMega                                int      `json:"fastest_win_mega"`
				KillstreakKitMegaMegaArmorsmith               int      `json:"killstreak_kit_mega_mega_armorsmith"`
				MobKillsMega                                  int      `json:"mob_kills_mega"`
				MobKillsKitMegaMegaArmorsmith                 int      `json:"mob_kills_kit_mega_mega_armorsmith"`
				WinstreakKitMegaMegaArmorsmith                int      `json:"winstreak_kit_mega_mega_armorsmith"`
				KitBasicSoloEnergixInventory                  struct {
					POTION9 string `json:"POTION:9"`
				} `json:"kit_basic_solo_energix_inventory"`
				KitAdvancedSoloEndermanInventory struct {
					ENDERPEARL0 string `json:"ENDER_PEARL:0"`
				} `json:"kit_advanced_solo_enderman_inventory"`
				KitBasicSoloArmorsmithInventory struct {
					ENCHANTEDBOOK0 string `json:"ENCHANTED_BOOK:0"`
					EXPBOTTLE0     string `json:"EXP_BOTTLE:0"`
					ANVIL0         string `json:"ANVIL:0"`
					DIAMONDHELMET0 string `json:"DIAMOND_HELMET:0"`
				} `json:"kit_basic_solo_armorsmith_inventory"`
				KitBasicSoloPrincessInventory struct {
					ARROW0      string `json:"ARROW:0"`
					GOLDHELMET0 string `json:"GOLD_HELMET:0"`
					BOW114      string `json:"BOW:114"`
				} `json:"kit_basic_solo_princess_inventory"`
				MobKillsKitBasicSoloGrenade            int `json:"mob_kills_kit_basic_solo_grenade"`
				LongestBowShotKitBasicSoloSnowman      int `json:"longest_bow_shot_kit_basic_solo_snowman"`
				ArrowsShotKitBasicSoloSnowman          int `json:"arrows_shot_kit_basic_solo_snowman"`
				ArrowsHitKitBasicSoloSnowman           int `json:"arrows_hit_kit_basic_solo_snowman"`
				AssistsKitBasicSoloSnowman             int `json:"assists_kit_basic_solo_snowman"`
				MostKillsGameKitAdvancedSoloKnight     int `json:"most_kills_game_kit_advanced_solo_knight"`
				SurvivedPlayersKitAdvancedSoloKnight   int `json:"survived_players_kit_advanced_solo_knight"`
				TimePlayedKitAdvancedSoloKnight        int `json:"time_played_kit_advanced_solo_knight"`
				KillsKitAdvancedSoloKnight             int `json:"kills_kit_advanced_solo_knight"`
				DeathsKitAdvancedSoloKnight            int `json:"deaths_kit_advanced_solo_knight"`
				MeleeKillsKitAdvancedSoloKnight        int `json:"melee_kills_kit_advanced_solo_knight"`
				LossesKitAdvancedSoloKnight            int `json:"losses_kit_advanced_solo_knight"`
				FastestWinKitAdvancedSoloKnight        int `json:"fastest_win_kit_advanced_solo_knight"`
				VoidKillsKitAdvancedSoloKnight         int `json:"void_kills_kit_advanced_solo_knight"`
				GamesKitAdvancedSoloKnight             int `json:"games_kit_advanced_solo_knight"`
				WinstreakKitAdvancedSoloKnight         int `json:"winstreak_kit_advanced_solo_knight"`
				KillstreakKitAdvancedSoloKnight        int `json:"killstreak_kit_advanced_solo_knight"`
				WinsKitAdvancedSoloKnight              int `json:"wins_kit_advanced_solo_knight"`
				ChestsOpenedKitAdvancedSoloKnight      int `json:"chests_opened_kit_advanced_solo_knight"`
				LongestBowShotKitBasicSoloSpeleologist int `json:"longest_bow_shot_kit_basic_solo_speleologist"`
				MeleeKillsKitBasicSoloSpeleologist     int `json:"melee_kills_kit_basic_solo_speleologist"`
				ArrowsHitKitBasicSoloSpeleologist      int `json:"arrows_hit_kit_basic_solo_speleologist"`
				DeathsKitBasicSoloSpeleologist         int `json:"deaths_kit_basic_solo_speleologist"`
				ArrowsShotKitBasicSoloSpeleologist     int `json:"arrows_shot_kit_basic_solo_speleologist"`
				MostKillsGameKitBasicSoloSpeleologist  int `json:"most_kills_game_kit_basic_solo_speleologist"`
				TimePlayedKitBasicSoloSpeleologist     int `json:"time_played_kit_basic_solo_speleologist"`
				ChestsOpenedKitBasicSoloSpeleologist   int `json:"chests_opened_kit_basic_solo_speleologist"`
				LossesKitBasicSoloSpeleologist         int `json:"losses_kit_basic_solo_speleologist"`
				VoidKillsKitBasicSoloSpeleologist      int `json:"void_kills_kit_basic_solo_speleologist"`
				FastestWinKitBasicSoloSpeleologist     int `json:"fastest_win_kit_basic_solo_speleologist"`
				ChestsOpenedKitAdvancedSoloArmorer     int `json:"chests_opened_kit_advanced_solo_armorer"`
				MostKillsGameKitAdvancedSoloArmorer    int `json:"most_kills_game_kit_advanced_solo_armorer"`
				MeleeKillsKitAdvancedSoloArmorer       int `json:"melee_kills_kit_advanced_solo_armorer"`
				TimePlayedKitAdvancedSoloArmorer       int `json:"time_played_kit_advanced_solo_armorer"`
				FastestWinKitAdvancedSoloArmorer       int `json:"fastest_win_kit_advanced_solo_armorer"`
				LongestBowShotKitAdvancedSoloArmorer   int `json:"longest_bow_shot_kit_advanced_solo_armorer"`
				ArrowsHitKitAdvancedSoloArmorer        int `json:"arrows_hit_kit_advanced_solo_armorer"`
				ArrowsShotKitAdvancedSoloArmorer       int `json:"arrows_shot_kit_advanced_solo_armorer"`
				VoidKillsKitAdvancedSoloArmorer        int `json:"void_kills_kit_advanced_solo_armorer"`
				FastestWinKitRankedRankedScout         int `json:"fastest_win_kit_ranked_ranked_scout"`
				GamesKitRankedRankedScout              int `json:"games_kit_ranked_ranked_scout"`
				SurvivedPlayersKitRankedRankedScout    int `json:"survived_players_kit_ranked_ranked_scout"`
				TimePlayedKitRankedRankedScout         int `json:"time_played_kit_ranked_ranked_scout"`
				MostKillsGameKitRankedRankedScout      int `json:"most_kills_game_kit_ranked_ranked_scout"`
				ChestsOpenedKitRankedRankedScout       int `json:"chests_opened_kit_ranked_ranked_scout"`
				MeleeKillsKitRankedRankedScout         int `json:"melee_kills_kit_ranked_ranked_scout"`
				ArrowsShotKitRankedRankedScout         int `json:"arrows_shot_kit_ranked_ranked_scout"`
				WinstreakKitRankedRankedScout          int `json:"winstreak_kit_ranked_ranked_scout"`
				KillstreakKitRankedRankedScout         int `json:"killstreak_kit_ranked_ranked_scout"`
				KillsKitRankedRankedScout              int `json:"kills_kit_ranked_ranked_scout"`
				WinsKitRankedRankedScout               int `json:"wins_kit_ranked_ranked_scout"`
				VoidKillsKitRankedRankedScout          int `json:"void_kills_kit_ranked_ranked_scout"`
				LongestBowShotKitRankedRankedScout     int `json:"longest_bow_shot_kit_ranked_ranked_scout"`
				ArrowsHitKitRankedRankedScout          int `json:"arrows_hit_kit_ranked_ranked_scout"`
				DeathsKitRankedRankedScout             int `json:"deaths_kit_ranked_ranked_scout"`
				LossesKitRankedRankedScout             int `json:"losses_kit_ranked_ranked_scout"`
				AssistsKitRankedRankedScout            int `json:"assists_kit_ranked_ranked_scout"`
				KitRankedRankedScoutInventory          struct {
					DIAMONDAXE0     string `json:"DIAMOND_AXE:0"`
					POTION2         string `json:"POTION:2"`
					IRONLEGGINGS0   string `json:"IRON_LEGGINGS:0"`
					IRONHELMET0     string `json:"IRON_HELMET:0"`
					IRONBOOTS0      string `json:"IRON_BOOTS:0"`
					IRONCHESTPLATE0 string `json:"IRON_CHESTPLATE:0"`
					DIAMONDPICKAXE0 string `json:"DIAMOND_PICKAXE:0"`
				} `json:"kit_ranked_ranked_scout_inventory"`
				FastestWinKitRankedRankedChampion            int `json:"fastest_win_kit_ranked_ranked_champion"`
				ChestsOpenedKitRankedRankedChampion          int `json:"chests_opened_kit_ranked_ranked_champion"`
				VoidKillsKitRankedRankedChampion             int `json:"void_kills_kit_ranked_ranked_champion"`
				MostKillsGameKitRankedRankedChampion         int `json:"most_kills_game_kit_ranked_ranked_champion"`
				TimePlayedKitRankedRankedChampion            int `json:"time_played_kit_ranked_ranked_champion"`
				LongestBowShotKitRankedRankedChampion        int `json:"longest_bow_shot_kit_ranked_ranked_champion"`
				ArrowsShotKitRankedRankedChampion            int `json:"arrows_shot_kit_ranked_ranked_champion"`
				ArrowsHitKitRankedRankedChampion             int `json:"arrows_hit_kit_ranked_ranked_champion"`
				MeleeKillsKitRankedRankedChampion            int `json:"melee_kills_kit_ranked_ranked_champion"`
				FastestWinKitBlacksmithRankedBlacksmith      int `json:"fastest_win_kit_blacksmith_ranked_blacksmith"`
				MostKillsGameKitBlacksmithRankedBlacksmith   int `json:"most_kills_game_kit_blacksmith_ranked_blacksmith"`
				GamesKitBlacksmithRankedBlacksmith           int `json:"games_kit_blacksmith_ranked_blacksmith"`
				MeleeKillsKitBlacksmithRankedBlacksmith      int `json:"melee_kills_kit_blacksmith_ranked_blacksmith"`
				KillstreakKitBlacksmithRankedBlacksmith      int `json:"killstreak_kit_blacksmith_ranked_blacksmith"`
				WinsKitBlacksmithRankedBlacksmith            int `json:"wins_kit_blacksmith_ranked_blacksmith"`
				SurvivedPlayersKitBlacksmithRankedBlacksmith int `json:"survived_players_kit_blacksmith_ranked_blacksmith"`
				TimePlayedKitBlacksmithRankedBlacksmith      int `json:"time_played_kit_blacksmith_ranked_blacksmith"`
				WinstreakKitBlacksmithRankedBlacksmith       int `json:"winstreak_kit_blacksmith_ranked_blacksmith"`
				KillsKitBlacksmithRankedBlacksmith           int `json:"kills_kit_blacksmith_ranked_blacksmith"`
				ChestsOpenedKitBlacksmithRankedBlacksmith    int `json:"chests_opened_kit_blacksmith_ranked_blacksmith"`
				LossesKitBlacksmithRankedBlacksmith          int `json:"losses_kit_blacksmith_ranked_blacksmith"`
				DeathsKitBlacksmithRankedBlacksmith          int `json:"deaths_kit_blacksmith_ranked_blacksmith"`
				LongestBowShotKitBlacksmithRankedBlacksmith  int `json:"longest_bow_shot_kit_blacksmith_ranked_blacksmith"`
				VoidKillsKitBlacksmithRankedBlacksmith       int `json:"void_kills_kit_blacksmith_ranked_blacksmith"`
				ArrowsShotKitBlacksmithRankedBlacksmith      int `json:"arrows_shot_kit_blacksmith_ranked_blacksmith"`
				ArrowsHitKitBlacksmithRankedBlacksmith       int `json:"arrows_hit_kit_blacksmith_ranked_blacksmith"`
				TimePlayedKitBasicSoloTroll                  int `json:"time_played_kit_basic_solo_troll"`
				ChestsOpenedKitBasicSoloTroll                int `json:"chests_opened_kit_basic_solo_troll"`
				TimePlayedKitBasicSoloPharaoh                int `json:"time_played_kit_basic_solo_pharaoh"`
				ChestsOpenedKitBasicSoloPharaoh              int `json:"chests_opened_kit_basic_solo_pharaoh"`
				MostKillsGameKitBasicSoloPharaoh             int `json:"most_kills_game_kit_basic_solo_pharaoh"`
				VoidKillsKitBasicSoloPharaoh                 int `json:"void_kills_kit_basic_solo_pharaoh"`
				MeleeKillsKitBasicSoloPharaoh                int `json:"melee_kills_kit_basic_solo_pharaoh"`
				FastestWinKitBasicSoloPharaoh                int `json:"fastest_win_kit_basic_solo_pharaoh"`
				KitBasicSoloScoutInventory                   struct {
					POTION2 string `json:"POTION:2"`
				} `json:"kit_basic_solo_scout_inventory"`
				KitAttackingTeamScoutInventory struct {
					POTION2     string `json:"POTION:2"`
					STONESWORD0 string `json:"STONE_SWORD:0"`
				} `json:"kit_attacking_team_scout_inventory"`
				ChestsOpenedKitAdvancedSoloFarmer   int `json:"chests_opened_kit_advanced_solo_farmer"`
				MostKillsGameKitAdvancedSoloFarmer  int `json:"most_kills_game_kit_advanced_solo_farmer"`
				TimePlayedKitAdvancedSoloFarmer     int `json:"time_played_kit_advanced_solo_farmer"`
				VoidKillsKitAdvancedSoloFarmer      int `json:"void_kills_kit_advanced_solo_farmer"`
				FastestWinKitAdvancedSoloFarmer     int `json:"fastest_win_kit_advanced_solo_farmer"`
				ArrowsShotKitAdvancedSoloFarmer     int `json:"arrows_shot_kit_advanced_solo_farmer"`
				MeleeKillsKitAdvancedSoloFarmer     int `json:"melee_kills_kit_advanced_solo_farmer"`
				LongestBowShotKitAdvancedSoloFarmer int `json:"longest_bow_shot_kit_advanced_solo_farmer"`
				ArrowsHitKitAdvancedSoloFarmer      int `json:"arrows_hit_kit_advanced_solo_farmer"`
				KitAdvancedSoloFarmerInventory      struct {
					IRONLEGGINGS0 string `json:"IRON_LEGGINGS:0"`
					EGG0          string `json:"EGG:0"`
					GOLDENAPPLE0  string `json:"GOLDEN_APPLE:0"`
				} `json:"kit_advanced_solo_farmer_inventory"`
				VoidKillsKitBasicSoloRookie                 int      `json:"void_kills_kit_basic_solo_rookie"`
				FastestWinKitBasicSoloRookie                int      `json:"fastest_win_kit_basic_solo_rookie"`
				WinsKitBasicSoloRookie                      int      `json:"wins_kit_basic_solo_rookie"`
				WinstreakKitBasicSoloRookie                 int      `json:"winstreak_kit_basic_solo_rookie"`
				KillstreakKitBasicSoloRookie                int      `json:"killstreak_kit_basic_solo_rookie"`
				LongestBowShotKitBasicSoloRookie            int      `json:"longest_bow_shot_kit_basic_solo_rookie"`
				ArrowsShotKitBasicSoloRookie                int      `json:"arrows_shot_kit_basic_solo_rookie"`
				ArrowsHitKitBasicSoloRookie                 int      `json:"arrows_hit_kit_basic_solo_rookie"`
				AssistsKitBasicSoloRookie                   int      `json:"assists_kit_basic_solo_rookie"`
				FastestWinKitAttackingTeamSnowman           int      `json:"fastest_win_kit_attacking_team_snowman"`
				VoidKillsKitAttackingTeamSnowman            int      `json:"void_kills_kit_attacking_team_snowman"`
				MeleeKillsKitAttackingTeamSnowman           int      `json:"melee_kills_kit_attacking_team_snowman"`
				MostKillsGameKitAttackingTeamSnowman        int      `json:"most_kills_game_kit_attacking_team_snowman"`
				WinstreakKitAttackingTeamSnowman            int      `json:"winstreak_kit_attacking_team_snowman"`
				KillstreakKitAttackingTeamSnowman           int      `json:"killstreak_kit_attacking_team_snowman"`
				TimePlayedKitAttackingTeamSnowman           int      `json:"time_played_kit_attacking_team_snowman"`
				ChestsOpenedKitAttackingTeamSnowman         int      `json:"chests_opened_kit_attacking_team_snowman"`
				TeamLuckyCharm                              int      `json:"team_lucky_charm"`
				TeamEnvironmentalExpert                     int      `json:"team_environmental_expert"`
				TeamBridger                                 int      `json:"team_bridger"`
				FastestWinKitSupportingTeamArmorsmith       int      `json:"fastest_win_kit_supporting_team_armorsmith"`
				KillstreakKitSupportingTeamArmorsmith       int      `json:"killstreak_kit_supporting_team_armorsmith"`
				WinstreakKitSupportingTeamArmorsmith        int      `json:"winstreak_kit_supporting_team_armorsmith"`
				MostKillsGameKitSupportingTeamArmorsmith    int      `json:"most_kills_game_kit_supporting_team_armorsmith"`
				ChestsOpenedKitSupportingTeamArmorsmith     int      `json:"chests_opened_kit_supporting_team_armorsmith"`
				TimePlayedKitSupportingTeamArmorsmith       int      `json:"time_played_kit_supporting_team_armorsmith"`
				MeleeKillsKitSupportingTeamArmorsmith       int      `json:"melee_kills_kit_supporting_team_armorsmith"`
				LongestBowShotKitSupportingTeamArmorsmith   int      `json:"longest_bow_shot_kit_supporting_team_armorsmith"`
				VoidKillsKitSupportingTeamArmorsmith        int      `json:"void_kills_kit_supporting_team_armorsmith"`
				ArrowsHitKitSupportingTeamArmorsmith        int      `json:"arrows_hit_kit_supporting_team_armorsmith"`
				ArrowsShotKitSupportingTeamArmorsmith       int      `json:"arrows_shot_kit_supporting_team_armorsmith"`
				MobKillsKitDefendingTeamArmorer             int      `json:"mob_kills_kit_defending_team_armorer"`
				MobKillsTeam                                int      `json:"mob_kills_team"`
				FastestWinKitAttackingTeamKnight            int      `json:"fastest_win_kit_attacking_team_knight"`
				TimePlayedKitAttackingTeamKnight            int      `json:"time_played_kit_attacking_team_knight"`
				VoidKillsKitAttackingTeamKnight             int      `json:"void_kills_kit_attacking_team_knight"`
				MeleeKillsKitAttackingTeamKnight            int      `json:"melee_kills_kit_attacking_team_knight"`
				KillstreakKitAttackingTeamKnight            int      `json:"killstreak_kit_attacking_team_knight"`
				MostKillsGameKitAttackingTeamKnight         int      `json:"most_kills_game_kit_attacking_team_knight"`
				WinstreakKitAttackingTeamKnight             int      `json:"winstreak_kit_attacking_team_knight"`
				ArrowsShotKitAttackingTeamKnight            int      `json:"arrows_shot_kit_attacking_team_knight"`
				ChestsOpenedKitAttackingTeamKnight          int      `json:"chests_opened_kit_attacking_team_knight"`
				FastestWinKitMiningTeamCannoneer            int      `json:"fastest_win_kit_mining_team_cannoneer"`
				MeleeKillsKitMiningTeamCannoneer            int      `json:"melee_kills_kit_mining_team_cannoneer"`
				MostKillsGameKitMiningTeamCannoneer         int      `json:"most_kills_game_kit_mining_team_cannoneer"`
				ChestsOpenedKitMiningTeamCannoneer          int      `json:"chests_opened_kit_mining_team_cannoneer"`
				VoidKillsKitMiningTeamCannoneer             int      `json:"void_kills_kit_mining_team_cannoneer"`
				TimePlayedKitMiningTeamCannoneer            int      `json:"time_played_kit_mining_team_cannoneer"`
				DeathsKitAttackingTeamEnergix               int      `json:"deaths_kit_attacking_team_energix"`
				LossesKitAttackingTeamEnergix               int      `json:"losses_kit_attacking_team_energix"`
				SurvivedPlayersKitAttackingTeamEnergix      int      `json:"survived_players_kit_attacking_team_energix"`
				ChestsOpenedKitAttackingTeamEnergix         int      `json:"chests_opened_kit_attacking_team_energix"`
				TimePlayedKitAttackingTeamEnergix           int      `json:"time_played_kit_attacking_team_energix"`
				FastestWinKitAdvancedSoloHunter             int      `json:"fastest_win_kit_advanced_solo_hunter"`
				LongestBowShotKitAdvancedSoloHunter         int      `json:"longest_bow_shot_kit_advanced_solo_hunter"`
				TimePlayedKitAdvancedSoloHunter             int      `json:"time_played_kit_advanced_solo_hunter"`
				ChestsOpenedKitAdvancedSoloHunter           int      `json:"chests_opened_kit_advanced_solo_hunter"`
				MeleeKillsKitAdvancedSoloHunter             int      `json:"melee_kills_kit_advanced_solo_hunter"`
				AssistsKitAdvancedSoloHunter                int      `json:"assists_kit_advanced_solo_hunter"`
				VoidKillsKitAdvancedSoloHunter              int      `json:"void_kills_kit_advanced_solo_hunter"`
				ArrowsHitKitAdvancedSoloHunter              int      `json:"arrows_hit_kit_advanced_solo_hunter"`
				MostKillsGameKitAdvancedSoloHunter          int      `json:"most_kills_game_kit_advanced_solo_hunter"`
				ArrowsShotKitAdvancedSoloHunter             int      `json:"arrows_shot_kit_advanced_solo_hunter"`
				FastestWinKitBasicSoloEcologist             int      `json:"fastest_win_kit_basic_solo_ecologist"`
				VoidKillsKitBasicSoloEcologist              int      `json:"void_kills_kit_basic_solo_ecologist"`
				MostKillsGameKitBasicSoloEcologist          int      `json:"most_kills_game_kit_basic_solo_ecologist"`
				TimePlayedKitBasicSoloEcologist             int      `json:"time_played_kit_basic_solo_ecologist"`
				ChestsOpenedKitBasicSoloEcologist           int      `json:"chests_opened_kit_basic_solo_ecologist"`
				WinstreakKitBasicSoloEcologist              int      `json:"winstreak_kit_basic_solo_ecologist"`
				KillstreakKitBasicSoloEcologist             int      `json:"killstreak_kit_basic_solo_ecologist"`
				LongestBowShotKitAttackingTeamGrenade       int      `json:"longest_bow_shot_kit_attacking_team_grenade"`
				FastestWinKitAttackingTeamGrenade           int      `json:"fastest_win_kit_attacking_team_grenade"`
				KillsKitAttackingTeamGrenade                int      `json:"kills_kit_attacking_team_grenade"`
				TimePlayedKitAttackingTeamGrenade           int      `json:"time_played_kit_attacking_team_grenade"`
				WinstreakKitAttackingTeamGrenade            int      `json:"winstreak_kit_attacking_team_grenade"`
				ArrowsShotKitAttackingTeamGrenade           int      `json:"arrows_shot_kit_attacking_team_grenade"`
				GamesKitAttackingTeamGrenade                int      `json:"games_kit_attacking_team_grenade"`
				ArrowsHitKitAttackingTeamGrenade            int      `json:"arrows_hit_kit_attacking_team_grenade"`
				ChestsOpenedKitAttackingTeamGrenade         int      `json:"chests_opened_kit_attacking_team_grenade"`
				AssistsKitAttackingTeamGrenade              int      `json:"assists_kit_attacking_team_grenade"`
				WinsKitAttackingTeamGrenade                 int      `json:"wins_kit_attacking_team_grenade"`
				MostKillsGameKitAttackingTeamGrenade        int      `json:"most_kills_game_kit_attacking_team_grenade"`
				SurvivedPlayersKitAttackingTeamGrenade      int      `json:"survived_players_kit_attacking_team_grenade"`
				KillstreakKitAttackingTeamGrenade           int      `json:"killstreak_kit_attacking_team_grenade"`
				MeleeKillsKitAttackingTeamGrenade           int      `json:"melee_kills_kit_attacking_team_grenade"`
				GamesKitAttackingTeamEnergix                int      `json:"games_kit_attacking_team_energix"`
				TimePlayedKitBasicSoloFisherman             int      `json:"time_played_kit_basic_solo_fisherman"`
				ChestsOpenedKitBasicSoloFisherman           int      `json:"chests_opened_kit_basic_solo_fisherman"`
				MostKillsGameKitBasicSoloFisherman          int      `json:"most_kills_game_kit_basic_solo_fisherman"`
				VoidKillsKitBasicSoloFisherman              int      `json:"void_kills_kit_basic_solo_fisherman"`
				LongestBowShotKitAdvancedSoloKnight         int      `json:"longest_bow_shot_kit_advanced_solo_knight"`
				ArrowsHitKitAdvancedSoloKnight              int      `json:"arrows_hit_kit_advanced_solo_knight"`
				ArrowsShotKitAdvancedSoloKnight             int      `json:"arrows_shot_kit_advanced_solo_knight"`
				LongestBowShotKitSupportingTeamRookie       int      `json:"longest_bow_shot_kit_supporting_team_rookie"`
				FastestWinKitSupportingTeamRookie           int      `json:"fastest_win_kit_supporting_team_rookie"`
				WinsKitSupportingTeamRookie                 int      `json:"wins_kit_supporting_team_rookie"`
				ArrowsHitKitSupportingTeamRookie            int      `json:"arrows_hit_kit_supporting_team_rookie"`
				MostKillsGameKitSupportingTeamRookie        int      `json:"most_kills_game_kit_supporting_team_rookie"`
				GamesKitSupportingTeamRookie                int      `json:"games_kit_supporting_team_rookie"`
				AssistsKitSupportingTeamRookie              int      `json:"assists_kit_supporting_team_rookie"`
				MeleeKillsKitSupportingTeamRookie           int      `json:"melee_kills_kit_supporting_team_rookie"`
				SurvivedPlayersKitSupportingTeamRookie      int      `json:"survived_players_kit_supporting_team_rookie"`
				VoidKillsKitSupportingTeamRookie            int      `json:"void_kills_kit_supporting_team_rookie"`
				WinstreakKitSupportingTeamRookie            int      `json:"winstreak_kit_supporting_team_rookie"`
				ChestsOpenedKitSupportingTeamRookie         int      `json:"chests_opened_kit_supporting_team_rookie"`
				TimePlayedKitSupportingTeamRookie           int      `json:"time_played_kit_supporting_team_rookie"`
				KillstreakKitSupportingTeamRookie           int      `json:"killstreak_kit_supporting_team_rookie"`
				KillsKitSupportingTeamRookie                int      `json:"kills_kit_supporting_team_rookie"`
				ArrowsShotKitSupportingTeamRookie           int      `json:"arrows_shot_kit_supporting_team_rookie"`
				FastestWinKitBasicSoloFisherman             int      `json:"fastest_win_kit_basic_solo_fisherman"`
				MeleeKillsKitBasicSoloFisherman             int      `json:"melee_kills_kit_basic_solo_fisherman"`
				MobKillsKitBasicSoloFisherman               int      `json:"mob_kills_kit_basic_solo_fisherman"`
				LongestBowShotKitAttackingTeamKnight        int      `json:"longest_bow_shot_kit_attacking_team_knight"`
				ArrowsHitKitAttackingTeamKnight             int      `json:"arrows_hit_kit_attacking_team_knight"`
				MeleeKillsKitAttackingTeamEnergix           int      `json:"melee_kills_kit_attacking_team_energix"`
				MostKillsGameKitAttackingTeamEnergix        int      `json:"most_kills_game_kit_attacking_team_energix"`
				KillsKitAttackingTeamEnergix                int      `json:"kills_kit_attacking_team_energix"`
				FastestWinKitBasicSoloScout                 int      `json:"fastest_win_kit_basic_solo_scout"`
				KillsKitBasicSoloScout                      int      `json:"kills_kit_basic_solo_scout"`
				MostKillsGameKitBasicSoloScout              int      `json:"most_kills_game_kit_basic_solo_scout"`
				WinsKitBasicSoloScout                       int      `json:"wins_kit_basic_solo_scout"`
				WinstreakKitBasicSoloScout                  int      `json:"winstreak_kit_basic_solo_scout"`
				KillstreakKitBasicSoloScout                 int      `json:"killstreak_kit_basic_solo_scout"`
				GamesKitBasicSoloScout                      int      `json:"games_kit_basic_solo_scout"`
				VoidKillsKitBasicSoloScout                  int      `json:"void_kills_kit_basic_solo_scout"`
				FastestWinKitMegaMegaPaladin                int      `json:"fastest_win_kit_mega_mega_paladin"`
				TimePlayedKitMegaMegaPaladin                int      `json:"time_played_kit_mega_mega_paladin"`
				MeleeKillsKitBasicSoloScout                 int      `json:"melee_kills_kit_basic_solo_scout"`
				LongestBowShotKitEnderchestSoloEnderchest   int      `json:"longest_bow_shot_kit_enderchest_solo_enderchest"`
				FastestWinKitEnderchestSoloEnderchest       int      `json:"fastest_win_kit_enderchest_solo_enderchest"`
				MostKillsGameKitEnderchestSoloEnderchest    int      `json:"most_kills_game_kit_enderchest_solo_enderchest"`
				ArrowsHitKitEnderchestSoloEnderchest        int      `json:"arrows_hit_kit_enderchest_solo_enderchest"`
				MeleeKillsKitEnderchestSoloEnderchest       int      `json:"melee_kills_kit_enderchest_solo_enderchest"`
				VoidKillsKitEnderchestSoloEnderchest        int      `json:"void_kills_kit_enderchest_solo_enderchest"`
				ArrowsShotKitEnderchestSoloEnderchest       int      `json:"arrows_shot_kit_enderchest_solo_enderchest"`
				LongestBowShotKitBasicSoloScout             int      `json:"longest_bow_shot_kit_basic_solo_scout"`
				ArrowsShotKitBasicSoloScout                 int      `json:"arrows_shot_kit_basic_solo_scout"`
				ArrowsHitKitBasicSoloScout                  int      `json:"arrows_hit_kit_basic_solo_scout"`
				LongestBowShotKitAttackingTeamScout         int      `json:"longest_bow_shot_kit_attacking_team_scout"`
				FastestWinKitAttackingTeamScout             int      `json:"fastest_win_kit_attacking_team_scout"`
				ChestsOpenedKitAttackingTeamScout           int      `json:"chests_opened_kit_attacking_team_scout"`
				MobKillsKitAttackingTeamScout               int      `json:"mob_kills_kit_attacking_team_scout"`
				TimePlayedKitAttackingTeamScout             int      `json:"time_played_kit_attacking_team_scout"`
				WinstreakKitAttackingTeamScout              int      `json:"winstreak_kit_attacking_team_scout"`
				KillstreakKitAttackingTeamScout             int      `json:"killstreak_kit_attacking_team_scout"`
				MostKillsGameKitAttackingTeamScout          int      `json:"most_kills_game_kit_attacking_team_scout"`
				ArrowsShotKitAttackingTeamScout             int      `json:"arrows_shot_kit_attacking_team_scout"`
				MeleeKillsKitAttackingTeamScout             int      `json:"melee_kills_kit_attacking_team_scout"`
				VoidKillsKitAttackingTeamScout              int      `json:"void_kills_kit_attacking_team_scout"`
				ArrowsHitKitAttackingTeamScout              int      `json:"arrows_hit_kit_attacking_team_scout"`
				VoidKillsKitAttackingTeamEnderman           int      `json:"void_kills_kit_attacking_team_enderman"`
				FloorIsLavaExplained                        int      `json:"floor_is_lava_explained"`
				RushExplained                               int      `json:"rush_explained"`
				BlizzardExplained                           int      `json:"blizzard_explained"`
				ColorManiaExplained                         int      `json:"color_mania_explained"`
				TntMadnessExplained                         int      `json:"tnt_madness_explained"`
				SoloBlackMagic                              int      `json:"solo_black_magic"`
				TntMadnessExplainedLast                     int64    `json:"tnt_madness_explained_last"`
				RushExplainedLast                           int64    `json:"rush_explained_last"`
				FloorIsLavaExplainedLast                    int64    `json:"floor_is_lava_explained_last"`
				BlizzardExplainedLast                       int64    `json:"blizzard_explained_last"`
				KillByColorExplainedLast                    int64    `json:"kill_by_color_explained_last"`
				KillByColorExplained                        int      `json:"kill_by_color_explained"`
				WinStreakLab                                int      `json:"win_streak_lab"`
				ChestsOpenedLabSolo                         int      `json:"chests_opened_lab_solo"`
				DeathsLab                                   int      `json:"deaths_lab"`
				DeathsLabSolo                               int      `json:"deaths_lab_solo"`
				TimePlayedLabSolo                           int      `json:"time_played_lab_solo"`
				DeathsLabKitAdvancedSoloEnderman            int      `json:"deaths_lab_kit_advanced_solo_enderman"`
				TimePlayedLab                               int      `json:"time_played_lab"`
				LossesLabKitAdvancedSoloEnderman            int      `json:"losses_lab_kit_advanced_solo_enderman"`
				CoinsGainedLab                              int      `json:"coins_gained_lab"`
				LossesLabSolo                               int      `json:"losses_lab_solo"`
				ChestsOpenedLab                             int      `json:"chests_opened_lab"`
				ChestsOpenedLabKitAdvancedSoloEnderman      int      `json:"chests_opened_lab_kit_advanced_solo_enderman"`
				LossesLab                                   int      `json:"losses_lab"`
				QuitsLab                                    int      `json:"quits_lab"`
				TimePlayedLabKitAdvancedSoloEnderman        int      `json:"time_played_lab_kit_advanced_solo_enderman"`
				SurvivedPlayersLabSolo                      int      `json:"survived_players_lab_solo"`
				KillsLabKitAdvancedSoloEnderman             int      `json:"kills_lab_kit_advanced_solo_enderman"`
				VoidKillsLab                                int      `json:"void_kills_lab"`
				GamesLab                                    int      `json:"games_lab"`
				SurvivedPlayersLab                          int      `json:"survived_players_lab"`
				VoidKillsLabKitAdvancedSoloEnderman         int      `json:"void_kills_lab_kit_advanced_solo_enderman"`
				ArrowsShotLabSolo                           int      `json:"arrows_shot_lab_solo"`
				KillsLab                                    int      `json:"kills_lab"`
				GamesLabSolo                                int      `json:"games_lab_solo"`
				BlocksPlacedLab                             int      `json:"blocks_placed_lab"`
				KillsLabSolo                                int      `json:"kills_lab_solo"`
				GamesLabKitAdvancedSoloEnderman             int      `json:"games_lab_kit_advanced_solo_enderman"`
				ArrowsShotLabKitAdvancedSoloEnderman        int      `json:"arrows_shot_lab_kit_advanced_solo_enderman"`
				ArrowsHitLabKitAdvancedSoloEnderman         int      `json:"arrows_hit_lab_kit_advanced_solo_enderman"`
				ArrowsHitLab                                int      `json:"arrows_hit_lab"`
				VoidKillsLabSolo                            int      `json:"void_kills_lab_solo"`
				ArrowsShotLab                               int      `json:"arrows_shot_lab"`
				ArrowsHitLabSolo                            int      `json:"arrows_hit_lab_solo"`
				SurvivedPlayersLabKitAdvancedSoloEnderman   int      `json:"survived_players_lab_kit_advanced_solo_enderman"`
				AssistsLab                                  int      `json:"assists_lab"`
				AssistsLabSolo                              int      `json:"assists_lab_solo"`
				AssistsLabKitAdvancedSoloEnderman           int      `json:"assists_lab_kit_advanced_solo_enderman"`
				KillstreakLab                               int      `json:"killstreak_lab"`
				WinsLab                                     int      `json:"wins_lab"`
				WinsLabKitAdvancedSoloEnderman              int      `json:"wins_lab_kit_advanced_solo_enderman"`
				KillstreakLabKitAdvancedSoloEnderman        int      `json:"killstreak_lab_kit_advanced_solo_enderman"`
				WinstreakLab                                int      `json:"winstreak_lab"`
				WinstreakLabKitAdvancedSoloEnderman         int      `json:"winstreak_lab_kit_advanced_solo_enderman"`
				WinstreakLabSolo                            int      `json:"winstreak_lab_solo"`
				KillstreakLabSolo                           int      `json:"killstreak_lab_solo"`
				WinsLabSolo                                 int      `json:"wins_lab_solo"`
				EnderpearlsThrownLab                        int      `json:"enderpearls_thrown_lab"`
				BlocksBrokenLab                             int      `json:"blocks_broken_lab"`
				MeleeKillsLabKitAdvancedSoloEnderman        int      `json:"melee_kills_lab_kit_advanced_solo_enderman"`
				MeleeKillsLabSolo                           int      `json:"melee_kills_lab_solo"`
				MeleeKillsLab                               int      `json:"melee_kills_lab"`
				MobKillsLabSolo                             int      `json:"mob_kills_lab_solo"`
				EggThrownLab                                int      `json:"egg_thrown_lab"`
				MobKillsLab                                 int      `json:"mob_kills_lab"`
				MobKillsLabKitAdvancedSoloEnderman          int      `json:"mob_kills_lab_kit_advanced_solo_enderman"`
				VoidKillsLabKitAttackingTeamDefault         int      `json:"void_kills_lab_kit_attacking_team_default"`
				SurvivedPlayersLabTeam                      int      `json:"survived_players_lab_team"`
				TimePlayedLabKitAttackingTeamDefault        int      `json:"time_played_lab_kit_attacking_team_default"`
				SurvivedPlayersLabKitAttackingTeamDefault   int      `json:"survived_players_lab_kit_attacking_team_default"`
				MostKillsGameLabKitAttackingTeamDefault     int      `json:"most_kills_game_lab_kit_attacking_team_default"`
				KillsLabKitAttackingTeamDefault             int      `json:"kills_lab_kit_attacking_team_default"`
				AssistsLabKitAttackingTeamDefault           int      `json:"assists_lab_kit_attacking_team_default"`
				DeathsLabKitAttackingTeamDefault            int      `json:"deaths_lab_kit_attacking_team_default"`
				LossesLabTeam                               int      `json:"losses_lab_team"`
				MostKillsGameLabTeam                        int      `json:"most_kills_game_lab_team"`
				KillsLabTeam                                int      `json:"kills_lab_team"`
				DeathsLabTeam                               int      `json:"deaths_lab_team"`
				ChestsOpenedLabKitAttackingTeamDefault      int      `json:"chests_opened_lab_kit_attacking_team_default"`
				ChestsOpenedLabTeam                         int      `json:"chests_opened_lab_team"`
				TimePlayedLabTeam                           int      `json:"time_played_lab_team"`
				AssistsLabTeam                              int      `json:"assists_lab_team"`
				MostKillsGameLab                            int      `json:"most_kills_game_lab"`
				LossesLabKitAttackingTeamDefault            int      `json:"losses_lab_kit_attacking_team_default"`
				VoidKillsLabTeam                            int      `json:"void_kills_lab_team"`
				FastestWinLabKitAttackingTeamDefault        int      `json:"fastest_win_lab_kit_attacking_team_default"`
				LongestBowShotLab                           int      `json:"longest_bow_shot_lab"`
				LongestBowShotLabKitAttackingTeamDefault    int      `json:"longest_bow_shot_lab_kit_attacking_team_default"`
				FastestWinLabTeam                           int      `json:"fastest_win_lab_team"`
				FastestWinLab                               int      `json:"fastest_win_lab"`
				LongestBowShotLabTeam                       int      `json:"longest_bow_shot_lab_team"`
				ArrowsHitLabTeam                            int      `json:"arrows_hit_lab_team"`
				KillstreakLabKitAttackingTeamDefault        int      `json:"killstreak_lab_kit_attacking_team_default"`
				ArrowsShotLabTeam                           int      `json:"arrows_shot_lab_team"`
				WinstreakLabKitAttackingTeamDefault         int      `json:"winstreak_lab_kit_attacking_team_default"`
				WinsLabKitAttackingTeamDefault              int      `json:"wins_lab_kit_attacking_team_default"`
				MeleeKillsLabTeam                           int      `json:"melee_kills_lab_team"`
				ArrowsShotLabKitAttackingTeamDefault        int      `json:"arrows_shot_lab_kit_attacking_team_default"`
				WinstreakLabTeam                            int      `json:"winstreak_lab_team"`
				GamesLabKitAttackingTeamDefault             int      `json:"games_lab_kit_attacking_team_default"`
				GamesLabTeam                                int      `json:"games_lab_team"`
				KillstreakLabTeam                           int      `json:"killstreak_lab_team"`
				WinsLabTeam                                 int      `json:"wins_lab_team"`
				ArrowsHitLabKitAttackingTeamDefault         int      `json:"arrows_hit_lab_kit_attacking_team_default"`
				MeleeKillsLabKitAttackingTeamDefault        int      `json:"melee_kills_lab_kit_attacking_team_default"`
				VoidKillsKitMegaMegaHellhound               int      `json:"void_kills_kit_mega_mega_hellhound"`
				ItemsEnchantedLab                           int      `json:"items_enchanted_lab"`
				SurvivedPlayersKitAdvancedSoloPigRider      int      `json:"survived_players_kit_advanced_solo_pig-rider"`
				ChestsOpenedKitAdvancedSoloPigRider         int      `json:"chests_opened_kit_advanced_solo_pig-rider"`
				LossesKitAdvancedSoloPigRider               int      `json:"losses_kit_advanced_solo_pig-rider"`
				TimePlayedKitAdvancedSoloPigRider           int      `json:"time_played_kit_advanced_solo_pig-rider"`
				DeathsKitAdvancedSoloPigRider               int      `json:"deaths_kit_advanced_solo_pig-rider"`
				ActiveKitTEAMS                              string   `json:"activeKit_TEAMS"`
				SlimeExplained                              int      `json:"slime_explained"`
				SlimeExplainedLast                          int64    `json:"slime_explained_last"`
				SoloNecromancer                             int      `json:"solo_necromancer"`
				LongestBowShotKitAdvancedSoloPigRider       int      `json:"longest_bow_shot_kit_advanced_solo_pig-rider"`
				ArrowsHitKitAdvancedSoloPigRider            int      `json:"arrows_hit_kit_advanced_solo_pig-rider"`
				VoidKillsKitAdvancedSoloPigRider            int      `json:"void_kills_kit_advanced_solo_pig-rider"`
				MostKillsGameKitAdvancedSoloPigRider        int      `json:"most_kills_game_kit_advanced_solo_pig-rider"`
				ArrowsShotKitAdvancedSoloPigRider           int      `json:"arrows_shot_kit_advanced_solo_pig-rider"`
				KillsKitAdvancedSoloPigRider                int      `json:"kills_kit_advanced_solo_pig-rider"`
				MeleeKillsKitAdvancedSoloPigRider           int      `json:"melee_kills_kit_advanced_solo_pig-rider"`
				GamesKitAdvancedSoloPigRider                int      `json:"games_kit_advanced_solo_pig-rider"`
				FastestWinKitAdvancedSoloPigRider           int      `json:"fastest_win_kit_advanced_solo_pig-rider"`
				WinsKitAdvancedSoloPigRider                 int      `json:"wins_kit_advanced_solo_pig-rider"`
				WinstreakKitAdvancedSoloPigRider            int      `json:"winstreak_kit_advanced_solo_pig-rider"`
				KillstreakKitAdvancedSoloPigRider           int      `json:"killstreak_kit_advanced_solo_pig-rider"`
				SurvivedPlayersLabKitAdvancedSoloPigRider   int      `json:"survived_players_lab_kit_advanced_solo_pig-rider"`
				TimePlayedLabKitAdvancedSoloPigRider        int      `json:"time_played_lab_kit_advanced_solo_pig-rider"`
				ChestsOpenedLabKitAdvancedSoloPigRider      int      `json:"chests_opened_lab_kit_advanced_solo_pig-rider"`
				LossesLabKitAdvancedSoloPigRider            int      `json:"losses_lab_kit_advanced_solo_pig-rider"`
				DeathsLabKitAdvancedSoloPigRider            int      `json:"deaths_lab_kit_advanced_solo_pig-rider"`
				KillsLabKitAdvancedSoloPigRider             int      `json:"kills_lab_kit_advanced_solo_pig-rider"`
				MeleeKillsLabKitAdvancedSoloPigRider        int      `json:"melee_kills_lab_kit_advanced_solo_pig-rider"`
				FastestWinLabSolo                           int      `json:"fastest_win_lab_solo"`
				FastestWinLabKitAdvancedSoloPigRider        int      `json:"fastest_win_lab_kit_advanced_solo_pig-rider"`
				WinsLabKitAdvancedSoloPigRider              int      `json:"wins_lab_kit_advanced_solo_pig-rider"`
				LabWinSlimeLab                              int      `json:"lab_win_slime_lab"`
				WinstreakLabKitAdvancedSoloPigRider         int      `json:"winstreak_lab_kit_advanced_solo_pig-rider"`
				ArrowsHitLabKitAdvancedSoloPigRider         int      `json:"arrows_hit_lab_kit_advanced_solo_pig-rider"`
				LabWinSlimeLabSolo                          int      `json:"lab_win_slime_lab_solo"`
				GamesLabKitAdvancedSoloPigRider             int      `json:"games_lab_kit_advanced_solo_pig-rider"`
				LabWinSlimeLabKitAdvancedSoloPigRider       int      `json:"lab_win_slime_lab_kit_advanced_solo_pig-rider"`
				VoidKillsLabKitAdvancedSoloPigRider         int      `json:"void_kills_lab_kit_advanced_solo_pig-rider"`
				ArrowsShotLabKitAdvancedSoloPigRider        int      `json:"arrows_shot_lab_kit_advanced_solo_pig-rider"`
				KillstreakLabKitAdvancedSoloPigRider        int      `json:"killstreak_lab_kit_advanced_solo_pig-rider"`
				LabWinTntMadnessLabSolo                     int      `json:"lab_win_tnt_madness_lab_solo"`
				LabWinTntMadnessLabKitAdvancedSoloPigRider  int      `json:"lab_win_tnt_madness_lab_kit_advanced_solo_pig-rider"`
				LabWinTntMadnessLab                         int      `json:"lab_win_tnt_madness_lab"`
				LongestBowShotLabKitAdvancedSoloPigRider    int      `json:"longest_bow_shot_lab_kit_advanced_solo_pig-rider"`
				LongestBowShotLabSolo                       int      `json:"longest_bow_shot_lab_solo"`
				AssistsLabKitAdvancedSoloPigRider           int      `json:"assists_lab_kit_advanced_solo_pig-rider"`
				LabWinTntMadnessLabKitAdvancedSoloEnderman  int      `json:"lab_win_tnt_madness_lab_kit_advanced_solo_enderman"`
				ChestsOpenedKitAdvancedSoloSalmon           int      `json:"chests_opened_kit_advanced_solo_salmon"`
				TimePlayedKitAdvancedSoloSalmon             int      `json:"time_played_kit_advanced_solo_salmon"`
				LossesKitAdvancedSoloSalmon                 int      `json:"losses_kit_advanced_solo_salmon"`
				DeathsKitAdvancedSoloSalmon                 int      `json:"deaths_kit_advanced_solo_salmon"`
				FastestWinKitAdvancedSoloSalmon             int      `json:"fastest_win_kit_advanced_solo_salmon"`
				KillstreakKitAdvancedSoloSalmon             int      `json:"killstreak_kit_advanced_solo_salmon"`
				KillsKitAdvancedSoloSalmon                  int      `json:"kills_kit_advanced_solo_salmon"`
				MeleeKillsKitAdvancedSoloSalmon             int      `json:"melee_kills_kit_advanced_solo_salmon"`
				WinstreakKitAdvancedSoloSalmon              int      `json:"winstreak_kit_advanced_solo_salmon"`
				SurvivedPlayersKitAdvancedSoloSalmon        int      `json:"survived_players_kit_advanced_solo_salmon"`
				MostKillsGameKitAdvancedSoloSalmon          int      `json:"most_kills_game_kit_advanced_solo_salmon"`
				WinsKitAdvancedSoloSalmon                   int      `json:"wins_kit_advanced_solo_salmon"`
				GamesKitAdvancedSoloSalmon                  int      `json:"games_kit_advanced_solo_salmon"`
				ArrowsShotKitAdvancedSoloSalmon             int      `json:"arrows_shot_kit_advanced_solo_salmon"`
				VoidKillsKitAdvancedSoloSalmon              int      `json:"void_kills_kit_advanced_solo_salmon"`
				AssistsKitAdvancedSoloSalmon                int      `json:"assists_kit_advanced_solo_salmon"`
				LongestBowShotKitAdvancedSoloSalmon         int      `json:"longest_bow_shot_kit_advanced_solo_salmon"`
				ArrowsHitKitAdvancedSoloSalmon              int      `json:"arrows_hit_kit_advanced_solo_salmon"`
				MobKillsKitAdvancedSoloEnderman             int      `json:"mob_kills_kit_advanced_solo_enderman"`
				MegaBlackMagic                              int      `json:"mega_black_magic"`
				MegaNecromancer                             int      `json:"mega_necromancer"`
				TeamNecromancer                             int      `json:"team_necromancer"`
				TeamBlackMagic                              int      `json:"team_black_magic"`
				MegaEnvironmentalExpert                     int      `json:"mega_environmental_expert"`
				KitMegaMegaPyromaniac                       int      `json:"kit_mega_mega_pyromaniac"`
				KitMegaMegaFisherman                        int      `json:"kit_mega_mega_fisherman"`
				LongestBowKillKitAdvancedSoloEnderman       int      `json:"longest_bow_kill_kit_advanced_solo_enderman"`
				LongestBowKill                              int      `json:"longest_bow_kill"`
				LongestBowKillSolo                          int      `json:"longest_bow_kill_solo"`
				InGamePresentsCap201719                     int      `json:"inGamePresentsCap_2017_19"`
				InGamePresentsCap201722                     int      `json:"inGamePresentsCap_2017_22"`
				LongestBowKillKitDefendingTeamArmorer       int      `json:"longest_bow_kill_kit_defending_team_armorer"`
				LongestBowKillTeam                          int      `json:"longest_bow_kill_team"`
				SkywarsChests                               int      `json:"skywars_chests"`
				SkyWarsOpenedChests                         int      `json:"SkyWars_openedChests"`
				SkywarsChestHistory                         []string `json:"skywars_chest_history"`
				SkyWarsOpenedCommons                        int      `json:"SkyWars_openedCommons"`
				SkyWarsOpenedRares                          int      `json:"SkyWars_openedRares"`
				CosmeticTokens                              int      `json:"cosmetic_tokens"`
				SkyWarsOpenedEpics                          int      `json:"SkyWars_openedEpics"`
				SkyWarsOpenedLegendaries                    int      `json:"SkyWars_openedLegendaries"`
				ActiveBalloon                               string   `json:"active_balloon"`
				ActiveSprays                                string   `json:"active_sprays"`
				ActiveVictorydance                          string   `json:"active_victorydance"`
				LossesKitAdvancedSoloSlime                  int      `json:"losses_kit_advanced_solo_slime"`
				GamesKitAdvancedSoloSlime                   int      `json:"games_kit_advanced_solo_slime"`
				TimePlayedKitAdvancedSoloSlime              int      `json:"time_played_kit_advanced_solo_slime"`
				DeathsKitAdvancedSoloSlime                  int      `json:"deaths_kit_advanced_solo_slime"`
				ChestsOpenedKitAdvancedSoloSlime            int      `json:"chests_opened_kit_advanced_solo_slime"`
				SurvivedPlayersKitAdvancedSoloSlime         int      `json:"survived_players_kit_advanced_solo_slime"`
				ActiveProjectiletrail                       string   `json:"active_projectiletrail"`
				ActiveKilleffect                            string   `json:"active_killeffect"`
				ActiveDeathcry                              string   `json:"active_deathcry"`
				LuckyExplained                              int      `json:"lucky_explained"`
				LuckyExplainedLast                          int64    `json:"lucky_explained_last"`
				TimePlayedLabKitAdvancedSoloSlime           int      `json:"time_played_lab_kit_advanced_solo_slime"`
				SurvivedPlayersLabKitAdvancedSoloSlime      int      `json:"survived_players_lab_kit_advanced_solo_slime"`
				MostKillsGameLabSolo                        int      `json:"most_kills_game_lab_solo"`
				LossesLabKitAdvancedSoloSlime               int      `json:"losses_lab_kit_advanced_solo_slime"`
				DeathsLabKitAdvancedSoloSlime               int      `json:"deaths_lab_kit_advanced_solo_slime"`
				ArrowsShotLabKitAdvancedSoloSlime           int      `json:"arrows_shot_lab_kit_advanced_solo_slime"`
				KillsLabKitAdvancedSoloSlime                int      `json:"kills_lab_kit_advanced_solo_slime"`
				MostKillsGameLabKitAdvancedSoloSlime        int      `json:"most_kills_game_lab_kit_advanced_solo_slime"`
				ChestsOpenedLabKitAdvancedSoloSlime         int      `json:"chests_opened_lab_kit_advanced_solo_slime"`
				LongestBowKillLabSolo                       int      `json:"longest_bow_kill_lab_solo"`
				FastestWinLabKitAdvancedSoloSlime           int      `json:"fastest_win_lab_kit_advanced_solo_slime"`
				LongestBowKillLab                           int      `json:"longest_bow_kill_lab"`
				LongestBowKillLabKitAdvancedSoloSlime       int      `json:"longest_bow_kill_lab_kit_advanced_solo_slime"`
				LongestBowShotLabKitAdvancedSoloSlime       int      `json:"longest_bow_shot_lab_kit_advanced_solo_slime"`
				WinsLabKitAdvancedSoloSlime                 int      `json:"wins_lab_kit_advanced_solo_slime"`
				LabWinLuckyBlocksLabKitAdvancedSoloSlime    int      `json:"lab_win_lucky_blocks_lab_kit_advanced_solo_slime"`
				VoidKillsLabKitAdvancedSoloSlime            int      `json:"void_kills_lab_kit_advanced_solo_slime"`
				WinstreakLabKitAdvancedSoloSlime            int      `json:"winstreak_lab_kit_advanced_solo_slime"`
				LabWinLuckyBlocksLabSolo                    int      `json:"lab_win_lucky_blocks_lab_solo"`
				LabWinLuckyBlocksLab                        int      `json:"lab_win_lucky_blocks_lab"`
				KillstreakLabKitAdvancedSoloSlime           int      `json:"killstreak_lab_kit_advanced_solo_slime"`
				GamesLabKitAdvancedSoloSlime                int      `json:"games_lab_kit_advanced_solo_slime"`
				MeleeKillsLabKitAdvancedSoloSlime           int      `json:"melee_kills_lab_kit_advanced_solo_slime"`
				ArrowsHitLabKitAdvancedSoloSlime            int      `json:"arrows_hit_lab_kit_advanced_solo_slime"`
				MobKillsLabKitAdvancedSoloSlime             int      `json:"mob_kills_lab_kit_advanced_solo_slime"`
				AssistsLabKitAdvancedSoloSlime              int      `json:"assists_lab_kit_advanced_solo_slime"`
				FreeLootChestNpc                            int64    `json:"freeLootChestNpc"`
				LongestBowKillKitBasicSoloScout             int      `json:"longest_bow_kill_kit_basic_solo_scout"`
				LongestBowKillKitAttackingTeamScout         int      `json:"longest_bow_kill_kit_attacking_team_scout"`
				TimePlayedLabKitBasicSoloScout              int      `json:"time_played_lab_kit_basic_solo_scout"`
				SurvivedPlayersLabKitBasicSoloScout         int      `json:"survived_players_lab_kit_basic_solo_scout"`
				ChestsOpenedLabKitBasicSoloScout            int      `json:"chests_opened_lab_kit_basic_solo_scout"`
				DeathsLabKitBasicSoloScout                  int      `json:"deaths_lab_kit_basic_solo_scout"`
				LossesLabKitBasicSoloScout                  int      `json:"losses_lab_kit_basic_solo_scout"`
				LongestBowKillLabKitBasicSoloScout          int      `json:"longest_bow_kill_lab_kit_basic_solo_scout"`
				WinstreakLabKitBasicSoloScout               int      `json:"winstreak_lab_kit_basic_solo_scout"`
				ArrowsHitLabKitBasicSoloScout               int      `json:"arrows_hit_lab_kit_basic_solo_scout"`
				KillsLabKitBasicSoloScout                   int      `json:"kills_lab_kit_basic_solo_scout"`
				VoidKillsLabKitBasicSoloScout               int      `json:"void_kills_lab_kit_basic_solo_scout"`
				GamesLabKitBasicSoloScout                   int      `json:"games_lab_kit_basic_solo_scout"`
				LabWinLuckyBlocksLabKitBasicSoloScout       int      `json:"lab_win_lucky_blocks_lab_kit_basic_solo_scout"`
				ArrowsShotLabKitBasicSoloScout              int      `json:"arrows_shot_lab_kit_basic_solo_scout"`
				WinsLabKitBasicSoloScout                    int      `json:"wins_lab_kit_basic_solo_scout"`
				MeleeKillsLabKitBasicSoloScout              int      `json:"melee_kills_lab_kit_basic_solo_scout"`
				KillstreakLabKitBasicSoloScout              int      `json:"killstreak_lab_kit_basic_solo_scout"`
				ActiveCage                                  string   `json:"active_cage"`
				VoidKillsLabKitAttackingTeamScout           int      `json:"void_kills_lab_kit_attacking_team_scout"`
				DeathsLabKitAttackingTeamScout              int      `json:"deaths_lab_kit_attacking_team_scout"`
				KillsLabKitAttackingTeamScout               int      `json:"kills_lab_kit_attacking_team_scout"`
				SurvivedPlayersLabKitAttackingTeamScout     int      `json:"survived_players_lab_kit_attacking_team_scout"`
				ChestsOpenedLabKitAttackingTeamScout        int      `json:"chests_opened_lab_kit_attacking_team_scout"`
				MeleeKillsLabKitAttackingTeamScout          int      `json:"melee_kills_lab_kit_attacking_team_scout"`
				LossesLabKitAttackingTeamScout              int      `json:"losses_lab_kit_attacking_team_scout"`
				TimePlayedLabKitAttackingTeamScout          int      `json:"time_played_lab_kit_attacking_team_scout"`
				ArrowsHitLabKitAttackingTeamScout           int      `json:"arrows_hit_lab_kit_attacking_team_scout"`
				ArrowsShotLabKitAttackingTeamScout          int      `json:"arrows_shot_lab_kit_attacking_team_scout"`
				LabWinTntMadnessLabTeam                     int      `json:"lab_win_tnt_madness_lab_team"`
				GamesLabKitAttackingTeamScout               int      `json:"games_lab_kit_attacking_team_scout"`
				WinstreakLabKitAttackingTeamScout           int      `json:"winstreak_lab_kit_attacking_team_scout"`
				LabWinTntMadnessLabKitAttackingTeamScout    int      `json:"lab_win_tnt_madness_lab_kit_attacking_team_scout"`
				WinsLabKitAttackingTeamScout                int      `json:"wins_lab_kit_attacking_team_scout"`
				KillstreakLabKitAttackingTeamScout          int      `json:"killstreak_lab_kit_attacking_team_scout"`
				MobKillsLabTeam                             int      `json:"mob_kills_lab_team"`
				AssistsLabKitAttackingTeamScout             int      `json:"assists_lab_kit_attacking_team_scout"`
				MobKillsLabKitAttackingTeamScout            int      `json:"mob_kills_lab_kit_attacking_team_scout"`
				LongestBowKillKitMegaMegaArmorsmith         int      `json:"longest_bow_kill_kit_mega_mega_armorsmith"`
				LongestBowKillMega                          int      `json:"longest_bow_kill_mega"`
				LongestBowShotLabKitAttackingTeamScout      int      `json:"longest_bow_shot_lab_kit_attacking_team_scout"`
				LabWinLuckyBlocksLabKitAdvancedSoloEnderman int      `json:"lab_win_lucky_blocks_lab_kit_advanced_solo_enderman"`
				BowKills                                    int      `json:"bow_kills"`
				BowKillsSolo                                int      `json:"bow_kills_solo"`
				BowKillsKitAdvancedSoloEnderman             int      `json:"bow_kills_kit_advanced_solo_enderman"`
				LabWinSlimeLabKitAdvancedSoloEnderman       int      `json:"lab_win_slime_lab_kit_advanced_solo_enderman"`
				SkywarsEasterBoxes                          int      `json:"skywars_easter_boxes"`
				ChestHistoryNew                             []string `json:"chest_history_new"`
				SoloRobbery                                 int      `json:"solo_robbery"`
				ActiveKitSOLORandom                         bool     `json:"activeKit_SOLO_random"`
				KitAdvancedSoloJesterInventory              struct {
					POTION1    string `json:"POTION:1"`
					WOODSWORD0 string `json:"WOOD_SWORD:0"`
				} `json:"kit_advanced_solo_jester_inventory"`
				LongestBowKillKitAdvancedSoloJester                       int   `json:"longest_bow_kill_kit_advanced_solo_jester"`
				FastestWinKitAdvancedSoloJester                           int   `json:"fastest_win_kit_advanced_solo_jester"`
				MostKillsGameKitAdvancedSoloJester                        int   `json:"most_kills_game_kit_advanced_solo_jester"`
				LongestBowShotKitAdvancedSoloJester                       int   `json:"longest_bow_shot_kit_advanced_solo_jester"`
				WinstreakKitAdvancedSoloJester                            int   `json:"winstreak_kit_advanced_solo_jester"`
				AssistsKitAdvancedSoloJester                              int   `json:"assists_kit_advanced_solo_jester"`
				KillstreakKitAdvancedSoloJester                           int   `json:"killstreak_kit_advanced_solo_jester"`
				KillsKitAdvancedSoloJester                                int   `json:"kills_kit_advanced_solo_jester"`
				SurvivedPlayersKitAdvancedSoloJester                      int   `json:"survived_players_kit_advanced_solo_jester"`
				MeleeKillsKitAdvancedSoloJester                           int   `json:"melee_kills_kit_advanced_solo_jester"`
				TimePlayedKitAdvancedSoloJester                           int   `json:"time_played_kit_advanced_solo_jester"`
				ArrowsShotKitAdvancedSoloJester                           int   `json:"arrows_shot_kit_advanced_solo_jester"`
				GamesKitAdvancedSoloJester                                int   `json:"games_kit_advanced_solo_jester"`
				WinsKitAdvancedSoloJester                                 int   `json:"wins_kit_advanced_solo_jester"`
				VoidKillsKitAdvancedSoloJester                            int   `json:"void_kills_kit_advanced_solo_jester"`
				ChestsOpenedKitAdvancedSoloJester                         int   `json:"chests_opened_kit_advanced_solo_jester"`
				ArrowsHitKitAdvancedSoloJester                            int   `json:"arrows_hit_kit_advanced_solo_jester"`
				LongestBowKillKitBasicSoloArmorsmith                      int   `json:"longest_bow_kill_kit_basic_solo_armorsmith"`
				BowKillsKitBasicSoloArmorsmith                            int   `json:"bow_kills_kit_basic_solo_armorsmith"`
				MobKillsKitBasicSoloArmorsmith                            int   `json:"mob_kills_kit_basic_solo_armorsmith"`
				ActiveKitTEAMSRandom                                      bool  `json:"activeKit_TEAMS_random"`
				BowKillsTeam                                              int   `json:"bow_kills_team"`
				BowKillsKitDefendingTeamArmorer                           int   `json:"bow_kills_kit_defending_team_armorer"`
				HuntersVsBeastsExplained                                  int   `json:"hunters_vs_beasts_explained"`
				HuntersVsBeastsExplainedLast                              int64 `json:"hunters_vs_beasts_explained_last"`
				BeastChance                                               int   `json:"beast_chance"`
				TimePlayedLabKitAdvancedSoloCannoneer                     int   `json:"time_played_lab_kit_advanced_solo_cannoneer"`
				KillsLabKitAdvancedSoloCannoneer                          int   `json:"kills_lab_kit_advanced_solo_cannoneer"`
				LossesLabKitAdvancedSoloCannoneer                         int   `json:"losses_lab_kit_advanced_solo_cannoneer"`
				GamesLabKitAdvancedSoloCannoneer                          int   `json:"games_lab_kit_advanced_solo_cannoneer"`
				DeathsLabKitAdvancedSoloCannoneer                         int   `json:"deaths_lab_kit_advanced_solo_cannoneer"`
				ChestsOpenedLabKitAdvancedSoloCannoneer                   int   `json:"chests_opened_lab_kit_advanced_solo_cannoneer"`
				SurvivedPlayersLabKitAdvancedSoloCannoneer                int   `json:"survived_players_lab_kit_advanced_solo_cannoneer"`
				FastestWinLabKitAdvancedSoloCannoneer                     int   `json:"fastest_win_lab_kit_advanced_solo_cannoneer"`
				LabHvbHunterWinsLab                                       int   `json:"lab_hvb_hunter_wins_lab"`
				WinsLabKitAdvancedSoloCannoneer                           int   `json:"wins_lab_kit_advanced_solo_cannoneer"`
				LabWinHuntersVsBeastsLabSolo                              int   `json:"lab_win_hunters_vs_beasts_lab_solo"`
				LabWinHuntersVsBeastsLab                                  int   `json:"lab_win_hunters_vs_beasts_lab"`
				LabWinHuntersVsBeastsLabKitAdvancedSoloCannoneer          int   `json:"lab_win_hunters_vs_beasts_lab_kit_advanced_solo_cannoneer"`
				WinstreakLabKitAdvancedSoloCannoneer                      int   `json:"winstreak_lab_kit_advanced_solo_cannoneer"`
				LongestBowKillLabKitAdvancedSoloCannoneer                 int   `json:"longest_bow_kill_lab_kit_advanced_solo_cannoneer"`
				MostKillsGameLabKitAdvancedSoloCannoneer                  int   `json:"most_kills_game_lab_kit_advanced_solo_cannoneer"`
				KillstreakLabKitAdvancedSoloCannoneer                     int   `json:"killstreak_lab_kit_advanced_solo_cannoneer"`
				MeleeKillsLabKitAdvancedSoloCannoneer                     int   `json:"melee_kills_lab_kit_advanced_solo_cannoneer"`
				LabHvbBeastWinsLab                                        int   `json:"lab_hvb_beast_wins_lab"`
				MobKillsLabKitAdvancedSoloCannoneer                       int   `json:"mob_kills_lab_kit_advanced_solo_cannoneer"`
				VoidKillsLabKitAdvancedSoloCannoneer                      int   `json:"void_kills_lab_kit_advanced_solo_cannoneer"`
				InGamePresentsCap20186                                    int   `json:"inGamePresentsCap_2018_6"`
				SkywarsExperience                                         int   `json:"skywars_experience"`
				TeamBulldozer                                             int   `json:"team_bulldozer"`
				TeamRobbery                                               int   `json:"team_robbery"`
				FastestWinKitAttackingTeamFisherman                       int   `json:"fastest_win_kit_attacking_team_fisherman"`
				MostKillsGameKitAttackingTeamFisherman                    int   `json:"most_kills_game_kit_attacking_team_fisherman"`
				KillstreakKitAttackingTeamFisherman                       int   `json:"killstreak_kit_attacking_team_fisherman"`
				KillsKitAttackingTeamFisherman                            int   `json:"kills_kit_attacking_team_fisherman"`
				GamesKitAttackingTeamFisherman                            int   `json:"games_kit_attacking_team_fisherman"`
				WinstreakKitAttackingTeamFisherman                        int   `json:"winstreak_kit_attacking_team_fisherman"`
				ChallengeWins                                             int   `json:"challenge_wins"`
				ChallengeWins0Solo                                        int   `json:"challenge_wins_0_solo"`
				ChallengeWins0KitAttackingTeamFisherman                   int   `json:"challenge_wins_0_kit_attacking_team_fisherman"`
				TimePlayedKitAttackingTeamFisherman                       int   `json:"time_played_kit_attacking_team_fisherman"`
				ChallengeWinsSolo                                         int   `json:"challenge_wins_solo"`
				SurvivedPlayersKitAttackingTeamFisherman                  int   `json:"survived_players_kit_attacking_team_fisherman"`
				VoidKillsKitAttackingTeamFisherman                        int   `json:"void_kills_kit_attacking_team_fisherman"`
				ChallengeWins0                                            int   `json:"challenge_wins_0"`
				WinsKitAttackingTeamFisherman                             int   `json:"wins_kit_attacking_team_fisherman"`
				ChestsOpenedKitAttackingTeamFisherman                     int   `json:"chests_opened_kit_attacking_team_fisherman"`
				ChallengeWinsKitAttackingTeamFisherman                    int   `json:"challenge_wins_kit_attacking_team_fisherman"`
				InGamePresentsCap201818                                   int   `json:"inGamePresentsCap_2018_18"`
				DeathsKitAttackingTeamFisherman                           int   `json:"deaths_kit_attacking_team_fisherman"`
				LossesKitAttackingTeamFisherman                           int   `json:"losses_kit_attacking_team_fisherman"`
				LongestBowShotKitAttackingTeamFisherman                   int   `json:"longest_bow_shot_kit_attacking_team_fisherman"`
				LongestBowKillKitAttackingTeamFisherman                   int   `json:"longest_bow_kill_kit_attacking_team_fisherman"`
				MeleeKillsKitAttackingTeamFisherman                       int   `json:"melee_kills_kit_attacking_team_fisherman"`
				ArrowsHitKitAttackingTeamFisherman                        int   `json:"arrows_hit_kit_attacking_team_fisherman"`
				MobKillsKitAttackingTeamFisherman                         int   `json:"mob_kills_kit_attacking_team_fisherman"`
				ArrowsShotKitAttackingTeamFisherman                       int   `json:"arrows_shot_kit_attacking_team_fisherman"`
				ChallengeAttemptsHalfHealth                               int   `json:"challenge_attempts_half_health"`
				ChallengeAttemptsRookieSolo                               int   `json:"challenge_attempts_rookie_solo"`
				ChallengeAttempts8Solo                                    int   `json:"challenge_attempts_8_solo"`
				ChallengeAttemptsArcherSolo                               int   `json:"challenge_attempts_archer_solo"`
				ChallengeAttemptsArcher                                   int   `json:"challenge_attempts_archer"`
				ChallengeAttemptsNoChestKitAdvancedSoloEnderman           int   `json:"challenge_attempts_no_chest_kit_advanced_solo_enderman"`
				ChallengeAttemptsHalfHealthKitAdvancedSoloEnderman        int   `json:"challenge_attempts_half_health_kit_advanced_solo_enderman"`
				ChallengeAttemptsUltimateWarrior                          int   `json:"challenge_attempts_ultimate_warrior"`
				ChallengeAttemptsNoChestSolo                              int   `json:"challenge_attempts_no_chest_solo"`
				ChallengeAttemptsHalfHealthSolo                           int   `json:"challenge_attempts_half_health_solo"`
				ChallengeAttemptsRookieKitAdvancedSoloEnderman            int   `json:"challenge_attempts_rookie_kit_advanced_solo_enderman"`
				ChallengeAttemptsNoBlockKitAdvancedSoloEnderman           int   `json:"challenge_attempts_no_block_kit_advanced_solo_enderman"`
				ChallengeAttemptsPaper                                    int   `json:"challenge_attempts_paper"`
				ChallengeAttemptsUhcSolo                                  int   `json:"challenge_attempts_uhc_solo"`
				ChallengeAttempts8KitAdvancedSoloEnderman                 int   `json:"challenge_attempts_8_kit_advanced_solo_enderman"`
				ChallengeAttemptsNoBlockSolo                              int   `json:"challenge_attempts_no_block_solo"`
				ChallengeAttemptsNoBlock                                  int   `json:"challenge_attempts_no_block"`
				ChallengeAttemptsNoChest                                  int   `json:"challenge_attempts_no_chest"`
				ChallengeAttemptsUhc                                      int   `json:"challenge_attempts_uhc"`
				ChallengeAttemptsPaperKitAdvancedSoloEnderman             int   `json:"challenge_attempts_paper_kit_advanced_solo_enderman"`
				ChallengeAttemptsPaperSolo                                int   `json:"challenge_attempts_paper_solo"`
				ChallengeAttemptsUltimateWarriorKitAdvancedSoloEnderman   int   `json:"challenge_attempts_ultimate_warrior_kit_advanced_solo_enderman"`
				ChallengeAttemptsSolo                                     int   `json:"challenge_attempts_solo"`
				ChallengeAttempts8                                        int   `json:"challenge_attempts_8"`
				ChallengeAttemptsArcherKitAdvancedSoloEnderman            int   `json:"challenge_attempts_archer_kit_advanced_solo_enderman"`
				ChallengeAttemptsRookie                                   int   `json:"challenge_attempts_rookie"`
				ChallengeAttemptsKitAdvancedSoloEnderman                  int   `json:"challenge_attempts_kit_advanced_solo_enderman"`
				ChallengeAttemptsUhcKitAdvancedSoloEnderman               int   `json:"challenge_attempts_uhc_kit_advanced_solo_enderman"`
				ChallengeAttemptsUltimateWarriorSolo                      int   `json:"challenge_attempts_ultimate_warrior_solo"`
				ChallengeAttempts                                         int   `json:"challenge_attempts"`
				ChallengeWinsNoChestSolo                                  int   `json:"challenge_wins_no_chest_solo"`
				ChallengeWinsArcher                                       int   `json:"challenge_wins_archer"`
				ChallengeWinsArcherSolo                                   int   `json:"challenge_wins_archer_solo"`
				ChallengeAttempts8KitAttackingTeamFisherman               int   `json:"challenge_attempts_8_kit_attacking_team_fisherman"`
				ChallengeWinsRookieKitAttackingTeamFisherman              int   `json:"challenge_wins_rookie_kit_attacking_team_fisherman"`
				ChallengeWinsUltimateWarriorKitAttackingTeamFisherman     int   `json:"challenge_wins_ultimate_warrior_kit_attacking_team_fisherman"`
				ChallengeWinsUltimateWarriorSolo                          int   `json:"challenge_wins_ultimate_warrior_solo"`
				ChallengeWinsNoBlockSolo                                  int   `json:"challenge_wins_no_block_solo"`
				ChallengeWinsUhcKitAttackingTeamFisherman                 int   `json:"challenge_wins_uhc_kit_attacking_team_fisherman"`
				ChallengeWinsHalfHealth                                   int   `json:"challenge_wins_half_health"`
				ChallengeAttemptsNoChestKitAttackingTeamFisherman         int   `json:"challenge_attempts_no_chest_kit_attacking_team_fisherman"`
				ChallengeWinsRookieSolo                                   int   `json:"challenge_wins_rookie_solo"`
				ChallengeWinsUhc                                          int   `json:"challenge_wins_uhc"`
				ChallengeWinsNoChestKitAttackingTeamFisherman             int   `json:"challenge_wins_no_chest_kit_attacking_team_fisherman"`
				ChallengeWinsNoBlock                                      int   `json:"challenge_wins_no_block"`
				ChallengeWinsNoChest                                      int   `json:"challenge_wins_no_chest"`
				ChallengeAttemptsPaperKitAttackingTeamFisherman           int   `json:"challenge_attempts_paper_kit_attacking_team_fisherman"`
				ChallengeAttemptsNoBlockKitAttackingTeamFisherman         int   `json:"challenge_attempts_no_block_kit_attacking_team_fisherman"`
				ChallengeAttemptsUltimateWarriorKitAttackingTeamFisherman int   `json:"challenge_attempts_ultimate_warrior_kit_attacking_team_fisherman"`
				ChallengeWinsPaperSolo                                    int   `json:"challenge_wins_paper_solo"`
				ChallengeWinsPaperKitAttackingTeamFisherman               int   `json:"challenge_wins_paper_kit_attacking_team_fisherman"`
				ChallengeWinsHalfHealthSolo                               int   `json:"challenge_wins_half_health_solo"`
				ChallengeWinsUhcSolo                                      int   `json:"challenge_wins_uhc_solo"`
				ChallengeWinsNoBlockKitAttackingTeamFisherman             int   `json:"challenge_wins_no_block_kit_attacking_team_fisherman"`
				ChallengeWinsUltimateWarrior                              int   `json:"challenge_wins_ultimate_warrior"`
				ChallengeAttemptsRookieKitAttackingTeamFisherman          int   `json:"challenge_attempts_rookie_kit_attacking_team_fisherman"`
				ChallengeWinsPaper                                        int   `json:"challenge_wins_paper"`
				ChallengeWinsRookie                                       int   `json:"challenge_wins_rookie"`
				ChallengeAttemptsArcherKitAttackingTeamFisherman          int   `json:"challenge_attempts_archer_kit_attacking_team_fisherman"`
				ChallengeAttemptsKitAttackingTeamFisherman                int   `json:"challenge_attempts_kit_attacking_team_fisherman"`
				ChallengeWinsArcherKitAttackingTeamFisherman              int   `json:"challenge_wins_archer_kit_attacking_team_fisherman"`
				ChallengeWinsHalfHealthKitAttackingTeamFisherman          int   `json:"challenge_wins_half_health_kit_attacking_team_fisherman"`
				ChallengeAttemptsUhcKitAttackingTeamFisherman             int   `json:"challenge_attempts_uhc_kit_attacking_team_fisherman"`
				ChallengeWins8                                            int   `json:"challenge_wins_8"`
				ChallengeWins8Solo                                        int   `json:"challenge_wins_8_solo"`
				ChallengeAttemptsHalfHealthKitAttackingTeamFisherman      int   `json:"challenge_attempts_half_health_kit_attacking_team_fisherman"`
				ChallengeWins8KitAttackingTeamFisherman                   int   `json:"challenge_wins_8_kit_attacking_team_fisherman"`
				ChallengeAttemptsHalfHealthKitAdvancedSoloHunter          int   `json:"challenge_attempts_half_health_kit_advanced_solo_hunter"`
				ChallengeAttempts4KitAdvancedSoloHunter                   int   `json:"challenge_attempts_4_kit_advanced_solo_hunter"`
				ChallengeAttemptsArcherKitAdvancedSoloHunter              int   `json:"challenge_attempts_archer_kit_advanced_solo_hunter"`
				ChallengeAttemptsUhcKitAdvancedSoloHunter                 int   `json:"challenge_attempts_uhc_kit_advanced_solo_hunter"`
				ChallengeAttempts4Solo                                    int   `json:"challenge_attempts_4_solo"`
				ChallengeAttemptsKitAdvancedSoloHunter                    int   `json:"challenge_attempts_kit_advanced_solo_hunter"`
				ChallengeAttempts4                                        int   `json:"challenge_attempts_4"`
				ChallengeAttemptsPaperKitAdvancedSoloHunter               int   `json:"challenge_attempts_paper_kit_advanced_solo_hunter"`
				ChallengeAttempts5Solo                                    int   `json:"challenge_attempts_5_solo"`
				ChallengeAttemptsNoChestKitAdvancedSoloHunter             int   `json:"challenge_attempts_no_chest_kit_advanced_solo_hunter"`
				ChallengeAttempts5KitAdvancedSoloHunter                   int   `json:"challenge_attempts_5_kit_advanced_solo_hunter"`
				ChallengeAttempts5                                        int   `json:"challenge_attempts_5"`
				ChallengeAttemptsKitAttackingTeamEnderman                 int   `json:"challenge_attempts_kit_attacking_team_enderman"`
				ChallengeAttempts1Solo                                    int   `json:"challenge_attempts_1_solo"`
				ChallengeAttempts1                                        int   `json:"challenge_attempts_1"`
				GamesKitAttackingTeamEnderman                             int   `json:"games_kit_attacking_team_enderman"`
				ChallengeAttemptsPaperKitAttackingTeamEnderman            int   `json:"challenge_attempts_paper_kit_attacking_team_enderman"`
				ChallengeAttempts1KitAttackingTeamEnderman                int   `json:"challenge_attempts_1_kit_attacking_team_enderman"`
				KitAttackingTeamEndermanInventory                         struct {
					ENDERPEARL0 string `json:"ENDER_PEARL:0"`
				} `json:"kit_attacking_team_enderman_inventory"`
				AssistsKitAttackingTeamEnderman                  int    `json:"assists_kit_attacking_team_enderman"`
				SelectedPrestigeIcon                             string `json:"selected_prestige_icon"`
				AngelOfDeathLevel                                int    `json:"angel_of_death_level"`
				LongestBowKillKitAttackingTeamEnderman           int    `json:"longest_bow_kill_kit_attacking_team_enderman"`
				FastestWinKitAttackingTeamEnderman               int    `json:"fastest_win_kit_attacking_team_enderman"`
				KillstreakKitAttackingTeamEnderman               int    `json:"killstreak_kit_attacking_team_enderman"`
				BowKillsKitAttackingTeamEnderman                 int    `json:"bow_kills_kit_attacking_team_enderman"`
				WinsKitAttackingTeamEnderman                     int    `json:"wins_kit_attacking_team_enderman"`
				WinstreakKitAttackingTeamEnderman                int    `json:"winstreak_kit_attacking_team_enderman"`
				LossesLabKitAttackingTeamEnderman                int    `json:"losses_lab_kit_attacking_team_enderman"`
				DeathsLabKitAttackingTeamEnderman                int    `json:"deaths_lab_kit_attacking_team_enderman"`
				TimePlayedLabKitAttackingTeamEnderman            int    `json:"time_played_lab_kit_attacking_team_enderman"`
				ChestsOpenedLabKitAttackingTeamEnderman          int    `json:"chests_opened_lab_kit_attacking_team_enderman"`
				ActiveKitMEGARandom                              bool   `json:"activeKit_MEGA_random"`
				LongestBowKillMegaDoubles                        int    `json:"longest_bow_kill_mega_doubles"`
				ArrowsShotMegaDoubles                            int    `json:"arrows_shot_mega_doubles"`
				TimePlayedMegaDoubles                            int    `json:"time_played_mega_doubles"`
				LossesMegaDoublesNormal                          int    `json:"losses_mega_doubles_normal"`
				GamesMegaDoubles                                 int    `json:"games_mega_doubles"`
				KillsMegaDoublesNormal                           int    `json:"kills_mega_doubles_normal"`
				DeathsMegaDoublesNormal                          int    `json:"deaths_mega_doubles_normal"`
				SurvivedPlayersMegaDoubles                       int    `json:"survived_players_mega_doubles"`
				LossesMegaDoubles                                int    `json:"losses_mega_doubles"`
				MeleeKillsMegaDoubles                            int    `json:"melee_kills_mega_doubles"`
				KillsMegaDoubles                                 int    `json:"kills_mega_doubles"`
				VoidKillsMegaDoubles                             int    `json:"void_kills_mega_doubles"`
				DeathsMegaDoubles                                int    `json:"deaths_mega_doubles"`
				ChestsOpenedMegaDoubles                          int    `json:"chests_opened_mega_doubles"`
				ArrowsHitMegaDoubles                             int    `json:"arrows_hit_mega_doubles"`
				WinsMegaDoublesNormal                            int    `json:"wins_mega_doubles_normal"`
				WinstreakMegaDoubles                             int    `json:"winstreak_mega_doubles"`
				KillstreakMegaDoubles                            int    `json:"killstreak_mega_doubles"`
				WinsMegaDoubles                                  int    `json:"wins_mega_doubles"`
				LongestBowKillKitMegaMegaScout                   int    `json:"longest_bow_kill_kit_mega_mega_scout"`
				LongestBowShotMegaDoubles                        int    `json:"longest_bow_shot_mega_doubles"`
				LongestBowShotKitMegaMegaScout                   int    `json:"longest_bow_shot_kit_mega_mega_scout"`
				MostKillsGameMegaDoubles                         int    `json:"most_kills_game_mega_doubles"`
				MostKillsGameKitMegaMegaScout                    int    `json:"most_kills_game_kit_mega_mega_scout"`
				ArrowsHitKitMegaMegaScout                        int    `json:"arrows_hit_kit_mega_mega_scout"`
				TimePlayedKitMegaMegaScout                       int    `json:"time_played_kit_mega_mega_scout"`
				MeleeKillsKitMegaMegaScout                       int    `json:"melee_kills_kit_mega_mega_scout"`
				ArrowsShotKitMegaMegaScout                       int    `json:"arrows_shot_kit_mega_mega_scout"`
				ChestsOpenedKitMegaMegaScout                     int    `json:"chests_opened_kit_mega_mega_scout"`
				MeleeKillsLabKitAttackingTeamEnderman            int    `json:"melee_kills_lab_kit_attacking_team_enderman"`
				VoidKillsLabKitAttackingTeamEnderman             int    `json:"void_kills_lab_kit_attacking_team_enderman"`
				KillsLabKitAttackingTeamEnderman                 int    `json:"kills_lab_kit_attacking_team_enderman"`
				SurvivedPlayersLabKitAttackingTeamEnderman       int    `json:"survived_players_lab_kit_attacking_team_enderman"`
				FastestWinLabKitAttackingTeamEnderman            int    `json:"fastest_win_lab_kit_attacking_team_enderman"`
				GamesLabKitAttackingTeamEnderman                 int    `json:"games_lab_kit_attacking_team_enderman"`
				LabWinTntMadnessLabKitAttackingTeamEnderman      int    `json:"lab_win_tnt_madness_lab_kit_attacking_team_enderman"`
				KillstreakLabKitAttackingTeamEnderman            int    `json:"killstreak_lab_kit_attacking_team_enderman"`
				WinsLabKitAttackingTeamEnderman                  int    `json:"wins_lab_kit_attacking_team_enderman"`
				WinstreakLabKitAttackingTeamEnderman             int    `json:"winstreak_lab_kit_attacking_team_enderman"`
				ArrowsHitLabKitAttackingTeamEnderman             int    `json:"arrows_hit_lab_kit_attacking_team_enderman"`
				ArrowsShotLabKitAttackingTeamEnderman            int    `json:"arrows_shot_lab_kit_attacking_team_enderman"`
				LongestBowShotLabKitAttackingTeamEnderman        int    `json:"longest_bow_shot_lab_kit_attacking_team_enderman"`
				AssistsLabKitAttackingTeamEnderman               int    `json:"assists_lab_kit_attacking_team_enderman"`
				LabWinSlimeLabKitAttackingTeamEnderman           int    `json:"lab_win_slime_lab_kit_attacking_team_enderman"`
				ActiveKitTEAMSTourneyRandom                      bool   `json:"activeKit_TEAMS_tourney_random"`
				ActiveKitTEAMSTourney                            string `json:"activeKit_TEAMS_tourney"`
				LevelFormatted                                   string `json:"levelFormatted"`
				LongestBowShotTourneyKitAttackingTeamEnderman    int    `json:"longest_bow_shot_tourney_kit_attacking_team_enderman"`
				LongestBowShotTourney                            int    `json:"longest_bow_shot_tourney"`
				WinStreakTourney                                 int    `json:"win_streak_tourney"`
				LongestBowShotTourneyCrazytourney                int    `json:"longest_bow_shot_tourney_crazytourney"`
				TourneySwCrazySolo0WinStreak                     int    `json:"tourney_sw_crazy_solo_0_win_streak"`
				TimePlayedTourney                                int    `json:"time_played_tourney"`
				EnderpearlsThrownTourney                         int    `json:"enderpearls_thrown_tourney"`
				ArrowsShotTourneyKitAttackingTeamEnderman        int    `json:"arrows_shot_tourney_kit_attacking_team_enderman"`
				TourneySwCrazySolo0MeleeKills                    int    `json:"tourney_sw_crazy_solo_0_melee_kills"`
				TourneySwCrazySolo0EnderpearlsThrown             int    `json:"tourney_sw_crazy_solo_0_enderpearls_thrown"`
				TourneySwCrazySolo0Quits                         int    `json:"tourney_sw_crazy_solo_0_quits"`
				KillsTourneyCrazytourney                         int    `json:"kills_tourney_crazytourney"`
				LossesTourneyCrazytourney                        int    `json:"losses_tourney_crazytourney"`
				SurvivedPlayersTourneyCrazytourney               int    `json:"survived_players_tourney_crazytourney"`
				TourneySwCrazySolo0ChestsOpened                  int    `json:"tourney_sw_crazy_solo_0_chests_opened"`
				TourneySwCrazySolo0ArrowsShot                    int    `json:"tourney_sw_crazy_solo_0_arrows_shot"`
				TourneySwCrazySolo0Kills                         int    `json:"tourney_sw_crazy_solo_0_kills"`
				BlocksPlacedTourney                              int    `json:"blocks_placed_tourney"`
				LossesTourney                                    int    `json:"losses_tourney"`
				TourneySwCrazySolo0SurvivedPlayers               int    `json:"tourney_sw_crazy_solo_0_survived_players"`
				CoinsGainedTourney                               int    `json:"coins_gained_tourney"`
				ChestsOpenedTourneyKitAttackingTeamEnderman      int    `json:"chests_opened_tourney_kit_attacking_team_enderman"`
				TourneySwCrazySolo0BlocksPlaced                  int    `json:"tourney_sw_crazy_solo_0_blocks_placed"`
				KillsTourney                                     int    `json:"kills_tourney"`
				TourneySwCrazySolo0Coins                         int    `json:"tourney_sw_crazy_solo_0_coins"`
				TourneySwCrazySolo0ArrowsHit                     int    `json:"tourney_sw_crazy_solo_0_arrows_hit"`
				SurvivedPlayersTourneyKitAttackingTeamEnderman   int    `json:"survived_players_tourney_kit_attacking_team_enderman"`
				ArrowsHitTourney                                 int    `json:"arrows_hit_tourney"`
				TimePlayedTourneyKitAttackingTeamEnderman        int    `json:"time_played_tourney_kit_attacking_team_enderman"`
				ArrowsHitTourneyCrazytourney                     int    `json:"arrows_hit_tourney_crazytourney"`
				QuitsTourney                                     int    `json:"quits_tourney"`
				LossesTourneyKitAttackingTeamEnderman            int    `json:"losses_tourney_kit_attacking_team_enderman"`
				ChestsOpenedTourneyCrazytourney                  int    `json:"chests_opened_tourney_crazytourney"`
				TimePlayedTourneyCrazytourney                    int    `json:"time_played_tourney_crazytourney"`
				ArrowsShotTourneyCrazytourney                    int    `json:"arrows_shot_tourney_crazytourney"`
				ArrowsHitTourneyKitAttackingTeamEnderman         int    `json:"arrows_hit_tourney_kit_attacking_team_enderman"`
				DeathsTourneyKitAttackingTeamEnderman            int    `json:"deaths_tourney_kit_attacking_team_enderman"`
				DeathsTourney                                    int    `json:"deaths_tourney"`
				DeathsTourneyCrazytourney                        int    `json:"deaths_tourney_crazytourney"`
				MeleeKillsTourneyCrazytourney                    int    `json:"melee_kills_tourney_crazytourney"`
				TourneySwCrazySolo0TimePlayed                    int    `json:"tourney_sw_crazy_solo_0_time_played"`
				TourneySwCrazySolo0Losses                        int    `json:"tourney_sw_crazy_solo_0_losses"`
				TourneySwCrazySolo0LongestBowShot                int    `json:"tourney_sw_crazy_solo_0_longest_bow_shot"`
				SurvivedPlayersTourney                           int    `json:"survived_players_tourney"`
				ChestsOpenedTourney                              int    `json:"chests_opened_tourney"`
				TourneySwCrazySolo0Deaths                        int    `json:"tourney_sw_crazy_solo_0_deaths"`
				MeleeKillsTourney                                int    `json:"melee_kills_tourney"`
				KillsTourneyKitAttackingTeamEnderman             int    `json:"kills_tourney_kit_attacking_team_enderman"`
				MeleeKillsTourneyKitAttackingTeamEnderman        int    `json:"melee_kills_tourney_kit_attacking_team_enderman"`
				ArrowsShotTourney                                int    `json:"arrows_shot_tourney"`
				TourneySwCrazySolo0CoinsGained                   int    `json:"tourney_sw_crazy_solo_0_coins_gained"`
				LongestBowKillTourneyCrazytourney                int    `json:"longest_bow_kill_tourney_crazytourney"`
				MostKillsGameTourneyKitDefendingTeamFarmer       int    `json:"most_kills_game_tourney_kit_defending_team_farmer"`
				MostKillsGameTourney                             int    `json:"most_kills_game_tourney"`
				LongestBowKillTourneyKitDefendingTeamFarmer      int    `json:"longest_bow_kill_tourney_kit_defending_team_farmer"`
				MostKillsGameTourneyCrazytourney                 int    `json:"most_kills_game_tourney_crazytourney"`
				LongestBowKillTourney                            int    `json:"longest_bow_kill_tourney"`
				KillsTourneyKitDefendingTeamFarmer               int    `json:"kills_tourney_kit_defending_team_farmer"`
				TourneySwCrazySolo0MostKillsGame                 int    `json:"tourney_sw_crazy_solo_0_most_kills_game"`
				MeleeKillsTourneyKitDefendingTeamFarmer          int    `json:"melee_kills_tourney_kit_defending_team_farmer"`
				SurvivedPlayersTourneyKitDefendingTeamFarmer     int    `json:"survived_players_tourney_kit_defending_team_farmer"`
				LossesTourneyKitDefendingTeamFarmer              int    `json:"losses_tourney_kit_defending_team_farmer"`
				DeathsTourneyKitDefendingTeamFarmer              int    `json:"deaths_tourney_kit_defending_team_farmer"`
				TourneySwCrazySolo0LongestBowKill                int    `json:"tourney_sw_crazy_solo_0_longest_bow_kill"`
				EggThrownTourney                                 int    `json:"egg_thrown_tourney"`
				ChestsOpenedTourneyKitDefendingTeamFarmer        int    `json:"chests_opened_tourney_kit_defending_team_farmer"`
				TourneySwCrazySolo0EggThrown                     int    `json:"tourney_sw_crazy_solo_0_egg_thrown"`
				TimePlayedTourneyKitDefendingTeamFarmer          int    `json:"time_played_tourney_kit_defending_team_farmer"`
				SurvivedPlayersTourneyKitSupportingTeamEcologist int    `json:"survived_players_tourney_kit_supporting_team_ecologist"`
				DeathsTourneyKitSupportingTeamEcologist          int    `json:"deaths_tourney_kit_supporting_team_ecologist"`
				LossesTourneyKitSupportingTeamEcologist          int    `json:"losses_tourney_kit_supporting_team_ecologist"`
				TimePlayedTourneyKitSupportingTeamEcologist      int    `json:"time_played_tourney_kit_supporting_team_ecologist"`
				LongestBowShotTourneyKitDefendingTeamFrog        int    `json:"longest_bow_shot_tourney_kit_defending_team_frog"`
				LongestBowKillTourneyKitDefendingTeamFrog        int    `json:"longest_bow_kill_tourney_kit_defending_team_frog"`
				FastestWinTourneyKitDefendingTeamFrog            int    `json:"fastest_win_tourney_kit_defending_team_frog"`
				FastestWinTourneyCrazytourney                    int    `json:"fastest_win_tourney_crazytourney"`
				MostKillsGameTourneyKitDefendingTeamFrog         int    `json:"most_kills_game_tourney_kit_defending_team_frog"`
				FastestWinTourney                                int    `json:"fastest_win_tourney"`
				SurvivedPlayersTourneyKitDefendingTeamFrog       int    `json:"survived_players_tourney_kit_defending_team_frog"`
				TimePlayedTourneyKitDefendingTeamFrog            int    `json:"time_played_tourney_kit_defending_team_frog"`
				MeleeKillsTourneyKitDefendingTeamFrog            int    `json:"melee_kills_tourney_kit_defending_team_frog"`
				BlocksBrokenTourney                              int    `json:"blocks_broken_tourney"`
				TourneySwCrazySolo0Winstreak                     int    `json:"tourney_sw_crazy_solo_0_winstreak"`
				KillstreakTourney                                int    `json:"killstreak_tourney"`
				WinstreakTourney                                 int    `json:"winstreak_tourney"`
				GamesTourneyCrazytourney                         int    `json:"games_tourney_crazytourney"`
				AssistsTourneyCrazytourney                       int    `json:"assists_tourney_crazytourney"`
				TourneySwCrazySolo0Killstreak                    int    `json:"tourney_sw_crazy_solo_0_killstreak"`
				KillstreakTourneyKitDefendingTeamFrog            int    `json:"killstreak_tourney_kit_defending_team_frog"`
				GamesTourney                                     int    `json:"games_tourney"`
				TourneySwCrazySolo0Games                         int    `json:"tourney_sw_crazy_solo_0_games"`
				TourneySwCrazySolo0Assists                       int    `json:"tourney_sw_crazy_solo_0_assists"`
				TourneySwCrazySolo0Wins                          int    `json:"tourney_sw_crazy_solo_0_wins"`
				WinstreakTourneyKitDefendingTeamFrog             int    `json:"winstreak_tourney_kit_defending_team_frog"`
				TourneySwCrazySolo0BlocksBroken                  int    `json:"tourney_sw_crazy_solo_0_blocks_broken"`
				KillstreakTourneyCrazytourney                    int    `json:"killstreak_tourney_crazytourney"`
				WinstreakTourneyCrazytourney                     int    `json:"winstreak_tourney_crazytourney"`
				ArrowsShotTourneyKitDefendingTeamFrog            int    `json:"arrows_shot_tourney_kit_defending_team_frog"`
				TourneySwCrazySolo0FastestWin                    int    `json:"tourney_sw_crazy_solo_0_fastest_win"`
				ChestsOpenedTourneyKitDefendingTeamFrog          int    `json:"chests_opened_tourney_kit_defending_team_frog"`
				WinsTourneyKitDefendingTeamFrog                  int    `json:"wins_tourney_kit_defending_team_frog"`
				WinsTourneyCrazytourney                          int    `json:"wins_tourney_crazytourney"`
				GamesTourneyKitDefendingTeamFrog                 int    `json:"games_tourney_kit_defending_team_frog"`
				WinsTourney                                      int    `json:"wins_tourney"`
				ArrowsHitTourneyKitDefendingTeamFrog             int    `json:"arrows_hit_tourney_kit_defending_team_frog"`
				AssistsTourney                                   int    `json:"assists_tourney"`
				AssistsTourneyKitDefendingTeamFrog               int    `json:"assists_tourney_kit_defending_team_frog"`
				KillsTourneyKitDefendingTeamFrog                 int    `json:"kills_tourney_kit_defending_team_frog"`
				VoidKillsTourneyKitDefendingTeamFrog             int    `json:"void_kills_tourney_kit_defending_team_frog"`
				VoidKillsTourneyCrazytourney                     int    `json:"void_kills_tourney_crazytourney"`
				VoidKillsTourney                                 int    `json:"void_kills_tourney"`
				TourneySwCrazySolo0VoidKills                     int    `json:"tourney_sw_crazy_solo_0_void_kills"`
				KitBasicSoloFrogInventory                        struct {
					POTION2            string `json:"POTION:2"`
					LEATHERCHESTPLATE0 string `json:"LEATHER_CHESTPLATE:0"`
					LEATHERLEGGINGS0   string `json:"LEATHER_LEGGINGS:0"`
					SKULLITEM3         string `json:"SKULL_ITEM:3"`
					LEATHERBOOTS0      string `json:"LEATHER_BOOTS:0"`
				} `json:"kit_basic_solo_frog_inventory"`
				DeathsTourneyKitDefendingTeamFrog   int `json:"deaths_tourney_kit_defending_team_frog"`
				LossesTourneyKitDefendingTeamFrog   int `json:"losses_tourney_kit_defending_team_frog"`
				MobKillsTourneyKitDefendingTeamFrog int `json:"mob_kills_tourney_kit_defending_team_frog"`
				MobKillsTourneyCrazytourney         int `json:"mob_kills_tourney_crazytourney"`
				TourneySwCrazySolo0MobKills         int `json:"tourney_sw_crazy_solo_0_mob_kills"`
				MobKillsTourney                     int `json:"mob_kills_tourney"`
				BowKillsTourneyKitDefendingTeamFrog int `json:"bow_kills_tourney_kit_defending_team_frog"`
				BowKillsTourney                     int `json:"bow_kills_tourney"`
				BowKillsTourneyCrazytourney         int `json:"bow_kills_tourney_crazytourney"`
				TourneySwCrazySolo0BowKills         int `json:"tourney_sw_crazy_solo_0_bow_kills"`
				ItemsEnchantedTourney               int `json:"items_enchanted_tourney"`
				TourneySwCrazySolo0ItemsEnchanted   int `json:"tourney_sw_crazy_solo_0_items_enchanted"`
				MostKillsGameKitDefendingTeamFrog   int `json:"most_kills_game_kit_defending_team_frog"`
				LongestBowKillKitDefendingTeamFrog  int `json:"longest_bow_kill_kit_defending_team_frog"`
				DeathsKitDefendingTeamFrog          int `json:"deaths_kit_defending_team_frog"`
				KillsKitDefendingTeamFrog           int `json:"kills_kit_defending_team_frog"`
				LossesKitDefendingTeamFrog          int `json:"losses_kit_defending_team_frog"`
				SurvivedPlayersKitDefendingTeamFrog int `json:"survived_players_kit_defending_team_frog"`
				TimePlayedKitDefendingTeamFrog      int `json:"time_played_kit_defending_team_frog"`
				MeleeKillsKitDefendingTeamFrog      int `json:"melee_kills_kit_defending_team_frog"`
				ChestsOpenedKitDefendingTeamFrog    int `json:"chests_opened_kit_defending_team_frog"`
				FastestWinKitDefendingTeamFrog      int `json:"fastest_win_kit_defending_team_frog"`
				VoidKillsKitDefendingTeamFrog       int `json:"void_kills_kit_defending_team_frog"`
				WinstreakKitDefendingTeamFrog       int `json:"winstreak_kit_defending_team_frog"`
				GamesKitDefendingTeamFrog           int `json:"games_kit_defending_team_frog"`
				WinsKitDefendingTeamFrog            int `json:"wins_kit_defending_team_frog"`
				KillstreakKitDefendingTeamFrog      int `json:"killstreak_kit_defending_team_frog"`
				TeamFrost                           int `json:"team_frost"`
				HeadsDecentKitDefendingTeamFrog     int `json:"heads_decent_kit_defending_team_frog"`
				HeadsSalty                          int `json:"heads_salty"`
				HeadsSaltyKitDefendingTeamFrog      int `json:"heads_salty_kit_defending_team_frog"`
				HeadsMeh                            int `json:"heads_meh"`
				Heads                               int `json:"heads"`
				HeadsYuckyKitDefendingTeamFrog      int `json:"heads_yucky_kit_defending_team_frog"`
				HeadsDivine                         int `json:"heads_divine"`
				HeadsYuckySolo                      int `json:"heads_yucky_solo"`
				HeadsDivineSolo                     int `json:"heads_divine_solo"`
				HeadsMehSolo                        int `json:"heads_meh_solo"`
				HeadsYucky                          int `json:"heads_yucky"`
				HeadsKitDefendingTeamFrog           int `json:"heads_kit_defending_team_frog"`
				HeadsDecent                         int `json:"heads_decent"`
				HeadsSaltySolo                      int `json:"heads_salty_solo"`
				HeadsDivineKitDefendingTeamFrog     int `json:"heads_divine_kit_defending_team_frog"`
				HeadsMehKitDefendingTeamFrog        int `json:"heads_meh_kit_defending_team_frog"`
				HeadsSolo                           int `json:"heads_solo"`
				HeadsDecentSolo                     int `json:"heads_decent_solo"`
				HeadCollection                      struct {
					Recent []struct {
						UUID      string `json:"uuid"`
						Timestamp int64  `json:"timestamp"`
						Mode      string `json:"mode"`
						Sacrifice string `json:"sacrifice"`
					} `json:"recent"`
					Prestigious []struct {
						UUID      string `json:"uuid"`
						Timestamp int64  `json:"timestamp"`
						Mode      string `json:"mode"`
						Sacrifice string `json:"sacrifice"`
					} `json:"prestigious"`
				} `json:"head_collection"`
				AssistsKitDefendingTeamFrog                  int `json:"assists_kit_defending_team_frog"`
				LongestBowShotKitDefendingTeamFrog           int `json:"longest_bow_shot_kit_defending_team_frog"`
				HeadsTastySolo                               int `json:"heads_tasty_solo"`
				HeadsTastyKitDefendingTeamFrog               int `json:"heads_tasty_kit_defending_team_frog"`
				ArrowsShotKitDefendingTeamFrog               int `json:"arrows_shot_kit_defending_team_frog"`
				HeadsTasty                                   int `json:"heads_tasty"`
				ArrowsHitKitDefendingTeamFrog                int `json:"arrows_hit_kit_defending_team_frog"`
				HeadsEwwSolo                                 int `json:"heads_eww_solo"`
				HeadsHeavenly                                int `json:"heads_heavenly"`
				HeadsMehKitAttackingTeamEnderman             int `json:"heads_meh_kit_attacking_team_enderman"`
				HeadsHeavenlyKitAttackingTeamEnderman        int `json:"heads_heavenly_kit_attacking_team_enderman"`
				HeadsEww                                     int `json:"heads_eww"`
				HeadsSaltyKitAttackingTeamEnderman           int `json:"heads_salty_kit_attacking_team_enderman"`
				HeadsEwwKitAttackingTeamEnderman             int `json:"heads_eww_kit_attacking_team_enderman"`
				HeadsKitAttackingTeamEnderman                int `json:"heads_kit_attacking_team_enderman"`
				HeadsHeavenlySolo                            int `json:"heads_heavenly_solo"`
				MostKillsGameKitDefendingTeamFarmer          int `json:"most_kills_game_kit_defending_team_farmer"`
				LongestBowKillKitDefendingTeamFarmer         int `json:"longest_bow_kill_kit_defending_team_farmer"`
				MeleeKillsKitDefendingTeamFarmer             int `json:"melee_kills_kit_defending_team_farmer"`
				TimePlayedKitDefendingTeamFarmer             int `json:"time_played_kit_defending_team_farmer"`
				SurvivedPlayersKitDefendingTeamFarmer        int `json:"survived_players_kit_defending_team_farmer"`
				LossesKitDefendingTeamFarmer                 int `json:"losses_kit_defending_team_farmer"`
				DeathsKitDefendingTeamFarmer                 int `json:"deaths_kit_defending_team_farmer"`
				GamesKitDefendingTeamFarmer                  int `json:"games_kit_defending_team_farmer"`
				KillsKitDefendingTeamFarmer                  int `json:"kills_kit_defending_team_farmer"`
				ChestsOpenedKitDefendingTeamFarmer           int `json:"chests_opened_kit_defending_team_farmer"`
				HeadsKitAdvancedSoloFarmer                   int `json:"heads_kit_advanced_solo_farmer"`
				HeadsHeavenlyKitAdvancedSoloFarmer           int `json:"heads_heavenly_kit_advanced_solo_farmer"`
				LabWinLuckyBlocksLabKitAttackingTeamEnderman int `json:"lab_win_lucky_blocks_lab_kit_attacking_team_enderman"`
				Privategames                                 struct {
					EnableTeleportMayhem  bool   `json:"enable_teleport_mayhem"`
					EnableMaxKitsAndPerks bool   `json:"enable_max_kits_and_perks"`
					ChestSwords           string `json:"chest_swords"`
					LowGravity            bool   `json:"low_gravity"`
					Dragons               string `json:"dragons"`
					EnableLegacyItems     bool   `json:"enable_legacy_items"`
					Speed                 string `json:"speed"`
					HealthBuff            string `json:"health_buff"`
					ChestArmour           string `json:"chest_armour"`
					ChestBows             string `json:"chest_bows"`
					OneHitOneKill         bool   `json:"one_hit_one_kill"`
					EnableNightTime       bool   `json:"enable_night_time"`
					NoKits                bool   `json:"no_kits"`
				} `json:"privategames"`
				HeadsSucculentKitAttackingTeamEnderman   int   `json:"heads_succulent_kit_attacking_team_enderman"`
				HeadsSucculentSolo                       int   `json:"heads_succulent_solo"`
				HeadsSucculent                           int   `json:"heads_succulent"`
				HeadsYuckyKitAttackingTeamEnderman       int   `json:"heads_yucky_kit_attacking_team_enderman"`
				MobKillsKitAttackingTeamEnderman         int   `json:"mob_kills_kit_attacking_team_enderman"`
				BowKillsKitDefendingTeamFrog             int   `json:"bow_kills_kit_defending_team_frog"`
				SkywarsHalloweenBoxes                    int   `json:"skywars_halloween_boxes"`
				ChestsOpenedLabKitDefendingTeamFarmer    int   `json:"chests_opened_lab_kit_defending_team_farmer"`
				DeathsLabKitDefendingTeamFarmer          int   `json:"deaths_lab_kit_defending_team_farmer"`
				LossesLabKitDefendingTeamFarmer          int   `json:"losses_lab_kit_defending_team_farmer"`
				SurvivedPlayersLabKitDefendingTeamFarmer int   `json:"survived_players_lab_kit_defending_team_farmer"`
				TimePlayedLabKitDefendingTeamFarmer      int   `json:"time_played_lab_kit_defending_team_farmer"`
				ChestsOpenedKitMegaMegaDefault           int   `json:"chests_opened_kit_mega_mega_default"`
				TimePlayedKitMegaMegaDefault             int   `json:"time_played_kit_mega_mega_default"`
				FreeEventKeySkywarsChristmasBoxes2019    bool  `json:"free_event_key_skywars_christmas_boxes_2019"`
				SkywarsChristmasBoxes                    int   `json:"skywars_christmas_boxes"`
				BowKillsLab                              int   `json:"bow_kills_lab"`
				BowKillsLabKitAttackingTeamEnderman      int   `json:"bow_kills_lab_kit_attacking_team_enderman"`
				BowKillsLabSolo                          int   `json:"bow_kills_lab_solo"`
				LabWinLuckyBlocksLabTeam                 int   `json:"lab_win_lucky_blocks_lab_team"`
				Opals                                    int   `json:"opals"`
				Shard                                    int   `json:"shard"`
				ShardKitAttackingTeamEnderman            int   `json:"shard_kit_attacking_team_enderman"`
				ShardSolo                                int   `json:"shard_solo"`
				ShardTeam                                int   `json:"shard_team"`
				LastTourneyAd                            int64 `json:"lastTourneyAd"`
			} `json:"SkyWars"`
			TrueCombat struct {
				CrazywallsKillsSoloChaos            int      `json:"crazywalls_kills_solo_chaos"`
				ItemsEnchanted                      int      `json:"items_enchanted"`
				ArrowsShot                          int      `json:"arrows_shot"`
				Coins                               int      `json:"coins"`
				Kills                               int      `json:"kills"`
				Wins                                int      `json:"wins"`
				SurvivedPlayers                     int      `json:"survived_players"`
				CrazywallsWinsSoloChaos             int      `json:"crazywalls_wins_solo_chaos"`
				WinStreak                           int      `json:"win_streak"`
				Games                               int      `json:"games"`
				CrazywallsLossesSoloChaos           int      `json:"crazywalls_losses_solo_chaos"`
				CrazywallsGamesSoloChaos            int      `json:"crazywalls_games_solo_chaos"`
				Deaths                              int      `json:"deaths"`
				Losses                              int      `json:"losses"`
				CrazywallsDeathsSoloChaos           int      `json:"crazywalls_deaths_solo_chaos"`
				CrazywallsDeathsSolo                int      `json:"crazywalls_deaths_solo"`
				CrazywallsGamesSolo                 int      `json:"crazywalls_games_solo"`
				CrazywallsLossesSolo                int      `json:"crazywalls_losses_solo"`
				CrazywallsKillsSolo                 int      `json:"crazywalls_kills_solo"`
				CrazywallsWinsSolo                  int      `json:"crazywalls_wins_solo"`
				GoldenSkulls                        int      `json:"golden_skulls"`
				GiantZombie                         int      `json:"giant_zombie"`
				SoloChaosBountyHunter               int      `json:"solo_chaos_bounty_hunter"`
				LiveCombat                          bool     `json:"live_combat"`
				CrazywallsKillsWeeklyBSolo          int      `json:"crazywalls_kills_weekly_b_solo"`
				CrazywallsKillsMonthlyBSolo         int      `json:"crazywalls_kills_monthly_b_solo"`
				SkullsGathered                      int      `json:"skulls_gathered"`
				TeamAdrenaline                      int      `json:"team_adrenaline"`
				SoloChaosTombRobber                 int      `json:"solo_chaos_tomb_robber"`
				GiantZombieLegendaries              int      `json:"giant_zombie_legendaries"`
				SoloChaosBusinessman                int      `json:"solo_chaos_businessman"`
				TeamBlazingArrow                    int      `json:"team_blazing_arrow"`
				KitBasicChaosAlchemist              int      `json:"kit_basic_chaos_alchemist"`
				ActiveKitSoloChaos                  string   `json:"activeKit_Solo_chaos"`
				CrazywallsKillsWeeklyBSoloChaos     int      `json:"crazywalls_kills_weekly_b_solo_chaos"`
				CrazywallsKillsMonthlyBSoloChaos    int      `json:"crazywalls_kills_monthly_b_solo_chaos"`
				Packages                            []string `json:"packages"`
				VotesGarden                         int      `json:"votes_Garden"`
				GiantZombieRares                    int      `json:"giant_zombie_rares"`
				SoloChaosRusher                     int      `json:"solo_chaos_rusher"`
				GoldDust                            int      `json:"gold_dust"`
				SoloChaosBerserk                    int      `json:"solo_chaos_berserk"`
				SoloChaosSuperLuck                  int      `json:"solo_chaos_super_luck"`
				TeamGroupHeal                       int      `json:"team_group_heal"`
				SoloBlazingArrow                    int      `json:"solo_blazing_arrow"`
				SoloLuckyDrops                      int      `json:"solo_lucky_drops"`
				SoloSpeedMining                     int      `json:"solo_speed_mining"`
				TeamRusher                          int      `json:"team_rusher"`
				SoloRusher                          int      `json:"solo_rusher"`
				SoloBountyHunter                    int      `json:"solo_bounty_hunter"`
				TeamSwiftness                       int      `json:"team_swiftness"`
				TeamBerserk                         int      `json:"team_berserk"`
				SoloAdrenaline                      int      `json:"solo_adrenaline"`
				KillsWeeklyB                        int      `json:"kills_weekly_b"`
				KillsMonthlyB                       int      `json:"kills_monthly_b"`
				TeamSmartMining                     int      `json:"team_smart_mining"`
				SoloVampirism                       int      `json:"solo_vampirism"`
				TeamBountyHunter                    int      `json:"team_bounty_hunter"`
				TeamPrecision                       int      `json:"team_precision"`
				KitBasicChaosWeaponsmith            int      `json:"kit_basic_chaos_weaponsmith"`
				KillsWeeklyA                        int      `json:"kills_weekly_a"`
				CrazywallsKillsMonthlyASoloChaos    int      `json:"crazywalls_kills_monthly_a_solo_chaos"`
				CrazywallsKillsWeeklyASoloChaos     int      `json:"crazywalls_kills_weekly_a_solo_chaos"`
				KillsMonthlyA                       int      `json:"kills_monthly_a"`
				TeamLastStand                       int      `json:"team_last_stand"`
				SoloPrecision                       int      `json:"solo_precision"`
				SoloKnowledge                       int      `json:"solo_knowledge"`
				KitBasicChaosArmorer                int      `json:"kit_basic_chaos_armorer"`
				CraftingLuckyEnchantedBookSharpness int      `json:"crafting_lucky_enchanted_book_sharpness"`
				SoloChaosBlazingArrow               int      `json:"solo_chaos_blazing_arrow"`
				SoloSwiftness                       int      `json:"solo_swiftness"`
				CraftingLuckyBrawler                int      `json:"crafting_lucky_brawler"`
				SoloChaosKnowledge                  int      `json:"solo_chaos_knowledge"`
				SoloChaosAdrenaline                 int      `json:"solo_chaos_adrenaline"`
				ActiveKitTeamChaos                  string   `json:"activeKit_Team_chaos"`
				CrazywallsWinsTeamChaos             int      `json:"crazywalls_wins_team_chaos"`
				CrazywallsKillsWeeklyBTeamChaos     int      `json:"crazywalls_kills_weekly_b_team_chaos"`
				CrazywallsKillsMonthlyBTeamChaos    int      `json:"crazywalls_kills_monthly_b_team_chaos"`
				CrazywallsKillsTeamChaos            int      `json:"crazywalls_kills_team_chaos"`
				ActiveKitSolo                       string   `json:"activeKit_Solo"`
				CrazywallsKillsWeeklyASolo          int      `json:"crazywalls_kills_weekly_a_solo"`
				TeamAnnoyOMite                      int      `json:"team_annoy-o-mite"`
				SoloChaosVampirism                  int      `json:"solo_chaos_vampirism"`
				SoloChaosSmartyPants                int      `json:"solo_chaos_smarty_pants"`
				ArrowsHit                           int      `json:"arrows_hit"`
				CrazywallsDeathsTeamChaos           int      `json:"crazywalls_deaths_team_chaos"`
				CrazywallsGamesTeamChaos            int      `json:"crazywalls_games_team_chaos"`
				CrazywallsKillsMonthlyATeamChaos    int      `json:"crazywalls_kills_monthly_a_team_chaos"`
				CrazywallsKillsWeeklyATeamChaos     int      `json:"crazywalls_kills_weekly_a_team_chaos"`
				CrazywallsLossesTeamChaos           int      `json:"crazywalls_losses_team_chaos"`
			} `json:"TrueCombat"`
			SuperSmash struct {
				SmashLevel       int    `json:"smashLevel"`
				LastLevelTHEBULK int    `json:"lastLevel_THE_BULK"`
				ActiveClass      string `json:"active_class"`
				Coins            int    `json:"coins"`
				WinStreak        int    `json:"win_streak"`
				DamageDealtTeams int    `json:"damage_dealt_teams"`
				ClassStats       struct {
					Tinman struct {
						Kills          int `json:"kills"`
						HomingMissiles struct {
							KillsTeams        int `json:"kills_teams"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							DamageDealt       int `json:"damage_dealt"`
							Kills             int `json:"kills"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							KillsNormal       int `json:"kills_normal"`
							Smasher           int `json:"smasher"`
							SmasherNormal     int `json:"smasher_normal"`
							SmasherTeams      int `json:"smasher_teams"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							Smasher2V2        int `json:"smasher_2v2"`
							Kills2V2          int `json:"kills_2v2"`
						} `json:"homing_missiles"`
						ThrowCake struct {
							SmashedTeams int `json:"smashed_teams"`
							Smashed      int `json:"smashed"`
						} `json:"throw_cake"`
						LaserCannon struct {
							SmasherTeams      int `json:"smasher_teams"`
							Smasher           int `json:"smasher"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							DamageDealt       int `json:"damage_dealt"`
							Kills             int `json:"kills"`
							KillsTeams        int `json:"kills_teams"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							Kills2V2          int `json:"kills_2v2"`
							Smasher2V2        int `json:"smasher_2v2"`
						} `json:"laser_cannon"`
						Overload struct {
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							DamageDealt       int `json:"damage_dealt"`
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							Kills2V2          int `json:"kills_2v2"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
						} `json:"overload"`
						GamesTeams  int `json:"games_teams"`
						WinsTeams   int `json:"wins_teams"`
						DeathsTeams int `json:"deaths_teams"`
						KillsTeams  int `json:"kills_teams"`
						RocketPunch struct {
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Smasher           int `json:"smasher"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
						} `json:"rocket_punch"`
						SmashedTeams     int `json:"smashed_teams"`
						SmasherTeams     int `json:"smasher_teams"`
						Smashed          int `json:"smashed"`
						WinStreakTeams   int `json:"win_streak_teams"`
						Games            int `json:"games"`
						Smasher          int `json:"smasher"`
						DamageDealtTeams int `json:"damage_dealt_teams"`
						WinStreak        int `json:"win_streak"`
						Wins             int `json:"wins"`
						Deaths           int `json:"deaths"`
						DamageDealt      int `json:"damage_dealt"`
						GamesNormal      int `json:"games_normal"`
						Losses           int `json:"losses"`
						EggBazooka       struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"egg_bazooka"`
						KillsNormal   int `json:"kills_normal"`
						LossesNormal  int `json:"losses_normal"`
						SmasherNormal int `json:"smasher_normal"`
						SmashedNormal int `json:"smashed_normal"`
						SeismicSlam   struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"seismic_slam"`
						DeathsNormal      int `json:"deaths_normal"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						WinStreakNormal   int `json:"win_streak_normal"`
						WinsNormal        int `json:"wins_normal"`
						Melee             struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							Smashed           int `json:"smashed"`
							SmashedTeams      int `json:"smashed_teams"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							Smasher           int `json:"smasher"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Kills             int `json:"kills"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							SmashedNormal     int `json:"smashed_normal"`
						} `json:"melee"`
						LossesTeams    int `json:"losses_teams"`
						ForceLightning struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"force_lightning"`
						Wins2V2      int `json:"wins_2v2"`
						Smashed2V2   int `json:"smashed_2v2"`
						Kills2V2     int `json:"kills_2v2"`
						WinStreak2V2 int `json:"win_streak_2v2"`
						CakeStorm    struct {
							Smashed2V2 int `json:"smashed_2v2"`
							Smashed    int `json:"smashed"`
						} `json:"cake_storm"`
						Deaths2V2      int `json:"deaths_2v2"`
						Smasher2V2     int `json:"smasher_2v2"`
						Games2V2       int `json:"games_2v2"`
						DamageDealt2V2 int `json:"damage_dealt_2v2"`
						Losses2V2      int `json:"losses_2v2"`
						Batarang       struct {
							Smashed2V2 int `json:"smashed_2v2"`
							Smashed    int `json:"smashed"`
						} `json:"batarang"`
						Telepunch struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"telepunch"`
					} `json:"TINMAN"`
					GeneralCluck struct {
						Bazooka struct {
							Kills             int `json:"kills"`
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							KillsNormal       int `json:"kills_normal"`
							Smasher           int `json:"smasher"`
						} `json:"bazooka"`
						WinStreakNormal int `json:"win_streak_normal"`
						EggBazooka      struct {
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
							SmasherNormal     int `json:"smasher_normal"`
							KillsNormal       int `json:"kills_normal"`
							Smasher           int `json:"smasher"`
							Kills             int `json:"kills"`
						} `json:"egg_bazooka"`
						Reinforcements struct {
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
						} `json:"reinforcements"`
						Melee struct {
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
							KillsNormal       int `json:"kills_normal"`
							SmashedNormal     int `json:"smashed_normal"`
							Smasher           int `json:"smasher"`
							Kills             int `json:"kills"`
							Smashed           int `json:"smashed"`
							SmasherNormal     int `json:"smasher_normal"`
						} `json:"melee"`
						DeathsNormal      int `json:"deaths_normal"`
						GamesNormal       int `json:"games_normal"`
						Games             int `json:"games"`
						Deaths            int `json:"deaths"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						WinStreak         int `json:"win_streak"`
						DamageDealt       int `json:"damage_dealt"`
						Wins              int `json:"wins"`
						Smasher           int `json:"smasher"`
						SmasherNormal     int `json:"smasher_normal"`
						KillsNormal       int `json:"kills_normal"`
						Kills             int `json:"kills"`
						WinsNormal        int `json:"wins_normal"`
						Losses            int `json:"losses"`
						LossesNormal      int `json:"losses_normal"`
						Smashed           int `json:"smashed"`
						SmashedNormal     int `json:"smashed_normal"`
					} `json:"GENERAL_CLUCK"`
					Skullfire struct {
						FlamingDesertEagle struct {
							Kills             int `json:"kills"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Smasher           int `json:"smasher"`
							DamageDealt       int `json:"damage_dealt"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
							Kills2V2          int `json:"kills_2v2"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
						} `json:"flaming_desert_eagle"`
						Kills int `json:"kills"`
						Melee struct {
							SmashedNormal     int `json:"smashed_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							Smashed           int `json:"smashed"`
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
						} `json:"melee"`
						Grenade struct {
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsNormal       int `json:"kills_normal"`
							Kills             int `json:"kills"`
							Smasher           int `json:"smasher"`
							SmasherNormal     int `json:"smasher_normal"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							Smashed           int `json:"smashed"`
							SmashedNormal     int `json:"smashed_normal"`
						} `json:"grenade"`
						Smasher           int `json:"smasher"`
						DeathsNormal      int `json:"deaths_normal"`
						Deaths            int `json:"deaths"`
						Smashed           int `json:"smashed"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						DesertEagle       struct {
							Kills             int `json:"kills"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Smashed           int `json:"smashed"`
							SmashedNormal     int `json:"smashed_normal"`
							Smasher           int `json:"smasher"`
							SmasherTeams      int `json:"smasher_teams"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
							Kills2V2          int `json:"kills_2v2"`
							Smasher2V2        int `json:"smasher_2v2"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
						} `json:"desert_eagle"`
						Games           int `json:"games"`
						WinStreakNormal int `json:"win_streak_normal"`
						DamageDealt     int `json:"damage_dealt"`
						WinsNormal      int `json:"wins_normal"`
						GamesNormal     int `json:"games_normal"`
						WinStreak       int `json:"win_streak"`
						SmasherNormal   int `json:"smasher_normal"`
						Wins            int `json:"wins"`
						SmashedNormal   int `json:"smashed_normal"`
						KillsNormal     int `json:"kills_normal"`
						LossesNormal    int `json:"losses_normal"`
						Losses          int `json:"losses"`
						Batarang        struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"batarang"`
						SmasherTeams int `json:"smasher_teams"`
						DeathsTeams  int `json:"deaths_teams"`
						LaserCannon  struct {
							Smashed      int `json:"smashed"`
							SmashedTeams int `json:"smashed_teams"`
							Smashed2V2   int `json:"smashed_2v2"`
						} `json:"laser_cannon"`
						KillsTeams       int `json:"kills_teams"`
						SmashedTeams     int `json:"smashed_teams"`
						WinStreakTeams   int `json:"win_streak_teams"`
						WinsTeams        int `json:"wins_teams"`
						GamesTeams       int `json:"games_teams"`
						DamageDealtTeams int `json:"damage_dealt_teams"`
						SwingPin         struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"swing_pin"`
						KiBlast struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"ki_blast"`
						Kills2V2       int `json:"kills_2v2"`
						Wins2V2        int `json:"wins_2v2"`
						Smasher2V2     int `json:"smasher_2v2"`
						Games2V2       int `json:"games_2v2"`
						WinStreak2V2   int `json:"win_streak_2v2"`
						DamageDealt2V2 int `json:"damage_dealt_2v2"`
						Deaths2V2      int `json:"deaths_2v2"`
						Smashed2V2     int `json:"smashed_2v2"`
						Losses2V2      int `json:"losses_2v2"`
					} `json:"SKULLFIRE"`
					Marauder struct {
						Melee struct {
							Smasher           int `json:"smasher"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							SmashedTeams      int `json:"smashed_teams"`
							Smashed           int `json:"smashed"`
							SmasherTeams      int `json:"smasher_teams"`
							KillsTeams        int `json:"kills_teams"`
						} `json:"melee"`
						ForceLightning struct {
							KillsNormal       int `json:"kills_normal"`
							Kills             int `json:"kills"`
							Smasher           int `json:"smasher"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
							SmasherNormal     int `json:"smasher_normal"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							Smashed           int `json:"smashed"`
							SmashedNormal     int `json:"smashed_normal"`
						} `json:"force_lightning"`
						WinsNormal  int `json:"wins_normal"`
						MonsterMash struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"monster_mash"`
						GamesNormal     int `json:"games_normal"`
						Smashed         int `json:"smashed"`
						SmasherNormal   int `json:"smasher_normal"`
						Smasher         int `json:"smasher"`
						DamageDealt     int `json:"damage_dealt"`
						WinStreakNormal int `json:"win_streak_normal"`
						ForcePull       struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							Kills             int `json:"kills"`
							Smasher           int `json:"smasher"`
							KillsTeams        int `json:"kills_teams"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							SmasherTeams      int `json:"smasher_teams"`
							SmasherNormal     int `json:"smasher_normal"`
							KillsNormal       int `json:"kills_normal"`
							Smashed           int `json:"smashed"`
							SmashedNormal     int `json:"smashed_normal"`
						} `json:"force_pull"`
						KillsNormal       int `json:"kills_normal"`
						Games             int `json:"games"`
						DeathsNormal      int `json:"deaths_normal"`
						SmashedNormal     int `json:"smashed_normal"`
						Wins              int `json:"wins"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						Kills             int `json:"kills"`
						Deaths            int `json:"deaths"`
						WinStreak         int `json:"win_streak"`
						KillsTeams        int `json:"kills_teams"`
						WinsTeams         int `json:"wins_teams"`
						CakeStorm         struct {
							Smashed      int `json:"smashed"`
							SmashedTeams int `json:"smashed_teams"`
						} `json:"cake_storm"`
						DamageDealtTeams int `json:"damage_dealt_teams"`
						WinStreakTeams   int `json:"win_streak_teams"`
						SmashedTeams     int `json:"smashed_teams"`
						DeathsTeams      int `json:"deaths_teams"`
						SmasherTeams     int `json:"smasher_teams"`
						GamesTeams       int `json:"games_teams"`
						Botmubile        struct {
							SmashedTeams  int `json:"smashed_teams"`
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"botmubile"`
						Batarang struct {
							SmashedTeams  int `json:"smashed_teams"`
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"batarang"`
						Frostbolt struct {
							Smashed      int `json:"smashed"`
							SmashedTeams int `json:"smashed_teams"`
						} `json:"frostbolt"`
						Losses       int `json:"losses"`
						LossesTeams  int `json:"losses_teams"`
						LossesNormal int `json:"losses_normal"`
						LaserCannon  struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"laser_cannon"`
						HomingMissiles struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"homing_missiles"`
					} `json:"MARAUDER"`
					Frosty struct {
						Melee struct {
							DamageDealt       int `json:"damage_dealt"`
							KillsTeams        int `json:"kills_teams"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							Kills             int `json:"kills"`
							SmasherTeams      int `json:"smasher_teams"`
							Smasher           int `json:"smasher"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
						} `json:"melee"`
						DeathsTeams int `json:"deaths_teams"`
						Frostbolt   struct {
							DamageDealt       int `json:"damage_dealt"`
							Kills             int `json:"kills"`
							Smasher           int `json:"smasher"`
							SmasherTeams      int `json:"smasher_teams"`
							KillsTeams        int `json:"kills_teams"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
						} `json:"frostbolt"`
						Losses           int `json:"losses"`
						DamageDealtTeams int `json:"damage_dealt_teams"`
						Smasher          int `json:"smasher"`
						SmashedTeams     int `json:"smashed_teams"`
						FreezingBreath   struct {
							KillsTeams        int `json:"kills_teams"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							Kills             int `json:"kills"`
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							KillsNormal       int `json:"kills_normal"`
						} `json:"freezing_breath"`
						Smashed int `json:"smashed"`
						Games   int `json:"games"`
						Bazooka struct {
							SmashedTeams int `json:"smashed_teams"`
							Smashed      int `json:"smashed"`
						} `json:"bazooka"`
						GamesTeams     int `json:"games_teams"`
						LossesTeams    int `json:"losses_teams"`
						KillsTeams     int `json:"kills_teams"`
						DamageDealt    int `json:"damage_dealt"`
						SmasherTeams   int `json:"smasher_teams"`
						Kills          int `json:"kills"`
						Deaths         int `json:"deaths"`
						SmashedNormal  int `json:"smashed_normal"`
						HomingMissiles struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"homing_missiles"`
						DeathsNormal      int `json:"deaths_normal"`
						WinStreak         int `json:"win_streak"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						GamesNormal       int `json:"games_normal"`
						Wins              int `json:"wins"`
						WinsNormal        int `json:"wins_normal"`
						SmasherNormal     int `json:"smasher_normal"`
						KillsNormal       int `json:"kills_normal"`
						WinStreakNormal   int `json:"win_streak_normal"`
					} `json:"FROSTY"`
					TheBulk struct {
						SeismicSlam struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Smasher           int `json:"smasher"`
							SmashedNormal     int `json:"smashed_normal"`
							Smashed           int `json:"smashed"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							Kills2V2          int `json:"kills_2v2"`
							Smasher2V2        int `json:"smasher_2v2"`
							KillsTeams        int `json:"kills_teams"`
							SmasherTeams      int `json:"smasher_teams"`
						} `json:"seismic_slam"`
						Boulder struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Kills             int `json:"kills"`
							Smasher           int `json:"smasher"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							Smasher2V2        int `json:"smasher_2v2"`
							Kills2V2          int `json:"kills_2v2"`
							KillsTeams        int `json:"kills_teams"`
							SmasherTeams      int `json:"smasher_teams"`
						} `json:"boulder"`
						MonsterCharge struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
							SmasherTeams      int `json:"smasher_teams"`
							Smasher           int `json:"smasher"`
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Smasher2V2        int `json:"smasher_2v2"`
							Kills2V2          int `json:"kills_2v2"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							Smashed           int `json:"smashed"`
							SmashedNormal     int `json:"smashed_normal"`
						} `json:"monster_charge"`
						Melee struct {
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
							Kills             int `json:"kills"`
							SmasherNormal     int `json:"smasher_normal"`
							KillsNormal       int `json:"kills_normal"`
							Smasher           int `json:"smasher"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							SmashedNormal     int `json:"smashed_normal"`
							Smashed           int `json:"smashed"`
							Smasher2V2        int `json:"smasher_2v2"`
							Kills2V2          int `json:"kills_2v2"`
							KillsTeams        int `json:"kills_teams"`
							SmasherTeams      int `json:"smasher_teams"`
							SmashedTeams      int `json:"smashed_teams"`
							Smashed2V2        int `json:"smashed_2v2"`
						} `json:"melee"`
						Games             int `json:"games"`
						DeathsNormal      int `json:"deaths_normal"`
						LossesNormal      int `json:"losses_normal"`
						DamageDealt       int `json:"damage_dealt"`
						Losses            int `json:"losses"`
						Deaths            int `json:"deaths"`
						GamesNormal       int `json:"games_normal"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						WinStreak         int `json:"win_streak"`
						Wins              int `json:"wins"`
						SmasherNormal     int `json:"smasher_normal"`
						Batarang          struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
							Smashed2V2    int `json:"smashed_2v2"`
						} `json:"batarang"`
						WinsNormal  int `json:"wins_normal"`
						Smasher     int `json:"smasher"`
						Kills       int `json:"kills"`
						MonsterMash struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							KillsNormal       int `json:"kills_normal"`
							Smasher           int `json:"smasher"`
							Kills             int `json:"kills"`
							SmasherNormal     int `json:"smasher_normal"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							Smasher2V2        int `json:"smasher_2v2"`
							Kills2V2          int `json:"kills_2v2"`
							SmasherTeams      int `json:"smasher_teams"`
							Smashed           int `json:"smashed"`
							SmashedNormal     int `json:"smashed_normal"`
							Smashed2V2        int `json:"smashed_2v2"`
						} `json:"monster_mash"`
						Smashed          int `json:"smashed"`
						SmashedNormal    int `json:"smashed_normal"`
						WinStreakNormal  int `json:"win_streak_normal"`
						KillsNormal      int `json:"kills_normal"`
						SmashedTeams     int `json:"smashed_teams"`
						WinsTeams        int `json:"wins_teams"`
						SmasherTeams     int `json:"smasher_teams"`
						DamageDealtTeams int `json:"damage_dealt_teams"`
						KillsTeams       int `json:"kills_teams"`
						ForceLightning   struct {
							SmashedTeams  int `json:"smashed_teams"`
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
							Smashed2V2    int `json:"smashed_2v2"`
						} `json:"force_lightning"`
						DeathsTeams    int `json:"deaths_teams"`
						GamesTeams     int `json:"games_teams"`
						WinStreakTeams int `json:"win_streak_teams"`
						Reinforcements struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"reinforcements"`
						KiBlast struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"ki_blast"`
						LaserCannon struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
							SmashedTeams  int `json:"smashed_teams"`
							Smashed2V2    int `json:"smashed_2v2"`
						} `json:"laser_cannon"`
						Botmubile struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"botmubile"`
						ThrowCake struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
							Smashed2V2    int `json:"smashed_2v2"`
						} `json:"throw_cake"`
						ForcePull struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
							Smashed2V2    int `json:"smashed_2v2"`
							SmashedTeams  int `json:"smashed_teams"`
						} `json:"force_pull"`
						SpooderBuddies struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"spooder_buddies"`
						Overload struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
							Smashed2V2    int `json:"smashed_2v2"`
						} `json:"overload"`
						DesertEagle struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"desert_eagle"`
						DamageDealt2V2 int `json:"damage_dealt_2v2"`
						Smasher2V2     int `json:"smasher_2v2"`
						Smashed2V2     int `json:"smashed_2v2"`
						Wins2V2        int `json:"wins_2v2"`
						WinStreak2V2   int `json:"win_streak_2v2"`
						Kills2V2       int `json:"kills_2v2"`
						Games2V2       int `json:"games_2v2"`
						Deaths2V2      int `json:"deaths_2v2"`
						SwingPin       struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"swing_pin"`
						Frostbolt struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"frostbolt"`
						HomingMissiles struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"homing_missiles"`
						SpiderKick struct {
							Smashed       int `json:"smashed"`
							Smashed2V2    int `json:"smashed_2v2"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"spider_kick"`
						LossesTeams int `json:"losses_teams"`
						Bazooka     struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
							Smashed2V2    int `json:"smashed_2v2"`
						} `json:"bazooka"`
						StaticLaser struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"static_laser"`
						FlamingDesertEagle struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"flaming_desert_eagle"`
						NotchedBow struct {
							Smashed      int `json:"smashed"`
							SmashedTeams int `json:"smashed_teams"`
						} `json:"notched_bow"`
						FriendWins       int `json:"friend_wins"`
						FriendWinsNormal int `json:"friend_wins_normal"`
						ChargedBeam      struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
							Smashed2V2    int `json:"smashed_2v2"`
						} `json:"charged_beam"`
						GrapplingHook struct {
							Smashed    int `json:"smashed"`
							Smashed2V2 int `json:"smashed_2v2"`
						} `json:"grappling_hook"`
						WebShot struct {
							Smashed2V2 int `json:"smashed_2v2"`
							Smashed    int `json:"smashed"`
						} `json:"web_shot"`
						Defecake struct {
							Smashed2V2 int `json:"smashed_2v2"`
							Smashed    int `json:"smashed"`
						} `json:"defecake"`
						EggBazooka struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"egg_bazooka"`
						FriendLosses        int `json:"friend_losses"`
						FriendLossesNormal  int `json:"friend_losses_normal"`
						OneVOneWinsNormal   int `json:"one_v_one_wins_normal"`
						OneVOneWins         int `json:"one_v_one_wins"`
						OneVOneLosses       int `json:"one_v_one_losses"`
						OneVOneLossesNormal int `json:"one_v_one_losses_normal"`
					} `json:"THE_BULK"`
					DuskCrawler struct {
						VoidSlash struct {
							DamageDealt       int `json:"damage_dealt"`
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							Kills3V3          int `json:"kills_3v3"`
							Smasher3V3        int `json:"smasher_3v3"`
							Smasher           int `json:"smasher"`
							DamageDealt3V3    int `json:"damage_dealt_3v3"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
							SmasherTeams      int `json:"smasher_teams"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							SmasherNormal     int `json:"smasher_normal"`
							Kills2V2          int `json:"kills_2v2"`
							Smasher2V2        int `json:"smasher_2v2"`
						} `json:"void_slash"`
						ForcePull struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"force_pull"`
						LossesNormal int `json:"losses_normal"`
						Melee        struct {
							SmashedNormal     int `json:"smashed_normal"`
							Smashed           int `json:"smashed"`
							DamageDealt       int `json:"damage_dealt"`
							Smasher3V3        int `json:"smasher_3v3"`
							Kills             int `json:"kills"`
							Smasher           int `json:"smasher"`
							Kills3V3          int `json:"kills_3v3"`
							DamageDealt3V3    int `json:"damage_dealt_3v3"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							Kills2V2          int `json:"kills_2v2"`
							Smasher2V2        int `json:"smasher_2v2"`
							SmasherTeams      int `json:"smasher_teams"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Smashed2V2        int `json:"smashed_2v2"`
						} `json:"melee"`
						Losses            int `json:"losses"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						Deaths            int `json:"deaths"`
						DamageDealt       int `json:"damage_dealt"`
						Teleboom          struct {
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
							DamageDealt3V3    int `json:"damage_dealt_3v3"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
							Kills             int `json:"kills"`
							SmasherTeams      int `json:"smasher_teams"`
							Smasher           int `json:"smasher"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Kills2V2          int `json:"kills_2v2"`
							Smasher2V2        int `json:"smasher_2v2"`
						} `json:"teleboom"`
						Games         int `json:"games"`
						DeathsNormal  int `json:"deaths_normal"`
						KillsNormal   int `json:"kills_normal"`
						Kills         int `json:"kills"`
						SmashedNormal int `json:"smashed_normal"`
						Smashed       int `json:"smashed"`
						GamesNormal   int `json:"games_normal"`
						Losses3V3     int `json:"losses_3v3"`
						Telepunch     struct {
							Kills3V3          int `json:"kills_3v3"`
							Smasher           int `json:"smasher"`
							DamageDealt       int `json:"damage_dealt"`
							Smasher3V3        int `json:"smasher_3v3"`
							DamageDealt3V3    int `json:"damage_dealt_3v3"`
							Kills             int `json:"kills"`
							KillsTeams        int `json:"kills_teams"`
							SmasherTeams      int `json:"smasher_teams"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							Smasher2V2        int `json:"smasher_2v2"`
							DamageDealt2V2    int `json:"damage_dealt_2v2"`
							Kills2V2          int `json:"kills_2v2"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							SmasherNormal     int `json:"smasher_normal"`
						} `json:"telepunch"`
						SpooderBuddies struct {
							Smashed    int `json:"smashed"`
							Smashed3V3 int `json:"smashed_3v3"`
							Smashed2V2 int `json:"smashed_2v2"`
						} `json:"spooder_buddies"`
						Games3V3       int `json:"games_3v3"`
						Smashed3V3     int `json:"smashed_3v3"`
						Smasher        int `json:"smasher"`
						Deaths3V3      int `json:"deaths_3v3"`
						Smasher3V3     int `json:"smasher_3v3"`
						Kills3V3       int `json:"kills_3v3"`
						DamageDealt3V3 int `json:"damage_dealt_3v3"`
						SmasherTeams   int `json:"smasher_teams"`
						WinsTeams      int `json:"wins_teams"`
						LaserCannon    struct {
							SmashedTeams  int `json:"smashed_teams"`
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
							Smashed2V2    int `json:"smashed_2v2"`
						} `json:"laser_cannon"`
						DamageDealtTeams int `json:"damage_dealt_teams"`
						KillsTeams       int `json:"kills_teams"`
						DeathsTeams      int `json:"deaths_teams"`
						WinStreak        int `json:"win_streak"`
						SmashedTeams     int `json:"smashed_teams"`
						GamesTeams       int `json:"games_teams"`
						WinStreakTeams   int `json:"win_streak_teams"`
						Wins             int `json:"wins"`
						Deaths2V2        int `json:"deaths_2v2"`
						Smasher2V2       int `json:"smasher_2v2"`
						Kills2V2         int `json:"kills_2v2"`
						DamageDealt2V2   int `json:"damage_dealt_2v2"`
						Losses2V2        int `json:"losses_2v2"`
						Smashed2V2       int `json:"smashed_2v2"`
						Games2V2         int `json:"games_2v2"`
						Wins2V2          int `json:"wins_2v2"`
						WinStreak2V2     int `json:"win_streak_2v2"`
						Boulder          struct {
							SmashedTeams  int `json:"smashed_teams"`
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"boulder"`
						WinsNormal      int `json:"wins_normal"`
						WinStreakNormal int `json:"win_streak_normal"`
						SmasherNormal   int `json:"smasher_normal"`
						Botmubile       struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"botmubile"`
						Frostbolt struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"frostbolt"`
						CakeStorm struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"cake_storm"`
						Reinforcements struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"reinforcements"`
						Batarang struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
							Smashed2V2    int `json:"smashed_2v2"`
						} `json:"batarang"`
						ForceLightning struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
							Smashed2V2    int `json:"smashed_2v2"`
						} `json:"force_lightning"`
						SeismicSlam struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
							Smashed2V2    int `json:"smashed_2v2"`
						} `json:"seismic_slam"`
						ThrowCake struct {
							Smashed2V2    int `json:"smashed_2v2"`
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"throw_cake"`
						HomingMissiles struct {
							Smashed2V2 int `json:"smashed_2v2"`
							Smashed    int `json:"smashed"`
						} `json:"homing_missiles"`
						KameBeam struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"kame_beam"`
						DesertEagle struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"desert_eagle"`
						Bazooka struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"bazooka"`
						KiBlast struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"ki_blast"`
						RocketPunch struct {
							Smashed int `json:"smashed"`
						} `json:"rocket_punch"`
						MonsterCharge struct {
							Smashed    int `json:"smashed"`
							Smashed2V2 int `json:"smashed_2v2"`
						} `json:"monster_charge"`
						SwingPin struct {
							Smashed    int `json:"smashed"`
							Smashed2V2 int `json:"smashed_2v2"`
						} `json:"swing_pin"`
						ChargedBeam struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"charged_beam"`
					} `json:"DUSK_CRAWLER"`
					Botmun struct {
						Batarang struct {
							Kills             int `json:"kills"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							DamageDealt       int `json:"damage_dealt"`
							Smasher           int `json:"smasher"`
							KillsNormal       int `json:"kills_normal"`
							KillsTeams        int `json:"kills_teams"`
							SmasherTeams      int `json:"smasher_teams"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
						} `json:"batarang"`
						Deaths          int `json:"deaths"`
						GamesNormal     int `json:"games_normal"`
						Games           int `json:"games"`
						WinStreak       int `json:"win_streak"`
						WinStreakNormal int `json:"win_streak_normal"`
						Kills           int `json:"kills"`
						Botmubile       struct {
							DamageDealtNormal int `json:"damage_dealt_normal"`
							Kills             int `json:"kills"`
							DamageDealt       int `json:"damage_dealt"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
							SmasherNormal     int `json:"smasher_normal"`
							Smasher           int `json:"smasher"`
							SmashedNormal     int `json:"smashed_normal"`
							Smashed           int `json:"smashed"`
							SmasherTeams      int `json:"smasher_teams"`
						} `json:"botmubile"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						Wins              int `json:"wins"`
						Smasher           int `json:"smasher"`
						Melee             struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							Smashed           int `json:"smashed"`
							SmashedNormal     int `json:"smashed_normal"`
							Smasher           int `json:"smasher"`
							SmasherTeams      int `json:"smasher_teams"`
							Kills             int `json:"kills"`
							KillsTeams        int `json:"kills_teams"`
							SmashedTeams      int `json:"smashed_teams"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
						} `json:"melee"`
						KillsNormal   int `json:"kills_normal"`
						SmasherNormal int `json:"smasher_normal"`
						DamageDealt   int `json:"damage_dealt"`
						DeathsNormal  int `json:"deaths_normal"`
						WinsNormal    int `json:"wins_normal"`
						WinsTeams     int `json:"wins_teams"`
						ForcePull     struct {
							SmashedTeams  int `json:"smashed_teams"`
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"force_pull"`
						KillsTeams       int `json:"kills_teams"`
						WinStreakTeams   int `json:"win_streak_teams"`
						DeathsTeams      int `json:"deaths_teams"`
						GamesTeams       int `json:"games_teams"`
						Smashed          int `json:"smashed"`
						SmasherTeams     int `json:"smasher_teams"`
						SmashedTeams     int `json:"smashed_teams"`
						DamageDealtTeams int `json:"damage_dealt_teams"`
						SmashedNormal    int `json:"smashed_normal"`
						LaserCannon      struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
							SmashedTeams  int `json:"smashed_teams"`
						} `json:"laser_cannon"`
						EggBazooka struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"egg_bazooka"`
						Bazooka struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
							SmashedTeams  int `json:"smashed_teams"`
						} `json:"bazooka"`
						Losses       int `json:"losses"`
						LossesNormal int `json:"losses_normal"`
						Frostbolt    struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"frostbolt"`
						ForceLightning struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"force_lightning"`
						DesertEagle struct {
							Smashed      int `json:"smashed"`
							SmashedTeams int `json:"smashed_teams"`
						} `json:"desert_eagle"`
						LossesTeams int `json:"losses_teams"`
						Boulder     struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"boulder"`
						Reinforcements struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"reinforcements"`
						FlamingDesertEagle struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"flaming_desert_eagle"`
						SpooderBuddies struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"spooder_buddies"`
						MonsterCharge struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"monster_charge"`
						KiBlast struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"ki_blast"`
						ThrowCake struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"throw_cake"`
						SpiderKick struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"spider_kick"`
						SwingPin struct {
							Smashed      int `json:"smashed"`
							SmashedTeams int `json:"smashed_teams"`
						} `json:"swing_pin"`
						SeismicSlam struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"seismic_slam"`
						HomingMissiles struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"homing_missiles"`
						FreezingBreath struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"freezing_breath"`
						OnionCannon struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"onion_cannon"`
					} `json:"BOTMUN"`
					Goku struct {
						Smasher int `json:"smasher"`
						KiBlast struct {
							Kills             int `json:"kills"`
							SmasherNormal     int `json:"smasher_normal"`
							DamageDealt       int `json:"damage_dealt"`
							Smasher           int `json:"smasher"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							SmashedNormal     int `json:"smashed_normal"`
							Smashed           int `json:"smashed"`
							SmasherTeams      int `json:"smasher_teams"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
						} `json:"ki_blast"`
						Wins  int `json:"wins"`
						Melee struct {
							DamageDealt       int `json:"damage_dealt"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							SmashedNormal     int `json:"smashed_normal"`
							Smashed           int `json:"smashed"`
							Kills             int `json:"kills"`
							SmasherNormal     int `json:"smasher_normal"`
							Smasher           int `json:"smasher"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							SmasherTeams      int `json:"smasher_teams"`
							KillsTeams        int `json:"kills_teams"`
						} `json:"melee"`
						Kills    int `json:"kills"`
						KameBeam struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
						} `json:"kame_beam"`
						Deaths            int `json:"deaths"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						WinStreak         int `json:"win_streak"`
						DamageDealt       int `json:"damage_dealt"`
						WinsNormal        int `json:"wins_normal"`
						Games             int `json:"games"`
						SmasherNormal     int `json:"smasher_normal"`
						WinStreakNormal   int `json:"win_streak_normal"`
						GamesNormal       int `json:"games_normal"`
						DeathsNormal      int `json:"deaths_normal"`
						KillsNormal       int `json:"kills_normal"`
						SmashedNormal     int `json:"smashed_normal"`
						Smashed           int `json:"smashed"`
						Reinforcements    struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"reinforcements"`
						Batarang struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"batarang"`
						Bazooka struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"bazooka"`
						GamesTeams       int `json:"games_teams"`
						SmasherTeams     int `json:"smasher_teams"`
						DamageDealtTeams int `json:"damage_dealt_teams"`
						WinStreakTeams   int `json:"win_streak_teams"`
						WinsTeams        int `json:"wins_teams"`
						KillsTeams       int `json:"kills_teams"`
						DeathsTeams      int `json:"deaths_teams"`
						CakeStorm        struct {
							SmashedTeams int `json:"smashed_teams"`
							Smashed      int `json:"smashed"`
						} `json:"cake_storm"`
						SmashedTeams   int `json:"smashed_teams"`
						SpooderBuddies struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"spooder_buddies"`
						Botmubile struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"botmubile"`
					} `json:"GOKU"`
					Spoderman struct {
						Melee struct {
							DamageDealt       int `json:"damage_dealt"`
							SmasherNormal     int `json:"smasher_normal"`
							KillsNormal       int `json:"kills_normal"`
							Smasher           int `json:"smasher"`
							Kills             int `json:"kills"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							KillsTeams        int `json:"kills_teams"`
							SmasherTeams      int `json:"smasher_teams"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
						} `json:"melee"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						SpooderBuddies    struct {
							SmasherNormal     int `json:"smasher_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
							Smasher           int `json:"smasher"`
							KillsNormal       int `json:"kills_normal"`
							Kills             int `json:"kills"`
							SmashedNormal     int `json:"smashed_normal"`
							Smashed           int `json:"smashed"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
							SmasherTeams      int `json:"smasher_teams"`
						} `json:"spooder_buddies"`
						WebShot struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Smasher           int `json:"smasher"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
							SmasherTeams      int `json:"smasher_teams"`
						} `json:"web_shot"`
						DeathsNormal int `json:"deaths_normal"`
						DamageDealt  int `json:"damage_dealt"`
						GamesNormal  int `json:"games_normal"`
						Batarang     struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
							SmashedTeams  int `json:"smashed_teams"`
						} `json:"batarang"`
						SpiderKick struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealtTeams  int `json:"damage_dealt_teams"`
							KillsTeams        int `json:"kills_teams"`
							Kills             int `json:"kills"`
						} `json:"spider_kick"`
						Smashed         int `json:"smashed"`
						WinStreakNormal int `json:"win_streak_normal"`
						WinsNormal      int `json:"wins_normal"`
						Wins            int `json:"wins"`
						Smasher         int `json:"smasher"`
						Kills           int `json:"kills"`
						WinStreak       int `json:"win_streak"`
						Games           int `json:"games"`
						Deaths          int `json:"deaths"`
						KillsNormal     int `json:"kills_normal"`
						SmashedNormal   int `json:"smashed_normal"`
						SmasherNormal   int `json:"smasher_normal"`
						WallClimber     struct {
							DamageDealt      int `json:"damage_dealt"`
							DamageDealtTeams int `json:"damage_dealt_teams"`
						} `json:"wall_climber"`
						DamageDealtTeams int `json:"damage_dealt_teams"`
						SmashedTeams     int `json:"smashed_teams"`
						GamesTeams       int `json:"games_teams"`
						WinsTeams        int `json:"wins_teams"`
						KillsTeams       int `json:"kills_teams"`
						ForceLightning   struct {
							SmashedTeams int `json:"smashed_teams"`
							Smashed      int `json:"smashed"`
						} `json:"force_lightning"`
						DeathsTeams    int `json:"deaths_teams"`
						SmasherTeams   int `json:"smasher_teams"`
						WinStreakTeams int `json:"win_streak_teams"`
						Bazooka        struct {
							Smashed      int `json:"smashed"`
							SmashedTeams int `json:"smashed_teams"`
						} `json:"bazooka"`
						LaserCannon struct {
							SmashedTeams int `json:"smashed_teams"`
							Smashed      int `json:"smashed"`
						} `json:"laser_cannon"`
					} `json:"SPODERMAN"`
					CakeMonster struct {
						CakeStorm struct {
							SmasherNormal     int `json:"smasher_normal"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							Kills             int `json:"kills"`
							DamageDealt       int `json:"damage_dealt"`
							Smasher           int `json:"smasher"`
						} `json:"cake_storm"`
						Defecake struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
						} `json:"defecake"`
						WinStreak int `json:"win_streak"`
						SwingPin  struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							Kills             int `json:"kills"`
							Smasher           int `json:"smasher"`
							KillsNormal       int `json:"kills_normal"`
						} `json:"swing_pin"`
						Smashed           int `json:"smashed"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						ThrowCake         struct {
							DamageDealt       int `json:"damage_dealt"`
							Smasher           int `json:"smasher"`
							KillsNormal       int `json:"kills_normal"`
							Kills             int `json:"kills"`
							SmasherNormal     int `json:"smasher_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
						} `json:"throw_cake"`
						KillsNormal     int `json:"kills_normal"`
						WinStreakNormal int `json:"win_streak_normal"`
						Melee           struct {
							DamageDealtNormal int `json:"damage_dealt_normal"`
							KillsNormal       int `json:"kills_normal"`
							Smasher           int `json:"smasher"`
							SmasherNormal     int `json:"smasher_normal"`
							Kills             int `json:"kills"`
							DamageDealt       int `json:"damage_dealt"`
							SmashedNormal     int `json:"smashed_normal"`
							Smashed           int `json:"smashed"`
						} `json:"melee"`
						Smasher    int `json:"smasher"`
						SpiderKick struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"spider_kick"`
						GamesNormal    int `json:"games_normal"`
						Deaths         int `json:"deaths"`
						WinsNormal     int `json:"wins_normal"`
						SmasherNormal  int `json:"smasher_normal"`
						SmashedNormal  int `json:"smashed_normal"`
						DamageDealt    int `json:"damage_dealt"`
						Kills          int `json:"kills"`
						Games          int `json:"games"`
						Wins           int `json:"wins"`
						DeathsNormal   int `json:"deaths_normal"`
						ForceLightning struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"force_lightning"`
						MonsterCharge struct {
							SmashedNormal int `json:"smashed_normal"`
							Smashed       int `json:"smashed"`
						} `json:"monster_charge"`
						LaserCannon struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"laser_cannon"`
					} `json:"CAKE_MONSTER"`
					ShoopDaWhoop struct {
						DamageDealt int `json:"damage_dealt"`
						Melee       struct {
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
						} `json:"melee"`
						StaticLaser struct {
							Smasher           int `json:"smasher"`
							Kills             int `json:"kills"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							SmasherNormal     int `json:"smasher_normal"`
							DamageDealt       int `json:"damage_dealt"`
							KillsNormal       int `json:"kills_normal"`
						} `json:"static_laser"`
						FirMaLazer struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
						} `json:"fir_ma_lazer"`
						ChargedBeam struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							KillsNormal       int `json:"kills_normal"`
							Kills             int `json:"kills"`
						} `json:"charged_beam"`
						WinsNormal        int `json:"wins_normal"`
						GamesNormal       int `json:"games_normal"`
						Smasher           int `json:"smasher"`
						KillsNormal       int `json:"kills_normal"`
						WinStreakNormal   int `json:"win_streak_normal"`
						Games             int `json:"games"`
						DeathsNormal      int `json:"deaths_normal"`
						SmasherNormal     int `json:"smasher_normal"`
						Wins              int `json:"wins"`
						WinStreak         int `json:"win_streak"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						Deaths            int `json:"deaths"`
						Kills             int `json:"kills"`
					} `json:"SHOOP_DA_WHOOP"`
					Sanic struct {
						Dash struct {
							DamageDealt       int `json:"damage_dealt"`
							Smasher           int `json:"smasher"`
							Smashed           int `json:"smashed"`
							SmashedNormal     int `json:"smashed_normal"`
							KillsNormal       int `json:"kills_normal"`
							Kills             int `json:"kills"`
							SmasherNormal     int `json:"smasher_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
						} `json:"dash"`
						Kills int `json:"kills"`
						Boom  struct {
							Kills             int `json:"kills"`
							KillsNormal       int `json:"kills_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
						} `json:"boom"`
						Melee struct {
							Kills             int `json:"kills"`
							Smasher           int `json:"smasher"`
							SmasherNormal     int `json:"smasher_normal"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
							DamageDealt       int `json:"damage_dealt"`
							KillsNormal       int `json:"kills_normal"`
						} `json:"melee"`
						DesertEagle struct {
							Smashed       int `json:"smashed"`
							SmashedNormal int `json:"smashed_normal"`
						} `json:"desert_eagle"`
						SmashedNormal     int `json:"smashed_normal"`
						SmasherNormal     int `json:"smasher_normal"`
						Deaths            int `json:"deaths"`
						KillsNormal       int `json:"kills_normal"`
						DamageDealtNormal int `json:"damage_dealt_normal"`
						OnionCannon       struct {
							DamageDealt       int `json:"damage_dealt"`
							DamageDealtNormal int `json:"damage_dealt_normal"`
						} `json:"onion_cannon"`
						DeathsNormal int `json:"deaths_normal"`
						Smashed      int `json:"smashed"`
						Losses       int `json:"losses"`
						GamesNormal  int `json:"games_normal"`
						LossesNormal int `json:"losses_normal"`
						DamageDealt  int `json:"damage_dealt"`
						Games        int `json:"games"`
						Smasher      int `json:"smasher"`
					} `json:"SANIC"`
				} `json:"class_stats"`
				Smasher                     int   `json:"smasher"`
				KillsMonthlyB               int   `json:"kills_monthly_b"`
				GamesMonthlyB               int   `json:"games_monthly_b"`
				Kills                       int   `json:"kills"`
				Wins                        int   `json:"wins"`
				GamesWeeklyB                int   `json:"games_weekly_b"`
				KillsWeeklyB                int   `json:"kills_weekly_b"`
				WinsTeams                   int   `json:"wins_teams"`
				DamageDealt                 int   `json:"damage_dealt"`
				SmashedTeams                int   `json:"smashed_teams"`
				KillsTeams                  int   `json:"kills_teams"`
				SmasherTeams                int   `json:"smasher_teams"`
				Smashed                     int   `json:"smashed"`
				DeathsTeams                 int   `json:"deaths_teams"`
				Deaths                      int   `json:"deaths"`
				WinsWeeklyB                 int   `json:"wins_weekly_b"`
				GamesTeams                  int   `json:"games_teams"`
				Games                       int   `json:"games"`
				WinsMonthlyB                int   `json:"wins_monthly_b"`
				XpTINMAN                    int   `json:"xp_TINMAN"`
				LastLevelTINMAN             int   `json:"lastLevel_TINMAN"`
				SmasherNormal               int   `json:"smasher_normal"`
				LossesNormal                int   `json:"losses_normal"`
				LossesWeeklyB               int   `json:"losses_weekly_b"`
				LossesMonthlyB              int   `json:"losses_monthly_b"`
				Losses                      int   `json:"losses"`
				SmashedNormal               int   `json:"smashed_normal"`
				DeathsNormal                int   `json:"deaths_normal"`
				GamesNormal                 int   `json:"games_normal"`
				KillsNormal                 int   `json:"kills_normal"`
				DamageDealtNormal           int   `json:"damage_dealt_normal"`
				SmashLevelTotal             int   `json:"smash_level_total"`
				WinsNormal                  int   `json:"wins_normal"`
				XpGENERALCLUCK              int   `json:"xp_GENERAL_CLUCK"`
				LastLevelGENERALCLUCK       int   `json:"lastLevel_GENERAL_CLUCK"`
				XpSKULLFIRE                 int   `json:"xp_SKULLFIRE"`
				LastLevelSKULLFIRE          int   `json:"lastLevel_SKULLFIRE"`
				ExpBoosterPurchases10Plays  int   `json:"expBooster_purchases_10_plays"`
				XpMARAUDER                  int   `json:"xp_MARAUDER"`
				LastLevelMARAUDER           int   `json:"lastLevel_MARAUDER"`
				Quits                       int   `json:"quits"`
				GamesWeeklyA                int   `json:"games_weekly_a"`
				KillsWeeklyA                int   `json:"kills_weekly_a"`
				WinsWeeklyA                 int   `json:"wins_weekly_a"`
				VotesStrawberryTowers       int   `json:"votes_Strawberry Towers"`
				LossesTeams                 int   `json:"losses_teams"`
				LossesWeeklyA               int   `json:"losses_weekly_a"`
				XpFROSTY                    int   `json:"xp_FROSTY"`
				LastLevelFROSTY             int   `json:"lastLevel_FROSTY"`
				ExpiredBooster              bool  `json:"expired_booster"`
				LastLevelDUSKCRAWLER        int   `json:"lastLevel_DUSK_CRAWLER"`
				XpBOTMUN                    int   `json:"xp_BOTMUN"`
				LastLevelBOTMUN             int   `json:"lastLevel_BOTMUN"`
				VotesCubed                  int   `json:"votes_Cubed"`
				ExpBoosterPurchases30Plays  int   `json:"expBooster_purchases_30_plays"`
				Smasher3V3                  int   `json:"smasher_3v3"`
				DamageDealt3V3              int   `json:"damage_dealt_3v3"`
				Kills3V3                    int   `json:"kills_3v3"`
				Deaths3V3                   int   `json:"deaths_3v3"`
				Games3V3                    int   `json:"games_3v3"`
				Smashed3V3                  int   `json:"smashed_3v3"`
				Losses3V3                   int   `json:"losses_3v3"`
				XpDUSKCRAWLER               int   `json:"xp_DUSK_CRAWLER"`
				VotesGunmetal               int   `json:"votes_Gunmetal"`
				XpGOKU                      int   `json:"xp_GOKU"`
				LastLevelGOKU               int   `json:"lastLevel_GOKU"`
				Deaths2V2                   int   `json:"deaths_2v2"`
				Kills2V2                    int   `json:"kills_2v2"`
				Games2V2                    int   `json:"games_2v2"`
				Smashed2V2                  int   `json:"smashed_2v2"`
				Smasher2V2                  int   `json:"smasher_2v2"`
				Losses2V2                   int   `json:"losses_2v2"`
				DamageDealt2V2              int   `json:"damage_dealt_2v2"`
				Wins2V2                     int   `json:"wins_2v2"`
				ExpBoosterPurchases100Plays int   `json:"expBooster_purchases_100_plays"`
				XpSPODERMAN                 int   `json:"xp_SPODERMAN"`
				LastLevelSPODERMAN          int   `json:"lastLevel_SPODERMAN"`
				VotesGrove                  int   `json:"votes_Grove"`
				VotesToybox                 int   `json:"votes_Toybox"`
				VotesColorClash             int   `json:"votes_Color Clash"`
				XpTHEBULK                   int   `json:"xp_THE_BULK"`
				PgDUSKCRAWLER               int   `json:"pg_DUSK_CRAWLER"`
				PgBOTMUN                    int   `json:"pg_BOTMUN"`
				XpCAKEMONSTER               int   `json:"xp_CAKE_MONSTER"`
				LastLevelCAKEMONSTER        int   `json:"lastLevel_CAKE_MONSTER"`
				PgTHEBULK                   int   `json:"pg_THE_BULK"`
				GamesMonthlyA               int   `json:"games_monthly_a"`
				KillsMonthlyA               int   `json:"kills_monthly_a"`
				WinsMonthlyA                int   `json:"wins_monthly_a"`
				XpSHOOPDAWHOOP              int   `json:"xp_SHOOP_DA_WHOOP"`
				LastLevelSHOOPDAWHOOP       int   `json:"lastLevel_SHOOP_DA_WHOOP"`
				LossesMonthlyA              int   `json:"losses_monthly_a"`
				LastLevelSANIC              int   `json:"lastLevel_SANIC"`
				FRIENDSGamesDay             int   `json:"FRIENDS_gamesDay"`
				FRIENDSFirstGame            int64 `json:"FRIENDS_firstGame"`
				FriendLosses                int   `json:"friend_losses"`
				FriendWins                  int   `json:"friend_wins"`
				FriendWinsNormal            int   `json:"friend_wins_normal"`
				Assists                     int   `json:"assists"`
				AssistsNormal               int   `json:"assists_normal"`
				FriendLossesNormal          int   `json:"friend_losses_normal"`
				HeroLevelBoosterActive      struct {
					Key        string `json:"key"`
					Multiplier int    `json:"multiplier"`
					Value      int    `json:"value"`
					Plays      int    `json:"plays"`
				} `json:"hero_level_booster_active"`
				ONEVJUANFirstGame   int64 `json:"ONE_V_JUAN_firstGame"`
				ONEVJUANGamesDay    int   `json:"ONE_V_JUAN_gamesDay"`
				OneVOneLossesNormal int   `json:"one_v_one_losses_normal"`
				OneVOneLosses       int   `json:"one_v_one_losses"`
				OneVOneWins         int   `json:"one_v_one_wins"`
				OneVOneWinsNormal   int   `json:"one_v_one_wins_normal"`
			} `json:"SuperSmash"`
			SpeedUHC struct {
				ActiveKitINSANE                      string   `json:"activeKit_INSANE"`
				Coins                                int      `json:"coins"`
				WinstreakKitBasicInsaneDefault       int      `json:"winstreak_kit_basic_insane_default"`
				GamesSolo                            int      `json:"games_solo"`
				KillsMonthlyA                        int      `json:"kills_monthly_a"`
				Wins                                 int      `json:"wins"`
				WinsMasteryWildSpecialist            int      `json:"wins_mastery_wild_specialist"`
				WinstreakInsane                      int      `json:"winstreak_insane"`
				KillstreakInsane                     int      `json:"killstreak_insane"`
				SurvivedPlayersSolo                  int      `json:"survived_players_solo"`
				WinsInsane                           int      `json:"wins_insane"`
				KillsMasteryWildSpecialist           int      `json:"kills_mastery_wild_specialist"`
				KillsWeeklyB                         int      `json:"kills_weekly_b"`
				Kills                                int      `json:"kills"`
				KillstreakSolo                       int      `json:"killstreak_solo"`
				TearsGathered                        int      `json:"tears_gathered"`
				Killstreak                           int      `json:"killstreak"`
				KillstreakKitBasicInsaneDefault      int      `json:"killstreak_kit_basic_insane_default"`
				ItemsEnchanted                       int      `json:"items_enchanted"`
				WinstreakSolo                        int      `json:"winstreak_solo"`
				WinsKitBasicInsaneDefault            int      `json:"wins_kit_basic_insane_default"`
				Winstreak                            int      `json:"winstreak"`
				BlocksBroken                         int      `json:"blocks_broken"`
				Tears                                int      `json:"tears"`
				KillsSolo                            int      `json:"kills_solo"`
				KillsInsane                          int      `json:"kills_insane"`
				WinsSolo                             int      `json:"wins_solo"`
				Games                                int      `json:"games"`
				SurvivedPlayers                      int      `json:"survived_players"`
				KillsSoloInsane                      int      `json:"kills_solo_insane"`
				SurvivedPlayersInsane                int      `json:"survived_players_insane"`
				GamesInsane                          int      `json:"games_insane"`
				GamesKitBasicInsaneDefault           int      `json:"games_kit_basic_insane_default"`
				SurvivedPlayersKitBasicInsaneDefault int      `json:"survived_players_kit_basic_insane_default"`
				WinsSoloInsane                       int      `json:"wins_solo_insane"`
				BlocksPlaced                         int      `json:"blocks_placed"`
				KillsKitBasicInsaneDefault           int      `json:"kills_kit_basic_insane_default"`
				SurvivedPlayersNormal                int      `json:"survived_players_normal"`
				WinsNormal                           int      `json:"wins_normal"`
				KillsSoloNormal                      int      `json:"kills_solo_normal"`
				GamesNormal                          int      `json:"games_normal"`
				WinsKitBasicNormalDefault            int      `json:"wins_kit_basic_normal_default"`
				WinsSoloNormal                       int      `json:"wins_solo_normal"`
				KillstreakKitBasicNormalDefault      int      `json:"killstreak_kit_basic_normal_default"`
				WinstreakKitBasicNormalDefault       int      `json:"winstreak_kit_basic_normal_default"`
				WinstreakNormal                      int      `json:"winstreak_normal"`
				SurvivedPlayersKitBasicNormalDefault int      `json:"survived_players_kit_basic_normal_default"`
				GamesKitBasicNormalDefault           int      `json:"games_kit_basic_normal_default"`
				KillstreakNormal                     int      `json:"killstreak_normal"`
				KillsNormal                          int      `json:"kills_normal"`
				KillsKitBasicNormalDefault           int      `json:"kills_kit_basic_normal_default"`
				HighestWinstreak                     int      `json:"highestWinstreak"`
				HighestKillstreak                    int      `json:"highestKillstreak"`
				WinStreak                            int      `json:"win_streak"`
				Losses                               int      `json:"losses"`
				DeathsSolo                           int      `json:"deaths_solo"`
				DeathsNormal                         int      `json:"deaths_normal"`
				LossesSoloNormal                     int      `json:"losses_solo_normal"`
				Quits                                int      `json:"quits"`
				DeathsSoloNormal                     int      `json:"deaths_solo_normal"`
				Deaths                               int      `json:"deaths"`
				DeathsKitBasicNormalDefault          int      `json:"deaths_kit_basic_normal_default"`
				DeathsMasteryWildSpecialist          int      `json:"deaths_mastery_wild_specialist"`
				LossesMasteryWildSpecialist          int      `json:"losses_mastery_wild_specialist"`
				LossesKitBasicNormalDefault          int      `json:"losses_kit_basic_normal_default"`
				LossesNormal                         int      `json:"losses_normal"`
				LossesSolo                           int      `json:"losses_solo"`
				ArrowsShot                           int      `json:"arrows_shot"`
				ArrowsHit                            int      `json:"arrows_hit"`
				Packages                             []string `json:"packages"`
				TearWellUses                         int      `json:"tearWellUses"`
				FoundCOMMON                          int      `json:"found_COMMON"`
				DeathsSoloInsane                     int      `json:"deaths_solo_insane"`
				LossesInsane                         int      `json:"losses_insane"`
				LossesSoloInsane                     int      `json:"losses_solo_insane"`
				DeathsInsane                         int      `json:"deaths_insane"`
				DeathsKitBasicInsaneDefault          int      `json:"deaths_kit_basic_insane_default"`
				LossesKitBasicInsaneDefault          int      `json:"losses_kit_basic_insane_default"`
				FirstJoinLobby                       bool     `json:"firstJoinLobby"`
				FirstJoinLobbyInt                    int      `json:"firstJoinLobbyInt"`
				NormalArrowRecovery                  int      `json:"normal_arrow_recovery"`
				KillsKitBasicInsaneMiner             int      `json:"kills_kit_basic_insane_miner"`
				KillsWeeklyA                         int      `json:"kills_weekly_a"`
				GamesKitBasicInsaneMiner             int      `json:"games_kit_basic_insane_miner"`
				KillstreakKitBasicInsaneMiner        int      `json:"killstreak_kit_basic_insane_miner"`
				SurvivedPlayersKitBasicInsaneMiner   int      `json:"survived_players_kit_basic_insane_miner"`
				WinstreakKitBasicInsaneMiner         int      `json:"winstreak_kit_basic_insane_miner"`
				WinsKitBasicInsaneMiner              int      `json:"wins_kit_basic_insane_miner"`
				AssistsKitBasicInsaneMiner           int      `json:"assists_kit_basic_insane_miner"`
				AssistsInsane                        int      `json:"assists_insane"`
				AssistsSolo                          int      `json:"assists_solo"`
				Assists                              int      `json:"assists"`
				Salt                                 int      `json:"salt"`
				ExtraWheels                          int      `json:"extra_wheels"`
				NormalMonsterTamer                   int      `json:"normal_monster_tamer"`
				FoundRARE                            int      `json:"found_RARE"`
				InsaneSwimmingChampion               int      `json:"insane_swimming_champion"`
				InsaneKnowledge                      int      `json:"insane_knowledge"`
				NormalSwimmingChampion               int      `json:"normal_swimming_champion"`
				InsanePortalProtection               int      `json:"insane_portal_protection"`
				MasteryFortune                       int      `json:"mastery_fortune"`
				FoundLEGENDARY                       int      `json:"found_LEGENDARY"`
				ActiveMasterPerk                     string   `json:"activeMasterPerk"`
				GamesKitBasicInsaneArcher            int      `json:"games_kit_basic_insane_archer"`
				WinsMasteryFortune                   int      `json:"wins_mastery_fortune"`
				KillstreakKitBasicInsaneArcher       int      `json:"killstreak_kit_basic_insane_archer"`
				KillsMasteryFortune                  int      `json:"kills_mastery_fortune"`
				WinsKitBasicInsaneArcher             int      `json:"wins_kit_basic_insane_archer"`
				KillsKitBasicInsaneArcher            int      `json:"kills_kit_basic_insane_archer"`
				SurvivedPlayersKitBasicInsaneArcher  int      `json:"survived_players_kit_basic_insane_archer"`
				WinstreakKitBasicInsaneArcher        int      `json:"winstreak_kit_basic_insane_archer"`
				ActiveKitNORMAL                      string   `json:"activeKit_NORMAL"`
				LossesMasteryFortune                 int      `json:"losses_mastery_fortune"`
				DeathsKitBasicNormalScout            int      `json:"deaths_kit_basic_normal_scout"`
				SurvivedPlayersKitBasicNormalScout   int      `json:"survived_players_kit_basic_normal_scout"`
				GamesKitBasicNormalScout             int      `json:"games_kit_basic_normal_scout"`
				DeathsMasteryFortune                 int      `json:"deaths_mastery_fortune"`
				KillsKitBasicNormalScout             int      `json:"kills_kit_basic_normal_scout"`
				LossesKitBasicNormalScout            int      `json:"losses_kit_basic_normal_scout"`
				KillstreakKitBasicNormalScout        int      `json:"killstreak_kit_basic_normal_scout"`
				WinsKitBasicNormalScout              int      `json:"wins_kit_basic_normal_scout"`
				WinstreakKitBasicNormalScout         int      `json:"winstreak_kit_basic_normal_scout"`
				LossesKitBasicInsaneHealer           int      `json:"losses_kit_basic_insane_healer"`
				DeathsKitBasicInsaneHealer           int      `json:"deaths_kit_basic_insane_healer"`
				SurvivedPlayersKitBasicInsaneHealer  int      `json:"survived_players_kit_basic_insane_healer"`
				WinsKitBasicInsaneHealer             int      `json:"wins_kit_basic_insane_healer"`
				KillsKitBasicInsaneHealer            int      `json:"kills_kit_basic_insane_healer"`
				GamesKitBasicInsaneHealer            int      `json:"games_kit_basic_insane_healer"`
				WinstreakKitBasicInsaneHealer        int      `json:"winstreak_kit_basic_insane_healer"`
				KillstreakKitBasicInsaneHealer       int      `json:"killstreak_kit_basic_insane_healer"`
				DeathsTeam                           int      `json:"deaths_team"`
				DeathsTeamNormal                     int      `json:"deaths_team_normal"`
				KillsTeamNormal                      int      `json:"kills_team_normal"`
				SurvivedPlayersTeam                  int      `json:"survived_players_team"`
				KillsTeam                            int      `json:"kills_team"`
				LossesTeam                           int      `json:"losses_team"`
				LossesTeamNormal                     int      `json:"losses_team_normal"`
				KillsMonthlyB                        int      `json:"kills_monthly_b"`
				GamesTeam                            int      `json:"games_team"`
				DeathsTeamInsane                     int      `json:"deaths_team_insane"`
				GamesKitBasicInsaneScout             int      `json:"games_kit_basic_insane_scout"`
				WinstreakKitBasicInsaneScout         int      `json:"winstreak_kit_basic_insane_scout"`
				DeathsKitBasicInsaneScout            int      `json:"deaths_kit_basic_insane_scout"`
				WinsKitBasicInsaneScout              int      `json:"wins_kit_basic_insane_scout"`
				WinstreakTeam                        int      `json:"winstreak_team"`
				WinsTeamInsane                       int      `json:"wins_team_insane"`
				WinsTeam                             int      `json:"wins_team"`
				SurvivedPlayersKitBasicInsaneScout   int      `json:"survived_players_kit_basic_insane_scout"`
				AssistsNormal                        int      `json:"assists_normal"`
				AssistsKitBasicNormalScout           int      `json:"assists_kit_basic_normal_scout"`
				InsaneTank                           int      `json:"insane_tank"`
				SurvivedPlayersKitBasicNormalArcher  int      `json:"survived_players_kit_basic_normal_archer"`
				WinstreakKitBasicNormalArcher        int      `json:"winstreak_kit_basic_normal_archer"`
				WinsKitBasicNormalArcher             int      `json:"wins_kit_basic_normal_archer"`
				KillstreakKitBasicNormalArcher       int      `json:"killstreak_kit_basic_normal_archer"`
				GamesKitBasicNormalArcher            int      `json:"games_kit_basic_normal_archer"`
				KillsKitBasicNormalArcher            int      `json:"kills_kit_basic_normal_archer"`
				LossesKitBasicInsaneArcher           int      `json:"losses_kit_basic_insane_archer"`
				DeathsKitBasicInsaneArcher           int      `json:"deaths_kit_basic_insane_archer"`
				InsaneNotoriety                      int      `json:"insane_notoriety"`
				NormalTelekinesis                    int      `json:"normal_telekinesis"`
				NormalNourishment                    int      `json:"normal_nourishment"`
				NormalNoMercy                        int      `json:"normal_no_mercy"`
				NormalLowGravity                     int      `json:"normal_low_gravity"`
				InsaneNoMercy                        int      `json:"insane_no_mercy"`
				NormalTenacity                       int      `json:"normal_tenacity"`
				InsaneAdrenaline                     int      `json:"insane_adrenaline"`
				NormalVitamins                       int      `json:"normal_vitamins"`
				MasteryHuntsman                      int      `json:"mastery_huntsman"`
				NormalColdBlood                      int      `json:"normal_cold_blood"`
				InsaneLowGravity                     int      `json:"insane_low_gravity"`
				InsaneMasterBreeder                  int      `json:"insane_master_breeder"`
				InsaneTelekinesis                    int      `json:"insane_telekinesis"`
				NormalBowFlex                        int      `json:"normal_bow_flex"`
				NormalMedicine                       int      `json:"normal_medicine"`
				InsaneVitamins                       int      `json:"insane_vitamins"`
				Score                                int      `json:"score"`
				MovedOver                            bool     `json:"movedOver"`
			} `json:"SpeedUHC"`
			SkyClash struct {
				CardPacks                      int      `json:"card_packs"`
				ActiveClass                    int      `json:"active_class"`
				Killstreak                     int      `json:"killstreak"`
				HighestKillstreak              int      `json:"highestKillstreak"`
				Playstreak                     int      `json:"playstreak"`
				LongestBowShotKitSwordsman     int      `json:"longest_bow_shot_kit_swordsman"`
				LongestBowShot                 int      `json:"longest_bow_shot"`
				WinStreak                      int      `json:"win_streak"`
				BowShotsKitSwordsman           int      `json:"bow_shots_kit_swordsman"`
				DeathsKitSwordsman             int      `json:"deaths_kit_swordsman"`
				Losses                         int      `json:"losses"`
				Kills                          int      `json:"kills"`
				KillsKitSwordsman              int      `json:"kills_kit_swordsman"`
				GamesPlayed                    int      `json:"games_played"`
				BowHitsKitSwordsman            int      `json:"bow_hits_kit_swordsman"`
				BowShots                       int      `json:"bow_shots"`
				LossesSolo                     int      `json:"losses_solo"`
				MostKillsGameKitSwordsman      int      `json:"most_kills_game_kit_swordsman"`
				EnderchestsOpened              int      `json:"enderchests_opened"`
				EnderchestsOpenedKitSwordsman  int      `json:"enderchests_opened_kit_swordsman"`
				BowHits                        int      `json:"bow_hits"`
				DeathsSolo                     int      `json:"deaths_solo"`
				MostKillsGame                  int      `json:"most_kills_game"`
				Quits                          int      `json:"quits"`
				MeleeKills                     int      `json:"melee_kills"`
				KillsSolo                      int      `json:"kills_solo"`
				KillsMonthlyB                  int      `json:"kills_monthly_b"`
				Deaths                         int      `json:"deaths"`
				GamesPlayedKitSwordsman        int      `json:"games_played_kit_swordsman"`
				Coins                          int      `json:"coins"`
				MeleeKillsKitSwordsman         int      `json:"melee_kills_kit_swordsman"`
				KillsWeeklyB                   int      `json:"kills_weekly_b"`
				PlayStreak                     int      `json:"play_streak"`
				FastestWinSoloKitSwordsman     int      `json:"fastest_win_solo_kit_swordsman"`
				FastestWinSolo                 int      `json:"fastest_win_solo"`
				VoidKills                      int      `json:"void_kills"`
				Games                          int      `json:"games"`
				Wins                           int      `json:"wins"`
				VoidKillsKitSwordsman          int      `json:"void_kills_kit_swordsman"`
				SoloWinsKitSwordsman           int      `json:"solo_wins_kit_swordsman"`
				SoloWins                       int      `json:"solo_wins"`
				WinsSolo                       int      `json:"wins_solo"`
				SpawnAtWitch                   int      `json:"spawn_at_witch"`
				AssistsKitSwordsman            int      `json:"assists_kit_swordsman"`
				Assists                        int      `json:"assists"`
				Packages                       []string `json:"packages"`
				Class2                         string   `json:"class_2"`
				AssistsKitAssassin             int      `json:"assists_kit_assassin"`
				GamesPlayedKitAssassin         int      `json:"games_played_kit_assassin"`
				DeathsKitAssassin              int      `json:"deaths_kit_assassin"`
				MostKillsGameKitAssassin       int      `json:"most_kills_game_kit_assassin"`
				MeleeKillsKitAssassin          int      `json:"melee_kills_kit_assassin"`
				VoidKillsKitAssassin           int      `json:"void_kills_kit_assassin"`
				KillsKitAssassin               int      `json:"kills_kit_assassin"`
				PerkSharpenedSword             int      `json:"perk_sharpened_sword"`
				PerkBlazingArrows              int      `json:"perk_blazing_arrows"`
				PerkSharpenedSwordNew          bool     `json:"perk_sharpened_sword_new"`
				PerkCreeperNew                 bool     `json:"perk_creeper_new"`
				PerkBlazingArrowsNew           bool     `json:"perk_blazing_arrows_new"`
				PerkCreeper                    int      `json:"perk_creeper"`
				PerkSupplyDrop                 int      `json:"perk_supply_drop"`
				PerkVoidWarranty               int      `json:"perk_void_warranty"`
				PerkVoidWarrantyNew            bool     `json:"perk_void_warranty_new"`
				PerkGuardian                   int      `json:"perk_guardian"`
				PerkGuardianNew                bool     `json:"perk_guardian_new"`
				PerkSupplyDropNew              bool     `json:"perk_supply_drop_new"`
				PerkHonedBow                   int      `json:"perk_honed_bow"`
				PerkSugarRushNew               bool     `json:"perk_sugar_rush_new"`
				PerkSugarRush                  int      `json:"perk_sugar_rush"`
				PerkHonedBowNew                bool     `json:"perk_honed_bow_new"`
				PerkMarksmanDuplicates         int      `json:"perk_marksman_duplicates"`
				LongestBowShotKitAssassin      int      `json:"longest_bow_shot_kit_assassin"`
				LongestBowKillKitAssassin      int      `json:"longest_bow_kill_kit_assassin"`
				LongestBowKill                 int      `json:"longest_bow_kill"`
				BowKills                       int      `json:"bow_kills"`
				BowKillsKitAssassin            int      `json:"bow_kills_kit_assassin"`
				BowHitsKitAssassin             int      `json:"bow_hits_kit_assassin"`
				EnderchestsOpenedKitAssassin   int      `json:"enderchests_opened_kit_assassin"`
				BowShotsKitAssassin            int      `json:"bow_shots_kit_assassin"`
				FastestWinSoloKitAssassin      int      `json:"fastest_win_solo_kit_assassin"`
				SoloWinsKitAssassin            int      `json:"solo_wins_kit_assassin"`
				MobsKilledKitAssassin          int      `json:"mobs_killed_kit_assassin"`
				MobsKilled                     int      `json:"mobs_killed"`
				PerkSugarRushDuplicates        int      `json:"perk_sugar_rush_duplicates"`
				PerkNutritious                 int      `json:"perk_nutritious"`
				PerkSkeletonJockey             int      `json:"perk_skeleton_jockey"`
				PerkNutritiousNew              bool     `json:"perk_nutritious_new"`
				PerkSkeletonJockeyNew          bool     `json:"perk_skeleton_jockey_new"`
				PerkHitAndRunDuplicates        int      `json:"perk_hit_and_run_duplicates"`
				PerkSkeletonJockeyDuplicates   int      `json:"perk_skeleton_jockey_duplicates"`
				PerkHonedBowDuplicates         int      `json:"perk_honed_bow_duplicates"`
				PerkHeartyStartDuplicates      int      `json:"perk_hearty_start_duplicates"`
				PerkFruitFinderNew             bool     `json:"perk_fruit_finder_new"`
				PerkSnowGolem                  int      `json:"perk_snow_golem"`
				PerkSnowGolemNew               bool     `json:"perk_snow_golem_new"`
				PerkFruitFinder                int      `json:"perk_fruit_finder"`
				PerkBlastProtectionNew         bool     `json:"perk_blast_protection_new"`
				PerkEnergyDrink                int      `json:"perk_energy_drink"`
				PerkEnergyDrinkNew             bool     `json:"perk_energy_drink_new"`
				PerkBlastProtection            int      `json:"perk_blast_protection"`
				PerkWitch                      int      `json:"perk_witch"`
				PerkWitchNew                   bool     `json:"perk_witch_new"`
				PerkBlastProtectionDuplicates  int      `json:"perk_blast_protection_duplicates"`
				PerkInvisibility               int      `json:"perk_invisibility"`
				PerkInvisibilityNew            bool     `json:"perk_invisibility_new"`
				PerkFruitFinderDuplicates      int      `json:"perk_fruit_finder_duplicates"`
				FastestWinFourTeams            int      `json:"fastest_win_four_teams"`
				FastestWinFourTeamsKitAssassin int      `json:"fastest_win_four_teams_kit_assassin"`
				KillsFourTeams                 int      `json:"kills_four_teams"`
				FourTeamsWinsKitAssassin       int      `json:"four_teams_wins_kit_assassin"`
				FourTeamsWins                  int      `json:"four_teams_wins"`
				WinsFourTeams                  int      `json:"wins_four_teams"`
				DeathsFourTeams                int      `json:"deaths_four_teams"`
				LossesFourTeams                int      `json:"losses_four_teams"`
				PerkSupplyDropDuplicates       int      `json:"perk_supply_drop_duplicates"`
				PerkPearlAbsorptionNew         bool     `json:"perk_pearl_absorption_new"`
				PerkSharpenedSwordDuplicates   int      `json:"perk_sharpened_sword_duplicates"`
				PerkPearlAbsorption            int      `json:"perk_pearl_absorption"`
				PerkPearlAbsorptionDuplicates  int      `json:"perk_pearl_absorption_duplicates"`
				PerkEndlessQuiverNew           bool     `json:"perk_endless_quiver_new"`
				PerkEndlessQuiver              int      `json:"perk_endless_quiver"`
				PerkEnergyDrinkDuplicates      int      `json:"perk_energy_drink_duplicates"`
				PerkGuardianDuplicates         int      `json:"perk_guardian_duplicates"`
				PerkEndlessQuiverDuplicates    int      `json:"perk_endless_quiver_duplicates"`
				PerkDamagePotion               int      `json:"perk_damage_potion"`
				PerkDamagePotionNew            bool     `json:"perk_damage_potion_new"`
				PerkPacify                     int      `json:"perk_pacify"`
				PerkPacifyNew                  bool     `json:"perk_pacify_new"`
				PerkCreeperDuplicates          int      `json:"perk_creeper_duplicates"`
				PerkResistantDuplicates        int      `json:"perk_resistant_duplicates"`
				PerkPacifyDuplicates           int      `json:"perk_pacify_duplicates"`
				PerkHitAndRun                  int      `json:"perk_hit_and_run"`
				KillsDoubles                   int      `json:"kills_doubles"`
				LossesDoubles                  int      `json:"losses_doubles"`
				DeathsDoubles                  int      `json:"deaths_doubles"`
				FastestWinDoubles              int      `json:"fastest_win_doubles"`
				FastestWinDoublesKitAssassin   int      `json:"fastest_win_doubles_kit_assassin"`
				DoublesWins                    int      `json:"doubles_wins"`
				WinsDoubles                    int      `json:"wins_doubles"`
				DoublesWinsKitAssassin         int      `json:"doubles_wins_kit_assassin"`
				PerkBlazingArrowsDuplicates    int      `json:"perk_blazing_arrows_duplicates"`
				PerkArrowDeflectionNew         bool     `json:"perk_arrow_deflection_new"`
				PerkArrowDeflection            int      `json:"perk_arrow_deflection"`
				PerkDamagePotionDuplicates     int      `json:"perk_damage_potion_duplicates"`
				PerkAlchemy                    int      `json:"perk_alchemy"`
				PerkInvisibilityDuplicates     int      `json:"perk_invisibility_duplicates"`
				PerkAlchemyNew                 bool     `json:"perk_alchemy_new"`
				KillsWeeklyA                   int      `json:"kills_weekly_a"`
				PerkChickenBow                 int      `json:"perk_chicken_bow"`
				PerkChickenBowNew              bool     `json:"perk_chicken_bow_new"`
				PerkIronGolem                  int      `json:"perk_iron_golem"`
				PerkIronGolemNew               bool     `json:"perk_iron_golem_new"`
				KitAssassinMinor               int      `json:"kit_assassin_minor"`
				PerkVoidWarrantyDuplicates     int      `json:"perk_void_warranty_duplicates"`
				PerkChickenBowDuplicates       int      `json:"perk_chicken_bow_duplicates"`
				PerkArrowDeflectionDuplicates  int      `json:"perk_arrow_deflection_duplicates"`
				PerkRampageDuplicates          int      `json:"perk_rampage_duplicates"`
				PerkAlchemyDuplicates          int      `json:"perk_alchemy_duplicates"`
				PerkVoidMagnet                 int      `json:"perk_void_magnet"`
				PerkVoidMagnetNew              bool     `json:"perk_void_magnet_new"`
				PerkVoidMagnetDuplicates       int      `json:"perk_void_magnet_duplicates"`
				Class0                         string   `json:"class_0"`
				PerkSnowGolemDuplicates        int      `json:"perk_snow_golem_duplicates"`
				PerkWitchDuplicates            int      `json:"perk_witch_duplicates"`
				PerkIronGolemDuplicates        int      `json:"perk_iron_golem_duplicates"`
				PerkTripleshot                 int      `json:"perk_tripleshot"`
				PerkTripleshotNew              bool     `json:"perk_tripleshot_new"`
				PerkTripleshotDuplicates       int      `json:"perk_tripleshot_duplicates"`
				PerkRegenerationDuplicates     int      `json:"perk_regeneration_duplicates"`
				PerkExplosiveBowNew            bool     `json:"perk_explosive_bow_new"`
				PerkExplosiveBow               int      `json:"perk_explosive_bow"`
				MobKills                       int      `json:"mob_kills"`
				MobKillsKitAssassin            int      `json:"mob_kills_kit_assassin"`
				PerkEnderman                   int      `json:"perk_enderman"`
				PerkEndermanNew                bool     `json:"perk_enderman_new"`
				PerkNutritiousDuplicates       int      `json:"perk_nutritious_duplicates"`
				PerkExplosiveBowDuplicates     int      `json:"perk_explosive_bow_duplicates"`
				PerkHeadstart                  int      `json:"perk_headstart"`
				PerkHeadstartNew               bool     `json:"perk_headstart_new"`
				LossesTeamWar                  int      `json:"losses_team_war"`
				KillsTeamWar                   int      `json:"kills_team_war"`
				KillsMonthlyA                  int      `json:"kills_monthly_a"`
				PerkEndermanDuplicates         int      `json:"perk_enderman_duplicates"`
				AssassinInventory              struct {
					LEATHERCHESTPLATE0 string `json:"LEATHER_CHESTPLATE:0"`
					ENDERPEARL0        string `json:"ENDER_PEARL:0"`
					SKULLITEM3         string `json:"SKULL_ITEM:3"`
					SKULLITEM2         string `json:"SKULL_ITEM:2"`
					SKULLITEM1         string `json:"SKULL_ITEM:1"`
					COMPASS0           string `json:"COMPASS:0"`
					STONESWORD0        string `json:"STONE_SWORD:0"`
					DIAMONDSWORD0      string `json:"DIAMOND_SWORD:0"`
					LEATHERHELMET0     string `json:"LEATHER_HELMET:0"`
					POTION14           string `json:"POTION:14"`
					POTION12           string `json:"POTION:12"`
					LEATHERLEGGINGS0   string `json:"LEATHER_LEGGINGS:0"`
					LEATHERBOOTS0      string `json:"LEATHER_BOOTS:0"`
				} `json:"assassin_inventory"`
				FastestWinMegaKitAssassin     int  `json:"fastest_win_mega_kit_assassin"`
				FastestWinMega                int  `json:"fastest_win_mega"`
				WinsMega                      int  `json:"wins_mega"`
				KillsMega                     int  `json:"kills_mega"`
				MegaWinsKitAssassin           int  `json:"mega_wins_kit_assassin"`
				MegaWins                      int  `json:"mega_wins"`
				KitAssassinMaster             int  `json:"kit_assassin_master"`
				PerkLifesteal                 int  `json:"perk_lifesteal"`
				PerkLifestealNew              bool `json:"perk_lifesteal_new"`
				PerkMonsterHunterNew          bool `json:"perk_monster_hunter_new"`
				PerkMonsterHunter             int  `json:"perk_monster_hunter"`
				PerkHeadHunterNew             bool `json:"perk_head_hunter_new"`
				PerkHeadHunter                int  `json:"perk_head_hunter"`
				PerkElvenArcher               int  `json:"perk_elven_archer"`
				PerkElvenArcherNew            bool `json:"perk_elven_archer_new"`
				PerkVoidChestNew              bool `json:"perk_void_chest_new"`
				PerkVoidChest                 int  `json:"perk_void_chest"`
				PerkDoubleJumpNew             bool `json:"perk_double_jump_new"`
				PerkDoubleJump                int  `json:"perk_double_jump"`
				PerkBiggerBangs               int  `json:"perk_bigger_bangs"`
				PerkBiggerBangsNew            bool `json:"perk_bigger_bangs_new"`
				PerkResistant                 int  `json:"perk_resistant"`
				PerkHeartyStart               int  `json:"perk_hearty_start"`
				PerkRampage                   int  `json:"perk_rampage"`
				PerkElvenArcherDuplicates     int  `json:"perk_elven_archer_duplicates"`
				PerkMonsterHunterDuplicates   int  `json:"perk_monster_hunter_duplicates"`
				PerkNuclearSolutionNew        bool `json:"perk_nuclear_solution_new"`
				PerkNuclearSolution           int  `json:"perk_nuclear_solution"`
				PerkMarksman                  int  `json:"perk_marksman"`
				PerkRegeneration              int  `json:"perk_regeneration"`
				PerkBountyHunterNew           bool `json:"perk_bounty_hunter_new"`
				PerkBountyHunter              int  `json:"perk_bounty_hunter"`
				PerkLifestealDuplicates       int  `json:"perk_lifesteal_duplicates"`
				PerkMushroomAura              int  `json:"perk_mushroom_aura"`
				PerkMushroomAuraNew           bool `json:"perk_mushroom_aura_new"`
				PerkBatShieldNew              bool `json:"perk_bat_shield_new"`
				PerkBatShield                 int  `json:"perk_bat_shield"`
				PerkMushroomAuraDuplicates    int  `json:"perk_mushroom_aura_duplicates"`
				PerkBatShieldDuplicates       int  `json:"perk_bat_shield_duplicates"`
				PerkWingedBootsNew            bool `json:"perk_winged_boots_new"`
				PerkNuclearSolutionDuplicates int  `json:"perk_nuclear_solution_duplicates"`
				PerkWingedBoots               int  `json:"perk_winged_boots"`
				PerkBountyHunterDuplicates    int  `json:"perk_bounty_hunter_duplicates"`
				PerkVoidChestDuplicates       int  `json:"perk_void_chest_duplicates"`
				PerkBiggerBangsDuplicates     int  `json:"perk_bigger_bangs_duplicates"`
				PerkNoChestChallenge          int  `json:"perk_no_chest_challenge"`
				PerkNoChestChallengeNew       bool `json:"perk_no_chest_challenge_new"`
				PerkFlowerPowerNew            bool `json:"perk_flower_power_new"`
				PerkFlowerPower               int  `json:"perk_flower_power"`
				PerkArcherChallengeNew        bool `json:"perk_archer_challenge_new"`
				PerkArcherChallenge           int  `json:"perk_archer_challenge"`
				PerkDoubleJumpDuplicates      int  `json:"perk_double_jump_duplicates"`
				PerkFlowerPowerDuplicates     int  `json:"perk_flower_power_duplicates"`
				FastestWinTeamWarKitAssassin  int  `json:"fastest_win_team_war_kit_assassin"`
				FastestWinTeamWar             int  `json:"fastest_win_team_war"`
				DeathsPerkHeadstart           int  `json:"deaths_perk_headstart"`
				DeathsPerkPearlAbsorption     int  `json:"deaths_perk_pearl_absorption"`
				KillsPerkSharpenedSword       int  `json:"kills_perk_sharpened_sword"`
				DeathsPerkSharpenedSword      int  `json:"deaths_perk_sharpened_sword"`
				WinsPerkSharpenedSword        int  `json:"wins_perk_sharpened_sword"`
				WinsTeamWar                   int  `json:"wins_team_war"`
				WinsPerkPearlAbsorption       int  `json:"wins_perk_pearl_absorption"`
				KillsPerkHeadstart            int  `json:"kills_perk_headstart"`
				WinsPerkHeadstart             int  `json:"wins_perk_headstart"`
				TeamWarWinsKitAssassin        int  `json:"team_war_wins_kit_assassin"`
				TeamWarWins                   int  `json:"team_war_wins"`
				KillsPerkPearlAbsorption      int  `json:"kills_perk_pearl_absorption"`
				DeathsTeamWar                 int  `json:"deaths_team_war"`
				LossesPerkHeadstart           int  `json:"losses_perk_headstart"`
				LossesPerkPearlAbsorption     int  `json:"losses_perk_pearl_absorption"`
				LossesPerkSharpenedSword      int  `json:"losses_perk_sharpened_sword"`
				SwordsmanInventory            struct {
					SKULLITEM3     string `json:"SKULL_ITEM:3"`
					POTION9        string `json:"POTION:9"`
					SKULLITEM2     string `json:"SKULL_ITEM:2"`
					SKULLITEM1     string `json:"SKULL_ITEM:1"`
					IRONSWORD0     string `json:"IRON_SWORD:0"`
					COMPASS0       string `json:"COMPASS:0"`
					LEATHERHELMET0 string `json:"LEATHER_HELMET:0"`
				} `json:"swordsman_inventory"`
				ClassesUnlocked         int    `json:"classes_unlocked"`
				Class3                  string `json:"class_3"`
				LossesPerkRampage       int    `json:"losses_perk_rampage"`
				DeathsPerkRampage       int    `json:"deaths_perk_rampage"`
				LossesPerkHitAndRun     int    `json:"losses_perk_hit_and_run"`
				DeathsPerkHitAndRun     int    `json:"deaths_perk_hit_and_run"`
				DeathsPerkRegeneration  int    `json:"deaths_perk_regeneration"`
				LossesPerkRegeneration  int    `json:"losses_perk_regeneration"`
				MobsSpawnedKitSwordsman int    `json:"mobs_spawned_kit_swordsman"`
				MobsSpawned             int    `json:"mobs_spawned"`
				NecromancerInventory    struct {
					LEATHERCHESTPLATE0 string `json:"LEATHER_CHESTPLATE:0"`
					CHAINMAILLEGGINGS0 string `json:"CHAINMAIL_LEGGINGS:0"`
					POTION12           string `json:"POTION:12"`
					ARROW0             string `json:"ARROW:0"`
					BOW0               string `json:"BOW:0"`
					MONSTEREGG51       string `json:"MONSTER_EGG:51"`
					SKULLITEM3         string `json:"SKULL_ITEM:3"`
					SKULLITEM2         string `json:"SKULL_ITEM:2"`
					SKULLITEM1         string `json:"SKULL_ITEM:1"`
					COMPASS0           string `json:"COMPASS:0"`
				} `json:"necromancer_inventory"`
				LossesPerkWitch                 int `json:"losses_perk_witch"`
				LossesPerkSnowGolem             int `json:"losses_perk_snow_golem"`
				MobsKilledKitNecromancer        int `json:"mobs_killed_kit_necromancer"`
				DeathsKitNecromancer            int `json:"deaths_kit_necromancer"`
				DeathsPerkCreeper               int `json:"deaths_perk_creeper"`
				EnderchestsOpenedKitNecromancer int `json:"enderchests_opened_kit_necromancer"`
				DeathsPerkWitch                 int `json:"deaths_perk_witch"`
				DeathsPerkSnowGolem             int `json:"deaths_perk_snow_golem"`
				MobsSpawnedKitNecromancer       int `json:"mobs_spawned_kit_necromancer"`
				GamesPlayedKitNecromancer       int `json:"games_played_kit_necromancer"`
				LossesPerkCreeper               int `json:"losses_perk_creeper"`
				BowShotsKitNecromancer          int `json:"bow_shots_kit_necromancer"`
				MobsSpawnedKitAssassin          int `json:"mobs_spawned_kit_assassin"`
				ArcherInventory                 struct {
					ARROW0               string `json:"ARROW:0"`
					BOW0                 string `json:"BOW:0"`
					CHAINMAILBOOTS0      string `json:"CHAINMAIL_BOOTS:0"`
					CHAINMAILCHESTPLATE0 string `json:"CHAINMAIL_CHESTPLATE:0"`
					SKULLITEM3           string `json:"SKULL_ITEM:3"`
					SKULLITEM2           string `json:"SKULL_ITEM:2"`
					SKULLITEM1           string `json:"SKULL_ITEM:1"`
					COMPASS0             string `json:"COMPASS:0"`
				} `json:"archer_inventory"`
				LongestBowShotKitArcher    int `json:"longest_bow_shot_kit_archer"`
				KillsPerkHonedBow          int `json:"kills_perk_honed_bow"`
				BowHitsKitArcher           int `json:"bow_hits_kit_archer"`
				KillsKitArcher             int `json:"kills_kit_archer"`
				MobKillsKitArcher          int `json:"mob_kills_kit_archer"`
				KillsPerkEndlessQuiver     int `json:"kills_perk_endless_quiver"`
				BowShotsKitArcher          int `json:"bow_shots_kit_archer"`
				GamesPlayedKitArcher       int `json:"games_played_kit_archer"`
				DeathsPerkEndlessQuiver    int `json:"deaths_perk_endless_quiver"`
				KillsPerkExplosiveBow      int `json:"kills_perk_explosive_bow"`
				DeathsPerkExplosiveBow     int `json:"deaths_perk_explosive_bow"`
				MostKillsGameKitArcher     int `json:"most_kills_game_kit_archer"`
				MobsKilledKitArcher        int `json:"mobs_killed_kit_archer"`
				DeathsKitArcher            int `json:"deaths_kit_archer"`
				EnderchestsOpenedKitArcher int `json:"enderchests_opened_kit_archer"`
				DeathsPerkHonedBow         int `json:"deaths_perk_honed_bow"`
				FastestWinSoloKitArcher    int `json:"fastest_win_solo_kit_archer"`
				VoidKillsKitArcher         int `json:"void_kills_kit_archer"`
				WinsPerkEndlessQuiver      int `json:"wins_perk_endless_quiver"`
				SoloWinsKitArcher          int `json:"solo_wins_kit_archer"`
				WinsPerkExplosiveBow       int `json:"wins_perk_explosive_bow"`
				MeleeKillsKitArcher        int `json:"melee_kills_kit_archer"`
				WinsPerkHonedBow           int `json:"wins_perk_honed_bow"`
				CutePantsFound             int `json:"cute_pants_found"`
				MobsSpawnedKitArcher       int `json:"mobs_spawned_kit_archer"`
				CutePantsFoundKitArcher    int `json:"cute_pants_found_kit_archer"`
				LossesPerkHonedBow         int `json:"losses_perk_honed_bow"`
				LossesPerkExplosiveBow     int `json:"losses_perk_explosive_bow"`
				LossesPerkEndlessQuiver    int `json:"losses_perk_endless_quiver"`
				WinsPerkEnergyDrink        int `json:"wins_perk_energy_drink"`
				KillsPerkWingedBoots       int `json:"kills_perk_winged_boots"`
				KillsPerkEnergyDrink       int `json:"kills_perk_energy_drink"`
				WinsPerkWingedBoots        int `json:"wins_perk_winged_boots"`
			} `json:"SkyClash"`
			Legacy struct {
				Coins                   int    `json:"coins"`
				TokensDaily             int    `json:"tokens_daily"`
				TokensLastReceivedStamp int64  `json:"tokens_last_received_stamp"`
				NextTokensSeconds       int    `json:"next_tokens_seconds"`
				PaintballTokens         int    `json:"paintball_tokens"`
				TotalTokens             int    `json:"total_tokens"`
				Tokens                  int    `json:"tokens"`
				VampirezTokens          int    `json:"vampirez_tokens"`
				QuakecraftTokens        int    `json:"quakecraft_tokens"`
				GingerbreadTokens       int    `json:"gingerbread_tokens"`
				PreferredChannel        string `json:"preferredChannel"`
			} `json:"Legacy"`
			Bedwars struct {
				BedwarsBoxes                                      int      `json:"bedwars_boxes"`
				Experience                                        int      `json:"Experience"`
				FirstJoin7                                        bool     `json:"first_join_7"`
				GamesPlayedBedwars1                               int      `json:"games_played_bedwars_1"`
				Winstreak                                         int      `json:"winstreak"`
				FinalDeathsBedwars                                int      `json:"final_deaths_bedwars"`
				GoldResourcesCollectedBedwars                     int      `json:"gold_resources_collected_bedwars"`
				VoidDeathsBedwars                                 int      `json:"void_deaths_bedwars"`
				FourThreeVoidKillsBedwars                         int      `json:"four_three_void_kills_bedwars"`
				FourThreeEntityAttackKillsBedwars                 int      `json:"four_three_entity_attack_kills_bedwars"`
				VoidKillsBedwars                                  int      `json:"void_kills_bedwars"`
				DiamondResourcesCollectedBedwars                  int      `json:"diamond_resources_collected_bedwars"`
				DeathsBedwars                                     int      `json:"deaths_bedwars"`
				EntityAttackFinalDeathsBedwars                    int      `json:"entity_attack_final_deaths_bedwars"`
				EmeraldResourcesCollectedBedwars                  int      `json:"emerald_resources_collected_bedwars"`
				FourThreeIronResourcesCollectedBedwars            int      `json:"four_three_iron_resources_collected_bedwars"`
				ResourcesCollectedBedwars                         int      `json:"resources_collected_bedwars"`
				FourThreeEntityAttackFinalDeathsBedwars           int      `json:"four_three_entity_attack_final_deaths_bedwars"`
				FallDeathsBedwars                                 int      `json:"fall_deaths_bedwars"`
				FourThreeGamesPlayedBedwars                       int      `json:"four_three_games_played_bedwars"`
				GamesPlayedBedwars                                int      `json:"games_played_bedwars"`
				FourThreeFallDeathsBedwars                        int      `json:"four_three_fall_deaths_bedwars"`
				FourThreeEmeraldResourcesCollectedBedwars         int      `json:"four_three_emerald_resources_collected_bedwars"`
				FourThreeVoidDeathsBedwars                        int      `json:"four_three_void_deaths_bedwars"`
				FourThreeEntityAttackDeathsBedwars                int      `json:"four_three_entity_attack_deaths_bedwars"`
				FourThreeBedsLostBedwars                          int      `json:"four_three_beds_lost_bedwars"`
				BedsLostBedwars                                   int      `json:"beds_lost_bedwars"`
				KillsBedwars                                      int      `json:"kills_bedwars"`
				FourThreeGoldResourcesCollectedBedwars            int      `json:"four_three_gold_resources_collected_bedwars"`
				EntityAttackKillsBedwars                          int      `json:"entity_attack_kills_bedwars"`
				FourThreeDeathsBedwars                            int      `json:"four_three_deaths_bedwars"`
				FourThreeLossesBedwars                            int      `json:"four_three_losses_bedwars"`
				EntityAttackDeathsBedwars                         int      `json:"entity_attack_deaths_bedwars"`
				LossesBedwars                                     int      `json:"losses_bedwars"`
				FourThreeDiamondResourcesCollectedBedwars         int      `json:"four_three_diamond_resources_collected_bedwars"`
				IronResourcesCollectedBedwars                     int      `json:"iron_resources_collected_bedwars"`
				FourThreeKillsBedwars                             int      `json:"four_three_kills_bedwars"`
				FourThreeFinalDeathsBedwars                       int      `json:"four_three_final_deaths_bedwars"`
				FourThreeResourcesCollectedBedwars                int      `json:"four_three_resources_collected_bedwars"`
				FourThreeVoidFinalKillsBedwars                    int      `json:"four_three_void_final_kills_bedwars"`
				FourThreeWinsBedwars                              int      `json:"four_three_wins_bedwars"`
				Coins                                             int      `json:"coins"`
				EntityAttackFinalKillsBedwars                     int      `json:"entity_attack_final_kills_bedwars"`
				FourThreePermanentItemsPurchasedBedwars           int      `json:"four_three_permanent _items_purchased_bedwars"`
				FourThreeBedsBrokenBedwars                        int      `json:"four_three_beds_broken_bedwars"`
				FourThreeEntityAttackFinalKillsBedwars            int      `json:"four_three_entity_attack_final_kills_bedwars"`
				ItemsPurchasedBedwars                             int      `json:"items_purchased_bedwars"`
				WinsBedwars                                       int      `json:"wins_bedwars"`
				FinalKillsBedwars                                 int      `json:"final_kills_bedwars"`
				VoidFinalKillsBedwars                             int      `json:"void_final_kills_bedwars"`
				FourThreeFinalKillsBedwars                        int      `json:"four_three_final_kills_bedwars"`
				FourThreeItemsPurchasedBedwars                    int      `json:"four_three_items_purchased_bedwars"`
				BedsBrokenBedwars                                 int      `json:"beds_broken_bedwars"`
				Packages                                          []string `json:"packages"`
				BedwarsBoxCommons                                 int      `json:"bedwars_box_commons"`
				ChestHistory                                      string   `json:"chest_history"`
				BedwarsBox                                        int      `json:"bedwars_box"`
				BedwarsBoxRares                                   int      `json:"bedwars_box_rares"`
				BedwarsBoxEpics                                   int      `json:"bedwars_box_epics"`
				BedwarsBoxLegendaries                             int      `json:"bedwars_box_legendaries"`
				FourThreeFallFinalDeathsBedwars                   int      `json:"four_three_fall_final_deaths_bedwars"`
				FallFinalDeathsBedwars                            int      `json:"fall_final_deaths_bedwars"`
				EightOneBedsLostBedwars                           int      `json:"eight_one_beds_lost_bedwars"`
				EightOneGamesPlayedBedwars                        int      `json:"eight_one_games_played_bedwars"`
				EightOneFinalKillsBedwars                         int      `json:"eight_one_final_kills_bedwars"`
				EightOneBedsBrokenBedwars                         int      `json:"eight_one_beds_broken_bedwars"`
				EightOneIronResourcesCollectedBedwars             int      `json:"eight_one_iron_resources_collected_bedwars"`
				EightOneEmeraldResourcesCollectedBedwars          int      `json:"eight_one_emerald_resources_collected_bedwars"`
				EightOneItemsPurchasedBedwars                     int      `json:"eight_one_items_purchased_bedwars"`
				EightOneDiamondResourcesCollectedBedwars          int      `json:"eight_one_diamond_resources_collected_bedwars"`
				EightOnePermanentItemsPurchasedBedwars            int      `json:"eight_one_permanent _items_purchased_bedwars"`
				EightOneEntityAttackKillsBedwars                  int      `json:"eight_one_entity_attack_kills_bedwars"`
				EightOneVoidKillsBedwars                          int      `json:"eight_one_void_kills_bedwars"`
				EightOneVoidFinalKillsBedwars                     int      `json:"eight_one_void_final_kills_bedwars"`
				EightOneWinsBedwars                               int      `json:"eight_one_wins_bedwars"`
				EightOneKillsBedwars                              int      `json:"eight_one_kills_bedwars"`
				EightOneResourcesCollectedBedwars                 int      `json:"eight_one_resources_collected_bedwars"`
				EightOneGoldResourcesCollectedBedwars             int      `json:"eight_one_gold_resources_collected_bedwars"`
				EightOneEntityAttackFinalKillsBedwars             int      `json:"eight_one_entity_attack_final_kills_bedwars"`
				EightOneFinalDeathsBedwars                        int      `json:"eight_one_final_deaths_bedwars"`
				EightOneLossesBedwars                             int      `json:"eight_one_losses_bedwars"`
				VoidFinalDeathsBedwars                            int      `json:"void_final_deaths_bedwars"`
				EightOneVoidFinalDeathsBedwars                    int      `json:"eight_one_void_final_deaths_bedwars"`
				EightOneDeathsBedwars                             int      `json:"eight_one_deaths_bedwars"`
				EightOneEntityAttackDeathsBedwars                 int      `json:"eight_one_entity_attack_deaths_bedwars"`
				EightTwoBedsBrokenBedwars                         int      `json:"eight_two_beds_broken_bedwars"`
				EightTwoEmeraldResourcesCollectedBedwars          int      `json:"eight_two_emerald_resources_collected_bedwars"`
				EightTwoResourcesCollectedBedwars                 int      `json:"eight_two_resources_collected_bedwars"`
				EightTwoDiamondResourcesCollectedBedwars          int      `json:"eight_two_diamond_resources_collected_bedwars"`
				EightTwoWinsBedwars                               int      `json:"eight_two_wins_bedwars"`
				EightTwoKillsBedwars                              int      `json:"eight_two_kills_bedwars"`
				EightTwoIronResourcesCollectedBedwars             int      `json:"eight_two_iron_resources_collected_bedwars"`
				EightTwoVoidKillsBedwars                          int      `json:"eight_two_void_kills_bedwars"`
				EightTwoDeathsBedwars                             int      `json:"eight_two_deaths_bedwars"`
				EightTwoGamesPlayedBedwars                        int      `json:"eight_two_games_played_bedwars"`
				EightTwoItemsPurchasedBedwars                     int      `json:"eight_two_items_purchased_bedwars"`
				EightTwoBedsLostBedwars                           int      `json:"eight_two_beds_lost_bedwars"`
				EightTwoEntityAttackDeathsBedwars                 int      `json:"eight_two_entity_attack_deaths_bedwars"`
				EightTwoGoldResourcesCollectedBedwars             int      `json:"eight_two_gold_resources_collected_bedwars"`
				EightTwoVoidFinalKillsBedwars                     int      `json:"eight_two_void_final_kills_bedwars"`
				EightTwoFinalKillsBedwars                         int      `json:"eight_two_final_kills_bedwars"`
				EightTwoVoidDeathsBedwars                         int      `json:"eight_two_void_deaths_bedwars"`
				EightTwoFallDeathsBedwars                         int      `json:"eight_two_fall_deaths_bedwars"`
				EightTwoFinalDeathsBedwars                        int      `json:"eight_two_final_deaths_bedwars"`
				EightTwoEntityAttackFinalDeathsBedwars            int      `json:"eight_two_entity_attack_final_deaths_bedwars"`
				EightTwoLossesBedwars                             int      `json:"eight_two_losses_bedwars"`
				EightTwoEntityAttackFinalKillsBedwars             int      `json:"eight_two_entity_attack_final_kills_bedwars"`
				EightTwoVoidFinalDeathsBedwars                    int      `json:"eight_two_void_final_deaths_bedwars"`
				EightTwoEntityAttackKillsBedwars                  int      `json:"eight_two_entity_attack_kills_bedwars"`
				FallKillsBedwars                                  int      `json:"fall_kills_bedwars"`
				EightTwoFallKillsBedwars                          int      `json:"eight_two_fall_kills_bedwars"`
				ActiveKillEffect                                  string   `json:"activeKillEffect"`
				ProjectileKillsBedwars                            int      `json:"projectile_kills_bedwars"`
				EightTwoProjectileKillsBedwars                    int      `json:"eight_two_projectile_kills_bedwars"`
				EightTwoFallFinalDeathsBedwars                    int      `json:"eight_two_fall_final_deaths_bedwars"`
				EightTwoProjectileFinalDeathsBedwars              int      `json:"eight_two_projectile_final_deaths_bedwars"`
				ProjectileFinalDeathsBedwars                      int      `json:"projectile_final_deaths_bedwars"`
				EightTwoFallFinalKillsBedwars                     int      `json:"eight_two_fall_final_kills_bedwars"`
				FallFinalKillsBedwars                             int      `json:"fall_final_kills_bedwars"`
				ProjectileFinalKillsBedwars                       int      `json:"projectile_final_kills_bedwars"`
				EightTwoProjectileFinalKillsBedwars               int      `json:"eight_two_projectile_final_kills_bedwars"`
				EightTwoEntityExplosionKillsBedwars               int      `json:"eight_two_entity_explosion_kills_bedwars"`
				EntityExplosionKillsBedwars                       int      `json:"entity_explosion_kills_bedwars"`
				ActiveVictoryDance                                string   `json:"activeVictoryDance"`
				ProjectileDeathsBedwars                           int      `json:"projectile_deaths_bedwars"`
				EightTwoProjectileDeathsBedwars                   int      `json:"eight_two_projectile_deaths_bedwars"`
				EightOneVoidDeathsBedwars                         int      `json:"eight_one_void_deaths_bedwars"`
				FourThreeFallKillsBedwars                         int      `json:"four_three_fall_kills_bedwars"`
				FourThreeVoidFinalDeathsBedwars                   int      `json:"four_three_void_final_deaths_bedwars"`
				ActiveProjectileTrail                             string   `json:"activeProjectileTrail"`
				ActiveDeathCry                                    string   `json:"activeDeathCry"`
				GlyphStorageNew                                   string   `json:"glyph_storage_new"`
				FourThreeEntityExplosionKillsBedwars              int      `json:"four_three_entity_explosion_kills_bedwars"`
				FourFourBedsBrokenBedwars                         int      `json:"four_four_beds_broken_bedwars"`
				FourFourDeathsBedwars                             int      `json:"four_four_deaths_bedwars"`
				FourFourGoldResourcesCollectedBedwars             int      `json:"four_four_gold_resources_collected_bedwars"`
				FourFourEntityAttackDeathsBedwars                 int      `json:"four_four_entity_attack_deaths_bedwars"`
				FourFourIronResourcesCollectedBedwars             int      `json:"four_four_iron_resources_collected_bedwars"`
				FourFourGamesPlayedBedwars                        int      `json:"four_four_games_played_bedwars"`
				FourFourWinsBedwars                               int      `json:"four_four_wins_bedwars"`
				FourFourVoidFinalKillsBedwars                     int      `json:"four_four_void_final_kills_bedwars"`
				FourFourEntityAttackFinalKillsBedwars             int      `json:"four_four_entity_attack_final_kills_bedwars"`
				FourFourResourcesCollectedBedwars                 int      `json:"four_four_resources_collected_bedwars"`
				FourFourItemsPurchasedBedwars                     int      `json:"four_four_items_purchased_bedwars"`
				FourFourFinalKillsBedwars                         int      `json:"four_four_final_kills_bedwars"`
				FourFourKillsBedwars                              int      `json:"four_four_kills_bedwars"`
				FourFourDiamondResourcesCollectedBedwars          int      `json:"four_four_diamond_resources_collected_bedwars"`
				FourFourEntityAttackKillsBedwars                  int      `json:"four_four_entity_attack_kills_bedwars"`
				FourFourVoidKillsBedwars                          int      `json:"four_four_void_kills_bedwars"`
				EightTwoEntityExplosionFinalKillsBedwars          int      `json:"eight_two_entity_explosion_final_kills_bedwars"`
				EntityExplosionFinalKillsBedwars                  int      `json:"entity_explosion_final_kills_bedwars"`
				ActiveNPCSkin                                     string   `json:"activeNPCSkin"`
				SprayGlyphField                                   string   `json:"spray_glyph_field"`
				ChestHistoryNew                                   []string `json:"chest_history_new"`
				EightOneEntityAttackFinalDeathsBedwars            int      `json:"eight_one_entity_attack_final_deaths_bedwars"`
				EightTwoEntityExplosionFinalDeathsBedwars         int      `json:"eight_two_entity_explosion_final_deaths_bedwars"`
				EntityExplosionFinalDeathsBedwars                 int      `json:"entity_explosion_final_deaths_bedwars"`
				EntityExplosionDeathsBedwars                      int      `json:"entity_explosion_deaths_bedwars"`
				EightTwoEntityExplosionDeathsBedwars              int      `json:"eight_two_entity_explosion_deaths_bedwars"`
				ActiveSprays                                      string   `json:"activeSprays"`
				FourFourPermanentItemsPurchasedBedwars            int      `json:"four_four_permanent _items_purchased_bedwars"`
				FourFourBedsLostBedwars                           int      `json:"four_four_beds_lost_bedwars"`
				FourFourEmeraldResourcesCollectedBedwars          int      `json:"four_four_emerald_resources_collected_bedwars"`
				FourFourVoidDeathsBedwars                         int      `json:"four_four_void_deaths_bedwars"`
				FourFourFallKillsBedwars                          int      `json:"four_four_fall_kills_bedwars"`
				FourFourFallFinalKillsBedwars                     int      `json:"four_four_fall_final_kills_bedwars"`
				FourFourFallDeathsBedwars                         int      `json:"four_four_fall_deaths_bedwars"`
				FourFourFinalDeathsBedwars                        int      `json:"four_four_final_deaths_bedwars"`
				FourFourVoidFinalDeathsBedwars                    int      `json:"four_four_void_final_deaths_bedwars"`
				FourFourLossesBedwars                             int      `json:"four_four_losses_bedwars"`
				FourFourEntityAttackFinalDeathsBedwars            int      `json:"four_four_entity_attack_final_deaths_bedwars"`
				FourFourEntityExplosionFinalKillsBedwars          int      `json:"four_four_entity_explosion_final_kills_bedwars"`
				FourThreeEntityExplosionFinalKillsBedwars         int      `json:"four_three_entity_explosion_final_kills_bedwars"`
				FourThreeFallFinalKillsBedwars                    int      `json:"four_three_fall_final_kills_bedwars"`
				FourFourEntityExplosionDeathsBedwars              int      `json:"four_four_entity_explosion_deaths_bedwars"`
				FourFourProjectileKillsBedwars                    int      `json:"four_four_projectile_kills_bedwars"`
				FourFourEntityExplosionKillsBedwars               int      `json:"four_four_entity_explosion_kills_bedwars"`
				FourThreeEntityExplosionDeathsBedwars             int      `json:"four_three_entity_explosion_deaths_bedwars"`
				FourThreeProjectileDeathsBedwars                  int      `json:"four_three_projectile_deaths_bedwars"`
				FourFourProjectileDeathsBedwars                   int      `json:"four_four_projectile_deaths_bedwars"`
				Favourites1                                       string   `json:"favourites_1"`
				BedwarsHalloweenBoxes                             int      `json:"bedwars_halloween_boxes"`
				BedwarsOpenedChests                               int      `json:"Bedwars_openedChests"`
				BedwarsOpenedCommons                              int      `json:"Bedwars_openedCommons"`
				BedwarsOpenedRares                                int      `json:"Bedwars_openedRares"`
				SpookyOpenAch                                     int      `json:"spooky_open_ach"`
				BedwarsOpenedLegendaries                          int      `json:"Bedwars_openedLegendaries"`
				BedwarsOpenedEpics                                int      `json:"Bedwars_openedEpics"`
				ActiveKillMessages                                string   `json:"activeKillMessages"`
				FreeEventKeyBedwarsHalloweenBoxes2017             bool     `json:"free_event_key_bedwars_halloween_boxes_2017"`
				FourFourProjectileFinalKillsBedwars               int      `json:"four_four_projectile_final_kills_bedwars"`
				FourFourFallFinalDeathsBedwars                    int      `json:"four_four_fall_final_deaths_bedwars"`
				FireTickDeathsBedwars                             int      `json:"fire_tick_deaths_bedwars"`
				EightTwoFireTickDeathsBedwars                     int      `json:"eight_two_fire_tick_deaths_bedwars"`
				EightOneFallKillsBedwars                          int      `json:"eight_one_fall_kills_bedwars"`
				EightTwoWrappedPresentResourcesCollectedBedwars   int      `json:"eight_two_wrapped_present_resources_collected_bedwars"`
				WrappedPresentResourcesCollectedBedwars           int      `json:"wrapped_present_resources_collected_bedwars"`
				BedwarsChristmasBoxes                             int      `json:"bedwars_christmas_boxes"`
				FreeEventKeyBedwarsChristmasBoxes2017             bool     `json:"free_event_key_bedwars_christmas_boxes_2017"`
				ActiveGlyph                                       string   `json:"activeGlyph"`
				EightOneWrappedPresentResourcesCollectedBedwars   int      `json:"eight_one_wrapped_present_resources_collected_bedwars"`
				FourFourWrappedPresentResourcesCollectedBedwars   int      `json:"four_four_wrapped_present_resources_collected_bedwars"`
				SeenBetaMenu                                      int      `json:"seen_beta_menu"`
				Favourites2                                       string   `json:"favourites_2"`
				BedwarsLunarBoxes                                 int      `json:"bedwars_lunar_boxes"`
				FreeEventKeyBedwarsLunarBoxes2018                 bool     `json:"free_event_key_bedwars_lunar_boxes_2018"`
				BedwarsEasterBoxes                                int      `json:"bedwars_easter_boxes"`
				ActiveBedDestroy                                  string   `json:"activeBedDestroy"`
				FireTickFinalKillsBedwars                         int      `json:"fire_tick_final_kills_bedwars"`
				EightTwoFireTickFinalKillsBedwars                 int      `json:"eight_two_fire_tick_final_kills_bedwars"`
				EightTwoUltimateWinstreak                         int      `json:"eight_two_ultimate_winstreak"`
				EightTwoUltimateVoidKillsBedwars                  int      `json:"eight_two_ultimate_void_kills_bedwars"`
				EightTwoUltimateResourcesCollectedBedwars         int      `json:"eight_two_ultimate_resources_collected_bedwars"`
				EightTwoUltimateBedsLostBedwars                   int      `json:"eight_two_ultimate_beds_lost_bedwars"`
				EightTwoUltimateGamesPlayedBedwars                int      `json:"eight_two_ultimate_games_played_bedwars"`
				EightTwoUltimateIronResourcesCollectedBedwars     int      `json:"eight_two_ultimate_iron_resources_collected_bedwars"`
				EightTwoUltimateFinalDeathsBedwars                int      `json:"eight_two_ultimate_final_deaths_bedwars"`
				EightTwoUltimateLossesBedwars                     int      `json:"eight_two_ultimate_losses_bedwars"`
				EightTwoUltimateGoldResourcesCollectedBedwars     int      `json:"eight_two_ultimate_gold_resources_collected_bedwars"`
				EightTwoUltimateKillsBedwars                      int      `json:"eight_two_ultimate_kills_bedwars"`
				EightTwoUltimateFinalKillsBedwars                 int      `json:"eight_two_ultimate_final_kills_bedwars"`
				EightTwoUltimateBedsBrokenBedwars                 int      `json:"eight_two_ultimate_beds_broken_bedwars"`
				EightTwoUltimateItemsPurchasedBedwars             int      `json:"eight_two_ultimate_items_purchased_bedwars"`
				EightTwoUltimateVoidFinalKillsBedwars             int      `json:"eight_two_ultimate_void_final_kills_bedwars"`
				EightTwoUltimateEntityAttackFinalDeathsBedwars    int      `json:"eight_two_ultimate_entity_attack_final_deaths_bedwars"`
				EightTwoUltimateDeathsBedwars                     int      `json:"eight_two_ultimate_deaths_bedwars"`
				EightTwoUltimateVoidDeathsBedwars                 int      `json:"eight_two_ultimate_void_deaths_bedwars"`
				EightTwoUltimateEntityAttackKillsBedwars          int      `json:"eight_two_ultimate_entity_attack_kills_bedwars"`
				EightTwoUltimateWinsBedwars                       int      `json:"eight_two_ultimate_wins_bedwars"`
				EightTwoUltimateEntityAttackDeathsBedwars         int      `json:"eight_two_ultimate_entity_attack_deaths_bedwars"`
				EightTwoUltimateEntityAttackFinalKillsBedwars     int      `json:"eight_two_ultimate_entity_attack_final_kills_bedwars"`
				EightTwoUltimateEmeraldResourcesCollectedBedwars  int      `json:"eight_two_ultimate_emerald_resources_collected_bedwars"`
				EightTwoUltimateDiamondResourcesCollectedBedwars  int      `json:"eight_two_ultimate_diamond_resources_collected_bedwars"`
				FourFourUltimateWinstreak                         int      `json:"four_four_ultimate_winstreak"`
				FourFourUltimateDeathsBedwars                     int      `json:"four_four_ultimate_deaths_bedwars"`
				FourFourUltimateVoidKillsBedwars                  int      `json:"four_four_ultimate_void_kills_bedwars"`
				FourFourUltimateFinalKillsBedwars                 int      `json:"four_four_ultimate_final_kills_bedwars"`
				FourFourUltimatePermanentItemsPurchasedBedwars    int      `json:"four_four_ultimate_permanent _items_purchased_bedwars"`
				FourFourUltimateEntityAttackKillsBedwars          int      `json:"four_four_ultimate_entity_attack_kills_bedwars"`
				FourFourUltimateResourcesCollectedBedwars         int      `json:"four_four_ultimate_resources_collected_bedwars"`
				FourFourUltimateBedsLostBedwars                   int      `json:"four_four_ultimate_beds_lost_bedwars"`
				FourFourUltimateWinsBedwars                       int      `json:"four_four_ultimate_wins_bedwars"`
				FourFourUltimateItemsPurchasedBedwars             int      `json:"four_four_ultimate_items_purchased_bedwars"`
				FourFourUltimateDiamondResourcesCollectedBedwars  int      `json:"four_four_ultimate_diamond_resources_collected_bedwars"`
				FourFourUltimateIronResourcesCollectedBedwars     int      `json:"four_four_ultimate_iron_resources_collected_bedwars"`
				FourFourUltimateGoldResourcesCollectedBedwars     int      `json:"four_four_ultimate_gold_resources_collected_bedwars"`
				FourFourUltimateBedsBrokenBedwars                 int      `json:"four_four_ultimate_beds_broken_bedwars"`
				FourFourUltimateGamesPlayedBedwars                int      `json:"four_four_ultimate_games_played_bedwars"`
				FourFourUltimateKillsBedwars                      int      `json:"four_four_ultimate_kills_bedwars"`
				FourFourUltimateEntityAttackFinalDeathsBedwars    int      `json:"four_four_ultimate_entity_attack_final_deaths_bedwars"`
				FourFourUltimateVoidDeathsBedwars                 int      `json:"four_four_ultimate_void_deaths_bedwars"`
				FourFourUltimateVoidFinalKillsBedwars             int      `json:"four_four_ultimate_void_final_kills_bedwars"`
				FourFourUltimateFinalDeathsBedwars                int      `json:"four_four_ultimate_final_deaths_bedwars"`
				FourFourUltimateEntityAttackFinalKillsBedwars     int      `json:"four_four_ultimate_entity_attack_final_kills_bedwars"`
				FourFourUltimateEntityAttackDeathsBedwars         int      `json:"four_four_ultimate_entity_attack_deaths_bedwars"`
				FourFourUltimateLossesBedwars                     int      `json:"four_four_ultimate_losses_bedwars"`
				FourFourUltimateVoidFinalDeathsBedwars            int      `json:"four_four_ultimate_void_final_deaths_bedwars"`
				FourFourUltimateEntityExplosionFinalDeathsBedwars int      `json:"four_four_ultimate_entity_explosion_final_deaths_bedwars"`
				UnderstandsStreaks                                bool     `json:"understands_streaks"`
				UnderstandsResourceBank                           bool     `json:"understands_resource_bank"`
				CastleWinstreak                                   int      `json:"castle_winstreak"`
				CastleDeathsBedwars                               int      `json:"castle_deaths_bedwars"`
				CastleWinsBedwars                                 int      `json:"castle_wins_bedwars"`
				CastleEntityAttackFinalKillsBedwars               int      `json:"castle_entity_attack_final_kills_bedwars"`
				CastleEntityExplosionFinalKillsBedwars            int      `json:"castle_entity_explosion_final_kills_bedwars"`
				CastleItemsPurchasedBedwars                       int      `json:"castle_items_purchased_bedwars"`
				CastleDiamondResourcesCollectedBedwars            int      `json:"castle_diamond_resources_collected_bedwars"`
				CastleGoldResourcesCollectedBedwars               int      `json:"castle_gold_resources_collected_bedwars"`
				CastleResourcesCollectedBedwars                   int      `json:"castle_resources_collected_bedwars"`
				CastleEntityAttackDeathsBedwars                   int      `json:"castle_entity_attack_deaths_bedwars"`
				CastleEmeraldResourcesCollectedBedwars            int      `json:"castle_emerald_resources_collected_bedwars"`
				CastleFinalKillsBedwars                           int      `json:"castle_final_kills_bedwars"`
				CastleVoidDeathsBedwars                           int      `json:"castle_void_deaths_bedwars"`
				CastleIronResourcesCollectedBedwars               int      `json:"castle_iron_resources_collected_bedwars"`
				CastleVoidKillsBedwars                            int      `json:"castle_void_kills_bedwars"`
				CastleKillsBedwars                                int      `json:"castle_kills_bedwars"`
				CastleBedsLostBedwars                             int      `json:"castle_beds_lost_bedwars"`
				CastleEntityExplosionDeathsBedwars                int      `json:"castle_entity_explosion_deaths_bedwars"`
				CastleGamesPlayedBedwars                          int      `json:"castle_games_played_bedwars"`
				CastleEntityAttackKillsBedwars                    int      `json:"castle_entity_attack_kills_bedwars"`
				CastleVoidFinalKillsBedwars                       int      `json:"castle_void_final_kills_bedwars"`
				CastleFinalDeathsBedwars                          int      `json:"castle_final_deaths_bedwars"`
				CastleVoidFinalDeathsBedwars                      int      `json:"castle_void_final_deaths_bedwars"`
				CastleFallKillsBedwars                            int      `json:"castle_fall_kills_bedwars"`
				CastleEntityExplosionKillsBedwars                 int      `json:"castle_entity_explosion_kills_bedwars"`
				CastleFallDeathsBedwars                           int      `json:"castle_fall_deaths_bedwars"`
				CastleProjectileKillsBedwars                      int      `json:"castle_projectile_kills_bedwars"`
				CastleLossesBedwars                               int      `json:"castle_losses_bedwars"`
				CastleFallFinalKillsBedwars                       int      `json:"castle_fall_final_kills_bedwars"`
				CastleBedsBrokenBedwars                           int      `json:"castle_beds_broken_bedwars"`
				CastleFallFinalDeathsBedwars                      int      `json:"castle_fall_final_deaths_bedwars"`
				FourFourWinstreak                                 int      `json:"four_four_winstreak"`
				EightTwoWinstreak                                 int      `json:"eight_two_winstreak"`
				SelectedUltimate                                  string   `json:"selected_ultimate"`
				EightTwoUltimatePermanentItemsPurchasedBedwars    int      `json:"eight_two_ultimate_permanent _items_purchased_bedwars"`
				CastleEntityAttackFinalDeathsBedwars              int      `json:"castle_entity_attack_final_deaths_bedwars"`
				FreeEventKeyBedwarsHalloweenBoxes2018             bool     `json:"free_event_key_bedwars_halloween_boxes_2018"`
				FreeEventKeyBedwarsChristmasBoxes2018             bool     `json:"free_event_key_bedwars_christmas_boxes_2018"`
				LastHytaleAd                                      int64    `json:"lastHytaleAd"`
				EightTwoRushWinstreak                             int      `json:"eight_two_rush_winstreak"`
				EightTwoRushEntityExplosionKillsBedwars           int      `json:"eight_two_rush_entity_explosion_kills_bedwars"`
				EightTwoRushItemsPurchasedBedwars                 int      `json:"eight_two_rush_items_purchased_bedwars"`
				EightTwoRushBedsBrokenBedwars                     int      `json:"eight_two_rush_beds_broken_bedwars"`
				EightTwoRushResourcesCollectedBedwars             int      `json:"eight_two_rush_resources_collected_bedwars"`
				EightTwoRushWinsBedwars                           int      `json:"eight_two_rush_wins_bedwars"`
				EightTwoRushEntityAttackFinalKillsBedwars         int      `json:"eight_two_rush_entity_attack_final_kills_bedwars"`
				EightTwoRushEmeraldResourcesCollectedBedwars      int      `json:"eight_two_rush_emerald_resources_collected_bedwars"`
				EightTwoRushKillsBedwars                          int      `json:"eight_two_rush_kills_bedwars"`
				EightTwoRushDeathsBedwars                         int      `json:"eight_two_rush_deaths_bedwars"`
				EightTwoRushVoidKillsBedwars                      int      `json:"eight_two_rush_void_kills_bedwars"`
				EightTwoRushVoidDeathsBedwars                     int      `json:"eight_two_rush_void_deaths_bedwars"`
				EightTwoRushFinalKillsBedwars                     int      `json:"eight_two_rush_final_kills_bedwars"`
				EightTwoRushGoldResourcesCollectedBedwars         int      `json:"eight_two_rush_gold_resources_collected_bedwars"`
				EightTwoRushGamesPlayedBedwars                    int      `json:"eight_two_rush_games_played_bedwars"`
				EightTwoRushEntityAttackKillsBedwars              int      `json:"eight_two_rush_entity_attack_kills_bedwars"`
				EightTwoRushIronResourcesCollectedBedwars         int      `json:"eight_two_rush_iron_resources_collected_bedwars"`
				EightTwoRushLossesBedwars                         int      `json:"eight_two_rush_losses_bedwars"`
				EightTwoRushVoidFinalDeathsBedwars                int      `json:"eight_two_rush_void_final_deaths_bedwars"`
				EightTwoRushDiamondResourcesCollectedBedwars      int      `json:"eight_two_rush_diamond_resources_collected_bedwars"`
				EightTwoRushVoidFinalKillsBedwars                 int      `json:"eight_two_rush_void_final_kills_bedwars"`
				EightTwoRushBedsLostBedwars                       int      `json:"eight_two_rush_beds_lost_bedwars"`
				EightTwoRushPermanentItemsPurchasedBedwars        int      `json:"eight_two_rush_permanent _items_purchased_bedwars"`
				EightTwoRushFinalDeathsBedwars                    int      `json:"eight_two_rush_final_deaths_bedwars"`
				EightTwoRushEntityAttackDeathsBedwars             int      `json:"eight_two_rush_entity_attack_deaths_bedwars"`
				Privategames                                      struct {
					OneHitOneKill          bool   `json:"one_hit_one_kill"`
					HealthBuff             string `json:"health_buff"`
					Speed                  string `json:"speed"`
					RespawnTime            string `json:"respawn_time"`
					DisableBlockProtection bool   `json:"disable_block_protection"`
					MaxTeamUpgrades        bool   `json:"max_team_upgrades"`
					BedInstabreak          bool   `json:"bed_instabreak"`
					EventTime              string `json:"event_time"`
					LowGravity             bool   `json:"low_gravity"`
				} `json:"privategames"`
				FourThreeWinstreak                             int    `json:"four_three_winstreak"`
				EightTwoLuckyBedsBrokenBedwars                 int    `json:"eight_two_lucky_beds_broken_bedwars"`
				EightTwoLuckyPermanentItemsPurchasedBedwars    int    `json:"eight_two_lucky_permanent _items_purchased_bedwars"`
				EightTwoLuckyIronResourcesCollectedBedwars     int    `json:"eight_two_lucky_iron_resources_collected_bedwars"`
				EightTwoLuckyWinsBedwars                       int    `json:"eight_two_lucky_wins_bedwars"`
				EightTwoLuckyEntityAttackKillsBedwars          int    `json:"eight_two_lucky_entity_attack_kills_bedwars"`
				EightTwoLuckyKillsBedwars                      int    `json:"eight_two_lucky_kills_bedwars"`
				EightTwoLuckyGoldResourcesCollectedBedwars     int    `json:"eight_two_lucky_gold_resources_collected_bedwars"`
				EightTwoLuckyResourcesCollectedBedwars         int    `json:"eight_two_lucky_resources_collected_bedwars"`
				EightTwoLuckyFinalKillsBedwars                 int    `json:"eight_two_lucky_final_kills_bedwars"`
				EightTwoLuckyEmeraldResourcesCollectedBedwars  int    `json:"eight_two_lucky_emerald_resources_collected_bedwars"`
				EightTwoLuckyItemsPurchasedBedwars             int    `json:"eight_two_lucky_items_purchased_bedwars"`
				EightTwoLuckyVoidFinalKillsBedwars             int    `json:"eight_two_lucky_void_final_kills_bedwars"`
				EightTwoLuckyBedsLostBedwars                   int    `json:"eight_two_lucky_beds_lost_bedwars"`
				EightTwoLuckyGamesPlayedBedwars                int    `json:"eight_two_lucky_games_played_bedwars"`
				EightTwoLuckyWinstreak                         int    `json:"eight_two_lucky_winstreak"`
				EightTwoLuckyVoidKillsBedwars                  int    `json:"eight_two_lucky_void_kills_bedwars"`
				EightTwoLuckyEntityAttackFinalDeathsBedwars    int    `json:"eight_two_lucky_entity_attack_final_deaths_bedwars"`
				EightTwoLuckyLossesBedwars                     int    `json:"eight_two_lucky_losses_bedwars"`
				EightTwoLuckyFireTickKillsBedwars              int    `json:"eight_two_lucky_fire_tick_kills_bedwars"`
				EightTwoLuckyFinalDeathsBedwars                int    `json:"eight_two_lucky_final_deaths_bedwars"`
				EightTwoLuckyDeathsBedwars                     int    `json:"eight_two_lucky_deaths_bedwars"`
				EightTwoLuckyDiamondResourcesCollectedBedwars  int    `json:"eight_two_lucky_diamond_resources_collected_bedwars"`
				EightTwoLuckyVoidDeathsBedwars                 int    `json:"eight_two_lucky_void_deaths_bedwars"`
				FourFourLuckyWinstreak                         int    `json:"four_four_lucky_winstreak"`
				FourFourLuckyDiamondResourcesCollectedBedwars  int    `json:"four_four_lucky_diamond_resources_collected_bedwars"`
				FourFourLuckyEntityAttackFinalKillsBedwars     int    `json:"four_four_lucky_entity_attack_final_kills_bedwars"`
				FourFourLuckyVoidDeathsBedwars                 int    `json:"four_four_lucky_void_deaths_bedwars"`
				FourFourLuckyPermanentItemsPurchasedBedwars    int    `json:"four_four_lucky_permanent _items_purchased_bedwars"`
				FourFourLuckyEmeraldResourcesCollectedBedwars  int    `json:"four_four_lucky_emerald_resources_collected_bedwars"`
				FourFourLuckyItemsPurchasedBedwars             int    `json:"four_four_lucky_items_purchased_bedwars"`
				FourFourLuckyGamesPlayedBedwars                int    `json:"four_four_lucky_games_played_bedwars"`
				FourFourLuckyEntityAttackKillsBedwars          int    `json:"four_four_lucky_entity_attack_kills_bedwars"`
				FourFourLuckyKillsBedwars                      int    `json:"four_four_lucky_kills_bedwars"`
				FourFourLuckyBedsBrokenBedwars                 int    `json:"four_four_lucky_beds_broken_bedwars"`
				FourFourLuckyResourcesCollectedBedwars         int    `json:"four_four_lucky_resources_collected_bedwars"`
				FourFourLuckyFinalKillsBedwars                 int    `json:"four_four_lucky_final_kills_bedwars"`
				FourFourLuckyDeathsBedwars                     int    `json:"four_four_lucky_deaths_bedwars"`
				FourFourLuckyIronResourcesCollectedBedwars     int    `json:"four_four_lucky_iron_resources_collected_bedwars"`
				FourFourLuckyVoidKillsBedwars                  int    `json:"four_four_lucky_void_kills_bedwars"`
				FourFourLuckyGoldResourcesCollectedBedwars     int    `json:"four_four_lucky_gold_resources_collected_bedwars"`
				FourFourLuckyWinsBedwars                       int    `json:"four_four_lucky_wins_bedwars"`
				EightTwoLuckyEntityAttackFinalKillsBedwars     int    `json:"eight_two_lucky_entity_attack_final_kills_bedwars"`
				EightTwoLuckyProjectileKillsBedwars            int    `json:"eight_two_lucky_projectile_kills_bedwars"`
				FourFourLuckyEntityAttackDeathsBedwars         int    `json:"four_four_lucky_entity_attack_deaths_bedwars"`
				FourFourLuckyVoidFinalKillsBedwars             int    `json:"four_four_lucky_void_final_kills_bedwars"`
				FourFourLuckyFireTickFinalKillsBedwars         int    `json:"four_four_lucky_fire_tick_final_kills_bedwars"`
				FourFourLuckyProjectileFinalKillsBedwars       int    `json:"four_four_lucky_projectile_final_kills_bedwars"`
				FourFourLuckyFinalDeathsBedwars                int    `json:"four_four_lucky_final_deaths_bedwars"`
				FourFourLuckyBedsLostBedwars                   int    `json:"four_four_lucky_beds_lost_bedwars"`
				FourFourLuckyVoidFinalDeathsBedwars            int    `json:"four_four_lucky_void_final_deaths_bedwars"`
				FourFourLuckyLossesBedwars                     int    `json:"four_four_lucky_losses_bedwars"`
				FourFourLuckyFireTickDeathsBedwars             int    `json:"four_four_lucky_fire_tick_deaths_bedwars"`
				EightTwoLuckyVoidFinalDeathsBedwars            int    `json:"eight_two_lucky_void_final_deaths_bedwars"`
				EightOneWinstreak                              int    `json:"eight_one_winstreak"`
				FourFourLuckyEntityAttackFinalDeathsBedwars    int    `json:"four_four_lucky_entity_attack_final_deaths_bedwars"`
				FourFourLuckyFallDeathsBedwars                 int    `json:"four_four_lucky_fall_deaths_bedwars"`
				FourFourLuckyEntityExplosionFinalDeathsBedwars int    `json:"four_four_lucky_entity_explosion_final_deaths_bedwars"`
				FourFourLuckyFallFinalKillsBedwars             int    `json:"four_four_lucky_fall_final_kills_bedwars"`
				FourFourLuckyFireTickKillsBedwars              int    `json:"four_four_lucky_fire_tick_kills_bedwars"`
				FourFourLuckyEntityExplosionDeathsBedwars      int    `json:"four_four_lucky_entity_explosion_deaths_bedwars"`
				FourFourLuckyProjectileDeathsBedwars           int    `json:"four_four_lucky_projectile_deaths_bedwars"`
				FourFourLuckyProjectileKillsBedwars            int    `json:"four_four_lucky_projectile_kills_bedwars"`
				FourFourLuckyFireFinalKillsBedwars             int    `json:"four_four_lucky_fire_final_kills_bedwars"`
				FourFourLuckyFireTickFinalDeathsBedwars        int    `json:"four_four_lucky_fire_tick_final_deaths_bedwars"`
				FourFourLuckyEntityExplosionFinalKillsBedwars  int    `json:"four_four_lucky_entity_explosion_final_kills_bedwars"`
				FourFourLuckyFallKillsBedwars                  int    `json:"four_four_lucky_fall_kills_bedwars"`
				EightTwoLuckyFallDeathsBedwars                 int    `json:"eight_two_lucky_fall_deaths_bedwars"`
				EightTwoLuckyEntityExplosionKillsBedwars       int    `json:"eight_two_lucky_entity_explosion_kills_bedwars"`
				ActiveIslandTopper                             string `json:"activeIslandTopper"`
				FourFourArmedWinstreak                         int    `json:"four_four_armed_winstreak"`
				FourFourArmedBedsBrokenBedwars                 int    `json:"four_four_armed_beds_broken_bedwars"`
				FourFourArmedBedsLostBedwars                   int    `json:"four_four_armed_beds_lost_bedwars"`
				FourFourArmedDeathsBedwars                     int    `json:"four_four_armed_deaths_bedwars"`
				FourFourArmedDiamondResourcesCollectedBedwars  int    `json:"four_four_armed_diamond_resources_collected_bedwars"`
				FourFourArmedEntityAttackFinalDeathsBedwars    int    `json:"four_four_armed_entity_attack_final_deaths_bedwars"`
				FourFourArmedEntityAttackKillsBedwars          int    `json:"four_four_armed_entity_attack_kills_bedwars"`
				FourFourArmedFinalDeathsBedwars                int    `json:"four_four_armed_final_deaths_bedwars"`
				FourFourArmedFinalKillsBedwars                 int    `json:"four_four_armed_final_kills_bedwars"`
				FourFourArmedGamesPlayedBedwars                int    `json:"four_four_armed_games_played_bedwars"`
				FourFourArmedGoldResourcesCollectedBedwars     int    `json:"four_four_armed_gold_resources_collected_bedwars"`
				FourFourArmedIronResourcesCollectedBedwars     int    `json:"four_four_armed_iron_resources_collected_bedwars"`
				FourFourArmedItemsPurchasedBedwars             int    `json:"four_four_armed_items_purchased_bedwars"`
				FourFourArmedKillsBedwars                      int    `json:"four_four_armed_kills_bedwars"`
				FourFourArmedLossesBedwars                     int    `json:"four_four_armed_losses_bedwars"`
				FourFourArmedProjectileDeathsBedwars           int    `json:"four_four_armed_projectile_deaths_bedwars"`
				FourFourArmedProjectileKillsBedwars            int    `json:"four_four_armed_projectile_kills_bedwars"`
				FourFourArmedResourcesCollectedBedwars         int    `json:"four_four_armed_resources_collected_bedwars"`
				FourFourArmedVoidDeathsBedwars                 int    `json:"four_four_armed_void_deaths_bedwars"`
				FourFourArmedVoidFinalKillsBedwars             int    `json:"four_four_armed_void_final_kills_bedwars"`
				FourFourArmedVoidKillsBedwars                  int    `json:"four_four_armed_void_kills_bedwars"`
				FourFourArmedEmeraldResourcesCollectedBedwars  int    `json:"four_four_armed_emerald_resources_collected_bedwars"`
				FourFourArmedEntityExplosionDeathsBedwars      int    `json:"four_four_armed_entity_explosion_deaths_bedwars"`
				FourFourArmedPermanentItemsPurchasedBedwars    int    `json:"four_four_armed_permanent _items_purchased_bedwars"`
				FourFourArmedEntityAttackDeathsBedwars         int    `json:"four_four_armed_entity_attack_deaths_bedwars"`
				FourFourArmedEntityExplosionKillsBedwars       int    `json:"four_four_armed_entity_explosion_kills_bedwars"`
				FourFourArmedFallKillsBedwars                  int    `json:"four_four_armed_fall_kills_bedwars"`
				FourFourArmedProjectileFinalKillsBedwars       int    `json:"four_four_armed_projectile_final_kills_bedwars"`
				FourFourArmedWinsBedwars                       int    `json:"four_four_armed_wins_bedwars"`
				FourFourArmedProjectileFinalDeathsBedwars      int    `json:"four_four_armed_projectile_final_deaths_bedwars"`
				FourFourArmedFallFinalDeathsBedwars            int    `json:"four_four_armed_fall_final_deaths_bedwars"`
				FourFourArmedEntityExplosionFinalKillsBedwars  int    `json:"four_four_armed_entity_explosion_final_kills_bedwars"`
				TwoFourWinstreak                               int    `json:"two_four_winstreak"`
				TwoFourEntityAttackFinalKillsBedwars           int    `json:"two_four_entity_attack_final_kills_bedwars"`
				TwoFourEntityAttackKillsBedwars                int    `json:"two_four_entity_attack_kills_bedwars"`
				TwoFourFinalKillsBedwars                       int    `json:"two_four_final_kills_bedwars"`
				TwoFourGamesPlayedBedwars                      int    `json:"two_four_games_played_bedwars"`
				TwoFourGoldResourcesCollectedBedwars           int    `json:"two_four_gold_resources_collected_bedwars"`
				TwoFourIronResourcesCollectedBedwars           int    `json:"two_four_iron_resources_collected_bedwars"`
				TwoFourItemsPurchasedBedwars                   int    `json:"two_four_items_purchased_bedwars"`
				TwoFourKillsBedwars                            int    `json:"two_four_kills_bedwars"`
				TwoFourResourcesCollectedBedwars               int    `json:"two_four_resources_collected_bedwars"`
				TwoFourVoidKillsBedwars                        int    `json:"two_four_void_kills_bedwars"`
				TwoFourWinsBedwars                             int    `json:"two_four_wins_bedwars"`
				TwoFourBedsLostBedwars                         int    `json:"two_four_beds_lost_bedwars"`
				TwoFourDeathsBedwars                           int    `json:"two_four_deaths_bedwars"`
				TwoFourEntityAttackDeathsBedwars               int    `json:"two_four_entity_attack_deaths_bedwars"`
				TwoFourEntityAttackFinalDeathsBedwars          int    `json:"two_four_entity_attack_final_deaths_bedwars"`
				TwoFourFinalDeathsBedwars                      int    `json:"two_four_final_deaths_bedwars"`
				TwoFourLossesBedwars                           int    `json:"two_four_losses_bedwars"`
				TwoFourVoidDeathsBedwars                       int    `json:"two_four_void_deaths_bedwars"`
				TwoFourVoidFinalKillsBedwars                   int    `json:"two_four_void_final_kills_bedwars"`
				TwoFourBedsBrokenBedwars                       int    `json:"two_four_beds_broken_bedwars"`
				CastleMagicKillsBedwars                        int    `json:"castle_magic_kills_bedwars"`
				EightTwoPermanentItemsPurchasedBedwars         int    `json:"eight_two_permanent_items_purchased_bedwars"`
				PermanentItemsPurchasedBedwars                 int    `json:"permanent_items_purchased_bedwars"`
				CastlePermanentItemsPurchasedBedwars           int    `json:"castle_permanent_items_purchased_bedwars"`
				Practice                                       struct {
					Selected string `json:"selected"`
					Mlg      struct {
						SuccessfulAttempts int `json:"successful_attempts"`
					} `json:"mlg"`
					FireballJumping struct {
						BlocksPlaced       int `json:"blocks_placed"`
						SuccessfulAttempts int `json:"successful_attempts"`
						FailedAttempts     int `json:"failed_attempts"`
					} `json:"fireball_jumping"`
					Records struct {
						BridgingDistance30ElevationNONEAngleSTRAIGHT int `json:"bridging_distance_30:elevation_NONE:angle_STRAIGHT:"`
					} `json:"records"`
					Bridging struct {
						FailedAttempts     int `json:"failed_attempts"`
						SuccessfulAttempts int `json:"successful_attempts"`
						BlocksPlaced       int `json:"blocks_placed"`
					} `json:"bridging"`
				} `json:"practice"`
			} `json:"Bedwars"`
			MurderMystery struct {
				Coins                                                    int      `json:"coins"`
				GamesMURDERASSASSINS                                     int      `json:"games_MURDER_ASSASSINS"`
				GamesCruiseShipMURDERASSASSINS                           int      `json:"games_cruise_ship_MURDER_ASSASSINS"`
				KillsCruiseShipMURDERASSASSINS                           int      `json:"kills_cruise_ship_MURDER_ASSASSINS"`
				CoinsPickedupCruiseShip                                  int      `json:"coins_pickedup_cruise_ship"`
				KnifeKills                                               int      `json:"knife_kills"`
				GamesCruiseShip                                          int      `json:"games_cruise_ship"`
				KillsCruiseShip                                          int      `json:"kills_cruise_ship"`
				DeathsCruiseShipMURDERASSASSINS                          int      `json:"deaths_cruise_ship_MURDER_ASSASSINS"`
				DeathsCruiseShip                                         int      `json:"deaths_cruise_ship"`
				KnifeKillsCruiseShip                                     int      `json:"knife_kills_cruise_ship"`
				Deaths                                                   int      `json:"deaths"`
				DeathsMURDERASSASSINS                                    int      `json:"deaths_MURDER_ASSASSINS"`
				CoinsPickedupCruiseShipMURDERASSASSINS                   int      `json:"coins_pickedup_cruise_ship_MURDER_ASSASSINS"`
				KnifeKillsCruiseShipMURDERASSASSINS                      int      `json:"knife_kills_cruise_ship_MURDER_ASSASSINS"`
				Games                                                    int      `json:"games"`
				Kills                                                    int      `json:"kills"`
				KnifeKillsMURDERASSASSINS                                int      `json:"knife_kills_MURDER_ASSASSINS"`
				CoinsPickedupMURDERASSASSINS                             int      `json:"coins_pickedup_MURDER_ASSASSINS"`
				CoinsPickedup                                            int      `json:"coins_pickedup"`
				KillsMURDERASSASSINS                                     int      `json:"kills_MURDER_ASSASSINS"`
				MurdermysteryBooks                                       []string `json:"murdermystery_books"`
				Wins                                                     int      `json:"wins"`
				WinsMURDERHARDCORE                                       int      `json:"wins_MURDER_HARDCORE"`
				GamesHeadquartersMURDERHARDCORE                          int      `json:"games_headquarters_MURDER_HARDCORE"`
				GamesMURDERHARDCORE                                      int      `json:"games_MURDER_HARDCORE"`
				CoinsPickedupHeadquarters                                int      `json:"coins_pickedup_headquarters"`
				WinsHeadquarters                                         int      `json:"wins_headquarters"`
				CoinsPickedupHeadquartersMURDERHARDCORE                  int      `json:"coins_pickedup_headquarters_MURDER_HARDCORE"`
				GamesHeadquarters                                        int      `json:"games_headquarters"`
				CoinsPickedupMURDERHARDCORE                              int      `json:"coins_pickedup_MURDER_HARDCORE"`
				WinsHeadquartersMURDERHARDCORE                           int      `json:"wins_headquarters_MURDER_HARDCORE"`
				GrantedChests                                            int      `json:"granted_chests"`
				MmChests                                                 int      `json:"mm_chests"`
				DetectiveChance                                          int      `json:"detective_chance"`
				WinsGoldRushMURDERCLASSIC                                int      `json:"wins_gold_rush_MURDER_CLASSIC"`
				GamesGoldRushMURDERCLASSIC                               int      `json:"games_gold_rush_MURDER_CLASSIC"`
				WasHero                                                  int      `json:"was_hero"`
				KillsGoldRushMURDERCLASSIC                               int      `json:"kills_gold_rush_MURDER_CLASSIC"`
				BowKillsMURDERCLASSIC                                    int      `json:"bow_kills_MURDER_CLASSIC"`
				WasHeroMURDERCLASSIC                                     int      `json:"was_hero_MURDER_CLASSIC"`
				BowKillsGoldRushMURDERCLASSIC                            int      `json:"bow_kills_gold_rush_MURDER_CLASSIC"`
				WasHeroGoldRushMURDERCLASSIC                             int      `json:"was_hero_gold_rush_MURDER_CLASSIC"`
				GamesGoldRush                                            int      `json:"games_gold_rush"`
				KillsMURDERCLASSIC                                       int      `json:"kills_MURDER_CLASSIC"`
				CoinsPickedupGoldRushMURDERCLASSIC                       int      `json:"coins_pickedup_gold_rush_MURDER_CLASSIC"`
				WinsMURDERCLASSIC                                        int      `json:"wins_MURDER_CLASSIC"`
				CoinsPickedupMURDERCLASSIC                               int      `json:"coins_pickedup_MURDER_CLASSIC"`
				DetectiveWinsGoldRushMURDERCLASSIC                       int      `json:"detective_wins_gold_rush_MURDER_CLASSIC"`
				BowKills                                                 int      `json:"bow_kills"`
				KillsGoldRush                                            int      `json:"kills_gold_rush"`
				DetectiveWinsGoldRush                                    int      `json:"detective_wins_gold_rush"`
				DetectiveWins                                            int      `json:"detective_wins"`
				WasHeroGoldRush                                          int      `json:"was_hero_gold_rush"`
				WinsGoldRush                                             int      `json:"wins_gold_rush"`
				DetectiveWinsMURDERCLASSIC                               int      `json:"detective_wins_MURDER_CLASSIC"`
				GamesMURDERCLASSIC                                       int      `json:"games_MURDER_CLASSIC"`
				CoinsPickedupGoldRush                                    int      `json:"coins_pickedup_gold_rush"`
				BowKillsGoldRush                                         int      `json:"bow_kills_gold_rush"`
				GamesAncientTomb                                         int      `json:"games_ancient_tomb"`
				DetectiveWinsAncientTombMURDERCLASSIC                    int      `json:"detective_wins_ancient_tomb_MURDER_CLASSIC"`
				GamesAncientTombMURDERCLASSIC                            int      `json:"games_ancient_tomb_MURDER_CLASSIC"`
				WinsAncientTombMURDERCLASSIC                             int      `json:"wins_ancient_tomb_MURDER_CLASSIC"`
				CoinsPickedupAncientTombMURDERCLASSIC                    int      `json:"coins_pickedup_ancient_tomb_MURDER_CLASSIC"`
				DetectiveWinsAncientTomb                                 int      `json:"detective_wins_ancient_tomb"`
				CoinsPickedupAncientTomb                                 int      `json:"coins_pickedup_ancient_tomb"`
				WinsAncientTomb                                          int      `json:"wins_ancient_tomb"`
				WinsHollywood                                            int      `json:"wins_hollywood"`
				WinsHollywoodMURDERCLASSIC                               int      `json:"wins_hollywood_MURDER_CLASSIC"`
				DetectiveWinsHollywood                                   int      `json:"detective_wins_hollywood"`
				CoinsPickedupHollywoodMURDERCLASSIC                      int      `json:"coins_pickedup_hollywood_MURDER_CLASSIC"`
				CoinsPickedupHollywood                                   int      `json:"coins_pickedup_hollywood"`
				GamesHollywood                                           int      `json:"games_hollywood"`
				GamesHollywoodMURDERCLASSIC                              int      `json:"games_hollywood_MURDER_CLASSIC"`
				DetectiveWinsHollywoodMURDERCLASSIC                      int      `json:"detective_wins_hollywood_MURDER_CLASSIC"`
				CoinsPickedupHollywoodMURDERHARDCORE                     int      `json:"coins_pickedup_hollywood_MURDER_HARDCORE"`
				WinsHollywoodMURDERHARDCORE                              int      `json:"wins_hollywood_MURDER_HARDCORE"`
				DetectiveWinsMURDERHARDCORE                              int      `json:"detective_wins_MURDER_HARDCORE"`
				GamesHollywoodMURDERHARDCORE                             int      `json:"games_hollywood_MURDER_HARDCORE"`
				DetectiveWinsHollywoodMURDERHARDCORE                     int      `json:"detective_wins_hollywood_MURDER_HARDCORE"`
				WinsLibraryMURDERCLASSIC                                 int      `json:"wins_library_MURDER_CLASSIC"`
				WinsLibrary                                              int      `json:"wins_library"`
				GamesLibraryMURDERCLASSIC                                int      `json:"games_library_MURDER_CLASSIC"`
				CoinsPickedupLibrary                                     int      `json:"coins_pickedup_library"`
				CoinsPickedupLibraryMURDERCLASSIC                        int      `json:"coins_pickedup_library_MURDER_CLASSIC"`
				GamesLibrary                                             int      `json:"games_library"`
				WasHeroCruiseShipMURDERCLASSIC                           int      `json:"was_hero_cruise_ship_MURDER_CLASSIC"`
				BowKillsCruiseShip                                       int      `json:"bow_kills_cruise_ship"`
				DetectiveWinsCruiseShip                                  int      `json:"detective_wins_cruise_ship"`
				WinsCruiseShipMURDERCLASSIC                              int      `json:"wins_cruise_ship_MURDER_CLASSIC"`
				CoinsPickedupCruiseShipMURDERCLASSIC                     int      `json:"coins_pickedup_cruise_ship_MURDER_CLASSIC"`
				GamesCruiseShipMURDERCLASSIC                             int      `json:"games_cruise_ship_MURDER_CLASSIC"`
				WasHeroCruiseShip                                        int      `json:"was_hero_cruise_ship"`
				BowKillsCruiseShipMURDERCLASSIC                          int      `json:"bow_kills_cruise_ship_MURDER_CLASSIC"`
				DetectiveWinsCruiseShipMURDERCLASSIC                     int      `json:"detective_wins_cruise_ship_MURDER_CLASSIC"`
				WinsCruiseShip                                           int      `json:"wins_cruise_ship"`
				KillsCruiseShipMURDERCLASSIC                             int      `json:"kills_cruise_ship_MURDER_CLASSIC"`
				MurdererChance                                           int      `json:"murderer_chance"`
				BowKillsLibraryMURDERCLASSIC                             int      `json:"bow_kills_library_MURDER_CLASSIC"`
				KillsLibraryMURDERCLASSIC                                int      `json:"kills_library_MURDER_CLASSIC"`
				KnifeKillsMURDERCLASSIC                                  int      `json:"knife_kills_MURDER_CLASSIC"`
				KnifeKillsLibraryMURDERCLASSIC                           int      `json:"knife_kills_library_MURDER_CLASSIC"`
				BowKillsLibrary                                          int      `json:"bow_kills_library"`
				MurdererWinsLibraryMURDERCLASSIC                         int      `json:"murderer_wins_library_MURDER_CLASSIC"`
				KillsAsMurdererMURDERCLASSIC                             int      `json:"kills_as_murderer_MURDER_CLASSIC"`
				KnifeKillsLibrary                                        int      `json:"knife_kills_library"`
				KillsAsMurdererLibrary                                   int      `json:"kills_as_murderer_library"`
				KillsAsMurdererLibraryMURDERCLASSIC                      int      `json:"kills_as_murderer_library_MURDER_CLASSIC"`
				KillsAsMurderer                                          int      `json:"kills_as_murderer"`
				MurdererWinsMURDERCLASSIC                                int      `json:"murderer_wins_MURDER_CLASSIC"`
				KillsLibrary                                             int      `json:"kills_library"`
				MurdererWinsLibrary                                      int      `json:"murderer_wins_library"`
				MurdererWins                                             int      `json:"murderer_wins"`
				MurdererWinsCruiseShip                                   int      `json:"murderer_wins_cruise_ship"`
				MurdererWinsCruiseShipMURDERCLASSIC                      int      `json:"murderer_wins_cruise_ship_MURDER_CLASSIC"`
				ThrownKnifeKillsCruiseShip                               int      `json:"thrown_knife_kills_cruise_ship"`
				ThrownKnifeKills                                         int      `json:"thrown_knife_kills"`
				KillsAsMurdererCruiseShip                                int      `json:"kills_as_murderer_cruise_ship"`
				ThrownKnifeKillsCruiseShipMURDERCLASSIC                  int      `json:"thrown_knife_kills_cruise_ship_MURDER_CLASSIC"`
				KillsAsMurdererCruiseShipMURDERCLASSIC                   int      `json:"kills_as_murderer_cruise_ship_MURDER_CLASSIC"`
				ThrownKnifeKillsMURDERCLASSIC                            int      `json:"thrown_knife_kills_MURDER_CLASSIC"`
				KnifeKillsCruiseShipMURDERCLASSIC                        int      `json:"knife_kills_cruise_ship_MURDER_CLASSIC"`
				BowKillsHollywoodMURDERCLASSIC                           int      `json:"bow_kills_hollywood_MURDER_CLASSIC"`
				KillsHollywood                                           int      `json:"kills_hollywood"`
				KillsHollywoodMURDERCLASSIC                              int      `json:"kills_hollywood_MURDER_CLASSIC"`
				WasHeroHollywood                                         int      `json:"was_hero_hollywood"`
				WasHeroHollywoodMURDERCLASSIC                            int      `json:"was_hero_hollywood_MURDER_CLASSIC"`
				BowKillsHollywood                                        int      `json:"bow_kills_hollywood"`
				DeathsGoldRush                                           int      `json:"deaths_gold_rush"`
				DeathsGoldRushMURDERCLASSIC                              int      `json:"deaths_gold_rush_MURDER_CLASSIC"`
				DeathsMURDERCLASSIC                                      int      `json:"deaths_MURDER_CLASSIC"`
				WinsTransport                                            int      `json:"wins_transport"`
				CoinsPickedupTransport                                   int      `json:"coins_pickedup_transport"`
				WinsTransportMURDERCLASSIC                               int      `json:"wins_transport_MURDER_CLASSIC"`
				GamesTransportMURDERCLASSIC                              int      `json:"games_transport_MURDER_CLASSIC"`
				GamesTransport                                           int      `json:"games_transport"`
				CoinsPickedupTransportMURDERCLASSIC                      int      `json:"coins_pickedup_transport_MURDER_CLASSIC"`
				CoinsPickedupHypixelWorld                                int      `json:"coins_pickedup_hypixel_world"`
				WasHeroHypixelWorldMURDERCLASSIC                         int      `json:"was_hero_hypixel_world_MURDER_CLASSIC"`
				WinsHypixelWorld                                         int      `json:"wins_hypixel_world"`
				KillsHypixelWorldMURDERCLASSIC                           int      `json:"kills_hypixel_world_MURDER_CLASSIC"`
				DetectiveWinsHypixelWorld                                int      `json:"detective_wins_hypixel_world"`
				KillsHypixelWorld                                        int      `json:"kills_hypixel_world"`
				GamesHypixelWorldMURDERCLASSIC                           int      `json:"games_hypixel_world_MURDER_CLASSIC"`
				BowKillsHypixelWorldMURDERCLASSIC                        int      `json:"bow_kills_hypixel_world_MURDER_CLASSIC"`
				WinsHypixelWorldMURDERCLASSIC                            int      `json:"wins_hypixel_world_MURDER_CLASSIC"`
				GamesHypixelWorld                                        int      `json:"games_hypixel_world"`
				BowKillsHypixelWorld                                     int      `json:"bow_kills_hypixel_world"`
				CoinsPickedupHypixelWorldMURDERCLASSIC                   int      `json:"coins_pickedup_hypixel_world_MURDER_CLASSIC"`
				WasHeroHypixelWorld                                      int      `json:"was_hero_hypixel_world"`
				DetectiveWinsHypixelWorldMURDERCLASSIC                   int      `json:"detective_wins_hypixel_world_MURDER_CLASSIC"`
				DetectiveWinsLibrary                                     int      `json:"detective_wins_library"`
				WasHeroLibraryMURDERCLASSIC                              int      `json:"was_hero_library_MURDER_CLASSIC"`
				WasHeroLibrary                                           int      `json:"was_hero_library"`
				DetectiveWinsLibraryMURDERCLASSIC                        int      `json:"detective_wins_library_MURDER_CLASSIC"`
				KnifeKillsHollywood                                      int      `json:"knife_kills_hollywood"`
				MurdererWinsHollywood                                    int      `json:"murderer_wins_hollywood"`
				ThrownKnifeKillsHollywood                                int      `json:"thrown_knife_kills_hollywood"`
				KnifeKillsHollywoodMURDERCLASSIC                         int      `json:"knife_kills_hollywood_MURDER_CLASSIC"`
				MurdererWinsHollywoodMURDERCLASSIC                       int      `json:"murderer_wins_hollywood_MURDER_CLASSIC"`
				ThrownKnifeKillsHollywoodMURDERCLASSIC                   int      `json:"thrown_knife_kills_hollywood_MURDER_CLASSIC"`
				KillsAsMurdererHollywoodMURDERCLASSIC                    int      `json:"kills_as_murderer_hollywood_MURDER_CLASSIC"`
				KillsAsMurdererHollywood                                 int      `json:"kills_as_murderer_hollywood"`
				Packages                                                 []string `json:"packages"`
				MmChestHistory                                           []string `json:"mm_chest_history"`
				MurderMysteryOpenedCommons                               int      `json:"MurderMystery_openedCommons"`
				MurderMysteryOpenedChests                                int      `json:"MurderMystery_openedChests"`
				MurderMysteryOpenedRares                                 int      `json:"MurderMystery_openedRares"`
				MurderMysteryOpenedEpics                                 int      `json:"MurderMystery_openedEpics"`
				MurderMysteryOpenedLegendaries                           int      `json:"MurderMystery_openedLegendaries"`
				ThrownKnifeKillsArchives                                 int      `json:"thrown_knife_kills_archives"`
				KillsAsMurdererArchives                                  int      `json:"kills_as_murderer_archives"`
				DeathsArchivesMURDERCLASSIC                              int      `json:"deaths_archives_MURDER_CLASSIC"`
				CoinsPickedupArchives                                    int      `json:"coins_pickedup_archives"`
				KnifeKillsArchives                                       int      `json:"knife_kills_archives"`
				KillsArchives                                            int      `json:"kills_archives"`
				GamesArchivesMURDERCLASSIC                               int      `json:"games_archives_MURDER_CLASSIC"`
				KillsArchivesMURDERCLASSIC                               int      `json:"kills_archives_MURDER_CLASSIC"`
				KnifeKillsArchivesMURDERCLASSIC                          int      `json:"knife_kills_archives_MURDER_CLASSIC"`
				ThrownKnifeKillsArchivesMURDERCLASSIC                    int      `json:"thrown_knife_kills_archives_MURDER_CLASSIC"`
				DeathsArchives                                           int      `json:"deaths_archives"`
				CoinsPickedupArchivesMURDERCLASSIC                       int      `json:"coins_pickedup_archives_MURDER_CLASSIC"`
				GamesArchives                                            int      `json:"games_archives"`
				KillsAsMurdererArchivesMURDERCLASSIC                     int      `json:"kills_as_murderer_archives_MURDER_CLASSIC"`
				KillsAsMurdererAncientTombMURDERCLASSIC                  int      `json:"kills_as_murderer_ancient_tomb_MURDER_CLASSIC"`
				DeathsAncientTomb                                        int      `json:"deaths_ancient_tomb"`
				KillsAncientTombMURDERCLASSIC                            int      `json:"kills_ancient_tomb_MURDER_CLASSIC"`
				KillsAncientTomb                                         int      `json:"kills_ancient_tomb"`
				KillsAsMurdererAncientTomb                               int      `json:"kills_as_murderer_ancient_tomb"`
				KnifeKillsAncientTombMURDERCLASSIC                       int      `json:"knife_kills_ancient_tomb_MURDER_CLASSIC"`
				KnifeKillsAncientTomb                                    int      `json:"knife_kills_ancient_tomb"`
				DeathsAncientTombMURDERCLASSIC                           int      `json:"deaths_ancient_tomb_MURDER_CLASSIC"`
				ThrownKnifeKillsAncientTombMURDERCLASSIC                 int      `json:"thrown_knife_kills_ancient_tomb_MURDER_CLASSIC"`
				ThrownKnifeKillsAncientTomb                              int      `json:"thrown_knife_kills_ancient_tomb"`
				MurdererWinsGoldRushMURDERCLASSIC                        int      `json:"murderer_wins_gold_rush_MURDER_CLASSIC"`
				MurdererWinsGoldRush                                     int      `json:"murderer_wins_gold_rush"`
				KillsAsMurdererGoldRushMURDERCLASSIC                     int      `json:"kills_as_murderer_gold_rush_MURDER_CLASSIC"`
				ThrownKnifeKillsGoldRush                                 int      `json:"thrown_knife_kills_gold_rush"`
				KillsAsMurdererGoldRush                                  int      `json:"kills_as_murderer_gold_rush"`
				ThrownKnifeKillsGoldRushMURDERCLASSIC                    int      `json:"thrown_knife_kills_gold_rush_MURDER_CLASSIC"`
				KnifeKillsGoldRush                                       int      `json:"knife_kills_gold_rush"`
				KnifeKillsGoldRushMURDERCLASSIC                          int      `json:"knife_kills_gold_rush_MURDER_CLASSIC"`
				MurdererWinsArchives                                     int      `json:"murderer_wins_archives"`
				MurdererWinsArchivesMURDERCLASSIC                        int      `json:"murderer_wins_archives_MURDER_CLASSIC"`
				WinsArchives                                             int      `json:"wins_archives"`
				WinsArchivesMURDERCLASSIC                                int      `json:"wins_archives_MURDER_CLASSIC"`
				DetectiveWinsArchivesMURDERCLASSIC                       int      `json:"detective_wins_archives_MURDER_CLASSIC"`
				DetectiveWinsArchives                                    int      `json:"detective_wins_archives"`
				ActiveVictoryDance                                       string   `json:"active_victory_dance"`
				ActiveProjectileTrail                                    string   `json:"active_projectile_trail"`
				ActiveLastWords                                          string   `json:"active_last_words"`
				ActiveKnifeSkin                                          string   `json:"active_knife_skin"`
				BowKillsArchives                                         int      `json:"bow_kills_archives"`
				WasHeroArchives                                          int      `json:"was_hero_archives"`
				WasHeroArchivesMURDERCLASSIC                             int      `json:"was_hero_archives_MURDER_CLASSIC"`
				BowKillsArchivesMURDERCLASSIC                            int      `json:"bow_kills_archives_MURDER_CLASSIC"`
				GamesTowerfallMURDERCLASSIC                              int      `json:"games_towerfall_MURDER_CLASSIC"`
				CoinsPickedupTowerfall                                   int      `json:"coins_pickedup_towerfall"`
				GamesTowerfall                                           int      `json:"games_towerfall"`
				WinsTowerfallMURDERCLASSIC                               int      `json:"wins_towerfall_MURDER_CLASSIC"`
				WinsTowerfall                                            int      `json:"wins_towerfall"`
				CoinsPickedupTowerfallMURDERCLASSIC                      int      `json:"coins_pickedup_towerfall_MURDER_CLASSIC"`
				MurdererWinsAncientTombMURDERCLASSIC                     int      `json:"murderer_wins_ancient_tomb_MURDER_CLASSIC"`
				MurdererWinsAncientTomb                                  int      `json:"murderer_wins_ancient_tomb"`
				WasHeroTransport                                         int      `json:"was_hero_transport"`
				KillsTransportMURDERCLASSIC                              int      `json:"kills_transport_MURDER_CLASSIC"`
				KillsTransport                                           int      `json:"kills_transport"`
				DetectiveWinsTransport                                   int      `json:"detective_wins_transport"`
				DetectiveWinsTransportMURDERCLASSIC                      int      `json:"detective_wins_transport_MURDER_CLASSIC"`
				BowKillsTransportMURDERCLASSIC                           int      `json:"bow_kills_transport_MURDER_CLASSIC"`
				BowKillsTransport                                        int      `json:"bow_kills_transport"`
				WasHeroTransportMURDERCLASSIC                            int      `json:"was_hero_transport_MURDER_CLASSIC"`
				GamesAncientTombMURDERASSASSINS                          int      `json:"games_ancient_tomb_MURDER_ASSASSINS"`
				ThrownKnifeKillsMURDERASSASSINS                          int      `json:"thrown_knife_kills_MURDER_ASSASSINS"`
				KillsAncientTombMURDERASSASSINS                          int      `json:"kills_ancient_tomb_MURDER_ASSASSINS"`
				CoinsPickedupAncientTombMURDERASSASSINS                  int      `json:"coins_pickedup_ancient_tomb_MURDER_ASSASSINS"`
				WinsAncientTombMURDERASSASSINS                           int      `json:"wins_ancient_tomb_MURDER_ASSASSINS"`
				KnifeKillsAncientTombMURDERASSASSINS                     int      `json:"knife_kills_ancient_tomb_MURDER_ASSASSINS"`
				WinsMURDERASSASSINS                                      int      `json:"wins_MURDER_ASSASSINS"`
				ThrownKnifeKillsAncientTombMURDERASSASSINS               int      `json:"thrown_knife_kills_ancient_tomb_MURDER_ASSASSINS"`
				BowKillsAncientTomb                                      int      `json:"bow_kills_ancient_tomb"`
				WasHeroAncientTomb                                       int      `json:"was_hero_ancient_tomb"`
				BowKillsAncientTombMURDERCLASSIC                         int      `json:"bow_kills_ancient_tomb_MURDER_CLASSIC"`
				WasHeroAncientTombMURDERCLASSIC                          int      `json:"was_hero_ancient_tomb_MURDER_CLASSIC"`
				Showqueuebook                                            bool     `json:"showqueuebook"`
				DeathsHollywoodMURDERCLASSIC                             int      `json:"deaths_hollywood_MURDER_CLASSIC"`
				DeathsHollywood                                          int      `json:"deaths_hollywood"`
				KillsAsMurdererTowerfallMURDERCLASSIC                    int      `json:"kills_as_murderer_towerfall_MURDER_CLASSIC"`
				KnifeKillsTowerfall                                      int      `json:"knife_kills_towerfall"`
				KillsTowerfall                                           int      `json:"kills_towerfall"`
				MurdererWinsTowerfallMURDERCLASSIC                       int      `json:"murderer_wins_towerfall_MURDER_CLASSIC"`
				KnifeKillsTowerfallMURDERCLASSIC                         int      `json:"knife_kills_towerfall_MURDER_CLASSIC"`
				ThrownKnifeKillsTowerfall                                int      `json:"thrown_knife_kills_towerfall"`
				ThrownKnifeKillsTowerfallMURDERCLASSIC                   int      `json:"thrown_knife_kills_towerfall_MURDER_CLASSIC"`
				KillsTowerfallMURDERCLASSIC                              int      `json:"kills_towerfall_MURDER_CLASSIC"`
				KillsAsMurdererTowerfall                                 int      `json:"kills_as_murderer_towerfall"`
				MurdererWinsTowerfall                                    int      `json:"murderer_wins_towerfall"`
				KnifeKillsHypixelWorld                                   int      `json:"knife_kills_hypixel_world"`
				ThrownKnifeKillsHypixelWorld                             int      `json:"thrown_knife_kills_hypixel_world"`
				KillsAsMurdererHypixelWorld                              int      `json:"kills_as_murderer_hypixel_world"`
				MurdererWinsHypixelWorldMURDERCLASSIC                    int      `json:"murderer_wins_hypixel_world_MURDER_CLASSIC"`
				ThrownKnifeKillsHypixelWorldMURDERCLASSIC                int      `json:"thrown_knife_kills_hypixel_world_MURDER_CLASSIC"`
				KillsAsMurdererHypixelWorldMURDERCLASSIC                 int      `json:"kills_as_murderer_hypixel_world_MURDER_CLASSIC"`
				MurdererWinsHypixelWorld                                 int      `json:"murderer_wins_hypixel_world"`
				KnifeKillsHypixelWorldMURDERCLASSIC                      int      `json:"knife_kills_hypixel_world_MURDER_CLASSIC"`
				DeathsHypixelWorldMURDERCLASSIC                          int      `json:"deaths_hypixel_world_MURDER_CLASSIC"`
				DeathsHypixelWorld                                       int      `json:"deaths_hypixel_world"`
				GamesHeadquartersMURDERCLASSIC                           int      `json:"games_headquarters_MURDER_CLASSIC"`
				WinsHeadquartersMURDERCLASSIC                            int      `json:"wins_headquarters_MURDER_CLASSIC"`
				DetectiveWinsHeadquartersMURDERCLASSIC                   int      `json:"detective_wins_headquarters_MURDER_CLASSIC"`
				DetectiveWinsHeadquarters                                int      `json:"detective_wins_headquarters"`
				CoinsPickedupHeadquartersMURDERCLASSIC                   int      `json:"coins_pickedup_headquarters_MURDER_CLASSIC"`
				KnifeKillsGoldRushMURDERASSASSINS                        int      `json:"knife_kills_gold_rush_MURDER_ASSASSINS"`
				GamesGoldRushMURDERASSASSINS                             int      `json:"games_gold_rush_MURDER_ASSASSINS"`
				CoinsPickedupGoldRushMURDERASSASSINS                     int      `json:"coins_pickedup_gold_rush_MURDER_ASSASSINS"`
				KillsGoldRushMURDERASSASSINS                             int      `json:"kills_gold_rush_MURDER_ASSASSINS"`
				WinsGoldRushMURDERASSASSINS                              int      `json:"wins_gold_rush_MURDER_ASSASSINS"`
				WinsHeadquartersMURDERASSASSINS                          int      `json:"wins_headquarters_MURDER_ASSASSINS"`
				KillsHeadquarters                                        int      `json:"kills_headquarters"`
				KillsHeadquartersMURDERASSASSINS                         int      `json:"kills_headquarters_MURDER_ASSASSINS"`
				CoinsPickedupHeadquartersMURDERASSASSINS                 int      `json:"coins_pickedup_headquarters_MURDER_ASSASSINS"`
				KnifeKillsHeadquarters                                   int      `json:"knife_kills_headquarters"`
				GamesHeadquartersMURDERASSASSINS                         int      `json:"games_headquarters_MURDER_ASSASSINS"`
				ThrownKnifeKillsHeadquarters                             int      `json:"thrown_knife_kills_headquarters"`
				ThrownKnifeKillsHeadquartersMURDERASSASSINS              int      `json:"thrown_knife_kills_headquarters_MURDER_ASSASSINS"`
				KnifeKillsHeadquartersMURDERASSASSINS                    int      `json:"knife_kills_headquarters_MURDER_ASSASSINS"`
				ThrownKnifeKillsCruiseShipMURDERASSASSINS                int      `json:"thrown_knife_kills_cruise_ship_MURDER_ASSASSINS"`
				WinsCruiseShipMURDERASSASSINS                            int      `json:"wins_cruise_ship_MURDER_ASSASSINS"`
				ThrownKnifeKillsGoldRushMURDERASSASSINS                  int      `json:"thrown_knife_kills_gold_rush_MURDER_ASSASSINS"`
				GamesTowerfallMURDERASSASSINS                            int      `json:"games_towerfall_MURDER_ASSASSINS"`
				CoinsPickedupTowerfallMURDERASSASSINS                    int      `json:"coins_pickedup_towerfall_MURDER_ASSASSINS"`
				KnifeKillsTowerfallMURDERASSASSINS                       int      `json:"knife_kills_towerfall_MURDER_ASSASSINS"`
				WinsTowerfallMURDERASSASSINS                             int      `json:"wins_towerfall_MURDER_ASSASSINS"`
				KillsTowerfallMURDERASSASSINS                            int      `json:"kills_towerfall_MURDER_ASSASSINS"`
				CoinsPickedupArchivesMURDERASSASSINS                     int      `json:"coins_pickedup_archives_MURDER_ASSASSINS"`
				DeathsArchivesMURDERASSASSINS                            int      `json:"deaths_archives_MURDER_ASSASSINS"`
				ThrownKnifeKillsArchivesMURDERASSASSINS                  int      `json:"thrown_knife_kills_archives_MURDER_ASSASSINS"`
				KnifeKillsArchivesMURDERASSASSINS                        int      `json:"knife_kills_archives_MURDER_ASSASSINS"`
				KillsArchivesMURDERASSASSINS                             int      `json:"kills_archives_MURDER_ASSASSINS"`
				GamesArchivesMURDERASSASSINS                             int      `json:"games_archives_MURDER_ASSASSINS"`
				DeathsAncientTombMURDERASSASSINS                         int      `json:"deaths_ancient_tomb_MURDER_ASSASSINS"`
				KillsTransportMURDERASSASSINS                            int      `json:"kills_transport_MURDER_ASSASSINS"`
				CoinsPickedupTransportMURDERASSASSINS                    int      `json:"coins_pickedup_transport_MURDER_ASSASSINS"`
				KnifeKillsTransport                                      int      `json:"knife_kills_transport"`
				GamesTransportMURDERASSASSINS                            int      `json:"games_transport_MURDER_ASSASSINS"`
				ThrownKnifeKillsTransport                                int      `json:"thrown_knife_kills_transport"`
				WinsTransportMURDERASSASSINS                             int      `json:"wins_transport_MURDER_ASSASSINS"`
				ThrownKnifeKillsTransportMURDERASSASSINS                 int      `json:"thrown_knife_kills_transport_MURDER_ASSASSINS"`
				KnifeKillsTransportMURDERASSASSINS                       int      `json:"knife_kills_transport_MURDER_ASSASSINS"`
				WinsHypixelWorldMURDERASSASSINS                          int      `json:"wins_hypixel_world_MURDER_ASSASSINS"`
				CoinsPickedupHypixelWorldMURDERASSASSINS                 int      `json:"coins_pickedup_hypixel_world_MURDER_ASSASSINS"`
				GamesHypixelWorldMURDERASSASSINS                         int      `json:"games_hypixel_world_MURDER_ASSASSINS"`
				ThrownKnifeKillsHypixelWorldMURDERASSASSINS              int      `json:"thrown_knife_kills_hypixel_world_MURDER_ASSASSINS"`
				KillsHypixelWorldMURDERASSASSINS                         int      `json:"kills_hypixel_world_MURDER_ASSASSINS"`
				KnifeKillsHypixelWorldMURDERASSASSINS                    int      `json:"knife_kills_hypixel_world_MURDER_ASSASSINS"`
				DeathsHypixelWorldMURDERASSASSINS                        int      `json:"deaths_hypixel_world_MURDER_ASSASSINS"`
				WinsArchivesMURDERASSASSINS                              int      `json:"wins_archives_MURDER_ASSASSINS"`
				BowKillsHypixelWorldMURDERASSASSINS                      int      `json:"bow_kills_hypixel_world_MURDER_ASSASSINS"`
				BowKillsMURDERASSASSINS                                  int      `json:"bow_kills_MURDER_ASSASSINS"`
				WinsHollywoodMURDERASSASSINS                             int      `json:"wins_hollywood_MURDER_ASSASSINS"`
				GamesHollywoodMURDERASSASSINS                            int      `json:"games_hollywood_MURDER_ASSASSINS"`
				CoinsPickedupHollywoodMURDERASSASSINS                    int      `json:"coins_pickedup_hollywood_MURDER_ASSASSINS"`
				ThrownKnifeKillsHollywoodMURDERASSASSINS                 int      `json:"thrown_knife_kills_hollywood_MURDER_ASSASSINS"`
				KillsHollywoodMURDERASSASSINS                            int      `json:"kills_hollywood_MURDER_ASSASSINS"`
				KnifeKillsHollywoodMURDERASSASSINS                       int      `json:"knife_kills_hollywood_MURDER_ASSASSINS"`
				DeathsGoldRushMURDERASSASSINS                            int      `json:"deaths_gold_rush_MURDER_ASSASSINS"`
				KnifeKillsLibraryMURDERASSASSINS                         int      `json:"knife_kills_library_MURDER_ASSASSINS"`
				KillsLibraryMURDERASSASSINS                              int      `json:"kills_library_MURDER_ASSASSINS"`
				ThrownKnifeKillsLibraryMURDERASSASSINS                   int      `json:"thrown_knife_kills_library_MURDER_ASSASSINS"`
				ThrownKnifeKillsLibrary                                  int      `json:"thrown_knife_kills_library"`
				CoinsPickedupLibraryMURDERASSASSINS                      int      `json:"coins_pickedup_library_MURDER_ASSASSINS"`
				GamesLibraryMURDERASSASSINS                              int      `json:"games_library_MURDER_ASSASSINS"`
				WinsLibraryMURDERASSASSINS                               int      `json:"wins_library_MURDER_ASSASSINS"`
				DeathsHollywoodMURDERASSASSINS                           int      `json:"deaths_hollywood_MURDER_ASSASSINS"`
				DeathsHeadquarters                                       int      `json:"deaths_headquarters"`
				DeathsHeadquartersMURDERASSASSINS                        int      `json:"deaths_headquarters_MURDER_ASSASSINS"`
				BowKillsCruiseShipMURDERASSASSINS                        int      `json:"bow_kills_cruise_ship_MURDER_ASSASSINS"`
				ThrownKnifeKillsTowerfallMURDERASSASSINS                 int      `json:"thrown_knife_kills_towerfall_MURDER_ASSASSINS"`
				DeathsTowerfallMURDERASSASSINS                           int      `json:"deaths_towerfall_MURDER_ASSASSINS"`
				DeathsTowerfall                                          int      `json:"deaths_towerfall"`
				DeathsLibraryMURDERASSASSINS                             int      `json:"deaths_library_MURDER_ASSASSINS"`
				DeathsLibrary                                            int      `json:"deaths_library"`
				BowKillsTowerfall                                        int      `json:"bow_kills_towerfall"`
				BowKillsTowerfallMURDERASSASSINS                         int      `json:"bow_kills_towerfall_MURDER_ASSASSINS"`
				DeathsCruiseShipMURDERCLASSIC                            int      `json:"deaths_cruise_ship_MURDER_CLASSIC"`
				TotalTimeSurvivedSecondsGoldRush                         int      `json:"total_time_survived_seconds_gold_rush"`
				TotalTimeSurvivedSeconds                                 int      `json:"total_time_survived_seconds"`
				LongestTimeAsSurvivorSeconds                             int      `json:"longest_time_as_survivor_seconds"`
				LongestTimeAsSurvivorSecondsGoldRush                     int      `json:"longest_time_as_survivor_seconds_gold_rush"`
				TotalTimeSurvivedSecondsGoldRushMURDERINFECTION          int      `json:"total_time_survived_seconds_gold_rush_MURDER_INFECTION"`
				LongestTimeAsSurvivorSecondsMURDERINFECTION              int      `json:"longest_time_as_survivor_seconds_MURDER_INFECTION"`
				TotalTimeSurvivedSecondsMURDERINFECTION                  int      `json:"total_time_survived_seconds_MURDER_INFECTION"`
				LongestTimeAsSurvivorSecondsGoldRushMURDERINFECTION      int      `json:"longest_time_as_survivor_seconds_gold_rush_MURDER_INFECTION"`
				LastOneAliveGoldRush                                     int      `json:"last_one_alive_gold_rush"`
				GamesMURDERINFECTION                                     int      `json:"games_MURDER_INFECTION"`
				KillsAsSurvivorMURDERINFECTION                           int      `json:"kills_as_survivor_MURDER_INFECTION"`
				KillsAsSurvivor                                          int      `json:"kills_as_survivor"`
				GamesGoldRushMURDERINFECTION                             int      `json:"games_gold_rush_MURDER_INFECTION"`
				KillsAsSurvivorGoldRushMURDERINFECTION                   int      `json:"kills_as_survivor_gold_rush_MURDER_INFECTION"`
				LastOneAliveGoldRushMURDERINFECTION                      int      `json:"last_one_alive_gold_rush_MURDER_INFECTION"`
				LastOneAliveMURDERINFECTION                              int      `json:"last_one_alive_MURDER_INFECTION"`
				LastOneAlive                                             int      `json:"last_one_alive"`
				KillsAsSurvivorGoldRush                                  int      `json:"kills_as_survivor_gold_rush"`
				MmChristmasChests                                        int      `json:"mm_christmas_chests"`
				QuickestShowdownWinTimeSecondsMURDERSHOWDOWN             int      `json:"quickest_showdown_win_time_seconds_MURDER_SHOWDOWN"`
				QuickestShowdownWinTimeSecondsMountain                   int      `json:"quickest_showdown_win_time_seconds_mountain"`
				QuickestShowdownWinTimeSeconds                           int      `json:"quickest_showdown_win_time_seconds"`
				QuickestShowdownWinTimeSecondsMountainMURDERSHOWDOWN     int      `json:"quickest_showdown_win_time_seconds_mountain_MURDER_SHOWDOWN"`
				GamesMountainMURDERSHOWDOWN                              int      `json:"games_mountain_MURDER_SHOWDOWN"`
				WinsMountain                                             int      `json:"wins_mountain"`
				WinsMountainMURDERSHOWDOWN                               int      `json:"wins_mountain_MURDER_SHOWDOWN"`
				ShowdownPotg                                             int      `json:"showdown_potg"`
				DeathsMountain                                           int      `json:"deaths_mountain"`
				ShowdownPotgMURDERSHOWDOWN                               int      `json:"showdown_potg_MURDER_SHOWDOWN"`
				CoinsPickedupMountain                                    int      `json:"coins_pickedup_mountain"`
				CoinsPickedupMountainMURDERSHOWDOWN                      int      `json:"coins_pickedup_mountain_MURDER_SHOWDOWN"`
				BowKillsMountainMURDERSHOWDOWN                           int      `json:"bow_kills_mountain_MURDER_SHOWDOWN"`
				KillsMountainMURDERSHOWDOWN                              int      `json:"kills_mountain_MURDER_SHOWDOWN"`
				CoinsPickedupMURDERSHOWDOWN                              int      `json:"coins_pickedup_MURDER_SHOWDOWN"`
				BowKillsMURDERSHOWDOWN                                   int      `json:"bow_kills_MURDER_SHOWDOWN"`
				ShowdownPotgMountainMURDERSHOWDOWN                       int      `json:"showdown_potg_mountain_MURDER_SHOWDOWN"`
				KillsMountain                                            int      `json:"kills_mountain"`
				KillsMURDERSHOWDOWN                                      int      `json:"kills_MURDER_SHOWDOWN"`
				DeathsMountainMURDERSHOWDOWN                             int      `json:"deaths_mountain_MURDER_SHOWDOWN"`
				GamesMountain                                            int      `json:"games_mountain"`
				GamesMURDERSHOWDOWN                                      int      `json:"games_MURDER_SHOWDOWN"`
				WinsMURDERSHOWDOWN                                       int      `json:"wins_MURDER_SHOWDOWN"`
				BowKillsMountain                                         int      `json:"bow_kills_mountain"`
				DeathsMURDERSHOWDOWN                                     int      `json:"deaths_MURDER_SHOWDOWN"`
				ShowdownPotgMountain                                     int      `json:"showdown_potg_mountain"`
				QuickestShowdownWinTimeSecondsArchives                   int      `json:"quickest_showdown_win_time_seconds_archives"`
				QuickestShowdownWinTimeSecondsArchivesMURDERSHOWDOWN     int      `json:"quickest_showdown_win_time_seconds_archives_MURDER_SHOWDOWN"`
				CoinsPickedupArchivesMURDERSHOWDOWN                      int      `json:"coins_pickedup_archives_MURDER_SHOWDOWN"`
				BowKillsArchivesMURDERSHOWDOWN                           int      `json:"bow_kills_archives_MURDER_SHOWDOWN"`
				DeathsArchivesMURDERSHOWDOWN                             int      `json:"deaths_archives_MURDER_SHOWDOWN"`
				KnifeKillsMURDERSHOWDOWN                                 int      `json:"knife_kills_MURDER_SHOWDOWN"`
				KillsArchivesMURDERSHOWDOWN                              int      `json:"kills_archives_MURDER_SHOWDOWN"`
				ShowdownPotgArchivesMURDERSHOWDOWN                       int      `json:"showdown_potg_archives_MURDER_SHOWDOWN"`
				GamesArchivesMURDERSHOWDOWN                              int      `json:"games_archives_MURDER_SHOWDOWN"`
				WinsArchivesMURDERSHOWDOWN                               int      `json:"wins_archives_MURDER_SHOWDOWN"`
				KnifeKillsArchivesMURDERSHOWDOWN                         int      `json:"knife_kills_archives_MURDER_SHOWDOWN"`
				ShowdownPotgArchives                                     int      `json:"showdown_potg_archives"`
				ShowdownPotgCruiseShip                                   int      `json:"showdown_potg_cruise_ship"`
				WinsCruiseShipMURDERSHOWDOWN                             int      `json:"wins_cruise_ship_MURDER_SHOWDOWN"`
				KillsCruiseShipMURDERSHOWDOWN                            int      `json:"kills_cruise_ship_MURDER_SHOWDOWN"`
				CoinsPickedupCruiseShipMURDERSHOWDOWN                    int      `json:"coins_pickedup_cruise_ship_MURDER_SHOWDOWN"`
				KnifeKillsCruiseShipMURDERSHOWDOWN                       int      `json:"knife_kills_cruise_ship_MURDER_SHOWDOWN"`
				ShowdownPotgCruiseShipMURDERSHOWDOWN                     int      `json:"showdown_potg_cruise_ship_MURDER_SHOWDOWN"`
				BowKillsCruiseShipMURDERSHOWDOWN                         int      `json:"bow_kills_cruise_ship_MURDER_SHOWDOWN"`
				GamesCruiseShipMURDERSHOWDOWN                            int      `json:"games_cruise_ship_MURDER_SHOWDOWN"`
				DeathsCruiseShipMURDERSHOWDOWN                           int      `json:"deaths_cruise_ship_MURDER_SHOWDOWN"`
				KnifeKillsMountainMURDERSHOWDOWN                         int      `json:"knife_kills_mountain_MURDER_SHOWDOWN"`
				KnifeKillsMountain                                       int      `json:"knife_kills_mountain"`
				BowKillsTowerfallMURDERSHOWDOWN                          int      `json:"bow_kills_towerfall_MURDER_SHOWDOWN"`
				KnifeKillsTowerfallMURDERSHOWDOWN                        int      `json:"knife_kills_towerfall_MURDER_SHOWDOWN"`
				CoinsPickedupTowerfallMURDERSHOWDOWN                     int      `json:"coins_pickedup_towerfall_MURDER_SHOWDOWN"`
				ShowdownPotgTowerfall                                    int      `json:"showdown_potg_towerfall"`
				GamesTowerfallMURDERSHOWDOWN                             int      `json:"games_towerfall_MURDER_SHOWDOWN"`
				DeathsTowerfallMURDERSHOWDOWN                            int      `json:"deaths_towerfall_MURDER_SHOWDOWN"`
				KillsTowerfallMURDERSHOWDOWN                             int      `json:"kills_towerfall_MURDER_SHOWDOWN"`
				WinsTowerfallMURDERSHOWDOWN                              int      `json:"wins_towerfall_MURDER_SHOWDOWN"`
				ShowdownPotgTowerfallMURDERSHOWDOWN                      int      `json:"showdown_potg_towerfall_MURDER_SHOWDOWN"`
				GamesHypixelWorldMURDERSHOWDOWN                          int      `json:"games_hypixel_world_MURDER_SHOWDOWN"`
				KillsHypixelWorldMURDERSHOWDOWN                          int      `json:"kills_hypixel_world_MURDER_SHOWDOWN"`
				CoinsPickedupHypixelWorldMURDERSHOWDOWN                  int      `json:"coins_pickedup_hypixel_world_MURDER_SHOWDOWN"`
				ShowdownPotgHypixelWorldMURDERSHOWDOWN                   int      `json:"showdown_potg_hypixel_world_MURDER_SHOWDOWN"`
				BowKillsHypixelWorldMURDERSHOWDOWN                       int      `json:"bow_kills_hypixel_world_MURDER_SHOWDOWN"`
				KnifeKillsHypixelWorldMURDERSHOWDOWN                     int      `json:"knife_kills_hypixel_world_MURDER_SHOWDOWN"`
				DeathsHypixelWorldMURDERSHOWDOWN                         int      `json:"deaths_hypixel_world_MURDER_SHOWDOWN"`
				ShowdownPotgHypixelWorld                                 int      `json:"showdown_potg_hypixel_world"`
				GamesLibraryMURDERSHOWDOWN                               int      `json:"games_library_MURDER_SHOWDOWN"`
				DeathsLibraryMURDERSHOWDOWN                              int      `json:"deaths_library_MURDER_SHOWDOWN"`
				KnifeKillsLibraryMURDERSHOWDOWN                          int      `json:"knife_kills_library_MURDER_SHOWDOWN"`
				KillsLibraryMURDERSHOWDOWN                               int      `json:"kills_library_MURDER_SHOWDOWN"`
				CoinsPickedupLibraryMURDERSHOWDOWN                       int      `json:"coins_pickedup_library_MURDER_SHOWDOWN"`
				BowKillsLibraryMURDERSHOWDOWN                            int      `json:"bow_kills_library_MURDER_SHOWDOWN"`
				WinsLibraryMURDERSHOWDOWN                                int      `json:"wins_library_MURDER_SHOWDOWN"`
				WinsHypixelWorldMURDERSHOWDOWN                           int      `json:"wins_hypixel_world_MURDER_SHOWDOWN"`
				CoinsPickedupAncientTombMURDERSHOWDOWN                   int      `json:"coins_pickedup_ancient_tomb_MURDER_SHOWDOWN"`
				KnifeKillsAncientTombMURDERSHOWDOWN                      int      `json:"knife_kills_ancient_tomb_MURDER_SHOWDOWN"`
				GamesAncientTombMURDERSHOWDOWN                           int      `json:"games_ancient_tomb_MURDER_SHOWDOWN"`
				BowKillsAncientTombMURDERSHOWDOWN                        int      `json:"bow_kills_ancient_tomb_MURDER_SHOWDOWN"`
				KillsAncientTombMURDERSHOWDOWN                           int      `json:"kills_ancient_tomb_MURDER_SHOWDOWN"`
				DeathsAncientTombMURDERSHOWDOWN                          int      `json:"deaths_ancient_tomb_MURDER_SHOWDOWN"`
				ShowdownPotgHollywood                                    int      `json:"showdown_potg_hollywood"`
				GamesHollywoodMURDERSHOWDOWN                             int      `json:"games_hollywood_MURDER_SHOWDOWN"`
				WinsHollywoodMURDERSHOWDOWN                              int      `json:"wins_hollywood_MURDER_SHOWDOWN"`
				BowKillsHollywoodMURDERSHOWDOWN                          int      `json:"bow_kills_hollywood_MURDER_SHOWDOWN"`
				ShowdownPotgHollywoodMURDERSHOWDOWN                      int      `json:"showdown_potg_hollywood_MURDER_SHOWDOWN"`
				KnifeKillsHollywoodMURDERSHOWDOWN                        int      `json:"knife_kills_hollywood_MURDER_SHOWDOWN"`
				KillsHollywoodMURDERSHOWDOWN                             int      `json:"kills_hollywood_MURDER_SHOWDOWN"`
				CoinsPickedupHollywoodMURDERSHOWDOWN                     int      `json:"coins_pickedup_hollywood_MURDER_SHOWDOWN"`
				QuickestShowdownWinTimeSecondsHypixelWorld               int      `json:"quickest_showdown_win_time_seconds_hypixel_world"`
				QuickestShowdownWinTimeSecondsHypixelWorldMURDERSHOWDOWN int      `json:"quickest_showdown_win_time_seconds_hypixel_world_MURDER_SHOWDOWN"`
				CoinsPickedupHeadquartersMURDERSHOWDOWN                  int      `json:"coins_pickedup_headquarters_MURDER_SHOWDOWN"`
				KnifeKillsHeadquartersMURDERSHOWDOWN                     int      `json:"knife_kills_headquarters_MURDER_SHOWDOWN"`
				BowKillsHeadquartersMURDERSHOWDOWN                       int      `json:"bow_kills_headquarters_MURDER_SHOWDOWN"`
				GamesHeadquartersMURDERSHOWDOWN                          int      `json:"games_headquarters_MURDER_SHOWDOWN"`
				KillsHeadquartersMURDERSHOWDOWN                          int      `json:"kills_headquarters_MURDER_SHOWDOWN"`
				WinsHeadquartersMURDERSHOWDOWN                           int      `json:"wins_headquarters_MURDER_SHOWDOWN"`
				DeathsHeadquartersMURDERSHOWDOWN                         int      `json:"deaths_headquarters_MURDER_SHOWDOWN"`
				BowKillsHeadquarters                                     int      `json:"bow_kills_headquarters"`
				DeathsHollywoodMURDERSHOWDOWN                            int      `json:"deaths_hollywood_MURDER_SHOWDOWN"`
				QuickestDetectiveWinTimeSecondsMountainMURDERCLASSIC     int      `json:"quickest_detective_win_time_seconds_mountain_MURDER_CLASSIC"`
				QuickestDetectiveWinTimeSecondsMURDERCLASSIC             int      `json:"quickest_detective_win_time_seconds_MURDER_CLASSIC"`
				QuickestDetectiveWinTimeSecondsMountain                  int      `json:"quickest_detective_win_time_seconds_mountain"`
				QuickestDetectiveWinTimeSeconds                          int      `json:"quickest_detective_win_time_seconds"`
				CoinsPickedupMountainMURDERCLASSIC                       int      `json:"coins_pickedup_mountain_MURDER_CLASSIC"`
				DetectiveWinsMountain                                    int      `json:"detective_wins_mountain"`
				GamesMountainMURDERCLASSIC                               int      `json:"games_mountain_MURDER_CLASSIC"`
				WinsMountainMURDERCLASSIC                                int      `json:"wins_mountain_MURDER_CLASSIC"`
				DetectiveWinsMountainMURDERCLASSIC                       int      `json:"detective_wins_mountain_MURDER_CLASSIC"`
				DetectiveWinsTowerfall                                   int      `json:"detective_wins_towerfall"`
				DetectiveWinsTowerfallMURDERCLASSIC                      int      `json:"detective_wins_towerfall_MURDER_CLASSIC"`
				QuickestMurdererWinTimeSecondsLibraryMURDERCLASSIC       int      `json:"quickest_murderer_win_time_seconds_library_MURDER_CLASSIC"`
				QuickestMurdererWinTimeSeconds                           int      `json:"quickest_murderer_win_time_seconds"`
				QuickestMurdererWinTimeSecondsLibrary                    int      `json:"quickest_murderer_win_time_seconds_library"`
				QuickestMurdererWinTimeSecondsMURDERCLASSIC              int      `json:"quickest_murderer_win_time_seconds_MURDER_CLASSIC"`
				ThrownKnifeKillsLibraryMURDERCLASSIC                     int      `json:"thrown_knife_kills_library_MURDER_CLASSIC"`
				DeathsLibraryMURDERCLASSIC                               int      `json:"deaths_library_MURDER_CLASSIC"`
				ChestHistoryNew                                          []string `json:"chest_history_new"`
				CoinsPickedupSanPeratico                                 int      `json:"coins_pickedup_san_peratico"`
				WinsSanPeratico                                          int      `json:"wins_san_peratico"`
				KnifeKillsSanPeraticoMURDERSHOWDOWN                      int      `json:"knife_kills_san_peratico_MURDER_SHOWDOWN"`
				KillsSanPeratico                                         int      `json:"kills_san_peratico"`
				BowKillsSanPeraticoMURDERSHOWDOWN                        int      `json:"bow_kills_san_peratico_MURDER_SHOWDOWN"`
				KillsSanPeraticoMURDERSHOWDOWN                           int      `json:"kills_san_peratico_MURDER_SHOWDOWN"`
				GamesSanPeratico                                         int      `json:"games_san_peratico"`
				BowKillsSanPeratico                                      int      `json:"bow_kills_san_peratico"`
				DeathsSanPeraticoMURDERSHOWDOWN                          int      `json:"deaths_san_peratico_MURDER_SHOWDOWN"`
				KnifeKillsSanPeratico                                    int      `json:"knife_kills_san_peratico"`
				WinsSanPeraticoMURDERSHOWDOWN                            int      `json:"wins_san_peratico_MURDER_SHOWDOWN"`
				GamesSanPeraticoMURDERSHOWDOWN                           int      `json:"games_san_peratico_MURDER_SHOWDOWN"`
				CoinsPickedupSanPeraticoMURDERSHOWDOWN                   int      `json:"coins_pickedup_san_peratico_MURDER_SHOWDOWN"`
				DeathsSanPeratico                                        int      `json:"deaths_san_peratico"`
				ShowdownPotgSanPeratico                                  int      `json:"showdown_potg_san_peratico"`
				ShowdownPotgSanPeraticoMURDERSHOWDOWN                    int      `json:"showdown_potg_san_peratico_MURDER_SHOWDOWN"`
				ShowdownPotgAncientTombMURDERSHOWDOWN                    int      `json:"showdown_potg_ancient_tomb_MURDER_SHOWDOWN"`
				WinsAncientTombMURDERSHOWDOWN                            int      `json:"wins_ancient_tomb_MURDER_SHOWDOWN"`
				ShowdownPotgAncientTomb                                  int      `json:"showdown_potg_ancient_tomb"`
				ActiveGesture                                            string   `json:"active_gesture"`
				GamesGoldRushMURDERSHOWDOWN                              int      `json:"games_gold_rush_MURDER_SHOWDOWN"`
				WinsGoldRushMURDERSHOWDOWN                               int      `json:"wins_gold_rush_MURDER_SHOWDOWN"`
				KillsGoldRushMURDERSHOWDOWN                              int      `json:"kills_gold_rush_MURDER_SHOWDOWN"`
				CoinsPickedupGoldRushMURDERSHOWDOWN                      int      `json:"coins_pickedup_gold_rush_MURDER_SHOWDOWN"`
				DeathsGoldRushMURDERSHOWDOWN                             int      `json:"deaths_gold_rush_MURDER_SHOWDOWN"`
				KnifeKillsGoldRushMURDERSHOWDOWN                         int      `json:"knife_kills_gold_rush_MURDER_SHOWDOWN"`
				BowKillsGoldRushMURDERSHOWDOWN                           int      `json:"bow_kills_gold_rush_MURDER_SHOWDOWN"`
				QuickestShowdownWinTimeSecondsSanPeratico                int      `json:"quickest_showdown_win_time_seconds_san_peratico"`
				QuickestShowdownWinTimeSecondsSanPeraticoMURDERSHOWDOWN  int      `json:"quickest_showdown_win_time_seconds_san_peratico_MURDER_SHOWDOWN"`
				WinsSanPeraticoMURDERCLASSIC                             int      `json:"wins_san_peratico_MURDER_CLASSIC"`
				GamesSanPeraticoMURDERCLASSIC                            int      `json:"games_san_peratico_MURDER_CLASSIC"`
				CoinsPickedupSanPeraticoMURDERCLASSIC                    int      `json:"coins_pickedup_san_peratico_MURDER_CLASSIC"`
				BowKillsTowerfallMURDERCLASSIC                           int      `json:"bow_kills_towerfall_MURDER_CLASSIC"`
				WasHeroTowerfall                                         int      `json:"was_hero_towerfall"`
				WasHeroTowerfallMURDERCLASSIC                            int      `json:"was_hero_towerfall_MURDER_CLASSIC"`
				DeathsHeadquartersMURDERCLASSIC                          int      `json:"deaths_headquarters_MURDER_CLASSIC"`
				KillsAsMurdererHeadquarters                              int      `json:"kills_as_murderer_headquarters"`
				KillsAsMurdererHeadquartersMURDERCLASSIC                 int      `json:"kills_as_murderer_headquarters_MURDER_CLASSIC"`
				KillsHeadquartersMURDERCLASSIC                           int      `json:"kills_headquarters_MURDER_CLASSIC"`
				KnifeKillsHeadquartersMURDERCLASSIC                      int      `json:"knife_kills_headquarters_MURDER_CLASSIC"`
				ThrownKnifeKillsHeadquartersMURDERCLASSIC                int      `json:"thrown_knife_kills_headquarters_MURDER_CLASSIC"`
				BowKillsArchivesTopFloor                                 int      `json:"bow_kills_archives_top_floor"`
				BowKillsArchivesTopFloorMURDERCLASSIC                    int      `json:"bow_kills_archives_top_floor_MURDER_CLASSIC"`
				CoinsPickedupArchivesTopFloor                            int      `json:"coins_pickedup_archives_top_floor"`
				CoinsPickedupArchivesTopFloorMURDERCLASSIC               int      `json:"coins_pickedup_archives_top_floor_MURDER_CLASSIC"`
				GamesArchivesTopFloor                                    int      `json:"games_archives_top_floor"`
				GamesArchivesTopFloorMURDERCLASSIC                       int      `json:"games_archives_top_floor_MURDER_CLASSIC"`
				KillsArchivesTopFloor                                    int      `json:"kills_archives_top_floor"`
				KillsArchivesTopFloorMURDERCLASSIC                       int      `json:"kills_archives_top_floor_MURDER_CLASSIC"`
				WasHeroArchivesTopFloor                                  int      `json:"was_hero_archives_top_floor"`
				WasHeroArchivesTopFloorMURDERCLASSIC                     int      `json:"was_hero_archives_top_floor_MURDER_CLASSIC"`
				WinsArchivesTopFloor                                     int      `json:"wins_archives_top_floor"`
				WinsArchivesTopFloorMURDERCLASSIC                        int      `json:"wins_archives_top_floor_MURDER_CLASSIC"`
				CoinsPickedupWidowSDen                                   int      `json:"coins_pickedup_widow's_den"`
				CoinsPickedupWidowSDenMURDERCLASSIC                      int      `json:"coins_pickedup_widow's_den_MURDER_CLASSIC"`
				DeathsWidowSDen                                          int      `json:"deaths_widow's_den"`
				DeathsWidowSDenMURDERCLASSIC                             int      `json:"deaths_widow's_den_MURDER_CLASSIC"`
				GamesWidowSDen                                           int      `json:"games_widow's_den"`
				GamesWidowSDenMURDERCLASSIC                              int      `json:"games_widow's_den_MURDER_CLASSIC"`
				WinsWidowSDen                                            int      `json:"wins_widow's_den"`
				WinsWidowSDenMURDERCLASSIC                               int      `json:"wins_widow's_den_MURDER_CLASSIC"`
				CoinsPickedupSanPeraticoV2                               int      `json:"coins_pickedup_san_peratico_v2"`
				CoinsPickedupSanPeraticoV2MURDERCLASSIC                  int      `json:"coins_pickedup_san_peratico_v2_MURDER_CLASSIC"`
				DeathsSanPeraticoV2                                      int      `json:"deaths_san_peratico_v2"`
				DeathsSanPeraticoV2MURDERCLASSIC                         int      `json:"deaths_san_peratico_v2_MURDER_CLASSIC"`
				GamesSanPeraticoV2                                       int      `json:"games_san_peratico_v2"`
				GamesSanPeraticoV2MURDERCLASSIC                          int      `json:"games_san_peratico_v2_MURDER_CLASSIC"`
				WinsSanPeraticoV2                                        int      `json:"wins_san_peratico_v2"`
				WinsSanPeraticoV2MURDERCLASSIC                           int      `json:"wins_san_peratico_v2_MURDER_CLASSIC"`
				MmHalloweenChests                                        int      `json:"mm_halloween_chests"`
				CoinsPickedupMURDERDOUBLEUP                              int      `json:"coins_pickedup_MURDER_DOUBLE_UP"`
				CoinsPickedupMountainMURDERDOUBLEUP                      int      `json:"coins_pickedup_mountain_MURDER_DOUBLE_UP"`
				DeathsMURDERDOUBLEUP                                     int      `json:"deaths_MURDER_DOUBLE_UP"`
				DeathsMountainMURDERDOUBLEUP                             int      `json:"deaths_mountain_MURDER_DOUBLE_UP"`
				GamesMURDERDOUBLEUP                                      int      `json:"games_MURDER_DOUBLE_UP"`
				GamesMountainMURDERDOUBLEUP                              int      `json:"games_mountain_MURDER_DOUBLE_UP"`
				WinsMURDERDOUBLEUP                                       int      `json:"wins_MURDER_DOUBLE_UP"`
				WinsMountainMURDERDOUBLEUP                               int      `json:"wins_mountain_MURDER_DOUBLE_UP"`
				CoinsPickedupArchivesTopFloorMURDERDOUBLEUP              int      `json:"coins_pickedup_archives_top_floor_MURDER_DOUBLE_UP"`
				GamesArchivesTopFloorMURDERDOUBLEUP                      int      `json:"games_archives_top_floor_MURDER_DOUBLE_UP"`
				KnifeKillsMURDERDOUBLEUP                                 int      `json:"knife_kills_MURDER_DOUBLE_UP"`
				KnifeKillsArchivesTopFloor                               int      `json:"knife_kills_archives_top_floor"`
				KnifeKillsArchivesTopFloorMURDERDOUBLEUP                 int      `json:"knife_kills_archives_top_floor_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsMURDERDOUBLEUP                           int      `json:"thrown_knife_kills_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsArchivesTopFloor                         int      `json:"thrown_knife_kills_archives_top_floor"`
				ThrownKnifeKillsArchivesTopFloorMURDERDOUBLEUP           int      `json:"thrown_knife_kills_archives_top_floor_MURDER_DOUBLE_UP"`
				CoinsPickedupSkywayPier                                  int      `json:"coins_pickedup_skyway_pier"`
				CoinsPickedupSkywayPierMURDERCLASSIC                     int      `json:"coins_pickedup_skyway_pier_MURDER_CLASSIC"`
				GamesSkywayPier                                          int      `json:"games_skyway_pier"`
				GamesSkywayPierMURDERCLASSIC                             int      `json:"games_skyway_pier_MURDER_CLASSIC"`
				CoinsPickedupSnowfall                                    int      `json:"coins_pickedup_snowfall"`
				CoinsPickedupSnowfallMURDERCLASSIC                       int      `json:"coins_pickedup_snowfall_MURDER_CLASSIC"`
				GamesSnowfall                                            int      `json:"games_snowfall"`
				GamesSnowfallMURDERCLASSIC                               int      `json:"games_snowfall_MURDER_CLASSIC"`
				CoinsPickedupAncientTombMURDERDOUBLEUP                   int      `json:"coins_pickedup_ancient_tomb_MURDER_DOUBLE_UP"`
				GamesAncientTombMURDERDOUBLEUP                           int      `json:"games_ancient_tomb_MURDER_DOUBLE_UP"`
				KnifeKillsAncientTombMURDERDOUBLEUP                      int      `json:"knife_kills_ancient_tomb_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsAncientTombMURDERDOUBLEUP                int      `json:"thrown_knife_kills_ancient_tomb_MURDER_DOUBLE_UP"`
				CoinsPickedupSanPeraticoV2MURDERDOUBLEUP                 int      `json:"coins_pickedup_san_peratico_v2_MURDER_DOUBLE_UP"`
				GamesSanPeraticoV2MURDERDOUBLEUP                         int      `json:"games_san_peratico_v2_MURDER_DOUBLE_UP"`
				WasHeroMURDERDOUBLEUP                                    int      `json:"was_hero_MURDER_DOUBLE_UP"`
				WasHeroSanPeraticoV2                                     int      `json:"was_hero_san_peratico_v2"`
				WasHeroSanPeraticoV2MURDERDOUBLEUP                       int      `json:"was_hero_san_peratico_v2_MURDER_DOUBLE_UP"`
				KnifeKillsSanPeraticoV2                                  int      `json:"knife_kills_san_peratico_v2"`
				KnifeKillsSanPeraticoV2MURDERDOUBLEUP                    int      `json:"knife_kills_san_peratico_v2_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsSanPeraticoV2                            int      `json:"thrown_knife_kills_san_peratico_v2"`
				ThrownKnifeKillsSanPeraticoV2MURDERDOUBLEUP              int      `json:"thrown_knife_kills_san_peratico_v2_MURDER_DOUBLE_UP"`
				BowKillsMURDERDOUBLEUP                                   int      `json:"bow_kills_MURDER_DOUBLE_UP"`
				BowKillsHeadquartersMURDERDOUBLEUP                       int      `json:"bow_kills_headquarters_MURDER_DOUBLE_UP"`
				CoinsPickedupHeadquartersMURDERDOUBLEUP                  int      `json:"coins_pickedup_headquarters_MURDER_DOUBLE_UP"`
				GamesHeadquartersMURDERDOUBLEUP                          int      `json:"games_headquarters_MURDER_DOUBLE_UP"`
				KnifeKillsHeadquartersMURDERDOUBLEUP                     int      `json:"knife_kills_headquarters_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsHeadquartersMURDERDOUBLEUP               int      `json:"thrown_knife_kills_headquarters_MURDER_DOUBLE_UP"`
				BowKillsSanPeraticoV2                                    int      `json:"bow_kills_san_peratico_v2"`
				BowKillsSanPeraticoV2MURDERDOUBLEUP                      int      `json:"bow_kills_san_peratico_v2_MURDER_DOUBLE_UP"`
				WinsSkywayPier                                           int      `json:"wins_skyway_pier"`
				WinsSkywayPierMURDERCLASSIC                              int      `json:"wins_skyway_pier_MURDER_CLASSIC"`
				CoinsPickedupCruiseShipMURDERDOUBLEUP                    int      `json:"coins_pickedup_cruise_ship_MURDER_DOUBLE_UP"`
				GamesCruiseShipMURDERDOUBLEUP                            int      `json:"games_cruise_ship_MURDER_DOUBLE_UP"`
				KillsMURDERDOUBLEUP                                      int      `json:"kills_MURDER_DOUBLE_UP"`
				KillsAsMurdererMURDERDOUBLEUP                            int      `json:"kills_as_murderer_MURDER_DOUBLE_UP"`
				KillsAsMurdererCruiseShipMURDERDOUBLEUP                  int      `json:"kills_as_murderer_cruise_ship_MURDER_DOUBLE_UP"`
				KillsCruiseShipMURDERDOUBLEUP                            int      `json:"kills_cruise_ship_MURDER_DOUBLE_UP"`
				KnifeKillsCruiseShipMURDERDOUBLEUP                       int      `json:"knife_kills_cruise_ship_MURDER_DOUBLE_UP"`
				MurdererWinsMURDERDOUBLEUP                               int      `json:"murderer_wins_MURDER_DOUBLE_UP"`
				MurdererWinsCruiseShipMURDERDOUBLEUP                     int      `json:"murderer_wins_cruise_ship_MURDER_DOUBLE_UP"`
				QuickestMurdererWinTimeSecondsMURDERDOUBLEUP             int      `json:"quickest_murderer_win_time_seconds_MURDER_DOUBLE_UP"`
				QuickestMurdererWinTimeSecondsCruiseShip                 int      `json:"quickest_murderer_win_time_seconds_cruise_ship"`
				QuickestMurdererWinTimeSecondsCruiseShipMURDERDOUBLEUP   int      `json:"quickest_murderer_win_time_seconds_cruise_ship_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsCruiseShipMURDERDOUBLEUP                 int      `json:"thrown_knife_kills_cruise_ship_MURDER_DOUBLE_UP"`
				WinsCruiseShipMURDERDOUBLEUP                             int      `json:"wins_cruise_ship_MURDER_DOUBLE_UP"`
				BowKillsMountainMURDERDOUBLEUP                           int      `json:"bow_kills_mountain_MURDER_DOUBLE_UP"`
				KillsMountainMURDERDOUBLEUP                              int      `json:"kills_mountain_MURDER_DOUBLE_UP"`
				CoinsPickedupGoldRushMURDERDOUBLEUP                      int      `json:"coins_pickedup_gold_rush_MURDER_DOUBLE_UP"`
				GamesGoldRushMURDERDOUBLEUP                              int      `json:"games_gold_rush_MURDER_DOUBLE_UP"`
				WasHeroGoldRushMURDERDOUBLEUP                            int      `json:"was_hero_gold_rush_MURDER_DOUBLE_UP"`
				WinsGoldRushMURDERDOUBLEUP                               int      `json:"wins_gold_rush_MURDER_DOUBLE_UP"`
			} `json:"MurderMystery"`
			BuildBattle struct {
				WinsSoloNormal     int      `json:"wins_solo_normal"`
				WinsTeamsNormal    int      `json:"wins_teams_normal"`
				Wins               int      `json:"wins"`
				BuildbattleLoadout []string `json:"buildbattle_loadout"`
				Coins              int      `json:"coins"`
				WinsGuessTheBuild  int      `json:"wins_guess_the_build"`
				GamesPlayed        int      `json:"games_played"`
				Score              int      `json:"score"`
				MonthlyCoinsA      int      `json:"monthly_coins_a"`
				WeeklyCoinsB       int      `json:"weekly_coins_b"`
				CorrectGuesses     int      `json:"correct_guesses"`
				SoloMostPoints     int      `json:"solo_most_points"`
				WeeklyCoinsA       int      `json:"weekly_coins_a"`
				TotalVotes         int      `json:"total_votes"`
				Packages           []string `json:"packages"`
				SuperVotes         int      `json:"super_votes"`
				NewSelectedHat     string   `json:"new_selected_hat"`
				MonthlyCoinsB      int      `json:"monthly_coins_b"`
			} `json:"BuildBattle"`
			Duels struct {
				GamesPlayedDuels               int    `json:"games_played_duels"`
				MeleeSwings                    int    `json:"melee_swings"`
				MeleeHits                      int    `json:"melee_hits"`
				Wins                           int    `json:"wins"`
				RoundsPlayed                   int    `json:"rounds_played"`
				UhcTournamentMeleeHits         int    `json:"uhc_tournament_melee_hits"`
				UhcTournamentHealthRegenerated int    `json:"uhc_tournament_health_regenerated"`
				UhcTournamentDamageDealt       int    `json:"uhc_tournament_damage_dealt"`
				BowShots                       int    `json:"bow_shots"`
				UhcTournamentKills             int    `json:"uhc_tournament_kills"`
				Kills                          int    `json:"kills"`
				BowHits                        int    `json:"bow_hits"`
				UhcTournamentMeleeSwings       int    `json:"uhc_tournament_melee_swings"`
				UhcTournamentRoundsPlayed      int    `json:"uhc_tournament_rounds_played"`
				HealthRegenerated              int    `json:"health_regenerated"`
				UhcTournamentBowHits           int    `json:"uhc_tournament_bow_hits"`
				DamageDealt                    int    `json:"damage_dealt"`
				UhcTournamentBowShots          int    `json:"uhc_tournament_bow_shots"`
				UhcTournamentWins              int    `json:"uhc_tournament_wins"`
				Coins                          int    `json:"coins"`
				Losses                         int    `json:"losses"`
				UhcTournamentLosses            int    `json:"uhc_tournament_losses"`
				UhcTournamentDeaths            int    `json:"uhc_tournament_deaths"`
				Deaths                         int    `json:"deaths"`
				Selected1                      string `json:"selected_1"`
				Selected2                      string `json:"selected_2"`
				KitMenuOption                  string `json:"kit_menu_option"`
				BlitzDuelsKit                  string `json:"blitz_duels_kit"`
				ShowLbOption                   string `json:"show_lb_option"`
				RematchOption1                 string `json:"rematch_option_1"`
				BlitzDuelMeleeHits             int    `json:"blitz_duel_melee_hits"`
				BlitzDuelMeleeSwings           int    `json:"blitz_duel_melee_swings"`
				BlitzDuelRoundsPlayed          int    `json:"blitz_duel_rounds_played"`
				BlitzDuelHealthRegenerated     int    `json:"blitz_duel_health_regenerated"`
				BlitzDuelDamageDealt           int    `json:"blitz_duel_damage_dealt"`
				BlitzDuelBowShots              int    `json:"blitz_duel_bow_shots"`
				BlitzDuelBowHits               int    `json:"blitz_duel_bow_hits"`
				BowspleefDuelRoundsPlayed      int    `json:"bowspleef_duel_rounds_played"`
				BowspleefDuelBowShots          int    `json:"bowspleef_duel_bow_shots"`
				DuelsRecentlyPlayed            string `json:"duels_recently_played"`
				SwDuelsKit                     string `json:"sw_duels_kit"`
				SwTournamentRoundsPlayed       int    `json:"sw_tournament_rounds_played"`
				SwTournamentMeleeHits          int    `json:"sw_tournament_melee_hits"`
				SwTournamentDamageDealt        int    `json:"sw_tournament_damage_dealt"`
				SwTournamentMeleeSwings        int    `json:"sw_tournament_melee_swings"`
				SwTournamentHealthRegenerated  int    `json:"sw_tournament_health_regenerated"`
				SwTournamentKills              int    `json:"sw_tournament_kills"`
				SwTournamentDeaths             int    `json:"sw_tournament_deaths"`
				CurrentBlitzWinstreak          int    `json:"current_blitz_winstreak"`
				CurrentWinstreak               int    `json:"current_winstreak"`
				BestOverallWinstreak           int    `json:"best_overall_winstreak"`
				BestBlitzWinstreak             int    `json:"best_blitz_winstreak"`
				BlitzDuelWins                  int    `json:"blitz_duel_wins"`
				BlitzDuelKills                 int    `json:"blitz_duel_kills"`
				DuelsWinstreakBestBlitzDuel    int    `json:"duels_winstreak_best_blitz_duel"`
				BlitzDuelLosses                int    `json:"blitz_duel_losses"`
				BlitzDuelDeaths                int    `json:"blitz_duel_deaths"`
				BlitzDuelLayoutHunter5         struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
					Num2 string `json:"2"`
					Num3 string `json:"3"`
					Num4 string `json:"4"`
					Num5 string `json:"5"`
					Num6 string `json:"6"`
				} `json:"blitz_duel_layout_hunter_5"`
				ChatEnabled                       string   `json:"chat_enabled"`
				BridgeDoublesWins                 int      `json:"bridge_doubles_wins"`
				BridgeDuelWins                    int      `json:"bridge_duel_wins"`
				BridgeDuelDeaths                  int      `json:"bridge_duel_deaths"`
				BridgeDoublesDeaths               int      `json:"bridge_doubles_deaths"`
				Goals                             int      `json:"goals"`
				BridgeDuelGoals                   int      `json:"bridge_duel_goals"`
				BridgeDoublesGoals                int      `json:"bridge_doubles_goals"`
				BridgeDoublesRoundsPlayed         int      `json:"bridge_doubles_rounds_played"`
				BridgeDoublesKills                int      `json:"bridge_doubles_kills"`
				BridgeDuelKills                   int      `json:"bridge_duel_kills"`
				BridgeDuelRoundsPlayed            int      `json:"bridge_duel_rounds_played"`
				Packages                          []string `json:"packages"`
				AllModesRookieTitlePrestige       int      `json:"all_modes_rookie_title_prestige"`
				NoDebuffRookieTitlePrestige       int      `json:"no_debuff_rookie_title_prestige"`
				UhcRookieTitlePrestige            int      `json:"uhc_rookie_title_prestige"`
				SkywarsRookieTitlePrestige        int      `json:"skywars_rookie_title_prestige"`
				SumoRookieTitlePrestige           int      `json:"sumo_rookie_title_prestige"`
				TournamentRookieTitlePrestige     int      `json:"tournament_rookie_title_prestige"`
				MegaWallsRookieTitlePrestige      int      `json:"mega_walls_rookie_title_prestige"`
				TntGamesRookieTitlePrestige       int      `json:"tnt_games_rookie_title_prestige"`
				BlitzRookieTitlePrestige          int      `json:"blitz_rookie_title_prestige"`
				BridgeRookieTitlePrestige         int      `json:"bridge_rookie_title_prestige"`
				ComboRookieTitlePrestige          int      `json:"combo_rookie_title_prestige"`
				BowRookieTitlePrestige            int      `json:"bow_rookie_title_prestige"`
				OpRookieTitlePrestige             int      `json:"op_rookie_title_prestige"`
				ClassicRookieTitlePrestige        int      `json:"classic_rookie_title_prestige"`
				Selected2New                      string   `json:"selected_2_new"`
				Selected1New                      string   `json:"selected_1_new"`
				BridgeMapWins                     []string `json:"bridgeMapWins"`
				MapsWonOn                         []string `json:"maps_won_on"`
				CurrentBridgeWinstreak            int      `json:"current_bridge_winstreak"`
				BestWinstreakModeBridgeDoubles    int      `json:"best_winstreak_mode_bridge_doubles"`
				BestBridgeWinstreak               int      `json:"best_bridge_winstreak"`
				CurrentWinstreakModeBridgeDoubles int      `json:"current_winstreak_mode_bridge_doubles"`
				DuelsChests                       int      `json:"duels_chests"`
				BridgeDoublesBridgeDeaths         int      `json:"bridge_doubles_bridge_deaths"`
				BlocksPlaced                      int      `json:"blocks_placed"`
				BridgeDoublesBridgeKills          int      `json:"bridge_doubles_bridge_kills"`
				BridgeDoublesDamageDealt          int      `json:"bridge_doubles_damage_dealt"`
				BridgeDoublesHealthRegenerated    int      `json:"bridge_doubles_health_regenerated"`
				BridgeDoublesBlocksPlaced         int      `json:"bridge_doubles_blocks_placed"`
				BridgeDoublesMeleeSwings          int      `json:"bridge_doubles_melee_swings"`
				BridgeDoublesMeleeHits            int      `json:"bridge_doubles_melee_hits"`
				BridgeKills                       int      `json:"bridge_kills"`
				BridgeDeaths                      int      `json:"bridge_deaths"`
				DuelsWinstreakBestBridgeDoubles   int      `json:"duels_winstreak_best_bridge_doubles"`
				BridgeDoublesBowShots             int      `json:"bridge_doubles_bow_shots"`
				BridgeDoublesBowHits              int      `json:"bridge_doubles_bow_hits"`
				BridgeFourRoundsPlayed            int      `json:"bridge_four_rounds_played"`
				BridgeFourDamageDealt             int      `json:"bridge_four_damage_dealt"`
				BridgeFourBlocksPlaced            int      `json:"bridge_four_blocks_placed"`
				BridgeFourBowShots                int      `json:"bridge_four_bow_shots"`
				BridgeFourMeleeSwings             int      `json:"bridge_four_melee_swings"`
				BridgeFourHealthRegenerated       int      `json:"bridge_four_health_regenerated"`
				BridgeFourMeleeHits               int      `json:"bridge_four_melee_hits"`
				SwDuelsKitNew3                    string   `json:"sw_duels_kit_new3"`
				CurrentWinstreakModeBridgeDuel    int      `json:"current_winstreak_mode_bridge_duel"`
				BestWinstreakModeBridgeDuel       int      `json:"best_winstreak_mode_bridge_duel"`
				BridgeDuelHealthRegenerated       int      `json:"bridge_duel_health_regenerated"`
				BridgeDuelDamageDealt             int      `json:"bridge_duel_damage_dealt"`
				BridgeDuelMeleeSwings             int      `json:"bridge_duel_melee_swings"`
				BridgeDuelBowHits                 int      `json:"bridge_duel_bow_hits"`
				BridgeDuelBridgeDeaths            int      `json:"bridge_duel_bridge_deaths"`
				BridgeDuelBridgeKills             int      `json:"bridge_duel_bridge_kills"`
				BridgeDuelMeleeHits               int      `json:"bridge_duel_melee_hits"`
				BridgeDuelBlocksPlaced            int      `json:"bridge_duel_blocks_placed"`
				BridgeDuelBowShots                int      `json:"bridge_duel_bow_shots"`
				DuelsWinstreakBestBridgeDuel      int      `json:"duels_winstreak_best_bridge_duel"`
				LayoutBridgeDuelLayout            struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
					Num2 string `json:"2"`
					Num3 string `json:"3"`
					Num4 string `json:"4"`
					Num5 string `json:"5"`
					Num7 string `json:"7"`
					Num8 string `json:"8"`
				} `json:"layout_bridge_duel_layout"`
				DuelsOpenedCommons            int      `json:"Duels_openedCommons"`
				DuelsOpenedChests             int      `json:"Duels_openedChests"`
				DuelsChestHistory             []string `json:"duels_chest_history"`
				DuelsOpenedRares              int      `json:"Duels_openedRares"`
				BestWinstreakModeBlitzDuel    int      `json:"best_winstreak_mode_blitz_duel"`
				CurrentWinstreakModeBlitzDuel int      `json:"current_winstreak_mode_blitz_duel"`
				ArcherKitWins                 int      `json:"archer_kit_wins"`
				BlitzDuelKitWins              int      `json:"blitz_duel_kit_wins"`
				BlitzDuelArcherKitWins        int      `json:"blitz_duel_archer_kit_wins"`
				KitWins                       int      `json:"kit_wins"`
				PigmanKitWins                 int      `json:"pigman_kit_wins"`
				BlitzDuelBlocksPlaced         int      `json:"blitz_duel_blocks_placed"`
				BlitzDuelPigmanKitWins        int      `json:"blitz_duel_pigman_kit_wins"`
				GolemKitWins                  int      `json:"golem_kit_wins"`
				BlitzDuelGolemKitWins         int      `json:"blitz_duel_golem_kit_wins"`
				BlitzDuelNecromancerKitWins   int      `json:"blitz_duel_necromancer_kit_wins"`
				NecromancerKitWins            int      `json:"necromancer_kit_wins"`
				ProgressMode                  string   `json:"progress_mode"`
				CurrentWinstreakModeComboDuel int      `json:"current_winstreak_mode_combo_duel"`
				CurrentComboWinstreak         int      `json:"current_combo_winstreak"`
				ComboDuelMeleeHits            int      `json:"combo_duel_melee_hits"`
				ComboDuelLosses               int      `json:"combo_duel_losses"`
				ComboDuelRoundsPlayed         int      `json:"combo_duel_rounds_played"`
				ComboDuelHealthRegenerated    int      `json:"combo_duel_health_regenerated"`
				GoldenApplesEaten             int      `json:"golden_apples_eaten"`
				ComboDuelDeaths               int      `json:"combo_duel_deaths"`
				ComboDuelMeleeSwings          int      `json:"combo_duel_melee_swings"`
				ComboDuelGoldenApplesEaten    int      `json:"combo_duel_golden_apples_eaten"`
				LeaderboardPageWinStreak      int      `json:"leaderboardPage_win_streak"`
				SumoDuelMeleeSwings           int      `json:"sumo_duel_melee_swings"`
				SumoDuelRoundsPlayed          int      `json:"sumo_duel_rounds_played"`
				SumoDuelMeleeHits             int      `json:"sumo_duel_melee_hits"`
				UhcDuelHealthRegenerated      int      `json:"uhc_duel_health_regenerated"`
				UhcDuelBowShots               int      `json:"uhc_duel_bow_shots"`
				UhcDuelMeleeSwings            int      `json:"uhc_duel_melee_swings"`
				UhcDuelBowHits                int      `json:"uhc_duel_bow_hits"`
				UhcDuelRoundsPlayed           int      `json:"uhc_duel_rounds_played"`
				UhcDuelDamageDealt            int      `json:"uhc_duel_damage_dealt"`
				UhcDuelBlocksPlaced           int      `json:"uhc_duel_blocks_placed"`
				UhcDuelGoldenApplesEaten      int      `json:"uhc_duel_golden_apples_eaten"`
				UhcDuelMeleeHits              int      `json:"uhc_duel_melee_hits"`
				CurrentSumoWinstreak          int      `json:"current_sumo_winstreak"`
				CurrentWinstreakModeSumoDuel  int      `json:"current_winstreak_mode_sumo_duel"`
				SumoDuelLosses                int      `json:"sumo_duel_losses"`
				DuelsRecentlyPlayed2          string   `json:"duels_recently_played2"`
				CurrentUhcWinstreak           int      `json:"current_uhc_winstreak"`
				CurrentWinstreakModeUhcDuel   int      `json:"current_winstreak_mode_uhc_duel"`
				UhcDuelDeaths                 int      `json:"uhc_duel_deaths"`
				UhcDuelLosses                 int      `json:"uhc_duel_losses"`
				SwDuelBlocksPlaced            int      `json:"sw_duel_blocks_placed"`
				SwDuelDamageDealt             int      `json:"sw_duel_damage_dealt"`
				SwDuelHealthRegenerated       int      `json:"sw_duel_health_regenerated"`
				SwDuelMeleeHits               int      `json:"sw_duel_melee_hits"`
				SwDuelMeleeSwings             int      `json:"sw_duel_melee_swings"`
				SwDuelRoundsPlayed            int      `json:"sw_duel_rounds_played"`
				BridgeFourBowHits             int      `json:"bridge_four_bow_hits"`
				BestSkywarsWinstreak          int      `json:"best_skywars_winstreak"`
				CurrentWinstreakModeSwDoubles int      `json:"current_winstreak_mode_sw_doubles"`
				CurrentSkywarsWinstreak       int      `json:"current_skywars_winstreak"`
				BestWinstreakModeSwDoubles    int      `json:"best_winstreak_mode_sw_doubles"`
				ScoutKitWins                  int      `json:"scout_kit_wins"`
				SwDoublesDamageDealt          int      `json:"sw_doubles_damage_dealt"`
				SwDoublesHealthRegenerated    int      `json:"sw_doubles_health_regenerated"`
				SwDoublesKills                int      `json:"sw_doubles_kills"`
				SwDoublesKitWins              int      `json:"sw_doubles_kit_wins"`
				SwDoublesMeleeHits            int      `json:"sw_doubles_melee_hits"`
				SwDoublesMeleeSwings          int      `json:"sw_doubles_melee_swings"`
				SwDoublesRoundsPlayed         int      `json:"sw_doubles_rounds_played"`
				SwDoublesScoutKitWins         int      `json:"sw_doubles_scout_kit_wins"`
				SwDoublesWins                 int      `json:"sw_doubles_wins"`
				LayoutUhcDuelLayout           struct {
					Num0  string `json:"0"`
					Num1  string `json:"1"`
					Num2  string `json:"2"`
					Num3  string `json:"3"`
					Num4  string `json:"4"`
					Num5  string `json:"5"`
					Num6  string `json:"6"`
					Num7  string `json:"7"`
					Num8  string `json:"8"`
					Num9  string `json:"9"`
					Num33 string `json:"33"`
					Num34 string `json:"34"`
					Num35 string `json:"35"`
				} `json:"layout_uhc_duel_layout"`
				SwDoublesBlocksPlaced           int  `json:"sw_doubles_blocks_placed"`
				DuelsWinstreakBestSwDoubles     int  `json:"duels_winstreak_best_sw_doubles"`
				SwDoublesBowHits                int  `json:"sw_doubles_bow_hits"`
				SwDoublesBowShots               int  `json:"sw_doubles_bow_shots"`
				SwDoublesDeaths                 int  `json:"sw_doubles_deaths"`
				SwDoublesLosses                 int  `json:"sw_doubles_losses"`
				OpDuelDamageDealt               int  `json:"op_duel_damage_dealt"`
				OpDuelHealthRegenerated         int  `json:"op_duel_health_regenerated"`
				OpDuelMeleeHits                 int  `json:"op_duel_melee_hits"`
				OpDuelMeleeSwings               int  `json:"op_duel_melee_swings"`
				OpDuelRoundsPlayed              int  `json:"op_duel_rounds_played"`
				ChallengesEnabled               bool `json:"challenges_enabled"`
				BowDuelBowHits                  int  `json:"bow_duel_bow_hits"`
				BowDuelBowShots                 int  `json:"bow_duel_bow_shots"`
				BowDuelDamageDealt              int  `json:"bow_duel_damage_dealt"`
				BowDuelHealthRegenerated        int  `json:"bow_duel_health_regenerated"`
				BowDuelRoundsPlayed             int  `json:"bow_duel_rounds_played"`
				BestClassicWinstreak            int  `json:"best_classic_winstreak"`
				BestWinstreakModeClassicDuel    int  `json:"best_winstreak_mode_classic_duel"`
				CurrentWinstreakModeClassicDuel int  `json:"current_winstreak_mode_classic_duel"`
				CurrentClassicWinstreak         int  `json:"current_classic_winstreak"`
				ClassicDuelBowHits              int  `json:"classic_duel_bow_hits"`
				ClassicDuelBowShots             int  `json:"classic_duel_bow_shots"`
				ClassicDuelDamageDealt          int  `json:"classic_duel_damage_dealt"`
				ClassicDuelHealthRegenerated    int  `json:"classic_duel_health_regenerated"`
				ClassicDuelKills                int  `json:"classic_duel_kills"`
				ClassicDuelMeleeHits            int  `json:"classic_duel_melee_hits"`
				ClassicDuelMeleeSwings          int  `json:"classic_duel_melee_swings"`
				ClassicDuelRoundsPlayed         int  `json:"classic_duel_rounds_played"`
				ClassicDuelWins                 int  `json:"classic_duel_wins"`
				DuelsWinstreakBestClassicDuel   int  `json:"duels_winstreak_best_classic_duel"`
				ClassicDuelDeaths               int  `json:"classic_duel_deaths"`
				ClassicDuelLosses               int  `json:"classic_duel_losses"`
				BestUhcWinstreak                int  `json:"best_uhc_winstreak"`
				BestWinstreakModeUhcDuel        int  `json:"best_winstreak_mode_uhc_duel"`
				UhcDuelKills                    int  `json:"uhc_duel_kills"`
				UhcDuelWins                     int  `json:"uhc_duel_wins"`
				DuelsWinstreakBestUhcDuel       int  `json:"duels_winstreak_best_uhc_duel"`
				BridgeDuelLosses                int  `json:"bridge_duel_losses"`
			} `json:"Duels"`
			Pit struct {
				Profile struct {
					OutgoingOffers  []interface{} `json:"outgoing_offers"`
					ContractChoices interface{}   `json:"contract_choices"`
					LastSave        int64         `json:"last_save"`
					KingQuest       struct {
						Kills  int `json:"kills"`
						Renown int `json:"renown"`
					} `json:"king_quest"`
					HatColor        int           `json:"hat_color"`
					TradeTimestamps []interface{} `json:"trade_timestamps"`
					Unlocks1        []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"unlocks_1"`
					DeathRecaps struct {
						Type int   `json:"type"`
						Data []int `json:"data"`
					} `json:"death_recaps"`
					InvEnderchest struct {
						Type int   `json:"type"`
						Data []int `json:"data"`
					} `json:"inv_enderchest"`
					Unlocks2 []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"unlocks_2"`
					Unlocks3 []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"unlocks_3"`
					Unlocks4 []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"unlocks_4"`
					Cash     float64 `json:"cash"`
					Unlocks5 []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"unlocks_5"`
					Unlocks6 []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"unlocks_6"`
					Unlocks7 []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"unlocks_7"`
					LeaderboardStats struct {
						PitRaffleJackpot2019Summer int `json:"Pit_raffle_jackpot_2019_summer"`
						PitKotlTime2019Spring      int `json:"Pit_kotl_time_2019_spring"`
						PitRaffleTickets2019Summer int `json:"Pit_raffle_tickets_2019_summer"`
						PitKotlGold2019Spring      int `json:"Pit_kotl_gold_2019_spring"`
						PitRaffleTickets2019Spring int `json:"Pit_raffle_tickets_2019_spring"`
					} `json:"leaderboard_stats"`
					SelectedPerk3       string  `json:"selected_perk_3"`
					SelectedPerk2       string  `json:"selected_perk_2"`
					SelectedPerk1       string  `json:"selected_perk_1"`
					SelectedPerk0       string  `json:"selected_perk_0"`
					LastContract        int64   `json:"last_contract"`
					CashDuringPrestige3 float64 `json:"cash_during_prestige_3"`
					GoldTransactions    []struct {
						Amount    int   `json:"amount"`
						Timestamp int64 `json:"timestamp"`
					} `json:"gold_transactions"`
					CashDuringPrestige2 float64 `json:"cash_during_prestige_2"`
					CashDuringPrestige5 float64 `json:"cash_during_prestige_5"`
					CashDuringPrestige4 float64 `json:"cash_during_prestige_4"`
					CashDuringPrestige7 float64 `json:"cash_during_prestige_7"`
					CashDuringPrestige6 float64 `json:"cash_during_prestige_6"`
					InvContents         struct {
						Type int   `json:"type"`
						Data []int `json:"data"`
					} `json:"inv_contents"`
					Unlocks []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"unlocks"`
					CashDuringPrestige1 float64 `json:"cash_during_prestige_1"`
					CashDuringPrestige0 float64 `json:"cash_during_prestige_0"`
					Renown              int     `json:"renown"`
					MovedAchievements1  bool    `json:"moved_achievements_1"`
					MovedAchievements2  bool    `json:"moved_achievements_2"`
					Prestiges           []struct {
						Index        int   `json:"index"`
						XpOnPrestige int   `json:"xp_on_prestige"`
						Timestamp    int64 `json:"timestamp"`
					} `json:"prestiges"`
					SpireStashInv struct {
						Type int   `json:"type"`
						Data []int `json:"data"`
					} `json:"spire_stash_inv"`
					ZeroPointThreeGoldTransfer bool `json:"zero_point_three_gold_transfer"`
					RenownUnlocks              []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"renown_unlocks"`
					SpireStashArmor struct {
						Type int   `json:"type"`
						Data []int `json:"data"`
					} `json:"spire_stash_armor"`
					LastMidfightDisconnect int64 `json:"last_midfight_disconnect"`
					InvArmor               struct {
						Type int   `json:"type"`
						Data []int `json:"data"`
					} `json:"inv_armor"`
					LoginMessages   []interface{} `json:"login_messages"`
					HotbarFavorites []int         `json:"hotbar_favorites"`
					Xp              int           `json:"xp"`
					EndedContracts  []struct {
						Difficulty   string `json:"difficulty"`
						GoldReward   int    `json:"gold_reward"`
						Requirements struct {
							Streak int `json:"streak"`
						} `json:"requirements"`
						Progress struct {
							Streak int `json:"streak"`
						} `json:"progress"`
						CompletionDate int64  `json:"completion_date"`
						RemainingTicks int    `json:"remaining_ticks"`
						Key            string `json:"key"`
					} `json:"ended_contracts"`
					Bounties []interface{} `json:"bounties"`
				} `json:"profile"`
				PitStatsPtl struct {
					Deaths                int `json:"deaths"`
					EnderchestOpened      int `json:"enderchest_opened"`
					MeleeDamageDealt      int `json:"melee_damage_dealt"`
					CashEarned            int `json:"cash_earned"`
					Joins                 int `json:"joins"`
					PlaytimeMinutes       int `json:"playtime_minutes"`
					BowDamageReceived     int `json:"bow_damage_received"`
					Kills                 int `json:"kills"`
					DamageReceived        int `json:"damage_received"`
					JumpedIntoPit         int `json:"jumped_into_pit"`
					MeleeDamageReceived   int `json:"melee_damage_received"`
					LeftClicks            int `json:"left_clicks"`
					DamageDealt           int `json:"damage_dealt"`
					Assists               int `json:"assists"`
					BlocksPlaced          int `json:"blocks_placed"`
					ContractsCompleted    int `json:"contracts_completed"`
					LaunchedByLaunchers   int `json:"launched_by_launchers"`
					ArrowsFired           int `json:"arrows_fired"`
					BowDamageDealt        int `json:"bow_damage_dealt"`
					ContractsStarted      int `json:"contracts_started"`
					ChatMessages          int `json:"chat_messages"`
					ArrowHits             int `json:"arrow_hits"`
					DiamondItemsPurchased int `json:"diamond_items_purchased"`
					MaxStreak             int `json:"max_streak"`
				} `json:"pit_stats_ptl"`
			} `json:"Pit"`
			Housing struct {
				LayoutItems struct {
				} `json:"layout_items"`
			} `json:"Housing"`
			SkyBlock struct {
				Profiles struct {
					B876Ec32E396476Ba1158438D83C67D4 struct {
						ProfileID string `json:"profile_id"`
						CuteName  string `json:"cute_name"`
					} `json:"b876ec32e396476ba1158438d83c67d4"`
					Eight92Dde8Fda9B494Aa4139040758E3Ae7 struct {
						ProfileID string `json:"profile_id"`
						CuteName  string `json:"cute_name"`
					} `json:"892dde8fda9b494aa4139040758e3ae7"`
					Six4722047F9B34E69B67B76A62351Eb05 struct {
						ProfileID string `json:"profile_id"`
						CuteName  string `json:"cute_name"`
					} `json:"64722047f9b34e69b67b76a62351eb05"`
					Three564B910B00C40698B6Cf652A344A41E struct {
						ProfileID string `json:"profile_id"`
						CuteName  string `json:"cute_name"`
					} `json:"3564b910b00c40698b6cf652a344a41e"`
					NineD3C26E0078D47A9Be7F0292B710Bed3 struct {
						ProfileID string `json:"profile_id"`
						CuteName  string `json:"cute_name"`
					} `json:"9d3c26e0078d47a9be7f0292b710bed3"`
				} `json:"profiles"`
			} `json:"SkyBlock"`
		} `json:"stats"`
		TestPass         bool   `json:"testPass"`
		ThanksReceived   int    `json:"thanksReceived"`
		ThanksSent       int    `json:"thanksSent"`
		TimePlaying      int    `json:"timePlaying"`
		TournamentTokens int    `json:"tournamentTokens"`
		UUID             string `json:"uuid"`
		VanityMeta       struct {
			Packages []string `json:"packages"`
		} `json:"vanityMeta"`
		Wardrobe string `json:"wardrobe"`
		Eugene   struct {
			DailyTwoKExp int64 `json:"dailyTwoKExp"`
		} `json:"eugene"`
		Voting struct {
			SecondaryMinestatus int   `json:"secondary_minestatus"`
			TotalMinestatus     int   `json:"total_minestatus"`
			Total               int   `json:"total"`
			LastMinestatus      int64 `json:"last_minestatus"`
			LastVote            int64 `json:"last_vote"`
			TotalMcsorg         int   `json:"total_mcsorg"`
			SecondaryMcsorg     int   `json:"secondary_mcsorg"`
			LastMcsorg          int64 `json:"last_mcsorg"`
			SecondaryMcsl       int   `json:"secondary_mcsl"`
			TotalMcsl           int   `json:"total_mcsl"`
			LastMcsl            int64 `json:"last_mcsl"`
			TotalPmc            int   `json:"total_pmc"`
			SecondaryPmc        int   `json:"secondary_pmc"`
			LastPmc             int64 `json:"last_pmc"`
			SecondaryTopg       int   `json:"secondary_topg"`
			TotalTopg           int   `json:"total_topg"`
			LastTopg            int64 `json:"last_topg"`
			VotesToday          int   `json:"votesToday"`
			TotalMcmp           int   `json:"total_mcmp"`
			SecondaryMcmp       int   `json:"secondary_mcmp"`
			LastMcmp            int64 `json:"last_mcmp"`
			SecondaryMcipl      int   `json:"secondary_mcipl"`
			TotalMcipl          int   `json:"total_mcipl"`
			LastMcipl           int64 `json:"last_mcipl"`
			SecondaryPact       int   `json:"secondary_pact"`
			TotalPact           int   `json:"total_pact"`
			LastPact            int64 `json:"last_pact"`
			TotalMcf            int   `json:"total_mcf"`
			SecondaryMcf        int   `json:"secondary_mcf"`
			LastMcf             int64 `json:"last_mcf"`
		} `json:"voting"`
		McVersionRp             string `json:"mcVersionRp"`
		MostRecentlyThankedUUID string `json:"mostRecentlyThankedUuid"`
		PetConsumables          struct {
			BakedPotato  int `json:"BAKED_POTATO"`
			Cookie       int `json:"COOKIE"`
			Feather      int `json:"FEATHER"`
			HayBlock     int `json:"HAY_BLOCK"`
			SlimeBall    int `json:"SLIME_BALL"`
			CookedBeef   int `json:"COOKED_BEEF"`
			RedRose      int `json:"RED_ROSE"`
			WaterBucket  int `json:"WATER_BUCKET"`
			Melon        int `json:"MELON"`
			Stick        int `json:"STICK"`
			WoodSword    int `json:"WOOD_SWORD"`
			MilkBucket   int `json:"MILK_BUCKET"`
			GoldRecord   int `json:"GOLD_RECORD"`
			Leash        int `json:"LEASH"`
			LavaBucket   int `json:"LAVA_BUCKET"`
			Bone         int `json:"BONE"`
			MagmaCream   int `json:"MAGMA_CREAM"`
			Wheat        int `json:"WHEAT"`
			MushroomSoup int `json:"MUSHROOM_SOUP"`
			Bread        int `json:"BREAD"`
			PumpkinPie   int `json:"PUMPKIN_PIE"`
			Apple        int `json:"APPLE"`
			CarrotItem   int `json:"CARROT_ITEM"`
			RawFish      int `json:"RAW_FISH"`
			Pork         int `json:"PORK"`
			Cake         int `json:"CAKE"`
			RottenFlesh  int `json:"ROTTEN_FLESH"`
		} `json:"petConsumables"`
		HousingMeta struct {
			TutorialStep       string   `json:"tutorialStep"`
			Packages           []string `json:"packages"`
			AllowedBlocks      []string `json:"allowedBlocks"`
			VisibilityDisabled bool     `json:"visibilityDisabled"`
			PlayerSettings     struct {
				Border     string `json:"BORDER"`
				YtRepulsor string `json:"YT_REPULSOR"`
				Tips       string `json:"TIPS"`
				Visibility string `json:"VISIBILITY"`
			} `json:"playerSettings"`
			FirstHouseJoinMs   int64    `json:"firstHouseJoinMs"`
			GivenCookies105000 []string `json:"given_cookies_105000"`
			GivenCookies105006 []string `json:"given_cookies_105006"`
			GivenCookies105016 []string `json:"given_cookies_105016"`
			GivenCookies105025 []string `json:"given_cookies_105025"`
			SelectedChannelsV3 []string `json:"selectedChannels_v3"`
			GivenCookies105036 []string `json:"given_cookies_105036"`
			GivenCookies105048 []string `json:"given_cookies_105048"`
			Playlist           string   `json:"playlist"`
		} `json:"housingMeta"`
		PetStats struct {
			Pig struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int    `json:"experience"`
				Name       string `json:"name"`
			} `json:"PIG"`
			IronGolem struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int    `json:"experience"`
				Name       string `json:"name"`
			} `json:"IRON_GOLEM"`
			Snowman struct {
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int    `json:"experience"`
				Name       string `json:"name"`
			} `json:"SNOWMAN"`
			HorseBrown struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"HORSE_BROWN"`
			HorseBrownBaby struct {
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"HORSE_BROWN_BABY"`
			SlimeBig struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Experience int `json:"experience"`
				Exercise   struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
			} `json:"SLIME_BIG"`
			Wolf struct {
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"WOLF"`
			Zombie struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"ZOMBIE"`
			Silverfish struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SILVERFISH"`
			SheepWhite struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_WHITE"`
			SheepGray struct {
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"SHEEP_GRAY"`
			SheepBrown struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_BROWN"`
			SheepSilver struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_SILVER"`
			SheepWhiteBaby struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_WHITE_BABY"`
			SheepGrayBaby struct {
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Experience int `json:"experience"`
			} `json:"SHEEP_GRAY_BABY"`
			SheepBrownBaby struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_BROWN_BABY"`
			SheepSilverBaby struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_SILVER_BABY"`
			SheepOrange struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_ORANGE"`
			SheepMagenta struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_MAGENTA"`
			SheepLightBlue struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_LIGHT_BLUE"`
			SheepYellow struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_YELLOW"`
			SheepLime struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_LIME"`
			SheepCyan struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_CYAN"`
			SheepPurple struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_PURPLE"`
			SheepBlue struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_BLUE"`
			SheepGreen struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_GREEN"`
			SheepRed struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_RED"`
			HorseGrayBaby struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"HORSE_GRAY_BABY"`
			HorseSkeleton struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"HORSE_SKELETON"`
			SlimeTiny struct {
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"SLIME_TINY"`
			HorseDarkBrown struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"HORSE_DARK_BROWN"`
			HorseGrey struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"HORSE_GREY"`
			HorseChestnut struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"HORSE_CHESTNUT"`
			HorseCreamy struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"HORSE_CREAMY"`
			HorseCreamyBaby struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"HORSE_CREAMY_BABY"`
			HorseChestnutBaby struct {
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"HORSE_CHESTNUT_BABY"`
			WildOcelot struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"WILD_OCELOT"`
			WildOcelotBaby struct {
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"WILD_OCELOT_BABY"`
			HorseUndead struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"HORSE_UNDEAD"`
			MooshroomBaby struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"MOOSHROOM_BABY"`
			SlimeSmall struct {
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SLIME_SMALL"`
			SheepPinkBaby struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_PINK_BABY"`
			SheepBlackBaby struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_BLACK_BABY"`
			HorseBlack struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"HORSE_BLACK"`
			HorseWhite struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"HORSE_WHITE"`
			SheepPink struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_PINK"`
			SheepOrangeBaby struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_ORANGE_BABY"`
			SheepMagentaBaby struct {
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"SHEEP_MAGENTA_BABY"`
			SheepBlack struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"SHEEP_BLACK"`
			SheepLightBlueBaby struct {
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SHEEP_LIGHT_BLUE_BABY"`
			PigZombie struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int    `json:"experience"`
				Name       string `json:"name"`
			} `json:"PIG_ZOMBIE"`
			CreeperPowered struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int    `json:"experience"`
				Name       string `json:"name"`
			} `json:"CREEPER_POWERED"`
			Witch struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"WITCH"`
			MagmaCubeBig struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"MAGMA_CUBE_BIG"`
			Herobrine struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int    `json:"experience"`
				Name       string `json:"name"`
			} `json:"HEROBRINE"`
			Creeper struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int    `json:"experience"`
				Name       string `json:"name"`
			} `json:"CREEPER"`
			Spider struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SPIDER"`
			Enderman struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int    `json:"experience"`
				Name       string `json:"name"`
			} `json:"ENDERMAN"`
			PigZombieBaby struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int    `json:"experience"`
				Name       string `json:"name"`
			} `json:"PIG_ZOMBIE_BABY"`
			PigBaby struct {
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"PIG_BABY"`
			Gorrila struct {
				Name string `json:"name"`
			} `json:"GORRILA"`
			Gifterino struct {
			} `json:"GIFTERINO"`
			Bat struct {
				Name   string `json:"name"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"BAT"`
			Elephant struct {
				Name string `json:"name"`
			} `json:"ELEPHANT"`
			Penguin struct {
				Name string `json:"name"`
			} `json:"PENGUIN"`
			CatBlack struct {
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Experience int `json:"experience"`
			} `json:"CAT_BLACK"`
			CatRed struct {
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"CAT_RED"`
			CatSiamese struct {
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Experience int `json:"experience"`
			} `json:"CAT_SIAMESE"`
			CatBlackBaby struct {
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Experience int `json:"experience"`
			} `json:"CAT_BLACK_BABY"`
			Chicken struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"CHICKEN"`
			ChickenBaby struct {
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"CHICKEN_BABY"`
			CatRedBaby struct {
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"CAT_RED_BABY"`
			CatSiameseBaby struct {
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"CAT_SIAMESE_BABY"`
			WolfBaby struct {
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"WOLF_BABY"`
			ZombieBaby struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"ZOMBIE_BABY"`
			CaveSpider struct {
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"CAVE_SPIDER"`
			BlackRabbit struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"BLACK_RABBIT"`
			BrownRabbit struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"BROWN_RABBIT"`
			GoldRabbit struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"GOLD_RABBIT"`
			SaltPepperRabbit struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"SALT_PEPPER_RABBIT"`
			WhiteRabbit struct {
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"WHITE_RABBIT"`
			VillagerBlacksmith struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"VILLAGER_BLACKSMITH"`
			VillagerBlacksmithBaby struct {
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"VILLAGER_BLACKSMITH_BABY"`
			VillagerButcher struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"VILLAGER_BUTCHER"`
			RedHelper struct {
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience int `json:"experience"`
			} `json:"RED_HELPER"`
			GreenHelper struct {
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"GREEN_HELPER"`
			Blaze struct {
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Experience int    `json:"experience"`
				Name       string `json:"name"`
			} `json:"BLAZE"`
			LittleYou struct {
				Name string `json:"name"`
			} `json:"LITTLE_YOU"`
		} `json:"petStats"`
		PetJourneyTimestamp int64  `json:"petJourneyTimestamp"`
		Transformation      string `json:"transformation"`
		SocialMedia         struct {
			Youtube string `json:"YOUTUBE"`
			Twitter string `json:"TWITTER"`
			Links   struct {
				Youtube string `json:"YOUTUBE"`
				Twitter string `json:"TWITTER"`
			} `json:"links"`
		} `json:"socialMedia"`
		DisableTipMessages bool `json:"disableTipMessages"`
		RewardConsumed     bool `json:"rewardConsumed"`
		GiftingMeta        struct {
			RealBundlesReceivedInc int `json:"realBundlesReceivedInc"`
			BundlesReceived        int `json:"bundlesReceived"`
			RealBundlesReceived    int `json:"realBundlesReceived"`
			GiftsGiven             int `json:"giftsGiven"`
			BundlesGiven           int `json:"bundlesGiven"`
			RealBundlesGiven       int `json:"realBundlesGiven"`
		} `json:"giftingMeta"`
		LevelingReward0         bool  `json:"levelingReward_0"`
		LevelingReward1         bool  `json:"levelingReward_1"`
		LevelingReward2         bool  `json:"levelingReward_2"`
		LevelingReward3         bool  `json:"levelingReward_3"`
		LevelingReward4         bool  `json:"levelingReward_4"`
		LevelingReward5         bool  `json:"levelingReward_5"`
		LevelingReward6         bool  `json:"levelingReward_6"`
		LevelingReward7         bool  `json:"levelingReward_7"`
		LevelingReward8         bool  `json:"levelingReward_8"`
		LevelingReward9         bool  `json:"levelingReward_9"`
		LevelingReward10        bool  `json:"levelingReward_10"`
		LevelingReward11        bool  `json:"levelingReward_11"`
		LevelingReward12        bool  `json:"levelingReward_12"`
		LevelingReward13        bool  `json:"levelingReward_13"`
		LevelingReward14        bool  `json:"levelingReward_14"`
		LevelingReward15        bool  `json:"levelingReward_15"`
		LevelingReward16        bool  `json:"levelingReward_16"`
		LevelingReward17        bool  `json:"levelingReward_17"`
		LevelingReward18        bool  `json:"levelingReward_18"`
		LevelingReward19        bool  `json:"levelingReward_19"`
		LevelingReward20        bool  `json:"levelingReward_20"`
		LevelingReward21        bool  `json:"levelingReward_21"`
		LevelingReward22        bool  `json:"levelingReward_22"`
		LevelingReward23        bool  `json:"levelingReward_23"`
		LevelingReward24        bool  `json:"levelingReward_24"`
		LevelingReward25        bool  `json:"levelingReward_25"`
		LevelingReward26        bool  `json:"levelingReward_26"`
		LevelingReward27        bool  `json:"levelingReward_27"`
		LevelingReward28        bool  `json:"levelingReward_28"`
		LevelingReward29        bool  `json:"levelingReward_29"`
		LevelingReward30        bool  `json:"levelingReward_30"`
		LevelingReward31        bool  `json:"levelingReward_31"`
		LevelingReward32        bool  `json:"levelingReward_32"`
		LevelingReward33        bool  `json:"levelingReward_33"`
		LevelingReward34        bool  `json:"levelingReward_34"`
		LevelingReward35        bool  `json:"levelingReward_35"`
		LevelingReward36        bool  `json:"levelingReward_36"`
		LevelingReward37        bool  `json:"levelingReward_37"`
		LevelingReward38        bool  `json:"levelingReward_38"`
		LevelingReward39        bool  `json:"levelingReward_39"`
		LevelingReward40        bool  `json:"levelingReward_40"`
		LevelingReward41        bool  `json:"levelingReward_41"`
		LevelingReward42        bool  `json:"levelingReward_42"`
		LevelingReward43        bool  `json:"levelingReward_43"`
		LevelingReward44        bool  `json:"levelingReward_44"`
		LevelingReward45        bool  `json:"levelingReward_45"`
		LevelingReward46        bool  `json:"levelingReward_46"`
		LevelingReward47        bool  `json:"levelingReward_47"`
		LevelingReward48        bool  `json:"levelingReward_48"`
		LevelingReward49        bool  `json:"levelingReward_49"`
		LevelingReward50        bool  `json:"levelingReward_50"`
		LevelingReward51        bool  `json:"levelingReward_51"`
		LevelingReward52        bool  `json:"levelingReward_52"`
		LevelingReward53        bool  `json:"levelingReward_53"`
		LevelingReward54        bool  `json:"levelingReward_54"`
		LevelingReward55        bool  `json:"levelingReward_55"`
		LevelingReward56        bool  `json:"levelingReward_56"`
		LevelingReward57        bool  `json:"levelingReward_57"`
		LevelingReward58        bool  `json:"levelingReward_58"`
		LevelingReward59        bool  `json:"levelingReward_59"`
		LevelingReward60        bool  `json:"levelingReward_60"`
		LevelingReward61        bool  `json:"levelingReward_61"`
		LevelingReward62        bool  `json:"levelingReward_62"`
		LevelingReward63        bool  `json:"levelingReward_63"`
		LevelingReward64        bool  `json:"levelingReward_64"`
		LevelingReward65        bool  `json:"levelingReward_65"`
		LevelingReward66        bool  `json:"levelingReward_66"`
		LevelingReward67        bool  `json:"levelingReward_67"`
		LevelingReward68        bool  `json:"levelingReward_68"`
		LevelingReward69        bool  `json:"levelingReward_69"`
		LevelingReward70        bool  `json:"levelingReward_70"`
		LevelingReward71        bool  `json:"levelingReward_71"`
		LevelingReward72        bool  `json:"levelingReward_72"`
		LevelingReward73        bool  `json:"levelingReward_73"`
		LevelingReward74        bool  `json:"levelingReward_74"`
		LevelingReward75        bool  `json:"levelingReward_75"`
		LevelingReward76        bool  `json:"levelingReward_76"`
		LevelingReward77        bool  `json:"levelingReward_77"`
		LevelingReward78        bool  `json:"levelingReward_78"`
		LevelingReward79        bool  `json:"levelingReward_79"`
		LevelingReward80        bool  `json:"levelingReward_80"`
		LevelingReward81        bool  `json:"levelingReward_81"`
		LevelingReward82        bool  `json:"levelingReward_82"`
		LevelingReward83        bool  `json:"levelingReward_83"`
		LevelingReward84        bool  `json:"levelingReward_84"`
		LevelingReward85        bool  `json:"levelingReward_85"`
		LevelingReward86        bool  `json:"levelingReward_86"`
		LevelingReward87        bool  `json:"levelingReward_87"`
		LevelingReward88        bool  `json:"levelingReward_88"`
		LevelingReward89        bool  `json:"levelingReward_89"`
		LevelingReward90        bool  `json:"levelingReward_90"`
		LevelingReward91        bool  `json:"levelingReward_91"`
		LevelingReward92        bool  `json:"levelingReward_92"`
		LevelingReward93        bool  `json:"levelingReward_93"`
		LevelingReward94        bool  `json:"levelingReward_94"`
		LevelingReward95        bool  `json:"levelingReward_95"`
		LevelingReward96        bool  `json:"levelingReward_96"`
		LevelingReward97        bool  `json:"levelingReward_97"`
		LevelingReward98        bool  `json:"levelingReward_98"`
		LevelingReward99        bool  `json:"levelingReward_99"`
		LevelingReward109       bool  `json:"levelingReward_109"`
		LevelingReward119       bool  `json:"levelingReward_119"`
		LevelingReward129       bool  `json:"levelingReward_129"`
		LevelingReward139       bool  `json:"levelingReward_139"`
		LevelingReward149       bool  `json:"levelingReward_149"`
		VanityConvertedBoxToday int   `json:"vanityConvertedBoxToday"`
		VanityFirstConvertedBox int64 `json:"vanityFirstConvertedBox"`
		SpecialtyCooldowns      struct {
			Mvp0     bool `json:"MVP0"`
			VipPlus0 bool `json:"VIP_PLUS0"`
			Vip0     bool `json:"VIP0"`
			MvpPlus0 bool `json:"MVP_PLUS0"`
			Vip1     bool `json:"VIP1"`
			VipPlus1 bool `json:"VIP_PLUS1"`
			Mvp1     bool `json:"MVP1"`
			MvpPlus1 bool `json:"MVP_PLUS1"`
			Vip2     bool `json:"VIP2"`
			VipPlus2 bool `json:"VIP_PLUS2"`
			Mvp2     bool `json:"MVP2"`
			MvpPlus2 bool `json:"MVP_PLUS2"`
			Vip3     bool `json:"VIP3"`
			VipPlus3 bool `json:"VIP_PLUS3"`
			Mvp3     bool `json:"MVP3"`
			MvpPlus3 bool `json:"MVP_PLUS3"`
			Mvp4     bool `json:"MVP4"`
			Vip4     bool `json:"VIP4"`
			VipPlus4 bool `json:"VIP_PLUS4"`
			MvpPlus4 bool `json:"MVP_PLUS4"`
			Vip5     bool `json:"VIP5"`
			VipPlus5 bool `json:"VIP_PLUS5"`
			Mvp5     bool `json:"MVP5"`
			MvpPlus5 bool `json:"MVP_PLUS5"`
			Vip6     bool `json:"VIP6"`
			VipPlus6 bool `json:"VIP_PLUS6"`
			Mvp6     bool `json:"MVP6"`
			MvpPlus6 bool `json:"MVP_PLUS6"`
		} `json:"specialtyCooldowns"`
		QuestSettings struct {
			AutoActivate bool `json:"autoActivate"`
		} `json:"questSettings"`
		LastAdsenseGenerateTime int64  `json:"lastAdsenseGenerateTime"`
		LastClaimedReward       int64  `json:"lastClaimedReward"`
		TotalRewards            int    `json:"totalRewards"`
		TotalDailyRewards       int    `json:"totalDailyRewards"`
		RewardStreak            int    `json:"rewardStreak"`
		RewardScore             int    `json:"rewardScore"`
		RewardHighScore         int    `json:"rewardHighScore"`
		RankPlusColor           string `json:"rankPlusColor"`
		AdsenseTokens           int    `json:"adsense_tokens"`
		LevelingReward159       bool   `json:"levelingReward_159"`
		FlashingSalePopup       int64  `json:"flashingSalePopup"`
		FlashingSalePoppedUp    int    `json:"flashingSalePoppedUp"`
		FlashingSaleOpens       int    `json:"flashingSaleOpens"`
		FlashingSaleClicks      int    `json:"flashingSaleClicks"`
		Halloween2016Cooldowns  struct {
			Vip0     bool `json:"VIP0"`
			VipPlus0 bool `json:"VIP_PLUS0"`
			MvpPlus0 bool `json:"MVP_PLUS0"`
			Mvp0     bool `json:"MVP0"`
			Vip1     bool `json:"VIP1"`
			VipPlus1 bool `json:"VIP_PLUS1"`
			Mvp1     bool `json:"MVP1"`
			MvpPlus1 bool `json:"MVP_PLUS1"`
			VipPlus2 bool `json:"VIP_PLUS2"`
			Vip2     bool `json:"VIP2"`
			Mvp2     bool `json:"MVP2"`
			MvpPlus2 bool `json:"MVP_PLUS2"`
		} `json:"halloween2016Cooldowns"`
		LevelingReward169    bool `json:"levelingReward_169"`
		SantaQuestStarted    bool `json:"SANTA_QUEST_STARTED"`
		Holiday2016Cooldowns struct {
			Vip0     bool `json:"VIP0"`
			Vip1     bool `json:"VIP1"`
			Vip2     bool `json:"VIP2"`
			Vip3     bool `json:"VIP3"`
			Vip4     bool `json:"VIP4"`
			Mvp0     bool `json:"MVP0"`
			Mvp1     bool `json:"MVP1"`
			Mvp2     bool `json:"MVP2"`
			Mvp3     bool `json:"MVP3"`
			VipPlus0 bool `json:"VIP_PLUS0"`
			VipPlus1 bool `json:"VIP_PLUS1"`
			VipPlus2 bool `json:"VIP_PLUS2"`
			VipPlus3 bool `json:"VIP_PLUS3"`
			MvpPlus0 bool `json:"MVP_PLUS0"`
			MvpPlus1 bool `json:"MVP_PLUS1"`
			MvpPlus2 bool `json:"MVP_PLUS2"`
			MvpPlus3 bool `json:"MVP_PLUS3"`
			VipPlus4 bool `json:"VIP_PLUS4"`
			MvpPlus4 bool `json:"MVP_PLUS4"`
			Mvp4     bool `json:"MVP4"`
			Vip5     bool `json:"VIP5"`
			VipPlus5 bool `json:"VIP_PLUS5"`
			Mvp5     bool `json:"MVP5"`
			MvpPlus5 bool `json:"MVP_PLUS5"`
		} `json:"holiday2016Cooldowns"`
		LevelingReward179 bool `json:"levelingReward_179"`
		CompassStats      struct {
			Compass struct {
				Battleground int `json:"battleground"`
				Arcade       int `json:"arcade"`
				Prototype    int `json:"prototype"`
				Skywars      int `json:"skywars"`
				Tntgames     int `json:"tntgames"`
			} `json:"compass"`
		} `json:"compassStats"`
		FortuneBuff           int      `json:"fortuneBuff"`
		NetworkUpdateBook     string   `json:"network_update_book"`
		LastLogout            int64    `json:"lastLogout"`
		AchievementTracking   []string `json:"achievementTracking"`
		AchievementRewardsNew struct {
			ForPoints200   int64 `json:"for_points_200"`
			ForPoints300   int64 `json:"for_points_300"`
			ForPoints400   int64 `json:"for_points_400"`
			ForPoints500   int64 `json:"for_points_500"`
			ForPoints600   int64 `json:"for_points_600"`
			ForPoints700   int64 `json:"for_points_700"`
			ForPoints800   int64 `json:"for_points_800"`
			ForPoints900   int64 `json:"for_points_900"`
			ForPoints1000  int64 `json:"for_points_1000"`
			ForPoints1100  int64 `json:"for_points_1100"`
			ForPoints1200  int64 `json:"for_points_1200"`
			ForPoints1300  int64 `json:"for_points_1300"`
			ForPoints1400  int64 `json:"for_points_1400"`
			ForPoints1500  int64 `json:"for_points_1500"`
			ForPoints1600  int64 `json:"for_points_1600"`
			ForPoints1700  int64 `json:"for_points_1700"`
			ForPoints1800  int64 `json:"for_points_1800"`
			ForPoints1900  int64 `json:"for_points_1900"`
			ForPoints2000  int64 `json:"for_points_2000"`
			ForPoints2100  int64 `json:"for_points_2100"`
			ForPoints2200  int64 `json:"for_points_2200"`
			ForPoints2300  int64 `json:"for_points_2300"`
			ForPoints2400  int64 `json:"for_points_2400"`
			ForPoints2500  int64 `json:"for_points_2500"`
			ForPoints2600  int64 `json:"for_points_2600"`
			ForPoints2700  int64 `json:"for_points_2700"`
			ForPoints2800  int64 `json:"for_points_2800"`
			ForPoints2900  int64 `json:"for_points_2900"`
			ForPoints3000  int64 `json:"for_points_3000"`
			ForPoints3100  int64 `json:"for_points_3100"`
			ForPoints3200  int64 `json:"for_points_3200"`
			ForPoints3300  int64 `json:"for_points_3300"`
			ForPoints3400  int64 `json:"for_points_3400"`
			ForPoints3500  int64 `json:"for_points_3500"`
			ForPoints3600  int64 `json:"for_points_3600"`
			ForPoints3700  int64 `json:"for_points_3700"`
			ForPoints3800  int64 `json:"for_points_3800"`
			ForPoints3900  int64 `json:"for_points_3900"`
			ForPoints4000  int64 `json:"for_points_4000"`
			ForPoints4100  int64 `json:"for_points_4100"`
			ForPoints4200  int64 `json:"for_points_4200"`
			ForPoints4300  int64 `json:"for_points_4300"`
			ForPoints4400  int64 `json:"for_points_4400"`
			ForPoints4500  int64 `json:"for_points_4500"`
			ForPoints4600  int64 `json:"for_points_4600"`
			ForPoints4700  int64 `json:"for_points_4700"`
			ForPoints4800  int64 `json:"for_points_4800"`
			ForPoints4900  int64 `json:"for_points_4900"`
			ForPoints5000  int64 `json:"for_points_5000"`
			ForPoints5100  int64 `json:"for_points_5100"`
			ForPoints5200  int64 `json:"for_points_5200"`
			ForPoints5300  int64 `json:"for_points_5300"`
			ForPoints5400  int64 `json:"for_points_5400"`
			ForPoints5500  int64 `json:"for_points_5500"`
			ForPoints5600  int64 `json:"for_points_5600"`
			ForPoints5700  int64 `json:"for_points_5700"`
			ForPoints5800  int64 `json:"for_points_5800"`
			ForPoints5900  int64 `json:"for_points_5900"`
			ForPoints6000  int64 `json:"for_points_6000"`
			ForPoints6100  int64 `json:"for_points_6100"`
			ForPoints6200  int64 `json:"for_points_6200"`
			ForPoints6300  int64 `json:"for_points_6300"`
			ForPoints6400  int64 `json:"for_points_6400"`
			ForPoints6500  int64 `json:"for_points_6500"`
			ForPoints6600  int64 `json:"for_points_6600"`
			ForPoints6700  int64 `json:"for_points_6700"`
			ForPoints6800  int64 `json:"for_points_6800"`
			ForPoints6900  int64 `json:"for_points_6900"`
			ForPoints7000  int64 `json:"for_points_7000"`
			ForPoints7100  int64 `json:"for_points_7100"`
			ForPoints7200  int64 `json:"for_points_7200"`
			ForPoints7300  int64 `json:"for_points_7300"`
			ForPoints7400  int64 `json:"for_points_7400"`
			ForPoints7500  int64 `json:"for_points_7500"`
			ForPoints7600  int64 `json:"for_points_7600"`
			ForPoints7700  int64 `json:"for_points_7700"`
			ForPoints7800  int64 `json:"for_points_7800"`
			ForPoints7900  int64 `json:"for_points_7900"`
			ForPoints8000  int64 `json:"for_points_8000"`
			ForPoints8100  int64 `json:"for_points_8100"`
			ForPoints8200  int64 `json:"for_points_8200"`
			ForPoints8300  int64 `json:"for_points_8300"`
			ForPoints8400  int64 `json:"for_points_8400"`
			ForPoints8500  int64 `json:"for_points_8500"`
			ForPoints8600  int64 `json:"for_points_8600"`
			ForPoints8700  int64 `json:"for_points_8700"`
			ForPoints8800  int64 `json:"for_points_8800"`
			ForPoints9000  int64 `json:"for_points_9000"`
			ForPoints8900  int64 `json:"for_points_8900"`
			ForPoints9100  int64 `json:"for_points_9100"`
			ForPoints9200  int64 `json:"for_points_9200"`
			ForPoints9300  int64 `json:"for_points_9300"`
			ForPoints9400  int64 `json:"for_points_9400"`
			ForPoints9500  int64 `json:"for_points_9500"`
			ForPoints9600  int64 `json:"for_points_9600"`
			ForPoints9700  int64 `json:"for_points_9700"`
			ForPoints9800  int64 `json:"for_points_9800"`
			ForPoints9900  int64 `json:"for_points_9900"`
			ForPoints10000 int64 `json:"for_points_10000"`
			ForPoints10100 int64 `json:"for_points_10100"`
			ForPoints10200 int64 `json:"for_points_10200"`
			ForPoints10300 int64 `json:"for_points_10300"`
			ForPoints10400 int64 `json:"for_points_10400"`
			ForPoints10500 int64 `json:"for_points_10500"`
			ForPoints10600 int64 `json:"for_points_10600"`
			ForPoints10700 int64 `json:"for_points_10700"`
			ForPoints10800 int64 `json:"for_points_10800"`
			ForPoints10900 int64 `json:"for_points_10900"`
			ForPoints11000 int64 `json:"for_points_11000"`
			ForPoints11100 int64 `json:"for_points_11100"`
			ForPoints11200 int64 `json:"for_points_11200"`
			ForPoints11300 int64 `json:"for_points_11300"`
			ForPoints11400 int64 `json:"for_points_11400"`
			ForPoints11500 int64 `json:"for_points_11500"`
			ForPoints11600 int64 `json:"for_points_11600"`
			ForPoints11700 int64 `json:"for_points_11700"`
			ForPoints11800 int64 `json:"for_points_11800"`
			ForPoints11900 int64 `json:"for_points_11900"`
			ForPoints12000 int64 `json:"for_points_12000"`
			ForPoints12100 int64 `json:"for_points_12100"`
			ForPoints12200 int64 `json:"for_points_12200"`
			ForPoints12300 int64 `json:"for_points_12300"`
		} `json:"achievementRewardsNew"`
		AchievementTotem struct {
			CanCustomize     bool     `json:"canCustomize"`
			AllowedMaxHeight int      `json:"allowed_max_height"`
			UnlockedParts    []string `json:"unlockedParts"`
			SelectedParts    struct {
				Slot0 string `json:"slot_0"`
				Slot1 string `json:"slot_1"`
				Slot2 string `json:"slot_2"`
				Slot3 string `json:"slot_3"`
			} `json:"selectedParts"`
			UnlockedColors []string `json:"unlockedColors"`
			SelectedColors struct {
				Slotcolor0 string `json:"slotcolor_0"`
			} `json:"selectedColors"`
		} `json:"achievementTotem"`
		LevelingReward189            bool `json:"levelingReward_189"`
		LevelingReward199            bool `json:"levelingReward_199"`
		LevelingReward209            bool `json:"levelingReward_209"`
		LevelingReward219            bool `json:"levelingReward_219"`
		LevelingReward224            bool `json:"levelingReward_224"`
		CompletedChristmasQuests2017 int  `json:"completed_christmas_quests_2017"`
		Challenges                   struct {
			AllTime struct {
				BEDWARSOffensive                 int `json:"BEDWARS__offensive"`
				BEDWARSSupport                   int `json:"BEDWARS__support"`
				SKYWARSRushChallenge             int `json:"SKYWARS__rush_challenge"`
				SKYWARSEndermanChallenge         int `json:"SKYWARS__enderman_challenge"`
				SKYWARSFeedingTheVoidChallenge   int `json:"SKYWARS__feeding_the_void_challenge"`
				SUPERSMASHLeaderboardChallenge   int `json:"SUPER_SMASH__leaderboard_challenge"`
				SUPERSMASHSmashChallenge         int `json:"SUPER_SMASH__smash_challenge"`
				SUPERSMASHFlawlessChallenge      int `json:"SUPER_SMASH__flawless_challenge"`
				ARCADEPartyGamesChallenge        int `json:"ARCADE__party_games_challenge"`
				BUILDBATTLEGuesserChallenge      int `json:"BUILD_BATTLE__guesser_challenge"`
				BUILDBATTLETop3Challenge         int `json:"BUILD_BATTLE__top_3_challenge"`
				QUAKECRAFTDonTBlinkChallenge     int `json:"QUAKECRAFT__don't_blink_challenge"`
				QUAKECRAFTComboChallenge         int `json:"QUAKECRAFT__combo_challenge"`
				QUAKECRAFTKillingStreakChallenge int `json:"QUAKECRAFT__killing_streak_challenge"`
				PAINTBALLKillStreakChallenge     int `json:"PAINTBALL__kill_streak_challenge"`
				PAINTBALLKillingSpreeChallenge   int `json:"PAINTBALL__killing_spree_challenge"`
				PAINTBALLFinishChallenge         int `json:"PAINTBALL__finish_challenge"`
				GINGERBREADFirstPlaceChallenge   int `json:"GINGERBREAD__first_place_challenge"`
				GINGERBREADCoinChallenge         int `json:"GINGERBREAD__coin_challenge"`
				MURDERMYSTERYMurderSpree         int `json:"MURDER_MYSTERY__murder_spree"`
				MURDERMYSTERYHero                int `json:"MURDER_MYSTERY__hero"`
				MCGOKillingSpreeChallenge        int `json:"MCGO__killing_spree_challenge"`
				PAINTBALLNukeChallenge           int `json:"PAINTBALL__nuke_challenge"`
				TNTGAMESTntWizardChallenge       int `json:"TNTGAMES__tnt_wizard_challenge"`
				GINGERBREADBananaChallenge       int `json:"GINGERBREAD__banana_challenge"`
				DUELSFeedTheVoidChallenge        int `json:"DUELS__feed_the_void_challenge"`
				ARCADEHoleInTheWallChallenge     int `json:"ARCADE__hole_in_the_wall_challenge"`
				ARCADEFootballChallenge          int `json:"ARCADE__football_challenge"`
				GINGERBREADLeaderboardChallenge  int `json:"GINGERBREAD__leaderboard_challenge"`
				ARCADEHypixelSaysChallenge       int `json:"ARCADE__hypixel_says_challenge"`
				DUELSTeamsChallenge              int `json:"DUELS__teams_challenge"`
				ARCADEZombiesChallenge           int `json:"ARCADE__zombies_challenge"`
				ARCADEPixelPaintersChallenge     int `json:"ARCADE__pixel_painters_challenge"`
				SKYCLASHEnderchestChallenge      int `json:"SKYCLASH__enderchest_challenge"`
				SKYCLASHFighterChallenge         int `json:"SKYCLASH__fighter_challenge"`
				SKYCLASHMonsterKillerChallenge   int `json:"SKYCLASH__monster_killer_challenge"`
				BATTLEGROUNDBruteChallenge       int `json:"BATTLEGROUND__brute_challenge"`
				BATTLEGROUNDCarryChallenge       int `json:"BATTLEGROUND__carry_challenge"`
				BATTLEGROUNDSupportChallenge     int `json:"BATTLEGROUND__support_challenge"`
				BATTLEGROUNDCaptureChallenge     int `json:"BATTLEGROUND__capture_challenge"`
				TNTGAMESTntRunChallenge          int `json:"TNTGAMES__tnt_run_challenge"`
				TNTGAMESBowSpleefChallenge       int `json:"TNTGAMES__bow_spleef_challenge"`
				DUELSTargetPracticeChallenge     int `json:"DUELS__target_practice_challenge"`
				TNTGAMESPvpRunChallenge          int `json:"TNTGAMES__pvp_run_challenge"`
				TNTGAMESTntWizardsChallenge      int `json:"TNTGAMES__tnt_wizards_challenge"`
				TNTGAMESTntTagChallenge          int `json:"TNTGAMES__tnt_tag_challenge"`
				VAMPIREZLastStandChallenge       int `json:"VAMPIREZ__last_stand_challenge"`
				VAMPIREZFangChallenge            int `json:"VAMPIREZ__fang_challenge"`
				ARCADEThrowOutChallenge          int `json:"ARCADE__throw_out_challenge"`
				VAMPIREZPurifyingChallenge       int `json:"VAMPIREZ__purifying_challenge"`
			} `json:"all_time"`
		} `json:"challenges"`
		AchievementSync struct {
			QuakeTiered int `json:"quake_tiered"`
		} `json:"achievementSync"`
		LevelingReward249 bool `json:"levelingReward_249"`
		AchievementPoints int  `json:"achievementPoints"`
		Tourney           struct {
			FirstJoinLobby int64 `json:"first_join_lobby"`
			SwCrazySolo0   struct {
				Playtime             int   `json:"playtime"`
				TributesEarned       int   `json:"tributes_earned"`
				FirstWin             int64 `json:"first_win"`
				ClaimedRankingReward int64 `json:"claimed_ranking_reward"`
			} `json:"sw_crazy_solo_0"`
			TotalTributes int `json:"total_tributes"`
			QuakeSolo20   struct {
				GamesPlayed          int   `json:"games_played"`
				Playtime             int   `json:"playtime"`
				TributesEarned       int   `json:"tributes_earned"`
				FirstGame            int64 `json:"first_game"`
				ClaimedRankingReward int64 `json:"claimed_ranking_reward"`
			} `json:"quake_solo2_0"`
			GingerbreadSolo0 struct {
				SeenRPbook           bool  `json:"seenRPbook"`
				GamesPlayed          int   `json:"games_played"`
				Playtime             int   `json:"playtime"`
				TributesEarned       int   `json:"tributes_earned"`
				FirstWin             int64 `json:"first_win"`
				ClaimedRankingReward int64 `json:"claimed_ranking_reward"`
			} `json:"gingerbread_solo_0"`
		} `json:"tourney"`
		CompletedChristmasQuests2018 int    `json:"completed_christmas_quests_2018"`
		VanityFavorites              string `json:"vanityFavorites"`
		CurrentClickEffect           string `json:"currentClickEffect"`
		ParkourCheckpointBests       struct {
			Warlords struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
			} `json:"Warlords"`
			BlitzLobby struct {
				Num0 int `json:"0"`
			} `json:"BlitzLobby"`
			SpeedUHC struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
			} `json:"SpeedUHC"`
			Tnt struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
				Num3 int `json:"3"`
			} `json:"TNT"`
			SkywarsAug2017 struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
				Num3 int `json:"3"`
			} `json:"SkywarsAug2017"`
			Bedwars struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
				Num3 int `json:"3"`
			} `json:"Bedwars"`
		} `json:"ParkourCheckpointBests"`
		BattlePassGlowStatus bool   `json:"battlePassGlowStatus"`
		Prefix               string `json:"prefix"`
		Dmcrates102019       struct {
			Regular bool `json:"REGULAR"`
			Vip     bool `json:"VIP"`
			VipPlus bool `json:"VIP_PLUS"`
			Mvp     bool `json:"MVP"`
			MvpPlus bool `json:"MVP_PLUS"`
		} `json:"dmcrates-10-2019"`
		Monthlycrates struct {
			One2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				MvpPlus bool `json:"MVP_PLUS"`
				Mvp     bool `json:"MVP"`
			} `json:"1-2017"`
			One2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"1-2018"`
			One2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"1-2019"`
			One02016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"10-2016"`
			One02017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"10-2017"`
			One02018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				Mvp     bool `json:"MVP"`
				VipPlus bool `json:"VIP_PLUS"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"10-2018"`
			One02019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"10-2019"`
			One12016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"11-2016"`
			One12017 struct {
				VipPlus bool `json:"VIP_PLUS"`
				Vip     bool `json:"VIP"`
				Regular bool `json:"REGULAR"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"11-2017"`
			One22016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"12-2016"`
			One22017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"12-2017"`
			One22018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"12-2018"`
			Two2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"2-2017"`
			Two2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"2-2018"`
			Two2019 struct {
				Regular bool `json:"REGULAR"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
				Vip     bool `json:"VIP"`
			} `json:"2-2019"`
			Three2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2017"`
			Three2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2018"`
			Three2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2019"`
			Four2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"4-2017"`
			Four2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"4-2018"`
			Four2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"4-2019"`
			Five2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"5-2017"`
			Five2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"5-2018"`
			Five2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"5-2019"`
			Six2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"6-2017"`
			Six2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"6-2019"`
			Seven2016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"7-2016"`
			Seven2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"7-2017"`
			Seven2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"7-2019"`
			Eight2016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"8-2016"`
			Eight2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"8-2017"`
			Eight2018 struct {
				Regular bool `json:"REGULAR"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
				Vip     bool `json:"VIP"`
			} `json:"8-2018"`
			Eight2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"8-2019"`
			Nine2016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"9-2016"`
			Nine2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"9-2017"`
			Nine2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"9-2019"`
			One22019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"12-2019"`
			One2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"1-2020"`
			Two2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"2-2020"`
			Three2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2020"`
			Five2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"5-2020"`
			Six2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"6-2020"`
			Seven2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"7-2020"`
			One12020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"11-2020"`
			Three2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2021"`
			Four2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"4-2021"`
			Five2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"5-2021"`
			Six2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"6-2021"`
			Nine2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"9-2021"`
		} `json:"monthlycrates"`
		Halloween2019Cooldowns struct {
			VipPlus3   bool `json:"VIP_PLUS3"`
			VipPlus2   bool `json:"VIP_PLUS2"`
			VipPlus1   bool `json:"VIP_PLUS1"`
			VipPlus0   bool `json:"VIP_PLUS0"`
			Normal0    bool `json:"NORMAL0"`
			Normal1    bool `json:"NORMAL1"`
			Normal2    bool `json:"NORMAL2"`
			Normal3    bool `json:"NORMAL3"`
			MvpPlus0   bool `json:"MVP_PLUS0"`
			MvpPlus1   bool `json:"MVP_PLUS1"`
			MvpPlus2   bool `json:"MVP_PLUS2"`
			MvpPlus3   bool `json:"MVP_PLUS3"`
			Superstar0 bool `json:"SUPERSTAR0"`
			Superstar1 bool `json:"SUPERSTAR1"`
			Superstar2 bool `json:"SUPERSTAR2"`
			Superstar3 bool `json:"SUPERSTAR3"`
			Vip3       bool `json:"VIP3"`
			Vip2       bool `json:"VIP2"`
			Vip1       bool `json:"VIP1"`
			Mvp3       bool `json:"MVP3"`
			Mvp2       bool `json:"MVP2"`
			Mvp1       bool `json:"MVP1"`
			Mvp0       bool `json:"MVP0"`
			Vip0       bool `json:"VIP0"`
		} `json:"halloween2019Cooldowns"`
		AdventRewards2017 struct {
			Day9  int64 `json:"day9"`
			Day10 int64 `json:"day10"`
			Day11 int64 `json:"day11"`
			Day12 int64 `json:"day12"`
			Day14 int64 `json:"day14"`
			Day15 int64 `json:"day15"`
			Day16 int64 `json:"day16"`
			Day17 int64 `json:"day17"`
			Day18 int64 `json:"day18"`
			Day19 int64 `json:"day19"`
			Day20 int64 `json:"day20"`
			Day22 int64 `json:"day22"`
			Day23 int64 `json:"day23"`
			Day24 int64 `json:"day24"`
			Day25 int64 `json:"day25"`
			Day1  int64 `json:"day1"`
			Day2  int64 `json:"day2"`
			Day3  int64 `json:"day3"`
			Day4  int64 `json:"day4"`
			Day5  int64 `json:"day5"`
			Day6  int64 `json:"day6"`
			Day7  int64 `json:"day7"`
			Day8  int64 `json:"day8"`
			Day13 int64 `json:"day13"`
			Day21 int64 `json:"day21"`
		} `json:"adventRewards2017"`
		AdventRewardsV22018 struct {
			Day1  int64 `json:"day1"`
			Day2  int64 `json:"day2"`
			Day6  int64 `json:"day6"`
			Day8  int64 `json:"day8"`
			Day9  int64 `json:"day9"`
			Day11 int64 `json:"day11"`
			Day16 int64 `json:"day16"`
			Day18 int64 `json:"day18"`
			Day19 int64 `json:"day19"`
			Day20 int64 `json:"day20"`
			Day22 int64 `json:"day22"`
			Day23 int64 `json:"day23"`
			Day24 int64 `json:"day24"`
			Day25 int64 `json:"day25"`
		} `json:"adventRewards_v2_2018"`
		AdventRewards2019 struct {
			Day1  int64 `json:"day1"`
			Day2  int64 `json:"day2"`
			Day4  int64 `json:"day4"`
			Day9  int64 `json:"day9"`
			Day20 int64 `json:"day20"`
			Day22 int64 `json:"day22"`
			Day24 int64 `json:"day24"`
			Day25 int64 `json:"day25"`
		} `json:"adventRewards2019"`
		GiftsGrinch                  int `json:"gifts_grinch"`
		CompletedChristmasQuests2019 int `json:"completed_christmas_quests_2019"`
		Christmas2019Cooldowns       struct {
			Superstar0 bool `json:"SUPERSTAR0"`
			MvpPlus0   bool `json:"MVP_PLUS0"`
			Normal0    bool `json:"NORMAL0"`
			Normal1    bool `json:"NORMAL1"`
			VipPlus0   bool `json:"VIP_PLUS0"`
			VipPlus1   bool `json:"VIP_PLUS1"`
			Mvp0       bool `json:"MVP0"`
			Mvp1       bool `json:"MVP1"`
			Vip0       bool `json:"VIP0"`
			Vip1       bool `json:"VIP1"`
			MvpPlus1   bool `json:"MVP_PLUS1"`
			Superstar1 bool `json:"SUPERSTAR1"`
			Normal2    bool `json:"NORMAL2"`
			MvpPlus3   bool `json:"MVP_PLUS3"`
			MvpPlus2   bool `json:"MVP_PLUS2"`
			Normal3    bool `json:"NORMAL3"`
			VipPlus2   bool `json:"VIP_PLUS2"`
			VipPlus3   bool `json:"VIP_PLUS3"`
			Mvp2       bool `json:"MVP2"`
			Mvp3       bool `json:"MVP3"`
			Vip2       bool `json:"VIP2"`
			Vip3       bool `json:"VIP3"`
			Superstar2 bool `json:"SUPERSTAR2"`
			Superstar3 bool `json:"SUPERSTAR3"`
		} `json:"christmas2019Cooldowns"`
		Xmas2019ARCADE2            bool   `json:"xmas2019_ARCADE_2"`
		Xmas2019ARCADE3            bool   `json:"xmas2019_ARCADE_3"`
		Xmas2019ARCADE1            bool   `json:"xmas2019_ARCADE_1"`
		Xmas2019PTL3               bool   `json:"xmas2019_PTL_3"`
		ClaimedSoloBank            int64  `json:"claimed_solo_bank"`
		CurrentGadget              string `json:"currentGadget"`
		AnniversaryNPCVisited2020  []int  `json:"anniversaryNPCVisited2020"`
		AnniversaryNPCProgress2020 int    `json:"anniversaryNPCProgress2020"`
		ClaimedPotatoTalisman      int64  `json:"claimed_potato_talisman"`
		ClaimPotatoWarCrown        int64  `json:"claim_potato_war_crown"`
		ClaimedPotatoBasket        int64  `json:"claimed_potato_basket"`
		Summer2020Cooldowns        struct {
			MvpPlus1   bool `json:"MVP_PLUS1"`
			MvpPlus2   bool `json:"MVP_PLUS2"`
			MvpPlus0   bool `json:"MVP_PLUS0"`
			Superstar2 bool `json:"SUPERSTAR2"`
			Superstar1 bool `json:"SUPERSTAR1"`
			Superstar0 bool `json:"SUPERSTAR0"`
			Vip2       bool `json:"VIP2"`
			Vip1       bool `json:"VIP1"`
			Vip0       bool `json:"VIP0"`
			Mvp2       bool `json:"MVP2"`
			Mvp1       bool `json:"MVP1"`
			Mvp0       bool `json:"MVP0"`
			VipPlus2   bool `json:"VIP_PLUS2"`
			VipPlus1   bool `json:"VIP_PLUS1"`
			VipPlus0   bool `json:"VIP_PLUS0"`
			Normal0    bool `json:"NORMAL0"`
			Normal1    bool `json:"NORMAL1"`
			Normal2    bool `json:"NORMAL2"`
		} `json:"summer2020Cooldowns"`
		Rank                 string `json:"rank"`
		SkyblockFreeCookie   int64  `json:"skyblock_free_cookie"`
		Easter2021Cooldowns2 struct {
			MvpPlus3   bool `json:"MVP_PLUS3"`
			MvpPlus2   bool `json:"MVP_PLUS2"`
			MvpPlus1   bool `json:"MVP_PLUS1"`
			MvpPlus0   bool `json:"MVP_PLUS0"`
			Normal3    bool `json:"NORMAL3"`
			Normal2    bool `json:"NORMAL2"`
			Normal1    bool `json:"NORMAL1"`
			Normal0    bool `json:"NORMAL0"`
			VipPlus0   bool `json:"VIP_PLUS0"`
			VipPlus1   bool `json:"VIP_PLUS1"`
			VipPlus2   bool `json:"VIP_PLUS2"`
			VipPlus3   bool `json:"VIP_PLUS3"`
			Mvp0       bool `json:"MVP0"`
			Mvp1       bool `json:"MVP1"`
			Mvp2       bool `json:"MVP2"`
			Mvp3       bool `json:"MVP3"`
			Vip0       bool `json:"VIP0"`
			Vip1       bool `json:"VIP1"`
			Vip2       bool `json:"VIP2"`
			Superstar3 bool `json:"SUPERSTAR3"`
			Superstar2 bool `json:"SUPERSTAR2"`
			Superstar1 bool `json:"SUPERSTAR1"`
			Superstar0 bool `json:"SUPERSTAR0"`
			Vip3       bool `json:"VIP3"`
		} `json:"easter2021Cooldowns2"`
		Outfit struct {
		} `json:"outfit"`
		CurrentPet         string `json:"currentPet"`
		ScorpiusBribe144   int64  `json:"scorpius_bribe_144"`
		ClaimedYear143Cake int64  `json:"claimed_year143_cake"`
		MostRecentGameType string `json:"mostRecentGameType"`
	} `json:"player"`
}
