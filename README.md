# Importador de Dados Go

## Features

- Importar dados de um arquivo TXT para o banco de dados
- Criação do branco de dados dentro da aplicação(função createDataBase()) 
- Uso do docker para configuração do ambiente

## Tecnologia
- [GoLang] 
- [Postgres] 

## Procedimentos

Primeiramente baixar do git os arquivos do projeto.
Rodar dentro da pasta do arquivos os comandos para subir o docker.

```sh
docker-compose up -d --build
```

Na sequencia rodar a aplicacao...

```sh
docker logs -f go
```

Após rodar a aplicação acessar o banco de dados (o banco de dados é criado pela aplicação)...

```sh
docker exec -it postgres psql -U user -d desafio
```
Nesse momento estará dentro do banco de dados, é possivel rodar um select diretamente dentro dele
```sh
select * from dados;
ou
select count(cpf) from dados;

```
## Banco de Dados

Script que o arquivo main.go roda para criar o banco de dados 
```sh
CREATE TABLE DADOS 
(CPF character varying, PRIVATE char, INCOMPLETO char, DATA_COMPRA date, TICKET_MEDIO double precision, TICKET_ULTIMA double precision, LOJA_FREQUENCIA character varying, ULTIMA_LOJA  character varying);
```

| CPF | PRIVATE | INCOMPLETO | DATA_COMPRA | TICKET_MEDIO | TICKET_ULTIMA | LOJA_FREQUENCIA | ULTIMA_LOJA |

###Referencias

[GoLang](https://go.dev/)

[Canal FullCycle](https://www.youtube.com/c/FullCycle)

[Canal Jason Cheung](https://www.youtube.com/channel/UCD2Mv8gCmF2kX4yS1zG7UsQ/featured)

