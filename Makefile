.PHONY: help test fmt vet build clean run-simple run-handler run-query

# デフォルトターゲット
help:
	@echo "利用可能なコマンド:"
	@echo "  make test        - テストを実行"
	@echo "  make fmt         - コードをフォーマット"
	@echo "  make vet         - go vetを実行"
	@echo "  make build       - すべてのコマンドをビルド"
	@echo "  make clean       - ビルド成果物を削除"
	@echo "  make run-simple  - simple_serverを起動"
	@echo "  make run-handler - handler_funcsを起動"
	@echo "  make run-query   - query_paramsを起動"

# テスト
test:
	go test -v ./...

# フォーマット
fmt:
	gofmt -s -w .

# 静的解析
vet:
	go vet ./...

# ビルド
build:
	@echo "ビルド中..."
	@mkdir -p bin
	@go build -o bin/simple_server ./cmd/simple_server
	@go build -o bin/handler_funcs ./cmd/handler_funcs
	@go build -o bin/query_params ./cmd/query_params
	@echo "完了: bin/"

# クリーンアップ
clean:
	rm -rf bin/
	go clean

# サーバー起動
run-simple:
	go run ./cmd/simple_server/main.go

run-handler:
	go run ./cmd/handler_funcs/main.go

run-query:
	go run ./cmd/query_params/main.go
