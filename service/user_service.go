package service

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"marketplace-mvc/model"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) Create(user model.User) error {
	s.DB.Create(&user)

	if user.ID == 0 {
		log.Error().Msg("error creating user")
		return errors.New("error creating user")
	}

	return nil
}

func (s *UserService) GetById(userId uint) (model.User, error) {
	user := model.User{}
	result := s.DB.First(&user, userId)

	if result.Error != nil {
		log.Error().Msg("error getting user")
		return model.User{}, errors.New(fmt.Sprintf("error: user not found with id: %d", userId))
	}
	return user, nil
}

func (s *UserService) GetAll() ([]model.User, error) {
	var users []model.User
	res := s.DB.Find(&users)

	if res.Error != nil {
		return []model.User{}, res.Error
	}

	return users, nil
}

func (s *UserService) Update(user model.User) error {
	_, err := s.GetById(user.ID)

	if err != nil {
		return err
	}

	s.DB.Save(&user)

	return nil
}

func (s *UserService) Delete(userId uint) error {
	_, err := s.GetById(userId)

	if err != nil {
		return err
	}

	res := s.DB.Delete(&model.User{}, userId)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("error: cannot delete user with id: %d", userId))
	}
	return nil
}

func (s *UserService) Login(user model.User) error {
	userFound := model.User{}
	result := s.DB.Where("password = ? and name = ?", user.Password, user.Name).First(&userFound)
	if result.Error != nil || userFound.ID == 0 {
		log.Error().Msg("error log user")
		return errors.New("password or user incorrect")
	}

	return nil
}
