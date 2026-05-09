package models

import (
	"errors"

	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	RG   string `json:"rg"`
	CPF  string `json:"cpf"`
}

func ValidaDadosDeAluno(aluno *Aluno) error {
	if aluno.Nome == "" {
		return errors.New("nome deve ser informado")
	}
	if len(aluno.RG) != 9 || !hasOnlyDigits(aluno.RG) {
		return errors.New("rg deve conter 9 digitos")
	}
	if len(aluno.CPF) != 11 || !hasOnlyDigits(aluno.CPF) {
		return errors.New("cpf deve conter 11 digitos")
	}
	return nil
}

func hasOnlyDigits(value string) bool {
	for _, char := range value {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
