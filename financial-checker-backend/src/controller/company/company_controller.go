package company

import (
	"net/http"

	Db "financial-checker-backend/connection/database"
	FinancialModel "financial-checker-backend/src/model/financial"

	"github.com/labstack/echo/v4"
)

// GetFinancialsResponse 財務データ取得レスポンス
type GetFinancialsResponse struct {
	Status  string                           `json:"status"`
	Message string                           `json:"message"`
	Data    []FinancialModel.FinancialStatement `json:"data,omitempty"`
}

// GetHealthResponse 健全性スコア取得レスポンス
type GetHealthResponse struct {
	Status  string                    `json:"status"`
	Message string                    `json:"message"`
	Data    *FinancialModel.HealthScore `json:"data,omitempty"`
}

// GetFinancials 企業の財務データを取得
func GetFinancials(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, GetFinancialsResponse{
			Status:  "error",
			Message: "証券コードが指定されていません",
		})
	}

	// DBハンドラーの取得
	dbValue := c.Get("db")
	if dbValue == nil {
		return c.JSON(http.StatusInternalServerError, GetFinancialsResponse{
			Status:  "error",
			Message: "データベース接続が取得できません",
		})
	}

	db, ok := dbValue.(*Db.PostgreSQLHandler)
	if !ok {
		return c.JSON(http.StatusInternalServerError, GetFinancialsResponse{
			Status:  "error",
			Message: "データベース型変換エラー",
		})
	}

	// モデルの初期化
	financialModel := FinancialModel.NewFinancialModel(db)

	// 財務データ取得（過去5年分）
	statements, err := financialModel.GetByCompanyCode(code, 5)
	if err != nil {
		// エラーログを出力
		c.Logger().Errorf("財務データ取得エラー: %v", err)
		return c.JSON(http.StatusInternalServerError, GetFinancialsResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, GetFinancialsResponse{
		Status:  "success",
		Message: "財務データを取得しました",
		Data:    statements,
	})
}

// GetHealth 企業の財務健全性スコアを取得
func GetHealth(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, GetHealthResponse{
			Status:  "error",
			Message: "証券コードが指定されていません",
		})
	}

	// DBハンドラーの取得
	dbValue := c.Get("db")
	if dbValue == nil {
		return c.JSON(http.StatusInternalServerError, GetHealthResponse{
			Status:  "error",
			Message: "データベース接続が取得できません",
		})
	}

	db, ok := dbValue.(*Db.PostgreSQLHandler)
	if !ok {
		return c.JSON(http.StatusInternalServerError, GetHealthResponse{
			Status:  "error",
			Message: "データベース型変換エラー",
		})
	}

	// モデルの初期化
	financialModel := FinancialModel.NewFinancialModel(db)

	// 健全性スコア計算
	score, err := financialModel.CalculateHealthScore(code)
	if err != nil {
		// エラーログを出力
		c.Logger().Errorf("健全性スコア計算エラー: %v", err)
		return c.JSON(http.StatusInternalServerError, GetHealthResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, GetHealthResponse{
		Status:  "success",
		Message: "健全性スコアを取得しました",
		Data:    score,
	})
}
