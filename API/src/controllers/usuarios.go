package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Criando Ususario"))

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Buscando Todos os usuarios"))

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Buscar um usuario"))

}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Atualizando Usuario"))

}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Deletando um  Ususario"))

}
