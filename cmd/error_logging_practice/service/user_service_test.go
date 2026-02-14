package service

import (
	"context"
	"testing"

	"github.com/Nautilus1011/GoLang/cmd/error_logging_practice/errors"
	"github.com/Nautilus1011/GoLang/cmd/error_logging_practice/logger"
)

func TestGetUser_Success(t *testing.T) {
	log := logger.NewLogger(nil, logger.LevelError) // テスト時はエラーのみ出力
	service := NewUserService(log)
	ctx := context.Background()

	user, err := service.GetUser(ctx, 1)
	if err != nil {
		t.Fatalf("エラーが発生すべきでない: %v", err)
	}

	if user == nil {
		t.Fatal("ユーザーがnilです")
	}

	if user.ID != 1 {
		t.Errorf("期待されるID: 1, 実際: %d", user.ID)
	}

	if user.Name == "" {
		t.Error("ユーザー名が空です")
	}
}

func TestGetUser_NotFound(t *testing.T) {
	log := logger.NewLogger(nil, logger.LevelError)
	service := NewUserService(log)
	ctx := context.Background()

	user, err := service.GetUser(ctx, 999)
	if err == nil {
		t.Fatal("エラーが発生すべき")
	}

	if user != nil {
		t.Error("ユーザーはnilであるべき")
	}

	// エラータイプの検証
	if !errors.IsNotFound(err) {
		t.Errorf("NotFoundエラーであるべき: %v", err)
	}
}

func TestGetUser_InvalidInput(t *testing.T) {
	log := logger.NewLogger(nil, logger.LevelError)
	service := NewUserService(log)
	ctx := context.Background()

	tests := []struct {
		name string
		id   int
	}{
		{"ゼロ", 0},
		{"負の数", -1},
		{"負の大きな数", -100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := service.GetUser(ctx, tt.id)
			if err == nil {
				t.Fatal("エラーが発生すべき")
			}

			if user != nil {
				t.Error("ユーザーはnilであるべき")
			}

			if !errors.IsInvalidInput(err) {
				t.Errorf("InvalidInputエラーであるべき: %v", err)
			}
		})
	}
}

func TestCreateUser_Success(t *testing.T) {
	log := logger.NewLogger(nil, logger.LevelError)
	service := NewUserService(log)
	ctx := context.Background()

	user, err := service.CreateUser(ctx, "テストユーザー", "test@example.com")
	if err != nil {
		t.Fatalf("エラーが発生すべきでない: %v", err)
	}

	if user == nil {
		t.Fatal("ユーザーがnilです")
	}

	if user.Name != "テストユーザー" {
		t.Errorf("期待される名前: テストユーザー, 実際: %s", user.Name)
	}

	if user.Email != "test@example.com" {
		t.Errorf("期待されるメール: test@example.com, 実際: %s", user.Email)
	}
}

func TestCreateUser_ValidationError(t *testing.T) {
	log := logger.NewLogger(nil, logger.LevelError)
	service := NewUserService(log)
	ctx := context.Background()

	tests := []struct {
		name  string
		uname string
		email string
	}{
		{"名前が空", "", "test@example.com"},
		{"メールが空", "テストユーザー", ""},
		{"両方空", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := service.CreateUser(ctx, tt.uname, tt.email)
			if err == nil {
				t.Fatal("エラーが発生すべき")
			}

			if user != nil {
				t.Error("ユーザーはnilであるべき")
			}

			if !errors.IsInvalidInput(err) {
				t.Errorf("InvalidInputエラーであるべき: %v", err)
			}
		})
	}
}

// テーブル駆動テストの例
func TestGetUser_TableDriven(t *testing.T) {
	log := logger.NewLogger(nil, logger.LevelError)
	service := NewUserService(log)
	ctx := context.Background()

	tests := []struct {
		name        string
		id          int
		expectError bool
		errorType   func(error) bool
	}{
		{
			name:        "正常なケース_ID1",
			id:          1,
			expectError: false,
		},
		{
			name:        "正常なケース_ID2",
			id:          2,
			expectError: false,
		},
		{
			name:        "NotFoundエラー",
			id:          999,
			expectError: true,
			errorType:   errors.IsNotFound,
		},
		{
			name:        "無効な入力_ゼロ",
			id:          0,
			expectError: true,
			errorType:   errors.IsInvalidInput,
		},
		{
			name:        "無効な入力_負の数",
			id:          -1,
			expectError: true,
			errorType:   errors.IsInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := service.GetUser(ctx, tt.id)

			if tt.expectError {
				if err == nil {
					t.Fatal("エラーが発生すべき")
				}
				if tt.errorType != nil && !tt.errorType(err) {
					t.Errorf("期待されるエラータイプでない: %v", err)
				}
				if user != nil {
					t.Error("ユーザーはnilであるべき")
				}
			} else {
				if err != nil {
					t.Fatalf("エラーが発生すべきでない: %v", err)
				}
				if user == nil {
					t.Fatal("ユーザーがnilです")
				}
				if user.ID != tt.id {
					t.Errorf("期待されるID: %d, 実際: %d", tt.id, user.ID)
				}
			}
		})
	}
}
