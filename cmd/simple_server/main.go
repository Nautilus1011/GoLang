package main

import (
	"fmt"
	"log"
	"net/http"
)

// 最もシンプルなHTTPサーバーの例
// このプログラムは http://localhost:8080 でアクセスできるサーバーを起動します

func main() {
	// ハンドラ関数を登録
	// "/" にアクセスしたときの処理を定義
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// レスポンスとして文字列を書き込む
		fmt.Fprintf(w, "こんにちは、世界！\n")
		fmt.Fprintf(w, "これは最もシンプルなGoのHTTPサーバーです。\n")
	})

	// サーバーを起動
	// ポート8080で待ち受ける
	log.Println("サーバーを起動しています: http://localhost:8080")

	// ListenAndServe はブロッキング関数で、エラーが発生するまで実行し続けます
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("サーバーエラー:", err)
	}
}
