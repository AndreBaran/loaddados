package date

import "time"

func ParseDate(conteudo string) (time.Time, error) {
	return time.Parse("2006-01-02", conteudo)
}
