package repository

import (
	"instagram-clone/internal/config"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatal("Cannot load config:", err)
	}

	db, err := NewDatabase(cfg)
	if err != nil {
		t.Fatal("Cannot connect to database:", err)
	}

	// Test connection
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal("Cannot get database instance:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		t.Fatal("Cannot ping database:", err)
	}

}
