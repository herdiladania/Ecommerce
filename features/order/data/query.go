package data

import (
	"e-commerce/config"
	carts "e-commerce/features/cart/data"
	"e-commerce/features/order"

	"fmt"
	"log"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type orderQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) order.OrderData {
	return &orderQuery{
		db: db,
	}
}
func (oq *orderQuery) Add(userID uint, cartID uint, address string) (order.Core, error) {
	oq.db.Begin()

	cart := []carts.CartHasProduct{}
	if err := oq.db.Where("cart_id=?", cartID).Find(&cart).Error; err != nil {
		oq.db.Rollback()
		log.Println("error retrieve user cart: ", err.Error())
		return order.Core{}, err
	}

	var totalPrice float64
	var quantity int
	for _, item := range cart {
		totalPrice += float64(item.TotalPrice)
		quantity += item.Quantity
	}

	orderInput := Order{
		UserID:          userID,
		Status:          "Waiting For Payment",
		TransactionDate: time.Now(),
		TotalPrice:      totalPrice,
		Address:         address,
		Quantity:        quantity,
	}

	product := carts.Cart{}
	oq.db.First(&product, cart[0].ProductID)
	orderInput.SellerID = product.UserID

	orderInput.TransactionCode = "Transaction-" + fmt.Sprint(orderInput.SellerID)
	s := config.MidtransSnapClient()
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderInput.TransactionCode,
			GrossAmt: int64(totalPrice),
		},
	}
	snapResp, _ := s.CreateTransaction(req)
	orderInput.PaymentUrl = snapResp.RedirectURL
	if err := oq.db.Create(&orderInput).Error; err != nil {
		oq.db.Rollback()
		log.Println("error add order query: ", err.Error())
		return order.Core{}, err
	}

	if err := oq.db.Where("cart_id", cartID).Delete(carts.CartHasProduct{}, carts.Cart{}).Error; err != nil {
		oq.db.Rollback()
		log.Println("error delete  cart: ", err.Error())
		return order.Core{}, err
	}

	oq.db.Save(&orderInput)
	oq.db.Commit()

	return DataToCore(orderInput), nil
}

func (oq *orderQuery) OrderHistory(userID uint) ([]order.Core, error) {
	orderHistory := []Order{}

	err := oq.db.Where("user_id = ?", userID).Find(&orderHistory).Error
	if err != nil {
		log.Println("order history query error", err.Error())
		return []order.Core{}, err
	}

	return ListOrderToCore(orderHistory), nil

}
