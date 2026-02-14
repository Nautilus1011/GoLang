package errors

import (
	"errors"
	"fmt"
)

// カスタムエラー型の定義
// 実務では、エラーの種類を区別できるようにすることが重要

// ErrorCode はエラーコードの型
type ErrorCode string

const (
	// エラーコードの定義
	CodeNotFound     ErrorCode = "NOT_FOUND"
	CodeInvalidInput ErrorCode = "INVALID_INPUT"
	CodeInternal     ErrorCode = "INTERNAL"
	CodeUnauthorized ErrorCode = "UNAUTHORIZED"
)

// AppError はアプリケーション固有のエラー
type AppError struct {
	Code    ErrorCode              // エラーコード
	Message string                 // ユーザー向けメッセージ
	Err     error                  // 元のエラー
	Details map[string]interface{} // 追加の詳細情報
}

// Error はerrorインターフェースの実装
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap はエラーチェーンのサポート (Go 1.13+)
func (e *AppError) Unwrap() error {
	return e.Err
}

// New はAppErrorを作成するヘルパー関数
func New(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: make(map[string]interface{}),
	}
}

// Wrap は既存のエラーをラップする
func Wrap(err error, code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
		Details: make(map[string]interface{}),
	}
}

// WithDetail は詳細情報を追加する
func (e *AppError) WithDetail(key string, value interface{}) *AppError {
	e.Details[key] = value
	return e
}

// 便利なコンストラクタ関数

// NotFound はNotFoundエラーを作成
func NotFound(message string) *AppError {
	return New(CodeNotFound, message)
}

// InvalidInput は無効な入力エラーを作成
func InvalidInput(message string) *AppError {
	return New(CodeInvalidInput, message)
}

// Internal は内部エラーを作成
func Internal(message string, err error) *AppError {
	return Wrap(err, CodeInternal, message)
}

// IsNotFound はNotFoundエラーかどうかを判定
func IsNotFound(err error) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == CodeNotFound
	}
	return false
}

// IsInvalidInput は無効な入力エラーかどうかを判定
func IsInvalidInput(err error) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == CodeInvalidInput
	}
	return false
}
