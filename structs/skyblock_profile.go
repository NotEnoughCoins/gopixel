package structs

type SkyblockProfile struct {
	Success bool `json:"success"`
	Profile struct {
		ProfileID string `json:"profile_id"`
		Members   struct {
			UUID struct {
				LastSave int `json:"last_save"`
				InvArmor struct {
					Type int    `json:"type"`
					Data string `json:"data"`
				} `json:"inv_armor"`
				FirstJoin    int `json:"first_join"`
				FirstJoinHub int `json:"first_join_hub"`
				Stats        struct {
					Deaths                                                int `json:"deaths"`
					DeathsVoid                                            int `json:"deaths_void"`
					Kills                                                 int `json:"kills"`
					KillsEmeraldSlime                                     int `json:"kills_emerald_slime"`
					AuctionsBids                                          int `json:"auctions_bids"`
					AuctionsHighestBid                                    int `json:"auctions_highest_bid"`
					KillsZombie                                           int `json:"kills_zombie"`
					AuctionsWon                                           int `json:"auctions_won"`
					AuctionsBoughtRare                                    int `json:"auctions_bought_rare"`
					AuctionsGoldSpent                                     int `json:"auctions_gold_spent"`
					KillsChicken                                          int `json:"kills_chicken"`
					DeathsZombie                                          int `json:"deaths_zombie"`
					DeathsSkeleton                                        int `json:"deaths_skeleton"`
					HighestCritDamage                                     int `json:"highest_crit_damage"`
					KillsSkeleton                                         int `json:"kills_skeleton"`
					KillsSpider                                           int `json:"kills_spider"`
					AuctionsBoughtUncommon                                int `json:"auctions_bought_uncommon"`
					KillsDiamondSkeleton                                  int `json:"kills_diamond_skeleton"`
					KillsDiamondZombie                                    int `json:"kills_diamond_zombie"`
					KillsZombieVillager                                   int `json:"kills_zombie_villager"`
					KillsRedstonePigman                                   int `json:"kills_redstone_pigman"`
					KillsInvisibleCreeper                                 int `json:"kills_invisible_creeper"`
					KillsWitch                                            int `json:"kills_witch"`
					ItemsFished                                           int `json:"items_fished"`
					ItemsFishedNormal                                     int `json:"items_fished_normal"`
					KillsSeaWalker                                        int `json:"kills_sea_walker"`
					KillsPondSquid                                        int `json:"kills_pond_squid"`
					ItemsFishedLargeTreasure                              int `json:"items_fished_large_treasure"`
					KillsSeaGuardian                                      int `json:"kills_sea_guardian"`
					ItemsFishedTreasure                                   int `json:"items_fished_treasure"`
					KillsUnburriedZombie                                  int `json:"kills_unburried_zombie"`
					DeathsUnburriedZombie                                 int `json:"deaths_unburried_zombie"`
					KillsRuinWolf                                         int `json:"kills_ruin_wolf"`
					KillsHorsemanZombie                                   int `json:"kills_horseman_zombie"`
					KillsLapisZombie                                      int `json:"kills_lapis_zombie"`
					DeathsFire                                            int `json:"deaths_fire"`
					KillsSplitterSpider                                   int `json:"kills_splitter_spider"`
					KillsWeaverSpider                                     int `json:"kills_weaver_spider"`
					KillsVoraciousSpider                                  int `json:"kills_voracious_spider"`
					KillsSplitterSpiderSilverfish                         int `json:"kills_splitter_spider_silverfish"`
					KillsJockeyShotSilverfish                             int `json:"kills_jockey_shot_silverfish"`
					KillsDasherSpider                                     int `json:"kills_dasher_spider"`
					KillsJockeySkeleton                                   int `json:"kills_jockey_skeleton"`
					KillsSpiderJockey                                     int `json:"kills_spider_jockey"`
					KillsWitherSkeleton                                   int `json:"kills_wither_skeleton"`
					DeathsWitherSkeleton                                  int `json:"deaths_wither_skeleton"`
					KillsFireballMagmaCube                                int `json:"kills_fireball_magma_cube"`
					KillsRabbit                                           int `json:"kills_rabbit"`
					KillsSheep                                            int `json:"kills_sheep"`
					EndRaceBestTime                                       int `json:"end_race_best_time"`
					DeathsFall                                            int `json:"deaths_fall"`
					DeathsSpider                                          int `json:"deaths_spider"`
					KillsPig                                              int `json:"kills_pig"`
					KillsCow                                              int `json:"kills_cow"`
					AuctionsBoughtEpic                                    int `json:"auctions_bought_epic"`
					KillsEnderman                                         int `json:"kills_enderman"`
					KillsRandomSlime                                      int `json:"kills_random_slime"`
					KillsRespawningSkeleton                               int `json:"kills_respawning_skeleton"`
					AuctionsCreated                                       int `json:"auctions_created"`
					AuctionsFees                                          int `json:"auctions_fees"`
					AuctionsCompleted                                     int `json:"auctions_completed"`
					AuctionsSoldCommon                                    int `json:"auctions_sold_common"`
					AuctionsGoldEarned                                    int `json:"auctions_gold_earned"`
					KillsWatcher                                          int `json:"kills_watcher"`
					KillsZealotEnderman                                   int `json:"kills_zealot_enderman"`
					KillsObsidianWither                                   int `json:"kills_obsidian_wither"`
					KillsEndermite                                        int `json:"kills_endermite"`
					DeathsUnknown                                         int `json:"deaths_unknown"`
					AuctionsSoldEpic                                      int `json:"auctions_sold_epic"`
					KillsBatPinata                                        int `json:"kills_bat_pinata"`
					DeathsDrowning                                        int `json:"deaths_drowning"`
					KillsBlaze                                            int `json:"kills_blaze"`
					AuctionsSoldSpecial                                   int `json:"auctions_sold_special"`
					KillsGeneratorGhast                                   int `json:"kills_generator_ghast"`
					KillsOldWolf                                          int `json:"kills_old_wolf"`
					AuctionsBoughtCommon                                  int `json:"auctions_bought_common"`
					DeathsWolf                                            int `json:"deaths_wolf"`
					KillsMagmaCube                                        int `json:"kills_magma_cube"`
					KillsPigman                                           int `json:"kills_pigman"`
					KillsPackSpirit                                       int `json:"kills_pack_spirit"`
					KillsHowlingSpirit                                    int `json:"kills_howling_spirit"`
					KillsSoulOfTheAlpha                                   int `json:"kills_soul_of_the_alpha"`
					KillsNightRespawningSkeleton                          int `json:"kills_night_respawning_skeleton"`
					AuctionsSoldRare                                      int `json:"auctions_sold_rare"`
					HighestCriticalDamage                                 int `json:"highest_critical_damage"`
					KillsSeaArcher                                        int `json:"kills_sea_archer"`
					KillsZombieDeep                                       int `json:"kills_zombie_deep"`
					KillsCatfish                                          int `json:"kills_catfish"`
					KillsChickenDeep                                      int `json:"kills_chicken_deep"`
					DeathsOldWolf                                         int `json:"deaths_old_wolf"`
					AuctionsSoldUncommon                                  int `json:"auctions_sold_uncommon"`
					AuctionsSoldLegendary                                 int `json:"auctions_sold_legendary"`
					EnderCrystalsDestroyed                                int `json:"ender_crystals_destroyed"`
					KillsWiseDragon                                       int `json:"kills_wise_dragon"`
					KillsUnstableDragon                                   int `json:"kills_unstable_dragon"`
					KillsStrongDragon                                     int `json:"kills_strong_dragon"`
					KillsProtectorDragon                                  int `json:"kills_protector_dragon"`
					GiftsReceived                                         int `json:"gifts_received"`
					GiftsGiven                                            int `json:"gifts_given"`
					KillsLiquidHotMagma                                   int `json:"kills_liquid_hot_magma"`
					MostWinterSnowballsHit                                int `json:"most_winter_snowballs_hit"`
					MostWinterDamageDealt                                 int `json:"most_winter_damage_dealt"`
					MostWinterMagmaDamageDealt                            int `json:"most_winter_magma_damage_dealt"`
					DeathsPlayer                                          int `json:"deaths_player"`
					DeathsLiquidHotMagma                                  int `json:"deaths_liquid_hot_magma"`
					DeathsMagmaCube                                       int `json:"deaths_magma_cube"`
					KillsNightSquid                                       int `json:"kills_night_squid"`
					DeathsSeaLeech                                        int `json:"deaths_sea_leech"`
					KillsOldDragon                                        int `json:"kills_old_dragon"`
					DeathsStrongDragon                                    int `json:"deaths_strong_dragon"`
					DeathsSuperiorDragon                                  int `json:"deaths_superior_dragon"`
					KillsSeaLeech                                         int `json:"kills_sea_leech"`
					KillsBroodMotherSpider                                int `json:"kills_brood_mother_spider"`
					KillsBroodMotherCaveSpider                            int `json:"kills_brood_mother_cave_spider"`
					AuctionsNoBids                                        int `json:"auctions_no_bids"`
					KillsYoungDragon                                      int `json:"kills_young_dragon"`
					KillsSuperiorDragon                                   int `json:"kills_superior_dragon"`
					AuctionsBoughtLegendary                               int `json:"auctions_bought_legendary"`
					KillsCaveSpider                                       int `json:"kills_cave_spider"`
					KillsPlayer                                           int `json:"kills_player"`
					DungeonHubGiantMushroomAnythingNoReturnBestTime       int `json:"dungeon_hub_giant_mushroom_anything_no_return_best_time"`
					DungeonHubPrecursorRuinsAnythingNoReturnBestTime      int `json:"dungeon_hub_precursor_ruins_anything_no_return_best_time"`
					DungeonHubCrystalCoreAnythingNoReturnBestTime         int `json:"dungeon_hub_crystal_core_anything_no_return_best_time"`
					KillsZombieGrunt                                      int `json:"kills_zombie_grunt"`
					KillsSkeletonGrunt                                    int `json:"kills_skeleton_grunt"`
					KillsDungeonRespawningSkeleton                        int `json:"kills_dungeon_respawning_skeleton"`
					KillsCryptLurker                                      int `json:"kills_crypt_lurker"`
					KillsCryptDreadlord                                   int `json:"kills_crypt_dreadlord"`
					KillsCryptTankZombie                                  int `json:"kills_crypt_tank_zombie"`
					KillsScaredSkeleton                                   int `json:"kills_scared_skeleton"`
					KillsDiamondGuy                                       int `json:"kills_diamond_guy"`
					DeathsLostAdventurer                                  int `json:"deaths_lost_adventurer"`
					KillsCryptSouleater                                   int `json:"kills_crypt_souleater"`
					KillsSkeletonSoldier                                  int `json:"kills_skeleton_soldier"`
					KillsCryptUndead                                      int `json:"kills_crypt_undead"`
					KillsWatcherSummonUndead                              int `json:"kills_watcher_summon_undead"`
					KillsBonzoSummonUndead                                int `json:"kills_bonzo_summon_undead"`
					KillsLostAdventurer                                   int `json:"kills_lost_adventurer"`
					DeathsBlaze                                           int `json:"deaths_blaze"`
					DeathsEnderman                                        int `json:"deaths_enderman"`
					DeathsLapisZombie                                     int `json:"deaths_lapis_zombie"`
					DeathsRuinWolf                                        int `json:"deaths_ruin_wolf"`
					DeathsEmeraldSlime                                    int `json:"deaths_emerald_slime"`
					DeathsWeaverSpider                                    int `json:"deaths_weaver_spider"`
					DeathsDasherSpider                                    int `json:"deaths_dasher_spider"`
					DeathsDiamondZombie                                   int `json:"deaths_diamond_zombie"`
					DeathsSplitterSpider                                  int `json:"deaths_splitter_spider"`
					DeathsSplitterSpiderSilverfish                        int `json:"deaths_splitter_spider_silverfish"`
					DeathsRedstonePigman                                  int `json:"deaths_redstone_pigman"`
					DeathsSpiderJockey                                    int `json:"deaths_spider_jockey"`
					DeathsDiamondSkeleton                                 int `json:"deaths_diamond_skeleton"`
					DeathsFireballMagmaCube                               int `json:"deaths_fireball_magma_cube"`
					DeathsZombieDeep                                      int `json:"deaths_zombie_deep"`
					DeathsWatcher                                         int `json:"deaths_watcher"`
					DeathsObsidianWither                                  int `json:"deaths_obsidian_wither"`
					DeathsEndermite                                       int `json:"deaths_endermite"`
					KillsGeneratorSlime                                   int `json:"kills_generator_slime"`
					KillsSlime                                            int `json:"kills_slime"`
					KillsGhast                                            int `json:"kills_ghast"`
					DeathsGeneratorSlime                                  int `json:"deaths_generator_slime"`
					DeathsZealotEnderman                                  int `json:"deaths_zealot_enderman"`
					DeathsOldDragon                                       int `json:"deaths_old_dragon"`
					DeathsWiseDragon                                      int `json:"deaths_wise_dragon"`
					KillsForestIslandBat                                  int `json:"kills_forest_island_bat"`
					KillsMagmaCubeBoss                                    int `json:"kills_magma_cube_boss"`
					DeathsMagmaCubeBoss                                   int `json:"deaths_magma_cube_boss"`
					KillsGeneratorMagmaCube                               int `json:"kills_generator_magma_cube"`
					DeathsCaveSpider                                      int `json:"deaths_cave_spider"`
					KillsSeaWitch                                         int `json:"kills_sea_witch"`
					KillsCreeper                                          int `json:"kills_creeper"`
					KillsGuardianDefender                                 int `json:"kills_guardian_defender"`
					KillsDeepSeaProtector                                 int `json:"kills_deep_sea_protector"`
					DeathsWaterHydra                                      int `json:"deaths_water_hydra"`
					KillsWaterHydra                                       int `json:"kills_water_hydra"`
					DeathsProtectorDragon                                 int `json:"deaths_protector_dragon"`
					ChickenRaceBestTime                                   int `json:"chicken_race_best_time"`
					KillsFrozenSteve                                      int `json:"kills_frozen_steve"`
					KillsFrostyTheSnowman                                 int `json:"kills_frosty_the_snowman"`
					ChickenRaceBestTime2                                  int `json:"chicken_race_best_time_2"`
					KillsGuardianEmperor                                  int `json:"kills_guardian_emperor"`
					KillsSkeletonEmperor                                  int `json:"kills_skeleton_emperor"`
					KillsCarrotKing                                       int `json:"kills_carrot_king"`
					KillsYeti                                             int `json:"kills_yeti"`
					DeathsYeti                                            int `json:"deaths_yeti"`
					DeathsPackSpirit                                      int `json:"deaths_pack_spirit"`
					DeathsSoulOfTheAlpha                                  int `json:"deaths_soul_of_the_alpha"`
					ShredderBait                                          int `json:"shredder_bait"`
					ShredderFished                                        int `json:"shredder_fished"`
					KillsGrinch                                           int `json:"kills_grinch"`
					DeathsGuardianEmperor                                 int `json:"deaths_guardian_emperor"`
					AuctionsBoughtSpecial                                 int `json:"auctions_bought_special"`
					ForagingRaceBestTime                                  int `json:"foraging_race_best_time"`
					PetMilestoneSeaCreaturesKilled                        int `json:"pet_milestone_sea_creatures_killed"`
					PetMilestoneOresMined                                 int `json:"pet_milestone_ores_mined"`
					DeathsRevenantZombie                                  int `json:"deaths_revenant_zombie"`
					KillsRevenantZombie                                   int `json:"kills_revenant_zombie"`
					DungeonHubCrystalCoreNoPearlsNoReturnBestTime         int `json:"dungeon_hub_crystal_core_no_pearls_no_return_best_time"`
					DungeonHubCrystalCoreNoAbilitiesNoReturnBestTime      int `json:"dungeon_hub_crystal_core_no_abilities_no_return_best_time"`
					KillsCellarSpider                                     int `json:"kills_cellar_spider"`
					KillsSniperSkeleton                                   int `json:"kills_sniper_skeleton"`
					DeathsWatcherSummonUndead                             int `json:"deaths_watcher_summon_undead"`
					DeathsCryptLurker                                     int `json:"deaths_crypt_lurker"`
					KillsCorruptedProtector                               int `json:"kills_corrupted_protector"`
					DeathsSkeletonEmperor                                 int `json:"deaths_skeleton_emperor"`
					KillsHorsemanBat                                      int `json:"kills_horseman_bat"`
					DeathsCorruptedProtector                              int `json:"deaths_corrupted_protector"`
					DungeonHubPrecursorRuinsNoPearlsNoReturnBestTime      int `json:"dungeon_hub_precursor_ruins_no_pearls_no_return_best_time"`
					DungeonHubPrecursorRuinsAnythingWithReturnBestTime    int `json:"dungeon_hub_precursor_ruins_anything_with_return_best_time"`
					DungeonHubPrecursorRuinsNoPearlsWithReturnBestTime    int `json:"dungeon_hub_precursor_ruins_no_pearls_with_return_best_time"`
					KillsHorsemanHorse                                    int `json:"kills_horseman_horse"`
					KillsDungeonSecretBat                                 int `json:"kills_dungeon_secret_bat"`
					KillsSkeletonMaster                                   int `json:"kills_skeleton_master"`
					DeathsScarfWarrior                                    int `json:"deaths_scarf_warrior"`
					KillsScarfWarrior                                     int `json:"kills_scarf_warrior"`
					KillsScarfMage                                        int `json:"kills_scarf_mage"`
					DeathsCryptDreadlord                                  int `json:"deaths_crypt_dreadlord"`
					DeathsSkeletonSoldier                                 int `json:"deaths_skeleton_soldier"`
					KillsParasite                                         int `json:"kills_parasite"`
					DeathsScarf                                           int `json:"deaths_scarf"`
					KillsLonelySpider                                     int `json:"kills_lonely_spider"`
					DeathsScarfMage                                       int `json:"deaths_scarf_mage"`
					KillsScarfPriest                                      int `json:"kills_scarf_priest"`
					KillsBlazeHigherOrLower                               int `json:"kills_blaze_higher_or_lower"`
					KillsZombieSoldier                                    int `json:"kills_zombie_soldier"`
					DeathsSkeletor                                        int `json:"deaths_skeletor"`
					DeathsSkeletonGrunt                                   int `json:"deaths_skeleton_grunt"`
					KillsScarfArcher                                      int `json:"kills_scarf_archer"`
					DeathsCryptSouleater                                  int `json:"deaths_crypt_souleater"`
					DeathsSkeletonMaster                                  int `json:"deaths_skeleton_master"`
					KillsDungeonRespawningSkeletonSkull                   int `json:"kills_dungeon_respawning_skeleton_skull"`
					DeathsTrap                                            int `json:"deaths_trap"`
					KillsCryptUndeadPieter                                int `json:"kills_crypt_undead_pieter"`
					KillsCryptUndeadValentin                              int `json:"kills_crypt_undead_valentin"`
					KillsShadowAssassin                                   int `json:"kills_shadow_assassin"`
					KillsSkeletor                                         int `json:"kills_skeletor"`
					DeathsShadowAssassin                                  int `json:"deaths_shadow_assassin"`
					DeathsDeathmite                                       int `json:"deaths_deathmite"`
					KillsWatcherBonzo                                     int `json:"kills_watcher_bonzo"`
					KillsProfessorGuardianSummon                          int `json:"kills_professor_guardian_summon"`
					DeathsProfessor                                       int `json:"deaths_professor"`
					KillsZombieKnight                                     int `json:"kills_zombie_knight"`
					DeathsProfessorMageGuardian                           int `json:"deaths_professor_mage_guardian"`
					DeathsScaredSkeleton                                  int `json:"deaths_scared_skeleton"`
					KillsCryptUndeadChristian                             int `json:"kills_crypt_undead_christian"`
					DeathsDiamondGuy                                      int `json:"deaths_diamond_guy"`
					DungeonHubGiantMushroomNoPearlsNoReturnBestTime       int `json:"dungeon_hub_giant_mushroom_no_pearls_no_return_best_time"`
					KillsCryptUndeadNicholas                              int `json:"kills_crypt_undead_nicholas"`
					KillsCryptUndeadBernhard                              int `json:"kills_crypt_undead_bernhard"`
					KillsCryptUndeadFriedrich                             int `json:"kills_crypt_undead_friedrich"`
					KillsCryptUndeadAlexander                             int `json:"kills_crypt_undead_alexander"`
					KillsCryptUndeadMarius                                int `json:"kills_crypt_undead_marius"`
					KillsKingMidas                                        int `json:"kills_king_midas"`
					DungeonHubGiantMushroomAnythingWithReturnBestTime     int `json:"dungeon_hub_giant_mushroom_anything_with_return_best_time"`
					DungeonHubGiantMushroomNoPearlsWithReturnBestTime     int `json:"dungeon_hub_giant_mushroom_no_pearls_with_return_best_time"`
					DungeonHubCrystalCoreAnythingWithReturnBestTime       int `json:"dungeon_hub_crystal_core_anything_with_return_best_time"`
					DungeonHubCrystalCoreNoPearlsWithReturnBestTime       int `json:"dungeon_hub_crystal_core_no_pearls_with_return_best_time"`
					KillsDeathmite                                        int `json:"kills_deathmite"`
					DeathsSuffocation                                     int `json:"deaths_suffocation"`
					DeathsYoungDragon                                     int `json:"deaths_young_dragon"`
					DeathsDeepSeaProtector                                int `json:"deaths_deep_sea_protector"`
					KillsTarantulaSpider                                  int `json:"kills_tarantula_spider"`
					DeathsTarantulaSpider                                 int `json:"deaths_tarantula_spider"`
					MostWinterCannonballsHit                              int `json:"most_winter_cannonballs_hit"`
					DeathsDungeonRespawningSkeleton                       int `json:"deaths_dungeon_respawning_skeleton"`
					DeathsProfessorGuardianSummon                         int `json:"deaths_professor_guardian_summon"`
					DeathsSniperSkeleton                                  int `json:"deaths_sniper_skeleton"`
					DeathsWither                                          int `json:"deaths_wither"`
					DeathsUnstableDragon                                  int `json:"deaths_unstable_dragon"`
					DeathsGeneratorGhast                                  int `json:"deaths_generator_ghast"`
					DungeonHubGiantMushroomNoAbilitiesNoReturnBestTime    int `json:"dungeon_hub_giant_mushroom_no_abilities_no_return_best_time"`
					DungeonHubGiantMushroomNothingNoReturnBestTime        int `json:"dungeon_hub_giant_mushroom_nothing_no_return_best_time"`
					DeathsPigman                                          int `json:"deaths_pigman"`
					DeathsCatfish                                         int `json:"deaths_catfish"`
					DeathsGuardianDefender                                int `json:"deaths_guardian_defender"`
					DungeonHubCrystalCoreNothingNoReturnBestTime          int `json:"dungeon_hub_crystal_core_nothing_no_return_best_time"`
					DungeonHubCrystalCoreNoAbilitiesWithReturnBestTime    int `json:"dungeon_hub_crystal_core_no_abilities_with_return_best_time"`
					DungeonHubCrystalCoreNothingWithReturnBestTime        int `json:"dungeon_hub_crystal_core_nothing_with_return_best_time"`
					DungeonHubGiantMushroomNoAbilitiesWithReturnBestTime  int `json:"dungeon_hub_giant_mushroom_no_abilities_with_return_best_time"`
					DungeonHubGiantMushroomNothingWithReturnBestTime      int `json:"dungeon_hub_giant_mushroom_nothing_with_return_best_time"`
					DungeonHubPrecursorRuinsNoAbilitiesNoReturnBestTime   int `json:"dungeon_hub_precursor_ruins_no_abilities_no_return_best_time"`
					DungeonHubPrecursorRuinsNothingNoReturnBestTime       int `json:"dungeon_hub_precursor_ruins_nothing_no_return_best_time"`
					DungeonHubPrecursorRuinsNoAbilitiesWithReturnBestTime int `json:"dungeon_hub_precursor_ruins_no_abilities_with_return_best_time"`
					DungeonHubPrecursorRuinsNothingWithReturnBestTime     int `json:"dungeon_hub_precursor_ruins_nothing_with_return_best_time"`
					DeathsZombieVillager                                  int `json:"deaths_zombie_villager"`
					DeathsHowlingSpirit                                   int `json:"deaths_howling_spirit"`
					DeathsProfessorArcherGuardian                         int `json:"deaths_professor_archer_guardian"`
					DeathsSeaGuardian                                     int `json:"deaths_sea_guardian"`
					DeathsRespawningSkeleton                              int `json:"deaths_respawning_skeleton"`
					DeathsJockeyShotSilverfish                            int `json:"deaths_jockey_shot_silverfish"`
					KillsHeadlessHorseman                                 int `json:"kills_headless_horseman"`
					DeathsKingMidas                                       int `json:"deaths_king_midas"`
					KillsSuperArcher                                      int `json:"kills_super_archer"`
					KillsCryptWitherskeleton                              int `json:"kills_crypt_witherskeleton"`
					KillsSpiritWolf                                       int `json:"kills_spirit_wolf"`
					KillsSpiritBull                                       int `json:"kills_spirit_bull"`
					KillsSpiritRabbit                                     int `json:"kills_spirit_rabbit"`
					KillsSpiritChicken                                    int `json:"kills_spirit_chicken"`
					KillsSpiritBat                                        int `json:"kills_spirit_bat"`
					KillsSpiritSheep                                      int `json:"kills_spirit_sheep"`
					DeathsSpiritChicken                                   int `json:"deaths_spirit_chicken"`
					KillsSpiritMiniboss                                   int `json:"kills_spirit_miniboss"`
					DeathsSpiritBat                                       int `json:"deaths_spirit_bat"`
					KillsSuperTankZombie                                  int `json:"kills_super_tank_zombie"`
					KillsWatcherScarf                                     int `json:"kills_watcher_scarf"`
					DeathsSpiritWolf                                      int `json:"deaths_spirit_wolf"`
					KillsThorn                                            int `json:"kills_thorn"`
					DeathsSpiritMiniboss                                  int `json:"deaths_spirit_miniboss"`
					DeathsSpiritRabbit                                    int `json:"deaths_spirit_rabbit"`
					DeathsSpiritBull                                      int `json:"deaths_spirit_bull"`
					DeathsWatcherScarf                                    int `json:"deaths_watcher_scarf"`
					DeathsWatcherGuardian                                 int `json:"deaths_watcher_guardian"`
					DeathsZombieKnight                                    int `json:"deaths_zombie_knight"`
					DeathsCryptTankZombie                                 int `json:"deaths_crypt_tank_zombie"`
					DeathsZombieSoldier                                   int `json:"deaths_zombie_soldier"`
					DeathsSuperArcher                                     int `json:"deaths_super_archer"`
					DeathsWatcherBonzo                                    int `json:"deaths_watcher_bonzo"`
					DeathsLonelySpider                                    int `json:"deaths_lonely_spider"`
					DeathsCryptUndead                                     int `json:"deaths_crypt_undead"`
					MythosBurrowsDugNext                                  int `json:"mythos_burrows_dug_next"`
					MythosBurrowsDugNextCOMMON                            int `json:"mythos_burrows_dug_next_COMMON"`
					MythosBurrowsDugCombat                                int `json:"mythos_burrows_dug_combat"`
					MythosBurrowsDugCombatCOMMON                          int `json:"mythos_burrows_dug_combat_COMMON"`
					MythosKills                                           int `json:"mythos_kills"`
					KillsSiameseLynx                                      int `json:"kills_siamese_lynx"`
					KillsMinosHunter                                      int `json:"kills_minos_hunter"`
					MythosBurrowsDugTreasure                              int `json:"mythos_burrows_dug_treasure"`
					MythosBurrowsDugTreasureCOMMON                        int `json:"mythos_burrows_dug_treasure_COMMON"`
					MythosBurrowsChainsComplete                           int `json:"mythos_burrows_chains_complete"`
					MythosBurrowsChainsCompleteCOMMON                     int `json:"mythos_burrows_chains_complete_COMMON"`
					MythosBurrowsDugNextUNCOMMON                          int `json:"mythos_burrows_dug_next_UNCOMMON"`
					MythosBurrowsDugCombatUNCOMMON                        int `json:"mythos_burrows_dug_combat_UNCOMMON"`
					KillsMinotaur                                         int `json:"kills_minotaur"`
					MythosBurrowsDugTreasureUNCOMMON                      int `json:"mythos_burrows_dug_treasure_UNCOMMON"`
					MythosBurrowsChainsCompleteUNCOMMON                   int `json:"mythos_burrows_chains_complete_UNCOMMON"`
					MythosBurrowsDugNextRARE                              int `json:"mythos_burrows_dug_next_RARE"`
					MythosBurrowsDugCombatRARE                            int `json:"mythos_burrows_dug_combat_RARE"`
					KillsGaiaConstruct                                    int `json:"kills_gaia_construct"`
					DeathsGaiaConstruct                                   int `json:"deaths_gaia_construct"`
					MythosBurrowsDugNextLEGENDARY                         int `json:"mythos_burrows_dug_next_LEGENDARY"`
					MythosBurrowsDugCombatLEGENDARY                       int `json:"mythos_burrows_dug_combat_LEGENDARY"`
					DeathsMinosChampion                                   int `json:"deaths_minos_champion"`
					DeathsMinotaur                                        int `json:"deaths_minotaur"`
					MythosBurrowsDugTreasureLEGENDARY                     int `json:"mythos_burrows_dug_treasure_LEGENDARY"`
					MythosBurrowsChainsCompleteLEGENDARY                  int `json:"mythos_burrows_chains_complete_LEGENDARY"`
					KillsMinosChampion                                    int `json:"kills_minos_champion"`
					KillsMinosInquisitor                                  int `json:"kills_minos_inquisitor"`
					MythosBurrowsDugNextNull                              int `json:"mythos_burrows_dug_next_null"`
					MythosBurrowsDugCombatNull                            int `json:"mythos_burrows_dug_combat_null"`
					MythosBurrowsDugTreasureRARE                          int `json:"mythos_burrows_dug_treasure_RARE"`
					MythosBurrowsChainsCompleteRARE                       int `json:"mythos_burrows_chains_complete_RARE"`
					MythosBurrowsDugNextEPIC                              int `json:"mythos_burrows_dug_next_EPIC"`
					MythosBurrowsDugTreasureEPIC                          int `json:"mythos_burrows_dug_treasure_EPIC"`
					MythosBurrowsDugCombatEPIC                            int `json:"mythos_burrows_dug_combat_EPIC"`
					MythosBurrowsChainsCompleteEPIC                       int `json:"mythos_burrows_chains_complete_EPIC"`
					MythosBurrowsDugTreasureNull                          int `json:"mythos_burrows_dug_treasure_null"`
					MythosBurrowsChainsCompleteNull                       int `json:"mythos_burrows_chains_complete_null"`
					DeathsMinosInquisitor                                 int `json:"deaths_minos_inquisitor"`
					KillsNurseShark                                       int `json:"kills_nurse_shark"`
					KillsTigerShark                                       int `json:"kills_tiger_shark"`
					KillsGreatWhiteShark                                  int `json:"kills_great_white_shark"`
					KillsBlueShark                                        int `json:"kills_blue_shark"`
					DeathsGreatWhiteShark                                 int `json:"deaths_great_white_shark"`
					KillsTentaclees                                       int `json:"kills_tentaclees"`
					DeathsLividClone                                      int `json:"deaths_livid_clone"`
					DeathsLivid                                           int `json:"deaths_livid"`
					DeathsTentaclees                                      int `json:"deaths_tentaclees"`
					DeathsBonzoSummonUndead                               int `json:"deaths_bonzo_summon_undead"`
					DeathsProfessorWarriorGuardian                        int `json:"deaths_professor_warrior_guardian"`
					AuctionsBoughtMythic                                  int `json:"auctions_bought_mythic"`
					KillsWatcherLivid                                     int `json:"kills_watcher_livid"`
					DeathsSadanStatue                                     int `json:"deaths_sadan_statue"`
					KillsZombieCommander                                  int `json:"kills_zombie_commander"`
					KillsSadanStatue                                      int `json:"kills_sadan_statue"`
					KillsSadanGiant                                       int `json:"kills_sadan_giant"`
					KillsSadanGolem                                       int `json:"kills_sadan_golem"`
					KillsSkeletorPrime                                    int `json:"kills_skeletor_prime"`
					DeathsSadanGiant                                      int `json:"deaths_sadan_giant"`
					DeathsSadanGolem                                      int `json:"deaths_sadan_golem"`
					DeathsSadan                                           int `json:"deaths_sadan"`
					KillsMimic                                            int `json:"kills_mimic"`
					DeathsSiameseLynx                                     int `json:"deaths_siamese_lynx"`
					DeathsBonzo                                           int `json:"deaths_bonzo"`
					DeathsSkeletorPrime                                   int `json:"deaths_skeletor_prime"`
					KillsWraith                                           int `json:"kills_wraith"`
					KillsWitherGourd                                      int `json:"kills_wither_gourd"`
					KillsScaryJerry                                       int `json:"kills_scary_jerry"`
					KillsPhantomSpirit                                    int `json:"kills_phantom_spirit"`
					KillsTrickOrTreater                                   int `json:"kills_trick_or_treater"`
					KillsBatSpooky                                        int `json:"kills_bat_spooky"`
					KillsBattyWitch                                       int `json:"kills_batty_witch"`
					KillsWitchBat                                         int `json:"kills_witch_bat"`
					KillsScarecrow                                        int `json:"kills_scarecrow"`
					KillsNightmare                                        int `json:"kills_nightmare"`
					KillsPhantomFisherman                                 int `json:"kills_phantom_fisherman"`
					KillsWerewolf                                         int `json:"kills_werewolf"`
					KillsGrimReaper                                       int `json:"kills_grim_reaper"`
					AuctionsSoldMythic                                    int `json:"auctions_sold_mythic"`
					DeathsCryptWitherskeleton                             int `json:"deaths_crypt_witherskeleton"`
					KillsZombieLord                                       int `json:"kills_zombie_lord"`
					DeathsWitherMiner                                     int `json:"deaths_wither_miner"`
					KillsWitherMiner                                      int `json:"kills_wither_miner"`
					KillsWitherGuard                                      int `json:"kills_wither_guard"`
					DeathsMaxor                                           int `json:"deaths_maxor"`
					DeathsWitherGuard                                     int `json:"deaths_wither_guard"`
					KillsSkeletonLord                                     int `json:"kills_skeleton_lord"`
					KillsWatcherGiantLaser                                int `json:"kills_watcher_giant_laser"`
					KillsWatcherGiantBoulder                              int `json:"kills_watcher_giant_boulder"`
					KillsWatcherGiantDiamond                              int `json:"kills_watcher_giant_diamond"`
					KillsNecronGuard                                      int `json:"kills_necron_guard"`
					DeathsWatcherLivid                                    int `json:"deaths_watcher_livid"`
					DeathsNecronGuard                                     int `json:"deaths_necron_guard"`
					KillsWatcherGiantBigfoot                              int `json:"kills_watcher_giant_bigfoot"`
					DeathsWatcherGiantBigfoot                             int `json:"deaths_watcher_giant_bigfoot"`
					DeathsWatcherGiantBoulder                             int `json:"deaths_watcher_giant_boulder"`
					DeathsSuperTankZombie                                 int `json:"deaths_super_tank_zombie"`
					DeathsCrushed                                         int `json:"deaths_crushed"`
					DeathsArmorStand                                      int `json:"deaths_armor_stand"`
					KillsMayorJerryGreen                                  int `json:"kills_mayor_jerry_green"`
					KillsMayorJerryBlue                                   int `json:"kills_mayor_jerry_blue"`
					KillsMayorJerryPurple                                 int `json:"kills_mayor_jerry_purple"`
					KillsMayorJerryGolden                                 int `json:"kills_mayor_jerry_golden"`
					KillsIceWalker                                        int `json:"kills_ice_walker"`
					KillsGoblin                                           int `json:"kills_goblin"`
					DeathsCavernsGhost                                    int `json:"deaths_caverns_ghost"`
					KillsGoblinKnifeThrower                               int `json:"kills_goblin_knife_thrower"`
					KillsGoblinWeaklingMelee                              int `json:"kills_goblin_weakling_melee"`
					KillsGoblinWeaklingBow                                int `json:"kills_goblin_weakling_bow"`
					KillsTreasureHoarder                                  int `json:"kills_treasure_hoarder"`
					KillsGoblinCreepertamer                               int `json:"kills_goblin_creepertamer"`
					KillsGoblinBattler                                    int `json:"kills_goblin_battler"`
					KillsGoblinMurderlover                                int `json:"kills_goblin_murderlover"`
					KillsCavernsGhost                                     int `json:"kills_caverns_ghost"`
					KillsGoblinCreeper                                    int `json:"kills_goblin_creeper"`
					KillsGoblinGolem                                      int `json:"kills_goblin_golem"`
					KillsCrystalSentry                                    int `json:"kills_crystal_sentry"`
					KillsPowderGhast                                      int `json:"kills_powder_ghast"`
					DeathsGoblinMurderlover                               int `json:"deaths_goblin_murderlover"`
					DeathsGoblinKnifeThrower                              int `json:"deaths_goblin_knife_thrower"`
					DeathsIceWalker                                       int `json:"deaths_ice_walker"`
					DeathsCryptUndeadHypixel                              int `json:"deaths_crypt_undead_hypixel"`
					DeathsCryptUndeadFlameboy101                          int `json:"deaths_crypt_undead_flameboy101"`
					DeathsGoblinWeaklingBow                               int `json:"deaths_goblin_weakling_bow"`
					KillsArachneBrood                                     int `json:"kills_arachne_brood"`
					KillsArachneKeeper                                    int `json:"kills_arachne_keeper"`
					DeathsArachne                                         int `json:"deaths_arachne"`
					KillsArachne                                          int `json:"kills_arachne"`
					DeathsArachneBrood                                    int `json:"deaths_arachne_brood"`
					DeathsArachneKeeper                                   int `json:"deaths_arachne_keeper"`
					DeathsCellarSpider                                    int `json:"deaths_cellar_spider"`
					KillsMasterSniperSkeleton                             int `json:"kills_master_sniper_skeleton"`
					KillsMasterCryptTankZombie                            int `json:"kills_master_crypt_tank_zombie"`
					KillsMasterZombieGrunt                                int `json:"kills_master_zombie_grunt"`
					KillsMasterCryptLurker                                int `json:"kills_master_crypt_lurker"`
					KillsMasterScaredSkeleton                             int `json:"kills_master_scared_skeleton"`
					KillsMasterSkeletonSoldier                            int `json:"kills_master_skeleton_soldier"`
					KillsMasterSkeletonGrunt                              int `json:"kills_master_skeleton_grunt"`
					KillsMasterCryptSouleater                             int `json:"kills_master_crypt_souleater"`
					KillsMasterDungeonRespawningSkeleton                  int `json:"kills_master_dungeon_respawning_skeleton"`
					KillsMasterLostAdventurer                             int `json:"kills_master_lost_adventurer"`
					KillsMasterCryptDreadlord                             int `json:"kills_master_crypt_dreadlord"`
					KillsMasterCellarSpider                               int `json:"kills_master_cellar_spider"`
					KillsMasterWatcherSummonUndead                        int `json:"kills_master_watcher_summon_undead"`
					DeathsMasterWatcherSummonUndead                       int `json:"deaths_master_watcher_summon_undead"`
					KillsMasterBonzoSummonUndead                          int `json:"kills_master_bonzo_summon_undead"`
					DeathsMasterBonzo                                     int `json:"deaths_master_bonzo"`
					DeathsSpiritSheep                                     int `json:"deaths_spirit_sheep"`
					KillsMasterCryptUndead                                int `json:"kills_master_crypt_undead"`
					KillsMasterDiamondGuy                                 int `json:"kills_master_diamond_guy"`
					KillsMasterSkeletonMaster                             int `json:"kills_master_skeleton_master"`
					DeathsMasterScarfArcher                               int `json:"deaths_master_scarf_archer"`
					KillsMasterScarfMage                                  int `json:"kills_master_scarf_mage"`
					DeathsMasterScarf                                     int `json:"deaths_master_scarf"`
					KillsDanteGoon                                        int `json:"kills_dante_goon"`
					KillsDanteSlimeGoon                                   int `json:"kills_dante_slime_goon"`
					KillsRat                                              int `json:"kills_rat"`
					KillsMushroomCow                                      int `json:"kills_mushroom_cow"`
					KillsTrapperPig                                       int `json:"kills_trapper_pig"`
					KillsTrapperChicken                                   int `json:"kills_trapper_chicken"`
					KillsTrapperSheep                                     int `json:"kills_trapper_sheep"`
					KillsTrapperCow                                       int `json:"kills_trapper_cow"`
					KillsTrapperRabbit                                    int `json:"kills_trapper_rabbit"`
					DeathsMasterLostAdventurer                            int `json:"deaths_master_lost_adventurer"`
					DeathsZombieGrunt                                     int `json:"deaths_zombie_grunt"`
					DeathsGrimReaper                                      int `json:"deaths_grim_reaper"`
					KillsOasisSheep                                       int `json:"kills_oasis_sheep"`
					DeathsMasterSkeletonSoldier                           int `json:"deaths_master_skeleton_soldier"`
					DeathsMasterScarfWarrior                              int `json:"deaths_master_scarf_warrior"`
					KillsMasterParasite                                   int `json:"kills_master_parasite"`
					KillsMasterScarfPriest                                int `json:"kills_master_scarf_priest"`
					DeathsMasterCryptLurker                               int `json:"deaths_master_crypt_lurker"`
					KillsMasterCryptUndeadBernhard                        int `json:"kills_master_crypt_undead_bernhard"`
					KillsMasterScarfWarrior                               int `json:"kills_master_scarf_warrior"`
					KillsMasterScarfArcher                                int `json:"kills_master_scarf_archer"`
					DeathsMasterScarfMage                                 int `json:"deaths_master_scarf_mage"`
					KillsMasterSkeletor                                   int `json:"kills_master_skeletor"`
					DeathsMasterSkeletor                                  int `json:"deaths_master_skeletor"`
					KillsMasterZombieKnight                               int `json:"kills_master_zombie_knight"`
					KillsMasterZombieSoldier                              int `json:"kills_master_zombie_soldier"`
					KillsMasterLonelySpider                               int `json:"kills_master_lonely_spider"`
					DeathsMasterWatcherBonzo                              int `json:"deaths_master_watcher_bonzo"`
					KillsMasterProfessorGuardianSummon                    int `json:"kills_master_professor_guardian_summon"`
					DeathsMasterProfessorMageGuardian                     int `json:"deaths_master_professor_mage_guardian"`
					KillsMasterSuperTankZombie                            int `json:"kills_master_super_tank_zombie"`
					KillsMasterSpiritBat                                  int `json:"kills_master_spirit_bat"`
					DeathsMasterZombieSoldier                             int `json:"deaths_master_zombie_soldier"`
					DeathsMasterLividClone                                int `json:"deaths_master_livid_clone"`
					DeathsMasterSniperSkeleton                            int `json:"deaths_master_sniper_skeleton"`
					KillsMasterWatcherBonzo                               int `json:"kills_master_watcher_bonzo"`
					DeathsMasterSpiritBat                                 int `json:"deaths_master_spirit_bat"`
					DeathsMasterSpiritSheep                               int `json:"deaths_master_spirit_sheep"`
					DeathsMasterSpiritRabbit                              int `json:"deaths_master_spirit_rabbit"`
					DeathsMasterProfessorGuardianSummon                   int `json:"deaths_master_professor_guardian_summon"`
					KillsMasterCryptWitherskeleton                        int `json:"kills_master_crypt_witherskeleton"`
					DeathsMasterWatcherScarf                              int `json:"deaths_master_watcher_scarf"`
					KillsMasterCryptUndeadPieter                          int `json:"kills_master_crypt_undead_pieter"`
					KillsMasterSpiritWolf                                 int `json:"kills_master_spirit_wolf"`
					KillsMasterSpiritRabbit                               int `json:"kills_master_spirit_rabbit"`
					KillsMasterSpiritSheep                                int `json:"kills_master_spirit_sheep"`
					KillsMasterSpiritBull                                 int `json:"kills_master_spirit_bull"`
					DeathsMasterShadowAssassin                            int `json:"deaths_master_shadow_assassin"`
					DeathsMasterSpiritChicken                             int `json:"deaths_master_spirit_chicken"`
					KillsMasterTentaclees                                 int `json:"kills_master_tentaclees"`
					DeathsMasterTentaclees                                int `json:"deaths_master_tentaclees"`
					KillsMasterSuperArcher                                int `json:"kills_master_super_archer"`
					KillsMasterShadowAssassin                             int `json:"kills_master_shadow_assassin"`
					DeathsMasterCryptWitherskeleton                       int `json:"deaths_master_crypt_witherskeleton"`
					DeathsMasterCryptDreadlord                            int `json:"deaths_master_crypt_dreadlord"`
					DeathsMasterSkeletonMaster                            int `json:"deaths_master_skeleton_master"`
					DeathsMasterZombieKnight                              int `json:"deaths_master_zombie_knight"`
					DeathsMasterLivid                                     int `json:"deaths_master_livid"`
					DeathsMasterCryptSouleater                            int `json:"deaths_master_crypt_souleater"`
					PumpkinLauncherCount                                  int `json:"pumpkin_launcher_count"`
					KillsShrineChargedCreeper                             int `json:"kills_shrine_charged_creeper"`
					KillsShrineSkeletonHorseman                           int `json:"kills_shrine_skeleton_horseman"`
					DeathsMasterDiamondGuy                                int `json:"deaths_master_diamond_guy"`
					DeathsMasterDungeonRespawningSkeleton                 int `json:"deaths_master_dungeon_respawning_skeleton"`
					DeathsMasterSkeletonGrunt                             int `json:"deaths_master_skeleton_grunt"`
					KillsOasisRabbit                                      int `json:"kills_oasis_rabbit"`
					KillsMasterSpiritChicken                              int `json:"kills_master_spirit_chicken"`
					KillsMasterSpiritMiniboss                             int `json:"kills_master_spirit_miniboss"`
					DeathsMasterSpiritMiniboss                            int `json:"deaths_master_spirit_miniboss"`
					KillsVoidlingFanatic                                  int `json:"kills_voidling_fanatic"`
					KillsVoidlingExtremist                                int `json:"kills_voidling_extremist"`
					DeathsVoidlingExtremist                               int `json:"deaths_voidling_extremist"`
					DeathsVoidlingFanatic                                 int `json:"deaths_voidling_fanatic"`
					KillsVoidlingEnderman                                 int `json:"kills_voidling_enderman"`
					KillsThyst                                            int `json:"kills_thyst"`
					KillsSludge                                           int `json:"kills_sludge"`
					KillsAutomaton                                        int `json:"kills_automaton"`
					KillsKeyGuardian                                      int `json:"kills_key_guardian"`
					DeathsAutomaton                                       int `json:"deaths_automaton"`
					KillsTeamTreasuriteViper                              int `json:"kills_team_treasurite_viper"`
					KillsTeamTreasuriteSebastian                          int `json:"kills_team_treasurite_sebastian"`
					KillsYog                                              int `json:"kills_yog"`
					KillsGoblinFlamethrower                               int `json:"kills_goblin_flamethrower"`
					KillsTeamTreasuriteWendy                              int `json:"kills_team_treasurite_wendy"`
					KillsBelle                                            int `json:"kills_belle"`
					KillsFireBat                                          int `json:"kills_fire_bat"`
					KillsWorm                                             int `json:"kills_worm"`
					KillsTeamTreasuriteGrunt                              int `json:"kills_team_treasurite_grunt"`
					DeathsYog                                             int `json:"deaths_yog"`
					KillsSilvo                                            int `json:"kills_silvo"`
					DeathsKalhuikiTribeMan                                int `json:"deaths_kalhuiki_tribe_man"`
					DeathsKalhuikiTribeWoman                              int `json:"deaths_kalhuiki_tribe_woman"`
					KillsButterfly                                        int `json:"kills_butterfly"`
					KillsCavitak                                          int `json:"kills_cavitak"`
					DeathsSludge                                          int `json:"deaths_sludge"`
					KillsTrappedSludge                                    int `json:"kills_trapped_sludge"`
					KillsTeamTreasuriteCorleone                           int `json:"kills_team_treasurite_corleone"`
					KillsSmog                                             int `json:"kills_smog"`
					KillsVittomite                                        int `json:"kills_vittomite"`
					KillsKalhuikiElder                                    int `json:"kills_kalhuiki_elder"`
					KillsKalhuikiTribeMan                                 int `json:"kills_kalhuiki_tribe_man"`
					KillsKalhuikiYoungling                                int `json:"kills_kalhuiki_youngling"`
					KillsScatha                                           int `json:"kills_scatha"`
					DeathsMasterCryptTankZombie                           int `json:"deaths_master_crypt_tank_zombie"`
					KillsLavaPigman                                       int `json:"kills_lava_pigman"`
					KillsLavaBlaze                                        int `json:"kills_lava_blaze"`
					KillsKalhuikiTribeWoman                               int `json:"kills_kalhuiki_tribe_woman"`
					DeathsZombieCommander                                 int `json:"deaths_zombie_commander"`
					DeathsEntity                                          int `json:"deaths_entity"`
					KillsFlamingWorm                                      int `json:"kills_flaming_worm"`
					DeathsTeamTreasuriteCorleone                          int `json:"deaths_team_treasurite_corleone"`
					DeathsMasterSpiritWolf                                int `json:"deaths_master_spirit_wolf"`
					KillsMasterCryptUndeadValentin                        int `json:"kills_master_crypt_undead_valentin"`
					KillsMasterCryptUndeadNicholas                        int `json:"kills_master_crypt_undead_nicholas"`
					KillsMasterCryptUndeadChristian                       int `json:"kills_master_crypt_undead_christian"`
					KillsMasterCryptUndeadFriedrich                       int `json:"kills_master_crypt_undead_friedrich"`
					DeathsMasterSadanGiant                                int `json:"deaths_master_sadan_giant"`
				} `json:"stats"`
				Objectives struct {
					CollectLog struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"collect_log"`
					TalkToGuide struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guide"`
					PublicIsland struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"public_island"`
					CraftWorkbench struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"craft_workbench"`
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
						CompletedAt int    `json:"completed_at"`
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
					TalkToLazyMiner struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_lazy_miner"`
					IncreaseMiningSkill5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"increase_mining_skill_5"`
					TalkToTelekinesisApplier struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_telekinesis_applier"`
					FindPickaxe struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"find_pickaxe"`
					CollectIngots struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						IronIngot   bool   `json:"IRON_INGOT"`
						GoldIngot   bool   `json:"GOLD_INGOT"`
					} `json:"collect_ingots"`
					WarpDeepCaverns struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"warp_deep_caverns"`
					TalkToLapisMiner struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_lapis_miner"`
					TalkToLiftOperator struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_lift_operator"`
					ReachLapisQuarry struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"reach_lapis_quarry"`
					CollectLapis struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						INKSACK4    bool   `json:"INK_SACK:4"`
					} `json:"collect_lapis"`
					DepositCoins struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"deposit_coins"`
					EnchantItem struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"enchant_item"`
					TalkToFarmhand1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_farmhand_1"`
					IncreaseFarmingSkill5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"increase_farming_skill_5"`
					WarpMushroomDesert struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"warp_mushroom_desert"`
					IncreaseForagingSkill5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"increase_foraging_skill_5"`
					TalkToGustave1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_gustave_1"`
					CollectDarkOakLogs struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"collect_dark_oak_logs"`
					TalkToCharlie2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_charlie_2"`
					CollectFarmingResources2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						INKSACK3    bool   `json:"INK_SACK:3"`
						Cactus      bool   `json:"CACTUS"`
						SugarCane   bool   `json:"SUGAR_CANE"`
					} `json:"collect_farming_resources_2"`
					IncreaseCombatSkill5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
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
					WarpBlazingFortress struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"warp_blazing_fortress"`
					CollectClay struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"collect_clay"`
					TalkToFisherman2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_fisherman_2"`
					GiveFairySouls struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"give_fairy_souls"`
					TalkToElle struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_elle"`
					TalkToGuber1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guber_1"`
					TalkToEndDealer struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_end_dealer"`
					CompleteTheEndRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_end_race_1"`
					TalkToGuber2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guber_2"`
					CompleteTheEndRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_end_race_2"`
					TalkToGuber3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guber_3"`
					CompleteTheEndRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_end_race_3"`
					TalkToGuber4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guber_4"`
					CompleteTheEndRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_end_race_4"`
					TalkToGuber5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guber_5"`
					MineCoal struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"mine_coal"`
					TalkToBlacksmith2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_blacksmith_2"`
					IncreaseMiningSkill struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"increase_mining_skill"`
					ReforgeItem struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"reforge_item"`
					WarpGoldMine struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"warp_gold_mine"`
					TalkToMelody struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_melody"`
					FightDragon struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"fight_dragon"`
					KillDangerMobs struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"kill_danger_mobs"`
					TalkToBartender struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_bartender"`
					CollectWheat struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"collect_wheat"`
					TalkToFarmer2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_farmer_2"`
					IncreaseFarmingSkill struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"increase_farming_skill"`
					WarpBarnIsland struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"warp_barn_island"`
					IncreaseCombatSkill struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"increase_combat_skill"`
					WarpSpidersDen struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"warp_spiders_den"`
					TalkToHaymitch struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_haymitch"`
					TalkToFrosty struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_frosty"`
					TalkToGulliver1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_gulliver_1"`
					CompleteTheChickenRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_chicken_race_1"`
					TalkToGulliver2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_gulliver_2"`
					CompleteTheChickenRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_chicken_race_2"`
					TalkToGulliver3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_gulliver_3"`
					CompleteTheChickenRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_chicken_race_3"`
					TalkToGulliver4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_gulliver_4"`
					CompleteTheChickenRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_chicken_race_4"`
					ChopTree struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"chop_tree"`
					TalkToLumberjack2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_lumberjack_2"`
					CollectFarmAnimalResources2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						Rabbit      bool   `json:"RABBIT"`
						Mutton      bool   `json:"MUTTON"`
					} `json:"collect_farm_animal_resources_2"`
					CollectWoolCarpenter struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"collect_wool_carpenter"`
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
					TalkToGuildford1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_1"`
					CompleteTheGiantMushroomAnythingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_with_return_race_1"`
					CompleteTheGiantMushroomNoPearlsWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_with_return_race_1"`
					CompleteTheGiantMushroomNoAbilitiesWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_with_return_race_1"`
					CompleteTheGiantMushroomNothingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_with_return_race_1"`
					CompleteThePrecursorRuinsAnythingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_with_return_race_1"`
					CompleteThePrecursorRuinsNoPearlsWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_with_return_race_1"`
					CompleteThePrecursorRuinsNoAbilitiesWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_with_return_race_1"`
					CompleteThePrecursorRuinsNothingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_with_return_race_1"`
					CompleteTheCrystalCoreAnythingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_with_return_race_1"`
					CompleteTheCrystalCoreNoPearlsWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_with_return_race_1"`
					CompleteTheCrystalCoreNoAbilitiesWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_with_return_race_1"`
					CompleteTheCrystalCoreNothingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_with_return_race_1"`
					CompleteTheGiantMushroomAnythingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_1"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_1"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_1"`
					CompleteTheGiantMushroomNothingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_1"`
					CompleteThePrecursorRuinsAnythingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_1"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_1"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_1"`
					CompleteThePrecursorRuinsNothingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_1"`
					CompleteTheCrystalCoreAnythingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_1"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_1"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_1"`
					CompleteTheCrystalCoreNothingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_1"`
					TalkToGuildfordGiantMushroomAnythingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_1"`
					CompleteTheGiantMushroomAnythingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_2"`
					TalkToGuildfordGiantMushroomAnythingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_2"`
					CompleteTheGiantMushroomAnythingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_3"`
					TalkToGuildfordGiantMushroomAnythingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_3"`
					CompleteTheGiantMushroomAnythingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_4"`
					TalkToGuildfordGiantMushroomAnythingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_4"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_1"`
					CompleteThePrecursorRuinsAnythingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_2"`
					CompleteThePrecursorRuinsAnythingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_3"`
					CompleteThePrecursorRuinsAnythingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_4"`
					TalkToGuildfordCrystalCoreAnythingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_1"`
					CompleteTheCrystalCoreAnythingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_2"`
					TalkToGuildfordCrystalCoreAnythingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_2"`
					CompleteTheCrystalCoreAnythingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_3"`
					TalkToGuildfordCrystalCoreAnythingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_3"`
					CompleteTheCrystalCoreAnythingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_4"`
					TalkToGuildfordCrystalCoreAnythingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_4"`
					CraftWoodPickaxe struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"craft_wood_pickaxe"`
					ReachPigmensDen struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"reach_pigmens_den"`
					CraftWheatMinion struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"craft_wheat_minion"`
					GivePickaxeLapisMiner struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"give_pickaxe_lapis_miner"`
					CollectRedstone struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						Redstone    bool   `json:"REDSTONE"`
					} `json:"collect_redstone"`
					ReachSlimehill struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"reach_slimehill"`
					CollectEmerald struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						Emerald     bool   `json:"EMERALD"`
					} `json:"collect_emerald"`
					ReachDiamondReserve struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"reach_diamond_reserve"`
					IncreaseForagingSkill struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"increase_foraging_skill"`
					WarpForagingIslands struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"warp_foraging_islands"`
					CollectBirchLogs struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"collect_birch_logs"`
					TalkToCharlie struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_charlie"`
					TalkToArtist2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_artist_2"`
					CollectDiamond struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						Diamond     bool   `json:"DIAMOND"`
					} `json:"collect_diamond"`
					ReachObsidianSanctuary struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"reach_obsidian_sanctuary"`
					CollectObsidian struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						Obsidian    bool   `json:"OBSIDIAN"`
					} `json:"collect_obsidian"`
					CollectSpider struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						String      bool   `json:"STRING"`
						SpiderEye   bool   `json:"SPIDER_EYE"`
					} `json:"collect_spider"`
					CollectNetherResources struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						BlazeRod    bool   `json:"BLAZE_ROD"`
						NetherStalk bool   `json:"NETHER_STALK"`
					} `json:"collect_nether_resources"`
					CollectNetherResources2 struct {
						Status        string `json:"status"`
						Progress      int    `json:"progress"`
						CompletedAt   int    `json:"completed_at"`
						Quartz        bool   `json:"QUARTZ"`
						GlowstoneDust bool   `json:"GLOWSTONE_DUST"`
						MagmaCream    bool   `json:"MAGMA_CREAM"`
					} `json:"collect_nether_resources_2"`
					CollectEndStone struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						EnderStone  bool   `json:"ENDER_STONE"`
					} `json:"collect_end_stone"`
					ReachDragonsNest struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"reach_dragons_nest"`
					GiveRickIngots struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"give_rick_ingots"`
					TalkToFarmhand2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_farmhand_2"`
					CollectFarmingResources struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						PotatoItem  bool   `json:"POTATO_ITEM"`
						CarrotItem  bool   `json:"CARROT_ITEM"`
						Pumpkin     bool   `json:"PUMPKIN"`
						Melon       bool   `json:"MELON"`
					} `json:"collect_farming_resources"`
					CollectFarmAnimalResources struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
						RawChicken  bool   `json:"RAW_CHICKEN"`
						Leather     bool   `json:"LEATHER"`
						Pork        bool   `json:"PORK"`
					} `json:"collect_farm_animal_resources"`
					CompleteTheWoodsRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_woods_race_1"`
					TalkToGustave2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_gustave_2"`
					CompleteTheWoodsRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_woods_race_2"`
					TalkToGustave3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_gustave_3"`
					CompleteTheWoodsRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_woods_race_3"`
					TalkToGustave4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_gustave_4"`
					CompleteTheWoodsRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_woods_race_4"`
					TalkToGustave5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_gustave_5"`
					TalkToGulliver5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_gulliver_5"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_1"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_2"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_2"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_3"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_3"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_4"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_4"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_1"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_2"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_2"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_3"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_3"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_4"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_4"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_1"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsAnythingWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_with_return_1"`
					CompleteThePrecursorRuinsAnythingWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_with_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoPearlsWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_with_return_1"`
					CompleteThePrecursorRuinsNoPearlsWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_with_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_2"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_3"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_4"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_1"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_2"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_2"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_3"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_3"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_4"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_4"`
					TalkToGuildfordPrecursorRuinsAnythingWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_with_return_2"`
					CompleteThePrecursorRuinsAnythingWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_with_return_race_3"`
					TalkToGuildfordPrecursorRuinsAnythingWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_with_return_3"`
					CompleteThePrecursorRuinsAnythingWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_with_return_race_4"`
					TalkToGuildfordPrecursorRuinsAnythingWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_with_return_4"`
					TalkToGuildfordPrecursorRuinsNoPearlsWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_with_return_2"`
					CompleteThePrecursorRuinsNoPearlsWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_with_return_race_3"`
					TalkToGuildfordPrecursorRuinsNoPearlsWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_with_return_3"`
					CompleteThePrecursorRuinsNoPearlsWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_with_return_race_4"`
					TalkToGuildfordPrecursorRuinsNoPearlsWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_with_return_4"`
					TalkToGuildfordGiantMushroomAnythingWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_with_return_1"`
					CompleteTheGiantMushroomAnythingWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_with_return_race_2"`
					TalkToGuildfordGiantMushroomAnythingWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_with_return_2"`
					CompleteTheGiantMushroomAnythingWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_with_return_race_3"`
					TalkToGuildfordGiantMushroomAnythingWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_with_return_3"`
					CompleteTheGiantMushroomAnythingWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_with_return_race_4"`
					TalkToGuildfordGiantMushroomNoPearlsWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_with_return_1"`
					CompleteTheGiantMushroomNoPearlsWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_with_return_race_2"`
					TalkToGuildfordGiantMushroomNoPearlsWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_with_return_2"`
					CompleteTheGiantMushroomNoPearlsWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_with_return_race_3"`
					TalkToGuildfordGiantMushroomNoPearlsWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_with_return_3"`
					CompleteTheGiantMushroomNoPearlsWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_with_return_race_4"`
					TalkToGuildfordGiantMushroomNoPearlsWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_with_return_4"`
					TalkToGuildfordCrystalCoreAnythingWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_with_return_1"`
					CompleteTheCrystalCoreAnythingWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_with_return_race_2"`
					TalkToGuildfordCrystalCoreAnythingWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_with_return_2"`
					CompleteTheCrystalCoreAnythingWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_with_return_race_3"`
					TalkToGuildfordCrystalCoreNoPearlsWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_with_return_1"`
					CompleteTheCrystalCoreNoPearlsWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_with_return_race_2"`
					TalkToGuildfordCrystalCoreNoPearlsWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_with_return_2"`
					CompleteTheCrystalCoreNoPearlsWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_with_return_race_3"`
					TalkToGuildfordCrystalCoreNoPearlsWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_with_return_3"`
					CompleteTheCrystalCoreNoPearlsWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_with_return_race_4"`
					TalkToCryptKeeper1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_crypt_keeper_1"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_1"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_2"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_2"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_3"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_3"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_4"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_4"`
					TalkToGuildfordGiantMushroomNothingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_1"`
					CompleteTheGiantMushroomNothingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_2"`
					TalkToGuildfordGiantMushroomNothingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_2"`
					CompleteTheGiantMushroomNothingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_3"`
					TalkToGuildfordGiantMushroomNothingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_3"`
					CompleteTheGiantMushroomNothingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_4"`
					TalkToGuildfordGiantMushroomNothingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_4"`
					AuctionItem struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"auction_item"`
					TalkToNocturne1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_nocturne_1"`
					CollectRunicFragments struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"collect_runic_fragments"`
					TalkToNocturne2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_nocturne_2"`
					TalkToDusk1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_dusk_1"`
					CollectRunes struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"collect_runes"`
					TalkToGuildfordCrystalCoreNothingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_1"`
					CompleteTheCrystalCoreNothingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_2"`
					TalkToGuildfordCrystalCoreNothingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_2"`
					CompleteTheCrystalCoreNothingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_3"`
					TalkToGuildfordCrystalCoreAnythingWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_with_return_3"`
					CompleteTheCrystalCoreAnythingWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_with_return_race_4"`
					TalkToGuildfordCrystalCoreAnythingWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_with_return_4"`
					TalkToGuildfordCrystalCoreNoAbilitiesWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_with_return_1"`
					CompleteTheCrystalCoreNoAbilitiesWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_with_return_race_2"`
					TalkToGuildfordCrystalCoreNoAbilitiesWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_with_return_2"`
					CompleteTheCrystalCoreNoAbilitiesWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_with_return_race_3"`
					TalkToGuildfordCrystalCoreNoAbilitiesWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_with_return_3"`
					CompleteTheCrystalCoreNoAbilitiesWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_with_return_race_4"`
					TalkToGuildfordCrystalCoreNothingWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_with_return_1"`
					CompleteTheCrystalCoreNothingWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_with_return_race_2"`
					TalkToGuildfordCrystalCoreNothingWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_with_return_2"`
					CompleteTheCrystalCoreNothingWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_with_return_race_3"`
					TalkToGuildfordCrystalCoreNothingWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_with_return_3"`
					CompleteTheCrystalCoreNothingWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_with_return_race_4"`
					TalkToGuildfordGiantMushroomNoAbilitiesWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_with_return_1"`
					CompleteTheGiantMushroomNoAbilitiesWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_with_return_race_2"`
					TalkToGuildfordGiantMushroomNoAbilitiesWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_with_return_2"`
					CompleteTheGiantMushroomNoAbilitiesWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_with_return_race_3"`
					TalkToGuildfordGiantMushroomNoAbilitiesWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_with_return_3"`
					CompleteTheGiantMushroomNoAbilitiesWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_with_return_race_4"`
					TalkToGuildfordGiantMushroomNothingWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_with_return_1"`
					CompleteTheGiantMushroomNothingWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_with_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_1"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_2"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_3"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_1"`
					CompleteThePrecursorRuinsNothingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_2"`
					CompleteThePrecursorRuinsNothingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_3"`
					CompleteThePrecursorRuinsNothingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_4"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_with_return_1"`
					CompleteThePrecursorRuinsNoAbilitiesWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_with_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_with_return_2"`
					CompleteThePrecursorRuinsNoAbilitiesWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_with_return_race_3"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_with_return_3"`
					CompleteThePrecursorRuinsNoAbilitiesWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_with_return_race_4"`
					TalkToGuildfordPrecursorRuinsNothingWithReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_with_return_1"`
					CompleteThePrecursorRuinsNothingWithReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_with_return_race_2"`
					TalkToGuildfordPrecursorRuinsNothingWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_with_return_2"`
					CompleteThePrecursorRuinsNothingWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_with_return_race_3"`
					TalkToGuildfordPrecursorRuinsNothingWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_with_return_3"`
					CompleteThePrecursorRuinsNothingWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_with_return_race_4"`
					TalkToGuildfordGiantMushroomNothingWithReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_with_return_2"`
					CompleteTheGiantMushroomNothingWithReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_with_return_race_3"`
					TalkToGuildfordCrystalCoreNoAbilitiesWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_with_return_4"`
					TalkToGuildfordCrystalCoreNothingWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_with_return_4"`
					TalkToGuildfordCrystalCoreNothingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_3"`
					CompleteTheCrystalCoreNothingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_4"`
					TalkToGuildfordCrystalCoreNothingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_4"`
					TalkToGuildfordGiantMushroomNothingWithReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_with_return_3"`
					CompleteTheGiantMushroomNothingWithReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_with_return_race_4"`
					TalkToGuildfordGiantMushroomNothingWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_with_return_4"`
					TalkToGuildfordPrecursorRuinsNothingWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_with_return_4"`
					TalkToGuildfordGiantMushroomAnythingWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_with_return_4"`
					TalkToGuildfordGiantMushroomNoAbilitiesWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_with_return_4"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_4"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_with_return_4"`
					TalkToGuildfordCrystalCoreNoPearlsWithReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_with_return_4"`
					TalkToRhys struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_rhys"`
					IncreaseMining12 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"increase_mining_12"`
					HotmGiveMaterials struct {
						Status               string `json:"status"`
						Progress             int    `json:"progress"`
						CompletedAt          int    `json:"completed_at"`
						Started              bool   `json:"started"`
						EnchantedRedstone    int    `json:"ENCHANTED_REDSTONE"`
						EnchantedIron        int    `json:"ENCHANTED_IRON"`
						EnchantedCoal        int    `json:"ENCHANTED_COAL"`
						EnchantedDiamond     int    `json:"ENCHANTED_DIAMOND"`
						EnchantedLapisLazuli int    `json:"ENCHANTED_LAPIS_LAZULI"`
					} `json:"hotm_give_materials"`
					Fetchur150 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-15-0"`
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
					FindRelics struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"find_relics"`
					FindUberRelics struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"find_uber_relics"`
					TalkToShaggy2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_shaggy_2"`
					TalkToArchaeologist2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_archaeologist_2"`
					TalkToGwendolyn struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_gwendolyn"`
					TalkToBraum struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_braum"`
					VisitGreaterMines struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"visit_greater_mines"`
					TalkToTheGoblinKing struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_the_goblin_king"`
					KillAutomatons struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"kill_automatons"`
					EnterDivanMines struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"enter_divan_mines"`
					FindAJungleKey struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"find_a_jungle_key"`
					MineRuby struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"mine_ruby"`
					CompleteTrialsOfJungleTemple struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_trials_of_jungle_temple"`
					MineAmethyst struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"mine_amethyst"`
					TalkToProfessorRobot struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_professor_robot"`
					FindTheGoblinQueen struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"find_the_goblin_queen"`
					MineAmber struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"mine_amber"`
					MineSapphire struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"mine_sapphire"`
					FindFourMissingPieces struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"find_four_missing_pieces"`
					MineJade struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"mine_jade"`
					Fetchur160 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-16-0"`
					Fetchur170 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-17-0"`
					Fetchur180 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-18-0"`
					Fetchur190 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-19-0"`
					Fetchur240 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-24-0"`
					Fetchur250 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-25-0"`
					Fetchur260 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-26-0"`
					Fetchur270 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-27-0"`
					Fetchur280 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-28-0"`
					Fetchur290 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-29-0"`
					Fetchur81 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-8-1"`
					Fetchur94 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-9-4"`
					Fetchur104 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-10-4"`
					Fetchur114 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-11-4"`
					Fetchur305 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-30-5"`
					Fetchur16 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-1-6"`
					Fetchur126 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-12-6"`
					Fetchur136 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-13-6"`
					Fetchur272 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-27-2"`
					Fetchur220 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-22-0"`
					Fetchur61 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-6-1"`
					Fetchur71 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-7-1"`
					Fetchur193 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-19-3"`
					Fetchur203 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-20-3"`
					Fetchur115 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-11-5"`
					Fetchur56 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-5-6"`
					Fetchur230 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-23-0"`
					Fetchur217 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-21-7"`
					Fetchur228 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-22-8"`
				} `json:"objectives"`
				Tutorial []string `json:"tutorial"`
				Quests   struct {
					CollectLog struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"collect_log"`
					ExploreHub struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"explore_hub"`
					ExploreVillage struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"explore_village"`
					TalkToLibrarian struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_librarian"`
					TalkToFarmer struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_farmer"`
					TalkToBlacksmith struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_blacksmith"`
					TalkToLumberjack struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_lumberjack"`
					TalkToAuctionMaster struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_auction_master"`
					TalkToBanker struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_banker"`
					TalkToCarpenter struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_carpenter"`
					TalkToArtist1 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_artist_1"`
					TalkToLazyMiner struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_lazy_miner"`
					IncreaseMiningSkill5 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"increase_mining_skill_5"`
					TalkToLapisMiner struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_lapis_miner"`
					TalkToFarmhand1 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_farmhand_1"`
					IncreaseFarmingSkill5 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"increase_farming_skill_5"`
					IncreaseForagingSkill5 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"increase_foraging_skill_5"`
					TalkToGustave1 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_gustave_1"`
					IncreaseCombatSkill5 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"increase_combat_skill_5"`
					TalkToRick struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_rick"`
					TalkToGuber1 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_guber_1"`
					TalkToEndDealer struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_end_dealer"`
					ReforgeItem struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"reforge_item"`
					KillDangerMobs struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"kill_danger_mobs"`
					TalkToGulliver1 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_gulliver_1"`
					TalkToGuildford1 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_guildford_1"`
					TalkToCryptKeeper1 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_crypt_keeper_1"`
					TalkToNocturne1 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_nocturne_1"`
					TalkToDusk1 struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_dusk_1"`
					TalkToRhys struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_rhys"`
					TalkToArchaeologist struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_archaeologist"`
					TalkToGwendolyn struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_gwendolyn"`
					TalkToTheGoblinKing struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_the_goblin_king"`
					KillAutomatons struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"kill_automatons"`
					EnterDivanMines struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"enter_divan_mines"`
					FindAJungleKey struct {
						Status        string `json:"status"`
						ActivatedAt   int    `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"find_a_jungle_key"`
				} `json:"quests"`
				CoinPurse                     int           `json:"coin_purse"`
				LastDeath                     int           `json:"last_death"`
				CraftedGenerators             []interface{} `json:"crafted_generators"`
				VisitedZones                  []string      `json:"visited_zones"`
				FairySoulsCollected           int           `json:"fairy_souls_collected"`
				FairySouls                    int           `json:"fairy_souls"`
				FairyExchanges                int           `json:"fairy_exchanges"`
				FishingTreasureCaught         int           `json:"fishing_treasure_caught"`
				DeathCount                    int           `json:"death_count"`
				AchievementSpawnedIslandTypes []string      `json:"achievement_spawned_island_types"`
				SlayerBosses                  struct {
					Spider struct {
						ClaimedLevels struct {
							Level1 bool `json:"level_1"`
							Level2 bool `json:"level_2"`
							Level3 bool `json:"level_3"`
							Level6 bool `json:"level_6"`
							Level4 bool `json:"level_4"`
							Level5 bool `json:"level_5"`
							Level7 bool `json:"level_7"`
							Level8 bool `json:"level_8"`
							Level9 bool `json:"level_9"`
						} `json:"claimed_levels"`
						BossKillsTier0 int `json:"boss_kills_tier_0"`
						Xp             int `json:"xp"`
						BossKillsTier1 int `json:"boss_kills_tier_1"`
						BossKillsTier2 int `json:"boss_kills_tier_2"`
						BossKillsTier3 int `json:"boss_kills_tier_3"`
					} `json:"spider"`
					Zombie struct {
						ClaimedLevels struct {
							Level1        bool `json:"level_1"`
							Level2        bool `json:"level_2"`
							Level3        bool `json:"level_3"`
							Level4        bool `json:"level_4"`
							Level5        bool `json:"level_5"`
							Level6        bool `json:"level_6"`
							Level7        bool `json:"level_7"`
							Level8        bool `json:"level_8"`
							Level9        bool `json:"level_9"`
							Level8Special bool `json:"level_8_special"`
							Level9Special bool `json:"level_9_special"`
							Level7Special bool `json:"level_7_special"`
						} `json:"claimed_levels"`
						BossKillsTier0 int `json:"boss_kills_tier_0"`
						Xp             int `json:"xp"`
						BossKillsTier1 int `json:"boss_kills_tier_1"`
						BossKillsTier2 int `json:"boss_kills_tier_2"`
						BossKillsTier3 int `json:"boss_kills_tier_3"`
						BossKillsTier4 int `json:"boss_kills_tier_4"`
					} `json:"zombie"`
					Wolf struct {
						ClaimedLevels struct {
							Level3 bool `json:"level_3"`
							Level1 bool `json:"level_1"`
							Level2 bool `json:"level_2"`
							Level4 bool `json:"level_4"`
							Level6 bool `json:"level_6"`
							Level5 bool `json:"level_5"`
							Level7 bool `json:"level_7"`
							Level8 bool `json:"level_8"`
							Level9 bool `json:"level_9"`
						} `json:"claimed_levels"`
						BossKillsTier0 int `json:"boss_kills_tier_0"`
						Xp             int `json:"xp"`
						BossKillsTier1 int `json:"boss_kills_tier_1"`
						BossKillsTier2 int `json:"boss_kills_tier_2"`
						BossKillsTier3 int `json:"boss_kills_tier_3"`
					} `json:"wolf"`
					Enderman struct {
						ClaimedLevels struct {
							Level1 bool `json:"level_1"`
							Level2 bool `json:"level_2"`
							Level3 bool `json:"level_3"`
							Level4 bool `json:"level_4"`
							Level5 bool `json:"level_5"`
							Level6 bool `json:"level_6"`
							Level7 bool `json:"level_7"`
						} `json:"claimed_levels"`
						BossKillsTier0 int `json:"boss_kills_tier_0"`
						Xp             int `json:"xp"`
						BossKillsTier1 int `json:"boss_kills_tier_1"`
						BossKillsTier2 int `json:"boss_kills_tier_2"`
						BossKillsTier3 int `json:"boss_kills_tier_3"`
					} `json:"enderman"`
				} `json:"slayer_bosses"`
				Pets []struct {
					Type      string `json:"type"`
					Exp       int    `json:"exp"`
					Active    bool   `json:"active"`
					Tier      string `json:"tier"`
					HeldItem  string `json:"heldItem"`
					CandyUsed int    `json:"candyUsed"`
				} `json:"pets"`
				CoopInvitation struct {
					Timestamp          int    `json:"timestamp"`
					InvitedBy          string `json:"invited_by"`
					Confirmed          bool   `json:"confirmed"`
					ConfirmedTimestamp int    `json:"confirmed_timestamp"`
				} `json:"coop_invitation"`
				ExperienceSkillCombat int      `json:"experience_skill_combat"`
				ExperienceSkillMining int      `json:"experience_skill_mining"`
				UnlockedCollTiers     []string `json:"unlocked_coll_tiers"`
				FishingBag            struct {
					Type int    `json:"type"`
					Data string `json:"data"`
				} `json:"fishing_bag"`
				ExperienceSkillAlchemy int `json:"experience_skill_alchemy"`
				ExperienceSkillFarming int `json:"experience_skill_farming"`
				Collection             struct {
					Log                         int `json:"LOG"`
					Seeds                       int `json:"SEEDS"`
					Cobblestone                 int `json:"COBBLESTONE"`
					Coal                        int `json:"COAL"`
					IronIngot                   int `json:"IRON_INGOT"`
					GoldIngot                   int `json:"GOLD_INGOT"`
					INKSACK4                    int `json:"INK_SACK:4"`
					Bone                        int `json:"BONE"`
					String                      int `json:"STRING"`
					SpiderEye                   int `json:"SPIDER_EYE"`
					CarrotItem                  int `json:"CARROT_ITEM"`
					LOG1                        int `json:"LOG:1"`
					Melon                       int `json:"MELON"`
					SugarCane                   int `json:"SUGAR_CANE"`
					Pumpkin                     int `json:"PUMPKIN"`
					Pork                        int `json:"PORK"`
					Wheat                       int `json:"WHEAT"`
					PotatoItem                  int `json:"POTATO_ITEM"`
					RawChicken                  int `json:"RAW_CHICKEN"`
					Feather                     int `json:"FEATHER"`
					Leather                     int `json:"LEATHER"`
					RottenFlesh                 int `json:"ROTTEN_FLESH"`
					MushroomCollection          int `json:"MUSHROOM_COLLECTION"`
					INKSACK3                    int `json:"INK_SACK:3"`
					Cactus                      int `json:"CACTUS"`
					Rabbit                      int `json:"RABBIT"`
					Mutton                      int `json:"MUTTON"`
					Diamond                     int `json:"DIAMOND"`
					Sand                        int `json:"SAND"`
					Gravel                      int `json:"GRAVEL"`
					EnderPearl                  int `json:"ENDER_PEARL"`
					EnderStone                  int `json:"ENDER_STONE"`
					LOG2                        int `json:"LOG:2"`
					LOG21                       int `json:"LOG_2:1"`
					LOG3                        int `json:"LOG:3"`
					Redstone                    int `json:"REDSTONE"`
					Emerald                     int `json:"EMERALD"`
					SlimeBall                   int `json:"SLIME_BALL"`
					Obsidian                    int `json:"OBSIDIAN"`
					NetherStalk                 int `json:"NETHER_STALK"`
					Sulphur                     int `json:"SULPHUR"`
					BlazeRod                    int `json:"BLAZE_ROD"`
					Netherrack                  int `json:"NETHERRACK"`
					Quartz                      int `json:"QUARTZ"`
					GlowstoneDust               int `json:"GLOWSTONE_DUST"`
					MagmaCream                  int `json:"MAGMA_CREAM"`
					RAWFISH3                    int `json:"RAW_FISH:3"`
					RawFish                     int `json:"RAW_FISH"`
					RAWFISH1                    int `json:"RAW_FISH:1"`
					ClayBall                    int `json:"CLAY_BALL"`
					WaterLily                   int `json:"WATER_LILY"`
					PrismarineShard             int `json:"PRISMARINE_SHARD"`
					PrismarineCrystals          int `json:"PRISMARINE_CRYSTALS"`
					InkSack                     int `json:"INK_SACK"`
					Log2                        int `json:"LOG_2"`
					RAWFISH2                    int `json:"RAW_FISH:2"`
					Sponge                      int `json:"SPONGE"`
					GhastTear                   int `json:"GHAST_TEAR"`
					Ice                         int `json:"ICE"`
					EnchantedLapisLazuli        int `json:"ENCHANTED_LAPIS_LAZULI"`
					EnchantedDiamond            int `json:"ENCHANTED_DIAMOND"`
					SnowBall                    int `json:"SNOW_BALL"`
					DOUBLEPLANT5                int `json:"DOUBLE_PLANT:5"`
					EnchantedSnowBlock          int `json:"ENCHANTED_SNOW_BLOCK"`
					EnchantedOakLog             int `json:"ENCHANTED_OAK_LOG"`
					RedRose                     int `json:"RED_ROSE"`
					DOUBLEPLANT1                int `json:"DOUBLE_PLANT:1"`
					DoublePlant                 int `json:"DOUBLE_PLANT"`
					Wool                        int `json:"WOOL"`
					EnchantedRawFish            int `json:"ENCHANTED_RAW_FISH"`
					EnchantedRawSalmon          int `json:"ENCHANTED_RAW_SALMON"`
					EnchantedSponge             int `json:"ENCHANTED_SPONGE"`
					EnchantedPufferfish         int `json:"ENCHANTED_PUFFERFISH"`
					EnchantedPrismarineShard    int `json:"ENCHANTED_PRISMARINE_SHARD"`
					EnchantedPrismarineCrystals int `json:"ENCHANTED_PRISMARINE_CRYSTALS"`
					MithrilOre                  int `json:"MITHRIL_ORE"`
					EnchantedClownfish          int `json:"ENCHANTED_CLOWNFISH"`
					EnchantedCookedSalmon       int `json:"ENCHANTED_COOKED_SALMON"`
					EnchantedMelonBlock         int `json:"ENCHANTED_MELON_BLOCK"`
					EnchantedMelon              int `json:"ENCHANTED_MELON"`
					RawSoulflow                 int `json:"RAW_SOULFLOW"`
					HardStone                   int `json:"HARD_STONE"`
					GemstoneCollection          int `json:"GEMSTONE_COLLECTION"`
					EnchantedBrownMushroom      int `json:"ENCHANTED_BROWN_MUSHROOM"`
					EnchantedRedMushroom        int `json:"ENCHANTED_RED_MUSHROOM"`
					Soulflow                    int `json:"SOULFLOW"`
				} `json:"collection"`
				Quiver struct {
					Type int    `json:"type"`
					Data string `json:"data"`
				} `json:"quiver"`
				EnderChestContents struct {
					Type int    `json:"type"`
					Data string `json:"data"`
				} `json:"ender_chest_contents"`
				PotionBag struct {
					Type int    `json:"type"`
					Data string `json:"data"`
				} `json:"potion_bag"`
				ExperienceSkillEnchanting int `json:"experience_skill_enchanting"`
				InvContents               struct {
					Type int    `json:"type"`
					Data string `json:"data"`
				} `json:"inv_contents"`
				TalismanBag struct {
					Type int    `json:"type"`
					Data string `json:"data"`
				} `json:"talisman_bag"`
				ExperienceSkillForaging     int `json:"experience_skill_foraging"`
				ExperienceSkillCarpentry    int `json:"experience_skill_carpentry"`
				ExperienceSkillRunecrafting int `json:"experience_skill_runecrafting"`
				WardrobeEquippedSlot        int `json:"wardrobe_equipped_slot"`
				WardrobeContents            struct {
					Type int    `json:"type"`
					Data string `json:"data"`
				} `json:"wardrobe_contents"`
				ExperienceSkillTaming int `json:"experience_skill_taming"`
				SacksCounts           struct {
					RawFish                int `json:"RAW_FISH"`
					RAWFISH1               int `json:"RAW_FISH:1"`
					Diamond                int `json:"DIAMOND"`
					Redstone               int `json:"REDSTONE"`
					INKSACK4               int `json:"INK_SACK:4"`
					Emerald                int `json:"EMERALD"`
					RAWFISH3               int `json:"RAW_FISH:3"`
					WaterLily              int `json:"WATER_LILY"`
					RAWFISH2               int `json:"RAW_FISH:2"`
					InkSack                int `json:"INK_SACK"`
					ClayBall               int `json:"CLAY_BALL"`
					PrismarineCrystals     int `json:"PRISMARINE_CRYSTALS"`
					PrismarineShard        int `json:"PRISMARINE_SHARD"`
					Sponge                 int `json:"SPONGE"`
					Bone                   int `json:"BONE"`
					RottenFlesh            int `json:"ROTTEN_FLESH"`
					Coal                   int `json:"COAL"`
					BlazeRod               int `json:"BLAZE_ROD"`
					Netherrack             int `json:"NETHERRACK"`
					MagmaCream             int `json:"MAGMA_CREAM"`
					EnderPearl             int `json:"ENDER_PEARL"`
					Obsidian               int `json:"OBSIDIAN"`
					Sand                   int `json:"SAND"`
					Ice                    int `json:"ICE"`
					String                 int `json:"STRING"`
					SpiderEye              int `json:"SPIDER_EYE"`
					RevenantFlesh          int `json:"REVENANT_FLESH"`
					Cobblestone            int `json:"COBBLESTONE"`
					IronIngot              int `json:"IRON_INGOT"`
					GoldIngot              int `json:"GOLD_INGOT"`
					Quartz                 int `json:"QUARTZ"`
					GlowstoneDust          int `json:"GLOWSTONE_DUST"`
					Gravel                 int `json:"GRAVEL"`
					EnderStone             int `json:"ENDER_STONE"`
					TarantulaWeb           int `json:"TARANTULA_WEB"`
					GhastTear              int `json:"GHAST_TEAR"`
					WolfTooth              int `json:"WOLF_TOOTH"`
					SlimeBall              int `json:"SLIME_BALL"`
					Sulphur                int `json:"SULPHUR"`
					SugarCane              int `json:"SUGAR_CANE"`
					Wheat                  int `json:"WHEAT"`
					CarrotItem             int `json:"CARROT_ITEM"`
					PotatoItem             int `json:"POTATO_ITEM"`
					Pumpkin                int `json:"PUMPKIN"`
					Melon                  int `json:"MELON"`
					Seeds                  int `json:"SEEDS"`
					MushroomCollection     int `json:"MUSHROOM_COLLECTION"`
					INKSACK3               int `json:"INK_SACK:3"`
					Cactus                 int `json:"CACTUS"`
					NetherStalk            int `json:"NETHER_STALK"`
					LOG3                   int `json:"LOG:3"`
					Feather                int `json:"FEATHER"`
					RawChicken             int `json:"RAW_CHICKEN"`
					LOG1                   int `json:"LOG:1"`
					LOG2                   int `json:"LOG:2"`
					Log                    int `json:"LOG"`
					LOG21                  int `json:"LOG_2:1"`
					Log2                   int `json:"LOG_2"`
					Pork                   int `json:"PORK"`
					Leather                int `json:"LEATHER"`
					Mutton                 int `json:"MUTTON"`
					Rabbit                 int `json:"RABBIT"`
					DungeonTrap            int `json:"DUNGEON_TRAP"`
					DungeonDecoy           int `json:"DUNGEON_DECOY"`
					SpiritLeap             int `json:"SPIRIT_LEAP"`
					InflatableJerry        int `json:"INFLATABLE_JERRY"`
					GreenCandy             int `json:"GREEN_CANDY"`
					PurpleCandy            int `json:"PURPLE_CANDY"`
					MithrilOre             int `json:"MITHRIL_ORE"`
					Starfall               int `json:"STARFALL"`
					TitaniumOre            int `json:"TITANIUM_ORE"`
					Treasurite             int `json:"TREASURITE"`
					RuneWhiteSpiral1       int `json:"RUNE_WHITE_SPIRAL_1"`
					RuneBlood21            int `json:"RUNE_BLOOD_2_1"`
					RuneHearts1            int `json:"RUNE_HEARTS_1"`
					RuneGem1               int `json:"RUNE_GEM_1"`
					RuneGolden1            int `json:"RUNE_GOLDEN_1"`
					RuneTidal1             int `json:"RUNE_TIDAL_1"`
					RuneClouds1            int `json:"RUNE_CLOUDS_1"`
					RuneSparkling1         int `json:"RUNE_SPARKLING_1"`
					RuneSmokey1            int `json:"RUNE_SMOKEY_1"`
					RuneZap1               int `json:"RUNE_ZAP_1"`
					RuneHot1               int `json:"RUNE_HOT_1"`
					RuneSpirit1            int `json:"RUNE_SPIRIT_1"`
					RuneIce1               int `json:"RUNE_ICE_1"`
					RuneFireSpiral1        int `json:"RUNE_FIRE_SPIRAL_1"`
					RuneMagic1             int `json:"RUNE_MAGIC_1"`
					RuneWake1              int `json:"RUNE_WAKE_1"`
					RuneLightning1         int `json:"RUNE_LIGHTNING_1"`
					RuneRainbow1           int `json:"RUNE_RAINBOW_1"`
					RuneLava1              int `json:"RUNE_LAVA_1"`
					RuneSnow1              int `json:"RUNE_SNOW_1"`
					RuneGem3               int `json:"RUNE_GEM_3"`
					RuneLava3              int `json:"RUNE_LAVA_3"`
					RuneSnow3              int `json:"RUNE_SNOW_3"`
					RuneWhiteSpiral3       int `json:"RUNE_WHITE_SPIRAL_3"`
					RuneHot3               int `json:"RUNE_HOT_3"`
					RuneBite3              int `json:"RUNE_BITE_3"`
					RuneIce3               int `json:"RUNE_ICE_3"`
					RuneMagic3             int `json:"RUNE_MAGIC_3"`
					RuneHearts3            int `json:"RUNE_HEARTS_3"`
					RuneZap3               int `json:"RUNE_ZAP_3"`
					RuneRedstone3          int `json:"RUNE_REDSTONE_3"`
					RuneBlood23            int `json:"RUNE_BLOOD_2_3"`
					RuneSparkling3         int `json:"RUNE_SPARKLING_3"`
					RuneSpirit3            int `json:"RUNE_SPIRIT_3"`
					RuneGolden3            int `json:"RUNE_GOLDEN_3"`
					RuneWake3              int `json:"RUNE_WAKE_3"`
					RuneClouds3            int `json:"RUNE_CLOUDS_3"`
					RuneZombieSlayer3      int `json:"RUNE_ZOMBIE_SLAYER_3"`
					RuneRainbow3           int `json:"RUNE_RAINBOW_3"`
					RuneWake2              int `json:"RUNE_WAKE_2"`
					RuneZap2               int `json:"RUNE_ZAP_2"`
					RuneSnake2             int `json:"RUNE_SNAKE_2"`
					RuneClouds2            int `json:"RUNE_CLOUDS_2"`
					RuneZombieSlayer2      int `json:"RUNE_ZOMBIE_SLAYER_2"`
					RuneHot2               int `json:"RUNE_HOT_2"`
					RuneBlood22            int `json:"RUNE_BLOOD_2_2"`
					RuneGem2               int `json:"RUNE_GEM_2"`
					RuneHearts2            int `json:"RUNE_HEARTS_2"`
					RuneSparkling2         int `json:"RUNE_SPARKLING_2"`
					RuneIce2               int `json:"RUNE_ICE_2"`
					RuneGolden2            int `json:"RUNE_GOLDEN_2"`
					RuneLightning2         int `json:"RUNE_LIGHTNING_2"`
					RuneRainbow2           int `json:"RUNE_RAINBOW_2"`
					RuneFireSpiral2        int `json:"RUNE_FIRE_SPIRAL_2"`
					RuneSnow2              int `json:"RUNE_SNOW_2"`
					RuneWhiteSpiral2       int `json:"RUNE_WHITE_SPIRAL_2"`
					RuneMusic2             int `json:"RUNE_MUSIC_2"`
					RuneCouture3           int `json:"RUNE_COUTURE_3"`
					RuneSnake3             int `json:"RUNE_SNAKE_3"`
					RuneCouture1           int `json:"RUNE_COUTURE_1"`
					RuneZombieSlayer1      int `json:"RUNE_ZOMBIE_SLAYER_1"`
					RuneBite1              int `json:"RUNE_BITE_1"`
					EnchantedSpiderEye     int `json:"ENCHANTED_SPIDER_EYE"`
					EnchantedString        int `json:"ENCHANTED_STRING"`
					RuneSnake1             int `json:"RUNE_SNAKE_1"`
					EnchantedBone          int `json:"ENCHANTED_BONE"`
					RuneRedstone1          int `json:"RUNE_REDSTONE_1"`
					RuneMusic1             int `json:"RUNE_MUSIC_1"`
					EnchantedEnderPearl    int `json:"ENCHANTED_ENDER_PEARL"`
					EnchantedMelon         int `json:"ENCHANTED_MELON"`
					EnchantedRottenFlesh   int `json:"ENCHANTED_ROTTEN_FLESH"`
					EnchantedCarrot        int `json:"ENCHANTED_CARROT"`
					EnchantedPotato        int `json:"ENCHANTED_POTATO"`
					EnchantedOakLog        int `json:"ENCHANTED_OAK_LOG"`
					EnchantedBirchLog      int `json:"ENCHANTED_BIRCH_LOG"`
					RedMushroom            int `json:"RED_MUSHROOM"`
					BrownMushroom          int `json:"BROWN_MUSHROOM"`
					WerewolfSkin           int `json:"WEREWOLF_SKIN"`
					RawBeef                int `json:"RAW_BEEF"`
					RabbitFoot             int `json:"RABBIT_FOOT"`
					RabbitHide             int `json:"RABBIT_HIDE"`
					Ectoplasm              int `json:"ECTOPLASM"`
					PumpkinGuts            int `json:"PUMPKIN_GUTS"`
					SnowBall               int `json:"SNOW_BALL"`
					RuneMagic2             int `json:"RUNE_MAGIC_2"`
					SnowBlock              int `json:"SNOW_BLOCK"`
					WhiteGift              int `json:"WHITE_GIFT"`
					IceHunk                int `json:"ICE_HUNK"`
					BlueIceHunk            int `json:"BLUE_ICE_HUNK"`
					RedGift                int `json:"RED_GIFT"`
					GreenGift              int `json:"GREEN_GIFT"`
					PackedIce              int `json:"PACKED_ICE"`
					EnchantedGunpowder     int `json:"ENCHANTED_GUNPOWDER"`
					RuneRedstone2          int `json:"RUNE_REDSTONE_2"`
					RuneSpirit2            int `json:"RUNE_SPIRIT_2"`
					RuneLava2              int `json:"RUNE_LAVA_2"`
					RuneBite2              int `json:"RUNE_BITE_2"`
					RuneLightning3         int `json:"RUNE_LIGHTNING_3"`
					EnchantedSugar         int `json:"ENCHANTED_SUGAR"`
					EnchantedCocoa         int `json:"ENCHANTED_COCOA"`
					NullSphere             int `json:"NULL_SPHERE"`
					RuneEndersnake1        int `json:"RUNE_ENDERSNAKE_1"`
					RuneEndersnake2        int `json:"RUNE_ENDERSNAKE_2"`
					RuneDragon1            int `json:"RUNE_DRAGON_1"`
					EnchantedNetherStalk   int `json:"ENCHANTED_NETHER_STALK"`
					EnchantedRedMushroom   int `json:"ENCHANTED_RED_MUSHROOM"`
					EnchantedBrownMushroom int `json:"ENCHANTED_BROWN_MUSHROOM"`
					HardStone              int `json:"HARD_STONE"`
					RoughRubyGem           int `json:"ROUGH_RUBY_GEM"`
					RoughAmethystGem       int `json:"ROUGH_AMETHYST_GEM"`
					RoughSapphireGem       int `json:"ROUGH_SAPPHIRE_GEM"`
					RoughJadeGem           int `json:"ROUGH_JADE_GEM"`
					RoughTopazGem          int `json:"ROUGH_TOPAZ_GEM"`
					RoughAmberGem          int `json:"ROUGH_AMBER_GEM"`
					FlawedAmberGem         int `json:"FLAWED_AMBER_GEM"`
					FlawedSapphireGem      int `json:"FLAWED_SAPPHIRE_GEM"`
					RoughJasperGem         int `json:"ROUGH_JASPER_GEM"`
					FlawedAmethystGem      int `json:"FLAWED_AMETHYST_GEM"`
					FlawedTopazGem         int `json:"FLAWED_TOPAZ_GEM"`
					FlawedRubyGem          int `json:"FLAWED_RUBY_GEM"`
					FlawedJasperGem        int `json:"FLAWED_JASPER_GEM"`
					EnchantedHardStone     int `json:"ENCHANTED_HARD_STONE"`
					EnchantedCobblestone   int `json:"ENCHANTED_COBBLESTONE"`
					FlawedJadeGem          int `json:"FLAWED_JADE_GEM"`
					EnchantedIron          int `json:"ENCHANTED_IRON"`
					EnchantedCoal          int `json:"ENCHANTED_COAL"`
					FineTopazGem           int `json:"FINE_TOPAZ_GEM"`
					FlawlessTopazGem       int `json:"FLAWLESS_TOPAZ_GEM"`
					PerfectTopazGem        int `json:"PERFECT_TOPAZ_GEM"`
					EnchantedRedstone      int `json:"ENCHANTED_REDSTONE"`
					FineJadeGem            int `json:"FINE_JADE_GEM"`
					EnchantedMithril       int `json:"ENCHANTED_MITHRIL"`
					EnchantedTitanium      int `json:"ENCHANTED_TITANIUM"`
					FineRubyGem            int `json:"FINE_RUBY_GEM"`
					EnchantedLapisLazuli   int `json:"ENCHANTED_LAPIS_LAZULI"`
					FineAmethystGem        int `json:"FINE_AMETHYST_GEM"`
					FlawlessAmethystGem    int `json:"FLAWLESS_AMETHYST_GEM"`
					PerfectAmethystGem     int `json:"PERFECT_AMETHYST_GEM"`
					FlawlessRubyGem        int `json:"FLAWLESS_RUBY_GEM"`
					PerfectRubyGem         int `json:"PERFECT_RUBY_GEM"`
					FlawlessJadeGem        int `json:"FLAWLESS_JADE_GEM"`
					PerfectJadeGem         int `json:"PERFECT_JADE_GEM"`
					FineAmberGem           int `json:"FINE_AMBER_GEM"`
					FlawlessAmberGem       int `json:"FLAWLESS_AMBER_GEM"`
					PerfectAmberGem        int `json:"PERFECT_AMBER_GEM"`
					FineSapphireGem        int `json:"FINE_SAPPHIRE_GEM"`
					FineJasperGem          int `json:"FINE_JASPER_GEM"`
					FlawlessSapphireGem    int `json:"FLAWLESS_SAPPHIRE_GEM"`
					PerfectSapphireGem     int `json:"PERFECT_SAPPHIRE_GEM"`
					EnchantedObsidian      int `json:"ENCHANTED_OBSIDIAN"`
					EnchantedGold          int `json:"ENCHANTED_GOLD"`
					EnchantedQuartz        int `json:"ENCHANTED_QUARTZ"`
					EnchantedGlowstoneDust int `json:"ENCHANTED_GLOWSTONE_DUST"`
					FlawlessJasperGem      int `json:"FLAWLESS_JASPER_GEM"`
					PerfectJasperGem       int `json:"PERFECT_JASPER_GEM"`
					RuneEndersnake3        int `json:"RUNE_ENDERSNAKE_3"`
					RuneJerry3             int `json:"RUNE_JERRY_3"`
				} `json:"sacks_counts"`
				ExperienceSkillFishing int `json:"experience_skill_fishing"`
				CandyInventoryContents struct {
					Type int    `json:"type"`
					Data string `json:"data"`
				} `json:"candy_inventory_contents"`
				SlayerQuest struct {
					Type            string `json:"type"`
					Tier            int    `json:"tier"`
					StartTimestamp  int    `json:"start_timestamp"`
					CompletionState int    `json:"completion_state"`
					CombatXp        int    `json:"combat_xp"`
					RecentMobKills  []struct {
						Xp        int `json:"xp"`
						Timestamp int `json:"timestamp"`
					} `json:"recent_mob_kills"`
					LastKilledMobIsland   string `json:"last_killed_mob_island"`
					XpOnLastFollowerSpawn int    `json:"xp_on_last_follower_spawn"`
					SpawnTimestamp        int    `json:"spawn_timestamp"`
					KillTimestamp         int    `json:"kill_timestamp"`
				} `json:"slayer_quest"`
				Dungeons struct {
					DungeonTypes struct {
						Catacombs struct {
							TimesPlayed struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"times_played"`
							Experience      int `json:"experience"`
							TierCompletions struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"tier_completions"`
							FastestTime struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"fastest_time"`
							BestRuns struct {
								Num0 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"0"`
								Num1 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"1"`
								Num2 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"2"`
								Num3 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"3"`
								Num4 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"4"`
								Num5 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"5"`
								Num6 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"6"`
								Num7 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"7"`
							} `json:"best_runs"`
							BestScore struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"best_score"`
							MobsKilled struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"mobs_killed"`
							MostMobsKilled struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"most_mobs_killed"`
							MostDamageBerserk struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"most_damage_berserk"`
							MostHealing struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"most_healing"`
							WatcherKills struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"watcher_kills"`
							MostDamageHealer struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"most_damage_healer"`
							HighestTierCompleted int `json:"highest_tier_completed"`
							MostDamageMage       struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"most_damage_mage"`
							FastestTimeS struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"fastest_time_s"`
							FastestTimeSPlus struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"fastest_time_s_plus"`
							MostDamageArcher struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num7 int `json:"7"`
							} `json:"most_damage_archer"`
							MostDamageTank struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"most_damage_tank"`
							MilestoneCompletions struct {
								Num0 int `json:"0"`
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
								Num7 int `json:"7"`
							} `json:"milestone_completions"`
						} `json:"catacombs"`
						MasterCatacombs struct {
							TierCompletions struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
							} `json:"tier_completions"`
							MilestoneCompletions struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
							} `json:"milestone_completions"`
							HighestTierCompleted int `json:"highest_tier_completed"`
							FastestTime          struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
							} `json:"fastest_time"`
							FastestTimeS struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
							} `json:"fastest_time_s"`
							BestRuns struct {
								Num1 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"1"`
								Num2 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"2"`
								Num3 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"3"`
								Num4 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"4"`
								Num5 []struct {
									Timestamp        int           `json:"timestamp"`
									ScoreExploration int           `json:"score_exploration"`
									ScoreSpeed       int           `json:"score_speed"`
									ScoreSkill       int           `json:"score_skill"`
									ScoreBonus       int           `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      int           `json:"elapsed_time"`
									DamageDealt      int           `json:"damage_dealt"`
									Deaths           int           `json:"deaths"`
									MobsKilled       int           `json:"mobs_killed"`
									SecretsFound     int           `json:"secrets_found"`
									DamageMitigated  int           `json:"damage_mitigated"`
									AllyHealing      int           `json:"ally_healing"`
								} `json:"5"`
							} `json:"best_runs"`
							BestScore struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
							} `json:"best_score"`
							MobsKilled struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
							} `json:"mobs_killed"`
							MostMobsKilled struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
							} `json:"most_mobs_killed"`
							MostDamageTank struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
							} `json:"most_damage_tank"`
							MostHealing struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
								Num6 int `json:"6"`
							} `json:"most_healing"`
							MostDamageMage struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
								Num4 int `json:"4"`
								Num5 int `json:"5"`
							} `json:"most_damage_mage"`
							MostDamageBerserk struct {
								Num1 int `json:"1"`
								Num2 int `json:"2"`
								Num3 int `json:"3"`
							} `json:"most_damage_berserk"`
							FastestTimeSPlus struct {
								Num1 int `json:"1"`
								Num3 int `json:"3"`
								Num5 int `json:"5"`
							} `json:"fastest_time_s_plus"`
							MostDamageArcher struct {
								Num1 int `json:"1"`
								Num3 int `json:"3"`
							} `json:"most_damage_archer"`
							MostDamageHealer struct {
								Num3 int `json:"3"`
								Num4 int `json:"4"`
							} `json:"most_damage_healer"`
						} `json:"master_catacombs"`
					} `json:"dungeon_types"`
					PlayerClasses struct {
						Healer struct {
							Experience int `json:"experience"`
						} `json:"healer"`
						Mage struct {
							Experience int `json:"experience"`
						} `json:"mage"`
						Berserk struct {
							Experience int `json:"experience"`
						} `json:"berserk"`
						Archer struct {
							Experience int `json:"experience"`
						} `json:"archer"`
						Tank struct {
							Experience int `json:"experience"`
						} `json:"tank"`
					} `json:"player_classes"`
					DungeonJournal struct {
						JournalEntries struct {
							KaryllesDiary      []interface{} `json:"karylles_diary"`
							TheStudy           []interface{} `json:"the_study"`
							ExpeditionVolume1  []interface{} `json:"expedition_volume_1"`
							UncannyRemains     []interface{} `json:"uncanny_remains"`
							ExpeditionVolume2  []interface{} `json:"expedition_volume_2"`
							GrimAdversity      []interface{} `json:"grim_adversity"`
							ExpeditionVolume3  []interface{} `json:"expedition_volume_3"`
							ExpeditionVolume4  []interface{} `json:"expedition_volume_4"`
							TheWalls           []interface{} `json:"the_walls"`
							TheEye             []interface{} `json:"the_eye"`
							Aftermath          []interface{} `json:"aftermath"`
							TheApprentice      []interface{} `json:"the_apprentice"`
							TheFollower        []interface{} `json:"the_follower"`
							TheApprentice2     []interface{} `json:"the_apprentice_2"`
							NecronsMagicScroll []interface{} `json:"necrons_magic_scroll"`
						} `json:"journal_entries"`
					} `json:"dungeon_journal"`
					DungeonsBlahBlah     []interface{} `json:"dungeons_blah_blah"`
					SelectedDungeonClass string        `json:"selected_dungeon_class"`
				} `json:"dungeons"`
				PersonalVaultContents struct {
					Type int    `json:"type"`
					Data string `json:"data"`
				} `json:"personal_vault_contents"`
				Griffin struct {
					Burrows []struct {
						Ts    int `json:"ts"`
						X     int `json:"x"`
						Y     int `json:"y"`
						Z     int `json:"z"`
						Type  int `json:"type"`
						Tier  int `json:"tier"`
						Chain int `json:"chain"`
					} `json:"burrows"`
				} `json:"griffin"`
				Jacob2 struct {
					MedalsInv struct {
						Silver int `json:"silver"`
						Bronze int `json:"bronze"`
						Gold   int `json:"gold"`
					} `json:"medals_inv"`
					Perks struct {
						DoubleDrops     int `json:"double_drops"`
						FarmingLevelCap int `json:"farming_level_cap"`
					} `json:"perks"`
					Talked       bool          `json:"talked"`
					UniqueGolds2 []interface{} `json:"unique_golds2"`
				} `json:"jacob2"`
				Experimentation struct {
					Simon struct {
						LastAttempt int `json:"last_attempt"`
						Attempts0   int `json:"attempts_0"`
						BonusClicks int `json:"bonus_clicks"`
						LastClaimed int `json:"last_claimed"`
						Claims0     int `json:"claims_0"`
						BestScore0  int `json:"best_score_0"`
						Attempts2   int `json:"attempts_2"`
						Claims2     int `json:"claims_2"`
						BestScore2  int `json:"best_score_2"`
						Attempts5   int `json:"attempts_5"`
						Claims5     int `json:"claims_5"`
						BestScore5  int `json:"best_score_5"`
						Attempts3   int `json:"attempts_3"`
						Attempts1   int `json:"attempts_1"`
						Claims1     int `json:"claims_1"`
						BestScore1  int `json:"best_score_1"`
						Claims3     int `json:"claims_3"`
						BestScore3  int `json:"best_score_3"`
					} `json:"simon"`
					Pairings struct {
						LastClaimed int `json:"last_claimed"`
						Claims2     int `json:"claims_2"`
						BestScore2  int `json:"best_score_2"`
						LastAttempt int `json:"last_attempt"`
						Claims3     int `json:"claims_3"`
						BestScore3  int `json:"best_score_3"`
						Claims4     int `json:"claims_4"`
						BestScore4  int `json:"best_score_4"`
						Claims0     int `json:"claims_0"`
						BestScore0  int `json:"best_score_0"`
						Claims1     int `json:"claims_1"`
						BestScore1  int `json:"best_score_1"`
						Claims5     int `json:"claims_5"`
						BestScore5  int `json:"best_score_5"`
					} `json:"pairings"`
					Numbers struct {
						LastAttempt int `json:"last_attempt"`
						Attempts3   int `json:"attempts_3"`
						Attempts2   int `json:"attempts_2"`
						BonusClicks int `json:"bonus_clicks"`
						LastClaimed int `json:"last_claimed"`
						Claims2     int `json:"claims_2"`
						BestScore2  int `json:"best_score_2"`
						Attempts1   int `json:"attempts_1"`
						Claims1     int `json:"claims_1"`
						BestScore1  int `json:"best_score_1"`
						Claims3     int `json:"claims_3"`
						BestScore3  int `json:"best_score_3"`
					} `json:"numbers"`
					ClaimsResets          int `json:"claims_resets"`
					ClaimsResetsTimestamp int `json:"claims_resets_timestamp"`
				} `json:"experimentation"`
				BackpackContents struct {
					Num0 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"0"`
					Num1 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"1"`
					Num2 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"2"`
					Num3 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"3"`
					Num4 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"4"`
					Num5 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"5"`
					Num6 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"6"`
					Num7 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"7"`
					Num8 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"8"`
					Num9 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"9"`
					Num10 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"10"`
					Num11 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"11"`
					Num12 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"12"`
					Num13 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"13"`
					Num14 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"14"`
					Num15 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"15"`
					Num16 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"16"`
					Num17 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"17"`
				} `json:"backpack_contents"`
				BackpackIcons struct {
					Num0 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"0"`
					Num1 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"1"`
					Num2 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"2"`
					Num3 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"3"`
					Num4 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"4"`
					Num5 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"5"`
					Num6 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"6"`
					Num7 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"7"`
					Num8 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"8"`
					Num9 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"9"`
					Num10 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"10"`
					Num11 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"11"`
					Num12 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"12"`
					Num13 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"13"`
					Num14 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"14"`
					Num15 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"15"`
					Num16 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"16"`
					Num17 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"17"`
				} `json:"backpack_icons"`
				Perks struct {
					PermanentHealth       int `json:"permanent_health"`
					CatacombsBossLuck     int `json:"catacombs_boss_luck"`
					CatacombsDefense      int `json:"catacombs_defense"`
					CatacombsStrength     int `json:"catacombs_strength"`
					CatacombsHealth       int `json:"catacombs_health"`
					CatacombsCritDamage   int `json:"catacombs_crit_damage"`
					CatacombsIntelligence int `json:"catacombs_intelligence"`
					PermanentDefense      int `json:"permanent_defense"`
					PermanentSpeed        int `json:"permanent_speed"`
					PermanentIntelligence int `json:"permanent_intelligence"`
					PermanentStrength     int `json:"permanent_strength"`
					ForbiddenBlessing     int `json:"forbidden_blessing"`
					CatacombsLooting      int `json:"catacombs_looting"`
				} `json:"perks"`
				ActiveEffects []struct {
					Effect         string        `json:"effect"`
					Level          int           `json:"level"`
					Modifiers      []interface{} `json:"modifiers"`
					TicksRemaining int           `json:"ticks_remaining"`
					Infinite       bool          `json:"infinite"`
				} `json:"active_effects"`
				VisitedModes []interface{} `json:"visited_modes"`
				HarpQuest    struct {
					SelectedSong                          string `json:"selected_song"`
					SelectedSongEpoch                     int    `json:"selected_song_epoch"`
					SongHymnJoyBestCompletion             int    `json:"song_hymn_joy_best_completion"`
					SongHymnJoyCompletions                int    `json:"song_hymn_joy_completions"`
					SongHymnJoyPerfectCompletions         int    `json:"song_hymn_joy_perfect_completions"`
					SongFrereJacquesBestCompletion        int    `json:"song_frere_jacques_best_completion"`
					SongFrereJacquesCompletions           int    `json:"song_frere_jacques_completions"`
					SongFrereJacquesPerfectCompletions    int    `json:"song_frere_jacques_perfect_completions"`
					SongAmazingGraceBestCompletion        int    `json:"song_amazing_grace_best_completion"`
					SongAmazingGraceCompletions           int    `json:"song_amazing_grace_completions"`
					SongAmazingGracePerfectCompletions    int    `json:"song_amazing_grace_perfect_completions"`
					SongBrahmsBestCompletion              int    `json:"song_brahms_best_completion"`
					SongBrahmsCompletions                 int    `json:"song_brahms_completions"`
					SongBrahmsPerfectCompletions          int    `json:"song_brahms_perfect_completions"`
					SongHappyBirthdayBestCompletion       int    `json:"song_happy_birthday_best_completion"`
					SongHappyBirthdayCompletions          int    `json:"song_happy_birthday_completions"`
					SongHappyBirthdayPerfectCompletions   int    `json:"song_happy_birthday_perfect_completions"`
					SongGreensleevesBestCompletion        int    `json:"song_greensleeves_best_completion"`
					SongGreensleevesCompletions           int    `json:"song_greensleeves_completions"`
					SongGreensleevesPerfectCompletions    int    `json:"song_greensleeves_perfect_completions"`
					SongJeopardyBestCompletion            int    `json:"song_jeopardy_best_completion"`
					SongJeopardyCompletions               int    `json:"song_jeopardy_completions"`
					SongJeopardyPerfectCompletions        int    `json:"song_jeopardy_perfect_completions"`
					SongMinuetBestCompletion              int    `json:"song_minuet_best_completion"`
					SongMinuetCompletions                 int    `json:"song_minuet_completions"`
					SongMinuetPerfectCompletions          int    `json:"song_minuet_perfect_completions"`
					SongJoyWorldBestCompletion            int    `json:"song_joy_world_best_completion"`
					SongJoyWorldCompletions               int    `json:"song_joy_world_completions"`
					SongJoyWorldPerfectCompletions        int    `json:"song_joy_world_perfect_completions"`
					SongPureImaginationBestCompletion     int    `json:"song_pure_imagination_best_completion"`
					SongPureImaginationCompletions        int    `json:"song_pure_imagination_completions"`
					SongPureImaginationPerfectCompletions int    `json:"song_pure_imagination_perfect_completions"`
					SongVieEnRoseBestCompletion           int    `json:"song_vie_en_rose_best_completion"`
					SongVieEnRoseCompletions              int    `json:"song_vie_en_rose_completions"`
					SongVieEnRosePerfectCompletions       int    `json:"song_vie_en_rose_perfect_completions"`
					ClaimedTalisman                       bool   `json:"claimed_talisman"`
					SongFireAndFlamesBestCompletion       int    `json:"song_fire_and_flames_best_completion"`
					SongFireAndFlamesCompletions          int    `json:"song_fire_and_flames_completions"`
					SongFireAndFlamesPerfectCompletions   int    `json:"song_fire_and_flames_perfect_completions"`
					SongPachelbelBestCompletion           int    `json:"song_pachelbel_best_completion"`
					SongPachelbelCompletions              int    `json:"song_pachelbel_completions"`
				} `json:"harp_quest"`
				PausedEffects []struct {
					Effect         string        `json:"effect"`
					Level          int           `json:"level"`
					Modifiers      []interface{} `json:"modifiers"`
					TicksRemaining int           `json:"ticks_remaining"`
					Infinite       bool          `json:"infinite"`
				} `json:"paused_effects"`
				DisabledPotionEffects []interface{} `json:"disabled_potion_effects"`
				MiningCore            struct {
					Nodes struct {
						MiningSpeed          int  `json:"mining_speed"`
						Special0             int  `json:"special_0"`
						MiningFortune        int  `json:"mining_fortune"`
						TitaniumInsanium     int  `json:"titanium_insanium"`
						MiningSpeedBoost     int  `json:"mining_speed_boost"`
						ForgeTime            int  `json:"forge_time"`
						DailyPowder          int  `json:"daily_powder"`
						EfficientMiner       int  `json:"efficient_miner"`
						MiningExperience     int  `json:"mining_experience"`
						MiningMadness        int  `json:"mining_madness"`
						DailyEffect          int  `json:"daily_effect"`
						ExperienceOrbs       int  `json:"experience_orbs"`
						FrontLoaded          int  `json:"front_loaded"`
						FallenStarBonus      int  `json:"fallen_star_bonus"`
						RandomEvent          int  `json:"random_event"`
						GoblinKiller         int  `json:"goblin_killer"`
						Mole                 int  `json:"mole"`
						Fortunate            int  `json:"fortunate"`
						Professional         int  `json:"professional"`
						PrecisionMining      int  `json:"precision_mining"`
						GreatExplorer        int  `json:"great_explorer"`
						LonesomeMiner        int  `json:"lonesome_miner"`
						MiningSpeed2         int  `json:"mining_speed_2"`
						PowderBuff           int  `json:"powder_buff"`
						VeinSeeker           int  `json:"vein_seeker"`
						ToggleFortunate      bool `json:"toggle_fortunate"`
						ToggleMiningSpeed    bool `json:"toggle_mining_speed"`
						MiningFortune2       int  `json:"mining_fortune_2"`
						ToggleEfficientMiner bool `json:"toggle_efficient_miner"`
					} `json:"nodes"`
					ReceivedFreeTier bool `json:"received_free_tier"`
					Crystals         struct {
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
					PowderMithril               int    `json:"powder_mithril"`
					PowderMithrilTotal          int    `json:"powder_mithril_total"`
					Experience                  int    `json:"experience"`
					TokensSpent                 int    `json:"tokens_spent"`
					PowderSpentMithril          int    `json:"powder_spent_mithril"`
					RetroactiveTier2Token       bool   `json:"retroactive_tier2_token"`
					Tokens                      int    `json:"tokens"`
					SelectedPickaxeAbility      string `json:"selected_pickaxe_ability"`
					DailyOresMinedDayMithrilOre int    `json:"daily_ores_mined_day_mithril_ore"`
					DailyOresMinedMithrilOre    int    `json:"daily_ores_mined_mithril_ore"`
					LastReset                   int    `json:"last_reset"`
					GreaterMinesLastAccess      int    `json:"greater_mines_last_access"`
					Biomes                      struct {
						Dwarven struct {
							StatuesPlaced []interface{} `json:"statues_placed"`
						} `json:"dwarven"`
						Precursor struct {
							PartsDelivered []interface{} `json:"parts_delivered"`
						} `json:"precursor"`
						Goblin struct {
							KingQuestActive     bool `json:"king_quest_active"`
							KingQuestsCompleted int  `json:"king_quests_completed"`
						} `json:"goblin"`
					} `json:"biomes"`
					PowderGemstone                int    `json:"powder_gemstone"`
					PowderGemstoneTotal           int    `json:"powder_gemstone_total"`
					DailyOresMinedDay             int    `json:"daily_ores_mined_day"`
					DailyOresMined                int    `json:"daily_ores_mined"`
					DailyOresMinedDayGemstone     int    `json:"daily_ores_mined_day_gemstone"`
					DailyOresMinedGemstone        int    `json:"daily_ores_mined_gemstone"`
					PowderSpentGemstone           int    `json:"powder_spent_gemstone"`
					CurrentDailyEffect            string `json:"current_daily_effect"`
					CurrentDailyEffectLastChanged int    `json:"current_daily_effect_last_changed"`
				} `json:"mining_core"`
				Forge struct {
					ForgeProcesses struct {
						Forge1 struct {
							Num1 struct {
								Type      string `json:"type"`
								ID        string `json:"id"`
								StartTime int    `json:"startTime"`
								Slot      int    `json:"slot"`
								Notified  bool   `json:"notified"`
							} `json:"1"`
							Num2 struct {
								Type      string `json:"type"`
								ID        string `json:"id"`
								StartTime int    `json:"startTime"`
								Slot      int    `json:"slot"`
								Notified  bool   `json:"notified"`
							} `json:"2"`
							Num3 struct {
								Type      string `json:"type"`
								ID        string `json:"id"`
								StartTime int    `json:"startTime"`
								Slot      int    `json:"slot"`
								Notified  bool   `json:"notified"`
							} `json:"3"`
							Num4 struct {
								Type      string `json:"type"`
								ID        string `json:"id"`
								StartTime int    `json:"startTime"`
								Slot      int    `json:"slot"`
								Notified  bool   `json:"notified"`
							} `json:"4"`
							Num5 struct {
								Type      string `json:"type"`
								ID        string `json:"id"`
								StartTime int    `json:"startTime"`
								Slot      int    `json:"slot"`
								Notified  bool   `json:"notified"`
							} `json:"5"`
						} `json:"forge_1"`
					} `json:"forge_processes"`
				} `json:"forge"`
				EssenceUndead  int `json:"essence_undead"`
				EssenceDiamond int `json:"essence_diamond"`
				EssenceDragon  int `json:"essence_dragon"`
				EssenceWither  int `json:"essence_wither"`
				EssenceSpider  int `json:"essence_spider"`
				TempStatBuffs  []struct {
					Stat     int    `json:"stat"`
					Key      string `json:"key"`
					Amount   int    `json:"amount"`
					ExpireAt int    `json:"expire_at"`
				} `json:"temp_stat_buffs"`
				EssenceGold           int `json:"essence_gold"`
				EssenceIce            int `json:"essence_ice"`
				ExperienceSkillSocial int `json:"experience_skill_social"`
			} `json:"uuid"`
		} `json:"members"`
		Banking struct {
			Balance      int `json:"balance"`
			Transactions []struct {
				Amount        int    `json:"amount"`
				Timestamp     int    `json:"timestamp"`
				Action        string `json:"action"`
				InitiatorName string `json:"initiator_name"`
			} `json:"transactions"`
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
				StartedMs   int    `json:"started_ms"`
				StartedBy   string `json:"started_by"`
				ClaimedMs   int    `json:"claimed_ms"`
				ClaimedBy   string `json:"claimed_by"`
				Fasttracked bool   `json:"fasttracked"`
			} `json:"upgrade_states"`
		} `json:"community_upgrades"`
	} `json:"profile"`
	Cause    string `json:"cause"`
	Throttle bool   `json:"throttle"`
}
