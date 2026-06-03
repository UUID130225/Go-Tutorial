package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	EnvInt   int
	EnvStr   string
	EnvBool  bool
	EnvFloat float64
	EnvList  []string
	EnvTime  int
}

func validateEnv(cfg *Config) error {
	if cfg.App.EnvInt <= 0 {
		return fmt.Errorf("ENV_INT phai la mot so nguyen duong")
	}

	if strings.TrimSpace(cfg.App.EnvStr) == "" {
		return fmt.Errorf("ENV_STR khong duoc de trong")
	}

	if cfg.App.EnvTime <= 0 {
		return fmt.Errorf("ENV_TIME phai la mot so nguyen duong")
	}

	return nil
}

func getEnvBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		log.Printf("Loi chuyen %s sang bool: %v", key, err)
		return defaultValue
	}
	return value
}

func getEnvFloat(key string, defaultValue float64) float64 {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Printf("Loi chuyen %s sang float: %v", key, err)
		return defaultValue
	}
	return value
}

func getEnvList(key, sep string) []string {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return []string{}
	}
	return strings.Split(valueStr, sep)
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Loi chuyen %s sang so nguyen: %v", key, err)
		return defaultValue
	}
	return value
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func loadEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Khong tim thay file .env, su dung gia tri mac dinh")
	}
	cfg := &Config{
		App: AppConfig{
			EnvInt:   getEnvInt("ENV_INT", 0),
			EnvStr:   getEnv("ENV_STR", ""),
			EnvBool:  getEnvBool("ENV_BOOL", false),
			EnvFloat: getEnvFloat("ENV_FLOAT", 0.0),
			EnvList:  getEnvList("ENV_LIST", ","),
			EnvTime:  getEnvInt("ENV_TIME", 0),
		},
	}
	if err := validateEnv(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func main() {

	cfg, err := loadEnv()
	if err != nil {
		log.Fatalf("Loi khi load env: %v", err)
	}

	fmt.Printf("ENV_INT: %d\n", cfg.App.EnvInt)
	fmt.Printf("ENV_STR: %s\n", cfg.App.EnvStr)
	fmt.Printf("ENV_BOOL: %t\n", cfg.App.EnvBool)
	fmt.Printf("ENV_FLOAT: %f\n", cfg.App.EnvFloat)
	fmt.Printf("ENV_LIST: %v\n", cfg.App.EnvList)
	fmt.Printf("ENV_TIME: %d\n", cfg.App.EnvTime)

}
