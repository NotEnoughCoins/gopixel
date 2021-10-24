package structs

type SkyblockAuctions struct {
	Success  bool `json:"success"`
	Auctions []struct {
		ID          string   `json:"_id"`
		UUID        string   `json:"uuid"`
		Auctioneer  string   `json:"auctioneer"`
		ProfileID   string   `json:"profile_id"`
		Coop        []string `json:"coop"`
		Start       int64    `json:"start"`
		End         int64    `json:"end"`
		ItemName    string   `json:"item_name"`
		ItemLore    string   `json:"item_lore"`
		Extra       string   `json:"extra"`
		Category    string   `json:"category"`
		Tier        string   `json:"tier"`
		StartingBid int      `json:"starting_bid"`
		ItemBytes   struct {
			Type int    `json:"type"`
			Data string `json:"data"`
		} `json:"item_bytes"`
		Claimed          bool          `json:"claimed"`
		ClaimedBidders   []interface{} `json:"claimed_bidders"`
		Bin              bool          `json:"bin"`
		HighestBidAmount int           `json:"highest_bid_amount"`
		Bids             []interface{} `json:"bids"`
	} `json:"auctions"`
}
