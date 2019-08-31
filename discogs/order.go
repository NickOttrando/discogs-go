package discogs

import (
	"fmt"
)

type OrderMessage struct {
	Subject   string `json:"subject"`
	Message   string `json:"message"`
	To        User   `json:"to"`
	Order     Order  `json:"order"`
	Timestamp string `json:"timestamp"`
}

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

type OrdersResponse struct {
	Orders []Order `json:"orders"`
	Pagination
}

type OrderMessagesResponse struct {
	Messages []OrderMessage `json:"messages"`
	Pagination
}

// View the data associated with an order.
// Authentication as the seller is required.
func (c *Client) GetOrder(orderID string) (out *Order, err error) {
	err = c.get(fmt.Sprintf("marketplace/orders/%s", orderID), nil, &out)
	return
}

// Returns a list of the authenticated user’s orders.
// Accepts Pagination parameters.
func (c *Client) GetOrders(opts *ListOptions) (out *OrdersResponse, err error) {
	var fmtedOpts *ListOptionsFmted
	if opts != nil {
		fmtedOpts, err = opts.Format()
	}
	if err != nil {
		return nil, err
	}
	err = c.get("marketplace/orders", &fmtedOpts, &out)
	return
}

// Returns a list of the order’s messages with the most recent first.
// Accepts Pagination parameters.
// Authentication as the seller is required.
func (c *Client) GetOrderMessages(orderID string, opts *ListOptions) (out *OrderMessagesResponse, err error) {
	var fmtedOpts *ListOptionsFmted
	if opts != nil {
		fmtedOpts, err = opts.Format()
	}
	if err != nil {
		return nil, err
	}
	err = c.get(fmt.Sprintf("marketplace/orders/%s/messages", orderID), &fmtedOpts, &out)
	return
}
