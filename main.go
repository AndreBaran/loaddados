package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"math"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "user"
	password = "mypassword"
	dbname   = "desafio"
)

var dados []Dado

type Dado struct {
	CPF             string  `json:"CPF"`
	PRIVATE         string  `json:"PRIVATE"`
	INCOMPLETO      string  `json:"INCOMPLETO"`
	DATA_COMPRA     time.Time  `json:"DATA_COMPRA"`
	TICKET_MEDIO    float64 `json:"TICKET_MEDIO"`
	TICKET_ULTIMA   float64 `json:"TICKET_ULTIMA"`
	LOJA_FREQUENCIA string  `json:"LOJA_FREQUENCIA"`
	ULTIMA_LOJA     string  `json:"ULTIMA_LOJA"`
}

func main() {
   //momento de criar banco de dados
	err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println("table DADOS is created")
	//leitura do arquivo
	readFile, err := os.Open("base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()

	for i, line := range lines {
		if i > 0 {
			conteudo := strings.Fields(line)
			
			fMedio, _ := strconv.ParseFloat(strings.ReplaceAll( conteudo[4],",","."), 32)
			fUltima, _ := strconv.ParseFloat(strings.ReplaceAll(conteudo[5],",","."), 32)

			fMedio= math.Round(fMedio*100)/100
			fUltima= math.Round(fUltima*100)/100
			
			currentTime := time.Now()
			currentTime,_ = time.Parse("2006-01-02",conteudo[3])
			
			//cria uma lista do tipo dado
			dado := Dado{conteudo[0], conteudo[1], conteudo[2], currentTime,
					fMedio, fUltima, conteudo[6], conteudo[7]}
			dados=append(dados,dado)
			if i == 1 {
				fmt.Println(conteudo)
				fmt.Println(fMedio)
				fmt.Println(fUltima)
				
				fmt.Println(dado)
			}
		
		}
		
	}
	//envia para o banco a lista de dados
	persistDados()
	//fmt.Println(dados)
}

func persistDados(){
	//fmt.Println(dados)
	start := time.Now()
	fmt.Println(start)
	//abre apenas uma conexao, para evitar a lentidao
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	   "password=%s dbname=%s sslmode=disable",
	   host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
	for _, iDado := range dados {
		//verifica se os dados de cpf e cnpj estao corretos
		bDadoValido:=isValidCpf(iDado.CPF) && ((iDado.LOJA_FREQUENCIA =="NULL") || (isValidCnpj(iDado.LOJA_FREQUENCIA))) && ((iDado.ULTIMA_LOJA =="NULL") || (isValidCnpj(iDado.ULTIMA_LOJA)))
			if bDadoValido {
		    //salva o dado no banco
			persistDado(iDado,db)
		}
	}
	db.Close()
	final := time.Now()
	diff := final.Sub(start)
	fmt.Println(diff)
}
func persistDado(perDado Dado,db *sql.DB ) error {
	//query que insere o dado no banco
	qry, err := db.Prepare("INSERT INTO DADOS (CPF, PRIVATE, INCOMPLETO, DATA_COMPRA, TICKET_MEDIO, TICKET_ULTIMA, LOJA_FREQUENCIA, ULTIMA_LOJA) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)")
	if err != nil {
		panic(err)
	}
	_, err = qry.Exec(perDado.CPF,perDado.PRIVATE,perDado.INCOMPLETO ,perDado.DATA_COMPRA,perDado.TICKET_MEDIO ,perDado.TICKET_ULTIMA,perDado.LOJA_FREQUENCIA ,perDado.ULTIMA_LOJA)
	if err != nil {
		panic(err)
	}
	return nil

}

func createDataBase() error {
	//funcao que cria o banco de dados
	fmt.Println("connecting")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	start := time.Now()
	for db.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("failed to connect after 10 secs.")
			break
		}
	}
	fmt.Println("connected:", db.Ping() == nil)
	_, err = db.Exec(`DROP TABLE IF EXISTS DADOS;`)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE DADOS (CPF character varying, PRIVATE char, INCOMPLETO char, DATA_COMPRA date, TICKET_MEDIO double precision, TICKET_ULTIMA double precision, LOJA_FREQUENCIA character varying, ULTIMA_LOJA  character varying);`)
	if err != nil {
		panic(err)
	}
	

	fmt.Println("table DADOS is created")
	return nil
}

func isValidCpf(sCpf string) bool{
	//regra validacao de cpf
	sCpf = strings.ReplaceAll(sCpf,"-","")
	sCpf = strings.ReplaceAll(sCpf,"/","")
	sCpf = strings.ReplaceAll(sCpf,".","")
    
	vCpf := strings.Split(sCpf,"")
	if (len(vCpf) != 11) {
		return false
	} else {
		sum:=0
		va:=0
		j:=10
		for i:=0;i<=8; i++ {
			va,_=strconv.Atoi(vCpf[i])				
			sum=sum+(va*j)
			j--			
		}
		dig1:=11-(sum%11)	
		if dig1>=10{
			dig1=0
		}	
		sum=0
		j=11
		for i:=0;i<=8; i++ {
			va,_=strconv.Atoi(vCpf[i])
			sum=sum+va*j
			j--
		}
        sum= sum+(dig1*2)
		dig2:=11-(sum%11)	
		if dig2>=10{
			dig2=0
		}	
		val1,_:=strconv.Atoi(vCpf[9])
		val2,_:=strconv.Atoi(vCpf[10])
		
		return (dig1 == val1) && (dig2 == val2)	
		

	}
			
	return true	
}

func isValidCnpj(sCnpj string) bool{
	//regra validacao de cnpj
	sCnpj = strings.ReplaceAll(sCnpj,"-","")
	sCnpj = strings.ReplaceAll(sCnpj,"/","")
	sCnpj = strings.ReplaceAll(sCnpj,".","")
    
	vCnpj := strings.Split(sCnpj,"")
	if (len(vCnpj) != 14) {
		return false
	} else {
		sum:=0
		va:=0
		j:=5
		for i:=0;i<=11; i++ {
			va,_=strconv.Atoi(vCnpj[i])				
			sum=sum+(va*j)
			j--	
			if j == 1 {
				j=9
			}		
		}
		dig1:=11-(sum%11)	
		if dig1>=10{
			dig1=0
		}	
		sum=0
		j=6
		for i:=0;i<=11; i++ {
			va,_=strconv.Atoi(vCnpj[i])
			sum=sum+va*j
			j--
			if j == 1 {
				j=9
			}	
		}
        sum= sum+(dig1*2)
		dig2:=11-(sum%11)	
		if dig2>=10{
			dig2=0
		}	
		val1,_:=strconv.Atoi(vCnpj[12])
		val2,_:=strconv.Atoi(vCnpj[13])
		
		return (dig1 == val1) && (dig2 == val2)	
		

	}
			
}
