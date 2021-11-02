package gopixel

import (
	"encoding/json"

	"fmt"

	"testing"

	"os"
)

// The config struct with all of the required variables
type Config struct {
	Key                 string `json:"key"`
	Player              string `json:"player"`
	SkyblockProfile     string `json:"skyblock_profile"`
	RankedSkywarsPlayer string `json:"ranked_skywars_player"`
	GuildId             string `json:"guild_id"`
	GuildName           string `json:"guild_name"`
}

// Get's the json data from the config.json file and unmarshals it to the Config struct
func getConfig() Config {
	var config Config

	bytes, err := os.ReadFile("config.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic(err)
	}

	return config
}

var config Config = getConfig()
var client *Client = NewClient(config.Key)

func TestKeyData(t *testing.T) {
	key, err := client.KeyData()
	if err != nil {
		t.Error((err))
	}
	fmt.Println(key)
}

func TestPlayerData(t *testing.T) {
	player, err := client.PlayerData(config.Player)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(player)
}

func TestAchievements(t *testing.T) {
	achievements, err := client.Achievements()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(achievements)
}

func TestBazaar(t *testing.T) {
	bazaar, err := client.Bazaar()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(bazaar)
}

func TestFriends(t *testing.T) {
	friends, err := client.Friends(config.Player)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(friends)
}

func TestQuests(t *testing.T) {
	quests, err := client.Quests()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(quests)
}

func TestChallenges(t *testing.T) {
	challenges, err := client.Challenges()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(challenges)
}

func TestGames(t *testing.T) {
	games, err := client.Games()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(games)
}

func TestGuildByPlayer(t *testing.T) {
	guild, err := client.GuildByPlayer(config.Player)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(guild)
}

func TestGuildById(t *testing.T) {
	guild, err := client.GuildById(config.GuildId)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(guild)
}

func TestGuildByName(t *testing.T) {
	guild, err := client.GuildByName(config.GuildName)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(guild)
}

func TestGuildAchievements(t *testing.T) {
	guildAchievements, err := client.GuildAchievements()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(guildAchievements)
}

func TestGuilPermissions(t *testing.T) {
	guildPermissions, err := client.GuildAchievements()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(guildPermissions)
}

func TestSkyblockCollections(t *testing.T) {
	skyblockCollections, err := client.SkyblockCollections()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(skyblockCollections)
}

func TestSkyblockProfiles(t *testing.T) {
	profiles, err := client.SkyblockProfiles(config.Player)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(profiles)
}

func TestSkyblockProfile(t *testing.T) {
	profiles, err := client.SkyblockProfile(config.SkyblockProfile)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(profiles)
}

func TestSkyblockItems(t *testing.T) {
	items, err := client.SkyblockItems()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(items)
}

func TestRecentGames(t *testing.T) {
	games, err := client.RecentGames(config.Player)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(games)
}

func TestPlayerStatus(t *testing.T) {
	status, err := client.PlayerStatus(config.Player)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(status)
}

func TestSkyblockSkills(t *testing.T) {
	skills, err := client.SkyblockSkills()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(skills)
}

func TestRankedSkywars(t *testing.T) {
	data, err := client.RankedSkywars(config.RankedSkywarsPlayer)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func TestSkyblockNews(t *testing.T) {
	news, err := client.SkyblockNews()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(news)
}

func TestSkyblockActiveAuctions(t *testing.T) {
	auctions, err := client.SkyblockActiveAuctions()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(auctions)
}

func TestSkyblockEndedAuctions(t *testing.T) {
	auctions, err := client.SkyblockActiveAuctions()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(auctions)
}

func TestBoosters(t *testing.T) {
	boosters, err := client.Boosters()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(boosters)
}

func TestPlayerCounts(t *testing.T) {
	counts, err := client.PlayerCounts()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(counts)
}

func TestLeaderboards(t *testing.T) {
	leaderboards, err := client.Leaderboards()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(leaderboards)
}

func TestPunishmentStats(t *testing.T) {
	stats, err := client.PunishmentStats()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(stats)
}
