package structs

type SkyblockSkill struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MaxLevel    float64 `json:"maxLevel"`
	Levels      []struct {
		Level            float64  `json:"level"`
		TotalExpRequired float64  `json:"totalExpRequired"`
		Unlocks          []string `json:"unlocks"`
	} `json:"levels"`
}

type SkyblockSkills struct {
	Success     bool    `json:"success"`
	LastUpdated float64 `json:"lastUpdated"`
	Version     string  `json:"version"`
	Collections struct {
		FARMING      SkyblockSkill `json:"FARMING"`
		MINING       SkyblockSkill `json:"MINING"`
		COMBAT       SkyblockSkill `json:"COMBAT"`
		FORAGING     SkyblockSkill `json:"FORAGING"`
		FISHING      SkyblockSkill `json:"FISHING"`
		ENCHANTING   SkyblockSkill `json:"ENCHANTING"`
		ALCHEMY      SkyblockSkill `json:"ALCHEMY"`
		CARPENTRY    SkyblockSkill `json:"CARPENTRY"`
		RUNECRAFTING SkyblockSkill `json:"RUNECRAFTING"`
		SOCIAL       SkyblockSkill `json:"SOCIAL"`
		TAMING       SkyblockSkill `json:"TAMING"`
	} `json:"collections"`
	Skills struct {
		FARMING      SkyblockSkill `json:"FARMING"`
		MINING       SkyblockSkill `json:"MINING"`
		COMBAT       SkyblockSkill `json:"COMBAT"`
		FORAGING     SkyblockSkill `json:"FORAGING"`
		FISHING      SkyblockSkill `json:"FISHING"`
		ENCHANTING   SkyblockSkill `json:"ENCHANTING"`
		ALCHEMY      SkyblockSkill `json:"ALCHEMY"`
		CARPENTRY    SkyblockSkill `json:"CARPENTRY"`
		RUNECRAFTING SkyblockSkill `json:"RUNECRAFTING"`
		SOCIAL       SkyblockSkill `json:"SOCIAL"`
		TAMING       SkyblockSkill `json:"TAMING"`
	} `json:"skills"`
}
