package discogs

import (
	"fmt"
)

type Order struct {
	ID                     string    `json:"id"`
	Status                 string    `json:"status"`
	MessagesURL            string    `json:"messages_url"`
	Fee                    Price     `json:"fee"`
	Created                string    `json:"created"`
	Items                  []Listing `json:"items"`
	ShippingAddress        string    `json:"shipping_address"`
	URI                    string    `json:"uri"`
	Shipping               Price     `json:"shipping"`
	Seller                 User      `json:"seller"`
	LastActivity           string    `json:"last_activity"`
	AdditionalInstructions string    `json:"additional_instructions"`
	NextStatus             []string  `json:"next_status"`
	Buyer                  User      `json:"buyer"`
	ResourceURL            string    `json:"resource_url"`
	Total                  Price     `json:"total"`
}

// todo: raise error unless Client is authenticated?
func (c *Client) GetOrder(orderID string) (out *Order, err error) {
	err = c.get(fmt.Sprintf("marketplace/orders/%s", orderID), nil, &out)
	return
}
