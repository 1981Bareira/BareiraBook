package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{

	{
		URI:                "/usuario",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuario/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuario/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuario/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteUsuario,
		RequerAutenticacao: false,
	},
}
