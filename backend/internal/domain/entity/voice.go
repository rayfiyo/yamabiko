package entity

import (
	"time"
)

// Voice はユーザーが投稿（shout）した内容を表すドメインエンティティ例
type Voice struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`

	// LLMで付加したメタ情報を保持（オプショナル）
	Analysis string `json:"analysis,omitempty"`
}
