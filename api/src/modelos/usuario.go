package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID          uint64    `json:"id,omitempty"`
	Nome        string    `json:"nome,omitempty"`
	Nick        string    `json:"nick,omitempty"`
	Email       string    `json:"email,omitempty"`
	Senha       string    `json:"senha,omitempty"`
	DataCriacao time.Time `json:"data_criacao,omitempty"`
}

// Validar verifica se os campos do usuário estão preenchidos corretamente
func (usuario Usuario) Validar() error {
	if usuario.Nome == "" {
		return errors.New("o campo nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("o campo nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o campo email é obrigatório e não pode estar em branco")
	}

	if usuario.Senha == "" {
		return errors.New("o campo senha é obrigatório e não pode estar em branco")
	}

	return nil
}

// Preparar vai chamar os métodos para validar e formatar o usuário
func (usuario *Usuario) Preparar() error {
	usuario.formatar()
	err := usuario.Validar()
	if err != nil {
		return err
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
