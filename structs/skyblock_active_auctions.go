package structs

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"github.com/Tnze/go-mc/nbt"
)

type ParsedNBT struct {
	Data map[string]interface{}
}

func (p *ParsedNBT) UnmarshalJSON(data []byte) error {
	var y string
	err := json.Unmarshal(data, &y)
	if err != nil {
		return err
	}
	var d map[string]interface{}
	n, err := base64.StdEncoding.DecodeString(y)
	r, err := gzip.NewReader(bytes.NewReader(n))
	_, err = nbt.NewDecoder(r).Decode(&d)
	if err != nil {
		return err
	}
	p.Data = d
	return nil
}

type SkyblockActiveAuctions struct {
	Success       bool   `json:"success"`
	Cause         string `json:"cause"`
	TotalPages    int    `json:"totalPages"`
	TotalAuctions int    `json:"totalAuctions"`
	LastUpdated   int64  `json:"lastUpdated"`
	Auctions      []struct {
		UUID             string        `json:"uuid"`
		Auctioneer       string        `json:"auctioneer"`
		ProfileID        string        `json:"profile_id"`
		Coop             []string      `json:"coop"`
		Start            int64         `json:"start"`
		End              int64         `json:"end"`
		ItemName         string        `json:"item_name"`
		ItemLore         string        `json:"item_lore"`
		Extra            string        `json:"extra"`
		Category         string        `json:"category"`
		Tier             string        `json:"tier"`
		StartingBid      int           `json:"starting_bid"`
		ItemBytes        ParsedNBT     `json:"item_bytes"`
		Claimed          bool          `json:"claimed"`
		ClaimedBidders   []interface{} `json:"claimed_bidders"`
		HighestBidAmount int           `json:"highest_bid_amount"`
		Bin              bool          `json:"bin,omitempty"`
		Bids             []interface{} `json:"bids"`
	} `json:"auctions"`
}

type SkyblockActiveAuctionsPage struct {
	Success       bool   `json:"success"`
	Cause         string `json:"cause"`
	Page          int    `json:"page"`
	TotalPages    int    `json:"totalPages"`
	TotalAuctions int    `json:"totalAuctions"`
	LastUpdated   int64  `json:"lastUpdated"`
	Auctions      []struct {
		UUID             string        `json:"uuid"`
		Auctioneer       string        `json:"auctioneer"`
		ProfileID        string        `json:"profile_id"`
		Coop             []string      `json:"coop"`
		Start            int64         `json:"start"`
		End              int64         `json:"end"`
		ItemName         string        `json:"item_name"`
		ItemLore         string        `json:"item_lore"`
		Extra            string        `json:"extra"`
		Category         string        `json:"category"`
		Tier             string        `json:"tier"`
		StartingBid      int           `json:"starting_bid"`
		ItemBytes        ParsedNBT     `json:"item_bytes"`
		Claimed          bool          `json:"claimed"`
		ClaimedBidders   []interface{} `json:"claimed_bidders"`
		HighestBidAmount int           `json:"highest_bid_amount"`
		Bin              bool          `json:"bin,omitempty"`
		Bids             []interface{} `json:"bids"`
	} `json:"auctions"`
}
