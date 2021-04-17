package controllers

import (
	"api/src/db"
	"api/src/model"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	var usuario model.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		{
			log.Fatal(erro)
		}
	}
	db, erro := db.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("Usuario criado com o ID %d", usuarioID)))
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
