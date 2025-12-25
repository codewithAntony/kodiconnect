package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	email VARCHAR(255) UNIQUE,
	password VARCHAR(255),
	role ENUM('LANDLORD', 'TENANT')
)`,
		`CREATE TABLE IF NOT EXISTS payments (
	id INT AUTO_INCREMENT PRIMARY KEY,
	tenant_id INT,
	amount DECIMAL(10, 2),
	mpesa_ref VARCHAR(50),
	status ENUM('PENDING', 'PAID'),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`,
		`CREATE TABLE IF NOT EXISTS maintenance_request (
	id INT AUTO_INCREMENT PRIMARY KEY,
	tenant_id INT,
	issue TEXT,
	urgency VARCHAR(20),
	image_url VARCHAR(255),
	status VARCHAR(20)
)`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}

	}
	return nil
}
