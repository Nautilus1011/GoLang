.PHONY: help test fmt vet build clean run-simple run-handler run-query check lint

# デフォルトターゲット
help:
	@echo "利用可能なコマンド:"
	@echo "  make test        - テストを実行"
	@echo "  make build       - ビルド確認"
	@echo "  make fmt         - コードをフォーマット"
	@echo "  make vet         - 静的解析を実行"
	@echo "  make lint        - Lintを実行（golangci-lint）"
	@echo "  make check       - すべてのチェックを実行"
	@echo "  make clean       - ビルド成果物を削除"
	@echo "  make run-simple  - simple_serverを起動"
	@echo "  make run-handler - handler_funcsを起動"
	@echo "  make run-query   - query_paramsを起動"

# 1. テスト実行
test:
	@echo "テストを実行中..."
	go test -v ./...

# 2. ビルド確認
build:
	@echo "ビルド確認中..."
	go build -v ./...
	@echo "✓ ビルド成功"

# 3. フォーマット
fmt:
	@echo "コードをフォーマット中..."
	gofmt -s -w .
	@echo "✓ フォーマット完了"

# 4. 静的解析
vet:
	@echo "静的解析中..."
	go vet ./...
	@echo "✓ go vet完了"

# 5. Lint
lint:
	@echo "Lint実行中..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run --timeout=3m; \
		echo "✓ Lint完了"; \
	else \
		echo "golangci-lintがインストールされていません"; \
		echo "インストール: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# すべてのチェックを実行（CI相当）
check: test build fmt vet
	@echo ""
	@echo "✓ すべてのチェックが完了しました！"

# クリーンアップ
clean:
	@echo "クリーンアップ中..."
	rm -rf bin/
	go clean
	@echo "✓ クリーンアップ完了"

# サーバー起動
run-simple:
	go run ./cmd/simple_server/main.go

run-handler:
	go run ./cmd/handler_funcs/main.go

run-query:
	go run ./cmd/query_params/main.go
