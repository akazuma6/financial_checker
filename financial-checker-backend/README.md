# Financial Checker Backend

日本株財務健全性可視化プラットフォームのバックエンドAPI

## 概要

EDINET（金融庁）の財務データを収集・解析し、企業の財務健全性を可視化するためのREST APIです。

## 技術スタック

- **言語**: Go 1.25.3
- **フレームワーク**: Echo v4
- **データベース**: PostgreSQL
- **アーキテクチャ**: MVC

## ディレクトリ構成

```
financial-checker-backend/
├── cmd/
│   └── api/
│       ├── main.go          # エントリーポイント
│       └── routes/
│           └── routes.go   # ルーティング定義
├── config/
│   └── env.go              # 環境変数管理
├── connection/
│   └── database/
│       └── database.go     # データベース接続
├── src/
│   ├── controller/
│   │   └── company/
│   │       └── company_controller.go  # コントローラー
│   └── model/
│       ├── company/
│       │   └── company_model.go       # 企業モデル
│       └── financial/
│           └── financial_model.go     # 財務データモデル
├── migration/
│   └── 001_create_tables.sql          # データベースマイグレーション
├── go.mod
└── README.md
```

## セットアップ

### 1. 依存関係のインストール

```bash
go mod download
```

### 2. 環境変数の設定

`.env`ファイルを作成し、以下の環境変数を設定してください：

```env
APP_ENV=development
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=password
DB_NAME=financial_checker
```

### 3. データベースのセットアップ

PostgreSQLに接続し、マイグレーションファイルを実行：

```bash
psql -U postgres -d financial_checker -f migration/001_create_tables.sql
```

### 4. サーバーの起動

```bash
go run cmd/api/main.go
```

サーバーはデフォルトでポート8080で起動します。

## APIエンドポイント

### 財務データ取得

```
GET /api/v1/companies/{code}/financials
```

指定した企業の過去5年分の財務データを取得します。

**レスポンス例:**
```json
{
  "status": "success",
  "message": "財務データを取得しました",
  "data": [
    {
      "id": 1,
      "companyCode": "7203",
      "fiscalYear": 2024,
      "sales": 1000000000000,
      "operatingIncome": 100000000000,
      "netIncome": 80000000000,
      "netAssets": 500000000000,
      "totalAssets": 1000000000000,
      "cashEquivalents": 100000000000,
      "isConsolidated": true
    }
  ]
}
```

### 財務健全性スコア取得

```
GET /api/v1/companies/{code}/health
```

指定した企業の財務健全性スコア（S~D判定）を取得します。

**レスポンス例:**
```json
{
  "status": "success",
  "message": "健全性スコアを取得しました",
  "data": {
    "companyCode": "7203",
    "score": 90,
    "grade": "S",
    "equityRatio": 50.0,
    "currentRatio": 0.0,
    "roe": 0.0,
    "comment": "非常に健全な財務状態です"
  }
}
```

### ヘルスチェック

```
GET /health
```

サーバーの稼働状況を確認します。

## 開発

### コーディング規約

- サンプルアプリ（`sample/apakan-next/go`）と同じMVC構造に従う
- コントローラーはHTTPリクエスト/レスポンスの処理のみ
- ビジネスロジックはモデル層に実装
- データベース操作はトランザクション内で実行

### 今後の拡張予定

- EDINET APIからのXBRLデータ取得バッチ処理
- 名寄せエンジンの実装
- より詳細な財務健全性スコアリングロジック
- 認証・認可機能
