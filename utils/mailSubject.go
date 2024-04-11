package utils

import (
	"fmt"

	"github.com/anjush-bhargavan/go_trade_product/pkg/model"
)

func NotifyWinner(product model.Product, userName string, amount float64) string {
	message := fmt.Sprintf("Hello %s,\n\n", userName)
	message += fmt.Sprintf("Congratulations! You won the bid for the product '%s'.\n\n", product.Name)
	message += fmt.Sprintf("Product Description: %s\n", product.Description)
	message += fmt.Sprintf("Seller ID: %d\n", product.SellerID)
	message += fmt.Sprintf("Product ID: %d\n", product.ID)
	message += fmt.Sprintf("Category ID: %d\n\n", product.Category)
	message += fmt.Sprintf("Please proceed to payment with the amount of %.2f.\n", amount)
	message += "Payment Link: [Insert Payment Link Here]\n\n"
	message += "Thank you for using our platform!\n\nBest regards,\n[Your Company Name]"

	return message
}


func ProductNotification(p model.Product, category string) string {
	startTime := p.BiddingStartTime.Format("2006-01-02 15:04:05")
	endTime := p.BiddingEndTime.Format("2006-01-02 15:04:05")

	biddingType := ""
	if p.BiddingType {
		biddingType = "Open" 
	}else{
		biddingType = "Closed"
	}

	message := fmt.Sprintf("Hello!\n\nA new product named '%s' has been added to the category '%s', which you are watching.\n\n"+
		"You might be interested in checking it out!\n\n"+
		"Product Details:\n"+
		"- Name: %s\n"+
		"- Description: %s\n"+
		"- Base Price: %d\n"+
		"- Bidding Type: %s\n"+
		"- Bidding Start Time: %s\n"+
		"- Bidding End Time: %s\n\n"+
		"Thank you for using our service!\n",
		p.Name, category,
		p.Name, p.Description, p.BasePrice, biddingType,
		startTime, endTime)

	return message
}

