.PHONY: help setup dev build run migrate clean

help:
	@echo "利用可能なコマンド:"
	@echo "  make setup    - プロジェクトのセットアップ（依存関係のインストール）"
	@echo "  make dev      - 開発サーバーの起動"
	@echo "  make build    - プロジェクトのビルド"
	@echo "  make run      - ビルド済みサーバーの起動"
	@echo "  make migrate  - データベースマイグレーションの実行"
	@echo "  make clean    - ビルド成果物の削除"

setup:
	@echo "依存関係をインストール中..."
	go mod download
	@echo "セットアップ完了！.envファイルを作成してください。"

dev:
	@echo "開発サーバーを起動中..."
	go run cmd/api/main.go

build:
	@echo "ビルド中..."
	go build -o bin/api cmd/api/main.go
	@echo "ビルド完了: bin/api"

run:
	@echo "サーバーを起動中..."
	./bin/api

migrate:
	@echo "データベースマイグレーションを実行中..."
	@if [ -z "$$DB_NAME" ]; then \
		echo "エラー: DB_NAME環境変数が設定されていません"; \
		exit 1; \
	fi
	psql -h $$DB_HOST -p $$DB_PORT -U $$DB_USER -d $$DB_NAME -f migration/001_create_tables.sql

clean:
	@echo "ビルド成果物を削除中..."
	rm -rf bin/
	@echo "クリーンアップ完了"
