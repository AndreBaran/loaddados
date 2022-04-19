package utils

import (
	"testing"
)

func TestCpf(t *testing.T) {

	tabelas := []struct {
		cpf     string
		isValid bool
	}{
		{"123.456.789-10", false},
		{"123.456.789-09", true},
		{"066.098.619-29", false},
		{"066.098.619-19", true},
	}

	for _, tabela := range tabelas {
		result := IsValidCpf(tabela.cpf)
		if result != tabela.isValid {
			t.Errorf("cpf invalido")
		}
	}

}

func TestCnpj(t *testing.T) {

	tabelas := []struct {
		cnpj    string
		isValid bool
	}{
		{"123.456.789-10", false},
		{"79.379.491/0008-50", true},
		{"79.379.491/0008-51", false},
		{"066.098.619-19", false},
	}

	for _, tabela := range tabelas {
		result := IsValidCnpj(tabela.cnpj)
		if result != tabela.isValid {
			t.Errorf("cnpj invalido")
		}
	}

}
