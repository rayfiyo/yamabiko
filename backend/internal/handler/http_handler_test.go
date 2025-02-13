// internal/handler/http_handler_test.go

package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rayfiyo/yamabiko/internal/domain"
	"github.com/rayfiyo/yamabiko/internal/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- モック定義 ---

type mockShoutUsecase struct {
	mock.Mock
}

func (m *mockShoutUsecase) Shout(voice string, isDemo bool) ([]string, error) {
	args := m.Called(voice, isDemo)
	return args.Get(0).([]string), args.Error(1)
}

type mockHistoryUsecase struct {
	mock.Mock
}

func (m *mockHistoryUsecase) GetHistory() ([]*domain.ShoutHistory, error) {
	args := m.Called()
	return args.Get(0).([]*domain.ShoutHistory), args.Error(1)
}

// --- テスト ---

// 観点: 正常系
// 内容: /api/shout に正しいJSONボディをPOSTした場合
// 期待する動作: 200が返り、レスポンスBODYに6個のレスポンスが含まれる
func TestShoutHandler_Normal(t *testing.T) {
	// モック用意
	mockUC := new(mockShoutUsecase)
	mockHistUC := new(mockHistoryUsecase)

	// 6個のレスポンスを返すモック設定
	mockUC.On("Shout", "hello", false).Return([]string{"r1", "r2", "r3", "r4", "r5", "r6"}, nil).Once()

	r := mux.NewRouter()
	handler.RegisterHTTPHandlers(r, mockUC, mockHistUC)

	// テスト用リクエスト
	body, _ := json.Marshal(map[string]interface{}{
		"voice":    "hello",
		"demoMode": false,
	})
	req := httptest.NewRequest(http.MethodPost, "/api/shout", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// レスポンス録画用
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp struct {
		Responses []string `json:"responses"`
	}
	json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Len(t, resp.Responses, 6)

	mockUC.AssertExpectations(t)
	mockHistUC.AssertExpectations(t)
}

// 観点: 異常系
// 内容: /api/shout に不正なJSON (voiceが空)
// 期待する動作: 400(BadRequest)が返る
func TestShoutHandler_BadRequest(t *testing.T) {
	mockUC := new(mockShoutUsecase)
	mockHistUC := new(mockHistoryUsecase)

	r := mux.NewRouter()
	handler.RegisterHTTPHandlers(r, mockUC, mockHistUC)

	body, _ := json.Marshal(map[string]interface{}{
		"voice": "",
	})
	req := httptest.NewRequest(http.MethodPost, "/api/shout", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// 観点: 異常系
// 内容: ユースケースがエラーを返す (DBエラーなど想定)
// 期待する動作: 500(InternalServerError)が返る
func TestShoutHandler_InternalServerError(t *testing.T) {
	mockUC := new(mockShoutUsecase)
	mockHistUC := new(mockHistoryUsecase)

	// ユースケースがエラーを返すモック
	mockUC.On("Shout", "hello", false).Return([]string(nil), errors.New("some error")).Once()

	r := mux.NewRouter()
	handler.RegisterHTTPHandlers(r, mockUC, mockHistUC)

	body, _ := json.Marshal(map[string]interface{}{
		"voice": "hello",
	})
	req := httptest.NewRequest(http.MethodPost, "/api/shout", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

// 観点: 正常系
// 内容: /api/history をGETしたとき、データが存在する
// 期待する動作: 200が返り、JSONに履歴一覧が含まれる
func TestHistoryHandler_Normal(t *testing.T) {
	mockUC := new(mockShoutUsecase)
	mockHistUC := new(mockHistoryUsecase)

	dummyData := []*domain.ShoutHistory{
		{ID: 1, Voice: "v1"},
		{ID: 2, Voice: "v2"},
	}
	mockHistUC.On("GetHistory").Return(dummyData, nil).Once()

	r := mux.NewRouter()
	handler.RegisterHTTPHandlers(r, mockUC, mockHistUC)

	req := httptest.NewRequest(http.MethodGet, "/api/history", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp []domain.ShoutHistory
	json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Len(t, resp, 2)
	assert.Equal(t, int64(1), resp[0].ID)

	mockUC.AssertExpectations(t)
	mockHistUC.AssertExpectations(t)
}

// 観点: 異常系
// 内容: リポジトリがエラーを返す → usecase がエラー → handler もエラー(500)を返す
// 期待する動作: 500が返る
func TestHistoryHandler_Error(t *testing.T) {
	mockUC := new(mockShoutUsecase)
	mockHistUC := new(mockHistoryUsecase)

	mockHistUC.On("GetHistory").Return([]*domain.ShoutHistory(nil), errors.New("db error")).Once()

	r := mux.NewRouter()
	handler.RegisterHTTPHandlers(r, mockUC, mockHistUC)

	req := httptest.NewRequest(http.MethodGet, "/api/history", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockUC.AssertExpectations(t)
	mockHistUC.AssertExpectations(t)
}
