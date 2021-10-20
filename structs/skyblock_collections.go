package structs

type SkyblockCollections struct {
	Success     bool   `json:"success"`
	LastUpdated int64  `json:"lastUpdated"`
	Version     string `json:"version"`
	Collections struct {
		Farming struct {
			Name  string `json:"name"`
			Items struct {
				INKSACK3 struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"INK_SACK:3"`
				CarrotItem struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"CARROT_ITEM"`
				Cactus struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"CACTUS"`
				RawChicken struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"RAW_CHICKEN"`
				SugarCane struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"SUGAR_CANE"`
				Pumpkin struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"PUMPKIN"`
				Wheat struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"WHEAT"`
				Seeds struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"SEEDS"`
				MushroomCollection struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"MUSHROOM_COLLECTION"`
				Rabbit struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"RABBIT"`
				NetherStalk struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"NETHER_STALK"`
				Mutton struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"MUTTON"`
				Melon struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"MELON"`
				PotatoItem struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"POTATO_ITEM"`
				Leather struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"LEATHER"`
				Pork struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"PORK"`
				Feather struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"FEATHER"`
			} `json:"items"`
		} `json:"FARMING"`
		Mining struct {
			Name  string `json:"name"`
			Items struct {
				INKSACK4 struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"INK_SACK:4"`
				Redstone struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"REDSTONE"`
				Coal struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"COAL"`
				EnderStone struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"ENDER_STONE"`
				Quartz struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"QUARTZ"`
				Sand struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"SAND"`
				IronIngot struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"IRON_INGOT"`
				GemstoneCollection struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"GEMSTONE_COLLECTION"`
				Obsidian struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"OBSIDIAN"`
				Diamond struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"DIAMOND"`
				Cobblestone struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"COBBLESTONE"`
				GlowstoneDust struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"GLOWSTONE_DUST"`
				GoldIngot struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"GOLD_INGOT"`
				Gravel struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"GRAVEL"`
				HardStone struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"HARD_STONE"`
				MithrilOre struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"MITHRIL_ORE"`
				Emerald struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"EMERALD"`
				Ice struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"ICE"`
				Netherrack struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"NETHERRACK"`
			} `json:"items"`
		} `json:"MINING"`
		Combat struct {
			Name  string `json:"name"`
			Items struct {
				EnderPearl struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"ENDER_PEARL"`
				SlimeBall struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"SLIME_BALL"`
				MagmaCream struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"MAGMA_CREAM"`
				GhastTear struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"GHAST_TEAR"`
				Sulphur struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"SULPHUR"`
				RottenFlesh struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"ROTTEN_FLESH"`
				SpiderEye struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"SPIDER_EYE"`
				Bone struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"BONE"`
				BlazeRod struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"BLAZE_ROD"`
				String struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"STRING"`
			} `json:"items"`
		} `json:"COMBAT"`
		Foraging struct {
			Name  string `json:"name"`
			Items struct {
				Log2 struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"LOG_2"`
				LOG1 struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"LOG:1"`
				LOG3 struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"LOG:3"`
				LOG2 struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"LOG:2"`
				Log struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"LOG"`
				LOG21 struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"LOG_2:1"`
			} `json:"items"`
		} `json:"FORAGING"`
		Fishing struct {
			Name  string `json:"name"`
			Items struct {
				WaterLily struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"WATER_LILY"`
				PrismarineShard struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"PRISMARINE_SHARD"`
				InkSack struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"INK_SACK"`
				RawFish struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"RAW_FISH"`
				RAWFISH3 struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"RAW_FISH:3"`
				RAWFISH2 struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"RAW_FISH:2"`
				RAWFISH1 struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"RAW_FISH:1"`
				PrismarineCrystals struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"PRISMARINE_CRYSTALS"`
				ClayBall struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"CLAY_BALL"`
				Sponge struct {
					Name     string `json:"name"`
					MaxTiers int    `json:"maxTiers"`
					Tiers    []struct {
						Tier           int      `json:"tier"`
						AmountRequired int      `json:"amountRequired"`
						Unlocks        []string `json:"unlocks"`
					} `json:"tiers"`
				} `json:"SPONGE"`
			} `json:"items"`
		} `json:"FISHING"`
	} `json:"collections"`
}
