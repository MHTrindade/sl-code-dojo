# sl-code-dojo

## Visão Geral

Este projeto é uma API de exemplo desenvolvida durante uma cerimônia de DOJO utilizando Golang, o framework Fiber e o ORM GORM. A aplicação gerencia registros de pessoas e seus contatos, utilizando o MySQL como banco de dados. O projeto foi estruturado para ser executado em containers Docker, facilitando o desenvolvimento e a implantação.

## Funcionalidades

- Criar registros de pessoas (com ou sem contatos).
- Endpoints RESTful para criação e recuperação de registros.
- Integração com um banco de dados MySQL.
- Migração automática do esquema utilizando `AutoMigrate` do GORM.

## Pré-requisitos

- [Docker](https://www.docker.com/) (ou uma alternativa como Colima) instalado.
- [Golang](https://golang.org/) instalado (para desenvolvimento local).
- MySQL (fornecido via container Docker conforme instruções abaixo).

## Instruções de Configuração

### 1. Construindo e Executando o Container MySQL

Navegue até o diretório onde está localizado o `Dockerfile` do MySQL e execute os seguintes comandos:

```bash
docker build -t mysql-dojo .
```

Depois, inicie o container com:

```bash
docker run -d --name mysql -p 3306:3306 mysql-dojo
```

### 2. Criando o Banco de Dados

Após iniciar o container do MySQL, conecte-se ao banco de dados e execute:

```sql
CREATE DATABASE sl_dojo;
```

Isso criará o banco de dados necessário para a aplicação.

### 3. Executando a API

Com o banco de dados configurado, você pode iniciar a API rodando:

```bash
go run main.go
```

A aplicação será iniciada e ficará escutando requisições HTTP na porta `3000`.

## Exemplos de Uso da API

### Exemplo 1: Criar um Registro de Pessoa Sem Contatos

```bash
curl --location 'http://localhost:3000/people' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Rapha",
    "age": 50
}'
```

### Exemplo 2: Criar um Registro de Pessoa Com Contatos

```bash
curl --location 'http://localhost:3000/people' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Rapha",
    "age": 50,
    "contacts": [
        {
            "type": "email",
            "value": "mail@mail.com"
        }
    ]
}'
```

## Informações Adicionais

- A API inclui um endpoint simples de teste (`/`), que retorna `"Hello, World 👋!"`.
- A aplicação utiliza `AutoMigrate` do GORM para criar ou atualizar automaticamente o esquema do banco de dados com as tabelas `Person` e `Contact`.
- Certifique-se de que a configuração do DSN (Data Source Name) está correta para conectar-se ao seu banco de dados MySQL:
  ```
  root:root@tcp(127.0.0.1:3306)/sl_dojo?charset=utf8mb4&parseTime=True&loc=Local
  ```

## Solução de Problemas

- **Problemas com o Container Docker:**  
  Verifique se o container do MySQL está rodando:
  ```bash
  docker ps
  ```
  Confira os logs do container para verificar erros:
  ```bash
  docker logs mysql
  ```

- **Erro de Conexão com o Banco de Dados:**  
  Confirme que o banco de dados `sl_dojo` existe e que a string de conexão (DSN) está correta.

- **Conflitos de Porta:**  
  Verifique se a porta `3306` não está sendo usada por outro serviço em sua máquina.

## Licença

Este projeto é de código aberto. Sinta-se à vontade para modificar e utilizar conforme necessário.

