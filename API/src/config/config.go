package config

import (
	"log"

	"github.com/joho/godotenv"
)

var (
	//String de conexao com o banco de dados
	StringConexaoBanco = ""

	//Porta onde a aplicacao esta rodando
	Porta = 0
)

//A função Carregar vai inicializar as variáveis de ambiente
func Carregar() {

	var erro error

	if erro = godotenv.Load(); erro != nil {

		log.Fatal()

	}
}
