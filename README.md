# GoLang 学習リポジトリ

Goプログラミング言語の学習用リポジトリです。基礎から応用まで、段階的に学習できる構成になっています。

## 📁 プロジェクト構成

```
.
├── cmd/                    # 実行可能プログラム
│   ├── simple_server/      # 01. 最もシンプルなHTTPサーバー
│   ├── handler_funcs/      # 02. 複数ハンドラを持つサーバー
│   └── query_params/       # 03. クエリパラメータを扱うサーバー
├── http_server/            # 高度なHTTPサーバー（中級者向け）
├── mylib/                  # ライブラリパッケージ
├── .github/workflows/      # CI/CD設定
│   └── ci.yml              # GitHub Actions
├── .golangci.yml           # Linter設定
├── Makefile                # タスクランナー
├── DEVELOPMENT.md          # 開発ガイド（詳細）
└── README.md               # このファイル
```

## 🚀 クイックスタート

### 基本的なサーバーを起動

```bash
# 最もシンプルなサーバー
go run cmd/simple_server/main.go

# 複数のハンドラを持つサーバー
go run cmd/handler_funcs/main.go

# クエリパラメータを扱うサーバー
go run cmd/query_params/main.go
```

### テストを実行

```bash
make test               # 全テストを実行
make test-verbose       # 詳細出力
make test-coverage      # カバレッジレポート生成
```

### コード品質チェック

```bash
make fmt    # コードをフォーマット
make lint   # リンターを実行
make vet    # 静的解析
```

## 📚 学習順序

### 1. 基礎（`mylib/`）
- `foundation.go` - 基本的な関数の書き方
- `animal.go` - インターフェースとポリモーフィズム
- `slice.go` - スライス操作

### 2. HTTPサーバー（`cmd/`）
1. **simple_server** - まず最もシンプルなサーバーから
2. **handler_funcs** - 複数のエンドポイント
3. **query_params** - パラメータの扱い方

### 3. 応用（`http_server/`）
- JSON処理
- ミドルウェア
- HTMLテンプレート
- グレースフルシャットダウン

詳細は各ディレクトリの README.md を参照してください。

## 🛠️ 開発環境

### 必要なツール

- Go 1.21 以上
- make
- git

### 開発ツールのインストール

```bash
make install-tools
```

これにより以下がインストールされます：
- golangci-lint（包括的なリンター）
- goimports（import文の自動整理）

### 依存関係

```bash
go mod download    # 依存パッケージをダウンロード
go mod tidy       # go.modファイルを整理
```

## 📋 Makefileコマンド

すべての利用可能なコマンドを表示：

```bash
make help
```

### よく使うコマンド

| コマンド | 説明 |
|---------|------|
| `make test` | テストを実行 |
| `make test-coverage` | カバレッジレポート付きテスト |
| `make lint` | golangci-lint を実行 |
| `make fmt` | コードをフォーマット |
| `make build` | すべてのコマンドをビルド |
| `make clean` | ビルド成果物を削除 |
| `make ci` | CI環境と同じチェックを実行 |
| `make run-simple` | simple_server を起動 |

## 🧪 テスト

### テストの実行

```bash
# すべてのテスト
make test

# 特定のパッケージ
go test ./mylib/...

# ベンチマーク
go test -bench=. ./mylib/...
```

### カバレッジレポート

```bash
make test-coverage
# coverage.html がブラウザで開けます
```

## 🔄 CI/CD

GitHub Actions により、以下が自動実行されます：

- ✅ **テスト** - 複数のGoバージョン（1.21, 1.22, 1.23）でテスト
- ✅ **Lint** - golangci-lint によるコード品質チェック
- ✅ **Build** - すべてのコマンドのビルド確認
- ✅ **Format Check** - コードフォーマットの確認

### ローカルでCIと同じチェックを実行

```bash
make ci
```

## 📖 詳細なドキュメント

より詳しい開発ガイドは以下を参照：

- **[DEVELOPMENT.md](DEVELOPMENT.md)** - 開発ガイド
  - テストの書き方
  - コード品質のベストプラクティス
  - CI/CD詳細
  - トラブルシューティング

- **[cmd/README.md](cmd/README.md)** - HTTPサーバーの学習ガイド

## 🤝 開発ワークフロー

1. ブランチを作成
   ```bash
   git checkout -b feature/your-feature
   ```

2. コードを書く（テスト駆動開発を推奨）
   ```bash
   # 1. テストを書く
   # 2. テストを実行（失敗を確認）
   make test
   # 3. 実装を書く
   # 4. テストを実行（成功を確認）
   make test
   ```

3. フォーマットとLint
   ```bash
   make fmt
   make lint
   ```

4. コミット
   ```bash
   git add .
   git commit -m "feat: 新機能の説明"
   git push origin feature/your-feature
   ```

## 🎯 学習のヒント

### 初心者の方へ

1. まず `cmd/simple_server` から始めてください
2. コードにたくさんコメントがあるので、読みながら理解しましょう
3. 実際に動かしてみることが大切です
4. テストコードも学習教材として活用してください

### テスト駆動開発（TDD）

このリポジトリではテストが充実しています：
- `mylib/foundation_test.go` - 基本的なテストの書き方
- `mylib/animal_test.go` - インターフェースのテスト

テストを見ながら、どのようにコードをテストするか学べます。

## 🔧 トラブルシューティング

### ポート8080が既に使用中

サーバーのポートを変更してください：
```go
err := http.ListenAndServe(":8081", nil)  // 8080 → 8081
```

### golangci-lint が見つからない

```bash
make install-tools
```

### テストキャッシュをクリア

```bash
go clean -testcache
```

## 📚 参考リンク

- [A Tour of Go](https://go.dev/tour/) - Go言語公式チュートリアル
- [Effective Go](https://go.dev/doc/effective_go) - Goのベストプラクティス
- [Go by Example](https://gobyexample.com/) - 実例で学ぶGo

## 📝 ライセンス

学習用リポジトリのため、自由に使用・改変してください。

---

**Happy Coding! 🎉**
