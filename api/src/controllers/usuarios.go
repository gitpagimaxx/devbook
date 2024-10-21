package controllers

import (
	"api/src/db"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpo, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario modelos.Usuario
	if err = json.Unmarshal(corpo, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorios := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, err = repositorios.Criar(usuario)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	db, err := db.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	repositorios := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, err := repositorios.BuscarUsuarios()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		log.Fatal(err)
	}

}

// BuscarUsuario busca um usuário salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

// AtualizarUsuario altera as informações de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário"))
}

// DeletarUsuario exclui as informações de um usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuário"))
}
