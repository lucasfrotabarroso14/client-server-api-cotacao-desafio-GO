# Desafio: Aplicação de Webserver HTTP, Contextos, Banco de Dados e Manipulação de Arquivos com Go

Aplicações: webserver HTTP, context, banco de dados e manipulação de arquivos com Go.

## Descrição

Serviços:

- `client.go`
- `server.go`

### Requisitos

1. **client.go**
   - Deverá realizar uma requisição HTTP para `server.go` solicitando a cotação do dólar.
   - Receber do `server.go` apenas o valor atual do câmbio (campo "bid" do JSON).
   - Usar o package `context` com timeout máximo de 300ms para receber o resultado de `server.go`.
   - Salvar a cotação atual em um arquivo `cotacao.txt` no formato: `Dólar: {valor}`.
   - Retornar erro nos logs caso o tempo de execução seja insuficiente.

2. **server.go**
   - Consumir a API de câmbio de Dólar e Real no endereço: [https://economia.awesomeapi.com.br/json/last/USD-BRL](https://economia.awesomeapi.com.br/json/last/USD-BRL).
   - Retornar o resultado no formato JSON para o cliente.
   - Usar o package `context` para registrar no banco de dados SQLite cada cotação recebida.
     - Timeout máximo para chamar a API de cotação do dólar: 200ms.
     - Timeout máximo para persistir os dados no banco: 10ms.
   - Criar o endpoint `/cotacao` e usar a porta 8080.
   - Retornar erro nos logs caso o tempo de execução seja insuficiente.

### Estrutura do Projeto

#### Servidor (`server.go`)

- [X] Fazer função que vai consumir a API e retornar a cotação do dólar.
- **Contexto do Servidor:**
  - [X] Criar o `context` do servidor.
  - [X] Timeout máximo para chamar a API: 200ms.
  - [X] Timeout máximo para persistir os dados no banco: 10ms.
- [X] Endpoint: `/cotacao` e porta: 8080.

#### Cliente (`client.go`)

- [X] Fazer a requisição GET para o servidor.
- [X] Salvar a cotação atual em um arquivo `cotacao.txt` (no formato: `Dólar: {valor}`).
- **Contexto do Cliente:**
  - [X] Criar o `context` do cliente.
  - [X] Timeout máximo de 300ms para receber o resultado do `server.go`.

---

Sinta-se à vontade para clonar o repositório e contribuir com melhorias!

