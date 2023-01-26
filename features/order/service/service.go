package service

import (
	"e-commerce/features/order"
	"e-commerce/helper"
	"errors"
	"log"
	"strings"
)

type orderSrv struct {
	qry order.OrderData
}

func New(data order.OrderData) order.OrderService {
	return &orderSrv{
		qry: data,
	}
}

func (os *orderSrv) Add(token interface{}, cartID uint, address string) (order.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		log.Println("user not found")
		return order.Core{}, errors.New("user not found")
	}

	res, err := os.qry.Add(uint(userID), cartID, address)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "order not found"
		} else {
			msg = "server problem"
		}
		return order.Core{}, errors.New(msg)
	}

	return res, nil
}

func (os *orderSrv) OrderHistory(token interface{}) ([]order.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		log.Println("error extract token")
		return []order.Core{}, errors.New("user not found")
	}

	res, err := os.qry.OrderHistory(uint(userID))

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error order history query: ", err.Error())
		return []order.Core{}, errors.New(msg)
	}
	return res, nil
}

func (os *orderSrv) UpdateOrderStatus(token interface{}, orderID uint, updatedStatus string) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("user not found")
	}

	err := os.qry.UpdateOrderStatus(uint(userID), orderID, updatedStatus)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "order data not found"
		} else {
			msg = "server problem"
		}
		return errors.New(msg)
	}
	return nil
}
