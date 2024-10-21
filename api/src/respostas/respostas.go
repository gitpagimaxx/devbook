package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna um JSON para a requisição
func JSON(w http.ResponseWriter, status int, dados interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Erro retorna um JSON com a mensagem de erro
func Erro(w http.ResponseWriter, status int, erro error) {
	JSON(w, status, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
