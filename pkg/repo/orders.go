package repo

import "github.com/anjush-bhargavan/go_trade_product/pkg/model"

// CreateOrders method for creating a new Orders in db
func (p *ProductRepository) CreateOrders(Orders *model.Orders) error {
	if err := p.DB.Create(&Orders).Error; err != nil {
		return err
	}
	return nil
}

// FindOrdersByID method finds the Orders from database using primary key
func (p *ProductRepository) FindOrdersByID(OrdersID uint) (*model.Orders, error) {
	var Orders model.Orders

	if err := p.DB.First(&Orders, OrdersID).Error; err != nil {
		return nil, err
	}
	return &Orders, nil
}

// UpdateOrders method update the details of Orders in database.
func (p *ProductRepository) UpdateOrders(Orders *model.Orders) error {
	if err := p.DB.Save(&Orders).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOrders method delete the Orders from database using Orders ID.
func (p *ProductRepository) DeleteOrders(OrdersID uint)  error {

	if err := p.DB.Delete(&model.Orders{},OrdersID).Error; err != nil {
		return  err
	}
	return nil
}

// FindAllOrders method finds all orders from database.
func (p *ProductRepository) FindAllOrders() (*[]model.Orders, error) {
	var orders []model.Orders

	if err := p.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}

// FindOrdersByUser method finds all orders of the user from database.
func (p *ProductRepository) FindOrdersByUser(userID uint) (*[]model.Orders, error) {
	var orders []model.Orders

	if err := p.DB.Where("user_id = ?",userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}

// FindOrdersBySeller method finds all orders of the seller from database.
func (p *ProductRepository) FindOrdersBySeller(sellerID uint) (*[]model.Orders, error) {
	var orders []model.Orders

	if err := p.DB.Where("user_id = ?",sellerID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}

// FindOrder method finds order of the user of specific product from database.
func (p *ProductRepository) FindOrder(userID,productID uint) (*model.Orders, error) {
	var order model.Orders

	if err := p.DB.Where("user_id = ? AND product_id = ?",userID,productID).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}