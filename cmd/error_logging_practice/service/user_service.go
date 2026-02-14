package service

import (
	"context"
	"fmt"

	"github.com/Nautilus1011/GoLang/cmd/error_logging_practice/errors"
	"github.com/Nautilus1011/GoLang/cmd/error_logging_practice/logger"
)

// User はユーザーのモデル
type User struct {
	ID    int
	Name  string
	Email string
}

// UserService はユーザー関連の業務ロジック
type UserService struct {
	logger *logger.Logger
}

// NewUserService は新しいUserServiceを作成
func NewUserService(log *logger.Logger) *UserService {
	return &UserService{
		logger: log,
	}
}

// GetUser はユーザーIDからユーザーを取得
func (s *UserService) GetUser(ctx context.Context, id int) (*User, error) {
	// バリデーション
	if id <= 0 {
		s.logger.Warn(ctx, "無効なユーザーID", "user_id", id)
		return nil, errors.InvalidInput("ユーザーIDは1以上である必要があります").
			WithDetail("user_id", id)
	}

	// データベースからの取得をシミュレート
	user, err := s.fetchUserFromDB(ctx, id)
	if err != nil {
		// エラーをラップして返す
		return nil, err
	}

	s.logger.Debug(ctx, "ユーザー取得完了", "user_id", id)
	return user, nil
}

// fetchUserFromDB はデータベースからユーザーを取得（モック）
func (s *UserService) fetchUserFromDB(ctx context.Context, id int) (*User, error) {
	// モックデータ
	mockUsers := map[int]*User{
		1: {ID: 1, Name: "田中太郎", Email: "tanaka@example.com"},
		2: {ID: 2, Name: "佐藤花子", Email: "sato@example.com"},
		3: {ID: 3, Name: "鈴木一郎", Email: "suzuki@example.com"},
	}

	user, exists := mockUsers[id]
	if !exists {
		return nil, errors.NotFound(fmt.Sprintf("ユーザーID %d が見つかりません", id)).
			WithDetail("user_id", id)
	}

	return user, nil
}

// CreateUser は新しいユーザーを作成
func (s *UserService) CreateUser(ctx context.Context, name, email string) (*User, error) {
	// バリデーション
	if name == "" {
		return nil, errors.InvalidInput("名前は必須です")
	}
	if email == "" {
		return nil, errors.InvalidInput("メールアドレスは必須です")
	}

	s.logger.Info(ctx, "ユーザー作成開始", "name", name, "email", email)

	// 作成処理をシミュレート
	user := &User{
		ID:    100, // 仮のID
		Name:  name,
		Email: email,
	}

	s.logger.Info(ctx, "ユーザー作成完了", "user_id", user.ID)
	return user, nil
}
