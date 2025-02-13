// internal/infra/db/postgres_test.go

package db_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rayfiyo/yamabiko/internal/domain"
	"github.com/rayfiyo/yamabiko/internal/infra/db"
	"github.com/stretchr/testify/require"
)

// テスト用のDB接続文字列を設定してください
const testConnString = "postgres://user:password@localhost:5432/test_db"

func setupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	pool, err := pgxpool.New(context.Background(), testConnString)
	require.NoError(t, err, "failed to connect to test DB")

	// 必要に応じてテーブル初期化・クリアなど
	_, err = pool.Exec(context.Background(), "TRUNCATE TABLE shout_history RESTART IDENTITY;")
	require.NoError(t, err, "failed to truncate table")

	return pool
}

// 観点: 正常系
// 内容: 新規レコードを保存後、FindAllで取得できる
// 期待する動作: Insertしたレコードが返ってくる
func TestPostgresHistoryRepo_SaveAndFindAll(t *testing.T) {
	if os.Getenv("INTEGRATION_TEST") != "1" {
		t.Skip("Skipping DB integration test because INTEGRATION_TEST != 1")
	}

	pool := setupTestDB(t)
	defer pool.Close()

	repo := db.NewPostgresHistoryRepo(pool)

	now := time.Now().Truncate(time.Second)
	shout := &domain.ShoutHistory{
		Voice:     "test voice",
		Response1: "r1",
		Response2: "r2",
		Response3: "r3",
		Response4: "r4",
		Response5: "r5",
		Response6: "r6",
		CreatedAt: now,
	}

	err := repo.Save(shout)
	require.NoError(t, err)
	require.NotZero(t, shout.ID, "ID should be set after insert")

	histories, err := repo.FindAll()
	require.NoError(t, err)
	require.Len(t, histories, 1)

	got := histories[0]
	require.Equal(t, "test voice", got.Voice)
	require.Equal(t, "r1", got.Response1)
	require.Equal(t, "r2", got.Response2)
	require.WithinDuration(t, now, got.CreatedAt, time.Second)
}

// 観点: 異常系
// 内容: 列の制約違反やDBエラーなどを発生させる(例: NULL制約違反にわざと空文字やNULLを渡す)
// 期待する動作: Saveがエラーを返す
func TestPostgresHistoryRepo_SaveError(t *testing.T) {
	if os.Getenv("INTEGRATION_TEST") != "1" {
		t.Skip("Skipping DB integration test because INTEGRATION_TEST != 1")
	}

	pool := setupTestDB(t)
	defer pool.Close()

	repo := db.NewPostgresHistoryRepo(pool)

	// voice_text が空(あるいは NOT NULL 制約があるなら)でエラーを期待
	shout := &domain.ShoutHistory{
		Voice:     "", // ここをわざと空にすることでエラー想定
		CreatedAt: time.Now(),
	}

	err := repo.Save(shout)
	require.Error(t, err, "expecting DB error due to missing required field")
}
