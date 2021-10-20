package structs

type SkyblockItems struct {
	Data []struct {
		AbilityDamageScaling  float64 `json:"ability_damage_scaling"`
		CatacombsRequirements struct {
			Dungeon struct {
				Level float64 `json:"level"`
				Type  string  `json:"type"`
			} `json:"dungeon"`
		} `json:"catacombs_requirements"`
		Category     string  `json:"category"`
		Color        string  `json:"color"`
		Crystal      string  `json:"crystal"`
		Description  string  `json:"description"`
		DungeonItem  bool    `json:"dungeon_item"`
		Durability   float64 `json:"durability"`
		Enchantments struct {
			AquaAffinity  float64 `json:"aqua_affinity"`
			BigBrain      float64 `json:"big_brain"`
			CounterStrike float64 `json:"counter_strike"`
			DepthStrider  float64 `json:"depth_strider"`
			Efficiency    float64 `json:"efficiency"`
			FirstStrike   float64 `json:"first_strike"`
			Power         float64 `json:"power"`
			Respiration   float64 `json:"respiration"`
			Scavenger     float64 `json:"scavenger"`
			Sharpness     float64 `json:"sharpness"`
			Telekinesis   float64 `json:"telekinesis"`
			Vampirism     float64 `json:"vampirism"`
			Vicious       float64 `json:"vicious"`
		} `json:"enchantments"`
		Essence struct {
			Costs       []float64 `json:"costs"`
			EssenceType string    `json:"essence_type"`
		} `json:"essence"`
		Furniture      string  `json:"furniture"`
		GearScore      float64 `json:"gear_score"`
		Generator      string  `json:"generator"`
		GeneratorTier  float64 `json:"generator_tier"`
		Glowing        bool    `json:"glowing"`
		ID             string  `json:"id"`
		ItemDurability float64 `json:"item_durability"`
		Material       string  `json:"material"`
		Name           string  `json:"name"`
		NpcSellPrice   float64 `json:"npc_sell_price"`
		PrivateIsland  string  `json:"private_island"`
		Requirements   struct {
			Collection struct {
				CollectionID string  `json:"collection_id"`
				RequiredTier float64 `json:"required_tier"`
			} `json:"collection"`
			Dungeon struct {
				Level float64 `json:"level"`
				Type  string  `json:"type"`
			} `json:"dungeon"`
			DungeonCompletion struct {
				Tier float64 `json:"tier"`
				Type string  `json:"type"`
			} `json:"dungeon_completion"`
			HeartOfTheMountain struct {
				Tier float64 `json:"tier"`
			} `json:"heart_of_the_mountain"`
			Skill struct {
				Level float64 `json:"level"`
				Type  string  `json:"type"`
			} `json:"skill"`
			Slayer struct {
				Level          float64 `json:"level"`
				SlayerBossType string  `json:"slayer_boss_type"`
			} `json:"slayer"`
			TargetPractice struct {
				Mode string `json:"mode"`
			} `json:"targetPractice"`
		} `json:"requirements"`
		Skin  string `json:"skin"`
		Stats struct {
			AbilityDamagePercent float64 `json:"ABILITY_DAMAGE_PERCENT"`
			AttackSpeed          float64 `json:"ATTACK_SPEED"`
			BreakingPower        float64 `json:"BREAKING_POWER"`
			CriticalChance       float64 `json:"CRITICAL_CHANCE"`
			CriticalDamage       float64 `json:"CRITICAL_DAMAGE"`
			Damage               float64 `json:"DAMAGE"`
			Defense              float64 `json:"DEFENSE"`
			Ferocity             float64 `json:"FEROCITY"`
			Health               float64 `json:"HEALTH"`
			Intelligence         float64 `json:"INTELLIGENCE"`
			MagicFind            float64 `json:"MAGIC_FIND"`
			MiningFortune        float64 `json:"MINING_FORTUNE"`
			MiningSpeed          float64 `json:"MINING_SPEED"`
			SeaCreatureChance    float64 `json:"SEA_CREATURE_CHANCE"`
			Strength             float64 `json:"STRENGTH"`
			TrueDefense          float64 `json:"TRUE_DEFENSE"`
			WalkSpeed            float64 `json:"WALK_SPEED"`
			WeaponAbilityDamage  float64 `json:"WEAPON_ABILITY_DAMAGE"`
		} `json:"stats"`
		Tier        string `json:"tier"`
		TieredStats struct {
			AttackSpeed         []float64 `json:"ATTACK_SPEED"`
			CriticalChance      []float64 `json:"CRITICAL_CHANCE"`
			CriticalDamage      []float64 `json:"CRITICAL_DAMAGE"`
			Damage              []float64 `json:"DAMAGE"`
			Defense             []float64 `json:"DEFENSE"`
			Health              []float64 `json:"HEALTH"`
			Intelligence        []float64 `json:"INTELLIGENCE"`
			Strength            []float64 `json:"STRENGTH"`
			WalkSpeed           []float64 `json:"WALK_SPEED"`
			WeaponAbilityDamage []float64 `json:"WEAPON_ABILITY_DAMAGE"`
		} `json:"tiered_stats"`
		Unstackable bool `json:"unstackable"`
	} `json:"items"`
	LastUpdated float64 `json:"lastUpdated"`
	Success     bool    `json:"success"`
}
