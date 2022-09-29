package service

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"marketplace-mvc/model"
)

type PaymentService struct {
	DB                  *gorm.DB
	UserService         *UserService
	ShoppingCartService *ShoppingCartService
}

func (s *PaymentService) Create(payment model.Payment) error {
	userFound, err := s.UserService.GetById(payment.PayerID)
	if err != nil {
		return err
	}
	payment.Payer = userFound

	s.DB.Create(&payment)

	if payment.ID == 0 {
		return errors.New("error creating payment")
	}

	return nil
}

func (s *PaymentService) GetById(paymentId uint) (model.Payment, error) {
	payment := model.Payment{}
	result := s.DB.First(&payment, paymentId)

	if result.Error != nil {
		return model.Payment{}, errors.New(fmt.Sprintf("error: payment not found with id: %d", paymentId))
	}
	return payment, nil
}

func (s *PaymentService) GetAll() ([]model.Payment, error) {
	var payments []model.Payment
	res := s.DB.Find(&payments)

	if res.Error != nil {
		return []model.Payment{}, res.Error
	}

	return payments, nil
}

func (s *PaymentService) Update(payment model.Payment) error {
	_, err := s.GetById(payment.ID)

	if err != nil {
		return err
	}

	s.DB.Save(&payment)

	return nil
}

func (s *PaymentService) Delete(paymentId uint) error {
	_, err := s.GetById(paymentId)

	if err != nil {
		return err
	}

	res := s.DB.Delete(&model.Payment{}, paymentId)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("error: cannot delete payment with id: %d", paymentId))
	}
	return nil
}

func (s *PaymentService) ApprovePayment(payment model.Payment) error {
	payment.Status = "approved"
	err := s.Update(payment)
	if err != nil {
		log.Error().Msg("error approving payment")
		return err
	}

	err = s.ShoppingCartService.BuyCart(payment.ShoppingCartID)
	if err != nil {
		log.Error().Msg("error creating address")
		return err
	}

	return nil
}
