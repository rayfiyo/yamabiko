// internal/usecase/mock_usecase.go

package usecase_test

import (
	"github.com/rayfiyo/yamabiko/internal/domain"
	"github.com/stretchr/testify/mock"
)

type mockGeminiClient struct {
	mock.Mock
}

func (m *mockGeminiClient) GenerateResponses(voice string) ([]string, error) {
	args := m.Called(voice)
	return args.Get(0).([]string), args.Error(1)
}

type mockHistoryRepository struct {
	mock.Mock
}

func (m *mockHistoryRepository) Save(h *domain.ShoutHistory) error {
	args := m.Called(h)
	return args.Error(0)
}

func (m *mockHistoryRepository) FindAll() ([]*domain.ShoutHistory, error) {
	// HistoryUsecase のテストで利用
	args := m.Called()
	return args.Get(0).([]*domain.ShoutHistory), args.Error(1)
}
