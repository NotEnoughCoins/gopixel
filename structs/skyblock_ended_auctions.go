package structs

type SkyblockEndedAuctions struct {
	Success bool `json:"success"`

	LastUpdated int64 `json:"lastUpdated"`
	Auctions    []struct {
		AuctionID     string  `json:"auction_id"`
		Seller        string  `json:"seller"`
		SellerProfile string  `json:"seller_profile"`
		Buyer         string  `json:"buyer"`
		Timestamp     int64   `json:"timestamp"`
		Price         float64 `json:"price"`
		Bin           bool    `json:"bin"`
		ItemBytes     string  `json:"item_bytes"`
	} `json:"auctions"`
}
