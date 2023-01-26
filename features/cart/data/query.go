package data

import (
	"e-commerce/features/cart"
	"e-commerce/features/product/data"
	"errors"
	"log"

	"gorm.io/gorm"
)

type cartQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.CartData {
	return &cartQuery{
		db: db,
	}
}

func (cq *cartQuery) AddCartData(newCart cart.Core) (cart.Core, error) {
	cnv := CoreToCart(newCart)
	err := cq.db.Create(&cnv).Error
	if err != nil {
		log.Println("Create query error", err)
		return cart.Core{}, err
	}
	return CartToCore(cnv), nil
}

func (cq *cartQuery) AddCHPData(cartID uint, newCart cart.Core) (cart.Core, error) {
	cnv := CoreToCHP(newCart)
	var compare data.Product
	if err := cq.db.Select("price").Where("id = ?", cnv.ProductID).First(&compare).Error; err != nil {
		log.Println("Select query error", err)
		return cart.Core{}, err
	}

	cnv.CartID = cartID
	cnv.TotalPrice = cnv.Quantity * int(compare.Price)

	err := cq.db.Create(&cnv).Error
	if err != nil {
		log.Println("Create query error", err)
		return cart.Core{}, err
	}
	return CHPToCore(cnv), nil
}

func (cq *cartQuery) MyCartData(userID uint) ([]cart.Core, error) {
	res := []cart.Core{}
	// qry := cq.db.Where("user_id = ?", userID).Find(&res)
	qry := cq.db.Raw("SELECT c.id, c.user_id, ch.quantity, ch.total_price, ch.product_id, p.name, p.image FROM carts c, cart_has_products ch, products p WHERE c.user_id = ? AND c.id = ch.cart_id AND ch.product_id = p.id AND c.deleted_at IS NULL;", userID).Find(&res)
	if qry.Error != nil {
		return nil, qry.Error
	}
	return res, nil
}

func (cq *cartQuery) UpdateCartData(userID, cartID uint, newCart cart.Core) (cart.Core, error) {
	cnv := CoreToCart(newCart)
	qry := cq.db.Where("id = ? AND user_id = ?", cartID, userID).First(&cnv)
	if qry.Error != nil {
		return cart.Core{}, qry.Error
	}
	if qry.RowsAffected < 1 {
		return cart.Core{}, errors.New("cart not found")
	}

	res := CoreToCHP(newCart)
	res.CartID = cnv.ID
	if err := cq.db.Select("product_id").Where("cart_id = ?", res.CartID).First(&res).Error; err != nil {
		log.Println("Select query error", err)
		return cart.Core{}, err
	}

	var compare data.Product
	if err := cq.db.Select("price").Where("id = ?", res.ProductID).First(&compare).Error; err != nil {
		log.Println("Select query error", err)
		return cart.Core{}, err
	}

	res.TotalPrice = res.Quantity * int(compare.Price)

	qryUp := cq.db.Model(&CartHasProduct{}).Where("cart_id = ?", res.CartID).Updates(&res)
	if qryUp.Error != nil {
		return cart.Core{}, qry.Error
	}
	return CHPToCore(res), nil
}

func (cq *cartQuery) DeleteCartData(userID, cartID uint) error {
	qryDel := cq.db.Where("id = ? AND user_id = ?", cartID, userID).Delete(&Cart{})
	if qryDel.Error != nil {
		return qryDel.Error
	}
	if qryDel.RowsAffected < 1 {
		return errors.New("no row deleted")
	}
	return nil
}
