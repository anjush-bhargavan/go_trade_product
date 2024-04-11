package repo


import "github.com/anjush-bhargavan/go_trade_product/pkg/model"

// CreateBid will create a Bid in the database.
func (u *ProductRepository) CreateBidTable(bid *model.Bidding) error {
	if err := u.DB.Create(&bid).Error; err != nil {
		return err
	}
	return nil
}

// FetchBids will fetch  Bids  by the product ID in the database.
func (u *ProductRepository) FindBidsOfProduct(productID uint) (*model.Bidding, error) {
	var bids *model.Bidding
	if err := u.DB.Where("product_id = ?", productID).Find(&bids).Error; err != nil {
		return nil, err
	}
	return bids, nil
}

// UpdateBid will update a Bid in the database.
func (u *ProductRepository) UpdateBidTable(bid *model.Bidding) error {
	if err := u.DB.Save(&bid).Error; err != nil {
		return err
	}
	return nil
}

// // FindBid will fetch  Bid  by the product ID & userID in the database.
// func (u *ProductRepository) FindBid(userID, productID uint) (*model.Bidding, error) {
// 	var bid *model.Bidding
// 	if err := u.DB.Where("product_id = ? AND user_id = ?", productID, userID).First(&bid).Error; err != nil {
// 		return nil, err
// 	}
// 	return bid, nil
// }

// // FindBidByAmount will fetch  Bid  by the product ID & amount in the database.
// func (u *ProductRepository) FindBidByAmount(productID uint,amount float64) (*model.Bidding, error) {
// 	var bid *model.Bidding
// 	if err := u.DB.Where("product_id = ? AND amount = ?", productID, amount).First(&bid).Error; err != nil {
// 		return nil, err
// 	}
// 	return bid, nil
// }
