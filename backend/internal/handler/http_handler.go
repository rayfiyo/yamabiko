// internal/handler/http_handler.go

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rayfiyo/yamabiko/internal/domain"
	"github.com/rayfiyo/yamabiko/utils/consts"
)

// 依存注入のためのインターフェース
type ShoutUsecase interface {
	Shout(voice string) ([]string, error)
}

type HistoryUsecase interface {
	GetHistory() ([]*domain.ShoutHistory, error)
}

func RegisterHTTPHandlers(r *mux.Router, s ShoutUsecase, h HistoryUsecase) {
	r.HandleFunc("/api/shout", shoutHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/api/history", historyHandler(h)).Methods(http.MethodGet)
}

// --------------------------------------------
// /api/shout
// --------------------------------------------
type shoutRequest struct {
	Voice    string `json:"voice"`
	DemoMode bool   `json:"demoMode"`
}

type shoutResponse struct {
	Responses []string `json:"responses"`
}

func shoutHandler(shoutUC ShoutUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエストボディをパース
		var req shoutRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// バリデーション
		if req.Voice == "" {
			http.Error(w, "voice is required", http.StatusBadRequest)
			return
		}

		// デモモードの場合
		if req.DemoMode {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(req.Voice + consts.DemoMsg)
			return
		}

		// ユースケース呼び出し
		responses, err := shoutUC.Shout(req.Voice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// JSON で応答
		resp := shoutResponse{Responses: responses}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

// --------------------------------------------
// /api/history
// --------------------------------------------
func historyHandler(historyUC HistoryUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ユースケース呼び出し
		histories, err := historyUC.GetHistory()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(histories)
	}
}
