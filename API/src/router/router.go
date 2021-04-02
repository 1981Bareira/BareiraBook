package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

//Esta funcao Gerar vai retornar um router com as rotaas configuradas
func Gerar() *mux.Router {

	r := mux.NewRouter()
	return rotas.Configurar(r)
}
