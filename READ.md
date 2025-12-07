# 日本株財務健全性可視化プラットフォーム 仕様書（案）

**プロジェクト名（仮）**: J-Financial Visualizer
**作成日**: 2025年12月2日
**作成者**: Kazuma Kojima

---

## 1. プロジェクト概要

日本の個人投資家およびエンジニア向けに、EDINET（金融庁）の複雑な財務データ（XBRL）を収集・解析し、直感的に理解できる「ビジュアル」と、扱いやすい「整形済みJSON」を提供するプラットフォーム。

### 解決する課題（Pain Points）

- **投資家**: 決算書（短信・有報）は文字ばかりで読むのが苦痛。どの企業が本当に安全か、直感的に分からない。
- **エンジニア**: EDINET APIはXML形式で扱いづらく、勘定科目の表記ゆれ（タクソノミの不統一）が激しいため、データ活用が進まない。

### 提供価値（Value Proposition）

- **For 投資家**: 信号機やメーター表示で、企業の「倒産リスク」や「割安度」が3秒でわかる。
- **For エンジニア**: 独自の名寄せエンジンにより標準化された財務データを、API経由で簡単に取得できる（将来的なマネタイズ源）。

---

## 2. 技術スタック

### バックエンド (API & Batch)

- **言語**: Go (1.23+)
- **フレームワーク**: Echo v4
- **役割**:
  - EDINET APIからのXBRLファイル定期取得（バッチ処理）
  - XBRLのパースおよび名寄せ処理（正規化）
  - フロントエンドへのデータ提供 (REST API)

### フロントエンド (UI)

- **言語**: TypeScript
- **フレームワーク**: Next.js (App Router)
- **UIライブラリ**: Tailwind CSS, Shadcn/ui (推奨)
- **グラフ描画**: Recharts (React用チャートライブラリ)

### データベース & インフラ

- **DB**: PostgreSQL (リレーショナルデータ管理)
- **AI/ML**: OpenAI API (gpt-4o-mini等) または Pythonマイクロサービス（表記ゆれの推論用）
- **環境**: Docker (開発環境)

---

## 3. システムアーキテクチャ図

```mermaid
graph TD
    User[ユーザー (投資家/Eng)] -->|閲覧| FE[Next.js フロントエンド]
    FE -->|JSON要求| API[Go (Echo) APIサーバー]

    subgraph "Backend System"
        API -->|クエリ| DB[(PostgreSQL)]

        Batch[Go バッチ処理] -->|1. XBRL取得| EDINET[EDINET API]
        Batch -->|2. 解析・抽出| Normalizer[名寄せエンジン]

        Normalizer -->|3. 未知のタグ判定| AI[OpenAI API / MLモデル]
        AI -->|4. 推論結果| Normalizer
        Normalizer -->|5. 正規化データ保存| DB
    end
```

---

## 4. 機能要件 (MVP: Minimum Viable Product)

### A. データ収集・正規化機能 (Backend)

#### XBRLダウンローダー
指定した証券コード（まずは有名企業数社に限定）の最新有価証券報告書を取得する。

#### ハイブリッド名寄せエンジン
- **辞書マッチング**: 事前に定義したマッピングテーブル（例: OperatingIncome = 営業利益）で高速変換。
- **AIフォールバック**: 辞書にないタグが出現した場合、LLMに問い合わせて「これは営業利益か？」を判定し、辞書を更新する。

### B. API機能 (Backend)

- `GET /api/v1/companies/{code}/financials`: 指定企業の過去数年分の主要財務データを返す。
- `GET /api/v1/companies/{code}/health`: 独自のスコアリング結果（S~D判定など）を返す。

### C. ユーザーインターフェース (Frontend)

#### 検索画面
証券コードまたは企業名での検索。

#### ダッシュボード
- **財務健全性メーター**: 自己資本比率などを元にした安全度メーター。
- **推移グラフ**: 売上高・営業利益の過去5年推移（棒グラフ＋折れ線）。
- **一言コメント**: 「この企業は昨年より利益率が改善しています」等の自動生成コメント。

---

## 5. データベース設計 (簡易版)

### companies (企業マスタ)

| カラム名 | 型 | 説明 |
|---------|-----|------|
| code | VARCHAR(10) | 証券コード (PK) |
| name | VARCHAR(255) | 企業名 |
| industry | VARCHAR(100) | 業種 |

### financial_statements (財務データ)

| カラム名 | 型 | 説明 |
|---------|-----|------|
| id | SERIAL | ID (PK) |
| company_code | VARCHAR(10) | 証券コード (FK) |
| fiscal_year | INT | 会計年度 (例: 2024) |
| sales | BIGINT | 売上高 |
| operating_income | BIGINT | 営業利益 |
| net_income | BIGINT | 当期純利益 |
| net_assets | BIGINT | 純資産 |
| total_assets | BIGINT | 総資産 |
| cash_equivalents | BIGINT | 現預金 |
| is_consolidated | BOOLEAN | 連結決算かどうか |

---

## 6. マネタイズ & ロードマップ

### Phase 1: MVP開発・ポートフォリオ化 (1ヶ月目)

- **目標**: 自分の就職活動・技術力アピールに使えるレベル。
- **機能**: 主要50社程度の手動バッチ取得。単純なグラフ表示。
- **収益**: 0円。

### Phase 2: 自動化 & エンジニア向け公開 (2~3ヶ月目)

- **目標**: エンジニア界隈での認知獲得。
- **機能**: 全上場企業の自動取得（Cron実行）。AIによる名寄せ精度の向上。
- **収益**: 技術ブログ(Zenn/Qiita)からのアフィリエイト、Noteでのデータセット販売（CSV）。

### Phase 3: プラットフォーム化 (半年〜)

- **目標**: 持続的な収益化。
- **機能**: APIの外部公開（RapidAPI等を使用）。一般投資家向けの「割安株通知機能」。
- **収益**: API利用料、証券口座開設アフィリエイト。

---

## 7. 開発ディレクトリ構成案

```
j-financial-visualizer/
├── backend/ (Go)
│   ├── cmd/
│   │   ├── api/          # APIサーバーのエントリーポイント
│   │   └── batch/        # EDINET取得バッチのエントリーポイント
│   ├── internal/
│   │   ├── models/       # Model層: DB構造体定義、ビジネスロジック
│   │   ├── controllers/  # Controller層: HTTPハンドラー、リクエスト/レスポンス処理
│   │   ├── repository/   # データアクセス層: DB操作
│   │   ├── edinet/       # EDINET API操作・XBRLパースロジック
│   │   └── normalizer/   # 名寄せロジック (辞書 + AI)
│   ├── go.mod
│   └── main.go
├── frontend/ (Next.js)
│   ├── src/
│   │   ├── app/          # App Router
│   │   ├── components/   # UIコンポーネント (Graphs, Cards)
│   │   └── lib/          # APIクライアント
│   └── package.json
└── docker-compose.yml    # DB立ち上げ用
```
