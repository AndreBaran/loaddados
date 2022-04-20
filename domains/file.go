package domains

import (
	"os"
	"time"
)

type Dado struct {
	CPF             string    `json:"CPF"`
	IS_PRIVATE      string    `json:"PRIVATE"`
	INCOMPLETO      string    `json:"INCOMPLETO"`
	DATA_COMPRA     time.Time `json:"DATA_COMPRA"`
	TICKET_MEDIO    float64   `json:"TICKET_MEDIO"`
	TICKET_ULTIMA   float64   `json:"TICKET_ULTIMA"`
	LOJA_FREQUENCIA string    `json:"LOJA_FREQUENCIA"`
	ULTIMA_LOJA     string    `json:"ULTIMA_LOJA"`
}

type FileService interface {
	Run() []string
	Close()
}

type FileClient interface {
	ReadFile() (*os.File, error)
	FileScanner(readFile *os.File) []string
	Close()
}
