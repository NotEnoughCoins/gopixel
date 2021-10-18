package structs

type Achievements struct {
	Success      bool  `json:"success"`
	LastUpdated  int64 `json:"lastUpdated"`
	Achievements struct {
		Arcade struct {
			OneTime struct {
				CreeperAttackSurvival struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CREEPER_ATTACK_SURVIVAL"`
				FootballFiveGoals struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FOOTBALL_FIVE_GOALS"`
				Overpowered struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OVERPOWERED"`
				WorldEconomics struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WORLD_ECONOMICS"`
				CantHideFromMe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CANT_HIDE_FROM_ME"`
				OverHere struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OVER_HERE"`
				GiddyUpHorsey struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GIDDY_UP_HORSEY"`
				Dragontamer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DRAGONTAMER"`
				CtwHeyThere struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CTW_HEY_THERE"`
				PixelPaintersOne struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PIXEL_PAINTERS_ONE"`
				ShootingRangeExplosiveArrow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHOOTING_RANGE_EXPLOSIVE_ARROW"`
				HypixelSaysMaster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HYPIXEL_SAYS_MASTER"`
				NoMutinyToday struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_MUTINY_TODAY"`
				DragonWarsBlast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DRAGON_WARS_BLAST"`
				CtwICanBeAnything struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CTW_I_CAN_BE_ANYTHING"`
				CtwNoNeed struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CTW_NO_NEED"`
				MiniWallsLastMan struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MINI_WALLS_LAST_MAN"`
				ZombiesFeelsGood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIES_FEELS_GOOD"`
				PtbRideBat struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"PTB_RIDE_BAT"`
				HoehoehoeScore struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HOEHOEHOE_SCORE"`
				CreeperAttackWaves struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CREEPER_ATTACK_WAVES"`
				HideAndSeekPropHunter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HIDE_AND_SEEK_PROP_HUNTER"`
				HoleScore struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HOLE_SCORE"`
				CtwMvp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CTW_MVP"`
				FootballSpeed struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FOOTBALL_SPEED"`
				MiniWallsDamage struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MINI_WALLS_DAMAGE"`
				ZombesSerialKiller struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBES_SERIAL_KILLER"`
				ProfessionalMower struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PROFESSIONAL_MOWER"`
				ZombiesSpeedRunner struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIES_SPEED_RUNNER"`
				ZombiesUltimate struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIES_ULTIMATE"`
				PigFishingSuperBacon struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PIG_FISHING_SUPER_BACON"`
				BountyHunterTargetKiller struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOUNTY_HUNTER_TARGET_KILLER"`
				WoopsDidntMeanTo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WOOPS_DIDNT_MEAN_TO"`
				MiniHunter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MINI_HUNTER"`
				HideAndSeekPartyPooper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HIDE_AND_SEEK_PARTY_POOPER"`
				GottaCatchThemAll struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOTTA_CATCH_THEM_ALL"`
				ZombiesTimeTrialBlood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIES_TIME_TRIAL_BLOOD"`
				ZombiesTeamPlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIES_TEAM_PLAYER"`
				HideAndSeekProp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HIDE_AND_SEEK_PROP"`
				Rpg16RocketPig struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RPG_16_ROCKET_PIG"`
				CtwRightPlaceRightTime struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CTW_RIGHT_PLACE_RIGHT_TIME"`
				ZombiesHerobrine struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIES_HEROBRINE"`
				FarmHuntDisguise struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FARM_HUNT_DISGUISE"`
				ZombiesPerkHoarder struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIES_PERK_HOARDER"`
				HoleFinals struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HOLE_FINALS"`
				CtwFashionablyLate struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CTW_FASHIONABLY_LATE"`
				CtwSafetyIsAnIllusion struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CTW_SAFETY_IS_AN_ILLUSION"`
				CtwMagician struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CTW_MAGICIAN"`
				CreeperAttackIronGolem struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CREEPER_ATTACK_IRON_GOLEM"`
				NoMercy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_MERCY"`
				ZombiesWin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIES_WIN"`
				ThrowOutPowerupKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THROW_OUT_POWERUP_KILL"`
				CtwFirst struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CTW_FIRST"`
				ZombiesElectrician struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIES_ELECTRICIAN"`
				TeamSlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEAM_SLAYER"`
				CtwComeback struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CTW_COMEBACK"`
				FarmHuntKiller struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FARM_HUNT_KILLER"`
				AvalanceWaves struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"AVALANCE_WAVES"`
				ZombiesTimeTrialDead struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIES_TIME_TRIAL_DEAD"`
				CtwNinja struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CTW_NINJA"`
				PartyGamesStars struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PARTY_GAMES_STARS"`
				TrampolinioRedWool struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRAMPOLINIO_RED_WOOL"`
				DragonSlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DRAGON_SLAYER"`
				Untouchable struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNTOUCHABLE"`
				LoneSurvivor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LONE_SURVIVOR"`
				AnimalSlaughter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ANIMAL_SLAUGHTER"`
				Savage struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SAVAGE"`
				DragonKiller struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DRAGON_KILLER"`
				CreeperAttackUpgrades struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CREEPER_ATTACK_UPGRADES"`
				ZombiesSurvivor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIES_SURVIVOR"`
			} `json:"one_time"`
			Tiered struct {
				ZombiesNiceShot struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ZOMBIES_NICE_SHOT"`
				HideAndSeekHiderKills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HIDE_AND_SEEK_HIDER_KILLS"`
				ZombieKiller struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ZOMBIE_KILLER"`
				ZombiesRoundProgression struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ZOMBIES_ROUND_PROGRESSION"`
				CtwOhSheep struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CTW_OH_SHEEP"`
				TeamWork struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TEAM_WORK"`
				ArcadeBanker struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ARCADE_BANKER"`
				MiniwallsWinner struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MINIWALLS_WINNER"`
				BountyHunter struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BOUNTY_HUNTER"`
				ArcadeWinner struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ARCADE_WINNER"`
				ZombiesHighRound struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ZOMBIES_HIGH_ROUND"`
				FootballPro struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"FOOTBALL_PRO"`
				CtwSlayer struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CTW_SLAYER"`
				FarmhuntDominator struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"FARMHUNT_DOMINATOR"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"arcade"`
		Arena struct {
			OneTime struct {
				MyNewHat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MY_NEW_HAT"`
				Spartacus struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPARTACUS"`
				NiceSpare struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NICE_SPARE"`
				NewToy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NEW_TOY"`
				Magical struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAGICAL"`
				HealthTotem struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HEALTH_TOTEM"`
				Pig struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PIG"`
				Melee struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MELEE"`
				Punisher struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PUNISHER"`
				Energy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENERGY"`
				DoomShroomGloom struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DOOM_SHROOM_GLOOM"`
				Environmentalist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENVIRONMENTALIST"`
				HardEarnedReward struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HARD_EARNED_REWARD"`
				DeadlyPumpkin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DEADLY_PUMPKIN"`
				Offensive struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OFFENSIVE"`
				DontTouchMe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DONT_TOUCH_ME"`
				PowerHungry struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"POWER_HUNGRY"`
				NotEvenClose struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_EVEN_CLOSE"`
				Cooldown struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COOLDOWN"`
				DragonWithin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DRAGON_WITHIN"`
				Pool struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"POOL"`
				Smash struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SMASH"`
				Flawless struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FLAWLESS"`
				Runic struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RUNIC"`
				Overkill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OVERKILL"`
				BigBrawler struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BIG_BRAWLER"`
				IronHeart struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"IRON_HEART"`
				Health struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HEALTH"`
				Pairs struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PAIRS"`
				Support struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUPPORT"`
				DiscoStar struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DISCO_STAR"`
				YouShallNotPass struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"YOU_SHALL_NOT_PASS"`
				TotemDestroyer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TOTEM_DESTROYER"`
				MaxRunicMagic struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_RUNIC_MAGIC"`
				NotToday struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_TODAY"`
				SpiritedAway struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPIRITED_AWAY"`
			} `json:"one_time"`
			Tiered struct {
				Powerup struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"POWERUP"`
				Bossed struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BOSSED"`
				GottaWearEmAll struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GOTTA_WEAR_EM_ALL"`
				MagicalBox struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MAGICAL_BOX"`
				Gladiator struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GLADIATOR"`
				ClimbTheRanks struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CLIMB_THE_RANKS"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"arena"`
		Bedwars struct {
			OneTime struct {
				IronPunch struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"IRON_PUNCH"`
				Strategist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"STRATEGIST"`
				GearedUp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GEARED_UP"`
				Distraction struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DISTRACTION"`
				Minefield struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MINEFIELD"`
				EmeraldHoarder struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EMERALD_HOARDER"`
				Ninja struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NINJA"`
				ThatsAFirst struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THATS_A_FIRST"`
				YouCantDoThat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"YOU_CANT_DO_THAT"`
				AlreadyOver struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ALREADY_OVER"`
				Sniper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SNIPER"`
				CuttingItClose struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CUTTING_IT_CLOSE"`
				BuggyBeds struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BUGGY_BEDS"`
				KatnissEverdeenStyle struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KATNISS_EVERDEEN_STYLE"`
				Builder struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BUILDER"`
				OutOfStock struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OUT_OF_STOCK"`
				SuperLooter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUPER_LOOTER"`
				Revenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REVENGE"`
				Alchemist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ALCHEMIST"`
				UltimateDefense struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ULTIMATE_DEFENSE"`
				Fireballs struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIREBALLS"`
				Bomber struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOMBER"`
				ShearLuck struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHEAR_LUCK"`
				Survivor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SURVIVOR"`
				DiamondHoarder struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DIAMOND_HOARDER"`
				Slayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SLAYER"`
				SavvyShopper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SAVVY_SHOPPER"`
				BedTrap struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BED_TRAP"`
				ItsDarkDownThere struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ITS_DARK_DOWN_THERE"`
				TeamPlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEAM_PLAYER"`
				MissionControl struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MISSION_CONTROL"`
				DontNeedBed struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DONT_NEED_BED"`
				Merciless struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MERCILESS"`
				StayAwayFromMe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"STAY_AWAY_FROM_ME"`
				Golem struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOLEM"`
				RejoiningTheDream struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REJOINING_THE_DREAM"`
				TheLastOfUs struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_LAST_OF_US"`
				SneakyRusher struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SNEAKY_RUSHER"`
				PickaxeChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PICKAXE_CHALLENGE"`
				First struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST"`
				DestroyBeds struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DESTROY_BEDS"`
				FirstBlood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_BLOOD"`
				GettingTheJobDoneBetter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GETTING_THE_JOB_DONE_BETTER"`
			} `json:"one_time"`
			Tiered struct {
				CollectorsEdition struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COLLECTORS_EDITION"`
				BedwarsKiller struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BEDWARS_KILLER"`
				Level struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"LEVEL"`
				Wins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS"`
				Beds struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BEDS"`
				LootBox struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"LOOT_BOX"`
				BedwarsChallenger struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BEDWARS_CHALLENGER"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"bedwars"`
		Blitz struct {
			OneTime struct {
				DonkeytamerMaster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DONKEYTAMER_MASTER"`
				FindBlitz struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIND_BLITZ"`
				Titanium struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TITANIUM"`
				Rampage struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RAMPAGE"`
				Nickname struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NICKNAME"`
				WinBeforeDeathmatch struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_BEFORE_DEATHMATCH"`
				LuckyMinions struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LUCKY_MINIONS"`
				FullInventory struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FULL_INVENTORY"`
				Finally struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FINALLY"`
				BlitzManiac struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLITZ_MANIAC"`
				UltimatePrestiger struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ULTIMATE_PRESTIGER"`
				SoShiny struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SO_SHINY"`
				EnchantSword struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENCHANT_SWORD"`
				EnchantedArmor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENCHANTED_ARMOR"`
				SpawnHorse struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPAWN_HORSE"`
				Champion struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHAMPION"`
				Invincible struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"INVINCIBLE"`
				FindHead struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIND_HEAD"`
				SevenKits struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SEVEN_KITS"`
				PhoenixMaster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PHOENIX_MASTER"`
				FishKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FISH_KILL"`
				NoRegrets struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_REGRETS"`
				UltimateCombatant struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ULTIMATE_COMBATANT"`
				UltimateCompletist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ULTIMATE_COMPLETIST"`
				CookingExpert struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COOKING_EXPERT"`
				FirstRanger struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_RANGER"`
				Pigrider struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PIGRIDER"`
				Rambo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RAMBO"`
				Collector struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COLLECTOR"`
				CraftBread struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CRAFT_BREAD"`
				FirstBlood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_BLOOD"`
				CloseCall struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CLOSE_CALL"`
				UseWolfTamer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"USE_WOLF_TAMER"`
				Experimentation struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXPERIMENTATION"`
				SafetyFirst struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SAFETY_FIRST"`
				FirstGame struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_GAME"`
				RabbitsFoot struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RABBITS_FOOT"`
				KillWithoutKit struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_WITHOUT_KIT"`
				Apocalypse struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"APOCALYPSE"`
				NotSkywars struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_SKYWARS"`
				Unfortunate struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNFORTUNATE"`
				BrutalWarrior struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BRUTAL_WARRIOR"`
				TooBasic struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TOO_BASIC"`
				SailAway struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SAIL_AWAY"`
				GetDiamondSword struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GET_DIAMOND_SWORD"`
				Unstoppable struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNSTOPPABLE"`
				NoProblem struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_PROBLEM"`
				Afterburner struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"AFTERBURNER"`
				MaxBlitz struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_BLITZ"`
				EvenShinier struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EVEN_SHINIER"`
				CoinFestival struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COIN_FESTIVAL"`
				Bomberman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOMBERMAN"`
				Preference struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PREFERENCE"`
				NoLooting struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_LOOTING"`
				LevelSeven struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEVEL_SEVEN"`
				UnderTheSea struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNDER_THE_SEA"`
			} `json:"one_time"`
			Tiered struct {
				KitCollector struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KIT_COLLECTOR"`
				KitExpert struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KIT_EXPERT"`
				RngMaster struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"RNG_MASTER"`
				MasterOfKits struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MASTER_OF_KITS"`
				WarVeteran struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WAR_VETERAN"`
				Kills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS"`
				RangedCombat struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"RANGED_COMBAT"`
				Coins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COINS"`
				Looter struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"LOOTER"`
				FightingExpert struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"FIGHTING_EXPERT"`
				TreasureSeeker struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TREASURE_SEEKER"`
				KitExperienceCollector struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KIT_EXPERIENCE_COLLECTOR"`
				MobMaster struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MOB_MASTER"`
				Wins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS"`
				WinsTeams struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS_TEAMS"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"blitz"`
		Buildbattle struct {
			OneTime struct {
				GuessingStreak struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GUESSING_STREAK"`
				OooShiny struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OOO_SHINY"`
				Mobster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MOBSTER"`
				Legendary struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY"`
				FastTyper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FAST_TYPER"`
				Intuition struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"INTUITION"`
				ProfessionalBuilder struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PROFESSIONAL_BUILDER"`
				ClassicMan struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CLASSIC_MAN"`
				Stenographer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"STENOGRAPHER"`
				Obvious struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OBVIOUS"`
				NoMistakes struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_MISTAKES"`
				Braniac struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BRANIAC"`
				Artist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ARTIST"`
				Fancy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FANCY"`
				Over99 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OVER_99"`
				ProWinner struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PRO_WINNER"`
				Teamwork struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEAMWORK"`
				ThatWoodBePerfect struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THAT_WOOD_BE_PERFECT"`
				EverySecondCounts struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EVERY_SECOND_COUNTS"`
				SuperiorVote struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUPERIOR_VOTE"`
				Musician struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MUSICIAN"`
				PerfectHarmony struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PERFECT_HARMONY"`
				TheHatMaster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_HAT_MASTER"`
			} `json:"one_time"`
			Tiered struct {
				BuildBattleVoter struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BUILD_BATTLE_VOTER"`
				GuessTheBuildGuesses struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GUESS_THE_BUILD_GUESSES"`
				BuildBattlePoints struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BUILD_BATTLE_POINTS"`
				GuessTheBuildWinner struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GUESS_THE_BUILD_WINNER"`
				BuildBattleScore struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BUILD_BATTLE_SCORE"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"buildbattle"`
		Christmas2017 struct {
			OneTime struct {
				GreedIncarnate struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GREED_INCARNATE"`
				RespectYourElder struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RESPECT_YOUR_ELDER"`
				ExplosiveCandy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXPLOSIVE_CANDY"`
				ChristmasIsSaved struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHRISTMAS_IS_SAVED"`
				SnowWars struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SNOW_WARS"`
				RealSanta2020 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REAL_SANTA_2020"`
				ToWar struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TO_WAR"`
				SpreadingLove struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPREADING_LOVE"`
				BigBagOGifts struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BIG_BAG_O_GIFTS"`
				Rush struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RUSH"`
				NewYearsCelebrations struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NEW_YEARS_CELEBRATIONS"`
				HolidayMiracle struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HOLIDAY_MIRACLE"`
				DoYouWannaBuildASnowman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DO_YOU_WANNA_BUILD_A_SNOWMAN"`
				BouncyCastle struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOUNCY_CASTLE"`
				NotMyMom struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_MY_MOM"`
				HolidaysRuined struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HOLIDAYS_RUINED"`
				CloseEnough struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CLOSE_ENOUGH"`
				WinterRunner struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WINTER_RUNNER"`
				Legendary struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY"`
				IceBreaker struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ICE_BREAKER"`
				GroopoReturns struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GROOPO_RETURNS"`
				ChristmasTopper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHRISTMAS_TOPPER"`
				FriendlyFire struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FRIENDLY_FIRE"`
				ChristmasQuest struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHRISTMAS_QUEST"`
				UhUh struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UH_UH"`
				HuntBegins2020 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HUNT_BEGINS_2020"`
				Blacksmith struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLACKSMITH"`
				SugarRush struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUGAR_RUSH"`
				DemCows struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DEM_COWS"`
				LetItSnow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LET_IT_SNOW"`
				MerryChristmas struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MERRY_CHRISTMAS"`
				TodaysTheDay struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TODAYS_THE_DAY"`
				Swegmas struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SWEGMAS"`
				IceSpreader struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ICE_SPREADER"`
				Festivities struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FESTIVITIES"`
				MeltingKiller struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MELTING_KILLER"`
				EmptyHouse struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EMPTY_HOUSE"`
				NomNom struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOM_NOM"`
				SharingIsCaring struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"SHARING_IS_CARING"`
				EatThis struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"EAT_THIS"`
				HuntBegins2019 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"HUNT_BEGINS_2019"`
				RealSanta2019 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"REAL_SANTA_2019"`
				RealSanta struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"REAL_SANTA"`
				StealFairly struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"STEAL_FAIRLY"`
				HuntBegins struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"HUNT_BEGINS"`
				SantaHelper struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"SANTA_HELPER"`
				SecretestSanta struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"SECRETEST_SANTA"`
				RealSanta2018 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"REAL_SANTA_2018"`
				HuntBegins2018 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"HUNT_BEGINS_2018"`
			} `json:"one_time"`
			Tiered struct {
				BestPresents struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BEST_PRESENTS"`
				Advent2020 struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ADVENT_2020"`
				PresentCollector struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"PRESENT_COLLECTOR"`
				SantaSaysRounds struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SANTA_SAYS_ROUNDS"`
				NoChristmas struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"NO_CHRISTMAS"`
				SecretSanta struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SECRET_SANTA"`
				Advent2019 struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ADVENT_2019"`
				Advent struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ADVENT"`
				Advent2018 struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ADVENT_2018"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"christmas2017"`
		Copsandcrims struct {
			OneTime struct {
				OhBabyATriple struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OH_BABY_A_TRIPLE"`
				GoldenDeagle struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOLDEN_DEAGLE"`
				OneShotOneKil struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ONE_SHOT_ONE_KIL"`
				RunAway struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RUN_AWAY"`
				ChallengeCompleted struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHALLENGE_COMPLETED"`
				ArmedAndDangerous struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ARMED_AND_DANGEROUS"`
				LikeMyGun struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LIKE_MY_GUN"`
				FancyNewToys struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FANCY_NEW_TOYS"`
				RaiseYourVoice struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RAISE_YOUR_VOICE"`
				Unstoppable struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNSTOPPABLE"`
				Umadbro struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UMADBRO"`
				Mp5Streak struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MP5_STREAK"`
				DarudeSandstorm struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DARUDE_SANDSTORM"`
				P90Streak struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"P90_STREAK"`
				GunMaster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GUN_MASTER"`
				WheresMyNuke struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WHERES_MY_NUKE"`
				HomingBullets struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HOMING_BULLETS"`
				Pentaaaa struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PENTAAAA"`
				HuntingSeason struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HUNTING_SEASON"`
				BlastedToTheMoon struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLASTED_TO_THE_MOON"`
				TheCarrier struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_CARRIER"`
				SpecialLooks struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPECIAL_LOOKS"`
				OnFire struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ON_FIRE"`
				ComebackStory struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COMEBACK_STORY"`
				SneakKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SNEAK_KILL"`
				WarfareStylist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WARFARE_STYLIST"`
				IsItGoodNow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"IS_IT_GOOD_NOW"`
				HalfTheBattle struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HALF_THE_BATTLE"`
				CharacterUpgrades struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHARACTER_UPGRADES"`
				FriendlyFire struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FRIENDLY_FIRE"`
				StreetArtist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"STREET_ARTIST"`
				NeverTooLate struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NEVER_TOO_LATE"`
				QuadFlash struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"QUAD_FLASH"`
				SamuraiWarrior struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SAMURAI_WARRIOR"`
				SecretOrder struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SECRET_ORDER"`
				TeamCommunicator struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEAM_COMMUNICATOR"`
				Ace struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ACE"`
				TooEasy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TOO_EASY"`
				JackOfAllTrades struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JACK_OF_ALL_TRADES"`
				WreckingMachine struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WRECKING_MACHINE"`
				Kamikaze struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KAMIKAZE"`
				CloseCall struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CLOSE_CALL"`
				Sniper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SNIPER"`
				GrafittiKing struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GRAFITTI_KING"`
				BlindKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLIND_KILL"`
				GrenadeKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GRENADE_KILL"`
				DeathmatchSpecialist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DEATHMATCH_SPECIALIST"`
				FriendlyGame struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FRIENDLY_GAME"`
				Sneaky struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SNEAKY"`
				Triipplee struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRIIPPLEE"`
				Flawless struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FLAWLESS"`
			} `json:"one_time"`
			Tiered struct {
				CacBanker struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CAC_BANKER"`
				BombSpecialist struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BOMB_SPECIALIST"`
				PracticeMakesPerfect struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"PRACTICE_MAKES_PERFECT"`
				SerialKiller struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SERIAL_KILLER"`
				HeadshotKills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HEADSHOT_KILLS"`
				GrenadeKills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GRENADE_KILLS"`
				HeroTerrorist struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HERO_TERRORIST"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"copsandcrims"`
		Duels struct {
			OneTime struct {
				GoneFishing struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GONE_FISHING"`
				OnFire struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ON_FIRE"`
				CloseCall struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CLOSE_CALL"`
				ShutDown struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHUT_DOWN"`
				MyPreference struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MY_PREFERENCE"`
				WellRounded struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WELL_ROUNDED"`
				CommunityOriented struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COMMUNITY_ORIENTED"`
				Gg struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GG"`
				Carried struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CARRIED"`
				Replay struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REPLAY"`
				Domination struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DOMINATION"`
				HawkEye struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAWK_EYE"`
				Revenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REVENGE"`
				TheWaitingGame struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_WAITING_GAME"`
				Fortification struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FORTIFICATION"`
				BurnBabyBurn struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BURN_BABY_BURN"`
				TrialByCombat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRIAL_BY_COMBAT"`
				BuildBattle struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BUILD_BATTLE"`
				Rematch struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REMATCH"`
				Summoner struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUMMONER"`
				Untouchable struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNTOUCHABLE"`
				ExpressYourself struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXPRESS_YOURSELF"`
				Hungry struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HUNGRY"`
				LobbySlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LOBBY_SLAYER"`
				SpeedDuel struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPEED_DUEL"`
				Ninja struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NINJA"`
				NotCloseAtAll struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_CLOSE_AT_ALL"`
				VoidArcher struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VOID_ARCHER"`
				SpeedySumo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPEEDY_SUMO"`
				GotYa struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOT_YA"`
				LastStand struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LAST_STAND"`
				HatTrick struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAT_TRICK"`
				GettingLoot struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GETTING_LOOT"`
				AxeYouAQuestion struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"AXE_YOU_A_QUESTION"`
				CleanSweep struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CLEAN_SWEEP"`
				HeartHoarder struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HEART_HOARDER"`
				Ace struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ACE"`
				NotHungary struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_HUNGARY"`
				TeamPlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEAM_PLAYER"`
				OneVOneMe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ONE_V_ONE_ME"`
				MoreDamage struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"MORE_DAMAGE"`
				SmashYourKeyboard struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"SMASH_YOUR_KEYBOARD"`
				WeAreNumberOne struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"WE_ARE_NUMBER_ONE"`
				Competitor struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"COMPETITOR"`
			} `json:"one_time"`
			Tiered struct {
				DuelsDivision struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"DUELS_DIVISION"`
				BridgeFourTeamsWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BRIDGE_FOUR_TEAMS_WINS"`
				DuelsWinner struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"DUELS_WINNER"`
				DuelsWinStreak struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"DUELS_WIN_STREAK"`
				BridgeWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BRIDGE_WINS"`
				BridgeDoublesWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BRIDGE_DOUBLES_WINS"`
				Goals struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GOALS"`
				BridgeTeamsWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BRIDGE_TEAMS_WINS"`
				DuelsTraveller struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"DUELS_TRAVELLER"`
				BridgeWinStreak struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BRIDGE_WIN_STREAK"`
				BridgeDuelsWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BRIDGE_DUELS_WINS"`
				UniqueMapWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"UNIQUE_MAP_WINS"`
				DuelsTrophies struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"DUELS_TROPHIES"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"duels"`
		Easter struct {
			OneTime struct {
				AllEggs2019 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ALL_EGGS_2019"`
				HappyEaster2019 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"HAPPY_EASTER_2019"`
				FirstEgg2019 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"FIRST_EGG_2019"`
				FirstEgg2020 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"FIRST_EGG_2020"`
				AllEggs2020 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ALL_EGGS_2020"`
				HappyEaster2020 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"HAPPY_EASTER_2020"`
				EggBetrayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EGG_BETRAYER"`
				SmashCluckKills struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SMASH_CLUCK_KILLS"`
				SpringMaiden struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPRING_MAIDEN"`
				EnvironmentProtection struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENVIRONMENT_PROTECTION"`
				AllEggs2021 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ALL_EGGS_2021"`
				EggMaster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EGG_MASTER"`
				CloseEnough struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CLOSE_ENOUGH"`
				BlitzSpawnRabbit struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLITZ_SPAWN_RABBIT"`
				CvcGrenades struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CVC_GRENADES"`
				EggAssault struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EGG_ASSAULT"`
				ArcadeChickenRace struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ARCADE_CHICKEN_RACE"`
				EggLayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EGG_LAYER"`
				WhatIsThis struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WHAT_IS_THIS"`
				EasterEgg struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EASTER_EGG"`
				PaintballLeeroy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PAINTBALL_LEEROY"`
				InnerRabbit struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"INNER_RABBIT"`
				MmCarrotKills struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MM_CARROT_KILLS"`
				UhcGoldenCarrot struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UHC_GOLDEN_CARROT"`
				SwEggVoid struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SW_EGG_VOID"`
				WhatHaveWeDone struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WHAT_HAVE_WE_DONE"`
				SpringHero struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPRING_HERO"`
				PitDragonEgg struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PIT_DRAGON_EGG"`
				WhatsThis struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WHATS_THIS"`
				SpringFishing struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPRING_FISHING"`
				HappyEaster2021 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAPPY_EASTER_2021"`
				TowerwarsChickens struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TOWERWARS_CHICKENS"`
				BwJumpBoost struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BW_JUMP_BOOST"`
				MegawallsJockey struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MEGAWALLS_JOCKEY"`
				FasterThanWind struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FASTER_THAN_WIND"`
				YouDidntSeeThatComing struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"YOU_DIDNT_SEE_THAT_COMING"`
				FirstEgg2021 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_EGG_2021"`
				VampirezCarrot struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VAMPIREZ_CARROT"`
			} `json:"one_time"`
			Tiered struct {
				ThrowEggs struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"THROW_EGGS"`
				MasterTracker struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MASTER_TRACKER"`
				EggFinder struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"EGG_FINDER"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"easter"`
		General struct {
			OneTime struct {
				TreasureHunt2021 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TREASURE_HUNT_2021"`
				Vip struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VIP"`
				Friends25 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FRIENDS_25"`
				Youtuber struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"YOUTUBER"`
				FirstJoin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_JOIN"`
				UsePortal struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"USE_PORTAL"`
				HotPotato struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HOT_POTATO"`
				FirstGame struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_GAME"`
				FirstChat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_CHAT"`
				FirstFriend struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_FRIEND"`
				AchievementNpc struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ACHIEVEMENT_NPC"`
				VipPlus struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VIP_PLUS"`
				Creeperbook struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CREEPERBOOK"`
				FirstParty struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_PARTY"`
				ALongJourneyBegins struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"A_LONG_JOURNEY_BEGINS"`
				TreasureHunt struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"TREASURE_HUNT"`
				Gifting struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"GIFTING"`
			} `json:"one_time"`
			Tiered struct {
				Challenger struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CHALLENGER"`
				Coins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COINS"`
				QuestMaster struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"QUEST_MASTER"`
				Wins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"general"`
		Gingerbread struct {
			OneTime struct {
				IsThisSurvivalGames struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"IS_THIS_SURVIVAL_GAMES"`
				Tryharder struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRYHARDER"`
				MissileMayhem struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MISSILE_MAYHEM"`
				Slippery struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SLIPPERY"`
				CantSeeAnything struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CANT_SEE_ANYTHING"`
				ShutDown struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHUT_DOWN"`
				TasteMyBanana struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TASTE_MY_BANANA"`
				NewStyle struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NEW_STYLE"`
				GettingGood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GETTING_GOOD"`
				CoinMagnet struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COIN_MAGNET"`
				EternallyAwesome struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ETERNALLY_AWESOME"`
				ChillOut struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHILL_OUT"`
				NeverLucky struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NEVER_LUCKY"`
				Modular struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MODULAR"`
				HonkingAmazing struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HONKING_AMAZING"`
				JavelinsRocket struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JAVELINS_ROCKET"`
				Mechanic struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MECHANIC"`
				SurfaceToAirMissile struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SURFACE_TO_AIR_MISSILE"`
				GetHitByMe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GET_HIT_BY_ME"`
				Ungrateful struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNGRATEFUL"`
				Contender struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CONTENDER"`
				GettinPaid struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GETTIN_PAID"`
				HonkingMad struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HONKING_MAD"`
				CoinsPlease struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COINS_PLEASE"`
				Lapped struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LAPPED"`
				SeenItAll struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SEEN_IT_ALL"`
				BoostOfConfidence struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOOST_OF_CONFIDENCE"`
				WellVersed struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WELL_VERSED"`
				SleeperAgent struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SLEEPER_AGENT"`
				ShowOff struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHOW_OFF"`
				FlowerPower struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FLOWER_POWER"`
				SeasonalRacer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SEASONAL_RACER"`
				Nope struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOPE"`
				VictorsPrize struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VICTORS_PRIZE"`
				EatMyDust struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EAT_MY_DUST"`
				ImLucky struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"IM_LUCKY"`
				FrontMan struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FRONT_MAN"`
			} `json:"one_time"`
			Tiered struct {
				Winner struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINNER"`
				Banker struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BANKER"`
				Racer struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"RACER"`
				Mystery struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MYSTERY"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"gingerbread"`
		Halloween2017 struct {
			OneTime struct {
				PumpkinDeath struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PUMPKIN_DEATH"`
				SugarRush2 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUGAR_RUSH_2"`
				FullBatMode struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FULL_BAT_MODE"`
				PumpkinVision struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PUMPKIN_VISION"`
				SpookyLooks struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPOOKY_LOOKS"`
				TrickOrTreatTime struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRICK_OR_TREAT_TIME"`
				BlameYourTeam struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLAME_YOUR_TEAM"`
				VampiresBeGone struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VAMPIRES_BE_GONE"`
				PumpkinDancer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PUMPKIN_DANCER"`
				RisingDead struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RISING_DEAD"`
				Gravedigger struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GRAVEDIGGER"`
				SurvivorsBeGone struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SURVIVORS_BE_GONE"`
				FiveBaskets2021 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIVE_BASKETS_2021"`
				ThatWasEasy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THAT_WAS_EASY"`
				MyHalloweenCostume struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MY_HALLOWEEN_COSTUME"`
				DevilsRide struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DEVILS_RIDE"`
				KillerLoot struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILLER_LOOT"`
				OctoberBetrayal struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OCTOBER_BETRAYAL"`
				PumpkinsWillRise struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PUMPKINS_WILL_RISE"`
				MasterAlchemist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MASTER_ALCHEMIST"`
				TisTheSeason struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TIS_THE_SEASON"`
				NotSoScary struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_SO_SCARY"`
				ThatTimeOfYear struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THAT_TIME_OF_YEAR"`
				SmokingVeil struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SMOKING_VEIL"`
				FearThePumpkinator struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FEAR_THE_PUMPKINATOR"`
				FireFromHell struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRE_FROM_HELL"`
				Tricked struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRICKED"`
				TbrSeasideDrive struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TBR_SEASIDE_DRIVE"`
				CorpseStillRuns struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CORPSE_STILL_RUNS"`
				TbrClockTowerTripFast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TBR_CLOCK_TOWER_TRIP_FAST"`
				LooseFloorboard struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LOOSE_FLOORBOARD"`
				FullMoon struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FULL_MOON"`
				TbrPumpkinJumpFast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TBR_PUMPKIN_JUMP_FAST"`
				BeFree struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BE_FREE"`
				MasterBobber struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MASTER_BOBBER"`
				SeeThrough struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SEE_THROUGH"`
				BurningTouch struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BURNING_TOUCH"`
				UndeadTargetPractice struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNDEAD_TARGET_PRACTICE"`
				TheCrawlingDead struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_CRAWLING_DEAD"`
				SpookyVictory struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPOOKY_VICTORY"`
				SpookyChest struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPOOKY_CHEST"`
				ClassyWither struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CLASSY_WITHER"`
				GoodTry struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOOD_TRY"`
				TbrSeasideDriveFast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TBR_SEASIDE_DRIVE_FAST"`
				HauntedMaps struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAUNTED_MAPS"`
				AllBaskets2021 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ALL_BASKETS_2021"`
				HiThere struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HI_THERE"`
				TbrClockTowerTrip struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TBR_CLOCK_TOWER_TRIP"`
				TbrPumpkinJump struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TBR_PUMPKIN_JUMP"`
				AllGhosts2018 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ALL_GHOSTS_2018"`
				FiveBaskets2020 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"FIVE_BASKETS_2020"`
				SecondGhost struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"SECOND_GHOST"`
				TbrKrakenAssault125 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"TBR_KRAKEN_ASSAULT_1_25"`
				TbrObservatorySpin struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"TBR_OBSERVATORY_SPIN"`
				TbrSharknadoJaunt struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"TBR_SHARKNADO_JAUNT"`
				AllBaskets struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ALL_BASKETS"`
				TbrSharknadoJauntFast struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"TBR_SHARKNADO_JAUNT_FAST"`
				AllBaskets2020 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ALL_BASKETS_2020"`
				AllGhosts struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ALL_GHOSTS"`
				TbrKrakenAssault struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"TBR_KRAKEN_ASSAULT"`
				TbrMidtownTrip struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"TBR_MIDTOWN_TRIP"`
				SugarRush struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"SUGAR_RUSH"`
				TbrObservatorySpin110 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"TBR_OBSERVATORY_SPIN_1_10"`
				FiveBaskets struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"FIVE_BASKETS"`
				SecondGhost2018 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"SECOND_GHOST_2018"`
				TbrMidtownTrip120 struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"TBR_MIDTOWN_TRIP_1_20"`
				ElSpooder struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"EL_SPOODER"`
				Necrotherapy struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"NECROTHERAPY"`
			} `json:"one_time"`
			Tiered struct {
				Pumpkinator struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"PUMPKINATOR"`
				PumpkinSmasher struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"PUMPKIN_SMASHER"`
				BraveTheDark struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BRAVE_THE_DARK"`
				GetThemApples struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GET_THEM_APPLES"`
				CandyHoarder struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CANDY_HOARDER"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"halloween2017"`
		Housing struct {
			OneTime struct {
				JoinFriend struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JOIN_FRIEND"`
				RecieveCookie struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RECIEVE_COOKIE"`
				JoinYoutube struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JOIN_YOUTUBE"`
				NewLook struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NEW_LOOK"`
				JoinRandom struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JOIN_RANDOM"`
				GiveCookie struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GIVE_COOKIE"`
				JoinStaff struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JOIN_STAFF"`
				MakeResident struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAKE_RESIDENT"`
				GrandOpening struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GRAND_OPENING"`
				BecomeResident struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BECOME_RESIDENT"`
				JoinGuild struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JOIN_GUILD"`
			} `json:"one_time"`
			Tiered struct {
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"housing"`
		Murdermystery struct {
			OneTime struct {
				BlessingAndCurse struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLESSING_AND_CURSE"`
				KillMurdererAsLastAlive struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_MURDERER_AS_LAST_ALIVE"`
				WinSurvivorDueToTime struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_SURVIVOR_DUE_TO_TIME"`
				ThirtyGoldPickedUp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THIRTY_GOLD_PICKED_UP"`
				ShootThrownKnife struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHOOT_THROWN_KNIFE"`
				TrapperTerritory struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRAPPER_TERRITORY"`
				SpecialTwoInARow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPECIAL_TWO_IN_A_ROW"`
				Jaws struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JAWS"`
				WinAsMurdererShortTime struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_AS_MURDERER_SHORT_TIME"`
				PlayBothGames struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PLAY_BOTH_GAMES"`
				KillInRapidTransport struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_IN_RAPID_TRANSPORT"`
				Disinfectant struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DISINFECTANT"`
				BlockWithBarrier struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLOCK_WITH_BARRIER"`
				SixthSense struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SIXTH_SENSE"`
				FiveCurses struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIVE_CURSES"`
				WinWhileInvincible struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_WHILE_INVINCIBLE"`
				FirstShotHit struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_SHOT_HIT"`
				MurdererFirstKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MURDERER_FIRST_KILL"`
				ThreeKnifeThrowKills struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THREE_KNIFE_THROW_KILLS"`
				KillMurdererAfterKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_MURDERER_AFTER_KILL"`
				DrinkManyPotions struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DRINK_MANY_POTIONS"`
				KillMurdererWhileBlinded struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_MURDERER_WHILE_BLINDED"`
				LastSurvivor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LAST_SURVIVOR"`
				PlayGameInLobby struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PLAY_GAME_IN_LOBBY"`
				SoldiersEliminated struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SOLDIERS_ELIMINATED"`
				Watson struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WATSON"`
				DoubleDuty struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DOUBLE_DUTY"`
				KillDetectiveFast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_DETECTIVE_FAST"`
				ClearCacti struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CLEAR_CACTI"`
				Killstreak struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILLSTREAK"`
				KaliFavor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KALI_FAVOR"`
				PickupGold struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PICKUP_GOLD"`
				BowKillstreak struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOW_KILLSTREAK"`
				BeTheHero struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BE_THE_HERO"`
				ThreeArrows struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THREE_ARROWS"`
				SwordKillLongRange struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SWORD_KILL_LONG_RANGE"`
				RideMonorail struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RIDE_MONORAIL"`
				BowKillOnDetective struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOW_KILL_ON_DETECTIVE"`
				ReichenbachFall struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REICHENBACH_FALL"`
				WinInnoMonorail struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_INNO_MONORAIL"`
				Emp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EMP"`
				Uncalculated struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNCALCULATED"`
				SecretChamber struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SECRET_CHAMBER"`
				Paranoid struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PARANOID"`
				LongrangeShot struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LONGRANGE_SHOT"`
				ReturnFromDeadWin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RETURN_FROM_DEAD_WIN"`
				WinMurdererFellInTrap struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_MURDERER_FELL_IN_TRAP"`
				NoGoldPickup struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_GOLD_PICKUP"`
				TopZombie struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TOP_ZOMBIE"`
				Dance struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DANCE"`
				SurviveStormOnTop struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SURVIVE_STORM_ON_TOP"`
				XomB struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"XOM_B"`
				AllDirections struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ALL_DIRECTIONS"`
				MurdererBowKills struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MURDERER_BOW_KILLS"`
				Stealth struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"STEALTH"`
				HitBySwordWhileInvinc struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HIT_BY_SWORD_WHILE_INVINC"`
				KillWhileMonorail struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_WHILE_MONORAIL"`
				TriggerHappyHavoc struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRIGGER_HAPPY_HAVOC"`
				Tricked struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRICKED"`
				KillOnHorse struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_ON_HORSE"`
			} `json:"one_time"`
			Tiered struct {
				Countermeasures struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COUNTERMEASURES"`
				Hitman struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HITMAN"`
				WinsAsSurvivor struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS_AS_SURVIVOR"`
				KillsAsMurderer struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS_AS_MURDERER"`
				WinsAsMurderer struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS_AS_MURDERER"`
				Brainiac struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BRAINIAC"`
				Hoarder struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HOARDER"`
				SurvivalSkills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SURVIVAL_SKILLS"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"murdermystery"`
		Paintball struct {
			OneTime struct {
				ActivateKillstreaks struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ACTIVATE_KILLSTREAKS"`
				CheatingDeath struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHEATING_DEATH"`
				ActivatePlusTen struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ACTIVATE_PLUS_TEN"`
				TriggerHappy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRIGGER_HAPPY"`
				EnergyDrink struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENERGY_DRINK"`
				ThunderStruck struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THUNDER_STRUCK"`
				Combo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COMBO"`
				GodKiller struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOD_KILLER"`
				LightsOut struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LIGHTS_OUT"`
				NowYouSeeMe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOW_YOU_SEE_ME"`
				NoKillstreaks struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_KILLSTREAKS"`
				LastKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LAST_KILL"`
				ActivateLeeroyJenkins struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ACTIVATE_LEEROY_JENKINS"`
				AtLeastITried struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"AT_LEAST_I_TRIED"`
				Hacker struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HACKER"`
				UnlockKillstreaks struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNLOCK_KILLSTREAKS"`
				UndercoverSloth struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNDERCOVER_SLOTH"`
				ExplosiveDeath struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXPLOSIVE_DEATH"`
				Sarcrifice struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SARCRIFICE"`
				FirstKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_KILL"`
				Closure struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CLOSURE"`
				OnTheBrinkOfDefeat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ON_THE_BRINK_OF_DEFEAT"`
				ShakyHands struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHAKY_HANDS"`
				HowDoesItFeel struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HOW_DOES_IT_FEEL"`
				Sampling struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SAMPLING"`
				GoHomeYoureDrunk struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GO_HOME_YOURE_DRUNK"`
				AdminHat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ADMIN_HAT"`
				UnlockHat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNLOCK_HAT"`
				Godfather struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GODFATHER"`
				WarfareTime struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WARFARE_TIME"`
				DoIGetMyNukeNow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DO_I_GET_MY_NUKE_NOW"`
				JacksonPollock struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JACKSON_POLLOCK"`
				Endurance struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENDURANCE"`
				Espionage struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ESPIONAGE"`
			} `json:"one_time"`
			Tiered struct {
				Wins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS"`
				HatCollector struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HAT_COLLECTOR"`
				Invincible struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"INVINCIBLE"`
				KillStreaks struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILL_STREAKS"`
				Coins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COINS"`
				Kills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"paintball"`
		Pit struct {
			OneTime struct {
				MineObsidian struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MINE_OBSIDIAN"`
				PunchSword struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PUNCH_SWORD"`
				WinAuction struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_AUCTION"`
				DoATrade struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DO_A_TRADE"`
				DeliverManyPizzas struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DELIVER_MANY_PIZZAS"`
				ColorFancyHat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COLOR_FANCY_HAT"`
				HealWithVampire struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HEAL_WITH_VAMPIRE"`
				EarnMysticItem struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EARN_MYSTIC_ITEM"`
				CollectTrickleDown struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COLLECT_TRICKLE_DOWN"`
				FallInVoid struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FALL_IN_VOID"`
				CompleteKingQuest struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COMPLETE_KING_QUEST"`
				BidAuction struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BID_AUCTION"`
				UseSelfCheckout struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"USE_SELF_CHECKOUT"`
				KillsWithRambo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILLS_WITH_RAMBO"`
				XpFromKoth struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"XP_FROM_KOTH"`
				TopThreeInEvent struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TOP_THREE_IN_EVENT"`
				BakeCake struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BAKE_CAKE"`
				ManyKillsAsBeast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MANY_KILLS_AS_BEAST"`
				EnchantTierThree struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENCHANT_TIER_THREE"`
				ReachLevel100 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REACH_LEVEL_100"`
				FullStorage struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FULL_STORAGE"`
				ManyClicksDragonEgg struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MANY_CLICKS_DRAGON_EGG"`
				KillHighBounty struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_HIGH_BOUNTY"`
				HighTierDemon struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HIGH_TIER_DEMON"`
				Pentakill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PENTAKILL"`
				TopThreeRagePit struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TOP_THREE_RAGE_PIT"`
				HaveALotOfGold struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAVE_A_LOT_OF_GOLD"`
				SmallStreak struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SMALL_STREAK"`
				SpireLastFloor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPIRE_LAST_FLOOR"`
				GetHighBounty struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GET_HIGH_BOUNTY"`
				HighTierAngel struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HIGH_TIER_ANGEL"`
				GetEndlessQuiverArrows struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GET_ENDLESS_QUIVER_ARROWS"`
				SquadsManyPoints struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SQUADS_MANY_POINTS"`
				ClaimBounty struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CLAIM_BOUNTY"`
				ConsumeEverything struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CONSUME_EVERYTHING"`
				EatLotsOfCake struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EAT_LOTS_OF_CAKE"`
				WinRaffle struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_RAFFLE"`
				CompleteNightQuest struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COMPLETE_NIGHT_QUEST"`
				GetLuckyDiamond struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GET_LUCKY_DIAMOND"`
				PickupIngots struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PICKUP_INGOTS"`
				PlaceObsidian struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PLACE_OBSIDIAN"`
				KillWithStrength struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_WITH_STRENGTH"`
				MaxOutElGato struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_OUT_EL_GATO"`
				ReachPrestige20 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REACH_PRESTIGE_20"`
				MaxOutRenown struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_OUT_RENOWN"`
				UnlockFastPass struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNLOCK_FAST_PASS"`
				FastQuickMaths struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FAST_QUICK_MATHS"`
				TimeOnKotl struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TIME_ON_KOTL"`
				HaveCobblestone struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAVE_COBBLESTONE"`
				UseGoldenHeads struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"USE_GOLDEN_HEADS"`
				PurchaseScamArtist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PURCHASE_SCAM_ARTIST"`
				BountiesClaimedWithBh struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOUNTIES_CLAIMED_WITH_BH"`
				EnchantRare struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENCHANT_RARE"`
				CompleteContractFast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COMPLETE_CONTRACT_FAST"`
				RobberyCap struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ROBBERY_CAP"`
				CollectFromCarePackage struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COLLECT_FROM_CARE_PACKAGE"`
				EatRagePotatoes struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EAT_RAGE_POTATOES"`
				EarnAquaPants struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EARN_AQUA_PANTS"`
				MaxOutPassives struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_OUT_PASSIVES"`
				StayOnKoth struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"STAY_ON_KOTH"`
				HighKillstreak struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HIGH_KILLSTREAK"`
				KillTwoBeasts struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_TWO_BEASTS"`
			} `json:"one_time"`
			Tiered struct {
				Mysticism struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MYSTICISM"`
				Kills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS"`
				Prestiges struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"PRESTIGES"`
				Contracts struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CONTRACTS"`
				Renown struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"RENOWN"`
				Gold struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GOLD"`
				Events struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"EVENTS"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"pit"`
		Quake struct {
			OneTime struct {
				ShowMeTheMoney struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHOW_ME_THE_MONEY"`
				WhatJustHappened struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WHAT_JUST_HAPPENED"`
				ThanksGrandma struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THANKS_GRANDMA"`
				BakingADozen struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BAKING_A_DOZEN"`
				DoubleTrouble struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DOUBLE_TROUBLE"`
				DoublingUp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DOUBLING_UP"`
				Perfectionnist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PERFECTIONNIST"`
				Frog struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FROG"`
				OneInTheChamber struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ONE_IN_THE_CHAMBER"`
				PlatinumPlated struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PLATINUM_PLATED"`
				Rip struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RIP"`
				LookingFancy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LOOKING_FANCY"`
				Tubular struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TUBULAR"`
				BeyondIncredible struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BEYOND_INCREDIBLE"`
				Fly struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FLY"`
				FirstKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_KILL"`
				ILiekTurtles struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"I_LIEK_TURTLES"`
				TeamPlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEAM_PLAYER"`
				DarkSide struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DARK_SIDE"`
				Presidential struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PRESIDENTIAL"`
				Wat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WAT"`
				LoveIsInTheAir struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LOVE_IS_IN_THE_AIR"`
				MyPrecious struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MY_PRECIOUS"`
				Bogof struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOGOF"`
				Incredible struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"INCREDIBLE"`
				ThinkFast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THINK_FAST"`
				BuffingUp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BUFFING_UP"`
				WantAWardrobe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WANT_A_WARDROBE"`
				Redstoner struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REDSTONER"`
				MyWay struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MY_WAY"`
				GoodGuyGamer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOOD_GUY_GAMER"`
				Minigun struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MINIGUN"`
				FabulousWin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FABULOUS_WIN"`
				BillyTalent struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BILLY_TALENT"`
				NotToday struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_TODAY"`
				WhatHaveIDone struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WHAT_HAVE_I_DONE"`
				Humiliation struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HUMILIATION"`
				OhBabyATriple struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OH_BABY_A_TRIPLE"`
				AppleCorer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"APPLE_CORER"`
				Perfection struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PERFECTION"`
				Fanatic struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FANATIC"`
				SimpleThings struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SIMPLE_THINGS"`
				Untouchable struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNTOUCHABLE"`
				GoingUpInLife struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOING_UP_IN_LIFE"`
				ItsSoShiny struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ITS_SO_SHINY"`
				GrabItAll struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GRAB_IT_ALL"`
				NotWorking struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_WORKING"`
				HeavyShoulders struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HEAVY_SHOULDERS"`
				LookinGood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LOOKIN_GOOD"`
				DiamondsToYou struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DIAMONDS_TO_YOU"`
				CalmDown struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CALM_DOWN"`
				Bubbly struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BUBBLY"`
				HeartStopper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HEART_STOPPER"`
				HowDidThatHappen struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HOW_DID_THAT_HAPPEN"`
				OneShot struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ONE_SHOT"`
				Squish struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SQUISH"`
				InMyWay struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"IN_MY_WAY"`
			} `json:"one_time"`
			Tiered struct {
				Coins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COINS"`
				Headshots struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HEADSHOTS"`
				Wins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS"`
				WeaponArsenal struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WEAPON_ARSENAL"`
				KillingSprees struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLING_SPREES"`
				Godlikes struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GODLIKES"`
				Kills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"quake"`
		Skyblock struct {
			OneTime struct {
				ExplosiveEnding struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXPLOSIVE_ENDING"`
				BeaconatorTwo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BEACONATOR_TWO"`
				SaddleUp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SADDLE_UP"`
				IAmSuperior struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"I_AM_SUPERIOR"`
				GoblinSlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOBLIN_SLAYER"`
				LostSoul struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LOST_SOUL"`
				Baited struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BAITED"`
				AGoodSpiderIsADeadSpider struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"A_GOOD_SPIDER_IS_A_DEAD_SPIDER"`
				WaterSword struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WATER_SWORD"`
				EmptyFlowerPot struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EMPTY_FLOWER_POT"`
				Zookeeper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOOKEEPER"`
				SafetyFirst struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SAFETY_FIRST"`
				EndRace struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"END_RACE"`
				Mystical struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MYSTICAL"`
				GlassCannon struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GLASS_CANNON"`
				WorthIt struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WORTH_IT"`
				IAmGroot struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"I_AM_GROOT"`
				ProductionExpanded struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PRODUCTION_EXPANDED"`
				NextGeneration struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NEXT_GENERATION"`
				ToSpaceWeGo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TO_SPACE_WE_GO"`
				RoyalMeeting struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ROYAL_MEETING"`
				SupremeFarmer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUPREME_FARMER"`
				Jerry struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JERRY"`
				MasterEnchanter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MASTER_ENCHANTER"`
				BrainPower struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BRAIN_POWER"`
				BiggerStorageIsSeeded struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BIGGER_STORAGE_IS_SEEDED"`
				FlaminHot struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FLAMIN_HOT"`
				Shrimp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHRIMP"`
				WelcomeToMyFactory struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WELCOME_TO_MY_FACTORY"`
				CuteLittleCube struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CUTE_LITTLE_CUBE"`
				AccessoriesGalore struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ACCESSORIES_GALORE"`
				WonderfulTreasures struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WONDERFUL_TREASURES"`
				EternalFlameRing struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ETERNAL_FLAME_RING"`
				TheRealZooShady struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_REAL_ZOO_SHADY"`
				AnimalFishing struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ANIMAL_FISHING"`
				Flawless struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FLAWLESS"`
				Seriously struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SERIOUSLY"`
				TreasureFishing struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TREASURE_FISHING"`
				DefeatingDeath struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DEFEATING_DEATH"`
				ExpensiveBrew struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXPENSIVE_BREW"`
				Overkill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OVERKILL"`
				Agile struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"AGILE"`
				HeartOfTheEnd struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HEART_OF_THE_END"`
				YourAdventureBegins struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"YOUR_ADVENTURE_BEGINS"`
				StorageForever struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"STORAGE_FOREVER"`
				HelpfulHand struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HELPFUL_HAND"`
				Hsssss struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HSSSSS"`
				Speedrunner struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPEEDRUNNER"`
				HiddenSecrets struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HIDDEN_SECRETS"`
				SecondChance struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SECOND_CHANCE"`
				SmellLikeRoses struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SMELL_LIKE_ROSES"`
				TimeToStartFishing struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TIME_TO_START_FISHING"`
				ThoughChoice struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THOUGH_CHOICE"`
				Resourceful struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RESOURCEFUL"`
				Rebirth struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REBIRTH"`
				UpgradesPeopleUpgrades struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UPGRADES_PEOPLE_UPGRADES"`
				IndianaBones struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"INDIANA_BONES"`
				KingOfTheChicks struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KING_OF_THE_CHICKS"`
				DragonSlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DRAGON_SLAYER"`
				PeakOfTheMountain struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PEAK_OF_THE_MOUNTAIN"`
				TrueAlchemist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRUE_ALCHEMIST"`
				Rainbow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RAINBOW"`
				AGoodReview struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"A_GOOD_REVIEW"`
				ThisIsFair struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THIS_IS_FAIR"`
				TheFlintBros struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_FLINT_BROS"`
				Librarian struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LIBRARIAN"`
				MoreSpace struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MORE_SPACE"`
				CombinedEfforts struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COMBINED_EFFORTS"`
				RoyalResidentDialogue struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ROYAL_RESIDENT_DIALOGUE"`
				TheProdigy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_PRODIGY"`
				KnowledgeIsPower struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KNOWLEDGE_IS_POWER"`
				LegendaryRod struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_ROD"`
				Caretaker struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CARETAKER"`
				NightEyes struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NIGHT_EYES"`
				KingOfThePets struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KING_OF_THE_PETS"`
				FullyEvolved struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FULLY_EVOLVED"`
				TheOneBottle struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_ONE_BOTTLE"`
				AbsorbItAll struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ABSORB_IT_ALL"`
				GhostBuster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GHOST_BUSTER"`
				PromisedFulfilled struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PROMISED_FULFILLED"`
				IBelieveICanFly struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"I_BELIEVE_I_CAN_FLY"`
				SPlusSquad struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"S_PLUS_SQUAD"`
				TheFlash struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_FLASH"`
				BatPinata struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BAT_PINATA"`
				StubbornGiver struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"STUBBORN_GIVER"`
				PreciousMinerals struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PRECIOUS_MINERALS"`
				IntoTheDeep struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"INTO_THE_DEEP"`
				Fortunate struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FORTUNATE"`
				NextLevel struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NEXT_LEVEL"`
				AChallengingClimb struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"A_CHALLENGING_CLIMB"`
				Spiky struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPIKY"`
				SpeedOfLight struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPEED_OF_LIGHT"`
				HigherEnchants struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HIGHER_ENCHANTS"`
				ThreeBirdsOneArrow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THREE_BIRDS_ONE_ARROW"`
				EveryLittleBitHelps struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EVERY_LITTLE_BIT_HELPS"`
				ImFastAsHeckBoy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"IM_FAST_AS_HECK_BOY"`
				CaughtTheGrinch struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CAUGHT_THE_GRINCH"`
				IKnewIt struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"I_KNEW_IT"`
				RoughDeal struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ROUGH_DEAL"`
				DoYouEvenVoodoo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DO_YOU_EVEN_VOODOO"`
				SuperFuel struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUPER_FUEL"`
				QuestComplete struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"QUEST_COMPLETE"`
				Nightmare struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NIGHTMARE"`
				FallenStarCult struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FALLEN_STAR_CULT"`
				FriendForLife struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FRIEND_FOR_LIFE"`
				FriarLawrence struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FRIAR_LAWRENCE"`
				SweetTooth struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SWEET_TOOTH"`
				UnitedInBlood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNITED_IN_BLOOD"`
				InfiniteDarkness struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"INFINITE_DARKNESS"`
				BigGameFisher struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BIG_GAME_FISHER"`
				FrozenMonster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FROZEN_MONSTER"`
				AdvancedTransportation struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ADVANCED_TRANSPORTATION"`
				TimeToGoOnVacation struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TIME_TO_GO_ON_VACATION"`
				SiriusBusiness struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SIRIUS_BUSINESS"`
				Arcadia struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ARCADIA"`
				Explorer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXPLORER"`
				SeaMonster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SEA_MONSTER"`
				LifelongContract struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LIFELONG_CONTRACT"`
				DeathFromAbove struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DEATH_FROM_ABOVE"`
				DungeonExplorer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DUNGEON_EXPLORER"`
				DeepStorage struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DEEP_STORAGE"`
				SacrificesMustBeMade struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SACRIFICES_MUST_BE_MADE"`
				WatchMeShine struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WATCH_ME_SHINE"`
				NoEnchantsNeeded struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_ENCHANTS_NEEDED"`
				KingOfTheSea struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KING_OF_THE_SEA"`
				GottaGoFast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOTTA_GO_FAST"`
				TrueAdventurer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRUE_ADVENTURER"`
				OhShiny struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OH_SHINY"`
				Businessman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BUSINESSMAN"`
				Dullahan struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DULLAHAN"`
				MassProduction struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MASS_PRODUCTION"`
				HappyNewYear struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAPPY_NEW_YEAR"`
				DragonsEgg struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DRAGONS_EGG"`
				Gottagofast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOTTAGOFAST"`
				SoulHunter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SOUL_HUNTER"`
				ExistentialRevelations struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXISTENTIAL_REVELATIONS"`
				ItNeverEnds struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"IT_NEVER_ENDS"`
				HappyHolidays struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAPPY_HOLIDAYS"`
				HigherThanARabbit struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HIGHER_THAN_A_RABBIT"`
				YourBigBreak struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"YOUR_BIG_BREAK"`
			} `json:"one_time"`
			Tiered struct {
				Angler struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ANGLER"`
				TreasureHunter struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TREASURE_HUNTER"`
				Excavator struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"EXCAVATOR"`
				Dungeoneer struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"DUNGEONEER"`
				Harvester struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HARVESTER"`
				Concoctor struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CONCOCTOR"`
				Domesticator struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"DOMESTICATOR"`
				Combat struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COMBAT"`
				Augmentation struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"AUGMENTATION"`
				Gatherer struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GATHERER"`
				GoblinKiller struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GOBLIN_KILLER"`
				Treasury struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TREASURY"`
				Slayer struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SLAYER"`
				UniqueGifts struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"UNIQUE_GIFTS"`
				MinionLover struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MINION_LOVER"`
				HardWorkingMiner struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HARD_WORKING_MINER"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"skyblock"`
		Skyclash struct {
			OneTime struct {
				WhatsNext struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"WHATS_NEXT"`
				BadKarma struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"BAD_KARMA"`
				Addicted struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ADDICTED"`
				HalfHealthChallenge struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"HALF_HEALTH_CHALLENGE"`
				Powerspike struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"POWERSPIKE"`
				AnvilKill struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ANVIL_KILL"`
				MaxPower struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"MAX_POWER"`
				HappyLittleAchievement struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"HAPPY_LITTLE_ACHIEVEMENT"`
				AreYouCrazy struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ARE_YOU_CRAZY"`
				VoidKiller struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"VOID_KILLER"`
				GoldenOne struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"GOLDEN_ONE"`
				NoChest struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"NO_CHEST"`
				Flying struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"FLYING"`
				NewPassive struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"NEW_PASSIVE"`
				SharpJustice struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"SHARP_JUSTICE"`
				SkyclashExpert struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"SKYCLASH_EXPERT"`
				UhcChallenge struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"UHC_CHALLENGE"`
				HeroicCharge struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"HEROIC_CHARGE"`
				NoAlchemy struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"NO_ALCHEMY"`
				Boo struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"BOO"`
				Assassination struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ASSASSINATION"`
				UltimateWarriorChallenge struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ULTIMATE_WARRIOR_CHALLENGE"`
				MyPlaystyle struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"MY_PLAYSTYLE"`
				BomberChallenge struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"BOMBER_CHALLENGE"`
				ArcherChallenge struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ARCHER_CHALLENGE"`
				Stay struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"STAY"`
				KawaiiPants struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"KAWAII_PANTS"`
				MissingApple struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"MISSING_APPLE"`
				DreamTeam struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"DREAM_TEAM"`
				KillSecured struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"KILL_SECURED"`
				Rusher struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"RUSHER"`
				Greatman struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"GREATMAN"`
				ImABetterWizard struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"IM_A_BETTER_WIZARD"`
				GodOfMonsters struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"GOD_OF_MONSTERS"`
				TrainedWell struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"TRAINED_WELL"`
				SkymonGo struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"SKYMON_GO"`
				RestInPepperoni struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"REST_IN_PEPPERONI"`
				EndlessQuiver struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ENDLESS_QUIVER"`
				Lightweight struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"LIGHTWEIGHT"`
				ImAWizard struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"IM_A_WIZARD"`
			} `json:"one_time"`
			Tiered struct {
				CardsUnlocked struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CARDS_UNLOCKED"`
				PacksOpened struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"PACKS_OPENED"`
				Wins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS"`
				Kills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS"`
				TreasureHunter struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TREASURE_HUNTER"`
				MobBeheading struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MOB_BEHEADING"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"skyclash"`
		Skywars struct {
			OneTime struct {
				Peacemaker struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PEACEMAKER"`
				Siege struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SIEGE"`
				ChallengeUhc struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHALLENGE_UHC"`
				Donator struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DONATOR"`
				ChallengeRookie struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHALLENGE_ROOKIE"`
				ChallengeMaster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHALLENGE_MASTER"`
				GoingHam struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOING_HAM"`
				ShinyStuff struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHINY_STUFF"`
				ChallengeNoBlock struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHALLENGE_NO_BLOCK"`
				Mythical struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MYTHICAL"`
				Baller struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BALLER"`
				TeamworkMakesTheDreamWork struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEAMWORK_MAKES_THE_DREAM_WORK"`
				TheAngelsJourney struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_ANGELS_JOURNEY"`
				AttentionSeeking struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ATTENTION_SEEKING"`
				NowImEnchanted struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOW_IM_ENCHANTED"`
				Enderdragon struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENDERDRAGON"`
				ChallengeUltimateWarrior struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHALLENGE_ULTIMATE_WARRIOR"`
				TouchOfDeath struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TOUCH_OF_DEATH"`
				Trolol struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TROLOL"`
				Gapple struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GAPPLE"`
				WellWell struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WELL_WELL"`
				MapSelect struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAP_SELECT"`
				Rng struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RNG"`
				NoChestChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_CHEST_CHALLENGE"`
				GoneFishing struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GONE_FISHING"`
				FastAndFurious struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FAST_AND_FURIOUS"`
				Gotcha struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOTCHA"`
				FistsOfFury struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FISTS_OF_FURY"`
				Legendary struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY"`
				MaxWell struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_WELL"`
				SpeedRunner struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPEED_RUNNER"`
				SpeedRun struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPEED_RUN"`
				MysteryMob struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MYSTERY_MOB"`
				ChallengeNoChest struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHALLENGE_NO_CHEST"`
				KitConoisseur struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KIT_CONOISSEUR"`
				Teamwork struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEAMWORK"`
				HappyMeal struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAPPY_MEAL"`
				Killstolen struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILLSTOLEN"`
				MobSpawner struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MOB_SPAWNER"`
				AshesToAshes struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ASHES_TO_ASHES"`
				MaxPerk struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_PERK"`
				KillStreak struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_STREAK"`
				LuckyCharm struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LUCKY_CHARM"`
				SoloWarrior struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SOLO_WARRIOR"`
				WhoNeedsTeammates struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WHO_NEEDS_TEAMMATES"`
				LuckySouls struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LUCKY_SOULS"`
				MegaWarrior struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MEGA_WARRIOR"`
				NickCage struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NICK_CAGE"`
				CorruptionLord struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CORRUPTION_LORD"`
				FearMe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FEAR_ME"`
				SlowSteady struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SLOW_STEADY"`
				PlayingItSafe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PLAYING_IT_SAFE"`
				ChallengePaper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHALLENGE_PAPER"`
				OpenChest struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OPEN_CHEST"`
				Criminal struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CRIMINAL"`
				HastaLaVista struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HASTA_LA_VISTA"`
				Sniper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SNIPER"`
				ChallengeArcher struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHALLENGE_ARCHER"`
				ChallengeHalfHealth struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHALLENGE_HALF_HEALTH"`
				PortalGame struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PORTAL_GAME"`
				ChallengePro struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHALLENGE_PRO"`
				WellDeserved struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WELL_DESERVED"`
			} `json:"one_time"`
			Tiered struct {
				KillsTeam struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS_TEAM"`
				KillsSolo struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS_SOLO"`
				KitsTeam struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KITS_TEAM"`
				KitsSolo struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KITS_SOLO"`
				KitsMega struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KITS_MEGA"`
				OpalObsession struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"OPAL_OBSESSION"`
				WinsTeam struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS_TEAM"`
				WinsLab struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS_LAB"`
				WinsSolo struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS_SOLO"`
				WinsMega struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS_MEGA"`
				Cages struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CAGES"`
				NewDayNewChallenge struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"NEW_DAY_NEW_CHALLENGE"`
				KillsMega struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS_MEGA"`
				Heads struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HEADS"`
				YouReAStar struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"YOU_RE_A_STAR"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"skywars"`
		Speeduhc struct {
			OneTime struct {
				MasteryUnlock struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"MASTERY_UNLOCK"`
				EnchantedBook struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ENCHANTED_BOOK"`
				PotionBrewer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"POTION_BREWER"`
				RidePig struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RIDE_PIG"`
				MasterOfMasters struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MASTER_OF_MASTERS"`
				UseEnderpearl struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"USE_ENDERPEARL"`
				GoldenApple struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOLDEN_APPLE"`
				HotHog struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HOT_HOG"`
				NaturalTalent struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NATURAL_TALENT"`
				Yeehaw struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"YEEHAW"`
				EnchantedArmor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENCHANTED_ARMOR"`
				BraveNewWorld struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BRAVE_NEW_WORLD"`
				SkeletonBow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SKELETON_BOW"`
				SnipePlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SNIPE_PLAYER"`
				KitSpecialist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KIT_SPECIALIST"`
				KitUnlock struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KIT_UNLOCK"`
				KillTnt struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_TNT"`
				Diamonds struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DIAMONDS"`
				MelonSmasher struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MELON_SMASHER"`
				GodApple struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOD_APPLE"`
				TearsOfLoneliness struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEARS_OF_LONELINESS"`
				PaperCut struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PAPER_CUT"`
				Rusher struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RUSHER"`
				Domination struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DOMINATION"`
				MyWay struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MY_WAY"`
				KillGhast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_GHAST"`
			} `json:"one_time"`
			Tiered struct {
				Salty struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SALTY"`
				Hunter struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HUNTER"`
				Collector struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COLLECTOR"`
				Promotion struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"PROMOTION"`
				UhcMaster struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"UHC_MASTER"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"speeduhc"`
		Summer struct {
			OneTime struct {
				HomeRun struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HOME_RUN"`
				CollectorsEdition struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Secret      bool   `json:"secret"`
					Legacy      bool   `json:"legacy"`
				} `json:"COLLECTORS_EDITION"`
				Shockwave struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHOCKWAVE"`
				BringTheStorm struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BRING_THE_STORM"`
				ThisIsntTheLobby struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THIS_ISNT_THE_LOBBY"`
				LightningRod struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LIGHTNING_ROD"`
				WrongSeason struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WRONG_SEASON"`
				OasisJake struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OASIS_JAKE"`
				StayHydrated struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"STAY_HYDRATED"`
				ExpertDiver struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"EXPERT_DIVER"`
				PoolParty struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"POOL_PARTY"`
				CoolOff struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COOL_OFF"`
				SunDenial struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUN_DENIAL"`
				RuleTheHeat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RULE_THE_HEAT"`
				OceanInABucket struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OCEAN_IN_A_BUCKET"`
				Grillmaster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GRILLMASTER"`
				BbqInTheSky struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BBQ_IN_THE_SKY"`
				Sunburnt struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUNBURNT"`
				BeatTheHeat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BEAT_THE_HEAT"`
				WinPinkSheep struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_PINK_SHEEP"`
				AustralianWinter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"AUSTRALIAN_WINTER"`
				SeasideShootout struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SEASIDE_SHOOTOUT"`
				Drought struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DROUGHT"`
				OutFishing struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OUT_FISHING"`
				PunchOut struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PUNCH_OUT"`
				KindRefreshment struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KIND_REFRESHMENT"`
				IceCold struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ICE_COLD"`
			} `json:"one_time"`
			Tiered struct {
				GoneFishing struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GONE_FISHING"`
				TreasureHoarder struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TREASURE_HOARDER"`
				TreasureMaster struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TREASURE_MASTER"`
				Shopaholic struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SHOPAHOLIC"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"summer"`
		Supersmash struct {
			OneTime struct {
				GencluckChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GENCLUCK_CHALLENGE"`
				PugChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PUG_CHALLENGE"`
				YoungApprentice struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"YOUNG_APPRENTICE"`
				ShoopChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHOOP_CHALLENGE"`
				ReachStars struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REACH_STARS"`
				ReachSky struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REACH_SKY"`
				TinmanChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TINMAN_CHALLENGE"`
				SpoodermanChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPOODERMAN_CHALLENGE"`
				BotmonVsSpooderman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOTMON_VS_SPOODERMAN"`
				FrostmageChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FROSTMAGE_CHALLENGE"`
				Domination struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DOMINATION"`
				MarauderChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MARAUDER_CHALLENGE"`
				TwoForOne struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TWO_FOR_ONE"`
				SecondRound struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SECOND_ROUND"`
				KarakotChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KARAKOT_CHALLENGE"`
				BotmonChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOTMON_CHALLENGE"`
				BulkChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BULK_CHALLENGE"`
				TooEasy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TOO_EASY"`
				VoidChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VOID_CHALLENGE"`
				GetOverHere struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GET_OVER_HERE"`
				CakeChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CAKE_CHALLENGE"`
				Supremacy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUPREMACY"`
				NeedAll struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NEED_ALL"`
				ExperienceExpress struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXPERIENCE_EXPRESS"`
				Teamwork struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEAMWORK"`
				SkullfireChallenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SKULLFIRE_CHALLENGE"`
				ThePeak struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THE_PEAK"`
			} `json:"one_time"`
			Tiered struct {
				SmashChampion struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SMASH_CHAMPION"`
				SmashWinner struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SMASH_WINNER"`
				HeroSlayer struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HERO_SLAYER"`
				Handyman struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HANDYMAN"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"supersmash"`
		Tntgames struct {
			OneTime struct {
				TntTagDifferenttags struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_TAG_DIFFERENTTAGS"`
				ProSurfer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PRO_SURFER"`
				WizardsNogood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_NOGOOD"`
				TntNpcs struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_NPCS"`
				DontGetClose struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DONT_GET_CLOSE"`
				SpleefRepulsor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPLEEF_REPULSOR"`
				WizardsFireWizardExplode struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_FIRE_WIZARD_EXPLODE"`
				SpleefHits struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPLEEF_HITS"`
				NoYou struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_YOU"`
				SpleefNoperks struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPLEEF_NOPERKS"`
				WizardsHurts struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_HURTS"`
				WinStreak struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_STREAK"`
				Spooky struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPOOKY"`
				Freeze struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FREEZE"`
				JackOfAll struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JACK_OF_ALL"`
				BowSpleefRepulsorUpgrade struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOW_SPLEEF_REPULSOR_UPGRADE"`
				PvpRunFlying struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PVP_RUN_FLYING"`
				WizardsTriple struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_TRIPLE"`
				RunPotions struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RUN_POTIONS"`
				WizardsThirtyKills struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_THIRTY_KILLS"`
				TntWinsInARow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_WINS_IN_A_ROW"`
				TntTagBlownup struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_TAG_BLOWNUP"`
				TntTagDangerous struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_TAG_DANGEROUS"`
				PvpRunTriple struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PVP_RUN_TRIPLE"`
				TntRunNodjs struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_RUN_NODJS"`
				ExtrmeSpeed struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXTRME_SPEED"`
				PvpRunNodjs struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PVP_RUN_NODJS"`
				WizardsBloodWizardRegen struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_BLOOD_WIZARD_REGEN"`
				TntTagSnowball struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_TAG_SNOWBALL"`
				TimeWizards struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TIME_WIZARDS"`
				TntTagTraveller struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_TAG_TRAVELLER"`
				TntTagAww struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_TAG_AWW"`
				BowSpleefTripleShotUpgrade struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOW_SPLEEF_TRIPLE_SHOT_UPGRADE"`
				Patience struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PATIENCE"`
				WizardsJumper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_JUMPER"`
				TntRunNomove struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_RUN_NOMOVE"`
				SpleefTriple struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPLEEF_TRIPLE"`
				BowSpleefDoubleJumpUpgrade struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOW_SPLEEF_DOUBLE_JUMP_UPGRADE"`
				SpleefShots struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPLEEF_SHOTS"`
				TntTagTagger struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_TAG_TAGGER"`
				WizardsCaphelp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_CAPHELP"`
				ThatWasClose struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THAT_WAS_CLOSE"`
				Lucky struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LUCKY"`
				WizardsAssistant struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_ASSISTANT"`
				ByeBye struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BYE_BYE"`
				PotionShower struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"POTION_SHOWER"`
				SpleefNoshots struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPLEEF_NOSHOTS"`
				TntRunFirstlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_RUN_FIRSTLAYER"`
				TntRunFirstWin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_RUN_FIRST_WIN"`
				TntRunNoSprinting struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_RUN_NO_SPRINTING"`
				SpleefExhausted struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPLEEF_EXHAUSTED"`
				Kangaroo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KANGAROO"`
				PvpRunManykills struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PVP_RUN_MANYKILLS"`
				Healer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HEALER"`
				TntPrestige struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_PRESTIGE"`
				TntRunShort struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_RUN_SHORT"`
				TntRunEffects struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_RUN_EFFECTS"`
				PvpRunPacifist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PVP_RUN_PACIFIST"`
				WizardsLeaderboard struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_LEADERBOARD"`
				FlyingMadman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FLYING_MADMAN"`
				TntTagFirstWin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_TAG_FIRST_WIN"`
				MagneticBoots struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAGNETIC_BOOTS"`
				TntParkour struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_PARKOUR"`
				PvpRunSabotage struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PVP_RUN_SABOTAGE"`
				OverpoweredArchery struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OVERPOWERED_ARCHERY"`
				Combo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COMBO"`
				PvpRunNohit struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PVP_RUN_NOHIT"`
				PvpRunFirstWin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PVP_RUN_FIRST_WIN"`
				TntTagClose struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_TAG_CLOSE"`
				WizardsHeal struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_HEAL"`
				WizardsAwizard struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_AWIZARD"`
				GrandMaster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GRAND_MASTER"`
				PvpRunHalfhearted struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PVP_RUN_HALFHEARTED"`
				PvpRunPotions struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PVP_RUN_POTIONS"`
				TntRunFlying struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_RUN_FLYING"`
				MegaTank struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MEGA_TANK"`
				WizardsLead struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_LEAD"`
				TntRunPurchasePotion struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_RUN_PURCHASE_POTION"`
				BowSpleefFirstDoubleJump struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOW_SPLEEF_FIRST_DOUBLE_JUMP"`
				Timer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TIMER"`
				WizardsFirstWin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIZARDS_FIRST_WIN"`
				Sniper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SNIPER"`
				TntTagDm struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TNT_TAG_DM"`
				PvpRunFists struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PVP_RUN_FISTS"`
				BowSpleefFirstWin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOW_SPLEEF_FIRST_WIN"`
				SecondChance struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SECOND_CHANCE"`
			} `json:"one_time"`
			Tiered struct {
				PvpRunWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"PVP_RUN_WINS"`
				TntWizardsCaps struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TNT_WIZARDS_CAPS"`
				TntTriathlon struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TNT_TRIATHLON"`
				BowSpleefWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BOW_SPLEEF_WINS"`
				TntRunWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TNT_RUN_WINS"`
				WizardsWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WIZARDS_WINS"`
				TntWizardsKills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TNT_WIZARDS_KILLS"`
				Clinic struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CLINIC"`
				BlockRunner struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BLOCK_RUNNER"`
				TntTagWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TNT_TAG_WINS"`
				PvpRunKiller struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"PVP_RUN_KILLER"`
				TntBanker struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TNT_BANKER"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"tntgames"`
		Truecombat struct {
			OneTime struct {
				GoldForager struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"GOLD_FORAGER"`
				Doctor struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"DOCTOR"`
				NotUhc struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"NOT_UHC"`
				Knight struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"KNIGHT"`
				RedstoneDealer struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"REDSTONE_DEALER"`
				GonePeaceful struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"GONE_PEACEFUL"`
				Multikill struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"MULTIKILL"`
				DiamondArmor struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"DIAMOND_ARMOR"`
				CrossFingers struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"CROSS_FINGERS"`
				DeadlyDonation struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"DEADLY_DONATION"`
				Invincible struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"INVINCIBLE"`
				NotToday struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"NOT_TODAY"`
				Rampage struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"RAMPAGE"`
				Legendary struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"LEGENDARY"`
				Boom struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"BOOM"`
				GoldenBounty struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"GOLDEN_BOUNTY"`
				IFna struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"I_FNA"`
				Dominating struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"DOMINATING"`
				MonsterHunter struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"MONSTER_HUNTER"`
				Strategist struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"STRATEGIST"`
				Monsterkill struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"MONSTERKILL"`
				DestinyCalls struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"DESTINY_CALLS"`
				StickySituation struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"STICKY_SITUATION"`
				Toxic struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"TOXIC"`
				ExplosiveArcher struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"EXPLOSIVE_ARCHER"`
				Chance struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"CHANCE"`
				EnderDragon struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"ENDER_DRAGON"`
				GoldDigger struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"GOLD_DIGGER"`
				DragonSlayer struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"DRAGON_SLAYER"`
				MaxedPerk struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"MAXED_PERK"`
				SharpenedKatana struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"SHARPENED_KATANA"`
				ThisIsTooEasy struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"THIS_IS_TOO_EASY"`
				Rushmode struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"RUSHMODE"`
				WhatHaveIDone struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"WHAT_HAVE_I_DONE"`
				LightningFast struct {
					Points      int    `json:"points"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
				} `json:"LIGHTNING_FAST"`
			} `json:"one_time"`
			Tiered struct {
				SoloKiller struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SOLO_KILLER"`
				KitHoarderSolo struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KIT_HOARDER_SOLO"`
				SoloWinner struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SOLO_WINNER"`
				FeelsLucky struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"FEELS_LUCKY"`
				King struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KING"`
				KitHoarderTeam struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KIT_HOARDER_TEAM"`
				TeamWinner struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TEAM_WINNER"`
				TeamKiller struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Legacy      bool   `json:"legacy"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TEAM_KILLER"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"truecombat"`
		Uhc struct {
			OneTime struct {
				CraftingRevolution struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CRAFTING_REVOLUTION"`
				Bloodthirsty struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLOODTHIRSTY"`
				FearedHunters struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FEARED_HUNTERS"`
				CraftingManiac struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CRAFTING_MANIAC"`
				KitMastery struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KIT_MASTERY"`
				Obliterate struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"OBLITERATE"`
				TreasureOfElDorado struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TREASURE_OF_EL_DORADO"`
				DeathsScythe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DEATHS_SCYTHE"`
				ProductionEnder struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PRODUCTION_ENDER"`
				DeuxExMachina struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DEUX_EX_MACHINA"`
				CrazyHardcore struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CRAZY_HARDCORE"`
				WealthyChampion struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WEALTHY_CHAMPION"`
				UltimatelyWealthy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ULTIMATELY_WEALTHY"`
				ModifierMaster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MODIFIER_MASTER"`
				AxeOfPerun struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"AXE_OF_PERUN"`
				ByeCruelWorld struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BYE_CRUEL_WORLD"`
				Prestigious struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PRESTIGIOUS"`
				OnePoundSlap struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					Secret                bool    `json:"secret"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ONE_POUND_SLAP"`
				DragonWarrior struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DRAGON_WARRIOR"`
				Seafood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SEAFOOD"`
				ShinyRock struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHINY_ROCK"`
				GhostRider struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GHOST_RIDER"`
				TabletsOfDestiny struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TABLETS_OF_DESTINY"`
				TotalDomination struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TOTAL_DOMINATION"`
				UhcExpert struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UHC_EXPERT"`
				Rampage struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RAMPAGE"`
				ChestOfFate struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHEST_OF_FATE"`
				ExtraPowerful struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXTRA_POWERFUL"`
				Adrenaline struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ADRENALINE"`
				WarmingUp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WARMING_UP"`
				AceKit struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ACE_KIT"`
				Puppy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PUPPY"`
				Powerhouse struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"POWERHOUSE"`
				RideAHorse struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RIDE_A_HORSE"`
				Bomberman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOMBERMAN"`
				ElitePrestige struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ELITE_PRESTIGE"`
				Eldorado struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ELDORADO"`
				Music struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MUSIC"`
				RobinHood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ROBIN_HOOD"`
				NothingCanStopUs struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOTHING_CAN_STOP_US"`
				ParkourMaster struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PARKOUR_MASTER"`
				Enderkind struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENDERKIND"`
				CallingUponThor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CALLING_UPON_THOR"`
				MasterMagician struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MASTER_MAGICIAN"`
				Exodus struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXODUS"`
				NoProblemHere struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_PROBLEM_HERE"`
				UltimateRecipe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ULTIMATE_RECIPE"`
				WrongMode struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WRONG_MODE"`
				ChampionRank struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHAMPION_RANK"`
				UltimateCrafting struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ULTIMATE_CRAFTING"`
				AccidentalSummoning struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ACCIDENTAL_SUMMONING"`
				Ghast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GHAST"`
				DrinkWithCaution struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DRINK_WITH_CAUTION"`
				Dong struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DONG"`
				DiceOfGod struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DICE_OF_GOD"`
			} `json:"one_time"`
			Tiered struct {
				Hunter struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"HUNTER"`
				MovingUp struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MOVING_UP"`
				Champion struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CHAMPION"`
				Consumer struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CONSUMER"`
				Bounty struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"BOUNTY"`
				Ultimatum struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ULTIMATUM"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"uhc"`
		Vampirez struct {
			OneTime struct {
				PurchaseBlood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PURCHASE_BLOOD"`
				Nightmare struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NIGHTMARE"`
				BloodThirsty struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLOOD_THIRSTY"`
				PurchaseGold struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PURCHASE_GOLD"`
				Pest struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PEST"`
				Blood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLOOD"`
				PurchaseFood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PURCHASE_FOOD"`
				SoleSurvivor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SOLE_SURVIVOR"`
				Robbed struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ROBBED"`
				VampireFangKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VAMPIRE_FANG_KILL"`
				VampireKillsOneRound struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VAMPIRE_KILLS_ONE_ROUND"`
				PurchaseArmor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PURCHASE_ARMOR"`
				DontNeedIt struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DONT_NEED_IT"`
				PurchaseSword struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PURCHASE_SWORD"`
				VampireShop struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VAMPIRE_SHOP"`
				Potions struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"POTIONS"`
				KillZombies struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_ZOMBIES"`
				Upgraded struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UPGRADED"`
				Undefeatable struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNDEFEATABLE"`
				WordPuns struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WORD_PUNS"`
				Purified struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PURIFIED"`
				ChestHunter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHEST_HUNTER"`
				Gold struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOLD"`
				LastChance struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LAST_CHANCE"`
				FirstWaveKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_WAVE_KILL"`
				TastesFunny struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TASTES_FUNNY"`
				ZombieSlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIE_SLAYER"`
				SurvivorKillsOneRound struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SURVIVOR_KILLS_ONE_ROUND"`
				ZombieWhisperer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ZOMBIE_WHISPERER"`
			} `json:"one_time"`
			Tiered struct {
				ZombieKiller struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ZOMBIE_KILLER"`
				Coins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COINS"`
				KillVampires struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILL_VAMPIRES"`
				KillSurvivors struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILL_SURVIVORS"`
				SurvivorWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SURVIVOR_WINS"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"vampirez"`
		Walls struct {
			OneTime struct {
				FirstKit struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_KIT"`
				GetDiamondSword struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GET_DIAMOND_SWORD"`
				CraftFlint struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CRAFT_FLINT"`
				TruePower struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRUE_POWER"`
				FindingNemo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FINDING_NEMO"`
				NoTeamDeaths struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NO_TEAM_DEATHS"`
				CatchFish struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CATCH_FISH"`
				KillCliff struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"KILL_CLIFF"`
				Engineer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ENGINEER"`
				WipedOut struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIPED_OUT"`
				Burning struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BURNING"`
				Revenge struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REVENGE"`
				NotAFish struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_A_FISH"`
				Customized struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CUSTOMIZED"`
				CreatePortal struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CREATE_PORTAL"`
				RideHorse struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RIDE_HORSE"`
				Untouchable struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNTOUCHABLE"`
				CraftBoat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CRAFT_BOAT"`
				TenKills struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEN_KILLS"`
				SoleSurvivor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SOLE_SURVIVOR"`
				Vampirism struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VAMPIRISM"`
				FirstPerk struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_PERK"`
				Berserk struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BERSERK"`
				RobinHood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ROBIN_HOOD"`
				NotToday struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_TODAY"`
				StarterKits struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"STARTER_KITS"`
				RecordLabel struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RECORD_LABEL"`
			} `json:"one_time"`
			Tiered struct {
				DiamondMiner struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"DIAMOND_MINER"`
				Kills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS"`
				Coins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COINS"`
				Wins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"walls"`
		Walls3 struct {
			OneTime struct {
				WhatsTheBigIdea struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WHATS_THE_BIG_IDEA"`
				ReadySetBoom struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"READY_SET_BOOM"`
				LegendaryAutomaton struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_AUTOMATON"`
				WolfSmash struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WOLF_SMASH"`
				HeavyEater struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HEAVY_EATER"`
				GoneVegan struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GONE_VEGAN"`
				GraveRobber struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GRAVE_ROBBER"`
				BullShark struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BULL_SHARK"`
				ThrowingCoconuts struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THROWING_COCONUTS"`
				BirdsEye struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BIRDS_EYE"`
				TreasureHunter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TREASURE_HUNTER"`
				Emp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EMP"`
				Recycling struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RECYCLING"`
				HammerHead struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAMMER_HEAD"`
				LegendarySquid struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_SQUID"`
				LegendaryHunter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_HUNTER"`
				SaveYourStuff struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SAVE_YOUR_STUFF"`
				Alotv1 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ALOTV1"`
				IFeelSick struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"I_FEEL_SICK"`
				FirstSkillUpgrade struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_SKILL_UPGRADE"`
				TigerShark struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TIGER_SHARK"`
				ComingThrough struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COMING_THROUGH"`
				Moodsetter struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MOODSETTER"`
				BlazeParty struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLAZE_PARTY"`
				RemoteDetonation struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REMOTE_DETONATION"`
				TimeToDiet struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TIME_TO_DIET"`
				CakeHunter2 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CAKE_HUNTER_2"`
				FailedExperiment struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FAILED_EXPERIMENT"`
				ToInfinity struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TO_INFINITY"`
				LegendaryPirate struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_PIRATE"`
				HuntingSeason struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HUNTING_SEASON"`
				CircleOfTrust struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CIRCLE_OF_TRUST"`
				LegendarySkeleton struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_SKELETON"`
				SneakAttack struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SNEAK_ATTACK"`
				BloodLust struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLOOD_LUST"`
				LegendarySpider struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_SPIDER"`
				LegendaryWerewolf struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_WEREWOLF"`
				Morra struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MORRA"`
				IAmCow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"I_AM_COW"`
				Sleepytime struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SLEEPYTIME"`
				LegendaryEnderman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_ENDERMAN"`
				LegendaryPhoenix struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_PHOENIX"`
				MuchDogs struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MUCH_DOGS"`
				Esc struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ESC"`
				CaptainCombo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CAPTAIN_COMBO"`
				Marksman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MARKSMAN"`
				LegendaryShark struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_SHARK"`
				LaserPrecision struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LASER_PRECISION"`
				HammerDown struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAMMER_DOWN"`
				Collector struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COLLECTOR"`
				Skitterama struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SKITTERAMA"`
				Surprise struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SURPRISE"`
				PotionsOfDeath struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"POTIONS_OF_DEATH"`
				LegendaryAssassin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_ASSASSIN"`
				LegendaryZombie struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_ZOMBIE"`
				LegendaryCreeper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_CREEPER"`
				HardAsSteel struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HARD_AS_STEEL"`
				MineDiamond struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MINE_DIAMOND"`
				Peacekreeper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PEACEKREEPER"`
				TrustMeImADoctor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TRUST_ME_IM_A_DOCTOR"`
				OnPoint struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ON_POINT"`
				MultiKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MULTI_KILL"`
				Yeehaw struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"YEEHAW"`
				Wunderbar struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WUNDERBAR"`
				Alchemy100 struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ALCHEMY_100"`
				LegendaryHerobrine struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_HEROBRINE"`
				ThisIsMyFinalForm struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THIS_IS_MY_FINAL_FORM"`
				BiologicalRestoration struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BIOLOGICAL_RESTORATION"`
				LegendaryDreadlord struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_DREADLORD"`
				BaBoom struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BA_BOOM"`
				InventoryManagement struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"INVENTORY_MANAGEMENT"`
				SpeedRun struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPEED_RUN"`
				BornTalented struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BORN_TALENTED"`
				FirstGatheringSkillUpgrade struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_GATHERING_SKILL_UPGRADE"`
				AttackWither struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ATTACK_WITHER"`
				LegendaryMoleman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_MOLEMAN"`
				YouShallNotPass struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"YOU_SHALL_NOT_PASS"`
				GreatWhite struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GREAT_WHITE"`
				ThanksConnor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THANKS_CONNOR"`
				MaxSkills struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_SKILLS"`
				Veteran struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"VETERAN"`
				ItsAllOgreNow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ITS_ALL_OGRE_NOW"`
				ChillSniper struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHILL_SNIPER"`
				FireInTheHole struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRE_IN_THE_HOLE"`
				SkeletonsBestFriend struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SKELETONS_BEST_FRIEND"`
				DeathFromAbove struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DEATH_FROM_ABOVE"`
				FeelsBad struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FEELS_BAD"`
				SpeedyMineman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SPEEDY_MINEMAN"`
				Competitive struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COMPETITIVE"`
				LuckySunny struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LUCKY_SUNNY"`
				AshesToBashes struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ASHES_TO_BASHES"`
				HugMe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HUG_ME"`
				SeasonsGreetings struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SEASONS_GREETINGS"`
				Magnate struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAGNATE"`
				Geronimo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GERONIMO"`
				BowDown struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOW_DOWN"`
				TeamPlayer struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TEAM_PLAYER"`
				LegendaryCow struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_COW"`
				Avalanche struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"AVALANCHE"`
				MaximumEffort struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAXIMUM_EFFORT"`
				RevengeOfTheWolves struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"REVENGE_OF_THE_WOLVES"`
				Rushlord struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RUSHLORD"`
				HighOnOres struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HIGH_ON_ORES"`
				FindChest struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIND_CHEST"`
				Whirlwind struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WHIRLWIND"`
				NightsRest struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NIGHTS_REST"`
				WinBeforeDeathmatch struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_BEFORE_DEATHMATCH"`
				Breadlord struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BREADLORD"`
				Whirlpool struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WHIRLPOOL"`
				Louis struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LOUIS"`
				FrostyFriendship struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FROSTY_FRIENDSHIP"`
				LegendaryBlaze struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_BLAZE"`
				DontBlink struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"DONT_BLINK"`
				Exchange struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXCHANGE"`
				Thunder struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THUNDER"`
				CruisingFlames struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CRUISING_FLAMES"`
				WinWithLivingWither struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WIN_WITH_LIVING_WITHER"`
				LegendaryShaman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_SHAMAN"`
				Timber struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TIMBER"`
				AtLeastHeTried struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"AT_LEAST_HE_TRIED"`
				Gotcha struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GOTCHA"`
				MaxRenderDistance struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_RENDER_DISTANCE"`
				Ninja7S struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NINJA_7S"`
				LegendaryGolem struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_GOLEM"`
				BlowingBubbles struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BLOWING_BUBBLES"`
				LegendaryRenegade struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_RENEGADE"`
				Masterpiece struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MASTERPIECE"`
				MaxZombieSkills struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_ZOMBIE_SKILLS"`
				Happy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAPPY"`
				LegendaryArcanist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_ARCANIST"`
				StayinAlive struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"STAYIN_ALIVE"`
				NotAGolem struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NOT_A_GOLEM"`
				SchoolCancelled struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SCHOOL_CANCELLED"`
				MassDestruction struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MASS_DESTRUCTION"`
				IntoTheFuture struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"INTO_THE_FUTURE"`
				LegendaryPigman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_PIGMAN"`
				TerminatedScript struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TERMINATED_SCRIPT"`
				Constructor struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CONSTRUCTOR"`
				MooBrawl struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MOO_BRAWL"`
				MaxSkeletonSkills struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_SKELETON_SKILLS"`
				TakingTheHeat struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TAKING_THE_HEAT"`
				AdvancedStrategies struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ADVANCED_STRATEGIES"`
				Untouchable struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNTOUCHABLE"`
				TimeToFeast struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TIME_TO_FEAST"`
				LegendarySnowman struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY_SNOWMAN"`
				MaxHerobrineSkills struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAX_HEROBRINE_SKILLS"`
			} `json:"one_time"`
			Tiered struct {
				Coins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COINS"`
				JackOfAllTrades struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"JACK_OF_ALL_TRADES"`
				Guardian struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"GUARDIAN"`
				Wins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WINS"`
				CakeHunterTiered struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CAKE_HUNTER_TIERED"`
				Kills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS"`
				Moctezuma struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MOCTEZUMA"`
				Rusher struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"RUSHER"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"walls3"`
		Warlords struct {
			OneTime struct {
				MediumRare struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MEDIUM_RARE"`
				ComingThrough struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COMING_THROUGH"`
				Supernatural struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUPERNATURAL"`
				SoloCarry struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SOLO_CARRY"`
				MyPleasure struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MY_PLEASURE"`
				Warlord struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WARLORD"`
				GreenIsLove struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GREEN_IS_LOVE"`
				MadSkillzYo struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAD_SKILLZ_YO"`
				HammerOfDestruction struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"HAMMER_OF_DESTRUCTION"`
				WorldTravel struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"WORLD_TRAVEL"`
				ExtensiveTraining struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EXTENSIVE_TRAINING"`
				EagleEyed struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"EAGLE_EYED"`
				ThatWasEasy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THAT_WAS_EASY"`
				CaptureTheWin struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CAPTURE_THE_WIN"`
				IHaveToConcentrate struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"I_HAVE_TO_CONCENTRATE"`
				JuicedUp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"JUICED_UP"`
				AreYouKiddingMe struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ARE_YOU_KIDDING_ME"`
				ShardsForGlory struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SHARDS_FOR_GLORY"`
				TotalDomination struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"TOTAL_DOMINATION"`
				Nifty struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"NIFTY"`
				MakinSomeRoom struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"MAKIN_SOME_ROOM"`
				Lifeleech struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LIFELEECH"`
				SlowDownThere struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SLOW_DOWN_THERE"`
				GiddyUp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"GIDDY_UP"`
				BeepBeep struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BEEP_BEEP"`
				AvengersWrath struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"AVENGERS_WRATH"`
				ThisBetterBeGood struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"THIS_BETTER_BE_GOOD"`
				RightOnTime struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"RIGHT_ON_TIME"`
				IGotYouBro struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"I_GOT_YOU_BRO"`
				PowerUp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"POWER_UP"`
				IMustResist struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"I_MUST_RESIST"`
				UndeadArmy struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"UNDEAD_ARMY"`
				BearingGifts struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BEARING_GIFTS"`
				CollateralDamage struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"COLLATERAL_DAMAGE"`
				FeelingSpecial struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FEELING_SPECIAL"`
				FirstOfMany struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"FIRST_OF_MANY"`
				PurplePower struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"PURPLE_POWER"`
				Legendary struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"LEGENDARY"`
				BootCamp struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"BOOT_CAMP"`
				ItShines struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"IT_SHINES"`
				OnTop struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"ON_TOP"`
				ChainKill struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"CHAIN_KILL"`
				SuperSoaker struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUPER_SOAKER"`
				SuperPowers struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SUPER_POWERS"`
				SurpriseAttack struct {
					Points                int     `json:"points"`
					Name                  string  `json:"name"`
					Description           string  `json:"description"`
					GamePercentUnlocked   float64 `json:"gamePercentUnlocked"`
					GlobalPercentUnlocked float64 `json:"globalPercentUnlocked"`
				} `json:"SURPRISE_ATTACK"`
			} `json:"one_time"`
			Tiered struct {
				Assist struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"ASSIST"`
				RepairWeapons struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"REPAIR_WEAPONS"`
				CtfWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CTF_WINS"`
				ShamanLevel struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"SHAMAN_LEVEL"`
				MageLevel struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"MAGE_LEVEL"`
				DomObjective struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"DOM_OBJECTIVE"`
				WarriorLevel struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"WARRIOR_LEVEL"`
				TdmWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"TDM_WINS"`
				Coins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"COINS"`
				Kills struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"KILLS"`
				CtfObjective struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"CTF_OBJECTIVE"`
				PaladinLevel struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"PALADIN_LEVEL"`
				DomWins struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Tiers       []struct {
						Tier   int `json:"tier"`
						Points int `json:"points"`
						Amount int `json:"amount"`
					} `json:"tiers"`
				} `json:"DOM_WINS"`
			} `json:"tiered"`
			TotalPoints       int `json:"total_points"`
			TotalLegacyPoints int `json:"total_legacy_points"`
		} `json:"warlords"`
	} `json:"achievements"`
}
