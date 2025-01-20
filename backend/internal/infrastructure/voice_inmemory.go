package infrastructure

import (
	"sync"

	"github.com/rayfiyo/yamabiko/internal/entity"
	"github.com/rayfiyo/yamabiko/internal/repository"
)

// InMemoryVoiceRepository はメモリ上にVoiceを保存する簡易実
type InMemoryVoiceRepository struct {
	mu     sync.Mutex
	voices []*entity.Voice
	nextID int64
}

func NewInMemoryVoiceRepository() repository.VoiceRepository {
	return &InMemoryVoiceRepository{
		voices: make([]*entity.Voice, 0),
		nextID: 1,
	}
}

func (r *InMemoryVoiceRepository) Save(voice *entity.Voice) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	voice.ID = r.nextID
	r.nextID++
	r.voices = append(r.voices, voice)

	return nil
}

func (r *InMemoryVoiceRepository) FindAll() ([]*entity.Voice, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// そのまま返すとスライスが書き換えられる可能性があるのでコピーするのが理想
	copied := make([]*entity.Voice, len(r.voices))
	copy(copied, r.voices)
	return copied, nil
}
