// cmd/main.go

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rayfiyo/yamabiko/internal/config"
	"github.com/rayfiyo/yamabiko/internal/handler"
	"github.com/rayfiyo/yamabiko/internal/infra/db"
	"github.com/rayfiyo/yamabiko/internal/infra/gemini"
	"github.com/rayfiyo/yamabiko/internal/infra/middleware"
	"github.com/rayfiyo/yamabiko/internal/usecase"
)

func main() {
	// 1) 環境変数読み込み
	cfg := config.Load()

	// 2) DB 接続プールを作成
	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, cfg.ConnString())
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}
	defer dbpool.Close()

	// 3) リポジトリ実装を初期化
	historyRepo := db.NewPostgresHistoryRepo(dbpool)

	// 4) Gemini クライアントを初期化 (実装は infra/gemini/gemini_client.go )
	geminiClient, err := gemini.NewGeminiClient(cfg.GoogleAPIKey)
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}

	// 5) ユースケースを組み立て
	shoutUsecase := usecase.NewShoutUsecase(geminiClient, historyRepo)
	historyUsecase := usecase.NewHistoryUsecase(historyRepo)

	// 6) ルータ生成
	r := mux.NewRouter()

	// レートリミットのミドルウェアを取り付け
	r.Use(middleware.NewRateLimitMiddleware(
		6*time.Second, // 最短間隔
		12,            // 1時間あたりの最大回数
		1*time.Hour,
	))

	// 7) ハンドラ登録
	handler.RegisterHTTPHandlers(r, shoutUsecase, historyUsecase)

	// 8) HTTPサーバ起動
	addr := fmt.Sprintf("0.0.0.0:%d", cfg.Port)
	log.Printf("Server running on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
