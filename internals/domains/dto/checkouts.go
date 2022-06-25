package dto

type Checkouts struct {
	UserID           uint               `json:"user_id" binding:"required"`
	CartID           uint               `json:"cart_id" binding:"required"`
	CheckoutProducts []CheckoutProducts `json:"checkout_products" binding:"required"`
}

type CheckoutProducts struct {
	CartProductID uint    `json:"cart_product_id" binding:"required"`
	ProductID     uint    `json:"product_id" binding:"required"`
	Quantity      int32   `json:"quantity" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
}
