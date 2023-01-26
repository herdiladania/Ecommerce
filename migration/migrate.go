package migration

import (
	cart "e-commerce/features/cart/data"
	order "e-commerce/features/order/data"
	product "e-commerce/features/product/data"
	user "e-commerce/features/user/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(user.User{})
	db.AutoMigrate(product.Product{})
	db.AutoMigrate(cart.Cart{})
	db.AutoMigrate(cart.CartHasProduct{})
	db.AutoMigrate(order.Order{})
}
