package structs

type SkyblockItems struct {
	Data []struct {
		AbilityDamageScaling  float64 `json:"ability_damage_scaling"`
		CatacombsRequirements struct {
			Dungeon struct {
				Level int64  `json:"level"`
				Type  string `json:"type"`
			} `json:"dungeon"`
		} `json:"catacombs_requirements"`
		Category     string `json:"category"`
		Color        string `json:"color"`
		Crystal      string `json:"crystal"`
		Description  string `json:"description"`
		DungeonItem  bool   `json:"dungeon_item"`
		Durability   int64  `json:"durability"`
		Enchantments struct {
			AquaAffinity  int64 `json:"aqua_affinity"`
			BigBrain      int64 `json:"big_brain"`
			CounterStrike int64 `json:"counter_strike"`
			DepthStrider  int64 `json:"depth_strider"`
			Efficiency    int64 `json:"efficiency"`
			FirstStrike   int64 `json:"first_strike"`
			Power         int64 `json:"power"`
			Respiration   int64 `json:"respiration"`
			Scavenger     int64 `json:"scavenger"`
			Sharpness     int64 `json:"sharpness"`
			Telekinesis   int64 `json:"telekinesis"`
			Vampirism     int64 `json:"vampirism"`
			Vicious       int64 `json:"vicious"`
		} `json:"enchantments"`
		Essence struct {
			Costs       []int64 `json:"costs"`
			EssenceType string  `json:"essence_type"`
		} `json:"essence"`
		Furniture      string `json:"furniture"`
		GearScore      int64  `json:"gear_score"`
		Generator      string `json:"generator"`
		GeneratorTier  int64  `json:"generator_tier"`
		Glowing        bool   `json:"glowing"`
		ID             string `json:"id"`
		ItemDurability int64  `json:"item_durability"`
		Material       string `json:"material"`
		Name           string `json:"name"`
		NpcSellPrice   int64  `json:"npc_sell_price"`
		PrivateIsland  string `json:"private_island"`
		Requirements   struct {
			Collection struct {
				CollectionID string `json:"collection_id"`
				RequiredTier int64  `json:"required_tier"`
			} `json:"collection"`
			Dungeon struct {
				Level int64  `json:"level"`
				Type  string `json:"type"`
			} `json:"dungeon"`
			DungeonCompletion struct {
				Tier int64  `json:"tier"`
				Type string `json:"type"`
			} `json:"dungeon_completion"`
			HeartOfTheMountain struct {
				Tier int64 `json:"tier"`
			} `json:"heart_of_the_mountain"`
			Skill struct {
				Level int64  `json:"level"`
				Type  string `json:"type"`
			} `json:"skill"`
			Slayer struct {
				Level          int64  `json:"level"`
				SlayerBossType string `json:"slayer_boss_type"`
			} `json:"slayer"`
			TargetPractice struct {
				Mode string `json:"mode"`
			} `json:"targetPractice"`
		} `json:"requirements"`
		Skin  string `json:"skin"`
		Stats struct {
			AbilityDamagePercent int64   `json:"ABILITY_DAMAGE_PERCENT"`
			AttackSpeed          int64   `json:"ATTACK_SPEED"`
			BreakingPower        int64   `json:"BREAKING_POWER"`
			CriticalChance       int64   `json:"CRITICAL_CHANCE"`
			CriticalDamage       int64   `json:"CRITICAL_DAMAGE"`
			Damage               int64   `json:"DAMAGE"`
			Defense              int64   `json:"DEFENSE"`
			Ferocity             int64   `json:"FEROCITY"`
			Health               int64   `json:"HEALTH"`
			Intelligence         int64   `json:"INTELLIGENCE"`
			MagicFind            int64   `json:"MAGIC_FIND"`
			MiningFortune        int64   `json:"MINING_FORTUNE"`
			MiningSpeed          int64   `json:"MINING_SPEED"`
			SeaCreatureChance    float64 `json:"SEA_CREATURE_CHANCE"`
			Strength             int64   `json:"STRENGTH"`
			TrueDefense          int64   `json:"TRUE_DEFENSE"`
			WalkSpeed            int64   `json:"WALK_SPEED"`
			WeaponAbilityDamage  int64   `json:"WEAPON_ABILITY_DAMAGE"`
		} `json:"stats"`
		Tier        string `json:"tier"`
		TieredStats struct {
			AttackSpeed         []int64 `json:"ATTACK_SPEED"`
			CriticalChance      []int64 `json:"CRITICAL_CHANCE"`
			CriticalDamage      []int64 `json:"CRITICAL_DAMAGE"`
			Damage              []int64 `json:"DAMAGE"`
			Defense             []int64 `json:"DEFENSE"`
			Health              []int64 `json:"HEALTH"`
			Intelligence        []int64 `json:"INTELLIGENCE"`
			Strength            []int64 `json:"STRENGTH"`
			WalkSpeed           []int64 `json:"WALK_SPEED"`
			WeaponAbilityDamage []int64 `json:"WEAPON_ABILITY_DAMAGE"`
		} `json:"tiered_stats"`
		Unstackable bool `json:"unstackable"`
	} `json:"items"`
	LastUpdated int64 `json:"lastUpdated"`
	Success     bool  `json:"success"`
}
