package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"marketplace-mvc/model"
)

type PurchaseService struct {
	DB             *gorm.DB
	UserService    *UserService
	ProductService *ProductService
}

func (s *PurchaseService) Create(purchase model.Purchase) error {
	userFound, err := s.UserService.GetById(purchase.BuyerID)
	if err != nil {
		return err
	}
	purchase.Buyer = userFound

	productFound, err := s.ProductService.GetById(purchase.ProductID)
	if err != nil {
		return err
	}
	purchase.Product = productFound

	s.DB.Create(&purchase)

	if purchase.ID == 0 {
		return errors.New("error creating purchase")
	}

	return nil
}

func (s *PurchaseService) GetById(purchaseId uint) (model.Purchase, error) {
	purchase := model.Purchase{}
	result := s.DB.First(&purchase, purchaseId)

	if result.Error != nil {
		return model.Purchase{}, errors.New(fmt.Sprintf("error: purchase not found with id: %d", purchaseId))
	}
	return purchase, nil
}

func (s *PurchaseService) GetAll() ([]model.Purchase, error) {
	var purchases []model.Purchase
	res := s.DB.Find(&purchases)

	if res.Error != nil {
		return []model.Purchase{}, res.Error
	}

	return purchases, nil
}

func (s *PurchaseService) Update(purchase model.Purchase) error {
	_, err := s.GetById(purchase.ID)

	if err != nil {
		return err
	}

	s.DB.Save(&purchase)

	return nil
}

func (s *PurchaseService) Delete(purchaseId uint) error {
	_, err := s.GetById(purchaseId)

	if err != nil {
		return err
	}

	res := s.DB.Delete(&model.Purchase{}, purchaseId)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("error: cannot delete purchase with id: %d", purchaseId))
	}
	return nil
}
