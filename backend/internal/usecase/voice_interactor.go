package usecase

import (
	"time"

	"github.com/rayfiyo/yamabiko/internal/domain/entity"
	"github.com/rayfiyo/yamabiko/internal/domain/repository"
)

// LLMClient は外部LLM（Geminiなど）を呼び出すためのインターフェース
type LLMClient interface {
	Analyze(content string) (string, error)
}

// VoiceInputData は、ユーザからの入力(Shout時)を受け取るための構造体
type VoiceInputData struct {
	Content string
	Demo    bool
}

// VoiceOutputData は、ユースケースが返却する出力データ例
type VoiceOutputData struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Analysis  string    `json:"analysis"`
}

// VoiceInteractor は、Voiceに関するビジネスロジックをまとめたユースケース
type VoiceInteractor struct {
	VoiceRepo repository.VoiceRepository
	LLM       LLMClient
}

func (vi *VoiceInteractor) Shout(in VoiceInputData) (*VoiceOutputData, error) {
	// Voice エンティティを生成
	voice := &entity.Voice{
		Content:   in.Content,
		CreatedAt: time.Now(),
	}

	// LLM による解析（analyze に結果を格納）
	if in.Demo {
		voice.Analysis = "デモメッセージです"
	} else {
		analysis, err := vi.LLM.Analyze(in.Content)
		if err != nil {
			// LLM 呼び出し失敗
			return nil, err
		}
		voice.Analysis = analysis
	}

	// Repository に保存
	err := vi.VoiceRepo.Save(voice)
	if err != nil {
		return nil, err
	}

	return &VoiceOutputData{
		ID:        voice.ID,
		Content:   voice.Content,
		CreatedAt: voice.CreatedAt,
		Analysis:  voice.Analysis,
	}, nil
}

// FindAllVoices はタイムラインの一覧を取得するユースケース例
func (vi *VoiceInteractor) FindAllVoices() ([]*VoiceOutputData, error) {
	voices, err := vi.VoiceRepo.FindAll()
	if err != nil {
		return nil, err
	}

	result := make([]*VoiceOutputData, 0, len(voices))
	for i, v := range voices {
		result[i] = &VoiceOutputData{
			ID:        v.ID,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			Analysis:  v.Analysis,
		}
	}

	return result, nil
}
