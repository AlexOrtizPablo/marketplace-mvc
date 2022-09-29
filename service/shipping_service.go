package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"marketplace-mvc/model"
)

type ShippingService struct {
	DB          *gorm.DB
	UserService *UserService
}

func (s *ShippingService) Create(shipping model.Shipping) error {
	userFound, err := s.UserService.GetById(shipping.UserID)
	if err != nil {
		return err
	}
	shipping.User = userFound

	shipping.Status = "pending"
	s.DB.Create(&shipping)

	if shipping.ID == 0 {
		return errors.New("error creating shipping")
	}

	return nil
}

func (s *ShippingService) GetById(shippingId uint) (model.Shipping, error) {
	shipping := model.Shipping{}
	result := s.DB.First(&shipping, shippingId)

	if result.Error != nil {
		return model.Shipping{}, errors.New(fmt.Sprintf("error: shipping not found with id: %d", shippingId))
	}
	return shipping, nil
}

func (s *ShippingService) GetAll() ([]model.Shipping, error) {
	var shippingList []model.Shipping
	res := s.DB.Find(&shippingList)

	if res.Error != nil {
		return []model.Shipping{}, res.Error
	}

	return shippingList, nil
}

func (s *ShippingService) Update(shipping model.Shipping) error {
	_, err := s.GetById(shipping.ID)
	if err != nil {
		return err
	}

	s.DB.Save(&shipping)

	return nil
}

func (s *ShippingService) Delete(shippingId uint) error {
	_, err := s.GetById(shippingId)

	if err != nil {
		return err
	}

	res := s.DB.Delete(&model.Shipping{}, shippingId)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("error: cannot delete shipping with id: %d", shippingId))
	}
	return nil
}
