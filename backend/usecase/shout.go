package usecase

import (
	"github.com/rayfiyo/yamabiko/entity"
	"github.com/rayfiyo/yamabiko/service"
	"github.com/rayfiyo/yamabiko/utils/errors"
)

type ShoutUseCase interface {
	Execute(voice string) (entity.Shout, error)
}

type shoutUseCase struct {
	echoService service.EchoService
}

func NewShoutUseCase(echoService service.EchoService) *shoutUseCase {
	return &shoutUseCase{echoService: echoService}
}

func (s *shoutUseCase) Execute(voice string) (entity.Shout, error) {
	// バリデーション
	if voice == "" {
		return entity.Shout{}, errors.ErrEmptyVoice
	}

	// echo を生成
	echo := s.echoService.GenerateEcho(voice)

	// echoes を生成
	var echoes []string
	echoes = append(echoes, echo)

	return entity.Shout{Voice: voice, Echo: echoes}, nil
}
