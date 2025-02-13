// internal/infra/gemini/gemini_client_test.go

package gemini_test

import (
	"testing"

	"github.com/rayfiyo/yamabiko/internal/infra/gemini"
	"github.com/stretchr/testify/assert"
)

// ここではAPIコールしない かつ， モック化するなら gemini_client_implの内部で
// genai.Client をモックにする必要があるので，あまり書くことは少ない

// 観点: 異常系
// 内容: APIKeyが空の場合
// 期待する動作: エラーが返る
func TestNewGeminiClient_EmptyAPIKey(t *testing.T) {
	cli, err := gemini.NewGeminiClient("")
	assert.Error(t, err)
	assert.Nil(t, cli)
}
