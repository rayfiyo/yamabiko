package repository

import "github.com/rayfiyo/yamabiko/internal/domain/entity"

// VoiceRepository は、Voiceエンティティを保存/取得するためのインターフェースを定義
// 実装は infrastructure 配下などで行う
type VoiceRepository interface {
	Save(voice *entity.Voice) error
	FindAll() ([]*entity.Voice, error)
}
