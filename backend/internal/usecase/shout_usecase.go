// internal/usecase/shout_usecase.go

package usecase

import (
	"fmt"
	"time"

	"github.com/rayfiyo/yamabiko/internal/domain"
	"github.com/rayfiyo/yamabiko/utils/consts"
)

// ShoutUsecase は /api/shout の処理を担う
type ShoutUsecase interface {
	Shout(voice string, isDemo bool) ([]string, error)
}

type shoutUsecaseImpl struct {
	geminiClient GeminiClient // Gemini API呼び出し用
	historyRepo  domain.HistoryRepository
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

func (u *shoutUsecaseImpl) Shout(voice string, isDemo bool) ([]string, error) {
	var responses []string

	if isDemo {
		// デモモードの場合は固定の応答を返す
		responses = []string{
			consts.DemoMsg1,
			consts.DemoMsg2,
			consts.DemoMsg3,
			consts.DemoMsg4,
			consts.DemoMsg5,
			consts.DemoMsg6,
		}
	} else {
		// デモモードでない場合は Gemini API から応答を取得する
		var err error
		responses, err = u.geminiClient.GenerateResponses(voice)
		if err != nil {
			return nil, err
		}
	}

	if len(responses) < 6 {
		// 6件未満の場合のエラー
		return nil, fmt.Errorf("response length is less than 6")
	}

	// DB 保存
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

	// 6件の応答を返す
	return responses, nil
}
