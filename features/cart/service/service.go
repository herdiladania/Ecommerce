package service

import (
	"e-commerce/features/cart"
	"errors"
	"log"
	"strings"
)

type cartService struct {
	qry cart.CartData
}

func New(cd cart.CartData) cart.CartService {
	return &cartService{
		qry: cd,
	}
}

func (cs *cartService) AddCartSrv(newCart cart.Core) (cart.Core, error) {
	resCart, err := cs.qry.AddCartData(newCart)
	if err != nil {
		log.Println(err)
		return cart.Core{}, errors.New("internal server error")
	}

	resCHP, err := cs.qry.AddCHPData(resCart.ID, newCart)
	if err != nil {
		log.Println(err)
		return cart.Core{}, errors.New("internal server error")
	}

	resCHP.UserID = resCart.UserID
	return resCHP, nil
}

func (cs *cartService) MyCartSrv(id uint) ([]cart.Core, error) {
	res, err := cs.qry.MyCartData(id)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "cart not found"
		} else {
			msg = "internal server error"
		}
		return nil, errors.New(msg)
	}
	if len(res) < 1 {
		return nil, errors.New("cart empty")
	}
	return res, nil
}

func (cs *cartService) UpdateCartSrv(userID, cartID uint, updateCart cart.Core) (cart.Core, error) {
	res, err := cs.qry.UpdateCartData(userID, cartID, updateCart)
	if err != nil {
	log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "cart not found"
		} else {
			msg = "internal server error"
		}
		return cart.Core{}, errors.New(msg)
	}
	return res, nil
}

func (cs *cartService) DeleteCartSrv(userID, cartID uint) (error) {
	if err := cs.qry.DeleteCartData(userID, cartID); err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "cart not found"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}
	return nil
}
