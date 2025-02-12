// internal/domain/repository.go

package domain

// HistoryRepository は ShoutHistory の永続化を行うためのインターフェース
type HistoryRepository interface {
	Save(history *ShoutHistory) error
	FindAll() ([]*ShoutHistory, error)
}
