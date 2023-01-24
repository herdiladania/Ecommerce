package data

import (
	"e-commerce/features/user"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Login(email string) (user.Core, error) {
	userLogin := User{}

	if err := uq.db.Where("email = ?", email).First(&userLogin).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}

	return ToCore(userLogin), nil
}
func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {

	cnv := CoreToData(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		return user.Core{}, err
	}

	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) Profile(id uint) (user.Core, error) {
	res := User{}
	if err := uq.db.Where("id = ?", id).First(&res).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return user.Core{}, err
	}

	return ToCore(res), nil
}

func (uq *userQuery) Update(id uint, updatedData user.Core) (user.Core, error) {
	cnvUpd := CoreToData(updatedData)
	qry := uq.db.Model(&User{}).Where("id = ?", id).Updates(cnvUpd)
	if err := qry.Error; err != nil {
		log.Println("error update user query : ", err)
		return updatedData, err
	}
	return updatedData, nil
}

func (uq *userQuery) Delete(id uint) (user.Core, error) {
	users := User{}

	delete := uq.db.Delete(&users, id)

	if delete.Error != nil {
		log.Println("Get By ID query error", delete.Error.Error())
		return user.Core{}, delete.Error
	}

	if delete.RowsAffected < 0 {
		log.Println("Rows affected delete error")
		return user.Core{}, errors.New("user not found")
	}

	return ToCore(users), nil
}
