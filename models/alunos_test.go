package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidaDadosDeAluno(t *testing.T) {
	tests := []struct {
		name    string
		aluno   Aluno
		wantErr bool
	}{
		{
			name: "aluno valido",
			aluno: Aluno{
				Nome: "Nome do Aluno",
				RG:   "123456789",
				CPF:  "12345678901",
			},
			wantErr: false,
		},
		{
			name: "nome vazio",
			aluno: Aluno{
				RG:  "123456789",
				CPF: "12345678901",
			},
			wantErr: true,
		},
		{
			name: "rg com tamanho invalido",
			aluno: Aluno{
				Nome: "Nome do Aluno",
				RG:   "123",
				CPF:  "12345678901",
			},
			wantErr: true,
		},
		{
			name: "rg com letras",
			aluno: Aluno{
				Nome: "Nome do Aluno",
				RG:   "12345678a",
				CPF:  "12345678901",
			},
			wantErr: true,
		},
		{
			name: "cpf com tamanho invalido",
			aluno: Aluno{
				Nome: "Nome do Aluno",
				RG:   "123456789",
				CPF:  "123",
			},
			wantErr: true,
		},
		{
			name: "cpf com letras",
			aluno: Aluno{
				Nome: "Nome do Aluno",
				RG:   "123456789",
				CPF:  "1234567890a",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidaDadosDeAluno(&tt.aluno)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
