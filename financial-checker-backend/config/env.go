package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var defaults = map[string]string{
	// アプリケーション設定
	"APP_ENV": "development",

	// データベース設定
	"DB_USER": "postgres",
	"DB_PASS": "",
	"DB_HOST": "localhost",
	"DB_PORT": "5432",
	"DB_NAME": "financial_checker",
}

func GetEnv(key string) (string, error) {
	if val := os.Getenv(key); val != "" {
		return val, nil
	}
	def, ok := defaults[key]
	if !ok {
		return "", fmt.Errorf("%sはデフォルトで定義されていません", key)
	}
	return def, nil
}

func GetEnvAsInt(key string) (int, error) {
	val := os.Getenv(key)
	if val == "" {
		var ok bool
		val, ok = defaults[key]
		if !ok {
			return 0, fmt.Errorf("%sはデフォルトで定義されていません", key)
		}
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("%s変換エラー: %v", key, err)
	}
	return intVal, nil
}

func GetEnvAsBool(key string) (bool, error) {
	val := os.Getenv(key)
	if val == "" {
		var ok bool
		val, ok = defaults[key]
		if !ok {
			return false, fmt.Errorf("%sはデフォルトで定義されていません", key)
		}
	}

	switch strings.ToLower(val) {
	case "true", "1", "yes":
		return true, nil
	case "false", "0", "no":
		return false, nil
	default:
		return false, fmt.Errorf("%s変換エラー: %s", key, val)
	}
}

// IsDevelopment 開発環境かどうかを判定
func IsDevelopment() bool {
	env, _ := GetEnv("APP_ENV")
	return env == "development" || env == ""
}
