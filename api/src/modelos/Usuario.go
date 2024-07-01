package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar() error {
	if erro := usuario.Validar(); erro != nil {
		return erro
	}

	usuario.Formatar()
	return nil
}

func (usuario *Usuario) Validar() error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório. Não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório. Não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o email é obrigatório. Não pode estar em branco")
	}

	if usuario.Senha == "" {
		return errors.New("a senha é obrigatório. Não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) Formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
