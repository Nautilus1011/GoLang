# シンプルなエラー処理・ロギング実装

小規模プロジェクト向けの、標準ライブラリを使ったシンプルな実装例です。

## 📁 プロジェクト構成

```
cmd/simple_error_logging/
├── main.go        # メインプログラム
├── user.go        # ユーザー関連のロジック
├── user_test.go   # テストコード
└── README.md      # このファイル
```

**特徴:**
- ✅ パッケージ分割なし（すべて`main`パッケージ）
- ✅ 標準ライブラリのみ使用
- ✅ シンプルで理解しやすい

## 🎯 使用している標準ライブラリ

### 1. `log/slog` - 構造化ログ（Go 1.21+）

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

logger.Info("メッセージ", "key1", "value1", "key2", 123)
logger.Error("エラー", "error", err)
```

**特徴:**
- JSON形式で出力
- キー・バリュー形式
- ログレベル管理（Debug, Info, Warn, Error）

### 2. `errors` - エラー処理（標準パッケージ）

#### センチネルエラー（定義済みエラー）

```go
var (
    ErrNotFound   = errors.New("not found")
    ErrValidation = errors.New("validation error")
)
```

#### エラーのラップ（`%w`）

```go
// エラーをラップして詳細情報を追加
return fmt.Errorf("failed to fetch user: %w", err)
```

#### エラーのチェック（`errors.Is`）

```go
if errors.Is(err, ErrNotFound) {
    // NotFoundエラーの処理
}
```

#### 複数エラーの結合（`errors.Join`、Go 1.20+）

```go
err := errors.Join(ErrValidation, ErrEmptyName, ErrEmptyEmail)
```

## 🚀 実行方法

### プログラムの実行

```bash
cd cmd/simple_error_logging
go run .
```

**出力例:**
```json
{"time":"2026-02-12T10:30:00+09:00","level":"INFO","msg":"アプリケーション開始"}
{"time":"2026-02-12T10:30:00+09:00","level":"INFO","msg":"ユーザー取得成功","user_id":1,"name":"田中太郎"}
取得成功: &{ID:1 Name:田中太郎 Email:tanaka@example.com}
{"time":"2026-02-12T10:30:00+09:00","level":"WARN","msg":"ユーザーが見つかりません","user_id":999}
エラー: failed to fetch user: user_id=999: not found
```

### テストの実行

```bash
# すべてのテストを実行
go test -v

# カバレッジ付き
go test -cover

# 詳細なカバレッジ
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 📚 コードの解説

### エラー処理のパターン

#### 1. センチネルエラーの定義

```go
var (
    ErrNotFound = errors.New("not found")
    ErrValidation = errors.New("validation error")
)
```

**使い所:**
- よく使うエラーを定数のように定義
- `errors.Is`でチェック可能

#### 2. エラーのラップ

```go
func GetUser(id int) (*User, error) {
    user, err := fetchUserFromDB(id)
    if err != nil {
        // コンテキスト情報を追加してラップ
        return nil, fmt.Errorf("failed to fetch user: %w", err)
    }
    return user, nil
}
```

**利点:**
- エラーの発生箇所がわかる
- 元のエラーは`errors.Unwrap`で取得可能
- `errors.Is`でチェック可能

#### 3. エラーのチェック

```go
if errors.Is(err, ErrNotFound) {
    // NotFoundエラー固有の処理
}
```

**従来の方法との違い:**
```go
// ❌ 古い方法（文字列比較）
if err.Error() == "not found" { ... }

// ✅ 新しい方法（型で判定）
if errors.Is(err, ErrNotFound) { ... }
```

#### 4. 複数エラーの結合

```go
var errs []error
if name == "" {
    errs = append(errs, ErrEmptyName)
}
if email == "" {
    errs = append(errs, ErrEmptyEmail)
}
if len(errs) > 0 {
    return errors.Join(errs...)
}
```

**利点:**
- 複数のバリデーションエラーをまとめて返せる
- 各エラーを個別にチェック可能

### ロギングのパターン

#### 構造化ログ

```go
logger.Info("ユーザー取得成功", 
    "user_id", user.ID, 
    "name", user.Name,
    "email", user.Email)
```

**出力:**
```json
{
  "time": "2026-02-12T10:30:00+09:00",
  "level": "INFO",
  "msg": "ユーザー取得成功",
  "user_id": 1,
  "name": "田中太郎",
  "email": "tanaka@example.com"
}
```

#### ログレベルの使い分け

```go
logger.Debug("デバッグ情報")   // 開発時の詳細情報
logger.Info("情報")           // 通常の情報
logger.Warn("警告")           // 警告（処理は継続）
logger.Error("エラー")        // エラー（処理失敗）
```

## 🔄 error_logging_practiceとの比較

| 項目 | simple_error_logging | error_logging_practice |
|---|---|---|
| **構成** | 1パッケージ、3ファイル | 4パッケージ、7ファイル |
| **エラー** | 標準`errors`+センチネル | カスタムエラー型 |
| **ログ** | 標準`log/slog` | カスタムロガー |
| **学習難易度** | ⭐️⭐️ | ⭐️⭐️⭐️⭐️ |
| **実務での使用** | 小規模プロジェクト向け | 大規模プロジェクト向け |

## 💡 実務での使い方

### いつこの方法を使うか

✅ **使うべきケース:**
- 小規模プロジェクト（~1000行）
- プロトタイプ・PoC
- 社内ツール
- 学習の最初のステップ

❌ **避けるべきケース:**
- 大規模プロジェクト
- マイクロサービス
- 複雑なエラー分類が必要な場合

### 次のステップ

このシンプル版で基礎を理解したら：

1. **error_logging_practiceを学ぶ**
   - カスタムエラー型の設計
   - より高度なロギング

2. **外部ライブラリを試す**
   - `zap`, `zerolog` (高速ログ)
   - `cockroachdb/errors` (高度なエラー)

3. **パッケージ設計を学ぶ**
   - `internal/`の使い方
   - Clean Architectureなど

## 📖 参考リソース

- [Go Blog - Error handling and Go](https://go.dev/blog/error-handling-and-go)
- [Go Blog - Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [log/slog パッケージ](https://pkg.go.dev/log/slog)
- [errors パッケージ](https://pkg.go.dev/errors)

## 🎓 学習のポイント

### まずこのファイルを読む順番

1. **README.md**（このファイル）- 全体像を理解
2. **user.go** - エラー処理の実装を見る
3. **main.go** - 使い方を見る
4. **user_test.go** - テストの書き方を学ぶ

### 理解すべき重要な概念

- ✅ センチネルエラー
- ✅ エラーのラップ（`%w`）
- ✅ `errors.Is`によるエラーチェック
- ✅ `errors.Join`による複数エラー
- ✅ 構造化ログ
- ✅ コンテキストの受け渡し

---

**作成日**: 2026年2月12日  
**目的**: Go言語の基本的なエラー処理・ロギングを標準ライブラリで学ぶ
