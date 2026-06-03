package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func dsn() string {
	databaseURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(os.Getenv("DB_USER"), os.Getenv("DB_PASS")),
		Host:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	}

	query := databaseURL.Query()
	query.Set("sslmode", os.Getenv("DB_SSLMODE"))
	databaseURL.RawQuery = query.Encode()

	return databaseURL.String()
}

func connectDB(context context.Context) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn())
	if err != nil {
		return nil, fmt.Errorf("Loi khi ket noi database: %v", err)
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		return nil, fmt.Errorf("DB_MAX_OPEN_CONNS không hợp lệ: %w", err)
	}

	maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		return nil, fmt.Errorf("DB_MAX_IDLE_CONNS không hợp lệ: %w", err)
	}

	connMaxLifetime, err := time.ParseDuration(os.Getenv("DB_CONN_MAX_LIFETIME"))
	if err != nil {
		return nil, fmt.Errorf("DB_CONN_MAX_LIFETIME không hợp lệ: %w", err)
	}

	connMaxIdleTime, err := time.ParseDuration(os.Getenv("DB_CONN_MAX_IDLE_TIME"))
	if err != nil {
		return nil, fmt.Errorf("DB_CONN_MAX_IDLE_TIME không hợp lệ: %w", err)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	if err := db.PingContext(context); err != nil {
		return nil, fmt.Errorf("Loi khi ping database: %v", err)
	}
	return db, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	db, err := connectDB(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	log.Println("Successfully connected to the database!")
}
