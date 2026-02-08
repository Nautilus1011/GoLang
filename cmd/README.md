# 基礎的なHTTPサーバー学習プログラム

このディレクトリには、Goで基本的なHTTPサーバーを学ぶための3つの段階的なプログラムが含まれています。

## ディレクトリ構成

```
cmd/
├── simple_server/    # 01. 最もシンプルなサーバー
│   └── main.go
├── handler_funcs/    # 02. 複数のハンドラを持つサーバー
│   └── main.go
└── query_params/     # 03. クエリパラメータを扱うサーバー
    └── main.go
```

## 学習の順序

### 1. simple_server - 最もシンプルなサーバー

**学べること:**
- `http.HandleFunc()` の基本的な使い方
- `http.ListenAndServe()` でサーバーを起動する方法
- レスポンスを返す基本的な方法

**実行方法:**
```bash
go run cmd/simple_server/main.go
```

**アクセス方法:**
```bash
# ブラウザで開く
open http://localhost:8080

# またはcurlコマンドで
curl http://localhost:8080
```

---

### 2. handler_funcs - 複数のハンドラを持つサーバー

**学べること:**
- 複数のパスに対してハンドラを登録する方法
- ハンドラ関数を分離して整理する方法
- 異なるページを提供する方法

**実行方法:**
```bash
go run cmd/handler_funcs/main.go
```

**アクセス方法:**
```bash
curl http://localhost:8080/
curl http://localhost:8080/about
curl http://localhost:8080/contact
```

---

### 3. query_params - クエリパラメータを扱うサーバー

**学べること:**
- URLクエリパラメータの取得方法 (`r.URL.Query().Get()`)
- パラメータの検証とデフォルト値の設定
- リクエスト情報の詳細な取得方法

**実行方法:**
```bash
go run cmd/query_params/main.go
```

**アクセス方法:**
```bash
curl "http://localhost:8080/hello?name=太郎"
curl "http://localhost:8080/add?a=10&b=20"
curl "http://localhost:8080/info?key1=value1&key2=value2"
```

## Tips

- サーバーを停止するには、ターミナルで `Ctrl + C` を押してください
- ポート8080が既に使用されている場合は、コード内の `:8080` を `:8081` などに変更してください
- 各プログラムは独立しているので、好きな順番で試せます（順序通りを推奨）

## より高度な機能を学びたい場合

基礎的なサーバーに慣れたら、`http_server/http_server.go` を見てみてください。
以下のような高度な機能が実装されています:
- JSONの送受信
- ミドルウェア
- HTMLテンプレート
- グレースフルシャットダウン
- 静的ファイルの配信

## プロジェクト構成について

この `cmd/` ディレクトリの構成は、Goのベストプラクティスに従っています：
- 各サブディレクトリが1つの実行可能プログラムを表す
- `go run cmd/<プログラム名>/main.go` で個別に実行できる
- コードの整理と管理がしやすい
