package structs

type SkyblockProfile struct {
	Success  bool   `json:"success"`
	Cause    string `json:"cause"`
	Throttle bool   `json:"throttle"`
	Profile  struct {
		ProfileID string `json:"profile_id"`
		Members   struct {
			UUID struct {
				LastSave float64 `json:"last_save"`
				InvArmor struct {
					Type float64 `json:"type"`
					Data string  `json:"data"`
				} `json:"inv_armor"`
				FirstJoin    float64 `json:"first_join"`
				FirstJoinHub float64 `json:"first_join_hub"`
				Stats        struct {
					Deaths                                                float64 `json:"deaths"`
					DeathsVoid                                            float64 `json:"deaths_void"`
					Kills                                                 float64 `json:"kills"`
					KillsEmeraldSlime                                     float64 `json:"kills_emerald_slime"`
					AuctionsBids                                          float64 `json:"auctions_bids"`
					AuctionsHighestBid                                    float64 `json:"auctions_highest_bid"`
					KillsZombie                                           float64 `json:"kills_zombie"`
					AuctionsWon                                           float64 `json:"auctions_won"`
					AuctionsBoughtRare                                    float64 `json:"auctions_bought_rare"`
					AuctionsGoldSpent                                     float64 `json:"auctions_gold_spent"`
					KillsChicken                                          float64 `json:"kills_chicken"`
					DeathsZombie                                          float64 `json:"deaths_zombie"`
					DeathsSkeleton                                        float64 `json:"deaths_skeleton"`
					HighestCritDamage                                     float64 `json:"highest_crit_damage"`
					KillsSkeleton                                         float64 `json:"kills_skeleton"`
					KillsSpider                                           float64 `json:"kills_spider"`
					AuctionsBoughtUncommon                                float64 `json:"auctions_bought_uncommon"`
					KillsDiamondSkeleton                                  float64 `json:"kills_diamond_skeleton"`
					KillsDiamondZombie                                    float64 `json:"kills_diamond_zombie"`
					KillsZombieVillager                                   float64 `json:"kills_zombie_villager"`
					KillsRedstonePigman                                   float64 `json:"kills_redstone_pigman"`
					KillsInvisibleCreeper                                 float64 `json:"kills_invisible_creeper"`
					KillsWitch                                            float64 `json:"kills_witch"`
					ItemsFished                                           float64 `json:"items_fished"`
					ItemsFishedNormal                                     float64 `json:"items_fished_normal"`
					KillsSeaWalker                                        float64 `json:"kills_sea_walker"`
					KillsPondSquid                                        float64 `json:"kills_pond_squid"`
					ItemsFishedLargeTreasure                              float64 `json:"items_fished_large_treasure"`
					KillsSeaGuardian                                      float64 `json:"kills_sea_guardian"`
					ItemsFishedTreasure                                   float64 `json:"items_fished_treasure"`
					KillsUnburriedZombie                                  float64 `json:"kills_unburried_zombie"`
					DeathsUnburriedZombie                                 float64 `json:"deaths_unburried_zombie"`
					KillsRuinWolf                                         float64 `json:"kills_ruin_wolf"`
					KillsHorsemanZombie                                   float64 `json:"kills_horseman_zombie"`
					KillsLapisZombie                                      float64 `json:"kills_lapis_zombie"`
					DeathsFire                                            float64 `json:"deaths_fire"`
					KillsSplitterSpider                                   float64 `json:"kills_splitter_spider"`
					KillsWeaverSpider                                     float64 `json:"kills_weaver_spider"`
					KillsVoraciousSpider                                  float64 `json:"kills_voracious_spider"`
					KillsSplitterSpiderSilverfish                         float64 `json:"kills_splitter_spider_silverfish"`
					KillsJockeyShotSilverfish                             float64 `json:"kills_jockey_shot_silverfish"`
					KillsDasherSpider                                     float64 `json:"kills_dasher_spider"`
					KillsJockeySkeleton                                   float64 `json:"kills_jockey_skeleton"`
					KillsSpiderJockey                                     float64 `json:"kills_spider_jockey"`
					KillsWitherSkeleton                                   float64 `json:"kills_wither_skeleton"`
					DeathsWitherSkeleton                                  float64 `json:"deaths_wither_skeleton"`
					KillsFireballMagmaCube                                float64 `json:"kills_fireball_magma_cube"`
					KillsRabbit                                           float64 `json:"kills_rabbit"`
					KillsSheep                                            float64 `json:"kills_sheep"`
					EndRaceBestTime                                       float64 `json:"end_race_best_time"`
					DeathsFall                                            float64 `json:"deaths_fall"`
					DeathsSpider                                          float64 `json:"deaths_spider"`
					KillsPig                                              float64 `json:"kills_pig"`
					KillsCow                                              float64 `json:"kills_cow"`
					AuctionsBoughtEpic                                    float64 `json:"auctions_bought_epic"`
					KillsEnderman                                         float64 `json:"kills_enderman"`
					KillsRandomSlime                                      float64 `json:"kills_random_slime"`
					KillsRespawningSkeleton                               float64 `json:"kills_respawning_skeleton"`
					AuctionsCreated                                       float64 `json:"auctions_created"`
					AuctionsFees                                          float64 `json:"auctions_fees"`
					AuctionsCompleted                                     float64 `json:"auctions_completed"`
					AuctionsSoldCommon                                    float64 `json:"auctions_sold_common"`
					AuctionsGoldEarned                                    float64 `json:"auctions_gold_earned"`
					KillsWatcher                                          float64 `json:"kills_watcher"`
					KillsZealotEnderman                                   float64 `json:"kills_zealot_enderman"`
					KillsObsidianWither                                   float64 `json:"kills_obsidian_wither"`
					KillsEndermite                                        float64 `json:"kills_endermite"`
					DeathsUnknown                                         float64 `json:"deaths_unknown"`
					AuctionsSoldEpic                                      float64 `json:"auctions_sold_epic"`
					KillsBatPinata                                        float64 `json:"kills_bat_pinata"`
					DeathsDrowning                                        float64 `json:"deaths_drowning"`
					KillsBlaze                                            float64 `json:"kills_blaze"`
					AuctionsSoldSpecial                                   float64 `json:"auctions_sold_special"`
					KillsGeneratorGhast                                   float64 `json:"kills_generator_ghast"`
					KillsOldWolf                                          float64 `json:"kills_old_wolf"`
					AuctionsBoughtCommon                                  float64 `json:"auctions_bought_common"`
					DeathsWolf                                            float64 `json:"deaths_wolf"`
					KillsMagmaCube                                        float64 `json:"kills_magma_cube"`
					KillsPigman                                           float64 `json:"kills_pigman"`
					KillsPackSpirit                                       float64 `json:"kills_pack_spirit"`
					KillsHowlingSpirit                                    float64 `json:"kills_howling_spirit"`
					KillsSoulOfTheAlpha                                   float64 `json:"kills_soul_of_the_alpha"`
					KillsNightRespawningSkeleton                          float64 `json:"kills_night_respawning_skeleton"`
					AuctionsSoldRare                                      float64 `json:"auctions_sold_rare"`
					HighestCriticalDamage                                 float64 `json:"highest_critical_damage"`
					KillsSeaArcher                                        float64 `json:"kills_sea_archer"`
					KillsZombieDeep                                       float64 `json:"kills_zombie_deep"`
					KillsCatfish                                          float64 `json:"kills_catfish"`
					KillsChickenDeep                                      float64 `json:"kills_chicken_deep"`
					DeathsOldWolf                                         float64 `json:"deaths_old_wolf"`
					AuctionsSoldUncommon                                  float64 `json:"auctions_sold_uncommon"`
					AuctionsSoldLegendary                                 float64 `json:"auctions_sold_legendary"`
					EnderCrystalsDestroyed                                float64 `json:"ender_crystals_destroyed"`
					KillsWiseDragon                                       float64 `json:"kills_wise_dragon"`
					KillsUnstableDragon                                   float64 `json:"kills_unstable_dragon"`
					KillsStrongDragon                                     float64 `json:"kills_strong_dragon"`
					KillsProtectorDragon                                  float64 `json:"kills_protector_dragon"`
					GiftsReceived                                         float64 `json:"gifts_received"`
					GiftsGiven                                            float64 `json:"gifts_given"`
					KillsLiquidHotMagma                                   float64 `json:"kills_liquid_hot_magma"`
					MostWinterSnowballsHit                                float64 `json:"most_winter_snowballs_hit"`
					MostWinterDamageDealt                                 float64 `json:"most_winter_damage_dealt"`
					MostWinterMagmaDamageDealt                            float64 `json:"most_winter_magma_damage_dealt"`
					DeathsPlayer                                          float64 `json:"deaths_player"`
					DeathsLiquidHotMagma                                  float64 `json:"deaths_liquid_hot_magma"`
					DeathsMagmaCube                                       float64 `json:"deaths_magma_cube"`
					KillsNightSquid                                       float64 `json:"kills_night_squid"`
					DeathsSeaLeech                                        float64 `json:"deaths_sea_leech"`
					KillsOldDragon                                        float64 `json:"kills_old_dragon"`
					DeathsStrongDragon                                    float64 `json:"deaths_strong_dragon"`
					DeathsSuperiorDragon                                  float64 `json:"deaths_superior_dragon"`
					KillsSeaLeech                                         float64 `json:"kills_sea_leech"`
					KillsBroodMotherSpider                                float64 `json:"kills_brood_mother_spider"`
					KillsBroodMotherCaveSpider                            float64 `json:"kills_brood_mother_cave_spider"`
					AuctionsNoBids                                        float64 `json:"auctions_no_bids"`
					KillsYoungDragon                                      float64 `json:"kills_young_dragon"`
					KillsSuperiorDragon                                   float64 `json:"kills_superior_dragon"`
					AuctionsBoughtLegendary                               float64 `json:"auctions_bought_legendary"`
					KillsCaveSpider                                       float64 `json:"kills_cave_spider"`
					KillsPlayer                                           float64 `json:"kills_player"`
					DungeonHubGiantMushroomAnythingNoReturnBestTime       float64 `json:"dungeon_hub_giant_mushroom_anything_no_return_best_time"`
					DungeonHubPrecursorRuinsAnythingNoReturnBestTime      float64 `json:"dungeon_hub_precursor_ruins_anything_no_return_best_time"`
					DungeonHubCrystalCoreAnythingNoReturnBestTime         float64 `json:"dungeon_hub_crystal_core_anything_no_return_best_time"`
					KillsZombieGrunt                                      float64 `json:"kills_zombie_grunt"`
					KillsSkeletonGrunt                                    float64 `json:"kills_skeleton_grunt"`
					KillsDungeonRespawningSkeleton                        float64 `json:"kills_dungeon_respawning_skeleton"`
					KillsCryptLurker                                      float64 `json:"kills_crypt_lurker"`
					KillsCryptDreadlord                                   float64 `json:"kills_crypt_dreadlord"`
					KillsCryptTankZombie                                  float64 `json:"kills_crypt_tank_zombie"`
					KillsScaredSkeleton                                   float64 `json:"kills_scared_skeleton"`
					KillsDiamondGuy                                       float64 `json:"kills_diamond_guy"`
					DeathsLostAdventurer                                  float64 `json:"deaths_lost_adventurer"`
					KillsCryptSouleater                                   float64 `json:"kills_crypt_souleater"`
					KillsSkeletonSoldier                                  float64 `json:"kills_skeleton_soldier"`
					KillsCryptUndead                                      float64 `json:"kills_crypt_undead"`
					KillsWatcherSummonUndead                              float64 `json:"kills_watcher_summon_undead"`
					KillsBonzoSummonUndead                                float64 `json:"kills_bonzo_summon_undead"`
					KillsLostAdventurer                                   float64 `json:"kills_lost_adventurer"`
					DeathsBlaze                                           float64 `json:"deaths_blaze"`
					DeathsEnderman                                        float64 `json:"deaths_enderman"`
					DeathsLapisZombie                                     float64 `json:"deaths_lapis_zombie"`
					DeathsRuinWolf                                        float64 `json:"deaths_ruin_wolf"`
					DeathsEmeraldSlime                                    float64 `json:"deaths_emerald_slime"`
					DeathsWeaverSpider                                    float64 `json:"deaths_weaver_spider"`
					DeathsDasherSpider                                    float64 `json:"deaths_dasher_spider"`
					DeathsDiamondZombie                                   float64 `json:"deaths_diamond_zombie"`
					DeathsSplitterSpider                                  float64 `json:"deaths_splitter_spider"`
					DeathsSplitterSpiderSilverfish                        float64 `json:"deaths_splitter_spider_silverfish"`
					DeathsRedstonePigman                                  float64 `json:"deaths_redstone_pigman"`
					DeathsSpiderJockey                                    float64 `json:"deaths_spider_jockey"`
					DeathsDiamondSkeleton                                 float64 `json:"deaths_diamond_skeleton"`
					DeathsFireballMagmaCube                               float64 `json:"deaths_fireball_magma_cube"`
					DeathsZombieDeep                                      float64 `json:"deaths_zombie_deep"`
					DeathsWatcher                                         float64 `json:"deaths_watcher"`
					DeathsObsidianWither                                  float64 `json:"deaths_obsidian_wither"`
					DeathsEndermite                                       float64 `json:"deaths_endermite"`
					KillsGeneratorSlime                                   float64 `json:"kills_generator_slime"`
					KillsSlime                                            float64 `json:"kills_slime"`
					KillsGhast                                            float64 `json:"kills_ghast"`
					DeathsGeneratorSlime                                  float64 `json:"deaths_generator_slime"`
					DeathsZealotEnderman                                  float64 `json:"deaths_zealot_enderman"`
					DeathsOldDragon                                       float64 `json:"deaths_old_dragon"`
					DeathsWiseDragon                                      float64 `json:"deaths_wise_dragon"`
					KillsForestIslandBat                                  float64 `json:"kills_forest_island_bat"`
					KillsMagmaCubeBoss                                    float64 `json:"kills_magma_cube_boss"`
					DeathsMagmaCubeBoss                                   float64 `json:"deaths_magma_cube_boss"`
					KillsGeneratorMagmaCube                               float64 `json:"kills_generator_magma_cube"`
					DeathsCaveSpider                                      float64 `json:"deaths_cave_spider"`
					KillsSeaWitch                                         float64 `json:"kills_sea_witch"`
					KillsCreeper                                          float64 `json:"kills_creeper"`
					KillsGuardianDefender                                 float64 `json:"kills_guardian_defender"`
					KillsDeepSeaProtector                                 float64 `json:"kills_deep_sea_protector"`
					DeathsWaterHydra                                      float64 `json:"deaths_water_hydra"`
					KillsWaterHydra                                       float64 `json:"kills_water_hydra"`
					DeathsProtectorDragon                                 float64 `json:"deaths_protector_dragon"`
					ChickenRaceBestTime                                   float64 `json:"chicken_race_best_time"`
					KillsFrozenSteve                                      float64 `json:"kills_frozen_steve"`
					KillsFrostyTheSnowman                                 float64 `json:"kills_frosty_the_snowman"`
					ChickenRaceBestTime2                                  float64 `json:"chicken_race_best_time_2"`
					KillsGuardianEmperor                                  float64 `json:"kills_guardian_emperor"`
					KillsSkeletonEmperor                                  float64 `json:"kills_skeleton_emperor"`
					KillsCarrotKing                                       float64 `json:"kills_carrot_king"`
					KillsYeti                                             float64 `json:"kills_yeti"`
					DeathsYeti                                            float64 `json:"deaths_yeti"`
					DeathsPackSpirit                                      float64 `json:"deaths_pack_spirit"`
					DeathsSoulOfTheAlpha                                  float64 `json:"deaths_soul_of_the_alpha"`
					ShredderBait                                          float64 `json:"shredder_bait"`
					ShredderFished                                        float64 `json:"shredder_fished"`
					KillsGrinch                                           float64 `json:"kills_grinch"`
					DeathsGuardianEmperor                                 float64 `json:"deaths_guardian_emperor"`
					AuctionsBoughtSpecial                                 float64 `json:"auctions_bought_special"`
					ForagingRaceBestTime                                  float64 `json:"foraging_race_best_time"`
					PetMilestoneSeaCreaturesKilled                        float64 `json:"pet_milestone_sea_creatures_killed"`
					PetMilestoneOresMined                                 float64 `json:"pet_milestone_ores_mined"`
					DeathsRevenantZombie                                  float64 `json:"deaths_revenant_zombie"`
					KillsRevenantZombie                                   float64 `json:"kills_revenant_zombie"`
					DungeonHubCrystalCoreNoPearlsNoReturnBestTime         float64 `json:"dungeon_hub_crystal_core_no_pearls_no_return_best_time"`
					DungeonHubCrystalCoreNoAbilitiesNoReturnBestTime      float64 `json:"dungeon_hub_crystal_core_no_abilities_no_return_best_time"`
					KillsCellarSpider                                     float64 `json:"kills_cellar_spider"`
					KillsSniperSkeleton                                   float64 `json:"kills_sniper_skeleton"`
					DeathsWatcherSummonUndead                             float64 `json:"deaths_watcher_summon_undead"`
					DeathsCryptLurker                                     float64 `json:"deaths_crypt_lurker"`
					KillsCorruptedProtector                               float64 `json:"kills_corrupted_protector"`
					DeathsSkeletonEmperor                                 float64 `json:"deaths_skeleton_emperor"`
					KillsHorsemanBat                                      float64 `json:"kills_horseman_bat"`
					DeathsCorruptedProtector                              float64 `json:"deaths_corrupted_protector"`
					DungeonHubPrecursorRuinsNoPearlsNoReturnBestTime      float64 `json:"dungeon_hub_precursor_ruins_no_pearls_no_return_best_time"`
					DungeonHubPrecursorRuinsAnythingWithReturnBestTime    float64 `json:"dungeon_hub_precursor_ruins_anything_with_return_best_time"`
					DungeonHubPrecursorRuinsNoPearlsWithReturnBestTime    float64 `json:"dungeon_hub_precursor_ruins_no_pearls_with_return_best_time"`
					KillsHorsemanHorse                                    float64 `json:"kills_horseman_horse"`
					KillsDungeonSecretBat                                 float64 `json:"kills_dungeon_secret_bat"`
					KillsSkeletonMaster                                   float64 `json:"kills_skeleton_master"`
					DeathsScarfWarrior                                    float64 `json:"deaths_scarf_warrior"`
					KillsScarfWarrior                                     float64 `json:"kills_scarf_warrior"`
					KillsScarfMage                                        float64 `json:"kills_scarf_mage"`
					DeathsCryptDreadlord                                  float64 `json:"deaths_crypt_dreadlord"`
					DeathsSkeletonSoldier                                 float64 `json:"deaths_skeleton_soldier"`
					KillsParasite                                         float64 `json:"kills_parasite"`
					DeathsScarf                                           float64 `json:"deaths_scarf"`
					KillsLonelySpider                                     float64 `json:"kills_lonely_spider"`
					DeathsScarfMage                                       float64 `json:"deaths_scarf_mage"`
					KillsScarfPriest                                      float64 `json:"kills_scarf_priest"`
					KillsBlazeHigherOrLower                               float64 `json:"kills_blaze_higher_or_lower"`
					KillsZombieSoldier                                    float64 `json:"kills_zombie_soldier"`
					DeathsSkeletor                                        float64 `json:"deaths_skeletor"`
					DeathsSkeletonGrunt                                   float64 `json:"deaths_skeleton_grunt"`
					KillsScarfArcher                                      float64 `json:"kills_scarf_archer"`
					DeathsCryptSouleater                                  float64 `json:"deaths_crypt_souleater"`
					DeathsSkeletonMaster                                  float64 `json:"deaths_skeleton_master"`
					KillsDungeonRespawningSkeletonSkull                   float64 `json:"kills_dungeon_respawning_skeleton_skull"`
					DeathsTrap                                            float64 `json:"deaths_trap"`
					KillsCryptUndeadPieter                                float64 `json:"kills_crypt_undead_pieter"`
					KillsCryptUndeadValentin                              float64 `json:"kills_crypt_undead_valentin"`
					KillsShadowAssassin                                   float64 `json:"kills_shadow_assassin"`
					KillsSkeletor                                         float64 `json:"kills_skeletor"`
					DeathsShadowAssassin                                  float64 `json:"deaths_shadow_assassin"`
					DeathsDeathmite                                       float64 `json:"deaths_deathmite"`
					KillsWatcherBonzo                                     float64 `json:"kills_watcher_bonzo"`
					KillsProfessorGuardianSummon                          float64 `json:"kills_professor_guardian_summon"`
					DeathsProfessor                                       float64 `json:"deaths_professor"`
					KillsZombieKnight                                     float64 `json:"kills_zombie_knight"`
					DeathsProfessorMageGuardian                           float64 `json:"deaths_professor_mage_guardian"`
					DeathsScaredSkeleton                                  float64 `json:"deaths_scared_skeleton"`
					KillsCryptUndeadChristian                             float64 `json:"kills_crypt_undead_christian"`
					DeathsDiamondGuy                                      float64 `json:"deaths_diamond_guy"`
					DungeonHubGiantMushroomNoPearlsNoReturnBestTime       float64 `json:"dungeon_hub_giant_mushroom_no_pearls_no_return_best_time"`
					KillsCryptUndeadNicholas                              float64 `json:"kills_crypt_undead_nicholas"`
					KillsCryptUndeadBernhard                              float64 `json:"kills_crypt_undead_bernhard"`
					KillsCryptUndeadFriedrich                             float64 `json:"kills_crypt_undead_friedrich"`
					KillsCryptUndeadAlexander                             float64 `json:"kills_crypt_undead_alexander"`
					KillsCryptUndeadMarius                                float64 `json:"kills_crypt_undead_marius"`
					KillsKingMidas                                        float64 `json:"kills_king_midas"`
					DungeonHubGiantMushroomAnythingWithReturnBestTime     float64 `json:"dungeon_hub_giant_mushroom_anything_with_return_best_time"`
					DungeonHubGiantMushroomNoPearlsWithReturnBestTime     float64 `json:"dungeon_hub_giant_mushroom_no_pearls_with_return_best_time"`
					DungeonHubCrystalCoreAnythingWithReturnBestTime       float64 `json:"dungeon_hub_crystal_core_anything_with_return_best_time"`
					DungeonHubCrystalCoreNoPearlsWithReturnBestTime       float64 `json:"dungeon_hub_crystal_core_no_pearls_with_return_best_time"`
					KillsDeathmite                                        float64 `json:"kills_deathmite"`
					DeathsSuffocation                                     float64 `json:"deaths_suffocation"`
					DeathsYoungDragon                                     float64 `json:"deaths_young_dragon"`
					DeathsDeepSeaProtector                                float64 `json:"deaths_deep_sea_protector"`
					KillsTarantulaSpider                                  float64 `json:"kills_tarantula_spider"`
					DeathsTarantulaSpider                                 float64 `json:"deaths_tarantula_spider"`
					MostWinterCannonballsHit                              float64 `json:"most_winter_cannonballs_hit"`
					DeathsDungeonRespawningSkeleton                       float64 `json:"deaths_dungeon_respawning_skeleton"`
					DeathsProfessorGuardianSummon                         float64 `json:"deaths_professor_guardian_summon"`
					DeathsSniperSkeleton                                  float64 `json:"deaths_sniper_skeleton"`
					DeathsWither                                          float64 `json:"deaths_wither"`
					DeathsUnstableDragon                                  float64 `json:"deaths_unstable_dragon"`
					DeathsGeneratorGhast                                  float64 `json:"deaths_generator_ghast"`
					DungeonHubGiantMushroomNoAbilitiesNoReturnBestTime    float64 `json:"dungeon_hub_giant_mushroom_no_abilities_no_return_best_time"`
					DungeonHubGiantMushroomNothingNoReturnBestTime        float64 `json:"dungeon_hub_giant_mushroom_nothing_no_return_best_time"`
					DeathsPigman                                          float64 `json:"deaths_pigman"`
					DeathsCatfish                                         float64 `json:"deaths_catfish"`
					DeathsGuardianDefender                                float64 `json:"deaths_guardian_defender"`
					DungeonHubCrystalCoreNothingNoReturnBestTime          float64 `json:"dungeon_hub_crystal_core_nothing_no_return_best_time"`
					DungeonHubCrystalCoreNoAbilitiesWithReturnBestTime    float64 `json:"dungeon_hub_crystal_core_no_abilities_with_return_best_time"`
					DungeonHubCrystalCoreNothingWithReturnBestTime        float64 `json:"dungeon_hub_crystal_core_nothing_with_return_best_time"`
					DungeonHubGiantMushroomNoAbilitiesWithReturnBestTime  float64 `json:"dungeon_hub_giant_mushroom_no_abilities_with_return_best_time"`
					DungeonHubGiantMushroomNothingWithReturnBestTime      float64 `json:"dungeon_hub_giant_mushroom_nothing_with_return_best_time"`
					DungeonHubPrecursorRuinsNoAbilitiesNoReturnBestTime   float64 `json:"dungeon_hub_precursor_ruins_no_abilities_no_return_best_time"`
					DungeonHubPrecursorRuinsNothingNoReturnBestTime       float64 `json:"dungeon_hub_precursor_ruins_nothing_no_return_best_time"`
					DungeonHubPrecursorRuinsNoAbilitiesWithReturnBestTime float64 `json:"dungeon_hub_precursor_ruins_no_abilities_with_return_best_time"`
					DungeonHubPrecursorRuinsNothingWithReturnBestTime     float64 `json:"dungeon_hub_precursor_ruins_nothing_with_return_best_time"`
					DeathsZombieVillager                                  float64 `json:"deaths_zombie_villager"`
					DeathsHowlingSpirit                                   float64 `json:"deaths_howling_spirit"`
					DeathsProfessorArcherGuardian                         float64 `json:"deaths_professor_archer_guardian"`
					DeathsSeaGuardian                                     float64 `json:"deaths_sea_guardian"`
					DeathsRespawningSkeleton                              float64 `json:"deaths_respawning_skeleton"`
					DeathsJockeyShotSilverfish                            float64 `json:"deaths_jockey_shot_silverfish"`
					KillsHeadlessHorseman                                 float64 `json:"kills_headless_horseman"`
					DeathsKingMidas                                       float64 `json:"deaths_king_midas"`
					KillsSuperArcher                                      float64 `json:"kills_super_archer"`
					KillsCryptWitherskeleton                              float64 `json:"kills_crypt_witherskeleton"`
					KillsSpiritWolf                                       float64 `json:"kills_spirit_wolf"`
					KillsSpiritBull                                       float64 `json:"kills_spirit_bull"`
					KillsSpiritRabbit                                     float64 `json:"kills_spirit_rabbit"`
					KillsSpiritChicken                                    float64 `json:"kills_spirit_chicken"`
					KillsSpiritBat                                        float64 `json:"kills_spirit_bat"`
					KillsSpiritSheep                                      float64 `json:"kills_spirit_sheep"`
					DeathsSpiritChicken                                   float64 `json:"deaths_spirit_chicken"`
					KillsSpiritMiniboss                                   float64 `json:"kills_spirit_miniboss"`
					DeathsSpiritBat                                       float64 `json:"deaths_spirit_bat"`
					KillsSuperTankZombie                                  float64 `json:"kills_super_tank_zombie"`
					KillsWatcherScarf                                     float64 `json:"kills_watcher_scarf"`
					DeathsSpiritWolf                                      float64 `json:"deaths_spirit_wolf"`
					KillsThorn                                            float64 `json:"kills_thorn"`
					DeathsSpiritMiniboss                                  float64 `json:"deaths_spirit_miniboss"`
					DeathsSpiritRabbit                                    float64 `json:"deaths_spirit_rabbit"`
					DeathsSpiritBull                                      float64 `json:"deaths_spirit_bull"`
					DeathsWatcherScarf                                    float64 `json:"deaths_watcher_scarf"`
					DeathsWatcherGuardian                                 float64 `json:"deaths_watcher_guardian"`
					DeathsZombieKnight                                    float64 `json:"deaths_zombie_knight"`
					DeathsCryptTankZombie                                 float64 `json:"deaths_crypt_tank_zombie"`
					DeathsZombieSoldier                                   float64 `json:"deaths_zombie_soldier"`
					DeathsSuperArcher                                     float64 `json:"deaths_super_archer"`
					DeathsWatcherBonzo                                    float64 `json:"deaths_watcher_bonzo"`
					DeathsLonelySpider                                    float64 `json:"deaths_lonely_spider"`
					DeathsCryptUndead                                     float64 `json:"deaths_crypt_undead"`
					MythosBurrowsDugNext                                  float64 `json:"mythos_burrows_dug_next"`
					MythosBurrowsDugNextCOMMON                            float64 `json:"mythos_burrows_dug_next_COMMON"`
					MythosBurrowsDugCombat                                float64 `json:"mythos_burrows_dug_combat"`
					MythosBurrowsDugCombatCOMMON                          float64 `json:"mythos_burrows_dug_combat_COMMON"`
					MythosKills                                           float64 `json:"mythos_kills"`
					KillsSiameseLynx                                      float64 `json:"kills_siamese_lynx"`
					KillsMinosHunter                                      float64 `json:"kills_minos_hunter"`
					MythosBurrowsDugTreasure                              float64 `json:"mythos_burrows_dug_treasure"`
					MythosBurrowsDugTreasureCOMMON                        float64 `json:"mythos_burrows_dug_treasure_COMMON"`
					MythosBurrowsChainsComplete                           float64 `json:"mythos_burrows_chains_complete"`
					MythosBurrowsChainsCompleteCOMMON                     float64 `json:"mythos_burrows_chains_complete_COMMON"`
					MythosBurrowsDugNextUNCOMMON                          float64 `json:"mythos_burrows_dug_next_UNCOMMON"`
					MythosBurrowsDugCombatUNCOMMON                        float64 `json:"mythos_burrows_dug_combat_UNCOMMON"`
					KillsMinotaur                                         float64 `json:"kills_minotaur"`
					MythosBurrowsDugTreasureUNCOMMON                      float64 `json:"mythos_burrows_dug_treasure_UNCOMMON"`
					MythosBurrowsChainsCompleteUNCOMMON                   float64 `json:"mythos_burrows_chains_complete_UNCOMMON"`
					MythosBurrowsDugNextRARE                              float64 `json:"mythos_burrows_dug_next_RARE"`
					MythosBurrowsDugCombatRARE                            float64 `json:"mythos_burrows_dug_combat_RARE"`
					KillsGaiaConstruct                                    float64 `json:"kills_gaia_construct"`
					DeathsGaiaConstruct                                   float64 `json:"deaths_gaia_construct"`
					MythosBurrowsDugNextLEGENDARY                         float64 `json:"mythos_burrows_dug_next_LEGENDARY"`
					MythosBurrowsDugCombatLEGENDARY                       float64 `json:"mythos_burrows_dug_combat_LEGENDARY"`
					DeathsMinosChampion                                   float64 `json:"deaths_minos_champion"`
					DeathsMinotaur                                        float64 `json:"deaths_minotaur"`
					MythosBurrowsDugTreasureLEGENDARY                     float64 `json:"mythos_burrows_dug_treasure_LEGENDARY"`
					MythosBurrowsChainsCompleteLEGENDARY                  float64 `json:"mythos_burrows_chains_complete_LEGENDARY"`
					KillsMinosChampion                                    float64 `json:"kills_minos_champion"`
					KillsMinosInquisitor                                  float64 `json:"kills_minos_inquisitor"`
					MythosBurrowsDugNextNull                              float64 `json:"mythos_burrows_dug_next_null"`
					MythosBurrowsDugCombatNull                            float64 `json:"mythos_burrows_dug_combat_null"`
					MythosBurrowsDugTreasureRARE                          float64 `json:"mythos_burrows_dug_treasure_RARE"`
					MythosBurrowsChainsCompleteRARE                       float64 `json:"mythos_burrows_chains_complete_RARE"`
					MythosBurrowsDugNextEPIC                              float64 `json:"mythos_burrows_dug_next_EPIC"`
					MythosBurrowsDugTreasureEPIC                          float64 `json:"mythos_burrows_dug_treasure_EPIC"`
					MythosBurrowsDugCombatEPIC                            float64 `json:"mythos_burrows_dug_combat_EPIC"`
					MythosBurrowsChainsCompleteEPIC                       float64 `json:"mythos_burrows_chains_complete_EPIC"`
					MythosBurrowsDugTreasureNull                          float64 `json:"mythos_burrows_dug_treasure_null"`
					MythosBurrowsChainsCompleteNull                       float64 `json:"mythos_burrows_chains_complete_null"`
					DeathsMinosInquisitor                                 float64 `json:"deaths_minos_inquisitor"`
					KillsNurseShark                                       float64 `json:"kills_nurse_shark"`
					KillsTigerShark                                       float64 `json:"kills_tiger_shark"`
					KillsGreatWhiteShark                                  float64 `json:"kills_great_white_shark"`
					KillsBlueShark                                        float64 `json:"kills_blue_shark"`
					DeathsGreatWhiteShark                                 float64 `json:"deaths_great_white_shark"`
					KillsTentaclees                                       float64 `json:"kills_tentaclees"`
					DeathsLividClone                                      float64 `json:"deaths_livid_clone"`
					DeathsLivid                                           float64 `json:"deaths_livid"`
					DeathsTentaclees                                      float64 `json:"deaths_tentaclees"`
					DeathsBonzoSummonUndead                               float64 `json:"deaths_bonzo_summon_undead"`
					DeathsProfessorWarriorGuardian                        float64 `json:"deaths_professor_warrior_guardian"`
					AuctionsBoughtMythic                                  float64 `json:"auctions_bought_mythic"`
					KillsWatcherLivid                                     float64 `json:"kills_watcher_livid"`
					DeathsSadanStatue                                     float64 `json:"deaths_sadan_statue"`
					KillsZombieCommander                                  float64 `json:"kills_zombie_commander"`
					KillsSadanStatue                                      float64 `json:"kills_sadan_statue"`
					KillsSadanGiant                                       float64 `json:"kills_sadan_giant"`
					KillsSadanGolem                                       float64 `json:"kills_sadan_golem"`
					KillsSkeletorPrime                                    float64 `json:"kills_skeletor_prime"`
					DeathsSadanGiant                                      float64 `json:"deaths_sadan_giant"`
					DeathsSadanGolem                                      float64 `json:"deaths_sadan_golem"`
					DeathsSadan                                           float64 `json:"deaths_sadan"`
					KillsMimic                                            float64 `json:"kills_mimic"`
					DeathsSiameseLynx                                     float64 `json:"deaths_siamese_lynx"`
					DeathsBonzo                                           float64 `json:"deaths_bonzo"`
					DeathsSkeletorPrime                                   float64 `json:"deaths_skeletor_prime"`
					KillsWraith                                           float64 `json:"kills_wraith"`
					KillsWitherGourd                                      float64 `json:"kills_wither_gourd"`
					KillsScaryJerry                                       float64 `json:"kills_scary_jerry"`
					KillsPhantomSpirit                                    float64 `json:"kills_phantom_spirit"`
					KillsTrickOrTreater                                   float64 `json:"kills_trick_or_treater"`
					KillsBatSpooky                                        float64 `json:"kills_bat_spooky"`
					KillsBattyWitch                                       float64 `json:"kills_batty_witch"`
					KillsWitchBat                                         float64 `json:"kills_witch_bat"`
					KillsScarecrow                                        float64 `json:"kills_scarecrow"`
					KillsNightmare                                        float64 `json:"kills_nightmare"`
					KillsPhantomFisherman                                 float64 `json:"kills_phantom_fisherman"`
					KillsWerewolf                                         float64 `json:"kills_werewolf"`
					KillsGrimReaper                                       float64 `json:"kills_grim_reaper"`
					AuctionsSoldMythic                                    float64 `json:"auctions_sold_mythic"`
					DeathsCryptWitherskeleton                             float64 `json:"deaths_crypt_witherskeleton"`
					KillsZombieLord                                       float64 `json:"kills_zombie_lord"`
					DeathsWitherMiner                                     float64 `json:"deaths_wither_miner"`
					KillsWitherMiner                                      float64 `json:"kills_wither_miner"`
					KillsWitherGuard                                      float64 `json:"kills_wither_guard"`
					DeathsMaxor                                           float64 `json:"deaths_maxor"`
					DeathsWitherGuard                                     float64 `json:"deaths_wither_guard"`
					KillsSkeletonLord                                     float64 `json:"kills_skeleton_lord"`
					KillsWatcherGiantLaser                                float64 `json:"kills_watcher_giant_laser"`
					KillsWatcherGiantBoulder                              float64 `json:"kills_watcher_giant_boulder"`
					KillsWatcherGiantDiamond                              float64 `json:"kills_watcher_giant_diamond"`
					KillsNecronGuard                                      float64 `json:"kills_necron_guard"`
					DeathsWatcherLivid                                    float64 `json:"deaths_watcher_livid"`
					DeathsNecronGuard                                     float64 `json:"deaths_necron_guard"`
					KillsWatcherGiantBigfoot                              float64 `json:"kills_watcher_giant_bigfoot"`
					DeathsWatcherGiantBigfoot                             float64 `json:"deaths_watcher_giant_bigfoot"`
					DeathsWatcherGiantBoulder                             float64 `json:"deaths_watcher_giant_boulder"`
					DeathsSuperTankZombie                                 float64 `json:"deaths_super_tank_zombie"`
					DeathsCrushed                                         float64 `json:"deaths_crushed"`
					DeathsArmorStand                                      float64 `json:"deaths_armor_stand"`
					KillsMayorJerryGreen                                  float64 `json:"kills_mayor_jerry_green"`
					KillsMayorJerryBlue                                   float64 `json:"kills_mayor_jerry_blue"`
					KillsMayorJerryPurple                                 float64 `json:"kills_mayor_jerry_purple"`
					KillsMayorJerryGolden                                 float64 `json:"kills_mayor_jerry_golden"`
					KillsIceWalker                                        float64 `json:"kills_ice_walker"`
					KillsGoblin                                           float64 `json:"kills_goblin"`
					DeathsCavernsGhost                                    float64 `json:"deaths_caverns_ghost"`
					KillsGoblinKnifeThrower                               float64 `json:"kills_goblin_knife_thrower"`
					KillsGoblinWeaklingMelee                              float64 `json:"kills_goblin_weakling_melee"`
					KillsGoblinWeaklingBow                                float64 `json:"kills_goblin_weakling_bow"`
					KillsTreasureHoarder                                  float64 `json:"kills_treasure_hoarder"`
					KillsGoblinCreepertamer                               float64 `json:"kills_goblin_creepertamer"`
					KillsGoblinBattler                                    float64 `json:"kills_goblin_battler"`
					KillsGoblinMurderlover                                float64 `json:"kills_goblin_murderlover"`
					KillsCavernsGhost                                     float64 `json:"kills_caverns_ghost"`
					KillsGoblinCreeper                                    float64 `json:"kills_goblin_creeper"`
					KillsGoblinGolem                                      float64 `json:"kills_goblin_golem"`
					KillsCrystalSentry                                    float64 `json:"kills_crystal_sentry"`
					KillsPowderGhast                                      float64 `json:"kills_powder_ghast"`
					DeathsGoblinMurderlover                               float64 `json:"deaths_goblin_murderlover"`
					DeathsGoblinKnifeThrower                              float64 `json:"deaths_goblin_knife_thrower"`
					DeathsIceWalker                                       float64 `json:"deaths_ice_walker"`
					DeathsCryptUndeadHypixel                              float64 `json:"deaths_crypt_undead_hypixel"`
					DeathsCryptUndeadFlameboy101                          float64 `json:"deaths_crypt_undead_flameboy101"`
					DeathsGoblinWeaklingBow                               float64 `json:"deaths_goblin_weakling_bow"`
					KillsArachneBrood                                     float64 `json:"kills_arachne_brood"`
					KillsArachneKeeper                                    float64 `json:"kills_arachne_keeper"`
					DeathsArachne                                         float64 `json:"deaths_arachne"`
					KillsArachne                                          float64 `json:"kills_arachne"`
					DeathsArachneBrood                                    float64 `json:"deaths_arachne_brood"`
					DeathsArachneKeeper                                   float64 `json:"deaths_arachne_keeper"`
					DeathsCellarSpider                                    float64 `json:"deaths_cellar_spider"`
					KillsMasterSniperSkeleton                             float64 `json:"kills_master_sniper_skeleton"`
					KillsMasterCryptTankZombie                            float64 `json:"kills_master_crypt_tank_zombie"`
					KillsMasterZombieGrunt                                float64 `json:"kills_master_zombie_grunt"`
					KillsMasterCryptLurker                                float64 `json:"kills_master_crypt_lurker"`
					KillsMasterScaredSkeleton                             float64 `json:"kills_master_scared_skeleton"`
					KillsMasterSkeletonSoldier                            float64 `json:"kills_master_skeleton_soldier"`
					KillsMasterSkeletonGrunt                              float64 `json:"kills_master_skeleton_grunt"`
					KillsMasterCryptSouleater                             float64 `json:"kills_master_crypt_souleater"`
					KillsMasterDungeonRespawningSkeleton                  float64 `json:"kills_master_dungeon_respawning_skeleton"`
					KillsMasterLostAdventurer                             float64 `json:"kills_master_lost_adventurer"`
					KillsMasterCryptDreadlord                             float64 `json:"kills_master_crypt_dreadlord"`
					KillsMasterCellarSpider                               float64 `json:"kills_master_cellar_spider"`
					KillsMasterWatcherSummonUndead                        float64 `json:"kills_master_watcher_summon_undead"`
					DeathsMasterWatcherSummonUndead                       float64 `json:"deaths_master_watcher_summon_undead"`
					KillsMasterBonzoSummonUndead                          float64 `json:"kills_master_bonzo_summon_undead"`
					DeathsMasterBonzo                                     float64 `json:"deaths_master_bonzo"`
					DeathsSpiritSheep                                     float64 `json:"deaths_spirit_sheep"`
					KillsMasterCryptUndead                                float64 `json:"kills_master_crypt_undead"`
					KillsMasterDiamondGuy                                 float64 `json:"kills_master_diamond_guy"`
					KillsMasterSkeletonMaster                             float64 `json:"kills_master_skeleton_master"`
					DeathsMasterScarfArcher                               float64 `json:"deaths_master_scarf_archer"`
					KillsMasterScarfMage                                  float64 `json:"kills_master_scarf_mage"`
					DeathsMasterScarf                                     float64 `json:"deaths_master_scarf"`
					KillsDanteGoon                                        float64 `json:"kills_dante_goon"`
					KillsDanteSlimeGoon                                   float64 `json:"kills_dante_slime_goon"`
					KillsRat                                              float64 `json:"kills_rat"`
					KillsMushroomCow                                      float64 `json:"kills_mushroom_cow"`
					KillsTrapperPig                                       float64 `json:"kills_trapper_pig"`
					KillsTrapperChicken                                   float64 `json:"kills_trapper_chicken"`
					KillsTrapperSheep                                     float64 `json:"kills_trapper_sheep"`
					KillsTrapperCow                                       float64 `json:"kills_trapper_cow"`
					KillsTrapperRabbit                                    float64 `json:"kills_trapper_rabbit"`
					DeathsMasterLostAdventurer                            float64 `json:"deaths_master_lost_adventurer"`
					DeathsZombieGrunt                                     float64 `json:"deaths_zombie_grunt"`
					DeathsGrimReaper                                      float64 `json:"deaths_grim_reaper"`
					KillsOasisSheep                                       float64 `json:"kills_oasis_sheep"`
					DeathsMasterSkeletonSoldier                           float64 `json:"deaths_master_skeleton_soldier"`
					DeathsMasterScarfWarrior                              float64 `json:"deaths_master_scarf_warrior"`
					KillsMasterParasite                                   float64 `json:"kills_master_parasite"`
					KillsMasterScarfPriest                                float64 `json:"kills_master_scarf_priest"`
					DeathsMasterCryptLurker                               float64 `json:"deaths_master_crypt_lurker"`
					KillsMasterCryptUndeadBernhard                        float64 `json:"kills_master_crypt_undead_bernhard"`
					KillsMasterScarfWarrior                               float64 `json:"kills_master_scarf_warrior"`
					KillsMasterScarfArcher                                float64 `json:"kills_master_scarf_archer"`
					DeathsMasterScarfMage                                 float64 `json:"deaths_master_scarf_mage"`
					KillsMasterSkeletor                                   float64 `json:"kills_master_skeletor"`
					DeathsMasterSkeletor                                  float64 `json:"deaths_master_skeletor"`
					KillsMasterZombieKnight                               float64 `json:"kills_master_zombie_knight"`
					KillsMasterZombieSoldier                              float64 `json:"kills_master_zombie_soldier"`
					KillsMasterLonelySpider                               float64 `json:"kills_master_lonely_spider"`
					DeathsMasterWatcherBonzo                              float64 `json:"deaths_master_watcher_bonzo"`
					KillsMasterProfessorGuardianSummon                    float64 `json:"kills_master_professor_guardian_summon"`
					DeathsMasterProfessorMageGuardian                     float64 `json:"deaths_master_professor_mage_guardian"`
					KillsMasterSuperTankZombie                            float64 `json:"kills_master_super_tank_zombie"`
					KillsMasterSpiritBat                                  float64 `json:"kills_master_spirit_bat"`
					DeathsMasterZombieSoldier                             float64 `json:"deaths_master_zombie_soldier"`
					DeathsMasterLividClone                                float64 `json:"deaths_master_livid_clone"`
					DeathsMasterSniperSkeleton                            float64 `json:"deaths_master_sniper_skeleton"`
					KillsMasterWatcherBonzo                               float64 `json:"kills_master_watcher_bonzo"`
					DeathsMasterSpiritBat                                 float64 `json:"deaths_master_spirit_bat"`
					DeathsMasterSpiritSheep                               float64 `json:"deaths_master_spirit_sheep"`
					DeathsMasterSpiritRabbit                              float64 `json:"deaths_master_spirit_rabbit"`
					DeathsMasterProfessorGuardianSummon                   float64 `json:"deaths_master_professor_guardian_summon"`
					KillsMasterCryptWitherskeleton                        float64 `json:"kills_master_crypt_witherskeleton"`
					DeathsMasterWatcherScarf                              float64 `json:"deaths_master_watcher_scarf"`
					KillsMasterCryptUndeadPieter                          float64 `json:"kills_master_crypt_undead_pieter"`
					KillsMasterSpiritWolf                                 float64 `json:"kills_master_spirit_wolf"`
					KillsMasterSpiritRabbit                               float64 `json:"kills_master_spirit_rabbit"`
					KillsMasterSpiritSheep                                float64 `json:"kills_master_spirit_sheep"`
					KillsMasterSpiritBull                                 float64 `json:"kills_master_spirit_bull"`
					DeathsMasterShadowAssassin                            float64 `json:"deaths_master_shadow_assassin"`
					DeathsMasterSpiritChicken                             float64 `json:"deaths_master_spirit_chicken"`
					KillsMasterTentaclees                                 float64 `json:"kills_master_tentaclees"`
					DeathsMasterTentaclees                                float64 `json:"deaths_master_tentaclees"`
					KillsMasterSuperArcher                                float64 `json:"kills_master_super_archer"`
					KillsMasterShadowAssassin                             float64 `json:"kills_master_shadow_assassin"`
					DeathsMasterCryptWitherskeleton                       float64 `json:"deaths_master_crypt_witherskeleton"`
					DeathsMasterCryptDreadlord                            float64 `json:"deaths_master_crypt_dreadlord"`
					DeathsMasterSkeletonMaster                            float64 `json:"deaths_master_skeleton_master"`
					DeathsMasterZombieKnight                              float64 `json:"deaths_master_zombie_knight"`
					DeathsMasterLivid                                     float64 `json:"deaths_master_livid"`
					DeathsMasterCryptSouleater                            float64 `json:"deaths_master_crypt_souleater"`
					PumpkinLauncherCount                                  float64 `json:"pumpkin_launcher_count"`
					KillsShrineChargedCreeper                             float64 `json:"kills_shrine_charged_creeper"`
					KillsShrineSkeletonHorseman                           float64 `json:"kills_shrine_skeleton_horseman"`
					DeathsMasterDiamondGuy                                float64 `json:"deaths_master_diamond_guy"`
					DeathsMasterDungeonRespawningSkeleton                 float64 `json:"deaths_master_dungeon_respawning_skeleton"`
					DeathsMasterSkeletonGrunt                             float64 `json:"deaths_master_skeleton_grunt"`
					KillsOasisRabbit                                      float64 `json:"kills_oasis_rabbit"`
					KillsMasterSpiritChicken                              float64 `json:"kills_master_spirit_chicken"`
					KillsMasterSpiritMiniboss                             float64 `json:"kills_master_spirit_miniboss"`
					DeathsMasterSpiritMiniboss                            float64 `json:"deaths_master_spirit_miniboss"`
					KillsVoidlingFanatic                                  float64 `json:"kills_voidling_fanatic"`
					KillsVoidlingExtremist                                float64 `json:"kills_voidling_extremist"`
					DeathsVoidlingExtremist                               float64 `json:"deaths_voidling_extremist"`
					DeathsVoidlingFanatic                                 float64 `json:"deaths_voidling_fanatic"`
					KillsVoidlingEnderman                                 float64 `json:"kills_voidling_enderman"`
					KillsThyst                                            float64 `json:"kills_thyst"`
					KillsSludge                                           float64 `json:"kills_sludge"`
					KillsAutomaton                                        float64 `json:"kills_automaton"`
					KillsKeyGuardian                                      float64 `json:"kills_key_guardian"`
					DeathsAutomaton                                       float64 `json:"deaths_automaton"`
					KillsTeamTreasuriteViper                              float64 `json:"kills_team_treasurite_viper"`
					KillsTeamTreasuriteSebastian                          float64 `json:"kills_team_treasurite_sebastian"`
					KillsYog                                              float64 `json:"kills_yog"`
					KillsGoblinFlamethrower                               float64 `json:"kills_goblin_flamethrower"`
					KillsTeamTreasuriteWendy                              float64 `json:"kills_team_treasurite_wendy"`
					KillsBelle                                            float64 `json:"kills_belle"`
					KillsFireBat                                          float64 `json:"kills_fire_bat"`
					KillsWorm                                             float64 `json:"kills_worm"`
					KillsTeamTreasuriteGrunt                              float64 `json:"kills_team_treasurite_grunt"`
					DeathsYog                                             float64 `json:"deaths_yog"`
					KillsSilvo                                            float64 `json:"kills_silvo"`
					DeathsKalhuikiTribeMan                                float64 `json:"deaths_kalhuiki_tribe_man"`
					DeathsKalhuikiTribeWoman                              float64 `json:"deaths_kalhuiki_tribe_woman"`
					KillsButterfly                                        float64 `json:"kills_butterfly"`
					KillsCavitak                                          float64 `json:"kills_cavitak"`
					DeathsSludge                                          float64 `json:"deaths_sludge"`
					KillsTrappedSludge                                    float64 `json:"kills_trapped_sludge"`
					KillsTeamTreasuriteCorleone                           float64 `json:"kills_team_treasurite_corleone"`
					KillsSmog                                             float64 `json:"kills_smog"`
					KillsVittomite                                        float64 `json:"kills_vittomite"`
					KillsKalhuikiElder                                    float64 `json:"kills_kalhuiki_elder"`
					KillsKalhuikiTribeMan                                 float64 `json:"kills_kalhuiki_tribe_man"`
					KillsKalhuikiYoungling                                float64 `json:"kills_kalhuiki_youngling"`
					KillsScatha                                           float64 `json:"kills_scatha"`
					DeathsMasterCryptTankZombie                           float64 `json:"deaths_master_crypt_tank_zombie"`
					KillsLavaPigman                                       float64 `json:"kills_lava_pigman"`
					KillsLavaBlaze                                        float64 `json:"kills_lava_blaze"`
					KillsKalhuikiTribeWoman                               float64 `json:"kills_kalhuiki_tribe_woman"`
					DeathsZombieCommander                                 float64 `json:"deaths_zombie_commander"`
					DeathsEntity                                          float64 `json:"deaths_entity"`
					KillsFlamingWorm                                      float64 `json:"kills_flaming_worm"`
					DeathsTeamTreasuriteCorleone                          float64 `json:"deaths_team_treasurite_corleone"`
					DeathsMasterSpiritWolf                                float64 `json:"deaths_master_spirit_wolf"`
					KillsMasterCryptUndeadValentin                        float64 `json:"kills_master_crypt_undead_valentin"`
					KillsMasterCryptUndeadNicholas                        float64 `json:"kills_master_crypt_undead_nicholas"`
					KillsMasterCryptUndeadChristian                       float64 `json:"kills_master_crypt_undead_christian"`
					KillsMasterCryptUndeadFriedrich                       float64 `json:"kills_master_crypt_undead_friedrich"`
					DeathsMasterSadanGiant                                float64 `json:"deaths_master_sadan_giant"`
				} `json:"stats"`
				Objectives struct {
					CollectLog struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"collect_log"`
					TalkToGuide struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guide"`
					PublicIsland struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"public_island"`
					CraftWorkbench struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"craft_workbench"`
					ExploreHub struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"explore_hub"`
					ExploreVillage struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"explore_village"`
					TalkToLibrarian struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_librarian"`
					TalkToFarmer struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_farmer"`
					TalkToBlacksmith struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_blacksmith"`
					TalkToLumberjack struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_lumberjack"`
					TalkToEventMaster struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_event_master"`
					TalkToAuctionMaster struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_auction_master"`
					TalkToBanker struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_banker"`
					TalkToFairy struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_fairy"`
					TalkToFisherman1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_fisherman_1"`
					TalkToCarpenter struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_carpenter"`
					TalkToArtist1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_artist_1"`
					PaintCanvas struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"paint_canvas"`
					TalkToLazyMiner struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_lazy_miner"`
					IncreaseMiningSkill5 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"increase_mining_skill_5"`
					TalkToTelekinesisApplier struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_telekinesis_applier"`
					FindPickaxe struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"find_pickaxe"`
					CollectIngots struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						IronIngot   bool    `json:"IRON_INGOT"`
						GoldIngot   bool    `json:"GOLD_INGOT"`
					} `json:"collect_ingots"`
					WarpDeepCaverns struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"warp_deep_caverns"`
					TalkToLapisMiner struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_lapis_miner"`
					TalkToLiftOperator struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_lift_operator"`
					ReachLapisQuarry struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"reach_lapis_quarry"`
					CollectLapis struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						INKSACK4    bool    `json:"INK_SACK:4"`
					} `json:"collect_lapis"`
					DepositCoins struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"deposit_coins"`
					EnchantItem struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"enchant_item"`
					TalkToFarmhand1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_farmhand_1"`
					IncreaseFarmingSkill5 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"increase_farming_skill_5"`
					WarpMushroomDesert struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"warp_mushroom_desert"`
					IncreaseForagingSkill5 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"increase_foraging_skill_5"`
					TalkToGustave1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_gustave_1"`
					CollectDarkOakLogs struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"collect_dark_oak_logs"`
					TalkToCharlie2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_charlie_2"`
					CollectFarmingResources2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						INKSACK3    bool    `json:"INK_SACK:3"`
						Cactus      bool    `json:"CACTUS"`
						SugarCane   bool    `json:"SUGAR_CANE"`
					} `json:"collect_farming_resources_2"`
					IncreaseCombatSkill5 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"increase_combat_skill_5"`
					TalkToRick struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_rick"`
					WarpTheEnd struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"warp_the_end"`
					WarpBlazingFortress struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"warp_blazing_fortress"`
					CollectClay struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"collect_clay"`
					TalkToFisherman2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_fisherman_2"`
					GiveFairySouls struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"give_fairy_souls"`
					TalkToElle struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_elle"`
					TalkToGuber1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guber_1"`
					TalkToEndDealer struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_end_dealer"`
					CompleteTheEndRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_end_race_1"`
					TalkToGuber2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guber_2"`
					CompleteTheEndRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_end_race_2"`
					TalkToGuber3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guber_3"`
					CompleteTheEndRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_end_race_3"`
					TalkToGuber4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guber_4"`
					CompleteTheEndRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_end_race_4"`
					TalkToGuber5 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guber_5"`
					MineCoal struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"mine_coal"`
					TalkToBlacksmith2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_blacksmith_2"`
					IncreaseMiningSkill struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"increase_mining_skill"`
					ReforgeItem struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"reforge_item"`
					WarpGoldMine struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"warp_gold_mine"`
					TalkToMelody struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_melody"`
					FightDragon struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"fight_dragon"`
					KillDangerMobs struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"kill_danger_mobs"`
					TalkToBartender struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_bartender"`
					CollectWheat struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"collect_wheat"`
					TalkToFarmer2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_farmer_2"`
					IncreaseFarmingSkill struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"increase_farming_skill"`
					WarpBarnIsland struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"warp_barn_island"`
					IncreaseCombatSkill struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"increase_combat_skill"`
					WarpSpidersDen struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"warp_spiders_den"`
					TalkToHaymitch struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_haymitch"`
					TalkToFrosty struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_frosty"`
					TalkToGulliver1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_gulliver_1"`
					CompleteTheChickenRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_chicken_race_1"`
					TalkToGulliver2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_gulliver_2"`
					CompleteTheChickenRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_chicken_race_2"`
					TalkToGulliver3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_gulliver_3"`
					CompleteTheChickenRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_chicken_race_3"`
					TalkToGulliver4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_gulliver_4"`
					CompleteTheChickenRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_chicken_race_4"`
					ChopTree struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"chop_tree"`
					TalkToLumberjack2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_lumberjack_2"`
					CollectFarmAnimalResources2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						Rabbit      bool    `json:"RABBIT"`
						Mutton      bool    `json:"MUTTON"`
					} `json:"collect_farm_animal_resources_2"`
					CollectWoolCarpenter struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"collect_wool_carpenter"`
					TalkToPetCollector struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_pet_collector"`
					TalkToPetSitter struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_pet_sitter"`
					TalkToGuildford1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_1"`
					CompleteTheGiantMushroomAnythingWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_with_return_race_1"`
					CompleteTheGiantMushroomNoPearlsWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_with_return_race_1"`
					CompleteTheGiantMushroomNoAbilitiesWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_with_return_race_1"`
					CompleteTheGiantMushroomNothingWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_with_return_race_1"`
					CompleteThePrecursorRuinsAnythingWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_with_return_race_1"`
					CompleteThePrecursorRuinsNoPearlsWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_with_return_race_1"`
					CompleteThePrecursorRuinsNoAbilitiesWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_with_return_race_1"`
					CompleteThePrecursorRuinsNothingWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_with_return_race_1"`
					CompleteTheCrystalCoreAnythingWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_with_return_race_1"`
					CompleteTheCrystalCoreNoPearlsWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_with_return_race_1"`
					CompleteTheCrystalCoreNoAbilitiesWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_with_return_race_1"`
					CompleteTheCrystalCoreNothingWithReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_with_return_race_1"`
					CompleteTheGiantMushroomAnythingNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_1"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_1"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_1"`
					CompleteTheGiantMushroomNothingNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_1"`
					CompleteThePrecursorRuinsAnythingNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_1"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_1"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_1"`
					CompleteThePrecursorRuinsNothingNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_1"`
					CompleteTheCrystalCoreAnythingNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_1"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_1"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_1"`
					CompleteTheCrystalCoreNothingNoReturnRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_1"`
					TalkToGuildfordGiantMushroomAnythingNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_1"`
					CompleteTheGiantMushroomAnythingNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_2"`
					TalkToGuildfordGiantMushroomAnythingNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_2"`
					CompleteTheGiantMushroomAnythingNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_3"`
					TalkToGuildfordGiantMushroomAnythingNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_3"`
					CompleteTheGiantMushroomAnythingNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_4"`
					TalkToGuildfordGiantMushroomAnythingNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_4"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_1"`
					CompleteThePrecursorRuinsAnythingNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_2"`
					CompleteThePrecursorRuinsAnythingNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_3"`
					CompleteThePrecursorRuinsAnythingNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_4"`
					TalkToGuildfordCrystalCoreAnythingNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_1"`
					CompleteTheCrystalCoreAnythingNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_2"`
					TalkToGuildfordCrystalCoreAnythingNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_2"`
					CompleteTheCrystalCoreAnythingNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_3"`
					TalkToGuildfordCrystalCoreAnythingNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_3"`
					CompleteTheCrystalCoreAnythingNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_4"`
					TalkToGuildfordCrystalCoreAnythingNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_4"`
					CraftWoodPickaxe struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"craft_wood_pickaxe"`
					ReachPigmensDen struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"reach_pigmens_den"`
					CraftWheatMinion struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"craft_wheat_minion"`
					GivePickaxeLapisMiner struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"give_pickaxe_lapis_miner"`
					CollectRedstone struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						Redstone    bool    `json:"REDSTONE"`
					} `json:"collect_redstone"`
					ReachSlimehill struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"reach_slimehill"`
					CollectEmerald struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						Emerald     bool    `json:"EMERALD"`
					} `json:"collect_emerald"`
					ReachDiamondReserve struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"reach_diamond_reserve"`
					IncreaseForagingSkill struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"increase_foraging_skill"`
					WarpForagingIslands struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"warp_foraging_islands"`
					CollectBirchLogs struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"collect_birch_logs"`
					TalkToCharlie struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_charlie"`
					TalkToArtist2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_artist_2"`
					CollectDiamond struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						Diamond     bool    `json:"DIAMOND"`
					} `json:"collect_diamond"`
					ReachObsidianSanctuary struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"reach_obsidian_sanctuary"`
					CollectObsidian struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						Obsidian    bool    `json:"OBSIDIAN"`
					} `json:"collect_obsidian"`
					CollectSpider struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						String      bool    `json:"STRING"`
						SpiderEye   bool    `json:"SPIDER_EYE"`
					} `json:"collect_spider"`
					CollectNetherResources struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						BlazeRod    bool    `json:"BLAZE_ROD"`
						NetherStalk bool    `json:"NETHER_STALK"`
					} `json:"collect_nether_resources"`
					CollectNetherResources2 struct {
						Status        string  `json:"status"`
						Progress      float64 `json:"progress"`
						CompletedAt   float64 `json:"completed_at"`
						Quartz        bool    `json:"QUARTZ"`
						GlowstoneDust bool    `json:"GLOWSTONE_DUST"`
						MagmaCream    bool    `json:"MAGMA_CREAM"`
					} `json:"collect_nether_resources_2"`
					CollectEndStone struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						EnderStone  bool    `json:"ENDER_STONE"`
					} `json:"collect_end_stone"`
					ReachDragonsNest struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"reach_dragons_nest"`
					GiveRickIngots struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"give_rick_ingots"`
					TalkToFarmhand2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_farmhand_2"`
					CollectFarmingResources struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						PotatoItem  bool    `json:"POTATO_ITEM"`
						CarrotItem  bool    `json:"CARROT_ITEM"`
						Pumpkin     bool    `json:"PUMPKIN"`
						Melon       bool    `json:"MELON"`
					} `json:"collect_farming_resources"`
					CollectFarmAnimalResources struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
						RawChicken  bool    `json:"RAW_CHICKEN"`
						Leather     bool    `json:"LEATHER"`
						Pork        bool    `json:"PORK"`
					} `json:"collect_farm_animal_resources"`
					CompleteTheWoodsRace1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_woods_race_1"`
					TalkToGustave2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_gustave_2"`
					CompleteTheWoodsRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_woods_race_2"`
					TalkToGustave3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_gustave_3"`
					CompleteTheWoodsRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_woods_race_3"`
					TalkToGustave4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_gustave_4"`
					CompleteTheWoodsRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_woods_race_4"`
					TalkToGustave5 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_gustave_5"`
					TalkToGulliver5 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_gulliver_5"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_1"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_2"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_2"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_3"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_3"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_4"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_4"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_1"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_2"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_2"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_3"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_3"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_4"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_4"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_1"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsAnythingWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_with_return_1"`
					CompleteThePrecursorRuinsAnythingWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_with_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoPearlsWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_with_return_1"`
					CompleteThePrecursorRuinsNoPearlsWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_with_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_2"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_3"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_4"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_1"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_2"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_2"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_3"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_3"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_4"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_4"`
					TalkToGuildfordPrecursorRuinsAnythingWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_with_return_2"`
					CompleteThePrecursorRuinsAnythingWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_with_return_race_3"`
					TalkToGuildfordPrecursorRuinsAnythingWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_with_return_3"`
					CompleteThePrecursorRuinsAnythingWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_with_return_race_4"`
					TalkToGuildfordPrecursorRuinsAnythingWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_with_return_4"`
					TalkToGuildfordPrecursorRuinsNoPearlsWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_with_return_2"`
					CompleteThePrecursorRuinsNoPearlsWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_with_return_race_3"`
					TalkToGuildfordPrecursorRuinsNoPearlsWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_with_return_3"`
					CompleteThePrecursorRuinsNoPearlsWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_with_return_race_4"`
					TalkToGuildfordPrecursorRuinsNoPearlsWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_with_return_4"`
					TalkToGuildfordGiantMushroomAnythingWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_with_return_1"`
					CompleteTheGiantMushroomAnythingWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_with_return_race_2"`
					TalkToGuildfordGiantMushroomAnythingWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_with_return_2"`
					CompleteTheGiantMushroomAnythingWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_with_return_race_3"`
					TalkToGuildfordGiantMushroomAnythingWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_with_return_3"`
					CompleteTheGiantMushroomAnythingWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_with_return_race_4"`
					TalkToGuildfordGiantMushroomNoPearlsWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_with_return_1"`
					CompleteTheGiantMushroomNoPearlsWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_with_return_race_2"`
					TalkToGuildfordGiantMushroomNoPearlsWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_with_return_2"`
					CompleteTheGiantMushroomNoPearlsWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_with_return_race_3"`
					TalkToGuildfordGiantMushroomNoPearlsWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_with_return_3"`
					CompleteTheGiantMushroomNoPearlsWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_with_return_race_4"`
					TalkToGuildfordGiantMushroomNoPearlsWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_with_return_4"`
					TalkToGuildfordCrystalCoreAnythingWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_with_return_1"`
					CompleteTheCrystalCoreAnythingWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_with_return_race_2"`
					TalkToGuildfordCrystalCoreAnythingWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_with_return_2"`
					CompleteTheCrystalCoreAnythingWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_with_return_race_3"`
					TalkToGuildfordCrystalCoreNoPearlsWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_with_return_1"`
					CompleteTheCrystalCoreNoPearlsWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_with_return_race_2"`
					TalkToGuildfordCrystalCoreNoPearlsWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_with_return_2"`
					CompleteTheCrystalCoreNoPearlsWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_with_return_race_3"`
					TalkToGuildfordCrystalCoreNoPearlsWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_with_return_3"`
					CompleteTheCrystalCoreNoPearlsWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_with_return_race_4"`
					TalkToCryptKeeper1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_crypt_keeper_1"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_1"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_2"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_2"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_3"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_3"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_4"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_4"`
					TalkToGuildfordGiantMushroomNothingNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_1"`
					CompleteTheGiantMushroomNothingNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_2"`
					TalkToGuildfordGiantMushroomNothingNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_2"`
					CompleteTheGiantMushroomNothingNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_3"`
					TalkToGuildfordGiantMushroomNothingNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_3"`
					CompleteTheGiantMushroomNothingNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_4"`
					TalkToGuildfordGiantMushroomNothingNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_4"`
					AuctionItem struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"auction_item"`
					TalkToNocturne1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_nocturne_1"`
					CollectRunicFragments struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"collect_runic_fragments"`
					TalkToNocturne2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_nocturne_2"`
					TalkToDusk1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_dusk_1"`
					CollectRunes struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"collect_runes"`
					TalkToGuildfordCrystalCoreNothingNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_1"`
					CompleteTheCrystalCoreNothingNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_2"`
					TalkToGuildfordCrystalCoreNothingNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_2"`
					CompleteTheCrystalCoreNothingNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_3"`
					TalkToGuildfordCrystalCoreAnythingWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_with_return_3"`
					CompleteTheCrystalCoreAnythingWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_with_return_race_4"`
					TalkToGuildfordCrystalCoreAnythingWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_with_return_4"`
					TalkToGuildfordCrystalCoreNoAbilitiesWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_with_return_1"`
					CompleteTheCrystalCoreNoAbilitiesWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_with_return_race_2"`
					TalkToGuildfordCrystalCoreNoAbilitiesWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_with_return_2"`
					CompleteTheCrystalCoreNoAbilitiesWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_with_return_race_3"`
					TalkToGuildfordCrystalCoreNoAbilitiesWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_with_return_3"`
					CompleteTheCrystalCoreNoAbilitiesWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_with_return_race_4"`
					TalkToGuildfordCrystalCoreNothingWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_with_return_1"`
					CompleteTheCrystalCoreNothingWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_with_return_race_2"`
					TalkToGuildfordCrystalCoreNothingWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_with_return_2"`
					CompleteTheCrystalCoreNothingWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_with_return_race_3"`
					TalkToGuildfordCrystalCoreNothingWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_with_return_3"`
					CompleteTheCrystalCoreNothingWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_with_return_race_4"`
					TalkToGuildfordGiantMushroomNoAbilitiesWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_with_return_1"`
					CompleteTheGiantMushroomNoAbilitiesWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_with_return_race_2"`
					TalkToGuildfordGiantMushroomNoAbilitiesWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_with_return_2"`
					CompleteTheGiantMushroomNoAbilitiesWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_with_return_race_3"`
					TalkToGuildfordGiantMushroomNoAbilitiesWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_with_return_3"`
					CompleteTheGiantMushroomNoAbilitiesWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_with_return_race_4"`
					TalkToGuildfordGiantMushroomNothingWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_with_return_1"`
					CompleteTheGiantMushroomNothingWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_with_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_1"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_2"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_3"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_1"`
					CompleteThePrecursorRuinsNothingNoReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_2"`
					CompleteThePrecursorRuinsNothingNoReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_3"`
					CompleteThePrecursorRuinsNothingNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_4"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_with_return_1"`
					CompleteThePrecursorRuinsNoAbilitiesWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_with_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_with_return_2"`
					CompleteThePrecursorRuinsNoAbilitiesWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_with_return_race_3"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_with_return_3"`
					CompleteThePrecursorRuinsNoAbilitiesWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_with_return_race_4"`
					TalkToGuildfordPrecursorRuinsNothingWithReturn1 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_with_return_1"`
					CompleteThePrecursorRuinsNothingWithReturnRace2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_with_return_race_2"`
					TalkToGuildfordPrecursorRuinsNothingWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_with_return_2"`
					CompleteThePrecursorRuinsNothingWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_with_return_race_3"`
					TalkToGuildfordPrecursorRuinsNothingWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_with_return_3"`
					CompleteThePrecursorRuinsNothingWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_with_return_race_4"`
					TalkToGuildfordGiantMushroomNothingWithReturn2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_with_return_2"`
					CompleteTheGiantMushroomNothingWithReturnRace3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_with_return_race_3"`
					TalkToGuildfordCrystalCoreNoAbilitiesWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_with_return_4"`
					TalkToGuildfordCrystalCoreNothingWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_with_return_4"`
					TalkToGuildfordCrystalCoreNothingNoReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_3"`
					CompleteTheCrystalCoreNothingNoReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_4"`
					TalkToGuildfordCrystalCoreNothingNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_4"`
					TalkToGuildfordGiantMushroomNothingWithReturn3 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_with_return_3"`
					CompleteTheGiantMushroomNothingWithReturnRace4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_with_return_race_4"`
					TalkToGuildfordGiantMushroomNothingWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_with_return_4"`
					TalkToGuildfordPrecursorRuinsNothingWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_with_return_4"`
					TalkToGuildfordGiantMushroomAnythingWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_with_return_4"`
					TalkToGuildfordGiantMushroomNoAbilitiesWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_with_return_4"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_4"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_with_return_4"`
					TalkToGuildfordCrystalCoreNoPearlsWithReturn4 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_with_return_4"`
					TalkToRhys struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_rhys"`
					IncreaseMining12 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"increase_mining_12"`
					HotmGiveMaterials struct {
						Status               string  `json:"status"`
						Progress             float64 `json:"progress"`
						CompletedAt          float64 `json:"completed_at"`
						Started              bool    `json:"started"`
						EnchantedRedstone    float64 `json:"ENCHANTED_REDSTONE"`
						EnchantedIron        float64 `json:"ENCHANTED_IRON"`
						EnchantedCoal        float64 `json:"ENCHANTED_COAL"`
						EnchantedDiamond     float64 `json:"ENCHANTED_DIAMOND"`
						EnchantedLapisLazuli float64 `json:"ENCHANTED_LAPIS_LAZULI"`
					} `json:"hotm_give_materials"`
					Fetchur150 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-15-0"`
					TalkToArchaeologist struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_archaeologist"`
					TalkToShaggy struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_shaggy"`
					FindRelics struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"find_relics"`
					FindUberRelics struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"find_uber_relics"`
					TalkToShaggy2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_shaggy_2"`
					TalkToArchaeologist2 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_archaeologist_2"`
					TalkToGwendolyn struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_gwendolyn"`
					TalkToBraum struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_braum"`
					VisitGreaterMines struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"visit_greater_mines"`
					TalkToTheGoblinKing struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_the_goblin_king"`
					KillAutomatons struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"kill_automatons"`
					EnterDivanMines struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"enter_divan_mines"`
					FindAJungleKey struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"find_a_jungle_key"`
					MineRuby struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"mine_ruby"`
					CompleteTrialsOfJungleTemple struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"complete_trials_of_jungle_temple"`
					MineAmethyst struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"mine_amethyst"`
					TalkToProfessorRobot struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"talk_to_professor_robot"`
					FindTheGoblinQueen struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"find_the_goblin_queen"`
					MineAmber struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"mine_amber"`
					MineSapphire struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"mine_sapphire"`
					FindFourMissingPieces struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"find_four_missing_pieces"`
					MineJade struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt float64 `json:"completed_at"`
					} `json:"mine_jade"`
					Fetchur160 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-16-0"`
					Fetchur170 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-17-0"`
					Fetchur180 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-18-0"`
					Fetchur190 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-19-0"`
					Fetchur240 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-24-0"`
					Fetchur250 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-25-0"`
					Fetchur260 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-26-0"`
					Fetchur270 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-27-0"`
					Fetchur280 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-28-0"`
					Fetchur290 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-29-0"`
					Fetchur81 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-8-1"`
					Fetchur94 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-9-4"`
					Fetchur104 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-10-4"`
					Fetchur114 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-11-4"`
					Fetchur305 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-30-5"`
					Fetchur16 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-1-6"`
					Fetchur126 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-12-6"`
					Fetchur136 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-13-6"`
					Fetchur272 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-27-2"`
					Fetchur220 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-22-0"`
					Fetchur61 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-6-1"`
					Fetchur71 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-7-1"`
					Fetchur193 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-19-3"`
					Fetchur203 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-20-3"`
					Fetchur115 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-11-5"`
					Fetchur56 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-5-6"`
					Fetchur230 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-23-0"`
					Fetchur217 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-21-7"`
					Fetchur228 struct {
						Status      string  `json:"status"`
						Progress    float64 `json:"progress"`
						CompletedAt int64   `json:"completed_at"`
					} `json:"fetchur-22-8"`
				} `json:"objectives"`
				Tutorial []string `json:"tutorial"`
				Quests   struct {
					CollectLog struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"collect_log"`
					ExploreHub struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"explore_hub"`
					ExploreVillage struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"explore_village"`
					TalkToLibrarian struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_librarian"`
					TalkToFarmer struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_farmer"`
					TalkToBlacksmith struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_blacksmith"`
					TalkToLumberjack struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_lumberjack"`
					TalkToAuctionMaster struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_auction_master"`
					TalkToBanker struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_banker"`
					TalkToCarpenter struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_carpenter"`
					TalkToArtist1 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_artist_1"`
					TalkToLazyMiner struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_lazy_miner"`
					IncreaseMiningSkill5 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"increase_mining_skill_5"`
					TalkToLapisMiner struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_lapis_miner"`
					TalkToFarmhand1 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_farmhand_1"`
					IncreaseFarmingSkill5 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"increase_farming_skill_5"`
					IncreaseForagingSkill5 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"increase_foraging_skill_5"`
					TalkToGustave1 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_gustave_1"`
					IncreaseCombatSkill5 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"increase_combat_skill_5"`
					TalkToRick struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_rick"`
					TalkToGuber1 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_guber_1"`
					TalkToEndDealer struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_end_dealer"`
					ReforgeItem struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"reforge_item"`
					KillDangerMobs struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"kill_danger_mobs"`
					TalkToGulliver1 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_gulliver_1"`
					TalkToGuildford1 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_guildford_1"`
					TalkToCryptKeeper1 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_crypt_keeper_1"`
					TalkToNocturne1 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_nocturne_1"`
					TalkToDusk1 struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_dusk_1"`
					TalkToRhys struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_rhys"`
					TalkToArchaeologist struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_archaeologist"`
					TalkToGwendolyn struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_gwendolyn"`
					TalkToTheGoblinKing struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"talk_to_the_goblin_king"`
					KillAutomatons struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"kill_automatons"`
					EnterDivanMines struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"enter_divan_mines"`
					FindAJungleKey struct {
						Status        string  `json:"status"`
						ActivatedAt   float64 `json:"activated_at"`
						ActivatedAtSb float64 `json:"activated_at_sb"`
						CompletedAt   float64 `json:"completed_at"`
						CompletedAtSb float64 `json:"completed_at_sb"`
					} `json:"find_a_jungle_key"`
				} `json:"quests"`
				CoinPurse                     float64       `json:"coin_purse"`
				LastDeath                     float64       `json:"last_death"`
				CraftedGenerators             []interface{} `json:"crafted_generators"`
				VisitedZones                  []string      `json:"visited_zones"`
				FairySoulsCollected           float64       `json:"fairy_souls_collected"`
				FairySouls                    float64       `json:"fairy_souls"`
				FairyExchanges                float64       `json:"fairy_exchanges"`
				FishingTreasureCaught         float64       `json:"fishing_treasure_caught"`
				DeathCount                    float64       `json:"death_count"`
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
						BossKillsTier0 float64 `json:"boss_kills_tier_0"`
						Xp             float64 `json:"xp"`
						BossKillsTier1 float64 `json:"boss_kills_tier_1"`
						BossKillsTier2 float64 `json:"boss_kills_tier_2"`
						BossKillsTier3 float64 `json:"boss_kills_tier_3"`
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
						BossKillsTier0 float64 `json:"boss_kills_tier_0"`
						Xp             float64 `json:"xp"`
						BossKillsTier1 float64 `json:"boss_kills_tier_1"`
						BossKillsTier2 float64 `json:"boss_kills_tier_2"`
						BossKillsTier3 float64 `json:"boss_kills_tier_3"`
						BossKillsTier4 float64 `json:"boss_kills_tier_4"`
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
						BossKillsTier0 float64 `json:"boss_kills_tier_0"`
						Xp             float64 `json:"xp"`
						BossKillsTier1 float64 `json:"boss_kills_tier_1"`
						BossKillsTier2 float64 `json:"boss_kills_tier_2"`
						BossKillsTier3 float64 `json:"boss_kills_tier_3"`
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
						BossKillsTier0 float64 `json:"boss_kills_tier_0"`
						Xp             float64 `json:"xp"`
						BossKillsTier1 float64 `json:"boss_kills_tier_1"`
						BossKillsTier2 float64 `json:"boss_kills_tier_2"`
						BossKillsTier3 float64 `json:"boss_kills_tier_3"`
					} `json:"enderman"`
				} `json:"slayer_bosses"`
				Pets []struct {
					Type      string  `json:"type"`
					Exp       float64 `json:"exp"`
					Active    bool    `json:"active"`
					Tier      string  `json:"tier"`
					HeldItem  string  `json:"heldItem"`
					CandyUsed float64 `json:"candyUsed"`
				} `json:"pets"`
				CoopInvitation struct {
					Timestamp          float64 `json:"timestamp"`
					InvitedBy          string  `json:"invited_by"`
					Confirmed          bool    `json:"confirmed"`
					ConfirmedTimestamp float64 `json:"confirmed_timestamp"`
				} `json:"coop_invitation"`
				ExperienceSkillCombat float64  `json:"experience_skill_combat"`
				ExperienceSkillMining float64  `json:"experience_skill_mining"`
				UnlockedCollTiers     []string `json:"unlocked_coll_tiers"`
				FishingBag            struct {
					Type float64 `json:"type"`
					Data string  `json:"data"`
				} `json:"fishing_bag"`
				ExperienceSkillAlchemy float64 `json:"experience_skill_alchemy"`
				ExperienceSkillFarming float64 `json:"experience_skill_farming"`
				Collection             struct {
					Log                         float64 `json:"LOG"`
					Seeds                       float64 `json:"SEEDS"`
					Cobblestone                 float64 `json:"COBBLESTONE"`
					Coal                        float64 `json:"COAL"`
					IronIngot                   float64 `json:"IRON_INGOT"`
					GoldIngot                   float64 `json:"GOLD_INGOT"`
					INKSACK4                    float64 `json:"INK_SACK:4"`
					Bone                        float64 `json:"BONE"`
					String                      float64 `json:"STRING"`
					SpiderEye                   float64 `json:"SPIDER_EYE"`
					CarrotItem                  float64 `json:"CARROT_ITEM"`
					LOG1                        float64 `json:"LOG:1"`
					Melon                       float64 `json:"MELON"`
					SugarCane                   float64 `json:"SUGAR_CANE"`
					Pumpkin                     float64 `json:"PUMPKIN"`
					Pork                        float64 `json:"PORK"`
					Wheat                       float64 `json:"WHEAT"`
					PotatoItem                  float64 `json:"POTATO_ITEM"`
					RawChicken                  float64 `json:"RAW_CHICKEN"`
					Feather                     float64 `json:"FEATHER"`
					Leather                     float64 `json:"LEATHER"`
					RottenFlesh                 float64 `json:"ROTTEN_FLESH"`
					MushroomCollection          float64 `json:"MUSHROOM_COLLECTION"`
					INKSACK3                    float64 `json:"INK_SACK:3"`
					Cactus                      float64 `json:"CACTUS"`
					Rabbit                      float64 `json:"RABBIT"`
					Mutton                      float64 `json:"MUTTON"`
					Diamond                     float64 `json:"DIAMOND"`
					Sand                        float64 `json:"SAND"`
					Gravel                      float64 `json:"GRAVEL"`
					EnderPearl                  float64 `json:"ENDER_PEARL"`
					EnderStone                  float64 `json:"ENDER_STONE"`
					LOG2                        float64 `json:"LOG:2"`
					LOG21                       float64 `json:"LOG_2:1"`
					LOG3                        float64 `json:"LOG:3"`
					Redstone                    float64 `json:"REDSTONE"`
					Emerald                     float64 `json:"EMERALD"`
					SlimeBall                   float64 `json:"SLIME_BALL"`
					Obsidian                    float64 `json:"OBSIDIAN"`
					NetherStalk                 float64 `json:"NETHER_STALK"`
					Sulphur                     float64 `json:"SULPHUR"`
					BlazeRod                    float64 `json:"BLAZE_ROD"`
					Netherrack                  float64 `json:"NETHERRACK"`
					Quartz                      float64 `json:"QUARTZ"`
					GlowstoneDust               float64 `json:"GLOWSTONE_DUST"`
					MagmaCream                  float64 `json:"MAGMA_CREAM"`
					RAWFISH3                    float64 `json:"RAW_FISH:3"`
					RawFish                     float64 `json:"RAW_FISH"`
					RAWFISH1                    float64 `json:"RAW_FISH:1"`
					ClayBall                    float64 `json:"CLAY_BALL"`
					WaterLily                   float64 `json:"WATER_LILY"`
					PrismarineShard             float64 `json:"PRISMARINE_SHARD"`
					PrismarineCrystals          float64 `json:"PRISMARINE_CRYSTALS"`
					InkSack                     float64 `json:"INK_SACK"`
					Log2                        float64 `json:"LOG_2"`
					RAWFISH2                    float64 `json:"RAW_FISH:2"`
					Sponge                      float64 `json:"SPONGE"`
					GhastTear                   float64 `json:"GHAST_TEAR"`
					Ice                         float64 `json:"ICE"`
					EnchantedLapisLazuli        float64 `json:"ENCHANTED_LAPIS_LAZULI"`
					EnchantedDiamond            float64 `json:"ENCHANTED_DIAMOND"`
					SnowBall                    float64 `json:"SNOW_BALL"`
					DOUBLEPLANT5                float64 `json:"DOUBLE_PLANT:5"`
					EnchantedSnowBlock          float64 `json:"ENCHANTED_SNOW_BLOCK"`
					EnchantedOakLog             float64 `json:"ENCHANTED_OAK_LOG"`
					RedRose                     float64 `json:"RED_ROSE"`
					DOUBLEPLANT1                float64 `json:"DOUBLE_PLANT:1"`
					DoublePlant                 float64 `json:"DOUBLE_PLANT"`
					Wool                        float64 `json:"WOOL"`
					EnchantedRawFish            float64 `json:"ENCHANTED_RAW_FISH"`
					EnchantedRawSalmon          float64 `json:"ENCHANTED_RAW_SALMON"`
					EnchantedSponge             float64 `json:"ENCHANTED_SPONGE"`
					EnchantedPufferfish         float64 `json:"ENCHANTED_PUFFERFISH"`
					EnchantedPrismarineShard    float64 `json:"ENCHANTED_PRISMARINE_SHARD"`
					EnchantedPrismarineCrystals float64 `json:"ENCHANTED_PRISMARINE_CRYSTALS"`
					MithrilOre                  float64 `json:"MITHRIL_ORE"`
					EnchantedClownfish          float64 `json:"ENCHANTED_CLOWNFISH"`
					EnchantedCookedSalmon       float64 `json:"ENCHANTED_COOKED_SALMON"`
					EnchantedMelonBlock         float64 `json:"ENCHANTED_MELON_BLOCK"`
					EnchantedMelon              float64 `json:"ENCHANTED_MELON"`
					RawSoulflow                 float64 `json:"RAW_SOULFLOW"`
					HardStone                   float64 `json:"HARD_STONE"`
					GemstoneCollection          float64 `json:"GEMSTONE_COLLECTION"`
					EnchantedBrownMushroom      float64 `json:"ENCHANTED_BROWN_MUSHROOM"`
					EnchantedRedMushroom        float64 `json:"ENCHANTED_RED_MUSHROOM"`
					Soulflow                    float64 `json:"SOULFLOW"`
				} `json:"collection"`
				Quiver struct {
					Type float64 `json:"type"`
					Data string  `json:"data"`
				} `json:"quiver"`
				EnderChestContents struct {
					Type float64 `json:"type"`
					Data string  `json:"data"`
				} `json:"ender_chest_contents"`
				PotionBag struct {
					Type float64 `json:"type"`
					Data string  `json:"data"`
				} `json:"potion_bag"`
				ExperienceSkillEnchanting float64 `json:"experience_skill_enchanting"`
				InvContents               struct {
					Type float64 `json:"type"`
					Data string  `json:"data"`
				} `json:"inv_contents"`
				TalismanBag struct {
					Type float64 `json:"type"`
					Data string  `json:"data"`
				} `json:"talisman_bag"`
				ExperienceSkillForaging     float64 `json:"experience_skill_foraging"`
				ExperienceSkillCarpentry    float64 `json:"experience_skill_carpentry"`
				ExperienceSkillRunecrafting float64 `json:"experience_skill_runecrafting"`
				WardrobeEquippedSlot        float64 `json:"wardrobe_equipped_slot"`
				WardrobeContents            struct {
					Type float64 `json:"type"`
					Data string  `json:"data"`
				} `json:"wardrobe_contents"`
				ExperienceSkillTaming float64 `json:"experience_skill_taming"`
				SacksCounts           struct {
					RawFish                float64 `json:"RAW_FISH"`
					RAWFISH1               float64 `json:"RAW_FISH:1"`
					Diamond                float64 `json:"DIAMOND"`
					Redstone               float64 `json:"REDSTONE"`
					INKSACK4               float64 `json:"INK_SACK:4"`
					Emerald                float64 `json:"EMERALD"`
					RAWFISH3               float64 `json:"RAW_FISH:3"`
					WaterLily              float64 `json:"WATER_LILY"`
					RAWFISH2               float64 `json:"RAW_FISH:2"`
					InkSack                float64 `json:"INK_SACK"`
					ClayBall               float64 `json:"CLAY_BALL"`
					PrismarineCrystals     float64 `json:"PRISMARINE_CRYSTALS"`
					PrismarineShard        float64 `json:"PRISMARINE_SHARD"`
					Sponge                 float64 `json:"SPONGE"`
					Bone                   float64 `json:"BONE"`
					RottenFlesh            float64 `json:"ROTTEN_FLESH"`
					Coal                   float64 `json:"COAL"`
					BlazeRod               float64 `json:"BLAZE_ROD"`
					Netherrack             float64 `json:"NETHERRACK"`
					MagmaCream             float64 `json:"MAGMA_CREAM"`
					EnderPearl             float64 `json:"ENDER_PEARL"`
					Obsidian               float64 `json:"OBSIDIAN"`
					Sand                   float64 `json:"SAND"`
					Ice                    float64 `json:"ICE"`
					String                 float64 `json:"STRING"`
					SpiderEye              float64 `json:"SPIDER_EYE"`
					RevenantFlesh          float64 `json:"REVENANT_FLESH"`
					Cobblestone            float64 `json:"COBBLESTONE"`
					IronIngot              float64 `json:"IRON_INGOT"`
					GoldIngot              float64 `json:"GOLD_INGOT"`
					Quartz                 float64 `json:"QUARTZ"`
					GlowstoneDust          float64 `json:"GLOWSTONE_DUST"`
					Gravel                 float64 `json:"GRAVEL"`
					EnderStone             float64 `json:"ENDER_STONE"`
					TarantulaWeb           float64 `json:"TARANTULA_WEB"`
					GhastTear              float64 `json:"GHAST_TEAR"`
					WolfTooth              float64 `json:"WOLF_TOOTH"`
					SlimeBall              float64 `json:"SLIME_BALL"`
					Sulphur                float64 `json:"SULPHUR"`
					SugarCane              float64 `json:"SUGAR_CANE"`
					Wheat                  float64 `json:"WHEAT"`
					CarrotItem             float64 `json:"CARROT_ITEM"`
					PotatoItem             float64 `json:"POTATO_ITEM"`
					Pumpkin                float64 `json:"PUMPKIN"`
					Melon                  float64 `json:"MELON"`
					Seeds                  float64 `json:"SEEDS"`
					MushroomCollection     float64 `json:"MUSHROOM_COLLECTION"`
					INKSACK3               float64 `json:"INK_SACK:3"`
					Cactus                 float64 `json:"CACTUS"`
					NetherStalk            float64 `json:"NETHER_STALK"`
					LOG3                   float64 `json:"LOG:3"`
					Feather                float64 `json:"FEATHER"`
					RawChicken             float64 `json:"RAW_CHICKEN"`
					LOG1                   float64 `json:"LOG:1"`
					LOG2                   float64 `json:"LOG:2"`
					Log                    float64 `json:"LOG"`
					LOG21                  float64 `json:"LOG_2:1"`
					Log2                   float64 `json:"LOG_2"`
					Pork                   float64 `json:"PORK"`
					Leather                float64 `json:"LEATHER"`
					Mutton                 float64 `json:"MUTTON"`
					Rabbit                 float64 `json:"RABBIT"`
					DungeonTrap            float64 `json:"DUNGEON_TRAP"`
					DungeonDecoy           float64 `json:"DUNGEON_DECOY"`
					SpiritLeap             float64 `json:"SPIRIT_LEAP"`
					InflatableJerry        float64 `json:"INFLATABLE_JERRY"`
					GreenCandy             float64 `json:"GREEN_CANDY"`
					PurpleCandy            float64 `json:"PURPLE_CANDY"`
					MithrilOre             float64 `json:"MITHRIL_ORE"`
					Starfall               float64 `json:"STARFALL"`
					TitaniumOre            float64 `json:"TITANIUM_ORE"`
					Treasurite             float64 `json:"TREASURITE"`
					RuneWhiteSpiral1       float64 `json:"RUNE_WHITE_SPIRAL_1"`
					RuneBlood21            float64 `json:"RUNE_BLOOD_2_1"`
					RuneHearts1            float64 `json:"RUNE_HEARTS_1"`
					RuneGem1               float64 `json:"RUNE_GEM_1"`
					RuneGolden1            float64 `json:"RUNE_GOLDEN_1"`
					RuneTidal1             float64 `json:"RUNE_TIDAL_1"`
					RuneClouds1            float64 `json:"RUNE_CLOUDS_1"`
					RuneSparkling1         float64 `json:"RUNE_SPARKLING_1"`
					RuneSmokey1            float64 `json:"RUNE_SMOKEY_1"`
					RuneZap1               float64 `json:"RUNE_ZAP_1"`
					RuneHot1               float64 `json:"RUNE_HOT_1"`
					RuneSpirit1            float64 `json:"RUNE_SPIRIT_1"`
					RuneIce1               float64 `json:"RUNE_ICE_1"`
					RuneFireSpiral1        float64 `json:"RUNE_FIRE_SPIRAL_1"`
					RuneMagic1             float64 `json:"RUNE_MAGIC_1"`
					RuneWake1              float64 `json:"RUNE_WAKE_1"`
					RuneLightning1         float64 `json:"RUNE_LIGHTNING_1"`
					RuneRainbow1           float64 `json:"RUNE_RAINBOW_1"`
					RuneLava1              float64 `json:"RUNE_LAVA_1"`
					RuneSnow1              float64 `json:"RUNE_SNOW_1"`
					RuneGem3               float64 `json:"RUNE_GEM_3"`
					RuneLava3              float64 `json:"RUNE_LAVA_3"`
					RuneSnow3              float64 `json:"RUNE_SNOW_3"`
					RuneWhiteSpiral3       float64 `json:"RUNE_WHITE_SPIRAL_3"`
					RuneHot3               float64 `json:"RUNE_HOT_3"`
					RuneBite3              float64 `json:"RUNE_BITE_3"`
					RuneIce3               float64 `json:"RUNE_ICE_3"`
					RuneMagic3             float64 `json:"RUNE_MAGIC_3"`
					RuneHearts3            float64 `json:"RUNE_HEARTS_3"`
					RuneZap3               float64 `json:"RUNE_ZAP_3"`
					RuneRedstone3          float64 `json:"RUNE_REDSTONE_3"`
					RuneBlood23            float64 `json:"RUNE_BLOOD_2_3"`
					RuneSparkling3         float64 `json:"RUNE_SPARKLING_3"`
					RuneSpirit3            float64 `json:"RUNE_SPIRIT_3"`
					RuneGolden3            float64 `json:"RUNE_GOLDEN_3"`
					RuneWake3              float64 `json:"RUNE_WAKE_3"`
					RuneClouds3            float64 `json:"RUNE_CLOUDS_3"`
					RuneZombieSlayer3      float64 `json:"RUNE_ZOMBIE_SLAYER_3"`
					RuneRainbow3           float64 `json:"RUNE_RAINBOW_3"`
					RuneWake2              float64 `json:"RUNE_WAKE_2"`
					RuneZap2               float64 `json:"RUNE_ZAP_2"`
					RuneSnake2             float64 `json:"RUNE_SNAKE_2"`
					RuneClouds2            float64 `json:"RUNE_CLOUDS_2"`
					RuneZombieSlayer2      float64 `json:"RUNE_ZOMBIE_SLAYER_2"`
					RuneHot2               float64 `json:"RUNE_HOT_2"`
					RuneBlood22            float64 `json:"RUNE_BLOOD_2_2"`
					RuneGem2               float64 `json:"RUNE_GEM_2"`
					RuneHearts2            float64 `json:"RUNE_HEARTS_2"`
					RuneSparkling2         float64 `json:"RUNE_SPARKLING_2"`
					RuneIce2               float64 `json:"RUNE_ICE_2"`
					RuneGolden2            float64 `json:"RUNE_GOLDEN_2"`
					RuneLightning2         float64 `json:"RUNE_LIGHTNING_2"`
					RuneRainbow2           float64 `json:"RUNE_RAINBOW_2"`
					RuneFireSpiral2        float64 `json:"RUNE_FIRE_SPIRAL_2"`
					RuneSnow2              float64 `json:"RUNE_SNOW_2"`
					RuneWhiteSpiral2       float64 `json:"RUNE_WHITE_SPIRAL_2"`
					RuneMusic2             float64 `json:"RUNE_MUSIC_2"`
					RuneCouture3           float64 `json:"RUNE_COUTURE_3"`
					RuneSnake3             float64 `json:"RUNE_SNAKE_3"`
					RuneCouture1           float64 `json:"RUNE_COUTURE_1"`
					RuneZombieSlayer1      float64 `json:"RUNE_ZOMBIE_SLAYER_1"`
					RuneBite1              float64 `json:"RUNE_BITE_1"`
					EnchantedSpiderEye     float64 `json:"ENCHANTED_SPIDER_EYE"`
					EnchantedString        float64 `json:"ENCHANTED_STRING"`
					RuneSnake1             float64 `json:"RUNE_SNAKE_1"`
					EnchantedBone          float64 `json:"ENCHANTED_BONE"`
					RuneRedstone1          float64 `json:"RUNE_REDSTONE_1"`
					RuneMusic1             float64 `json:"RUNE_MUSIC_1"`
					EnchantedEnderPearl    float64 `json:"ENCHANTED_ENDER_PEARL"`
					EnchantedMelon         float64 `json:"ENCHANTED_MELON"`
					EnchantedRottenFlesh   float64 `json:"ENCHANTED_ROTTEN_FLESH"`
					EnchantedCarrot        float64 `json:"ENCHANTED_CARROT"`
					EnchantedPotato        float64 `json:"ENCHANTED_POTATO"`
					EnchantedOakLog        float64 `json:"ENCHANTED_OAK_LOG"`
					EnchantedBirchLog      float64 `json:"ENCHANTED_BIRCH_LOG"`
					RedMushroom            float64 `json:"RED_MUSHROOM"`
					BrownMushroom          float64 `json:"BROWN_MUSHROOM"`
					WerewolfSkin           float64 `json:"WEREWOLF_SKIN"`
					RawBeef                float64 `json:"RAW_BEEF"`
					RabbitFoot             float64 `json:"RABBIT_FOOT"`
					RabbitHide             float64 `json:"RABBIT_HIDE"`
					Ectoplasm              float64 `json:"ECTOPLASM"`
					PumpkinGuts            float64 `json:"PUMPKIN_GUTS"`
					SnowBall               float64 `json:"SNOW_BALL"`
					RuneMagic2             float64 `json:"RUNE_MAGIC_2"`
					SnowBlock              float64 `json:"SNOW_BLOCK"`
					WhiteGift              float64 `json:"WHITE_GIFT"`
					IceHunk                float64 `json:"ICE_HUNK"`
					BlueIceHunk            float64 `json:"BLUE_ICE_HUNK"`
					RedGift                float64 `json:"RED_GIFT"`
					GreenGift              float64 `json:"GREEN_GIFT"`
					PackedIce              float64 `json:"PACKED_ICE"`
					EnchantedGunpowder     float64 `json:"ENCHANTED_GUNPOWDER"`
					RuneRedstone2          float64 `json:"RUNE_REDSTONE_2"`
					RuneSpirit2            float64 `json:"RUNE_SPIRIT_2"`
					RuneLava2              float64 `json:"RUNE_LAVA_2"`
					RuneBite2              float64 `json:"RUNE_BITE_2"`
					RuneLightning3         float64 `json:"RUNE_LIGHTNING_3"`
					EnchantedSugar         float64 `json:"ENCHANTED_SUGAR"`
					EnchantedCocoa         float64 `json:"ENCHANTED_COCOA"`
					NullSphere             float64 `json:"NULL_SPHERE"`
					RuneEndersnake1        float64 `json:"RUNE_ENDERSNAKE_1"`
					RuneEndersnake2        float64 `json:"RUNE_ENDERSNAKE_2"`
					RuneDragon1            float64 `json:"RUNE_DRAGON_1"`
					EnchantedNetherStalk   float64 `json:"ENCHANTED_NETHER_STALK"`
					EnchantedRedMushroom   float64 `json:"ENCHANTED_RED_MUSHROOM"`
					EnchantedBrownMushroom float64 `json:"ENCHANTED_BROWN_MUSHROOM"`
					HardStone              float64 `json:"HARD_STONE"`
					RoughRubyGem           float64 `json:"ROUGH_RUBY_GEM"`
					RoughAmethystGem       float64 `json:"ROUGH_AMETHYST_GEM"`
					RoughSapphireGem       float64 `json:"ROUGH_SAPPHIRE_GEM"`
					RoughJadeGem           float64 `json:"ROUGH_JADE_GEM"`
					RoughTopazGem          float64 `json:"ROUGH_TOPAZ_GEM"`
					RoughAmberGem          float64 `json:"ROUGH_AMBER_GEM"`
					FlawedAmberGem         float64 `json:"FLAWED_AMBER_GEM"`
					FlawedSapphireGem      float64 `json:"FLAWED_SAPPHIRE_GEM"`
					RoughJasperGem         float64 `json:"ROUGH_JASPER_GEM"`
					FlawedAmethystGem      float64 `json:"FLAWED_AMETHYST_GEM"`
					FlawedTopazGem         float64 `json:"FLAWED_TOPAZ_GEM"`
					FlawedRubyGem          float64 `json:"FLAWED_RUBY_GEM"`
					FlawedJasperGem        float64 `json:"FLAWED_JASPER_GEM"`
					EnchantedHardStone     float64 `json:"ENCHANTED_HARD_STONE"`
					EnchantedCobblestone   float64 `json:"ENCHANTED_COBBLESTONE"`
					FlawedJadeGem          float64 `json:"FLAWED_JADE_GEM"`
					EnchantedIron          float64 `json:"ENCHANTED_IRON"`
					EnchantedCoal          float64 `json:"ENCHANTED_COAL"`
					FineTopazGem           float64 `json:"FINE_TOPAZ_GEM"`
					FlawlessTopazGem       float64 `json:"FLAWLESS_TOPAZ_GEM"`
					PerfectTopazGem        float64 `json:"PERFECT_TOPAZ_GEM"`
					EnchantedRedstone      float64 `json:"ENCHANTED_REDSTONE"`
					FineJadeGem            float64 `json:"FINE_JADE_GEM"`
					EnchantedMithril       float64 `json:"ENCHANTED_MITHRIL"`
					EnchantedTitanium      float64 `json:"ENCHANTED_TITANIUM"`
					FineRubyGem            float64 `json:"FINE_RUBY_GEM"`
					EnchantedLapisLazuli   float64 `json:"ENCHANTED_LAPIS_LAZULI"`
					FineAmethystGem        float64 `json:"FINE_AMETHYST_GEM"`
					FlawlessAmethystGem    float64 `json:"FLAWLESS_AMETHYST_GEM"`
					PerfectAmethystGem     float64 `json:"PERFECT_AMETHYST_GEM"`
					FlawlessRubyGem        float64 `json:"FLAWLESS_RUBY_GEM"`
					PerfectRubyGem         float64 `json:"PERFECT_RUBY_GEM"`
					FlawlessJadeGem        float64 `json:"FLAWLESS_JADE_GEM"`
					PerfectJadeGem         float64 `json:"PERFECT_JADE_GEM"`
					FineAmberGem           float64 `json:"FINE_AMBER_GEM"`
					FlawlessAmberGem       float64 `json:"FLAWLESS_AMBER_GEM"`
					PerfectAmberGem        float64 `json:"PERFECT_AMBER_GEM"`
					FineSapphireGem        float64 `json:"FINE_SAPPHIRE_GEM"`
					FineJasperGem          float64 `json:"FINE_JASPER_GEM"`
					FlawlessSapphireGem    float64 `json:"FLAWLESS_SAPPHIRE_GEM"`
					PerfectSapphireGem     float64 `json:"PERFECT_SAPPHIRE_GEM"`
					EnchantedObsidian      float64 `json:"ENCHANTED_OBSIDIAN"`
					EnchantedGold          float64 `json:"ENCHANTED_GOLD"`
					EnchantedQuartz        float64 `json:"ENCHANTED_QUARTZ"`
					EnchantedGlowstoneDust float64 `json:"ENCHANTED_GLOWSTONE_DUST"`
					FlawlessJasperGem      float64 `json:"FLAWLESS_JASPER_GEM"`
					PerfectJasperGem       float64 `json:"PERFECT_JASPER_GEM"`
					RuneEndersnake3        float64 `json:"RUNE_ENDERSNAKE_3"`
					RuneJerry3             float64 `json:"RUNE_JERRY_3"`
				} `json:"sacks_counts"`
				ExperienceSkillFishing float64 `json:"experience_skill_fishing"`
				CandyInventoryContents struct {
					Type float64 `json:"type"`
					Data string  `json:"data"`
				} `json:"candy_inventory_contents"`
				SlayerQuest struct {
					Type            string  `json:"type"`
					Tier            float64 `json:"tier"`
					StartTimestamp  float64 `json:"start_timestamp"`
					CompletionState float64 `json:"completion_state"`
					CombatXp        float64 `json:"combat_xp"`
					RecentMobKills  []struct {
						Xp        float64 `json:"xp"`
						Timestamp float64 `json:"timestamp"`
					} `json:"recent_mob_kills"`
					LastKilledMobIsland   string  `json:"last_killed_mob_island"`
					XpOnLastFollowerSpawn float64 `json:"xp_on_last_follower_spawn"`
					SpawnTimestamp        float64 `json:"spawn_timestamp"`
					KillTimestamp         float64 `json:"kill_timestamp"`
				} `json:"slayer_quest"`
				Dungeons struct {
					DungeonTypes struct {
						Catacombs struct {
							TimesPlayed struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"times_played"`
							Experience      float64 `json:"experience"`
							TierCompletions struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"tier_completions"`
							FastestTime struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"fastest_time"`
							BestRuns struct {
								Num0 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"0"`
								Num1 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"1"`
								Num2 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"2"`
								Num3 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"3"`
								Num4 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"4"`
								Num5 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"5"`
								Num6 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"6"`
								Num7 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"7"`
							} `json:"best_runs"`
							BestScore struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"best_score"`
							MobsKilled struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"mobs_killed"`
							MostMobsKilled struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"most_mobs_killed"`
							MostDamageBerserk struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"most_damage_berserk"`
							MostHealing struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"most_healing"`
							WatcherKills struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"watcher_kills"`
							MostDamageHealer struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"most_damage_healer"`
							HighestTierCompleted float64 `json:"highest_tier_completed"`
							MostDamageMage       struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"most_damage_mage"`
							FastestTimeS struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"fastest_time_s"`
							FastestTimeSPlus struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"fastest_time_s_plus"`
							MostDamageArcher struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num7 float64 `json:"7"`
							} `json:"most_damage_archer"`
							MostDamageTank struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"most_damage_tank"`
							MilestoneCompletions struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"milestone_completions"`
						} `json:"catacombs"`
						MasterCatacombs struct {
							TierCompletions struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
							} `json:"tier_completions"`
							MilestoneCompletions struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
							} `json:"milestone_completions"`
							HighestTierCompleted float64 `json:"highest_tier_completed"`
							FastestTime          struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
							} `json:"fastest_time"`
							FastestTimeS struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
							} `json:"fastest_time_s"`
							BestRuns struct {
								Num1 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"1"`
								Num2 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"2"`
								Num3 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"3"`
								Num4 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"4"`
								Num5 []struct {
									Timestamp        float64       `json:"timestamp"`
									ScoreExploration float64       `json:"score_exploration"`
									ScoreSpeed       float64       `json:"score_speed"`
									ScoreSkill       float64       `json:"score_skill"`
									ScoreBonus       float64       `json:"score_bonus"`
									DungeonClass     string        `json:"dungeon_class"`
									Teammates        []interface{} `json:"teammates"`
									ElapsedTime      float64       `json:"elapsed_time"`
									DamageDealt      float64       `json:"damage_dealt"`
									Deaths           float64       `json:"deaths"`
									MobsKilled       float64       `json:"mobs_killed"`
									SecretsFound     float64       `json:"secrets_found"`
									DamageMitigated  float64       `json:"damage_mitigated"`
									AllyHealing      float64       `json:"ally_healing"`
								} `json:"5"`
							} `json:"best_runs"`
							BestScore struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
							} `json:"best_score"`
							MobsKilled struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
							} `json:"mobs_killed"`
							MostMobsKilled struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
							} `json:"most_mobs_killed"`
							MostDamageTank struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
							} `json:"most_damage_tank"`
							MostHealing struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
							} `json:"most_healing"`
							MostDamageMage struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
							} `json:"most_damage_mage"`
							MostDamageBerserk struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
							} `json:"most_damage_berserk"`
							FastestTimeSPlus struct {
								Num1 float64 `json:"1"`
								Num3 float64 `json:"3"`
								Num5 float64 `json:"5"`
							} `json:"fastest_time_s_plus"`
							MostDamageArcher struct {
								Num1 float64 `json:"1"`
								Num3 float64 `json:"3"`
							} `json:"most_damage_archer"`
							MostDamageHealer struct {
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
							} `json:"most_damage_healer"`
						} `json:"master_catacombs"`
					} `json:"dungeon_types"`
					PlayerClasses struct {
						Healer struct {
							Experience float64 `json:"experience"`
						} `json:"healer"`
						Mage struct {
							Experience float64 `json:"experience"`
						} `json:"mage"`
						Berserk struct {
							Experience float64 `json:"experience"`
						} `json:"berserk"`
						Archer struct {
							Experience float64 `json:"experience"`
						} `json:"archer"`
						Tank struct {
							Experience float64 `json:"experience"`
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
					Type float64 `json:"type"`
					Data string  `json:"data"`
				} `json:"personal_vault_contents"`
				Griffin struct {
					Burrows []struct {
						Ts    float64 `json:"ts"`
						X     float64 `json:"x"`
						Y     float64 `json:"y"`
						Z     float64 `json:"z"`
						Type  float64 `json:"type"`
						Tier  float64 `json:"tier"`
						Chain float64 `json:"chain"`
					} `json:"burrows"`
				} `json:"griffin"`
				Jacob2 struct {
					MedalsInv struct {
						Silver float64 `json:"silver"`
						Bronze float64 `json:"bronze"`
						Gold   float64 `json:"gold"`
					} `json:"medals_inv"`
					Perks struct {
						DoubleDrops     float64 `json:"double_drops"`
						FarmingLevelCap float64 `json:"farming_level_cap"`
					} `json:"perks"`
					Talked       bool          `json:"talked"`
					UniqueGolds2 []interface{} `json:"unique_golds2"`
				} `json:"jacob2"`
				Experimentation struct {
					Simon struct {
						LastAttempt float64 `json:"last_attempt"`
						Attempts0   float64 `json:"attempts_0"`
						BonusClicks float64 `json:"bonus_clicks"`
						LastClaimed float64 `json:"last_claimed"`
						Claims0     float64 `json:"claims_0"`
						BestScore0  float64 `json:"best_score_0"`
						Attempts2   float64 `json:"attempts_2"`
						Claims2     float64 `json:"claims_2"`
						BestScore2  float64 `json:"best_score_2"`
						Attempts5   float64 `json:"attempts_5"`
						Claims5     float64 `json:"claims_5"`
						BestScore5  float64 `json:"best_score_5"`
						Attempts3   float64 `json:"attempts_3"`
						Attempts1   float64 `json:"attempts_1"`
						Claims1     float64 `json:"claims_1"`
						BestScore1  float64 `json:"best_score_1"`
						Claims3     float64 `json:"claims_3"`
						BestScore3  float64 `json:"best_score_3"`
					} `json:"simon"`
					Pairings struct {
						LastClaimed float64 `json:"last_claimed"`
						Claims2     float64 `json:"claims_2"`
						BestScore2  float64 `json:"best_score_2"`
						LastAttempt float64 `json:"last_attempt"`
						Claims3     float64 `json:"claims_3"`
						BestScore3  float64 `json:"best_score_3"`
						Claims4     float64 `json:"claims_4"`
						BestScore4  float64 `json:"best_score_4"`
						Claims0     float64 `json:"claims_0"`
						BestScore0  float64 `json:"best_score_0"`
						Claims1     float64 `json:"claims_1"`
						BestScore1  float64 `json:"best_score_1"`
						Claims5     float64 `json:"claims_5"`
						BestScore5  float64 `json:"best_score_5"`
					} `json:"pairings"`
					Numbers struct {
						LastAttempt float64 `json:"last_attempt"`
						Attempts3   float64 `json:"attempts_3"`
						Attempts2   float64 `json:"attempts_2"`
						BonusClicks float64 `json:"bonus_clicks"`
						LastClaimed float64 `json:"last_claimed"`
						Claims2     float64 `json:"claims_2"`
						BestScore2  float64 `json:"best_score_2"`
						Attempts1   float64 `json:"attempts_1"`
						Claims1     float64 `json:"claims_1"`
						BestScore1  float64 `json:"best_score_1"`
						Claims3     float64 `json:"claims_3"`
						BestScore3  float64 `json:"best_score_3"`
					} `json:"numbers"`
					ClaimsResets          float64 `json:"claims_resets"`
					ClaimsResetsTimestamp float64 `json:"claims_resets_timestamp"`
				} `json:"experimentation"`
				BackpackContents struct {
					Num0 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"0"`
					Num1 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"1"`
					Num2 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"2"`
					Num3 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"3"`
					Num4 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"4"`
					Num5 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"5"`
					Num6 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"6"`
					Num7 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"7"`
					Num8 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"8"`
					Num9 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"9"`
					Num10 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"10"`
					Num11 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"11"`
					Num12 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"12"`
					Num13 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"13"`
					Num14 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"14"`
					Num15 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"15"`
					Num16 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"16"`
					Num17 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"17"`
				} `json:"backpack_contents"`
				BackpackIcons struct {
					Num0 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"0"`
					Num1 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"1"`
					Num2 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"2"`
					Num3 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"3"`
					Num4 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"4"`
					Num5 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"5"`
					Num6 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"6"`
					Num7 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"7"`
					Num8 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"8"`
					Num9 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"9"`
					Num10 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"10"`
					Num11 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"11"`
					Num12 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"12"`
					Num13 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"13"`
					Num14 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"14"`
					Num15 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"15"`
					Num16 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"16"`
					Num17 struct {
						Type float64 `json:"type"`
						Data string  `json:"data"`
					} `json:"17"`
				} `json:"backpack_icons"`
				Perks struct {
					PermanentHealth       float64 `json:"permanent_health"`
					CatacombsBossLuck     float64 `json:"catacombs_boss_luck"`
					CatacombsDefense      float64 `json:"catacombs_defense"`
					CatacombsStrength     float64 `json:"catacombs_strength"`
					CatacombsHealth       float64 `json:"catacombs_health"`
					CatacombsCritDamage   float64 `json:"catacombs_crit_damage"`
					CatacombsIntelligence float64 `json:"catacombs_intelligence"`
					PermanentDefense      float64 `json:"permanent_defense"`
					PermanentSpeed        float64 `json:"permanent_speed"`
					PermanentIntelligence float64 `json:"permanent_intelligence"`
					PermanentStrength     float64 `json:"permanent_strength"`
					ForbiddenBlessing     float64 `json:"forbidden_blessing"`
					CatacombsLooting      float64 `json:"catacombs_looting"`
				} `json:"perks"`
				ActiveEffects []struct {
					Effect         string        `json:"effect"`
					Level          float64       `json:"level"`
					Modifiers      []interface{} `json:"modifiers"`
					TicksRemaining float64       `json:"ticks_remaining"`
					Infinite       bool          `json:"infinite"`
				} `json:"active_effects"`
				VisitedModes []interface{} `json:"visited_modes"`
				HarpQuest    struct {
					SelectedSong                          string  `json:"selected_song"`
					SelectedSongEpoch                     float64 `json:"selected_song_epoch"`
					SongHymnJoyBestCompletion             float64 `json:"song_hymn_joy_best_completion"`
					SongHymnJoyCompletions                float64 `json:"song_hymn_joy_completions"`
					SongHymnJoyPerfectCompletions         float64 `json:"song_hymn_joy_perfect_completions"`
					SongFrereJacquesBestCompletion        float64 `json:"song_frere_jacques_best_completion"`
					SongFrereJacquesCompletions           float64 `json:"song_frere_jacques_completions"`
					SongFrereJacquesPerfectCompletions    float64 `json:"song_frere_jacques_perfect_completions"`
					SongAmazingGraceBestCompletion        float64 `json:"song_amazing_grace_best_completion"`
					SongAmazingGraceCompletions           float64 `json:"song_amazing_grace_completions"`
					SongAmazingGracePerfectCompletions    float64 `json:"song_amazing_grace_perfect_completions"`
					SongBrahmsBestCompletion              float64 `json:"song_brahms_best_completion"`
					SongBrahmsCompletions                 float64 `json:"song_brahms_completions"`
					SongBrahmsPerfectCompletions          float64 `json:"song_brahms_perfect_completions"`
					SongHappyBirthdayBestCompletion       float64 `json:"song_happy_birthday_best_completion"`
					SongHappyBirthdayCompletions          float64 `json:"song_happy_birthday_completions"`
					SongHappyBirthdayPerfectCompletions   float64 `json:"song_happy_birthday_perfect_completions"`
					SongGreensleevesBestCompletion        float64 `json:"song_greensleeves_best_completion"`
					SongGreensleevesCompletions           float64 `json:"song_greensleeves_completions"`
					SongGreensleevesPerfectCompletions    float64 `json:"song_greensleeves_perfect_completions"`
					SongJeopardyBestCompletion            float64 `json:"song_jeopardy_best_completion"`
					SongJeopardyCompletions               float64 `json:"song_jeopardy_completions"`
					SongJeopardyPerfectCompletions        float64 `json:"song_jeopardy_perfect_completions"`
					SongMinuetBestCompletion              float64 `json:"song_minuet_best_completion"`
					SongMinuetCompletions                 float64 `json:"song_minuet_completions"`
					SongMinuetPerfectCompletions          float64 `json:"song_minuet_perfect_completions"`
					SongJoyWorldBestCompletion            float64 `json:"song_joy_world_best_completion"`
					SongJoyWorldCompletions               float64 `json:"song_joy_world_completions"`
					SongJoyWorldPerfectCompletions        float64 `json:"song_joy_world_perfect_completions"`
					SongPureImaginationBestCompletion     float64 `json:"song_pure_imagination_best_completion"`
					SongPureImaginationCompletions        float64 `json:"song_pure_imagination_completions"`
					SongPureImaginationPerfectCompletions float64 `json:"song_pure_imagination_perfect_completions"`
					SongVieEnRoseBestCompletion           float64 `json:"song_vie_en_rose_best_completion"`
					SongVieEnRoseCompletions              float64 `json:"song_vie_en_rose_completions"`
					SongVieEnRosePerfectCompletions       float64 `json:"song_vie_en_rose_perfect_completions"`
					ClaimedTalisman                       bool    `json:"claimed_talisman"`
					SongFireAndFlamesBestCompletion       float64 `json:"song_fire_and_flames_best_completion"`
					SongFireAndFlamesCompletions          float64 `json:"song_fire_and_flames_completions"`
					SongFireAndFlamesPerfectCompletions   float64 `json:"song_fire_and_flames_perfect_completions"`
					SongPachelbelBestCompletion           float64 `json:"song_pachelbel_best_completion"`
					SongPachelbelCompletions              float64 `json:"song_pachelbel_completions"`
				} `json:"harp_quest"`
				PausedEffects []struct {
					Effect         string        `json:"effect"`
					Level          float64       `json:"level"`
					Modifiers      []interface{} `json:"modifiers"`
					TicksRemaining float64       `json:"ticks_remaining"`
					Infinite       bool          `json:"infinite"`
				} `json:"paused_effects"`
				DisabledPotionEffects []interface{} `json:"disabled_potion_effects"`
				MiningCore            struct {
					Nodes struct {
						MiningSpeed          float64 `json:"mining_speed"`
						Special0             float64 `json:"special_0"`
						MiningFortune        float64 `json:"mining_fortune"`
						TitaniumInsanium     float64 `json:"titanium_insanium"`
						MiningSpeedBoost     float64 `json:"mining_speed_boost"`
						ForgeTime            float64 `json:"forge_time"`
						DailyPowder          float64 `json:"daily_powder"`
						EfficientMiner       float64 `json:"efficient_miner"`
						MiningExperience     float64 `json:"mining_experience"`
						MiningMadness        float64 `json:"mining_madness"`
						DailyEffect          float64 `json:"daily_effect"`
						ExperienceOrbs       float64 `json:"experience_orbs"`
						FrontLoaded          float64 `json:"front_loaded"`
						FallenStarBonus      float64 `json:"fallen_star_bonus"`
						RandomEvent          float64 `json:"random_event"`
						GoblinKiller         float64 `json:"goblin_killer"`
						Mole                 float64 `json:"mole"`
						Fortunate            float64 `json:"fortunate"`
						Professional         float64 `json:"professional"`
						PrecisionMining      float64 `json:"precision_mining"`
						GreatExplorer        float64 `json:"great_explorer"`
						LonesomeMiner        float64 `json:"lonesome_miner"`
						MiningSpeed2         float64 `json:"mining_speed_2"`
						PowderBuff           float64 `json:"powder_buff"`
						VeinSeeker           float64 `json:"vein_seeker"`
						ToggleFortunate      bool    `json:"toggle_fortunate"`
						ToggleMiningSpeed    bool    `json:"toggle_mining_speed"`
						MiningFortune2       float64 `json:"mining_fortune_2"`
						ToggleEfficientMiner bool    `json:"toggle_efficient_miner"`
					} `json:"nodes"`
					ReceivedFreeTier bool `json:"received_free_tier"`
					Crystals         struct {
						JadeCrystal struct {
							State       string  `json:"state"`
							TotalPlaced float64 `json:"total_placed"`
						} `json:"jade_crystal"`
						AmberCrystal struct {
							State       string  `json:"state"`
							TotalPlaced float64 `json:"total_placed"`
						} `json:"amber_crystal"`
						AmethystCrystal struct {
							State       string  `json:"state"`
							TotalPlaced float64 `json:"total_placed"`
						} `json:"amethyst_crystal"`
						SapphireCrystal struct {
							State       string  `json:"state"`
							TotalPlaced float64 `json:"total_placed"`
						} `json:"sapphire_crystal"`
						TopazCrystal struct {
							State       string  `json:"state"`
							TotalPlaced float64 `json:"total_placed"`
						} `json:"topaz_crystal"`
						JasperCrystal struct {
							State string `json:"state"`
						} `json:"jasper_crystal"`
						RubyCrystal struct {
							State string `json:"state"`
						} `json:"ruby_crystal"`
					} `json:"crystals"`
					PowderMithril               float64 `json:"powder_mithril"`
					PowderMithrilTotal          float64 `json:"powder_mithril_total"`
					Experience                  float64 `json:"experience"`
					TokensSpent                 float64 `json:"tokens_spent"`
					PowderSpentMithril          float64 `json:"powder_spent_mithril"`
					RetroactiveTier2Token       bool    `json:"retroactive_tier2_token"`
					Tokens                      float64 `json:"tokens"`
					SelectedPickaxeAbility      string  `json:"selected_pickaxe_ability"`
					DailyOresMinedDayMithrilOre float64 `json:"daily_ores_mined_day_mithril_ore"`
					DailyOresMinedMithrilOre    float64 `json:"daily_ores_mined_mithril_ore"`
					LastReset                   float64 `json:"last_reset"`
					GreaterMinesLastAccess      float64 `json:"greater_mines_last_access"`
					Biomes                      struct {
						Dwarven struct {
							StatuesPlaced []interface{} `json:"statues_placed"`
						} `json:"dwarven"`
						Precursor struct {
							PartsDelivered []interface{} `json:"parts_delivered"`
						} `json:"precursor"`
						Goblin struct {
							KingQuestActive     bool    `json:"king_quest_active"`
							KingQuestsCompleted float64 `json:"king_quests_completed"`
						} `json:"goblin"`
					} `json:"biomes"`
					PowderGemstone                float64 `json:"powder_gemstone"`
					PowderGemstoneTotal           float64 `json:"powder_gemstone_total"`
					DailyOresMinedDay             float64 `json:"daily_ores_mined_day"`
					DailyOresMined                float64 `json:"daily_ores_mined"`
					DailyOresMinedDayGemstone     float64 `json:"daily_ores_mined_day_gemstone"`
					DailyOresMinedGemstone        float64 `json:"daily_ores_mined_gemstone"`
					PowderSpentGemstone           float64 `json:"powder_spent_gemstone"`
					CurrentDailyEffect            string  `json:"current_daily_effect"`
					CurrentDailyEffectLastChanged float64 `json:"current_daily_effect_last_changed"`
				} `json:"mining_core"`
				Forge struct {
					ForgeProcesses struct {
						Forge1 struct {
							Num1 struct {
								Type      string  `json:"type"`
								ID        string  `json:"id"`
								StartTime float64 `json:"startTime"`
								Slot      float64 `json:"slot"`
								Notified  bool    `json:"notified"`
							} `json:"1"`
							Num2 struct {
								Type      string  `json:"type"`
								ID        string  `json:"id"`
								StartTime float64 `json:"startTime"`
								Slot      float64 `json:"slot"`
								Notified  bool    `json:"notified"`
							} `json:"2"`
							Num3 struct {
								Type      string  `json:"type"`
								ID        string  `json:"id"`
								StartTime float64 `json:"startTime"`
								Slot      float64 `json:"slot"`
								Notified  bool    `json:"notified"`
							} `json:"3"`
							Num4 struct {
								Type      string  `json:"type"`
								ID        string  `json:"id"`
								StartTime float64 `json:"startTime"`
								Slot      float64 `json:"slot"`
								Notified  bool    `json:"notified"`
							} `json:"4"`
							Num5 struct {
								Type      string  `json:"type"`
								ID        string  `json:"id"`
								StartTime float64 `json:"startTime"`
								Slot      float64 `json:"slot"`
								Notified  bool    `json:"notified"`
							} `json:"5"`
						} `json:"forge_1"`
					} `json:"forge_processes"`
				} `json:"forge"`
				EssenceUndead  float64 `json:"essence_undead"`
				EssenceDiamond float64 `json:"essence_diamond"`
				EssenceDragon  float64 `json:"essence_dragon"`
				EssenceWither  float64 `json:"essence_wither"`
				EssenceSpider  float64 `json:"essence_spider"`
				TempStatBuffs  []struct {
					Stat     float64 `json:"stat"`
					Key      string  `json:"key"`
					Amount   float64 `json:"amount"`
					ExpireAt float64 `json:"expire_at"`
				} `json:"temp_stat_buffs"`
				EssenceGold           float64 `json:"essence_gold"`
				EssenceIce            float64 `json:"essence_ice"`
				ExperienceSkillSocial float64 `json:"experience_skill_social"`
			} `json:"uuid"`
		} `json:"members"`
		Banking struct {
			Balance      float64 `json:"balance"`
			Transactions []struct {
				Amount        float64 `json:"amount"`
				Timestamp     float64 `json:"timestamp"`
				Action        string  `json:"action"`
				InitiatorName string  `json:"initiator_name"`
			} `json:"transactions"`
		} `json:"banking"`
		CommunityUpgrades struct {
			CurrentlyUpgrading struct {
				Upgrade    string  `json:"upgrade"`
				NewTier    float64 `json:"new_tier"`
				StartMs    int64   `json:"start_ms"`
				WhoStarted string  `json:"who_started"`
			} `json:"currently_upgrading"`
			UpgradeStates []struct {
				Upgrade     string  `json:"upgrade"`
				Tier        float64 `json:"tier"`
				StartedMs   float64 `json:"started_ms"`
				StartedBy   string  `json:"started_by"`
				ClaimedMs   float64 `json:"claimed_ms"`
				ClaimedBy   string  `json:"claimed_by"`
				Fasttracked bool    `json:"fasttracked"`
			} `json:"upgrade_states"`
		} `json:"community_upgrades"`
	} `json:"profile"`
}
