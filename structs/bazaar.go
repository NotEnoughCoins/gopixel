package structs

type Sellsummary []struct {
	Amount       int     `json:"amount"`
	PricePerUnit float64 `json:"pricePerUnit"`
	Orders       int     `json:"orders"`
}

type BuySummary []struct {
	Amount       int     `json:"amount"`
	PricePerUnit float64 `json:"pricePerUnit"`
	Orders       int     `json:"orders"`
}

type QuickStatus struct {
	ProductID      string  `json:"productId"`
	SellPrice      float64 `json:"sellPrice"`
	SellVolume     int     `json:"sellVolume"`
	SellMovingWeek int     `json:"sellMovingWeek"`
	SellOrders     int     `json:"sellOrders"`
	BuyPrice       float64 `json:"buyPrice"`
	BuyVolume      int     `json:"buyVolume"`
	BuyMovingWeek  int     `json:"buyMovingWeek"`
	BuyOrders      int     `json:"buyOrders"`
}

type Product struct {
	ProductID   string `json:"product_id"`
	Sellsummary `json:"sell_summary"`
	BuySummary  `json:"buy_summary"`
	QuickStatus `json:"quick_status"`
}

type Bazaar struct {
	Success     bool   `json:"success"`
	Cause       string `json:"cause"`
	LastUpdated int64  `json:"lastUpdated"`
	Products    struct {
		INKSACK3 struct {
			Product
		} `json:"INK_SACK:3"`
		BrownMushroom struct {
			Product
		} `json:"BROWN_MUSHROOM"`
		SpookyShard struct {
			Product
		} `json:"SPOOKY_SHARD"`
		INKSACK4 struct {
			Product
		} `json:"INK_SACK:4"`
		TarantulaWeb struct {
			Product
		} `json:"TARANTULA_WEB"`
		CarrotItem struct {
			Product
		} `json:"CARROT_ITEM"`
		EnchantedPotato struct {
			Product
		} `json:"ENCHANTED_POTATO"`
		ExpBottle struct {
			Product
		} `json:"EXP_BOTTLE"`
		JerryBoxGreen struct {
			Product
		} `json:"JERRY_BOX_GREEN"`
		EnchantedSlimeBall struct {
			Product
		} `json:"ENCHANTED_SLIME_BALL"`
		EnchantedRedMushroom struct {
			Product
		} `json:"ENCHANTED_RED_MUSHROOM"`
		EnchantedGoldenCarrot struct {
			Product
		} `json:"ENCHANTED_GOLDEN_CARROT"`
		EnchantedRabbitHide struct {
			Product
		} `json:"ENCHANTED_RABBIT_HIDE"`
		FlawedAmethystGem struct {
			Product
		} `json:"FLAWED_AMETHYST_GEM"`
		PerfectJadeGem struct {
			Product
		} `json:"PERFECT_JADE_GEM"`
		EnchantedBirchLog struct {
			Product
		} `json:"ENCHANTED_BIRCH_LOG"`
		EnchantedGunpowder struct {
			Product
		} `json:"ENCHANTED_GUNPOWDER"`
		EnchantedMelon struct {
			Product
		} `json:"ENCHANTED_MELON"`
		EnchantedSugar struct {
			Product
		} `json:"ENCHANTED_SUGAR"`
		Cactus struct {
			Product
		} `json:"CACTUS"`
		FlawedJasperGem struct {
			Product
		} `json:"FLAWED_JASPER_GEM"`
		EnchantedBlazeRod struct {
			Product
		} `json:"ENCHANTED_BLAZE_ROD"`
		EnchantedCake struct {
			Product
		} `json:"ENCHANTED_CAKE"`
		Pumpkin struct {
			Product
		} `json:"PUMPKIN"`
		EnchantedBrownMushroom struct {
			Product
		} `json:"ENCHANTED_BROWN_MUSHROOM"`
		Wheat struct {
			Product
		} `json:"WHEAT"`
		NurseSharkTooth struct {
			Product
		} `json:"NURSE_SHARK_TOOTH"`
		FlawedAmberGem struct {
			Product
		} `json:"FLAWED_AMBER_GEM"`
		EnchantedRawSalmon struct {
			Product
		} `json:"ENCHANTED_RAW_SALMON"`
		EnchantedGlisteringMelon struct {
			Product
		} `json:"ENCHANTED_GLISTERING_MELON"`
		PrismarineShard struct {
			Product
		} `json:"PRISMARINE_SHARD"`
		EnchantedEmerald struct {
			Product
		} `json:"ENCHANTED_EMERALD"`
		ProtectorFragment struct {
			Product
		} `json:"PROTECTOR_FRAGMENT"`
		EnchantedSpiderEye struct {
			Product
		} `json:"ENCHANTED_SPIDER_EYE"`
		RedMushroom struct {
			Product
		} `json:"RED_MUSHROOM"`
		GrandExpBottle struct {
			Product
		} `json:"GRAND_EXP_BOTTLE"`
		Mutton struct {
			Product
		} `json:"MUTTON"`
		EnchantedMelonBlock struct {
			Product
		} `json:"ENCHANTED_MELON_BLOCK"`
		PowerCrystal struct {
			Product
		} `json:"POWER_CRYSTAL"`
		RawSoulflow struct {
			Product
		} `json:"RAW_SOULFLOW"`
		SharkFin struct {
			Product
		} `json:"SHARK_FIN"`
		Diamond struct {
			Product
		} `json:"DIAMOND"`
		WiseFragment struct {
			Product
		} `json:"WISE_FRAGMENT"`
		Cobblestone struct {
			Product
		} `json:"COBBLESTONE"`
		RefinedMithril struct {
			Product
		} `json:"REFINED_MITHRIL"`
		SpiderEye struct {
			Product
		} `json:"SPIDER_EYE"`
		RawFish struct {
			Product
		} `json:"RAW_FISH"`
		EnchantedPufferfish struct {
			Product
		} `json:"ENCHANTED_PUFFERFISH"`
		PerfectRubyGem struct {
			Product
		} `json:"PERFECT_RUBY_GEM"`
		Yoggie struct {
			Product
		} `json:"YOGGIE"`
		PerfectJasperGem struct {
			Product
		} `json:"PERFECT_JASPER_GEM"`
		PotatoItem struct {
			Product
		} `json:"POTATO_ITEM"`
		EnchantedNetherrack struct {
			Product
		} `json:"ENCHANTED_NETHERRACK"`
		EnchantedHardStone struct {
			Product
		} `json:"ENCHANTED_HARD_STONE"`
		EnchantedHugeMushroom1 struct {
			Product
		} `json:"ENCHANTED_HUGE_MUSHROOM_1"`
		RefinedDiamond struct {
			Product
		} `json:"REFINED_DIAMOND"`
		TightlyTiedHayBale struct {
			Product
		} `json:"TIGHTLY_TIED_HAY_BALE"`
		EnchantedCobblestone struct {
			Product
		} `json:"ENCHANTED_COBBLESTONE"`
		EnchantedHugeMushroom2 struct {
			Product
		} `json:"ENCHANTED_HUGE_MUSHROOM_2"`
		Pork struct {
			Product
		} `json:"PORK"`
		PrismarineCrystals struct {
			Product
		} `json:"PRISMARINE_CRYSTALS"`
		Ice struct {
			Product
		} `json:"ICE"`
		TigerSharkTooth struct {
			Product
		} `json:"TIGER_SHARK_TOOTH"`
		HugeMushroom1 struct {
			Product
		} `json:"HUGE_MUSHROOM_1"`
		IceBait struct {
			Product
		} `json:"ICE_BAIT"`
		HugeMushroom2 struct {
			Product
		} `json:"HUGE_MUSHROOM_2"`
		LOG21 struct {
			Product
		} `json:"LOG_2:1"`
		EnchantedSnowBlock struct {
			Product
		} `json:"ENCHANTED_SNOW_BLOCK"`
		GoldenTooth struct {
			Product
		} `json:"GOLDEN_TOOTH"`
		String struct {
			Product
		} `json:"STRING"`
		HyperCatalyst struct {
			Product
		} `json:"HYPER_CATALYST"`
		RabbitFoot struct {
			Product
		} `json:"RABBIT_FOOT"`
		Redstone struct {
			Product
		} `json:"REDSTONE"`
		JerryBoxGolden struct {
			Product
		} `json:"JERRY_BOX_GOLDEN"`
		PumpkinGuts struct {
			Product
		} `json:"PUMPKIN_GUTS"`
		EnchantedCactusGreen struct {
			Product
		} `json:"ENCHANTED_CACTUS_GREEN"`
		BoosterCookie struct {
			Product
		} `json:"BOOSTER_COOKIE"`
		EnchantedCarrotOnAStick struct {
			Product
		} `json:"ENCHANTED_CARROT_ON_A_STICK"`
		EnchantedLapisLazuliBlock struct {
			Product
		} `json:"ENCHANTED_LAPIS_LAZULI_BLOCK"`
		EnchantedCookie struct {
			Product
		} `json:"ENCHANTED_COOKIE"`
		EnchantedEndstone struct {
			Product
		} `json:"ENCHANTED_ENDSTONE"`
		ColossalExpBottle struct {
			Product
		} `json:"COLOSSAL_EXP_BOTTLE"`
		EnchantedSand struct {
			Product
		} `json:"ENCHANTED_SAND"`
		EnchantedString struct {
			Product
		} `json:"ENCHANTED_STRING"`
		StrongFragment struct {
			Product
		} `json:"STRONG_FRAGMENT"`
		SlimeBall struct {
			Product
		} `json:"SLIME_BALL"`
		SnowBall struct {
			Product
		} `json:"SNOW_BALL"`
		HolyFragment struct {
			Product
		} `json:"HOLY_FRAGMENT"`
		EnchantedAcaciaLog struct {
			Product
		} `json:"ENCHANTED_ACACIA_LOG"`
		EnchantedEgg struct {
			Product
		} `json:"ENCHANTED_EGG"`
		Sand struct {
			Product
		} `json:"SAND"`
		FlawedRubyGem struct {
			Product
		} `json:"FLAWED_RUBY_GEM"`
		SoulFragment struct {
			Product
		} `json:"SOUL_FRAGMENT"`
		FineJadeGem struct {
			Product
		} `json:"FINE_JADE_GEM"`
		RawChicken struct {
			Product
		} `json:"RAW_CHICKEN"`
		PlasmaBucket struct {
			Product
		} `json:"PLASMA_BUCKET"`
		FlawlessJasperGem struct {
			Product
		} `json:"FLAWLESS_JASPER_GEM"`
		AncientClaw struct {
			Product
		} `json:"ANCIENT_CLAW"`
		EnchantedGhastTear struct {
			Product
		} `json:"ENCHANTED_GHAST_TEAR"`
		EnchantedLapisLazuli struct {
			Product
		} `json:"ENCHANTED_LAPIS_LAZULI"`
		EnchantedCocoa struct {
			Product
		} `json:"ENCHANTED_COCOA"`
		CarrotBait struct {
			Product
		} `json:"CARROT_BAIT"`
		FineTopazGem struct {
			Product
		} `json:"FINE_TOPAZ_GEM"`
		Seeds struct {
			Product
		} `json:"SEEDS"`
		FineRubyGem struct {
			Product
		} `json:"FINE_RUBY_GEM"`
		EnchantedLeather struct {
			Product
		} `json:"ENCHANTED_LEATHER"`
		EnchantedSponge struct {
			Product
		} `json:"ENCHANTED_SPONGE"`
		EnchantedSharkFin struct {
			Product
		} `json:"ENCHANTED_SHARK_FIN"`
		PerfectAmberGem struct {
			Product
		} `json:"PERFECT_AMBER_GEM"`
		HayBlock struct {
			Product
		} `json:"HAY_BLOCK"`
		Flint struct {
			Product
		} `json:"FLINT"`
		InkSack struct {
			Product
		} `json:"INK_SACK"`
		EnchantedSpruceLog struct {
			Product
		} `json:"ENCHANTED_SPRUCE_LOG"`
		EnchantedRottenFlesh struct {
			Product
		} `json:"ENCHANTED_ROTTEN_FLESH"`
		WolfTooth struct {
			Product
		} `json:"WOLF_TOOTH"`
		EnchantedGrilledPork struct {
			Product
		} `json:"ENCHANTED_GRILLED_PORK"`
		EnchantedNetherStalk struct {
			Product
		} `json:"ENCHANTED_NETHER_STALK"`
		EnchantedRedstoneBlock struct {
			Product
		} `json:"ENCHANTED_REDSTONE_BLOCK"`
		EnchantedQuartzBlock struct {
			Product
		} `json:"ENCHANTED_QUARTZ_BLOCK"`
		EnchantedAncientClaw struct {
			Product
		} `json:"ENCHANTED_ANCIENT_CLAW"`
		GreenCandy struct {
			Product
		} `json:"GREEN_CANDY"`
		EnchantedRedstone struct {
			Product
		} `json:"ENCHANTED_REDSTONE"`
		EnchantedRedstoneLamp struct {
			Product
		} `json:"ENCHANTED_REDSTONE_LAMP"`
		Treasurite struct {
			Product
		} `json:"TREASURITE"`
		DwarvenCompactor struct {
			Product
		} `json:"DWARVEN_COMPACTOR"`
		GreatWhiteSharkTooth struct {
			Product
		} `json:"GREAT_WHITE_SHARK_TOOTH"`
		Gravel struct {
			Product
		} `json:"GRAVEL"`
		Melon struct {
			Product
		} `json:"MELON"`
		EnchantedLavaBucket struct {
			Product
		} `json:"ENCHANTED_LAVA_BUCKET"`
		EnchantedPackedIce struct {
			Product
		} `json:"ENCHANTED_PACKED_ICE"`
		RAWFISH3 struct {
			Product
		} `json:"RAW_FISH:3"`
		EnchantedPrismarineShard struct {
			Product
		} `json:"ENCHANTED_PRISMARINE_SHARD"`
		EnchantedCarrotStick struct {
			Product
		} `json:"ENCHANTED_CARROT_STICK"`
		Recombobulator3000 struct {
			Product
		} `json:"RECOMBOBULATOR_3000"`
		EnchantedIronBlock struct {
			Product
		} `json:"ENCHANTED_IRON_BLOCK"`
		Bone struct {
			Product
		} `json:"BONE"`
		RAWFISH2 struct {
			Product
		} `json:"RAW_FISH:2"`
		RAWFISH1 struct {
			Product
		} `json:"RAW_FISH:1"`
		RevenantFlesh struct {
			Product
		} `json:"REVENANT_FLESH"`
		EnchantedGlowstone struct {
			Product
		} `json:"ENCHANTED_GLOWSTONE"`
		EnchantedPork struct {
			Product
		} `json:"ENCHANTED_PORK"`
		RoughJasperGem struct {
			Product
		} `json:"ROUGH_JASPER_GEM"`
		Feather struct {
			Product
		} `json:"FEATHER"`
		WhaleBait struct {
			Product
		} `json:"WHALE_BAIT"`
		Netherrack struct {
			Product
		} `json:"NETHERRACK"`
		Sponge struct {
			Product
		} `json:"SPONGE"`
		BlazeRod struct {
			Product
		} `json:"BLAZE_ROD"`
		EnchantedDarkOakLog struct {
			Product
		} `json:"ENCHANTED_DARK_OAK_LOG"`
		FlawlessTopazGem struct {
			Product
		} `json:"FLAWLESS_TOPAZ_GEM"`
		YoungFragment struct {
			Product
		} `json:"YOUNG_FRAGMENT"`
		EnchantedClownfish struct {
			Product
		} `json:"ENCHANTED_CLOWNFISH"`
		RefinedMineral struct {
			Product
		} `json:"REFINED_MINERAL"`
		EnchantedGold struct {
			Product
		} `json:"ENCHANTED_GOLD"`
		EnchantedRawChicken struct {
			Product
		} `json:"ENCHANTED_RAW_CHICKEN"`
		EnchantedWaterLily struct {
			Product
		} `json:"ENCHANTED_WATER_LILY"`
		RoughAmethystGem struct {
			Product
		} `json:"ROUGH_AMETHYST_GEM"`
		RoughRubyGem struct {
			Product
		} `json:"ROUGH_RUBY_GEM"`
		LOG1 struct {
			Product
		} `json:"LOG:1"`
		FlawlessRubyGem struct {
			Product
		} `json:"FLAWLESS_RUBY_GEM"`
		NullAtom struct {
			Product
		} `json:"NULL_ATOM"`
		TitaniumOre struct {
			Product
		} `json:"TITANIUM_ORE"`
		BlueSharkTooth struct {
			Product
		} `json:"BLUE_SHARK_TOOTH"`
		Catalyst struct {
			Product
		} `json:"CATALYST"`
		LOG3 struct {
			Product
		} `json:"LOG:3"`
		LOG2 struct {
			Product
		} `json:"LOG:2"`
		BlessedBait struct {
			Product
		} `json:"BLESSED_BAIT"`
		EnchantedGlowstoneDust struct {
			Product
		} `json:"ENCHANTED_GLOWSTONE_DUST"`
		EnchantedInkSack struct {
			Product
		} `json:"ENCHANTED_INK_SACK"`
		EnchantedCactus struct {
			Product
		} `json:"ENCHANTED_CACTUS"`
		EnchantedSugarCane struct {
			Product
		} `json:"ENCHANTED_SUGAR_CANE"`
		FlawlessSapphireGem struct {
			Product
		} `json:"FLAWLESS_SAPPHIRE_GEM"`
		EnchantedCookedSalmon struct {
			Product
		} `json:"ENCHANTED_COOKED_SALMON"`
		EnchantedSeeds struct {
			Product
		} `json:"ENCHANTED_SEEDS"`
		ConcentratedStone struct {
			Product
		} `json:"CONCENTRATED_STONE"`
		Log struct {
			Product
		} `json:"LOG"`
		JacobsTicket struct {
			Product
		} `json:"JACOBS_TICKET"`
		EnchantedBoneBlock struct {
			Product
		} `json:"ENCHANTED_BONE_BLOCK"`
		GhastTear struct {
			Product
		} `json:"GHAST_TEAR"`
		AbsoluteEnderPearl struct {
			Product
		} `json:"ABSOLUTE_ENDER_PEARL"`
		UnstableFragment struct {
			Product
		} `json:"UNSTABLE_FRAGMENT"`
		EnchantedEnderPearl struct {
			Product
		} `json:"ENCHANTED_ENDER_PEARL"`
		PurpleCandy struct {
			Product
		} `json:"PURPLE_CANDY"`
		SpikedBait struct {
			Product
		} `json:"SPIKED_BAIT"`
		EnchantedFermentedSpiderEye struct {
			Product
		} `json:"ENCHANTED_FERMENTED_SPIDER_EYE"`
		PolishedPumpkin struct {
			Product
		} `json:"POLISHED_PUMPKIN"`
		EnchantedGoldBlock struct {
			Product
		} `json:"ENCHANTED_GOLD_BLOCK"`
		RoughJadeGem struct {
			Product
		} `json:"ROUGH_JADE_GEM"`
		EnchantedJungleLog struct {
			Product
		} `json:"ENCHANTED_JUNGLE_LOG"`
		EnchantedFlint struct {
			Product
		} `json:"ENCHANTED_FLINT"`
		IronIngot struct {
			Product
		} `json:"IRON_INGOT"`
		EnchantedEmeraldBlock struct {
			Product
		} `json:"ENCHANTED_EMERALD_BLOCK"`
		NullOvoid struct {
			Product
		} `json:"NULL_OVOID"`
		EnchantedClayBall struct {
			Product
		} `json:"ENCHANTED_CLAY_BALL"`
		RoughSapphireGem struct {
			Product
		} `json:"ROUGH_SAPPHIRE_GEM"`
		GlowstoneDust struct {
			Product
		} `json:"GLOWSTONE_DUST"`
		GoldIngot struct {
			Product
		} `json:"GOLD_INGOT"`
		RevenantViscera struct {
			Product
		} `json:"REVENANT_VISCERA"`
		TarantulaSilk struct {
			Product
		} `json:"TARANTULA_SILK"`
		PerfectAmethystGem struct {
			Product
		} `json:"PERFECT_AMETHYST_GEM"`
		TitanicExpBottle struct {
			Product
		} `json:"TITANIC_EXP_BOTTLE"`
		EnchantedMutton struct {
			Product
		} `json:"ENCHANTED_MUTTON"`
		NullSphere struct {
			Product
		} `json:"NULL_SPHERE"`
		SuperCompactor3000 struct {
			Product
		} `json:"SUPER_COMPACTOR_3000"`
		EnchantedIron struct {
			Product
		} `json:"ENCHANTED_IRON"`
		SuperEgg struct {
			Product
		} `json:"SUPER_EGG"`
		StockOfStonks struct {
			Product
		} `json:"STOCK_OF_STONKS"`
		MithrilOre struct {
			Product
		} `json:"MITHRIL_ORE"`
		EnchantedHayBlock struct {
			Product
		} `json:"ENCHANTED_HAY_BLOCK"`
		EnchantedBone struct {
			Product
		} `json:"ENCHANTED_BONE"`
		EnchantedPaper struct {
			Product
		} `json:"ENCHANTED_PAPER"`
		EnchantedTitanium struct {
			Product
		} `json:"ENCHANTED_TITANIUM"`
		EnchantedDiamondBlock struct {
			Product
		} `json:"ENCHANTED_DIAMOND_BLOCK"`
		SpookyBait struct {
			Product
		} `json:"SPOOKY_BAIT"`
		SuperiorFragment struct {
			Product
		} `json:"SUPERIOR_FRAGMENT"`
		MagmaBucket struct {
			Product
		} `json:"MAGMA_BUCKET"`
		Emerald struct {
			Product
		} `json:"EMERALD"`
		EnchantedRabbitFoot struct {
			Product
		} `json:"ENCHANTED_RABBIT_FOOT"`
		LightBait struct {
			Product
		} `json:"LIGHT_BAIT"`
		HotPotatoBook struct {
			Product
		} `json:"HOT_POTATO_BOOK"`
		EnchantedIce struct {
			Product
		} `json:"ENCHANTED_ICE"`
		ClayBall struct {
			Product
		} `json:"CLAY_BALL"`
		OldFragment struct {
			Product
		} `json:"OLD_FRAGMENT"`
		GreenGift struct {
			Product
		} `json:"GREEN_GIFT"`
		FlawlessAmethystGem struct {
			Product
		} `json:"FLAWLESS_AMETHYST_GEM"`
		RoughTopazGem struct {
			Product
		} `json:"ROUGH_TOPAZ_GEM"`
		PackedIce struct {
			Product
		} `json:"PACKED_ICE"`
		RoughAmberGem struct {
			Product
		} `json:"ROUGH_AMBER_GEM"`
		WaterLily struct {
			Product
		} `json:"WATER_LILY"`
		Log2 struct {
			Product
		} `json:"LOG_2"`
		HamsterWheel struct {
			Product
		} `json:"HAMSTER_WHEEL"`
		EnchantedObsidian struct {
			Product
		} `json:"ENCHANTED_OBSIDIAN"`
		FineAmberGem struct {
			Product
		} `json:"FINE_AMBER_GEM"`
		EnchantedCoal struct {
			Product
		} `json:"ENCHANTED_COAL"`
		Coal struct {
			Product
		} `json:"COAL"`
		EnchantedQuartz struct {
			Product
		} `json:"ENCHANTED_QUARTZ"`
		EnderPearl struct {
			Product
		} `json:"ENDER_PEARL"`
		EnchantedCoalBlock struct {
			Product
		} `json:"ENCHANTED_COAL_BLOCK"`
		WerewolfSkin struct {
			Product
		} `json:"WEREWOLF_SKIN"`
		EnchantedPrismarineCrystals struct {
			Product
		} `json:"ENCHANTED_PRISMARINE_CRYSTALS"`
		PerfectTopazGem struct {
			Product
		} `json:"PERFECT_TOPAZ_GEM"`
		DaedalusStick struct {
			Product
		} `json:"DAEDALUS_STICK"`
		EnchantedWetSponge struct {
			Product
		} `json:"ENCHANTED_WET_SPONGE"`
		FlawedJadeGem struct {
			Product
		} `json:"FLAWED_JADE_GEM"`
		EnderStone struct {
			Product
		} `json:"ENDER_STONE"`
		EnchantedRawFish struct {
			Product
		} `json:"ENCHANTED_RAW_FISH"`
		Quartz struct {
			Product
		} `json:"QUARTZ"`
		FoulFlesh struct {
			Product
		} `json:"FOUL_FLESH"`
		JerryBoxPurple struct {
			Product
		} `json:"JERRY_BOX_PURPLE"`
		RawBeef struct {
			Product
		} `json:"RAW_BEEF"`
		SludgeJuice struct {
			Product
		} `json:"SLUDGE_JUICE"`
		EnchantedEyeOfEnder struct {
			Product
		} `json:"ENCHANTED_EYE_OF_ENDER"`
		Ectoplasm struct {
			Product
		} `json:"ECTOPLASM"`
		SugarCane struct {
			Product
		} `json:"SUGAR_CANE"`
		MagmaCream struct {
			Product
		} `json:"MAGMA_CREAM"`
		SharkBait struct {
			Product
		} `json:"SHARK_BAIT"`
		RedGift struct {
			Product
		} `json:"RED_GIFT"`
		EnchantedMithril struct {
			Product
		} `json:"ENCHANTED_MITHRIL"`
		JerryBoxBlue struct {
			Product
		} `json:"JERRY_BOX_BLUE"`
		EnchantedRawBeef struct {
			Product
		} `json:"ENCHANTED_RAW_BEEF"`
		EnchantedSlimeBlock struct {
			Product
		} `json:"ENCHANTED_SLIME_BLOCK"`
		EnchantedFeather struct {
			Product
		} `json:"ENCHANTED_FEATHER"`
		EnchantedOakLog struct {
			Product
		} `json:"ENCHANTED_OAK_LOG"`
		RabbitHide struct {
			Product
		} `json:"RABBIT_HIDE"`
		WhiteGift struct {
			Product
		} `json:"WHITE_GIFT"`
		Sulphur struct {
			Product
		} `json:"SULPHUR"`
		Rabbit struct {
			Product
		} `json:"RABBIT"`
		FineSapphireGem struct {
			Product
		} `json:"FINE_SAPPHIRE_GEM"`
		NetherStalk struct {
			Product
		} `json:"NETHER_STALK"`
		DarkBait struct {
			Product
		} `json:"DARK_BAIT"`
		EnchantedCarrot struct {
			Product
		} `json:"ENCHANTED_CARROT"`
		EnchantedPumpkin struct {
			Product
		} `json:"ENCHANTED_PUMPKIN"`
		GriffinFeather struct {
			Product
		} `json:"GRIFFIN_FEATHER"`
		RottenFlesh struct {
			Product
		} `json:"ROTTEN_FLESH"`
		EnchantedCookedFish struct {
			Product
		} `json:"ENCHANTED_COOKED_FISH"`
		Obsidian struct {
			Product
		} `json:"OBSIDIAN"`
		Soulflow struct {
			Product
		} `json:"SOULFLOW"`
		MinnowBait struct {
			Product
		} `json:"MINNOW_BAIT"`
		EnchantedMagmaCream struct {
			Product
		} `json:"ENCHANTED_MAGMA_CREAM"`
		EnchantedFireworkRocket struct {
			Product
		} `json:"ENCHANTED_FIREWORK_ROCKET"`
		Starfall struct {
			Product
		} `json:"STARFALL"`
		FlawedTopazGem struct {
			Product
		} `json:"FLAWED_TOPAZ_GEM"`
		HardStone struct {
			Product
		} `json:"HARD_STONE"`
		FlawlessJadeGem struct {
			Product
		} `json:"FLAWLESS_JADE_GEM"`
		Leather struct {
			Product
		} `json:"LEATHER"`
		EnchantedCookedMutton struct {
			Product
		} `json:"ENCHANTED_COOKED_MUTTON"`
		FineAmethystGem struct {
			Product
		} `json:"FINE_AMETHYST_GEM"`
		RefinedTitanium struct {
			Product
		} `json:"REFINED_TITANIUM"`
		EnchantedRabbit struct {
			Product
		} `json:"ENCHANTED_RABBIT"`
		MutantNetherStalk struct {
			Product
		} `json:"MUTANT_NETHER_STALK"`
		EnchantedBread struct {
			Product
		} `json:"ENCHANTED_BREAD"`
		FumingPotatoBook struct {
			Product
		} `json:"FUMING_POTATO_BOOK"`
		FineJasperGem struct {
			Product
		} `json:"FINE_JASPER_GEM"`
		FlawedSapphireGem struct {
			Product
		} `json:"FLAWED_SAPPHIRE_GEM"`
		EnchantedCharcoal struct {
			Product
		} `json:"ENCHANTED_CHARCOAL"`
		FlawlessAmberGem struct {
			Product
		} `json:"FLAWLESS_AMBER_GEM"`
		EnchantedBlazePowder struct {
			Product
		} `json:"ENCHANTED_BLAZE_POWDER"`
		SummoningEye struct {
			Product
		} `json:"SUMMONING_EYE"`
		PerfectSapphireGem struct {
			Product
		} `json:"PERFECT_SAPPHIRE_GEM"`
		FishBait struct {
			Product
		} `json:"FISH_BAIT"`
		SnowBlock struct {
			Product
		} `json:"SNOW_BLOCK"`
		EnchantedBakedPotato struct {
			Product
		} `json:"ENCHANTED_BAKED_POTATO"`
		Compactor struct {
			Product
		} `json:"COMPACTOR"`
		EnchantedDiamond struct {
			Product
		} `json:"ENCHANTED_DIAMOND"`
		BazaarCookie struct {
			Product
		} `json:"BAZAAR_COOKIE"`
	} `json:"products"`
}
