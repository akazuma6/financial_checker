# Financial Checker Backend セットアップガイド

## 前提条件

- Go 1.25.3以上
- PostgreSQL 12以上

## セットアップ手順

### 1. 環境変数の設定

`.env`ファイルをプロジェクトルートに作成し、以下の内容を設定してください：

```env
APP_ENV=development
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=your_password
DB_NAME=financial_checker
PORT=8080
```

**注意**: PostgreSQLのデフォルトユーザーは`postgres`です。環境に応じて変更してください。

### 2. PostgreSQLのセットアップ

#### データベースの作成

```bash
# PostgreSQLに接続
psql -U postgres

# データベースを作成
CREATE DATABASE financial_checker;

# 接続を確認
\c financial_checker

# マイグレーションを実行
\i migration/001_create_tables.sql
\i migration/002_sample_data.sql
```

または、コマンドラインから直接実行：

```bash
# データベース作成
createdb -U postgres financial_checker

# マイグレーション実行
psql -U postgres -d financial_checker -f migration/001_create_tables.sql
psql -U postgres -d financial_checker -f migration/002_sample_data.sql
```

### 3. 依存関係のインストール

```bash
go mod download
```

### 4. サーバーの起動

```bash
go run cmd/api/main.go
```

サーバーはデフォルトでポート8080で起動します。

### 5. 動作確認

```bash
# ヘルスチェック
curl http://localhost:8080/health

# 財務データ取得（サンプル）
curl http://localhost:8080/api/v1/companies/7203/financials

# 健全性スコア取得（サンプル）
curl http://localhost:8080/api/v1/companies/7203/health
```

## トラブルシューティング

### PostgreSQL接続エラー

**エラー**: `role "root" does not exist`

**解決方法**:
- `.env`ファイルで`DB_USER=postgres`を設定
- または、PostgreSQLに`root`ユーザーを作成

### データベースが存在しない

**エラー**: `database "financial_checker" does not exist`

**解決方法**:
```bash
createdb -U postgres financial_checker
```

### パスワード認証エラー

**エラー**: `password authentication failed`

**解決方法**:
- `.env`ファイルで正しいパスワードを設定
- または、PostgreSQLの`pg_hba.conf`で認証方式を確認

## サンプルデータ

マイグレーションファイル`002_sample_data.sql`には以下の企業のサンプルデータが含まれています：

- 7203: トヨタ自動車
- 6758: ソニーグループ
- 9984: ソフトバンクグループ
- 9434: KDDI
- 4063: 信越化学工業

これらの証券コードでAPIをテストできます。
