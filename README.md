# Financial Checker

日本株財務健全性可視化プラットフォーム

## 概要

EDINET（金融庁）の財務データを収集・解析し、企業の財務健全性を可視化するためのWebアプリケーションです。

## 技術スタック

### バックエンド
- Go 1.25.3
- Echo v4
- PostgreSQL

### フロントエンド
- Next.js 16
- TypeScript
- Material-UI
- Recharts

## クイックスタート

### 1. PostgreSQLの起動（Docker Compose）

```bash
# プロジェクトルートで実行
docker-compose up -d

# ログ確認
docker-compose logs -f postgres
```

### 2. バックエンドのセットアップ

```bash
cd financial-checker-backend

# .envファイルを作成
cp .env.example .env
# .envファイルを編集して必要に応じて設定を変更

# 依存関係のインストール
go mod download

# サーバー起動
go run cmd/api/main.go
```

バックエンドは `http://localhost:8080` で起動します。

### 3. フロントエンドのセットアップ

```bash
cd financial-checker-frontend

# 依存関係のインストール
npm install

# .env.localファイルを作成
echo "NEXT_PUBLIC_API_BASE_URL=http://localhost:8080" > .env.local

# 開発サーバー起動
npm run dev
```

フロントエンドは `http://localhost:3000` で起動します。

## データベース

### 初回セットアップ

Docker ComposeでPostgreSQLを起動すると、`migration/`ディレクトリ内のSQLファイルが自動的に実行されます：

- `001_create_tables.sql` - テーブル作成
- `002_sample_data.sql` - サンプルデータ投入

### サンプルデータ

以下の企業のサンプルデータが含まれています：

- 7203: トヨタ自動車
- 6758: ソニーグループ
- 9984: ソフトバンクグループ
- 9434: KDDI
- 4063: 信越化学工業

### データベース接続

```bash
# Dockerコンテナ内のpsqlを使用
docker-compose exec postgres psql -U postgres -d financial_checker

# または、ローカルのpsqlを使用（インストール済みの場合）
psql -h localhost -U postgres -d financial_checker
```

## APIエンドポイント

- `GET /health` - ヘルスチェック
- `GET /api/v1/companies/{code}/financials` - 財務データ取得
- `GET /api/v1/companies/{code}/health` - 健全性スコア取得

## 開発

### バックエンド

```bash
cd financial-checker-backend
go run cmd/api/main.go
```

### フロントエンド

```bash
cd financial-checker-frontend
npm run dev
```

### データベースのリセット

```bash
# コンテナとボリュームを削除
docker-compose down -v

# 再起動（マイグレーションが自動実行される）
docker-compose up -d
```

## トラブルシューティング

### PostgreSQL接続エラー

`.env`ファイルで以下の設定を確認してください：

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=financial_checker
```

### ポートが既に使用されている

```bash
# ポート5432が使用されている場合
docker-compose down
# docker-compose.ymlのportsセクションを変更（例: "5433:5432"）
```

## ディレクトリ構成

```
financialchecker/
├── docker-compose.yml          # Docker Compose設定
├── financial-checker-backend/  # バックエンド（Go）
│   ├── cmd/api/               # APIエントリーポイント
│   ├── src/                   # ソースコード
│   ├── migration/             # データベースマイグレーション
│   └── .env                   # 環境変数
└── financial-checker-frontend/ # フロントエンド（Next.js）
    ├── app/                   # Next.js App Router
    ├── services/              # APIクライアント
    └── .env.local             # 環境変数
```
