package database

import (
	"database/sql"
	"log"
	"time"
	"github.com/ClickHouse/clickhouse-go"
)

type ClickHouseConfig struct {
	URL string
	Database string
	Username string
	Password string
}

func NewClickHouse() (*sql.DB, error) {
	connStr := "tps://localhost:9000?database=auth_service&username=default&password="

	db, err := sql.Open("clickhouse", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to ClickHouse successfully")
	return db, nil
}

func RunMigrations(db *sql.DB) error {
	tenantTable := `
		CREATE TABLE IF NOT EXISTS tenants (
			id String,
			name String,
			description String,
			plan String,
			max_users Int32,
			max_printers Int32,
			is_active Boolean,
			created_at DateTime,
			updated_at DateTime
		) ENGINE = MergeTree()
		PARTITION BY toYYYYMM(created_at)
		ORDER BY (id, created_at)
	`
	userTable := `
		CREATE TABLE IF NOT EXISTS users (
			id String,
			tenant_id String,
			email String,
			password_hash String,
			first_name String,
			last_name String,
			role String,
			created_at DateTime,
			updated_at DateTime,
			is_active Boolean
		) ENGINE = MergeTree()
		PARTITION BY toYYYYMM(created_at)
		ORDER BY (tenant_id, email, created_at)
	`
	sessionTable := `
		CREATE TABLE IF NOT EXISTS sessions (
			id String,
			user_id String,
			tenant_id String,
			token String,
			expires_at DateTime,
			created_at DateTime
		) ENGINE = MergeTree()
		PARTITION BY toYYYYMM(created_at)
		ORDER BY (user_id, created_at)
	`

	_, err := db.Exec(tenantTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(userTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(sessionTable)
	return err
}