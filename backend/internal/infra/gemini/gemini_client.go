// internal/infra/gemini/gemini_client.go

package gemini

import (
	"context"
	"fmt"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// GeminiClient インターフェース
type GeminiClient interface {
	GenerateResponses(voice string) ([]string, error)
}

// geminiClientImpl は実際に Gemini(PaLM) API を呼び出す実装
type geminiClientImpl struct {
	client *genai.Client
}

// NewGeminiClient : SDKクライアント初期化 (要: GOOGLE_API_KEY)
func NewGeminiClient(apiKey string) (GeminiClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("gemini API key is empty or not set")
	}
	c, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &geminiClientImpl{client: c}, nil
}

// GenerateResponses : voice に対する6通りの応答を生成して返す
func (g *geminiClientImpl) GenerateResponses(voice string) ([]string, error) {
	// コンテキスト(タイムアウト10秒)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 使用したいモデルを指定:
	model := g.client.GenerativeModel("gemini-1.5-pro-latest")

	// 生成パラメータを設定
	model.ResponseMIMEType = "application/json"
	model.SetCandidateCount(6)
	resp, err := model.GenerateContent(ctx, genai.Text(voice))
	if err != nil {
		return nil, err
	}

	// Gemini から複数候補が返ってくる -> 取り出し
	// resp.Candidates は slice で、たとえば 6通り分が入っている
	// 各 Candidate の Content.Parts に実際のテキスト (genai.Text) が含まれる
	var results []string
	for _, c := range resp.Candidates {
		// 参考: https://pkg.go.dev/github.com/google/generative-ai-go/genai#Candidate
		// Candidate.Content.Parts の中に genai.Text が入っている
		// 1つのCandidateに複数partがある可能性があるので連結する
		textParts := ""
		for _, part := range c.Content.Parts {
			if txt, ok := part.(genai.Text); ok {
				textParts += string(txt)
			}
		}
		results = append(results, textParts)
	}

	// もし 6通り未満しか返ってこなかった場合のチェック (状況に応じて)
	if len(results) < 6 {
		// 必要に応じてエラーにする or 足りないぶんダミー埋め
		return nil, fmt.Errorf("only %d results returned from Gemini", len(results))
	}

	return results, nil
}
