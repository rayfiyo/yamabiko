// internal/domain/entity.go

package domain

import "time"

// ShoutHistory は 1回のシャウトで生成された履歴を表すドメインモデル。
// 一般には JSONB や別テーブル化も検討できるが、ここでは単純化。
type ShoutHistory struct {
	ID        int64     `db:"id"`
	Voice     string    `db:"voice_text"`
	Response1 string    `db:"response1"`
	Response2 string    `db:"response2"`
	Response3 string    `db:"response3"`
	Response4 string    `db:"response4"`
	Response5 string    `db:"response5"`
	Response6 string    `db:"response6"`
	CreatedAt time.Time `db:"created_at"`
}
