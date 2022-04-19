package services

import "loaddados/domains"

type dataHandlerService struct {
	dataHandlerClient domains.DataHandlerClient
}

func NewDataHandlerService(dataHandlerClient domains.DataHandlerClient) *dataHandlerService {
	return &dataHandlerService{
		dataHandlerClient: dataHandlerClient,
	}
}

func (d *dataHandlerService) GetData(lines []string) []domains.Dado {
	return d.dataHandlerClient.Get(lines)
}
