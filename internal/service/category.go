package service

import (
	"shop/internal/model"
	"shop/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) Create(category *model.Category) error {
	return nil
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return nil, nil
}

func (s *CategoryService) GetByID(id int64) (*model.Category, error) {
	return nil, nil
}

func (s *CategoryService) Update(category *model.Category) error {
	return nil
}

func (s *CategoryService) ChangeStatus(id int64, active bool) error {
	return nil
}

func (s *CategoryService) Delete(id int64) error {
	return nil
}
