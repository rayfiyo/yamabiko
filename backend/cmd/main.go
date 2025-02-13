// cmd/main.go

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rayfiyo/yamabiko/utils/config"
	"github.com/rayfiyo/yamabiko/internal/handler"
	"github.com/rayfiyo/yamabiko/internal/infra/db"
	"github.com/rayfiyo/yamabiko/internal/infra/gemini"
	"github.com/rayfiyo/yamabiko/internal/usecase"
	"github.com/rs/cors"
)

func main() {
	// 環境変数読み込み
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// DB 接続プールを作成
	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, cfg.ConnString())
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}
	defer dbpool.Close()

	// リポジトリ実装を初期化
	historyRepo := db.NewPostgresHistoryRepo(dbpool)

	// Gemini クライアントを初期化 (実装は infra/gemini/gemini_client.go )
	geminiClient, err := gemini.NewGeminiClient(cfg.GoogleAPIKey)
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}

	// ユースケースを組み立て
	shoutUsecase := usecase.NewShoutUsecase(geminiClient, historyRepo)
	historyUsecase := usecase.NewHistoryUsecase(historyRepo)

	// ルータ生成
	r := mux.NewRouter()

	// ハンドラ登録
	handler.RegisterHTTPHandlers(r, shoutUsecase, historyUsecase)

	// ライブラリで CORS ハンドラを作る
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
	})

	// mux.Router を包む形でハンドラに
	h := c.Handler(r)

	// HTTPサーバ起動
	addr := fmt.Sprintf("0.0.0.0:%d", cfg.Port)
	log.Printf("Server running on %s", addr)
	if err := http.ListenAndServe(addr, h); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
