package routes

import (
	"github.com/labstack/echo/v4"

	CompanyController "financial-checker-backend/src/controller/company"
)

// RegisterRoutes すべてのルートを登録
func RegisterRoutes(e *echo.Echo) {
	// API v1グループ
	v1 := e.Group("/api/v1")

	// 企業関連のルート
	RegisterCompanyRoutes(v1)
}

// RegisterCompanyRoutes 企業関連のルートを登録
func RegisterCompanyRoutes(g *echo.Group) {
	companies := g.Group("/companies")

	// 財務データ取得
	companies.GET("/:code/financials", CompanyController.GetFinancials)

	// 健全性スコア取得
	companies.GET("/:code/health", CompanyController.GetHealth)
}
