// internal/infra/middleware/rate_limit_test.go

package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rayfiyo/yamabiko/internal/infra/middleware"
	"github.com/stretchr/testify/assert"
)

// 観点: 正常系
// 内容: 最低間隔を満たしていればOK、回数制限にも引っかからない
// 期待する動作: 全部ステータス200で通る
func TestRateLimitMiddleware_Normal(t *testing.T) {
	// 例: minInterval=1ms, maxRequests=3, window=1h
	rl := middleware.NewRateLimitMiddleware(1*time.Millisecond, 3, time.Hour)
	handler := rl(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/api/history", nil)
	rr := httptest.NewRecorder()

	// 1回目
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	// 少し待機して再度(間隔クリア)
	time.Sleep(2 * time.Millisecond)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	// 3回目
	time.Sleep(2 * time.Millisecond)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

// 観点: 異常系
// 内容: minIntervalを満たさずに短時間に連続リクエスト
// 期待する動作: 429(TooManyRequests)が返る
func TestRateLimitMiddleware_TooFrequent(t *testing.T) {
	rl := middleware.NewRateLimitMiddleware(100*time.Millisecond, 3, time.Hour)
	handler := rl(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/api/history", nil)
	rr := httptest.NewRecorder()

	// 1回目はOK
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	// すぐに2回目(minIntervalを満たしていない)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusTooManyRequests, rr.Code)
}

// 観点: 異常系
// 内容: 1時間ウィンドウ中に maxRequests を超えるリクエスト
// 期待する動作: 429(TooManyRequests)が返る
func TestRateLimitMiddleware_MaxRequests(t *testing.T) {
	rl := middleware.NewRateLimitMiddleware(1*time.Millisecond, 2, time.Hour)
	handler := rl(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/api/history", nil)

	// 1回目
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	// 2回目(OK)
	time.Sleep(2 * time.Millisecond)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	// 3回目 → maxRequests=2 を超える
	time.Sleep(2 * time.Millisecond)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusTooManyRequests, rr.Code)
}
