package mocks

import (
	"github.com/Flyewzz/golang-itv/models"
)

type MockStoreController struct {
	last *models.Request
}

func NewMockStoreController() *MockStoreController {
	return &MockStoreController{}
}

func (st *MockStoreController) Add(request *models.Request) int { st.last = request; return 0 }
func (st *MockStoreController) GetAll() []models.Request        { return []models.Request{} }
func (st *MockStoreController) GetByPage(page int) ([]models.Request, error) {
	return []models.Request{}, nil
}
func (st *MockStoreController) GetById(id int) (*models.Request, error) {
	if id == 0 {
		return st.last, nil
	}
	return nil, STORE_INCORRECT_INDEX
}
func (st *MockStoreController) RemoveById(id int) error { return nil }
func (st *MockStoreController) RemoveAll()              {}
