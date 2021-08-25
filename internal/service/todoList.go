package service

import (
	"github.com/hostingvk4/badgerList/internal/models"
	"github.com/hostingvk4/badgerList/internal/repository"
)

type ListService struct {
	repo repository.List
}

func NewListService(repo repository.List) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) Create(list models.List) (int, error) {
	return s.repo.Create(list)
}

func (s *ListService) GetAll(userId int) ([]models.List, error) {
	return s.repo.GetAll(userId)
}
func (s *ListService) GetListById(userId, listId int) (models.List, error) {
	return s.repo.GetListById(userId, listId)
}
func (s *ListService) Update(userId, listId int, list models.UpdateListInput) error {
	return s.repo.Update(userId, listId, list)
}
func (s *ListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
