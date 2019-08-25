package discogs

import (
	"fmt"
)

type Order struct {
	ID int64 `json:"id"`
}

// todo: raise error unless Client is authenticated?
func (c *Client) GetOrder(orderID string) (out *Order, err error) {
	err = c.get(fmt.Sprintf("marketplace/orders/%s", orderID), nil, &out)
	return
}
