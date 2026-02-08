package main

import (
	"fmt"
	"log"
	"net/http"
)

// クエリパラメータを扱うHTTPサーバーの例
// URLからパラメータを取得して処理します

// 挨拶ハンドラ - クエリパラメータ "name" を使用
// 例: http://localhost:8080/hello?name=太郎
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	name := r.URL.Query().Get("name")

	// パラメータが空の場合はデフォルト値を使用
	if name == "" {
		name = "ゲスト"
	}

	fmt.Fprintf(w, "こんにちは、%sさん！\n", name)
}

// 計算ハンドラ - 2つの数値を受け取って合計を返す
// 例: http://localhost:8080/add?a=5&b=3
func addHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	// パラメータが空かチェック
	if aStr == "" || bStr == "" {
		fmt.Fprintf(w, "エラー: パラメータ 'a' と 'b' が必要です\n")
		fmt.Fprintf(w, "例: /add?a=5&b=3\n")
		return
	}

	// 文字列を整数に変換
	var a, b int
	_, err1 := fmt.Sscanf(aStr, "%d", &a)
	_, err2 := fmt.Sscanf(bStr, "%d", &b)

	if err1 != nil || err2 != nil {
		fmt.Fprintf(w, "エラー: 数値を入力してください\n")
		return
	}

	// 計算して結果を返す
	result := a + b
	fmt.Fprintf(w, "%d + %d = %d\n", a, b, result)
}

// リクエスト情報を表示するハンドラ
func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "リクエスト情報:\n")
	fmt.Fprintf(w, "==================\n")
	fmt.Fprintf(w, "メソッド: %s\n", r.Method)
	fmt.Fprintf(w, "パス: %s\n", r.URL.Path)
	fmt.Fprintf(w, "クエリ文字列: %s\n", r.URL.RawQuery)
	fmt.Fprintf(w, "\nすべてのクエリパラメータ:\n")

	// すべてのクエリパラメータをループで表示
	for key, values := range r.URL.Query() {
		for _, value := range values {
			fmt.Fprintf(w, "  %s = %s\n", key, value)
		}
	}
}

func main() {
	// ハンドラを登録
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/info", infoHandler)

	// ホームページ
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "クエリパラメータのサンプル\n")
		fmt.Fprintf(w, "========================\n\n")
		fmt.Fprintf(w, "試してみてください:\n")
		fmt.Fprintf(w, "1. http://localhost:8080/hello?name=太郎\n")
		fmt.Fprintf(w, "2. http://localhost:8080/add?a=10&b=20\n")
		fmt.Fprintf(w, "3. http://localhost:8080/info?key1=value1&key2=value2\n")
	})

	// サーバーを起動
	log.Println("サーバーを起動しています: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("サーバーエラー:", err)
	}
}
