package main

import (
	"context"
	"fmt"
	"log/slog"
)

func main() {
	// slogのデフォルトログレベルを設定
	// 開発時: LevelDebug、本番: LevelInfo
	slog.SetLogLoggerLevel(slog.LevelDebug)

	ctx := context.Background()

	slog.Info("アプリケーション開始")

	// 正常なケース
	user, err := GetUser(ctx, 1)
	if err != nil {
		slog.Error("ユーザー取得失敗", "error", err, "user_id", 1)
		fmt.Printf("エラー: %v\n", err)
	} else {
		slog.Info("ユーザー取得成功", "user_id", user.ID, "name", user.Name)
		fmt.Printf("取得成功: %+v\n", user)
	}

	// エラーケース: ユーザーが見つからない
	_, err = GetUser(ctx, 999)
	if err != nil {
		// errors.Isで標準エラーをチェック
		if IsNotFoundError(err) {
			slog.Warn("ユーザーが見つかりません", "user_id", 999)
		} else {
			slog.Error("予期しないエラー", "error", err, "user_id", 999)
		}
		fmt.Printf("エラー: %v\n", err)
	}

	// エラーケース: 無効なID
	_, err = GetUser(ctx, -1)
	if err != nil {
		if IsValidationError(err) {
			slog.Warn("バリデーションエラー", "error", err, "user_id", -1)
		} else {
			slog.Error("予期しないエラー", "error", err, "user_id", -1)
		}
		fmt.Printf("エラー: %v\n", err)
	}

	// ユーザー作成
	newUser, err := CreateUser(ctx, "新規ユーザー", "new@example.com")
	if err != nil {
		slog.Error("ユーザー作成失敗", "error", err)
		fmt.Printf("エラー: %v\n", err)
	} else {
		slog.Info("ユーザー作成成功", "user_id", newUser.ID, "name", newUser.Name)
		fmt.Printf("作成成功: %+v\n", newUser)
	}

	slog.Info("アプリケーション終了")
}
