// internal/usecase/history_usecase.go

package usecase

import (
	"github.com/rayfiyo/yamabiko/internal/domain"
)

// HistoryUsecase は /api/history の処理を担う
type HistoryUsecase interface {
	GetHistory() ([]*domain.ShoutHistory, error)
}

type historyUsecaseImpl struct {
	historyRepo domain.HistoryRepository
}

func NewHistoryUsecase(hr domain.HistoryRepository) HistoryUsecase {
	return &historyUsecaseImpl{
		historyRepo: hr,
	}
}

func (u *historyUsecaseImpl) GetHistory() ([]*domain.ShoutHistory, error) {
	return u.historyRepo.FindAll()
}
