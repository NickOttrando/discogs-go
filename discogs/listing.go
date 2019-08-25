package discogs

import (
	"fmt"
)

type Price struct {
	Currency string  `json:"currency"`
	Method   string  `json:"method"`
	Value    float32 `json:"value"`
}

type Listing struct {
	ID                    int64   `json:"id"`
	Status                string  `json:"status"`
	ShipsFrom             string  `json:"ships_from"`
	AllowOffers           bool    `json:"allow_offers"`
	URI                   string  `json:"uri"`
	Comments              string  `json:"comments"`
	Seller                User    `json:"seller"`
	Price                 Price   `json:"price"`
	OriginalShippingPrice Price   `json:"original_shipping_price"`
	SleeveCondition       string  `json:"sleeve_condition"`
	ShippingPrice         Price   `json:"shipping_price"`
	Release               Release `json:"release"`
	ResourceURL           string  `json:"resource_url"`
	Audio                 bool    `json:'audio"`
	Condition             string  `json:"condition"`
	Posted                string  `json:"posted"`
}

type InventoryResponse struct {
	Listings   []Listing `json:"listings"`
	Pagination `json:"pagination"`
}

func (c *Client) GetListing(listingID int64) (out *Listing, err error) {
	err = c.get(fmt.Sprintf("marketplace/listings/%d", listingID), nil, &out)
	return
}

func (c *Client) GetUserInventory(userName string, opts *ListOptions) (out *InventoryResponse, err error) {
	var fmtedOpts *ListOptionsFmted
	if opts != nil {
		fmtedOpts, err = opts.Format()
	}
	if err != nil {
		return nil, err
	}
	err = c.get(fmt.Sprintf("users/%s/inventory", userName), &fmtedOpts, &out)
	return
}
