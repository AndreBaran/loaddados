package services

import (
	"loaddados/domains"
	"loaddados/internal/clients/database/postgres"
	"loaddados/pkg/utils"
	"log"
	"sync"
)

type Storage struct {
	dbClient *postgres.Database
}

func NewStorageService(dbClient *postgres.Database) *Storage {
	return &Storage{
		dbClient: dbClient,
	}
}

func (s *Storage) SendToInsert(dados []domains.Dado) error {
	var wg sync.WaitGroup
	wg.Add(5)

	//quebrando a massa de dados
	//em 5 blocos para execucao paralera
	//evitando demasiados acesso ao banco
	//o qual bloquei
	for i := 0; i <= 4; i++ {
		iniFor := (i * 10000)
		endFor := ((i + 1) * 10000)
		if endFor > len(dados) {
			endFor = len(dados)
		}
		//go routine
		go func() {
			for j := iniFor; j < endFor; j++ {
				//validacao cpf/cnpj
				bDadoValido := utils.IsValidCpf(dados[j].CPF) && ((dados[j].LOJA_FREQUENCIA == "NULL") || (utils.IsValidCnpj(dados[j].LOJA_FREQUENCIA))) && ((dados[j].ULTIMA_LOJA == "NULL") || (utils.IsValidCnpj(dados[j].ULTIMA_LOJA)))
				if bDadoValido {
					if err := s.InsertData(dados[j]); err != nil {
						log.Println(err)
					}
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return nil
}

func (s *Storage) InsertData(dado domains.Dado) error {
	return s.dbClient.Insert(dado)
}

func (s *Storage) Close() {
	s.dbClient.Close()
}
