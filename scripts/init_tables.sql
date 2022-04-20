--Schemas

CREATE SCHEMA IF NOT EXISTS desafio;

drop table if exists desafio.dados;

CREATE table if not exists desafio.dados (
    CPF character varying,
    IS_PRIVATE char,
    INCOMPLETO char,
    DATA_COMPRA date,
    TICKET_MEDIO double precision,
    TICKET_ULTIMA double precision,
    LOJA_FREQUENCIA character varying,
    ULTIMA_LOJA  character varying
);