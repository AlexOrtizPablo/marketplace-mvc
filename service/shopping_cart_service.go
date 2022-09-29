package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"marketplace-mvc/model"
	"time"
)

type ShoppingCartService struct {
	DB              *gorm.DB
	UserService     *UserService
	ProductService  *ProductService
	PurchaseService *PurchaseService
}

func (s *ShoppingCartService) Create(shoppingCart model.ShoppingCart) error {
	userFound, err := s.UserService.GetById(shoppingCart.BuyerID)
	if err != nil {
		return err
	}
	shoppingCart.Buyer = userFound

	for _, product := range shoppingCart.Products {
		_, err := s.ProductService.GetById(product.ID)
		if err != nil {
			return err
		}
	}

	s.DB.Create(&shoppingCart)

	if shoppingCart.ID == 0 {
		return errors.New("error creating shoppingCart")
	}

	return nil
}

func (s *ShoppingCartService) GetById(shoppingCartId uint) (model.ShoppingCart, error) {
	shoppingCart := model.ShoppingCart{}
	result := s.DB.First(&shoppingCart, shoppingCartId)

	if result.Error != nil {
		return model.ShoppingCart{}, errors.New(fmt.Sprintf("error: shoppingCart not found with id: %d", shoppingCartId))
	}
	return shoppingCart, nil
}

func (s *ShoppingCartService) GetAll() ([]model.ShoppingCart, error) {
	var shoppingCarts []model.ShoppingCart
	res := s.DB.Find(&shoppingCarts)

	if res.Error != nil {
		return []model.ShoppingCart{}, res.Error
	}

	return shoppingCarts, nil
}

func (s *ShoppingCartService) Update(shoppingCart model.ShoppingCart) error {
	_, err := s.GetById(shoppingCart.ID)

	if err != nil {
		return err
	}

	s.DB.Save(&shoppingCart)

	return nil
}

func (s *ShoppingCartService) Delete(shoppingCartId uint) error {
	_, err := s.GetById(shoppingCartId)
	if err != nil {
		return err
	}

	res := s.DB.Delete(&model.ShoppingCart{}, shoppingCartId)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("error: cannot delete shoppingCart with id: %d", shoppingCartId))
	}
	return nil
}

func (s *ShoppingCartService) BuyCart(shoppingCartId uint) error {
	shoppingCart, err := s.GetById(shoppingCartId)
	if err != nil {
		return err
	}

	for _, product := range shoppingCart.Products {
		err := s.PurchaseService.Create(model.Purchase{
			BuyerID:   shoppingCart.BuyerID,
			Buyer:     shoppingCart.Buyer,
			ProductID: product.ID,
			Product:   product,
			CreatedAt: time.Time{},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
