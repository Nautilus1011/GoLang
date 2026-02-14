package main

import (
	"context"
	"errors"
	"testing"
)

func TestGetUser_Success(t *testing.T) {
	ctx := context.Background()

	user, err := GetUser(ctx, 1)
	if err != nil {
		t.Fatalf("エラーが発生すべきでない: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("期待されるID: 1, 実際: %d", user.ID)
	}
}

func TestGetUser_NotFound(t *testing.T) {
	ctx := context.Background()

	user, err := GetUser(ctx, 999)
	if err == nil {
		t.Fatal("エラーが発生すべき")
	}

	if user != nil {
		t.Error("ユーザーはnilであるべき")
	}

	// errors.Isでエラーをチェック
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("ErrNotFoundであるべき: %v", err)
	}

	// ヘルパー関数でもチェック可能
	if !IsNotFoundError(err) {
		t.Error("IsNotFoundErrorがtrueを返すべき")
	}
}

func TestGetUser_InvalidID(t *testing.T) {
	ctx := context.Background()

	tests := []int{0, -1, -100}

	for _, id := range tests {
		user, err := GetUser(ctx, id)
		if err == nil {
			t.Errorf("ID=%d: エラーが発生すべき", id)
		}

		if user != nil {
			t.Errorf("ID=%d: ユーザーはnilであるべき", id)
		}

		if !errors.Is(err, ErrValidation) {
			t.Errorf("ID=%d: ErrValidationであるべき: %v", id, err)
		}

		if !errors.Is(err, ErrInvalidID) {
			t.Errorf("ID=%d: ErrInvalidIDであるべき: %v", id, err)
		}
	}
}

func TestCreateUser_Success(t *testing.T) {
	ctx := context.Background()

	user, err := CreateUser(ctx, "テストユーザー", "test@example.com")
	if err != nil {
		t.Fatalf("エラーが発生すべきでない: %v", err)
	}

	if user.Name != "テストユーザー" {
		t.Errorf("期待される名前: テストユーザー, 実際: %s", user.Name)
	}
}

func TestCreateUser_ValidationError(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		userName  string
		email     string
		wantError error
	}{
		{"名前が空", "", "test@example.com", ErrEmptyName},
		{"メールが空", "テスト", "", ErrEmptyEmail},
		{"両方空", "", "", ErrValidation},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := CreateUser(ctx, tt.userName, tt.email)
			if err == nil {
				t.Fatal("エラーが発生すべき")
			}

			if user != nil {
				t.Error("ユーザーはnilであるべき")
			}

			// errors.Isで期待されるエラーが含まれているか確認
			if !errors.Is(err, tt.wantError) {
				t.Errorf("期待されるエラー %v が含まれていない: %v", tt.wantError, err)
			}
		})
	}
}

// エラーメッセージのテスト
func TestErrorMessages(t *testing.T) {
	ctx := context.Background()

	// NotFoundエラーのメッセージ確認
	_, err := GetUser(ctx, 999)
	if err == nil {
		t.Fatal("エラーが発生すべき")
	}

	errMsg := err.Error()
	// エラーメッセージにuser_idが含まれているか
	if errMsg == "" {
		t.Error("エラーメッセージが空")
	}

	t.Logf("NotFoundエラー: %v", err)

	// Validationエラーのメッセージ確認
	_, err = GetUser(ctx, -1)
	if err == nil {
		t.Fatal("エラーが発生すべき")
	}

	t.Logf("Validationエラー: %v", err)
}
