package datahandler

import (
	"loaddados/domains"
	"loaddados/pkg/date"
	"loaddados/pkg/utils"
	"log"
	"strings"
)

type dataHandler struct{}

func NewDataHandlerClient() *dataHandler {
	return &dataHandler{}
}

func (d *dataHandler) Get(lines []string) []domains.Dado {
	dados := make([]domains.Dado, 0)
	for i, line := range lines {
		if i > 0 {
			conteudo := strings.Fields(line)

			fMedio, err := utils.ParseFloat(conteudo[4])
			if err != nil {
				log.Println("an error was found: ", err)
			}

			fUltima, err := utils.ParseFloat(conteudo[5])
			if err != nil {
				log.Println("an error was found: ", err)
			}

			currentTime, err := date.ParseDate(conteudo[3])
			if err != nil {
				log.Println("an error was found: ", err)
			}

			//cria uma lista do tipo dado
			dado := domains.Dado{
				CPF:             conteudo[0],
				IS_PRIVATE:      conteudo[1],
				INCOMPLETO:      conteudo[2],
				DATA_COMPRA:     currentTime,
				TICKET_MEDIO:    fMedio,
				TICKET_ULTIMA:   fUltima,
				LOJA_FREQUENCIA: conteudo[6],
				ULTIMA_LOJA:     conteudo[7],
			}
			dados = append(dados, dado)
		}
	}
	return dados
}
