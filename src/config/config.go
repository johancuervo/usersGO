package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("❌ Error cargando el archivo .env")
	}
	log.Println("✅ Variables de entorno cargadas correctamente")
}
