# エラー処理・ロギング・テスト設計 実践プロジェクト

このプロジェクトは、Go言語における実務で重視される以下の技術を学ぶためのサンプル実装です：

- **エラー処理**: カスタムエラー型、エラーラップ、エラー判定
- **ロギング**: 構造化ログ、ログレベル、JSON形式出力
- **テスト設計**: テーブル駆動テスト、モックデータ、テストカバレッジ

## 📁 プロジェクト構成

```
cmd/error_logging_practice/
├── main.go                     # エントリーポイント
├── errors/                     # カスタムエラーパッケージ
│   ├── errors.go              # エラー型定義
│   └── errors_test.go         # エラーのテスト
├── logger/                     # ロギングパッケージ
│   ├── logger.go              # 構造化ロガー実装
│   └── logger_test.go         # ロガーのテスト
├── service/                    # ビジネスロジック
│   ├── user_service.go        # ユーザーサービス
│   └── user_service_test.go   # サービスのテスト
└── README.md                   # このファイル
```

## 🎯 学べる実務スキル

### 1. エラー処理のベストプラクティス

**特徴:**
- カスタムエラー型による型安全なエラー処理
- エラーコードによる分類（NOT_FOUND, INVALID_INPUT等）
- `errors.Is`, `errors.As`を使った標準的なエラーチェック
- エラーコンテキスト情報の付加

**実装例:**
```go
// エラーの作成
err := errors.NotFound("ユーザーが見つかりません").
    WithDetail("user_id", 123)

// エラーのチェック
if errors.IsNotFound(err) {
    // NotFoundに特化した処理
}
```

### 2. 構造化ロギング

**特徴:**
- JSON形式での構造化ログ出力
- ログレベル管理（DEBUG, INFO, WARN, ERROR）
- キー・バリュー形式でのメタデータ追加
- タイムスタンプの自動付加

**実装例:**
```go
logger.Info(ctx, "ユーザー取得成功", 
    "user_id", user.ID, 
    "name", user.Name)

// 出力例:
// {"timestamp":"2026-02-12T10:30:45+09:00","level":"INFO","message":"ユーザー取得成功","fields":{"user_id":1,"name":"田中太郎"}}
```

### 3. テスト設計

**特徴:**
- テーブル駆動テスト（Table-Driven Tests）
- 正常系・異常系の網羅的なテスト
- エラー型の検証
- サブテストによる構造化

**実装例:**
```go
tests := []struct {
    name        string
    id          int
    expectError bool
    errorType   func(error) bool
}{
    {"正常なケース", 1, false, nil},
    {"NotFoundエラー", 999, true, errors.IsNotFound},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // テスト実行
    })
}
```

## 🚀 実行方法

### プログラムの実行

```bash
cd cmd/error_logging_practice
go run main.go
```

**出力例:**
```json
{"timestamp":"2026-02-12T10:30:00+09:00","level":"INFO","message":"アプリケーション開始","fields":{}}
{"timestamp":"2026-02-12T10:30:00+09:00","level":"DEBUG","message":"ユーザー取得完了","fields":{"user_id":1}}
{"timestamp":"2026-02-12T10:30:00+09:00","level":"INFO","message":"ユーザー取得成功","fields":{"user_id":1,"name":"田中太郎"}}
取得成功: &{ID:1 Name:田中太郎 Email:tanaka@example.com}
{"timestamp":"2026-02-12T10:30:00+09:00","level":"ERROR","message":"ユーザー取得エラー","fields":{"error":"[NOT_FOUND] ユーザーID 999 が見つかりません","user_id":999}}
エラー: [NOT_FOUND] ユーザーID 999 が見つかりません
```

### テストの実行

**全テスト実行:**
```bash
cd cmd/error_logging_practice
go test ./...
```

**詳細出力:**
```bash
go test -v ./...
```

**カバレッジ確認:**
```bash
go test -cover ./...
```

**カバレッジ詳細（HTML）:**
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 📚 各パッケージの解説

### errors パッケージ

カスタムエラー型を定義し、エラーコードと詳細情報を付加できます。

**主な型・関数:**
- `AppError`: カスタムエラー型
- `New()`: 新しいエラーを作成
- `Wrap()`: 既存のエラーをラップ
- `NotFound()`, `InvalidInput()`: エラーコンストラクタ
- `IsNotFound()`, `IsInvalidInput()`: エラー型判定

### logger パッケージ

構造化ログを出力するシンプルなロガー実装です。

**主な型・関数:**
- `Logger`: ロガー本体
- `NewLogger()`: ロガーの作成
- `Debug()`, `Info()`, `Warn()`, `Error()`: ログ出力
- `Level`: ログレベル（DEBUG, INFO, WARN, ERROR）

### service パッケージ

ビジネスロジックを実装し、エラー処理とロギングを統合します。

**主な型・関数:**
- `UserService`: ユーザー関連の処理
- `GetUser()`: ユーザー取得
- `CreateUser()`: ユーザー作成

## 💡 実務で活かせるポイント

### 1. エラーハンドリングの一貫性

実務では、エラーの種類を適切に分類し、一貫した方法で処理することが重要です。
このプロジェクトでは：
- エラーコードによる分類
- エラーラップによるコンテキスト保持
- 型安全なエラーチェック

を実装しています。

### 2. デバッグ可能なログ

本番環境でのトラブルシューティングには、適切なログが不可欠です。
このプロジェクトでは：
- JSON形式で機械可読
- 構造化されたメタデータ
- ログレベルによるフィルタリング

を実現しています。

### 3. テスト駆動開発（TDD）

品質の高いコードを維持するには、包括的なテストが必要です。
このプロジェクトでは：
- テーブル駆動テストによる効率化
- 正常系・異常系の網羅
- エラー型の厳密な検証

を示しています。

## 🎓 学習の進め方

1. **コードを読む**: 各パッケージの実装を読み、構造を理解する
2. **テストを実行**: テストを実行し、どのようなケースをカバーしているか確認
3. **コードを改変**: 新しい機能やテストケースを追加してみる
4. **実験**: エラーケースを増やしたり、ログレベルを変更してみる

## 🔧 カスタマイズ例

### 新しいエラーコードの追加

```go
const (
    CodeUnauthorized ErrorCode = "UNAUTHORIZED"
    CodeForbidden    ErrorCode = "FORBIDDEN"
)

func Unauthorized(message string) *AppError {
    return New(CodeUnauthorized, message)
}
```

### ログの出力先を変更

```go
// ファイルへの出力
f, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
logger := logger.NewLogger(f, logger.LevelInfo)
```

### 新しいサービスメソッドの追加

```go
func (s *UserService) UpdateUser(ctx context.Context, id int, name string) error {
    if id <= 0 {
        return errors.InvalidInput("無効なユーザーID")
    }
    // 実装...
    s.logger.Info(ctx, "ユーザー更新完了", "user_id", id)
    return nil
}
```

## 📖 参考リソース

- [Effective Go - Error handling](https://go.dev/doc/effective_go#errors)
- [Go Blog - Error handling and Go](https://go.dev/blog/error-handling-and-go)
- [Go Blog - Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [Go Testing Package](https://pkg.go.dev/testing)

## 🎯 インターンでアピールできるポイント

このプロジェクトで学んだことをインターンで活かすには：

1. **エラーハンドリングの重要性を理解している**
   - 適切なエラー分類と処理ができる
   - デバッグしやすいエラーメッセージを書ける

2. **運用を意識したコーディング**
   - 構造化ログで障害調査をサポート
   - ログレベルで環境に応じた出力制御

3. **テスタブルなコード設計**
   - テストしやすい設計を心がける
   - カバレッジを意識した開発

4. **実務で即戦力になる技術**
   - Go標準のエラーハンドリングパターン
   - プロダクションレベルのログ設計
   - 保守性の高いテストコード

---

**作成日**: 2026年2月12日  
**目的**: Go言語における実務寄りのエラー処理・ロギング・テスト設計の習得
