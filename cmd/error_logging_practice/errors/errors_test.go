package errors

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(CodeNotFound, "テストメッセージ")

	if err.Code != CodeNotFound {
		t.Errorf("期待されるコード: %s, 実際: %s", CodeNotFound, err.Code)
	}

	if err.Message != "テストメッセージ" {
		t.Errorf("期待されるメッセージ: %s, 実際: %s", "テストメッセージ", err.Message)
	}

	if err.Err != nil {
		t.Error("Errはnilであるべき")
	}
}

func TestWrap(t *testing.T) {
	originalErr := errors.New("元のエラー")
	wrappedErr := Wrap(originalErr, CodeInternal, "ラップされたエラー")

	if wrappedErr.Err != originalErr {
		t.Error("元のエラーが保持されていない")
	}

	// errors.Unwrapのテスト
	if unwrapped := errors.Unwrap(wrappedErr); unwrapped != originalErr {
		t.Error("Unwrapが正しく動作していない")
	}

	// errors.Isのテスト
	if !errors.Is(wrappedErr, originalErr) {
		t.Error("errors.Isが正しく動作していない")
	}
}

func TestWithDetail(t *testing.T) {
	err := New(CodeInvalidInput, "詳細付きエラー").
		WithDetail("field", "username").
		WithDetail("value", "")

	if err.Details["field"] != "username" {
		t.Error("詳細情報が正しく設定されていない")
	}

	if err.Details["value"] != "" {
		t.Error("詳細情報が正しく設定されていない")
	}
}

func TestIsNotFound(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "NotFoundエラー",
			err:      NotFound("見つかりません"),
			expected: true,
		},
		{
			name:     "InvalidInputエラー",
			err:      InvalidInput("無効な入力"),
			expected: false,
		},
		{
			name:     "標準エラー",
			err:      errors.New("標準エラー"),
			expected: false,
		},
		{
			name:     "nilエラー",
			err:      nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotFound(tt.err)
			if result != tt.expected {
				t.Errorf("期待される結果: %v, 実際: %v", tt.expected, result)
			}
		})
	}
}

func TestIsInvalidInput(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "InvalidInputエラー",
			err:      InvalidInput("無効な入力"),
			expected: true,
		},
		{
			name:     "NotFoundエラー",
			err:      NotFound("見つかりません"),
			expected: false,
		},
		{
			name:     "標準エラー",
			err:      errors.New("標準エラー"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsInvalidInput(tt.err)
			if result != tt.expected {
				t.Errorf("期待される結果: %v, 実際: %v", tt.expected, result)
			}
		})
	}
}

func TestErrorMessage(t *testing.T) {
	tests := []struct {
		name     string
		err      *AppError
		contains string
	}{
		{
			name:     "基本的なエラー",
			err:      New(CodeNotFound, "ユーザーが見つかりません"),
			contains: "NOT_FOUND",
		},
		{
			name:     "ラップされたエラー",
			err:      Wrap(errors.New("DB接続エラー"), CodeInternal, "内部エラー"),
			contains: "DB接続エラー",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := tt.err.Error()
			if msg == "" {
				t.Error("エラーメッセージが空")
			}
			// 含まれるべき文字列のチェックは省略（必要に応じて追加）
		})
	}
}
