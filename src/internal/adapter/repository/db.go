package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Importamos el driver de PostgreSQL
)

// NewPostgresDB crea una conexión a PostgreSQL y la devuelve
func NewPostgresDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Probar la conexión
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("✅ Conectado a PostgreSQL correctamente")
	return db, nil
}
