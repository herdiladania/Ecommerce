package data

import (
	"e-commerce/features/cart"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID         uint
	CartHasProduct []CartHasProduct `gorm:"foreignKey:CartID"`
}

type CartHasProduct struct {
	CartID     uint
	ProductID  uint
	Quantity   int
	TotalPrice int
}

func CartToCore(data Cart) cart.Core {
	return cart.Core{
		ID:     data.ID,
		UserID: data.UserID,
	}
}

func CoreToCart(data cart.Core) Cart {
	return Cart{
		Model:  gorm.Model{ID: data.ID},
		UserID: data.UserID,
	}
}

func CHPToCore(data CartHasProduct) cart.Core {
	return cart.Core{
		ID:         data.CartID,
		ProductID:  data.ProductID,
		Quantity:   data.Quantity,
		TotalPrice: data.TotalPrice,
	}
}

func CoreToCHP(data cart.Core) CartHasProduct {
	return CartHasProduct{
		CartID:     data.ID,
		ProductID:  data.ProductID,
		Quantity:   data.Quantity,
		TotalPrice: data.TotalPrice,
	}
}
