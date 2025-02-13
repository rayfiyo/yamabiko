// internal/usecase/shout_usecase_test.go

package usecase_test

import (
	"errors"
	"testing"

	"github.com/rayfiyo/yamabiko/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 観点: 正常系
// 内容: 非デモモードで呼び出し → GeminiClientから6個の応答が返り、DB保存も成功
// 期待する動作: Shoutがエラーなく6つのレスポンスを返す & DB保存成功
func TestShoutUsecase_Shout_Normal(t *testing.T) {
	mockGC := new(mockGeminiClient)
	mockRepo := new(mockHistoryRepository)

	// 6件のレスポンスをモックで用意
	responses := []string{"res1", "res2", "res3", "res4", "res5", "res6"}
	mockGC.On("GenerateResponses", "hello").Return(responses, nil).Once()

	// 保存が正常に完了する
	mockRepo.On("Save", mock.AnythingOfType("*domain.ShoutHistory")).Return(nil).Once()

	uc := usecase.NewShoutUsecase(mockGC, mockRepo)

	result, err := uc.Shout("hello", false)
	assert.NoError(t, err, "error should not occur in normal case")
	assert.Len(t, result, 6, "should return 6 responses")

	// モック呼び出し回数の検証
	mockGC.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

// 観点: 正常系 (ただしDemoモード)
// 内容: Demoモードの場合はGeminiClientを呼ばずに固定メッセージが返る
// 期待する動作: Demo用の固定レスポンスが6つ返却 & DB保存される
func TestShoutUsecase_Shout_DemoMode(t *testing.T) {
	mockGC := new(mockGeminiClient)
	mockRepo := new(mockHistoryRepository)

	// DemoモードならGeminiClient.GenerateResponsesは呼ばれない
	// -> 何も設定しない(あるいは呼ばれないことを明示的に期待する)
	mockGC.AssertNotCalled(t, "GenerateResponses", mock.Anything)

	// DB保存が正常完了する想定
	mockRepo.On("Save", mock.AnythingOfType("*domain.ShoutHistory")).Return(nil).Once()

	uc := usecase.NewShoutUsecase(mockGC, mockRepo)

	result, err := uc.Shout("dummy voice", true)
	assert.NoError(t, err)
	assert.Len(t, result, 6, "demo mode should return 6 fixed responses")

	mockRepo.AssertExpectations(t)
}

// 観点: 異常系
// 内容: GeminiClient.GenerateResponses でエラーが返ってくる場合
// 期待する動作: ShoutUsecase.Shout がエラーを返す (DB保存されない)
func TestShoutUsecase_Shout_ErrorFromGeminiClient(t *testing.T) {
	mockGC := new(mockGeminiClient)
	mockRepo := new(mockHistoryRepository)

	mockGC.On("GenerateResponses", "hello").Return([]string{}, errors.New("gemini error")).Once()
	// DB Saveは呼ばれない想定
	mockRepo.AssertNotCalled(t, "Save", mock.Anything)

	uc := usecase.NewShoutUsecase(mockGC, mockRepo)

	result, err := uc.Shout("hello", false)
	assert.Error(t, err, "error should occur from gemini client")
	assert.Nil(t, result, "result should be nil when error occurs")
	mockGC.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

// 観点: 異常系
// 内容: GeminiClient.GenerateResponses が6件未満のレスポンスを返した場合
// 期待する動作: エラーが返り、DB保存されない
func TestShoutUsecase_Shout_InsufficientResponses(t *testing.T) {
	mockGC := new(mockGeminiClient)
	mockRepo := new(mockHistoryRepository)

	mockGC.On("GenerateResponses", "hello").Return([]string{"res1", "res2"}, nil).Once()

	// DB Saveは呼ばれない想定
	mockRepo.AssertNotCalled(t, "Save", mock.Anything)

	uc := usecase.NewShoutUsecase(mockGC, mockRepo)

	result, err := uc.Shout("hello", false)
	assert.Error(t, err)
	assert.Nil(t, result)
	mockGC.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

// 観点: 異常系
// 内容: DB保存時にエラーが発生した場合
// 期待する動作: エラーが返り、レスポンスは nil
func TestShoutUsecase_Shout_DBSaveError(t *testing.T) {
	mockGC := new(mockGeminiClient)
	mockRepo := new(mockHistoryRepository)

	responses := []string{"r1", "r2", "r3", "r4", "r5", "r6"}
	mockGC.On("GenerateResponses", "hello").Return(responses, nil).Once()

	mockRepo.On("Save", mock.AnythingOfType("*domain.ShoutHistory")).Return(errors.New("db error")).Once()

	uc := usecase.NewShoutUsecase(mockGC, mockRepo)

	result, err := uc.Shout("hello", false)
	assert.Error(t, err)
	assert.Nil(t, result)
	mockGC.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
