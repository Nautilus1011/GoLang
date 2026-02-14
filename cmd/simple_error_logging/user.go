package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
)

// User はユーザーの構造体
type User struct {
	ID    int
	Name  string
	Email string
}

// よく使うエラーを定義（センチネルエラー）
var (
	ErrNotFound   = errors.New("not found")
	ErrValidation = errors.New("validation error")
	ErrInvalidID  = errors.New("invalid id")
	ErrEmptyName  = errors.New("name is required")
	ErrEmptyEmail = errors.New("email is required")
)

// IsNotFoundError はNotFoundエラーかどうかを判定
func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// IsValidationError はバリデーションエラーかどうかを判定
func IsValidationError(err error) bool {
	return errors.Is(err, ErrValidation) ||
		errors.Is(err, ErrInvalidID) ||
		errors.Is(err, ErrEmptyName) ||
		errors.Is(err, ErrEmptyEmail)
}

// GetUser はユーザーIDからユーザーを取得
func GetUser(ctx context.Context, id int) (*User, error) {
	// バリデーション
	if id <= 0 {
		// errors.Joinで複数のエラーを結合（Go 1.20+）
		return nil, errors.Join(ErrValidation, ErrInvalidID, fmt.Errorf("user_id must be positive: got %d", id))
	}

	// データベースからの取得をシミュレート
	user, err := fetchUserFromDB(ctx, id)
	if err != nil {
		// fmt.Errorfの%wでエラーをラップ（Go 1.13+）
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	slog.Debug("ユーザー取得完了", "user_id", id)
	return user, nil
}

// fetchUserFromDB はデータベースからユーザーを取得（モック）
func fetchUserFromDB(ctx context.Context, id int) (*User, error) {
	// モックデータ
	mockUsers := map[int]*User{
		1: {ID: 1, Name: "田中太郎", Email: "tanaka@example.com"},
		2: {ID: 2, Name: "佐藤花子", Email: "sato@example.com"},
		3: {ID: 3, Name: "鈴木一郎", Email: "suzuki@example.com"},
	}

	user, exists := mockUsers[id]
	if !exists {
		// センチネルエラーとカスタムメッセージを組み合わせ
		return nil, fmt.Errorf("user_id=%d: %w", id, ErrNotFound)
	}

	return user, nil
}

// CreateUser は新しいユーザーを作成
func CreateUser(ctx context.Context, name, email string) (*User, error) {
	// バリデーション
	var errs []error

	if name == "" {
		errs = append(errs, ErrEmptyName)
	}
	if email == "" {
		errs = append(errs, ErrEmptyEmail)
	}

	if len(errs) > 0 {
		// 複数のエラーを結合
		return nil, errors.Join(append([]error{ErrValidation}, errs...)...)
	}

	slog.Debug("ユーザー作成開始", "name", name, "email", email)

	// 作成処理をシミュレート
	user := &User{
		ID:    100, // 仮のID
		Name:  name,
		Email: email,
	}

	slog.Debug("ユーザー作成完了", "user_id", user.ID)
	return user, nil
}
