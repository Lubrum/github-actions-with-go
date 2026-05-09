package database

import (
	"fmt"
	"os"

	"github.com/Lubrum/github-actions-with-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectaComBancoDeDados() error {
	db, err := gorm.Open(postgres.Open(connectionString()))
	if err != nil {
		return fmt.Errorf("conectar com banco de dados: %w", err)
	}

	if err := db.AutoMigrate(&models.Aluno{}); err != nil {
		return fmt.Errorf("migrar banco de dados: %w", err)
	}

	DB = db
	return nil
}

func connectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
		os.Getenv("DBPORT"),
	)
}
