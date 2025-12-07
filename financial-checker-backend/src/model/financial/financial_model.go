package financial

import (
	"database/sql"
	"fmt"

	Db "financial-checker-backend/connection/database"
)

// FinancialStatement 財務データの構造体
type FinancialStatement struct {
	ID               int    `json:"id" db:"id"`
	CompanyCode      string `json:"companyCode" db:"company_code"`
	FiscalYear       int    `json:"fiscalYear" db:"fiscal_year"`
	Sales            *int64 `json:"sales" db:"sales"`
	OperatingIncome  *int64 `json:"operatingIncome" db:"operating_income"`
	NetIncome        *int64 `json:"netIncome" db:"net_income"`
	NetAssets        *int64 `json:"netAssets" db:"net_assets"`
	TotalAssets      *int64 `json:"totalAssets" db:"total_assets"`
	CashEquivalents  *int64 `json:"cashEquivalents" db:"cash_equivalents"`
	IsConsolidated   bool   `json:"isConsolidated" db:"is_consolidated"`
}

// HealthScore 財務健全性スコア
type HealthScore struct {
	CompanyCode     string  `json:"companyCode"`
	Score           int     `json:"score"`           // 0-100のスコア
	Grade           string  `json:"grade"`           // S, A, B, C, D
	EquityRatio     float64 `json:"equityRatio"`     // 自己資本比率
	CurrentRatio    float64 `json:"currentRatio"`    // 流動比率
	ROE             float64 `json:"roe"`             // 自己資本利益率
	Comment         string  `json:"comment"`         // コメント
}

// FinancialModel 財務モデル
type FinancialModel struct {
	DB *Db.PostgreSQLHandler
}

// NewFinancialModel 新しい財務モデルを生成
func NewFinancialModel(db *Db.PostgreSQLHandler) *FinancialModel {
	return &FinancialModel{DB: db}
}

// GetByCompanyCode 企業コードで財務データを取得（過去数年分）
func (m *FinancialModel) GetByCompanyCode(companyCode string, years int) ([]FinancialStatement, error) {
	var statements []FinancialStatement

	err := m.DB.Transaction(func(tx *sql.Tx) error {
		query := `
			SELECT id, company_code, fiscal_year, sales, operating_income, net_income,
			       net_assets, total_assets, cash_equivalents, is_consolidated
			FROM financial_statements
			WHERE company_code = $1
			ORDER BY fiscal_year DESC
			LIMIT $2
		`
		rows, err := tx.Query(query, companyCode, years)
		if err != nil {
			return fmt.Errorf("財務データ取得エラー (query: %s, code: %s, years: %d): %w", query, companyCode, years, err)
		}
		defer rows.Close()

		for rows.Next() {
			var stmt FinancialStatement
			err := rows.Scan(
				&stmt.ID,
				&stmt.CompanyCode,
				&stmt.FiscalYear,
				&stmt.Sales,
				&stmt.OperatingIncome,
				&stmt.NetIncome,
				&stmt.NetAssets,
				&stmt.TotalAssets,
				&stmt.CashEquivalents,
				&stmt.IsConsolidated,
			)
			if err != nil {
				return fmt.Errorf("データ読み取りエラー: %w", err)
			}
			statements = append(statements, stmt)
		}

		return rows.Err()
	})

	if err != nil {
		return nil, err
	}

	return statements, nil
}

// CalculateHealthScore 財務健全性スコアを計算
func (m *FinancialModel) CalculateHealthScore(companyCode string) (*HealthScore, error) {
	var score HealthScore
	score.CompanyCode = companyCode

	err := m.DB.Transaction(func(tx *sql.Tx) error {
		// 最新の財務データを取得
		query := `
			SELECT net_assets, total_assets, cash_equivalents
			FROM financial_statements
			WHERE company_code = $1
			ORDER BY fiscal_year DESC
			LIMIT 1
		`
		var netAssets, totalAssets, cashEquivalents sql.NullInt64
		err := tx.QueryRow(query, companyCode).Scan(&netAssets, &totalAssets, &cashEquivalents)
		if err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("財務データが見つかりません: code=%s", companyCode)
			}
			return fmt.Errorf("財務データ取得エラー (query: %s, code: %s): %w", query, companyCode, err)
		}

		// 自己資本比率を計算
		if totalAssets.Valid && totalAssets.Int64 > 0 && netAssets.Valid {
			score.EquityRatio = float64(netAssets.Int64) / float64(totalAssets.Int64) * 100
		}

		// スコア計算（簡易版）
		// 自己資本比率に基づいてスコアを算出
		if score.EquityRatio >= 50 {
			score.Score = 90
			score.Grade = "S"
			score.Comment = "非常に健全な財務状態です"
		} else if score.EquityRatio >= 30 {
			score.Score = 75
			score.Grade = "A"
			score.Comment = "健全な財務状態です"
		} else if score.EquityRatio >= 20 {
			score.Score = 60
			score.Grade = "B"
			score.Comment = "やや不安定な財務状態です"
		} else if score.EquityRatio >= 10 {
			score.Score = 40
			score.Grade = "C"
			score.Comment = "不安定な財務状態です"
		} else {
			score.Score = 20
			score.Grade = "D"
			score.Comment = "非常に不安定な財務状態です"
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &score, nil
}
