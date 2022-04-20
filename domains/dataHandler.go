package domains

type DataHandlerService interface {
	GetData(lines []string) []Dado
}

type DataHandlerClient interface {
	Get(lines []string) []Dado
}
