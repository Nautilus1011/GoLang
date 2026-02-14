package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Nautilus1011/GoLang/cmd/error_logging_practice/logger"
	"github.com/Nautilus1011/GoLang/cmd/error_logging_practice/service"
)

func main() {
	// ロガーの初期化
	log := logger.NewLogger(os.Stdout, logger.LevelInfo)

	ctx := context.Background()

	log.Info(ctx, "アプリケーション開始")

	// ユーザーサービスの作成
	userService := service.NewUserService(log)

	// 正常なケース
	user, err := userService.GetUser(ctx, 1)
	if err != nil {
		log.Error(ctx, "ユーザー取得エラー", "error", err, "user_id", 1)
		fmt.Printf("エラー: %v\n", err)
	} else {
		log.Info(ctx, "ユーザー取得成功", "user_id", user.ID, "name", user.Name)
		fmt.Printf("取得成功: %+v\n", user)
	}

	// エラーケース: ユーザーが見つかりません
	_, err = userService.GetUser(ctx, 999)
	if err != nil {
		log.Error(ctx, "ユーザー取得エラー", "error", err, "user_id", 999)
		fmt.Printf("エラー: %v\n", err)
	}

	// エラーケース: 無効なID
	_, err = userService.GetUser(ctx, -1)
	if err != nil {
		log.Error(ctx, "ユーザー取得エラー", "error", err, "user_id", -1)
		fmt.Printf("エラー: %v\n", err)
	}

	log.Info(ctx, "アプリケーション終了")
}
