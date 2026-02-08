# CI/CDガイド - テストとチェックの解説

このドキュメントでは、CI/CDで実行される各テストの意味と重要性を説明します。

## 🔍 実行されるチェック一覧

### 1️⃣ テスト実行（`go test`）

```bash
make test
# または
go test -v ./...
```

**何をチェック？**
- ユニットテストを実行
- コードが期待通りに動作するか確認

**重要性**
- ✅ バグを早期発見
- ✅ リファクタリング時の安全性確保
- ✅ 仕様の明確化

**学習ポイント**
```go
// テストの例
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("期待値: 5, 実際: %d", result)
    }
}
```

---

### 2️⃣ ビルド確認（`go build`）

```bash
make build
# または
go build -v ./...
```

**何をチェック？**
- コードがコンパイル可能か確認
- 実行ファイルを作成できるか検証

**重要性**
- ✅ 構文エラーの検出
- ✅ 型エラーの検出
- ✅ インポートの問題を発見

**学習ポイント**
- Goはコンパイル言語なので、実行前にビルドが必要
- ビルド時にエラーが出れば、実行できない
- コンパイラが多くのエラーを事前にキャッチ

**よくあるエラー例**
```
# 型エラー
cannot use "hello" (type string) as type int

# 未使用の変数
declared and not used

# インポート漏れ
undefined: fmt
```

---

### 3️⃣ フォーマットチェック（`gofmt`）

```bash
make fmt
# または
gofmt -s -w .
```

**何をチェック？**
- コードが標準フォーマットに従っているか
- インデント、空白、改行などのスタイル

**重要性**
- ✅ チーム全体で統一されたコードスタイル
- ✅ レビュー時にスタイルの議論を避ける
- ✅ 可読性の向上

**学習ポイント**
- Goには公式のフォーマッター（`gofmt`）がある
- コミット前に必ず実行すべき
- VSCodeなどのエディタで保存時に自動実行可能

**フォーマット前後の例**
```go
// フォーマット前
func main(){
x:=1+2
fmt.Println(x)}

// フォーマット後
func main() {
    x := 1 + 2
    fmt.Println(x)
}
```

---

### 4️⃣ 静的解析（`go vet`）

```bash
make vet
# または
go vet ./...
```

**何をチェック？**
- コードの潜在的な問題を検出
- よくあるミスやバグの可能性を指摘

**重要性**
- ✅ 実行時エラーを事前に防ぐ
- ✅ よくあるミスを学べる
- ✅ コードの品質向上

**学習ポイント**
検出される問題の例：
- `Printf`のフォーマット指定子の誤り
- 構造体のコピーミス
- 到達不可能なコード
- アンロック忘れ

**検出される問題の例**
```go
// NG: フォーマット指定子の誤り
name := "太郎"
fmt.Printf("%d", name)  // %sが正しい

// NG: エラーチェック漏れの可能性
file, _ := os.Open("test.txt")  // エラーを無視している

// OK: 適切なエラーハンドリング
file, err := os.Open("test.txt")
if err != nil {
    log.Fatal(err)
}
```

---

### 5️⃣ Lint（`golangci-lint`）

```bash
make lint
# または
golangci-lint run
```

**何をチェック？**
- コードの品質
- ベストプラクティスへの準拠
- 複数のリンターを統合実行

**重要性**
- ✅ より高度なコード品質チェック
- ✅ Goらしい書き方を学べる
- ✅ セキュリティ問題の検出

**学習ポイント**
検出される問題の例：
- エラーチェック漏れ（`errcheck`）
- 未使用のコード（`unused`）
- コードの複雑さ（`gocyclo`）
- セキュリティリスク（`gosec`）

**検出される問題の例**
```go
// errcheck: エラーチェック漏れ
file.Close()  // エラーを無視している

// 改善版
if err := file.Close(); err != nil {
    log.Printf("close error: %v", err)
}

// unused: 未使用の変数
func example() {
    x := 10  // 使われていない
    y := 20
    return y
}
```

---

## 🚀 使い方

### ローカルで実行

```bash
# すべてのチェックを実行
make check

# 個別に実行
make test    # テスト
make build   # ビルド確認
make fmt     # フォーマット
make vet     # 静的解析
make lint    # Lint（golangci-lintが必要）
```

### GitHub Actionsで自動実行

プッシュやプルリクエスト時に自動で実行されます。

```yaml
# .github/workflows/ci.yml
jobs:
  test:
    steps:
      - name: Run tests        # 1. テスト
      - name: Build            # 2. ビルド
      - name: Check formatting # 3. フォーマット
      - name: Run go vet       # 4. 静的解析
      - name: Run golangci-lint # 5. Lint
```

---

## 📝 チェックリスト

コミット前に確認：

- [ ] `make test` - すべてのテストが通る
- [ ] `make build` - ビルドが成功する
- [ ] `make fmt` - フォーマット済み
- [ ] `make vet` - 警告なし
- [ ] `make lint` - 重大な問題なし（オプション）

または一括で：
```bash
make check
```

---

## 🎓 学習のポイント

### 初心者の方へ

1. **まず `make test` と `make fmt` から始める**
   - この2つが最も重要
   - 習慣化すると良い

2. **エラーメッセージを読む**
   - どのチェックでエラーが出たか確認
   - エラーメッセージから学ぶ

3. **エディタと連携**
   - VSCodeなどで保存時に自動フォーマット
   - リアルタイムでエラー表示

### 各チェックの優先度

| チェック | 優先度 | 必須度 |
|---------|--------|--------|
| test | 🔴 高 | 必須 |
| build | 🔴 高 | 必須 |
| fmt | 🟠 中 | 推奨 |
| vet | 🟠 中 | 推奨 |
| lint | 🟢 低 | オプション |

---

## ❓ よくある質問

### Q1: Lintで警告が多すぎる場合は？

**A:** 学習段階では警告を全て修正する必要はありません。
- まずはエラー（error）を修正
- 警告（warning）は徐々に対応
- CI/CDでは`continue-on-error: true`で失敗しない設定

### Q2: `go vet`と`golangci-lint`の違いは？

**A:**
- `go vet`: Go公式の基本的な静的解析ツール
- `golangci-lint`: 複数のリンターを統合した高機能ツール
- `golangci-lint`は`go vet`を含む

### Q3: フォーマットチェックで失敗したら？

**A:** `make fmt`を実行するだけで自動修正されます。

---

## 🔗 参考リンク

- [Testing in Go](https://go.dev/doc/tutorial/add-a-test)
- [go vet](https://pkg.go.dev/cmd/vet)
- [gofmt](https://pkg.go.dev/cmd/gofmt)
- [golangci-lint](https://golangci-lint.run/)

---

**これらのチェックを習慣化することで、高品質なGoコードを書けるようになります！** 🎉
