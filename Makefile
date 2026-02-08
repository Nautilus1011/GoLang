.PHONY: help test test-verbose test-coverage lint fmt vet build build-all clean run-simple run-handler run-query install-tools

# デフォルトターゲット
help:
	@echo "利用可能なコマンド:"
	@echo "  make test          - テストを実行"
	@echo "  make test-verbose  - 詳細出力でテストを実行"
	@echo "  make test-coverage - カバレッジレポート付きでテストを実行"
	@echo "  make lint          - golangci-lintを実行"
	@echo "  make fmt           - コードをフォーマット"
	@echo "  make vet           - go vetを実行"
	@echo "  make build         - すべてのコマンドをビルド"
	@echo "  make build-all     - すべてのパッケージをビルド"
	@echo "  make clean         - ビルド成果物を削除"
	@echo "  make run-simple    - simple_serverを起動"
	@echo "  make run-handler   - handler_funcsを起動"
	@echo "  make run-query     - query_paramsを起動"
	@echo "  make install-tools - 開発ツールをインストール"

# テスト関連
test:
	go test ./...

test-verbose:
	go test -v ./...

test-coverage:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "カバレッジレポート: coverage.html"

# Lint関連
lint:
	golangci-lint run --timeout=5m

fmt:
	gofmt -s -w .
	goimports -w .

vet:
	go vet ./...

# ビルド関連
build:
	@echo "simple_serverをビルド中..."
	@go build -o bin/simple_server ./cmd/simple_server
	@echo "handler_funcsをビルド中..."
	@go build -o bin/handler_funcs ./cmd/handler_funcs
	@echo "query_paramsをビルド中..."
	@go build -o bin/query_params ./cmd/query_params
	@echo "ビルド完了: bin/"

build-all:
	go build -v ./...

clean:
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean

# サーバー起動
run-simple:
	go run ./cmd/simple_server/main.go

run-handler:
	go run ./cmd/handler_funcs/main.go

run-query:
	go run ./cmd/query_params/main.go

# 開発ツールのインストール
install-tools:
	@echo "開発ツールをインストール中..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	@echo "インストール完了"

# 依存関係の管理
mod-download:
	go mod download

mod-tidy:
	go mod tidy

mod-verify:
	go mod verify

# CI環境で実行されるすべてのチェック
ci: fmt vet lint test

# すべてをチェックしてビルド
all: ci build
