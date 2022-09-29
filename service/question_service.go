package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"marketplace-mvc/model"
)

type QuestionService struct {
	DB             *gorm.DB
	UserService    *UserService
	ProductService *ProductService
}

func (s *QuestionService) Create(question model.Question) error {
	userFound, err := s.UserService.GetById(question.UserID)
	if err != nil {
		return err
	}
	question.User = userFound

	productFound, err := s.ProductService.GetById(question.ProductID)
	if err != nil {
		return err
	}
	question.Product = productFound

	s.DB.Create(&question)

	if question.ID == 0 {
		return errors.New("error creating question")
	}

	return nil
}

func (s *QuestionService) GetById(questionId uint) (model.Question, error) {
	question := model.Question{}
	result := s.DB.First(&question, questionId)

	if result.Error != nil {
		return model.Question{}, errors.New(fmt.Sprintf("error: question not found with id: %d", questionId))
	}
	return question, nil
}

func (s *QuestionService) GetAll() ([]model.Question, error) {
	var questions []model.Question
	res := s.DB.Find(&questions)

	if res.Error != nil {
		return []model.Question{}, res.Error
	}

	return questions, nil
}

func (s *QuestionService) Update(question model.Question) error {
	_, err := s.GetById(question.ID)

	if err != nil {
		return err
	}

	s.DB.Save(&question)

	return nil
}

func (s *QuestionService) Delete(questionId uint) error {
	_, err := s.GetById(questionId)

	if err != nil {
		return err
	}

	res := s.DB.Delete(&model.Question{}, questionId)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("error: cannot delete question with id: %d", questionId))
	}
	return nil
}
