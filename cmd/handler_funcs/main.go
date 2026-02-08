package main

import (
	"fmt"
	"log"
	"net/http"
)

// 複数のハンドラ関数を持つHTTPサーバーの例
// 異なるパスごとに異なる処理を実装します

// ホームページのハンドラ
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ようこそホームページへ！\n")
	fmt.Fprintf(w, "\n利用可能なページ:\n")
	fmt.Fprintf(w, "- / (このページ)\n")
	fmt.Fprintf(w, "- /about\n")
	fmt.Fprintf(w, "- /contact\n")
}

// Aboutページのハンドラ
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About ページ\n")
	fmt.Fprintf(w, "これは学習用のGoサーバーです。\n")
}

// Contactページのハンドラ
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact ページ\n")
	fmt.Fprintf(w, "お問い合わせ先: example@example.com\n")
}

func main() {
	// 各パスに対してハンドラ関数を登録
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	// サーバーを起動
	log.Println("サーバーを起動しています: http://localhost:8080")
	log.Println("試してみてください:")
	log.Println("  http://localhost:8080/")
	log.Println("  http://localhost:8080/about")
	log.Println("  http://localhost:8080/contact")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("サーバーエラー:", err)
	}
}
