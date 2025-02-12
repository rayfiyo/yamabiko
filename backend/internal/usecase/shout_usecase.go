// internal/usecase/shout_usecase.go

package usecase

import (
    "time"

    "github.com/rayfiyo/yamabiko/internal/domain"
)

// ShoutUsecase は /api/shout の処理を担う
type ShoutUsecase interface {
    Shout(voice string) ([]string, error)
}

type shoutUsecaseImpl struct {
    geminiClient   GeminiClient  // Gemini API呼び出し用
    historyRepo    domain.HistoryRepository
}

// GeminiClient は usecase 層が必要とする外部AI呼び出しのインターフェース
// (internal/infra/gemini/GeminiClient と一致していればよい)
type GeminiClient interface {
    GenerateResponses(voice string) ([]string, error)
}

func NewShoutUsecase(gc GeminiClient, hr domain.HistoryRepository) ShoutUsecase {
    return &shoutUsecaseImpl{
        geminiClient: gc,
        historyRepo:  hr,
    }
}

func (u *shoutUsecaseImpl) Shout(voice string) ([]string, error) {
    // 1) Gemini で応答生成
    responses, err := u.geminiClient.GenerateResponses(voice)
    if err != nil {
        return nil, err
    }

    if len(responses) < 6 {
        // 万が一APIが6件返せなかった場合のエラー
        return nil, err
    }

    // 2) DB 保存
    history := &domain.ShoutHistory{
        Voice:     voice,
        Response1: responses[0],
        Response2: responses[1],
        Response3: responses[2],
        Response4: responses[3],
        Response5: responses[4],
        Response6: responses[5],
        CreatedAt: time.Now(),
    }

    if err := u.historyRepo.Save(history); err != nil {
        return nil, err
    }

    // 3) 6件の応答を返す
    return responses, nil
}

