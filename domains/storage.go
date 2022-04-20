package domains

type StorageService interface {
	SendToInsert(dados []Dado) error
	InsertData(dado Dado) error
	Close()
}

type DatabaseClient interface {
	Insert(dado Dado) error
	Close()
}
