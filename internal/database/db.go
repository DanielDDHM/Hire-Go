package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDsn() string {
	err := godotenv.Load()

	if err != nil {
		fmt.Print("Error When Reload .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DB")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		user, password, dbName, host, port, sslMode,
	)

	return dsn
}

func ConnectDB() (*gorm.DB, error) {
	dsn := GetDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter *sql.DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao pingar o banco de dados: %v", err)
	}

	fmt.Print("Db Connected")
	return db, nil
}
