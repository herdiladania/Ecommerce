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

func (os *orderSrv) Add(token interface{}) (order.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		log.Println("user not found")
		return order.Core{}, errors.New("user not found")
	}

	res, err := os.qry.Add(uint(userID))
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
