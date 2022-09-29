package service

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"marketplace-mvc/model"
)

type AddressService struct {
	DB *gorm.DB
}

func (s *AddressService) Create(address model.Address) error {
	s.DB.Create(&address)

	if address.ID == 0 {
		log.Error().Msg("error creating address")
		return errors.New("error creating address")
	}

	return nil
}

func (s *AddressService) GetById(addressId uint) (model.Address, error) {
	address := model.Address{}
	result := s.DB.First(&address, addressId)

	if result.Error != nil {
		return model.Address{}, errors.New(fmt.Sprintf("error: address not found with id: %d", addressId))
	}
	return address, nil
}

func (s *AddressService) GetAll() ([]model.Address, error) {
	var addresses []model.Address
	res := s.DB.Find(&addresses)

	if res.Error != nil {
		log.Error().Msg("error getting addresses")
		return []model.Address{}, res.Error
	}

	return addresses, nil
}

func (s *AddressService) Update(address model.Address) error {
	_, err := s.GetById(address.ID)

	if err != nil {
		log.Error().Msg("error updating address")
		return err
	}

	s.DB.Save(&address)

	return nil
}

func (s *AddressService) Delete(addressId uint) error {
	_, err := s.GetById(addressId)

	if err != nil {
		log.Error().Msg("error deleting address")
		return err
	}

	res := s.DB.Delete(&model.Address{}, addressId)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("error: cannot delete address with id: %d", addressId))
	}
	return nil
}
