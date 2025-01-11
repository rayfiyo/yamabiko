package service

type EchoService interface {
	GenerateEcho(voice string) string
}

type echoService struct{}

func NewEchoService() *echoService {
	return &echoService{}
}

func (e *echoService) GenerateEcho(voice string) string {
	// 例: Echoを生成するシンプルなロジック
	return "Echo: " + voice
}
