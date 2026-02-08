# 開発ガイド

このドキュメントでは、GoLang リポジトリでの開発方法とCI/CD環境について説明します。

## 目次

- [環境セットアップ](#環境セットアップ)
- [開発ワークフロー](#開発ワークフロー)
- [テスト](#テスト)
- [コード品質](#コード品質)
- [CI/CD](#cicd)
- [Makefile コマンド](#makefile-コマンド)

---

## 環境セットアップ

### 必要なツール

- Go 1.21 以上
- make
- git

### 開発ツールのインストール

プロジェクトで使用する開発ツールをインストールします：

```bash
make install-tools
```

これにより以下がインストールされます：
- `golangci-lint` - 包括的なリンター
- `goimports` - import文の自動整理

### 依存関係のインストール

```bash
go mod download
```

---

## 開発ワークフロー

### 1. ブランチの作成

```bash
git checkout -b feature/your-feature-name
```

### 2. コードの作成

コードを書く前にテストを書くことを推奨します（TDD）。

### 3. フォーマットとLint

コードを書いたら、フォーマットとリンターを実行：

```bash
make fmt    # コードをフォーマット
make lint   # リンターを実行
```

### 4. テストの実行

```bash
make test           # 通常のテスト
make test-verbose   # 詳細表示
make test-coverage  # カバレッジレポート生成
```

### 5. ビルドの確認

```bash
make build  # すべてのコマンドをビルド
```

### 6. コミットとプッシュ

```bash
git add .
git commit -m "feat: 新機能の追加"
git push origin feature/your-feature-name
```

---

## テスト

### テストの書き方

テストファイルは `*_test.go` という名前で作成します。

例：`foundation_test.go`

```go
package mylib

import "testing"

func TestAdd(t *testing.T) {
    result := Add(3, 4)
    expected := 7
    
    if result != expected {
        t.Errorf("Add(3, 4) = %d; want %d", result, expected)
    }
}
```

### テーブル駆動テスト

複数のケースをテストする場合は、テーブル駆動テストを使用：

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        x, y     int
        expected int
    }{
        {"正の数", 3, 4, 7},
        {"負の数", -1, 1, 0},
        {"ゼロ", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.x, tt.y)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}
```

### ベンチマークテスト

パフォーマンスを測定したい場合：

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(3, 4)
    }
}
```

実行方法：

```bash
go test -bench=. -benchmem
```

### カバレッジレポート

```bash
make test-coverage
```

これで `coverage.html` が生成されます。ブラウザで開いて確認できます。

---

## コード品質

### golangci-lint

包括的なリンターで、複数のリンターを統合しています。

```bash
make lint
```

設定は [.golangci.yml](.golangci.yml) で管理されています。

有効なリンター：
- `errcheck` - エラーチェック漏れ
- `gosimple` - コードの簡略化
- `govet` - 静的解析
- `gofmt` - フォーマット
- `gosec` - セキュリティ
- その他多数

### コードフォーマット

```bash
make fmt
```

これにより：
- `gofmt -s` でフォーマット
- `goimports` で import文を整理

### go vet

静的解析ツール：

```bash
make vet
```

---

## CI/CD

### GitHub Actions

プッシュまたはプルリクエスト時に自動的に以下が実行されます：

1. **Test** - 複数のGoバージョンでテスト
   - Go 1.21, 1.22, 1.23
   - レースディテクタ有効
   - カバレッジレポート生成

2. **Lint** - golangci-lint でコード品質チェック

3. **Build** - すべてのコマンドのビルド確認

4. **Format Check** - コードフォーマットの確認

### ワークフロー設定

[.github/workflows/ci.yml](.github/workflows/ci.yml) で設定されています。

### ローカルでCIと同じチェックを実行

```bash
make ci
```

これで以下が実行されます：
- フォーマット
- go vet
- golangci-lint
- テスト

---

## Makefile コマンド

すべての利用可能なコマンドを表示：

```bash
make help
```

### よく使うコマンド

| コマンド | 説明 |
|---------|------|
| `make test` | テストを実行 |
| `make test-coverage` | カバレッジレポート付きテスト |
| `make lint` | golangci-lint 実行 |
| `make fmt` | コードをフォーマット |
| `make build` | すべてのコマンドをビルド |
| `make clean` | ビルド成果物を削除 |
| `make ci` | CI環境と同じチェックを実行 |
| `make run-simple` | simple_server を起動 |
| `make run-handler` | handler_funcs を起動 |
| `make run-query` | query_params を起動 |

---

## プロジェクト構成

```
.
├── .github/
│   └── workflows/
│       └── ci.yml           # GitHub Actions CI設定
├── .golangci.yml            # golangci-lint 設定
├── Makefile                 # タスク定義
├── go.mod                   # Go モジュール定義
├── main.go                  # メインプログラム
├── cmd/                     # 実行可能プログラム
│   ├── simple_server/       # シンプルなHTTPサーバー
│   ├── handler_funcs/       # ハンドラー関数の例
│   └── query_params/        # クエリパラメータの例
├── http_server/             # 高度なHTTPサーバー
│   └── http_server.go
├── mylib/                   # ライブラリパッケージ
│   ├── animal.go
│   ├── animal_test.go       # テスト
│   ├── foundation.go
│   └── foundation_test.go   # テスト
└── DEVELOPMENT.md           # このファイル
```

---

## ベストプラクティス

### 1. テスト駆動開発（TDD）

1. テストを書く
2. テストを実行（失敗することを確認）
3. 実装を書く
4. テストを実行（成功することを確認）
5. リファクタリング

### 2. コミットメッセージ

Conventional Commits形式を推奨：

- `feat:` - 新機能
- `fix:` - バグ修正
- `docs:` - ドキュメント
- `test:` - テストの追加・修正
- `refactor:` - リファクタリング
- `chore:` - その他の変更

例：
```
feat: クエリパラメータ処理を追加
fix: エラーハンドリングのバグを修正
docs: READMEを更新
```

### 3. プルリクエスト前のチェックリスト

- [ ] `make fmt` でフォーマット済み
- [ ] `make lint` で警告なし
- [ ] `make test` ですべてのテストが通る
- [ ] 新しいコードにテストを追加
- [ ] ドキュメントを更新（必要な場合）

### 4. コードレビュー

- 読みやすさを重視
- テストがあることを確認
- エラーハンドリングが適切か確認

---

## トラブルシューティング

### golangci-lint のインストールに失敗する

```bash
# 手動インストール
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### テストが遅い

キャッシュを利用：

```bash
go test -count=1 ./...  # キャッシュを無視
go clean -testcache     # テストキャッシュをクリア
```

### ポート8080が使用中

サーバーのポートを変更：

```go
// :8080 を :8081 などに変更
err := http.ListenAndServe(":8081", nil)
```

---

## 参考リンク

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [golangci-lint](https://golangci-lint.run/)
- [Testing in Go](https://go.dev/doc/tutorial/add-a-test)

---

## 質問やフィードバック

問題や提案がある場合は、Issueを作成してください。
