package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"marketplace-mvc/model"
)

type RatingService struct {
	DB             *gorm.DB
	UserService    *UserService
	ProductService *ProductService
}

func (s *RatingService) Create(rating model.Rating) error {
	userFound, err := s.UserService.GetById(rating.UserID)
	if err != nil {
		return err
	}
	rating.User = userFound

	productFound, err := s.ProductService.GetById(rating.ProductRatedID)
	if err != nil {
		return err
	}
	rating.ProductRated = productFound

	s.DB.Create(&rating)

	if rating.ID == 0 {
		return errors.New("error creating rating")
	}

	return nil
}

func (s *RatingService) GetById(ratingId uint) (model.Rating, error) {
	rating := model.Rating{}
	result := s.DB.First(&rating, ratingId)

	if result.Error != nil {
		return model.Rating{}, errors.New(fmt.Sprintf("error: rating not found with id: %d", ratingId))
	}
	return rating, nil
}

func (s *RatingService) GetAll() ([]model.Rating, error) {
	var ratings []model.Rating
	res := s.DB.Find(&ratings)

	if res.Error != nil {
		return []model.Rating{}, res.Error
	}

	return ratings, nil
}

func (s *RatingService) Update(rating model.Rating) error {
	_, err := s.GetById(rating.ID)

	if err != nil {
		return err
	}

	s.DB.Save(&rating)

	return nil
}

func (s *RatingService) Delete(ratingId uint) error {
	_, err := s.GetById(ratingId)

	if err != nil {
		return err
	}

	res := s.DB.Delete(&model.Rating{}, ratingId)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("error: cannot delete rating with id: %d", ratingId))
	}
	return nil
}
