// internal/infra/middleware/rate_limit.go

package middleware

import (
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

// ClientRateInfo は特定クライアント(IP)のリクエスト履歴を管理するための構造
type ClientRateInfo struct {
	lastRequest time.Time
	requests    []time.Time // 過去1時間のリクエスト時刻を格納
}

// RateLimitMiddleware は レート制限ミドルウェア
type RateLimitMiddleware struct {
	mu          sync.Mutex
	clients     map[string]*ClientRateInfo
	minInterval time.Duration
	maxRequests int
	window      time.Duration
}

// NewRateLimitMiddleware のコンストラクタ
func NewRateLimitMiddleware(minInterval time.Duration, maxRequests int, window time.Duration) func(http.Handler) http.Handler {
	r := &RateLimitMiddleware{
		clients:     make(map[string]*ClientRateInfo),
		minInterval: minInterval,
		maxRequests: maxRequests,
		window:      window,
	}

	return r.handle
}

func (r *RateLimitMiddleware) handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ip := clientIP(req)
		now := time.Now()

		r.mu.Lock()
		info, ok := r.clients[ip]
		if !ok {
			info = &ClientRateInfo{}
			r.clients[ip] = info
		}

		// Check minimum interval
		if !info.lastRequest.IsZero() && now.Sub(info.lastRequest) < r.minInterval {
			r.mu.Unlock()
			log.Printf("[RateLimit] IP=%s: too frequent requests", ip)
			http.Error(w, "Too Many Requests (min interval)", http.StatusTooManyRequests)
			return
		}

		// Cleanup old requests in the 1-hour window
		cutoff := now.Add(-r.window)
		filtered := make([]time.Time, 0, len(info.requests))
		for _, t := range info.requests {
			if t.After(cutoff) {
				filtered = append(filtered, t)
			}
		}
		info.requests = filtered

		// Check max requests in window
		if len(info.requests) >= r.maxRequests {
			r.mu.Unlock()
			log.Printf("[RateLimit] IP=%s: exceed max requests (%d) in 1 hour", ip, r.maxRequests)
			http.Error(w, "Too Many Requests (over limit)", http.StatusTooManyRequests)
			return
		}

		// Update info
		info.requests = append(info.requests, now)
		info.lastRequest = now

		r.mu.Unlock()

		next.ServeHTTP(w, req)
	})
}

// clientIP は X-Forwarded-For や RemoteAddr からクライアントIPを取得する。
func clientIP(r *http.Request) string {
	// X-Forwarded-For ヘッダの先頭を取得
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}

	// RemoteAddr から IP を抽出
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
