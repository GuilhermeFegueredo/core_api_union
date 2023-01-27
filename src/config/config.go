package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexao é a string de conexão com o MySQL
	DNS = ""

	// Port onde a API vai estar rodando
	Port = 0

	// SecretKey é a chave que vai assinar o token
	SecretKey []byte
)

// Load vai inicializar as variaveis de ambiente
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	DNS = fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
