package structs

type SkyblockProfiles struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause"`
	Data    []struct {
		ProfileID string            `json:"profile_id"`
		Members   map[string]Member `json:"members"`
		CuteName  string            `json:"cute_name"`
		Banking   struct {
			Balance      float64       `json:"balance"`
			Transactions []interface{} `json:"transactions"`
		} `json:"banking"`
		CommunityUpgrades struct {
			CurrentlyUpgrading struct {
				Upgrade    string `json:"upgrade"`
				NewTier    int    `json:"new_tier"`
				StartMs    int64  `json:"start_ms"`
				WhoStarted string `json:"who_started"`
			} `json:"currently_upgrading"`
			UpgradeStates []struct {
				Upgrade     string `json:"upgrade"`
				Tier        int    `json:"tier"`
				StartedMs   int64  `json:"started_ms"`
				StartedBy   string `json:"started_by"`
				ClaimedMs   int64  `json:"claimed_ms"`
				ClaimedBy   string `json:"claimed_by"`
				Fasttracked bool   `json:"fasttracked"`
			} `json:"upgrade_states"`
		} `json:"community_upgrades"`
		GameMode string `json:"game_mode"`
	} `json:"profiles"`
}

type Member struct {
	LastSave int64 `json:"last_save"`
	InvArmor struct {
		Type int    `json:"type"`
		Data string `json:"data"`
	} `json:"inv_armor"`
	CoopInvitation struct {
		Timestamp          int64  `json:"timestamp"`
		InvitedBy          string `json:"invited_by"`
		Confirmed          bool   `json:"confirmed"`
		ConfirmedTimestamp int64  `json:"confirmed_timestamp"`
	} `json:"coop_invitation"`
	FirstJoin    int64 `json:"first_join"`
	FirstJoinHub int   `json:"first_join_hub"`
	Stats        struct {
		Deaths                        float64 `json:"deaths"`
		DeathsVoid                    float64 `json:"deaths_void"`
		AuctionsBids                  float64 `json:"auctions_bids"`
		AuctionsHighestBid            float64 `json:"auctions_highest_bid"`
		AuctionsWon                   float64 `json:"auctions_won"`
		AuctionsBoughtSpecial         float64 `json:"auctions_bought_special"`
		AuctionsGoldSpent             float64 `json:"auctions_gold_spent"`
		AuctionsBoughtRare            float64 `json:"auctions_bought_rare"`
		HighestCriticalDamage         float64 `json:"highest_critical_damage"`
		Kills                         float64 `json:"kills"`
		KillsSpider                   float64 `json:"kills_spider"`
		KillsDiamondZombie            float64 `json:"kills_diamond_zombie"`
		PetMilestoneOresMined         float64 `json:"pet_milestone_ores_mined"`
		DeathsCavernsGhost            float64 `json:"deaths_caverns_ghost"`
		AuctionsBoughtUncommon        float64 `json:"auctions_bought_uncommon"`
		KillsGoblin                   float64 `json:"kills_goblin"`
		KillsGoblinWeaklingMelee      float64 `json:"kills_goblin_weakling_melee"`
		KillsGoblinKnifeThrower       float64 `json:"kills_goblin_knife_thrower"`
		KillsGoblinWeaklingBow        float64 `json:"kills_goblin_weakling_bow"`
		DeathsGoblinWeaklingBow       float64 `json:"deaths_goblin_weakling_bow"`
		DeathsGoblinKnifeThrower      float64 `json:"deaths_goblin_knife_thrower"`
		DeathsGoblinWeaklingMelee     float64 `json:"deaths_goblin_weakling_melee"`
		KillsIceWalker                float64 `json:"kills_ice_walker"`
		KillsGoblinBattler            float64 `json:"kills_goblin_battler"`
		KillsGoblinCreeper            float64 `json:"kills_goblin_creeper"`
		KillsGoblinGolem              float64 `json:"kills_goblin_golem"`
		DeathsIceWalker               float64 `json:"deaths_ice_walker"`
		KillsCrystalSentry            float64 `json:"kills_crystal_sentry"`
		AuctionsBoughtEpic            float64 `json:"auctions_bought_epic"`
		AuctionsCreated               float64 `json:"auctions_created"`
		AuctionsFees                  float64 `json:"auctions_fees"`
		KillsZombie                   float64 `json:"kills_zombie"`
		KillsUnburriedZombie          float64 `json:"kills_unburried_zombie"`
		AuctionsCompleted             float64 `json:"auctions_completed"`
		AuctionsSoldCommon            float64 `json:"auctions_sold_common"`
		AuctionsGoldEarned            float64 `json:"auctions_gold_earned"`
		KillsSpiderJockey             float64 `json:"kills_spider_jockey"`
		KillsJockeySkeleton           float64 `json:"kills_jockey_skeleton"`
		KillsSplitterSpiderSilverfish float64 `json:"kills_splitter_spider_silverfish"`
		KillsVoraciousSpider          float64 `json:"kills_voracious_spider"`
		KillsWeaverSpider             float64 `json:"kills_weaver_spider"`
		KillsDasherSpider             float64 `json:"kills_dasher_spider"`
		KillsSplitterSpider           float64 `json:"kills_splitter_spider"`
		KillsJockeyShotSilverfish     float64 `json:"kills_jockey_shot_silverfish"`
		KillsRuinWolf                 float64 `json:"kills_ruin_wolf"`
		KillsOldWolf                  float64 `json:"kills_old_wolf"`
		DeathsWolf                    float64 `json:"deaths_wolf"`
		DeathsSpider                  float64 `json:"deaths_spider"`
		KillsSkeleton                 float64 `json:"kills_skeleton"`
		KillsArachneBrood             float64 `json:"kills_arachne_brood"`
		KillsTeamTreasuriteWendy      float64 `json:"kills_team_treasurite_wendy"`
		DeathsTeamTreasuriteCorleone  float64 `json:"deaths_team_treasurite_corleone"`
		KillsTeamTreasuriteGrunt      float64 `json:"kills_team_treasurite_grunt"`
		KillsTeamTreasuriteSebastian  float64 `json:"kills_team_treasurite_sebastian"`
		KillsBelle                    float64 `json:"kills_belle"`
		KillsSmog                     float64 `json:"kills_smog"`
		KillsTeamTreasuriteCorleone   float64 `json:"kills_team_treasurite_corleone"`
		DeathsYog                     float64 `json:"deaths_yog"`
		KillsBlaze                    float64 `json:"kills_blaze"`
		KillsFireBat                  float64 `json:"kills_fire_bat"`
		KillsTeamTreasuriteViper      float64 `json:"kills_team_treasurite_viper"`
		KillsYog                      float64 `json:"kills_yog"`
		KillsAutomaton                float64 `json:"kills_automaton"`
		KillsSludge                   float64 `json:"kills_sludge"`
		DeathsAutomaton               float64 `json:"deaths_automaton"`
		KillsWorm                     float64 `json:"kills_worm"`
		KillsCavitak                  float64 `json:"kills_cavitak"`
		KillsThyst                    float64 `json:"kills_thyst"`
		KillsScatha                   float64 `json:"kills_scatha"`
		KillsGoblinCreepertamer       float64 `json:"kills_goblin_creepertamer"`
		KillsSilvo                    float64 `json:"kills_silvo"`
		AuctionsSoldRare              float64 `json:"auctions_sold_rare"`
		KillsGoblinMurderlover        float64 `json:"kills_goblin_murderlover"`
		KillsGoblinFlamethrower       float64 `json:"kills_goblin_flamethrower"`
		AuctionsBoughtCommon          float64 `json:"auctions_bought_common"`
		AuctionsSoldEpic              float64 `json:"auctions_sold_epic"`
	} `json:"stats"`
	Objectives struct {
		CollectLog struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"collect_log"`
		TalkToGuide struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"talk_to_guide"`
		PublicIsland struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"public_island"`
		ExploreHub struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"explore_hub"`
		ExploreVillage struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"explore_village"`
		TalkToLibrarian struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"talk_to_librarian"`
		TalkToFarmer struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_farmer"`
		TalkToBlacksmith struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_blacksmith"`
		TalkToLumberjack struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_lumberjack"`
		TalkToEventMaster struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_event_master"`
		TalkToAuctionMaster struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_auction_master"`
		TalkToBanker struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_banker"`
		TalkToFairy struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_fairy"`
		TalkToFisherman1 struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_fisherman_1"`
		TalkToCarpenter struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_carpenter"`
		TalkToArtist1 struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_artist_1"`
		PaintCanvas struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"paint_canvas"`
		TalkToPetCollector struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_pet_collector"`
		TalkToPetSitter struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_pet_sitter"`
		TalkToFarmhand1 struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_farmhand_1"`
		IncreaseFarmingSkill5 struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"increase_farming_skill_5"`
		WarpMushroomDesert struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"warp_mushroom_desert"`
		EnchantItem struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"enchant_item"`
		CollectFarmingResources2 struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"collect_farming_resources_2"`
		CraftWorkbench struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"craft_workbench"`
		TalkToLapisMiner struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_lapis_miner"`
		TalkToRhys struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"talk_to_rhys"`
		IncreaseMining12 struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"increase_mining_12"`
		HotmGiveMaterials struct {
			Status        string `json:"status"`
			Progress      int    `json:"progress"`
			CompletedAt   int64  `json:"completed_at"`
			ENCHANTEDIRON int    `json:"ENCHANTED_IRON"`
			ENCHANTEDCOAL int    `json:"ENCHANTED_COAL"`
			Started       bool   `json:"started"`
		} `json:"hotm_give_materials"`
		Fetchur258 struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"fetchur-25-8"`
		IncreaseCombatSkill5 struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"increase_combat_skill_5"`
		TalkToRick struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_rick"`
		WarpTheEnd struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"warp_the_end"`
		TalkToArchaeologist struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_archaeologist"`
		TalkToShaggy struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_shaggy"`
		WarpBlazingFortress struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"warp_blazing_fortress"`
		TalkToGwendolyn struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"talk_to_gwendolyn"`
		TalkToBraum struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"talk_to_braum"`
		VisitGreaterMines struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"visit_greater_mines"`
		TalkToTheGoblinKing struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"talk_to_the_goblin_king"`
		KillAutomatons struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"kill_automatons"`
		EnterDivanMines struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"enter_divan_mines"`
		FindAJungleKey struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"find_a_jungle_key"`
		MineRuby struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"mine_ruby"`
		CompleteTrialsOfJungleTemple struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"complete_trials_of_jungle_temple"`
		MineAmethyst struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"mine_amethyst"`
		FindFourMissingPieces struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"find_four_missing_pieces"`
		MineJade struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"mine_jade"`
		FindTheGoblinQueen struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"find_the_goblin_queen"`
		MineAmber struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"mine_amber"`
		TalkToProfessorRobot struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int64  `json:"completed_at"`
		} `json:"talk_to_professor_robot"`
		MineSapphire struct {
			Status      string `json:"status"`
			Progress    int    `json:"progress"`
			CompletedAt int    `json:"completed_at"`
		} `json:"mine_sapphire"`
	} `json:"objectives"`
	Tutorial []string `json:"tutorial"`
	Quests   struct {
		CollectLog struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int64  `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"collect_log"`
		ExploreHub struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"explore_hub"`
		ExploreVillage struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"explore_village"`
		TalkToLibrarian struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int64  `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_librarian"`
		TalkToFarmer struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_farmer"`
		TalkToBlacksmith struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_blacksmith"`
		TalkToLumberjack struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_lumberjack"`
		TalkToAuctionMaster struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_auction_master"`
		TalkToBanker struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_banker"`
		TalkToCarpenter struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_carpenter"`
		TalkToArtist1 struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_artist_1"`
		TalkToFarmhand1 struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_farmhand_1"`
		IncreaseFarmingSkill5 struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"increase_farming_skill_5"`
		TalkToLapisMiner struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_lapis_miner"`
		TalkToRhys struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int64  `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_rhys"`
		IncreaseCombatSkill5 struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"increase_combat_skill_5"`
		TalkToRick struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_rick"`
		TalkToArchaeologist struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_archaeologist"`
		TalkToGwendolyn struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_gwendolyn"`
		TalkToTheGoblinKing struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"talk_to_the_goblin_king"`
		KillAutomatons struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"kill_automatons"`
		EnterDivanMines struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"enter_divan_mines"`
		FindAJungleKey struct {
			Status        string `json:"status"`
			ActivatedAt   int64  `json:"activated_at"`
			ActivatedAtSb int    `json:"activated_at_sb"`
			CompletedAt   int    `json:"completed_at"`
			CompletedAtSb int    `json:"completed_at_sb"`
		} `json:"find_a_jungle_key"`
	} `json:"quests"`
	CoinPurse           float64  `json:"coin_purse"`
	LastDeath           int      `json:"last_death"`
	CraftedGenerators   []string `json:"crafted_generators"`
	VisitedZones        []string `json:"visited_zones"`
	FairySoulsCollected int      `json:"fairy_souls_collected"`
	DeathCount          int      `json:"death_count"`
	SlayerBosses        struct {
		Zombie struct {
			ClaimedLevels struct {
				Level2 bool `json:"level_2"`
				Level1 bool `json:"level_1"`
			} `json:"claimed_levels"`
			BossKillsTier0 int `json:"boss_kills_tier_0"`
			Xp             int `json:"xp"`
			BossKillsTier1 int `json:"boss_kills_tier_1"`
		} `json:"zombie"`
		Spider struct {
			ClaimedLevels struct {
				Level1 bool `json:"level_1"`
				Level2 bool `json:"level_2"`
				Level3 bool `json:"level_3"`
				Level4 bool `json:"level_4"`
				Level5 bool `json:"level_5"`
			} `json:"claimed_levels"`
			BossKillsTier0 int `json:"boss_kills_tier_0"`
			Xp             int `json:"xp"`
			BossKillsTier1 int `json:"boss_kills_tier_1"`
			BossKillsTier2 int `json:"boss_kills_tier_2"`
			BossKillsTier3 int `json:"boss_kills_tier_3"`
		} `json:"spider"`
		Wolf struct {
			ClaimedLevels struct {
				Level3 bool `json:"level_3"`
			} `json:"claimed_levels"`
			BossKillsTier0 int `json:"boss_kills_tier_0"`
			Xp             int `json:"xp"`
			BossKillsTier1 int `json:"boss_kills_tier_1"`
			BossKillsTier2 int `json:"boss_kills_tier_2"`
		} `json:"wolf"`
		Enderman struct {
			ClaimedLevels struct {
			} `json:"claimed_levels"`
		} `json:"enderman"`
	} `json:"slayer_bosses"`
	Pets     []interface{} `json:"pets"`
	Dungeons struct {
		DungeonTypes struct {
			Catacombs struct {
			} `json:"catacombs"`
			MasterCatacombs struct {
			} `json:"master_catacombs"`
		} `json:"dungeon_types"`
		PlayerClasses struct {
			Healer struct {
			} `json:"healer"`
			Mage struct {
			} `json:"mage"`
			Berserk struct {
			} `json:"berserk"`
			Archer struct {
			} `json:"archer"`
			Tank struct {
			} `json:"tank"`
		} `json:"player_classes"`
		DungeonJournal struct {
		} `json:"dungeon_journal"`
	} `json:"dungeons"`
	Griffin struct {
		Burrows []interface{} `json:"burrows"`
	} `json:"griffin"`
	Jacob2 struct {
		MedalsInv struct {
		} `json:"medals_inv"`
		Perks struct {
		} `json:"perks"`
		Contests struct {
		} `json:"contests"`
	} `json:"jacob2"`
	Perks struct {
	} `json:"perks"`
	ActiveEffects         []interface{} `json:"active_effects"`
	DisabledPotionEffects []string      `json:"disabled_potion_effects"`
	VisitedModes          []string      `json:"visited_modes"`
	TempStatBuffs         []struct {
		Stat     int    `json:"stat"`
		Key      string `json:"key"`
		Amount   int    `json:"amount"`
		ExpireAt int64  `json:"expire_at"`
	} `json:"temp_stat_buffs"`
	MiningCore struct {
		Nodes struct {
			MiningSpeed      int `json:"mining_speed"`
			MiningFortune    int `json:"mining_fortune"`
			TitaniumInsanium int `json:"titanium_insanium"`
			MiningSpeedBoost int `json:"mining_speed_boost"`
		} `json:"nodes"`
		ReceivedFreeTier   bool    `json:"received_free_tier"`
		PowderMithril      int     `json:"powder_mithril"`
		PowderMithrilTotal int     `json:"powder_mithril_total"`
		Experience         float64 `json:"experience"`
		Crystals           struct {
			JadeCrystal struct {
				State       string `json:"state"`
				TotalPlaced int    `json:"total_placed"`
			} `json:"jade_crystal"`
			AmberCrystal struct {
				State       string `json:"state"`
				TotalPlaced int    `json:"total_placed"`
			} `json:"amber_crystal"`
			AmethystCrystal struct {
				State       string `json:"state"`
				TotalPlaced int    `json:"total_placed"`
			} `json:"amethyst_crystal"`
			SapphireCrystal struct {
				State       string `json:"state"`
				TotalPlaced int    `json:"total_placed"`
			} `json:"sapphire_crystal"`
			TopazCrystal struct {
				State       string `json:"state"`
				TotalPlaced int    `json:"total_placed"`
			} `json:"topaz_crystal"`
			JasperCrystal struct {
				State string `json:"state"`
			} `json:"jasper_crystal"`
			RubyCrystal struct {
				State string `json:"state"`
			} `json:"ruby_crystal"`
		} `json:"crystals"`
		RetroactiveTier2Token  bool   `json:"retroactive_tier2_token"`
		TokensSpent            int    `json:"tokens_spent"`
		PowderSpentMithril     int    `json:"powder_spent_mithril"`
		SelectedPickaxeAbility string `json:"selected_pickaxe_ability"`
		GreaterMinesLastAccess int64  `json:"greater_mines_last_access"`
		Biomes                 struct {
			Dwarven struct {
			} `json:"dwarven"`
			Precursor struct {
				PartsDelivered []interface{} `json:"parts_delivered"`
			} `json:"precursor"`
			Goblin struct {
				KingQuestActive     bool `json:"king_quest_active"`
				KingQuestsCompleted int  `json:"king_quests_completed"`
			} `json:"goblin"`
		} `json:"biomes"`
		PowderGemstone      int   `json:"powder_gemstone"`
		PowderGemstoneTotal int   `json:"powder_gemstone_total"`
		LastReset           int64 `json:"last_reset"`
	} `json:"mining_core"`
	Forge struct {
		ForgeProcesses struct {
		} `json:"forge_processes"`
	} `json:"forge"`
}
