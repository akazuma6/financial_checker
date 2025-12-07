package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"financial-checker-backend/config"

	_ "github.com/lib/pq"
)

// DBConfig データベース設定用の構造体
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// PostgreSQLHandler PostgreSQL接続を管理する構造体
type PostgreSQLHandler struct {
	DB     *sql.DB
	Config DBConfig
}

// NewPostgreSQLHandler DBハンドラーを生成
func NewPostgreSQLHandler(config DBConfig) (*PostgreSQLHandler, error) {
	// DSN (Data Source Name) の構築
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	// デバッグ用: 接続情報をログ出力（パスワードは除く）
	fmt.Printf("Connecting to database: host=%s port=%d user=%s dbname=%s\n",
		config.Host, config.Port, config.User, config.DBName)

	// データベース接続
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// 接続設定
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(15 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)

	// 接続テスト
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgreSQLHandler{
		DB:     db,
		Config: config,
	}, nil
}

// GetDBConfig 環境変数からDB設定を取得
func GetDBConfig() (DBConfig, error) {
	host, err := config.GetEnv("DB_HOST")
	if err != nil {
		return DBConfig{}, err
	}

	port, err := config.GetEnvAsInt("DB_PORT")
	if err != nil {
		return DBConfig{}, err
	}

	user, err := config.GetEnv("DB_USER")
	if err != nil {
		return DBConfig{}, err
	}

	password, err := config.GetEnv("DB_PASS")
	if err != nil {
		return DBConfig{}, err
	}

	dbName, err := config.GetEnv("DB_NAME")
	if err != nil {
		return DBConfig{}, err
	}

	return DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}, nil
}

// GetDB DB接続を取得
func (h *PostgreSQLHandler) GetDB() *sql.DB {
	return h.DB
}

// Close データベース接続を閉じる
func (h *PostgreSQLHandler) Close() error {
	if h.DB != nil {
		return h.DB.Close()
	}
	return nil
}

// Transaction トランザクションを実行
func (h *PostgreSQLHandler) Transaction(txFunc func(*sql.Tx) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	tx, err := h.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := txFunc(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("failed to rollback: %v (original error: %w)", rbErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// Ping データベース接続を確認
func (h *PostgreSQLHandler) Ping() error {
	return h.DB.Ping()
}
