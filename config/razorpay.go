package config

import (
	"fmt"
	razorpay "github.com/razorpay/razorpay-go"
)

type RazorpayClient struct {
	client *razorpay.Client
}

func NewRazorpayClient(cfg Config) *RazorpayClient {
	client := razorpay.NewClient(cfg.APIKey, cfg.APISecret)

	return &RazorpayClient{
		client: client,
	}
}

// type pageVariables struct {
// 	OrderID string
// }

func (r *RazorpayClient) CreateOrder(amount float64) (string, error) {
	amountInPaisa := int(amount * 100)

	data := map[string]interface{}{
		"amount":   amountInPaisa,
		"currency": "INR",
		"receipt":  "go_trade",
	}
	body, err := r.client.Order.Create(data, nil)
	if err != nil {
		return "", fmt.Errorf("problem creating order: %v", err)
	}

	value := body["id"]
	orderID := value.(string)

	return orderID, nil
}
