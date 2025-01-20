package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/rayfiyo/yamabiko/internal/usecase"
)

// VoiceHandler はHTTPハンドラー
type VoiceHandler struct {
	VoiceUsecase usecase.VoiceInteractor
}

// shoutRequest はPOST /api/shout用のリクエストボディを受け取る
type shoutRequest struct {
	Voice    string `json:"voice"`
	DemoMode bool   `json:"demo_mode"`
}

// ShoutHandler はPOST /api/shoutを処理するハンドラ
func (vh *VoiceHandler) ShoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req shoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// バリデーション
	if req.Voice == "" {
		http.Error(w, `{"message":"テキストボックスを空にすることはできません"}`, http.StatusBadRequest)
		return
	}

	// ユースケースを呼び出す
	out, err := vh.VoiceUsecase.Shout(usecase.VoiceInputData{
		Content: req.Voice,
		Demo:    req.DemoMode,
	})
	if err != nil {
		http.Error(w, `{"message":"shout に失敗"}`, http.StatusInternalServerError)
		return
	}

	// 成功: 200 OK で応答(JSON)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}

// TimelineHandler はGET /api/timelineを処理するハンドラ
func (vh *VoiceHandler) TimelineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	voices, err := vh.VoiceUsecase.FindAllVoices()
	if err != nil {
		http.Error(w, `{"message":"failed to fetch timeline"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(voices)
}
