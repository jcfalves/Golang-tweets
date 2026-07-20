package internalsql

import (
	"database/sql"
	"fmt"
	"go-tweets/internal/config"
	"log"
)

func ConnectMySql(cfg *config.Config) (*sql.DB, error) {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	log.Println("Connected to MySQL database successfully")

	return db, nil
}
