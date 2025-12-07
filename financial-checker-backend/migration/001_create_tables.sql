-- 企業マスタテーブル
CREATE TABLE IF NOT EXISTS companies (
    code VARCHAR(10) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    industry VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 財務データテーブル
CREATE TABLE IF NOT EXISTS financial_statements (
    id SERIAL PRIMARY KEY,
    company_code VARCHAR(10) NOT NULL REFERENCES companies(code) ON DELETE CASCADE,
    fiscal_year INT NOT NULL,
    sales BIGINT,
    operating_income BIGINT,
    net_income BIGINT,
    net_assets BIGINT,
    total_assets BIGINT,
    cash_equivalents BIGINT,
    is_consolidated BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(company_code, fiscal_year)
);

-- インデックス作成
CREATE INDEX IF NOT EXISTS idx_financial_statements_company_code ON financial_statements(company_code);
CREATE INDEX IF NOT EXISTS idx_financial_statements_fiscal_year ON financial_statements(fiscal_year);
