package pricingsdk

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
)

type Pricing struct {
	ROI          float64 `json:"roi"`
	MarketPrice  float64 `json:"marketPrice"`
	BondingPrice float64 `json:"bondingPrice"`
	MaxPayout    float64 `json:"maxPayout"`
	PayoutAmount float64 `json:"payoutAmount"`
}

func (c *Client) FetchPricing(account solana.PublicKey) (Pricing, error) {
	path := fmt.Sprintf("/bonding/%s", account)
	var result Pricing
	err := c.get(path, &result)
	return result, err
}
