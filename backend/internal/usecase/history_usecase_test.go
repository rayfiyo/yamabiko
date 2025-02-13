// internal/usecase/history_usecase_test.go

package usecase_test

import (
	"errors"
	"testing"

	"github.com/rayfiyo/yamabiko/internal/domain"
	"github.com/rayfiyo/yamabiko/internal/usecase"
	"github.com/stretchr/testify/assert"
)

// 観点: 正常系
// 内容: リポジトリが問題なくデータを返す
// 期待する動作: 取得した履歴を正しく返す
func TestHistoryUsecase_GetHistory_Normal(t *testing.T) {
	mockRepo := new(mockHistoryRepository)
	// テスト用ダミーデータ
	dummyHistories := []*domain.ShoutHistory{
		{ID: 1, Voice: "voice1"},
		{ID: 2, Voice: "voice2"},
	}

	// FindAllが正常に返る
	mockRepo.On("FindAll").Return(dummyHistories, nil).Once()

	uc := usecase.NewHistoryUsecase(mockRepo)
	result, err := uc.GetHistory()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, dummyHistories[0].Voice, result[0].Voice)

	mockRepo.AssertExpectations(t)
}

// 観点: 異常系
// 内容: リポジトリがエラーを返す
// 期待する動作: ユースケースもエラーを返す
func TestHistoryUsecase_GetHistory_RepoError(t *testing.T) {
	mockRepo := new(mockHistoryRepository)
	mockRepo.On("FindAll").Return([]*domain.ShoutHistory(nil), errors.New("db error")).Once()

	uc := usecase.NewHistoryUsecase(mockRepo)
	result, err := uc.GetHistory()
	assert.Error(t, err)
	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}
