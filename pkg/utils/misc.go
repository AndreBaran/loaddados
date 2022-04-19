package utils

import (
	"math"
	"strconv"
	"strings"
)

func ParseFloat(conteudo string) (float64, error) {
	newFloat, err := strconv.ParseFloat(strings.ReplaceAll(conteudo, ",", "."), 32)
	if err != nil {
		return 0, err
	}
	return math.Round(newFloat*100) / 100, nil
}

func IsValidCpf(sCpf string) bool {
	//regra validacao de cpf
	sCpf = strings.ReplaceAll(sCpf, "-", "")
	sCpf = strings.ReplaceAll(sCpf, "/", "")
	sCpf = strings.ReplaceAll(sCpf, ".", "")

	vCpf := strings.Split(sCpf, "")
	if len(vCpf) != 11 {
		return false
	}

	sum := 0
	va := 0
	j := 10
	for i := 0; i <= 8; i++ {
		va, _ = strconv.Atoi(vCpf[i])
		sum = sum + (va * j)
		j--
	}
	dig1 := 11 - (sum % 11)
	if dig1 >= 10 {
		dig1 = 0
	}
	sum = 0
	j = 11
	for i := 0; i <= 8; i++ {
		va, _ = strconv.Atoi(vCpf[i])
		sum = sum + va*j
		j--
	}
	sum = sum + (dig1 * 2)
	dig2 := 11 - (sum % 11)
	if dig2 >= 10 {
		dig2 = 0
	}
	val1, _ := strconv.Atoi(vCpf[9])
	val2, _ := strconv.Atoi(vCpf[10])

	return (dig1 == val1) && (dig2 == val2)
}

func IsValidCnpj(sCnpj string) bool {
	//regra validacao de cnpj
	sCnpj = strings.ReplaceAll(sCnpj, "-", "")
	sCnpj = strings.ReplaceAll(sCnpj, "/", "")
	sCnpj = strings.ReplaceAll(sCnpj, ".", "")

	vCnpj := strings.Split(sCnpj, "")
	if len(vCnpj) != 14 {
		return false
	}

	sum := 0
	va := 0
	j := 5
	for i := 0; i <= 11; i++ {
		va, _ = strconv.Atoi(vCnpj[i])
		sum = sum + (va * j)
		j--
		if j == 1 {
			j = 9
		}
	}
	dig1 := 11 - (sum % 11)
	if dig1 >= 10 {
		dig1 = 0
	}
	sum = 0
	j = 6
	for i := 0; i <= 11; i++ {
		va, _ = strconv.Atoi(vCnpj[i])
		sum = sum + va*j
		j--
		if j == 1 {
			j = 9
		}
	}
	sum = sum + (dig1 * 2)
	dig2 := 11 - (sum % 11)
	if dig2 >= 10 {
		dig2 = 0
	}
	val1, _ := strconv.Atoi(vCnpj[12])
	val2, _ := strconv.Atoi(vCnpj[13])

	return (dig1 == val1) && (dig2 == val2)
}
