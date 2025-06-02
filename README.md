# Lanchonete API - Tech Challenge 2

## Visão Geral

Este projeto implementa uma API para gerenciamento de pedidos de uma lanchonete, seguindo princípios de Clean Architecture. O sistema é composto por múltiplos módulos organizados em camadas, facilitando manutenção, testes e escalabilidade.

---

## Estrutura de Diretórios

- **bootstrap/**  
  Inicialização da aplicação, configuração de ambiente e injeção de dependências.

- **db/**  
  Scripts de inicialização do banco de dados (ex: `init.sql`).

- **docs/**  
  Documentação Swagger gerada automaticamente.

- **infra/**  
  Implementações de infraestrutura, como repositórios de banco de dados.

- **internal/**  
  Código de domínio e regras de negócio:
  - **application/**: Casos de uso e presenters.
  - **domain/**: Entidades e interfaces de repositórios.
  - **infrastructure/**: Implementações específicas de infraestrutura.
  - **interfaces/**: Camada de entrada (HTTP, handlers, rotas).

- **usecases/**  
  Casos de uso específicos.

- **main.go**  
  Ponto de entrada da aplicação.

---

## Esquema do Banco de Dados

Exemplo simplificado das principais tabelas:

```sql
CREATE TABLE Cliente (
    cpfCliente VARCHAR(11) PRIMARY KEY,
    nomeCliente VARCHAR(100),
    emailCliente VARCHAR(100)
);

CREATE TABLE Produto (
    idProduto INT AUTO_INCREMENT PRIMARY KEY,
    nomeProduto VARCHAR(100),
    descricaoProduto TEXT,
    precoProduto FLOAT,
    personalizacaoProduto VARCHAR(255),
    categoriaProduto VARCHAR(50)
);

CREATE TABLE Pedido (
    idPedido INT AUTO_INCREMENT PRIMARY KEY,
    cliente VARCHAR(11),
    totalPedido FLOAT,
    tempoEstimado VARCHAR(8),
    status VARCHAR(50),
    statusPagamento VARCHAR(50),
    FOREIGN KEY (cliente) REFERENCES Cliente(cpfCliente)
);

CREATE TABLE Pedido_Produto (
    idPedido INT,
    idProduto INT,
    quantidade INT,
    FOREIGN KEY (idPedido) REFERENCES Pedido(idPedido),
    FOREIGN KEY (idProduto) REFERENCES Produto(idProduto)
);

CREATE TABLE Pagamento (
    idPagamento INT AUTO_INCREMENT PRIMARY KEY,
    dataCriacao DATETIME,
    Status VARCHAR(50),
    idPedido INT,
    FOREIGN KEY (idPedido) REFERENCES Pedido(idPedido)
);

CREATE TABLE Acompanhamento (
    idAcompanhamento INT AUTO_INCREMENT PRIMARY KEY,
    tempoEstimado VARCHAR(8)
);

CREATE TABLE Acompanhamento_Pedido (
    idAcompanhamento INT,
    idPedido INT,
    FOREIGN KEY (idAcompanhamento) REFERENCES Acompanhamento(idAcompanhamento),
    FOREIGN KEY (idPedido) REFERENCES Pedido(idPedido)
);
```
## Como Rodar a Aplicação

### Pré-requisitos

- Go 1.20+
- Docker e Docker Compose (opcional, recomendado)
- MySQL

### Passos

1. **Clone o repositório:**
   ```sh
   git clone <repo-url>
   cd CleanArch
   ```

2. **Configure variáveis de ambiente:**
   - Edite o arquivo `.env` conforme necessário.

3. **Suba o banco de dados (opcional):**
   ```sh
   docker-compose up -d
   ```

4. **Instale as dependências:**
   ```sh
   go mod tidy
   ```

5. **Gere a documentação Swagger:**
   ```sh
   swag init
   ```

6. **Rode a aplicação:**
   ```sh
   go run main.go
   ```

7. **Acesse a documentação:**
   - [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html)

---

## Observações

- O formato recomendado para datas é: `"2025-03-19T12:34:56Z"`
- Para rodar os testes, utilize:
  ```sh
  go test ./...
  ```

---

## Licença

Este projeto é parte de um desafio técnico e possui fins educacionais.
