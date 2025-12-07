package company

import (
	"database/sql"
	"fmt"

	Db "financial-checker-backend/connection/database"
)

// Company 企業情報の構造体
type Company struct {
	Code     string `json:"code" db:"code"`
	Name     string `json:"name" db:"name"`
	Industry string `json:"industry" db:"industry"`
}

// CompanyModel 企業モデル
type CompanyModel struct {
	DB *Db.PostgreSQLHandler
}

// NewCompanyModel 新しい企業モデルを生成
func NewCompanyModel(db *Db.PostgreSQLHandler) *CompanyModel {
	return &CompanyModel{DB: db}
}

// GetByCode 証券コードで企業を取得
func (m *CompanyModel) GetByCode(code string) (*Company, error) {
	var company Company

	err := m.DB.Transaction(func(tx *sql.Tx) error {
		query := `SELECT code, name, industry FROM companies WHERE code = $1`
		err := tx.QueryRow(query, code).Scan(&company.Code, &company.Name, &company.Industry)
		if err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("企業が見つかりません: code=%s", code)
			}
			return fmt.Errorf("企業取得エラー: %w", err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &company, nil
}
