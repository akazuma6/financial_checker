# Financial Checker Frontend

日本株財務健全性可視化プラットフォームのフロントエンド

## 概要

Next.js 16を使用したReactアプリケーション。企業の財務データと健全性スコアを可視化します。

## 技術スタック

- **フレームワーク**: Next.js 16 (App Router)
- **言語**: TypeScript
- **UIライブラリ**: Material-UI (MUI)
- **スタイリング**: Tailwind CSS
- **グラフ**: Recharts
- **状態管理**: Zustand
- **通知**: Notistack

## セットアップ

### 1. 依存関係のインストール

```bash
npm install
```

### 2. 環境変数の設定

`.env.local`ファイルを作成し、以下の環境変数を設定してください：

```env
NEXT_PUBLIC_API_BASE_URL=http://localhost:8080
```

### 3. 開発サーバーの起動

```bash
npm run dev
```

ブラウザで [http://localhost:3000](http://localhost:3000) を開きます。

## 機能

### 企業検索

- 証券コードで企業を検索
- サンプルコード: 7203（トヨタ自動車）、6758（ソニーグループ）、9984（ソフトバンクグループ）

### 財務データ表示

- 過去5年分の財務データをグラフで表示
- 売上高、営業利益、当期純利益の推移を可視化
- 詳細な財務データをテーブル形式で表示

### 健全性スコア

- S~Dの5段階評価
- 自己資本比率に基づくスコアリング
- コメント表示

## ディレクトリ構成

```
financial-checker-frontend/
├── src/
│   ├── app/
│   │   ├── layout.tsx          # ルートレイアウト
│   │   ├── page.tsx            # トップページ（検索）
│   │   ├── companies/
│   │   │   └── [code]/
│   │   │       └── page.tsx   # 企業詳細ページ
│   │   ├── theme.tsx           # MUIテーマ設定
│   │   └── globals.css         # グローバルスタイル
│   └── services/
│       ├── utils.ts            # APIユーティリティ
│       └── company.ts          # 企業APIクライアント
├── public/                     # 静的ファイル
├── package.json
└── next.config.ts
```

## 開発

### ビルド

```bash
npm run build
```

### 本番環境での起動

```bash
npm start
```

### リント

```bash
npm run lint
```

## バックエンドとの連携

フロントエンドは以下のAPIエンドポイントを使用します：

- `GET /api/v1/companies/{code}/financials` - 財務データ取得
- `GET /api/v1/companies/{code}/health` - 健全性スコア取得

バックエンドサーバーが起動していることを確認してください。
