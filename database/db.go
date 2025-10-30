package database

import (
	"log"
	"os"
	"time"
	"fmt"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ...existing code...
var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		DB, err = gorm.Open(postgres.Open(dsn))
		if err == nil {
			break
		}
		log.Printf("Tentativa %d/%d: erro ao conectar com banco: %v", attempt, maxAttempts, err)
		if attempt == maxAttempts {
			log.Panic("Erro ao conectar com banco de dados")
		}
		time.Sleep(3 * time.Second)
	}

	if DB == nil {
		log.Panic("Erro ao conectar com banco de dados: DB nil após tentativas")
	}

	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Printf("AutoMigrate error: %v", err)
	}
}


