package data

import (
	"e-commerce/config"
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
func (oq *orderQuery) Add(userID uint) (order.Core, error) {
	transaksi := oq.db.Begin()

	cart := []userCart{}
	if err := transaksi.Where("user_id = ?", userID).Find(&cart).Error; err != nil {
		transaksi.Rollback()
		log.Println("error retrieve user cart: ", err.Error())
		return order.Core{}, err
	}

	var totalPrice float64
	for _, item := range cart {
		totalPrice += item.TotalPrice
	}

	orderInput := Order{
		UserID:          userID,
		Status:          "Waiting For Payment",
		TransactionDate: time.Now(),
		TotalPrice:      totalPrice,
	}

	product := userProduct{}
	transaksi.First(&product, cart[0].ProductID)
	orderInput.SellerID = product.UserID

	if err := transaksi.Create(&orderInput).Error; err != nil {
		transaksi.Rollback()
		log.Println("error add order query: ", err.Error())
		return order.Core{}, err
	}

	orderInput.TransactionCode = "Transaction-" + fmt.Sprint(orderInput.ID)
	transaksi.Save(&orderInput)

	listProducts := []orderProduct{}
	for _, item := range cart {
		orderProduct := orderProduct{
			OrderID:    orderInput.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			TotalPrice: item.TotalPrice,
		}
		listProducts = append(listProducts, orderProduct)
	}
	if err := transaksi.Create(&listProducts).Error; err != nil {
		transaksi.Rollback()
		log.Println("error create orderproduct: ", err.Error())
		return order.Core{}, err
	}

	if err := transaksi.Where("user_id = ?", userID).Delete(userCart{}).Error; err != nil {
		transaksi.Rollback()
		log.Println("error delete user cart: ", err.Error())
		return order.Core{}, err
	}

	s := config.MidtransSnapClient()
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderInput.TransactionCode,
			GrossAmt: int64(totalPrice),
		},
	}
	snapResp, _ := s.CreateTransaction(req)
	orderInput.PaymentUrl = snapResp.RedirectURL
	transaksi.Commit()

	return DataToCore(orderInput), nil
}
