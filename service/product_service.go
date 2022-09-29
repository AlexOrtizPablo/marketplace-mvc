package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"marketplace-mvc/model"
)

type ProductService struct {
	DB          *gorm.DB
	UserService *UserService
}

func (s *ProductService) PublishProduct(product model.Product) error {
	userFound, err := s.UserService.GetById(product.SellerID)
	if err != nil {
		return err
	}
	product.Seller = userFound

	s.DB.Create(&product)

	if product.ID == 0 {
		return errors.New("error creating product")
	}

	return nil
}

func (s *ProductService) GetById(productId uint) (model.Product, error) {
	product := model.Product{}
	result := s.DB.First(&product, productId)

	if result.Error != nil {
		return model.Product{}, errors.New(fmt.Sprintf("error: product not found with id: %d", productId))
	}
	return product, nil
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	var products []model.Product
	res := s.DB.Find(&products)

	if res.Error != nil {
		return []model.Product{}, res.Error
	}

	return products, nil
}

func (s *ProductService) Update(product model.Product) error {
	_, err := s.GetById(product.ID)

	if err != nil {
		return err
	}

	s.DB.Save(&product)

	return nil
}

func (s *ProductService) Delete(productId uint) error {
	_, err := s.GetById(productId)

	if err != nil {
		return err
	}

	res := s.DB.Delete(&model.Product{}, productId)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("error: cannot delete product with id: %d", productId))
	}
	return nil
}
