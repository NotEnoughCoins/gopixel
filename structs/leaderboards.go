package structs

type Leaderboards struct {
	Success bool `json:"success"`
	Data    struct {
		UHC []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"UHC"`
		VAMPIREZ []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"VAMPIREZ"`
		TRUECOMBAT []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"TRUE_COMBAT"`
		SURVIVALGAMES []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"SURVIVAL_GAMES"`
		WALLS3 []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"WALLS3"`
		QUAKECRAFT []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"QUAKECRAFT"`
		TNTGAMES []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"TNTGAMES"`
		SUPERSMASH []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"SUPER_SMASH"`
		WALLS []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"WALLS"`
		MURDERMYSTERY []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"MURDER_MYSTERY"`
		BUILDBATTLE []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"BUILD_BATTLE"`
		ARCADE []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"ARCADE"`
		GINGERBREAD []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"GINGERBREAD"`
		SKYCLASH []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"SKYCLASH"`
		PROTOTYPE []interface{} `json:"PROTOTYPE"`
		BEDWARS   []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"BEDWARS"`
		DUELS []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"DUELS"`
		ARENA []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"ARENA"`
		MCGO []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"MCGO"`
		SPEEDUHC []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"SPEED_UHC"`
		PAINTBALL []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"PAINTBALL"`
		BATTLEGROUND []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"BATTLEGROUND"`
		SKYWARS []struct {
			Path     string   `json:"path"`
			Prefix   string   `json:"prefix"`
			Title    string   `json:"title"`
			Location string   `json:"location"`
			Count    int      `json:"count"`
			Leaders  []string `json:"leaders"`
		} `json:"SKYWARS"`
	} `json:"leaderboards"`
}
