# Desafio Backend GoLang Jr

## Descrição do Projeto
Este projeto é uma API REST para gerenciar livros de usuários. Com esta API, você pode:
- Cadastrar um livro
- Atualizar um livro
- Listar um livro específico
- Listar todos os livros
- Deletar um livro

Os livros possuem os seguintes atributos:
- **Título**
- **Categoria**
- **Autor**
- **Sinopse**

## Tecnologias Utilizadas
- **Linguagem:** GoLang
- **Banco de Dados:** MongoDB
- **Gerenciamento de Ambiente:** Variáveis de ambiente com `godotenv`
- **Containerização:** Docker e Docker Compose
- **Gerenciamento de Dependências:** Go Modules

## Estrutura do Projeto
```
.
├── config           
│   └── config.go
├── controllers       
│   └── book.go
├── database          
│   └── mongodb.go
├── models           
│   └── book.go
├── routes            
│   └── routes.go
├── main.go          
├── docker-compose.yml
├── Dockerfile
├── .env              
└── go.mod            
```

## Pré-requisitos
1. **Docker e Docker Compose** instalados.
2. **GoLang** instalado (versão 1.17 ou superior).
3. Um arquivo `.env` configurado conforme o exemplo abaixo.

## Configuração do Arquivo `.env`
Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:
```env
MONGODB_URI=mongodb://db:27017
MONGODB_DB=books
MONGODB_COLLECTION=books
```

## Instruções para Executar o Projeto

### Usando Docker
1. **Construa e inicie os containers:**
   ```bash
   docker-compose up
   ```
2. Acesse a aplicação em `http://localhost:8080`.

### Sem Docker
1. **Instale as dependências:**
   ```bash
   go mod tidy
   ```
2. **Inicie o MongoDB localmente:** Certifique-se de que um servidor MongoDB esteja em execução e configurado na URL definida em `MONGODB_URI`.
3. **Execute a aplicação:**
   ```bash
   go run main.go
   ```
4. Acesse a aplicação em `http://localhost:8080`.

## Endpoints da API

### Base URL
`http://localhost:8080`

### Rotas Disponíveis

#### 1. **Criar Livro**
- **URL:** `/books`
- **Método:** `POST`
- **Corpo da Requisição:**
  ```json
  {
    "title": "Livro Exemplo",
    "category": "Ficção",
    "author": "Autor Exemplo",
    "synopsis": "Esta é a sinopse do livro."
  }
  ```
- **Resposta de Sucesso:** `201 Created`

#### 2. **Atualizar Livro**
- **URL:** `/books/{id}`
- **Método:** `PUT`
- **Corpo da Requisição:**
  ```json
  {
    "title": "Livro Atualizado",
    "category": "Drama",
    "author": "Novo Autor",
    "synopsis": "Sinopse atualizada."
  }
  ```
- **Resposta de Sucesso:** `200 OK`

#### 3. **Listar Livro Específico**
- **URL:** `/books/{id}`
- **Método:** `GET`
- **Resposta de Sucesso:**
  ```json
  {
    "id": "1",
    "title": "Livro Exemplo",
    "category": "Ficção",
    "author": "Autor Exemplo",
    "synopsis": "Esta é a sinopse do livro."
  }
  ```

#### 4. **Listar Todos os Livros**
- **URL:** `/books`
- **Método:** `GET`
- **Resposta de Sucesso:**
  ```json
  [
    {
      "id": "1",
      "title": "Livro Exemplo",
      "category": "Ficção",
      "author": "Autor Exemplo",
      "synopsis": "Esta é a sinopse do livro."
    },
    {
      "id": "2",
      "title": "Outro Livro",
      "category": "Aventura",
      "author": "Outro Autor",
      "synopsis": "Sinopse de outro livro."
    }
  ]
  ```

#### 5. **Deletar Livro**
- **URL:** `/books/{id}`
- **Método:** `DELETE`
- **Resposta de Sucesso:** `200 Book deleted successfully`

## Justificativa do Banco de Dados
- **Estrutura Flexível:** Como o desafio não exige a implementação de relacionamentos complexos, o MongoDB, sendo um banco de dados NoSQL, oferece uma estrutura de dados flexível. Ele permite armazenar documentos JSON, o que é ideal para esse tipo de aplicação onde os dados (livros) podem ter diferentes atributos, como título, categoria, autor e sinopse, sem a necessidade de tabelas rígidas ou chaves estrangeiras.

- **Escalabilidade:** O MongoDB é conhecido por sua alta escalabilidade, permitindo que a aplicação cresça facilmente conforme a quantidade de livros ou usuários aumente. Isso é vantajoso em sistemas que podem precisar de performance em grandes volumes de dados.

- **Desempenho:** Em casos onde não há a necessidade de realizar consultas com joins complexos, o MongoDB pode oferecer um desempenho melhor, já que ele não precisa fazer operações de junção entre tabelas, como no caso de bancos relacionais, o que pode reduzir a sobrecarga de processamento.

- **Desenvolvimento Ágil:** A flexibilidade do MongoDB também facilita o desenvolvimento rápido e a adaptação a mudanças no modelo de dados, o que é vantajoso em projetos que precisam ser entregues rapidamente, como esse desafio.

## Informações Adicionais
- **Escolha do MongoDB:** Optei pelo MongoDB devido à sua flexibilidade, já que o desafio não exigia relacionamentos complexos entre dados, e sua estrutura de documentos JSON permite um modelo de dados mais ágil e escalável.
- **Estrutura do Projeto:** Defini a estrutura de diretórios de forma a organizar claramente as responsabilidades de cada parte da aplicação (configuração, controladores, modelos e rotas).
- **Uso de Docker:** Containerizei a aplicação para garantir que ela funcionasse de forma consistente em diferentes ambientes, escolhendo uma imagem Docker leve (golang:alpine) para otimizar o desempenho.
- **Variáveis de Ambiente:** Utilizei o godotenv para gerenciar variáveis de ambiente de maneira segura e prática, centralizando a configuração em um arquivo .env.
- **Docker Compose:** Implementei o Docker Compose para simplificar a orquestração da aplicação e do banco de dados, facilitando o processo de execução do projeto em diferentes ambientes.
- **Gin:** Optei pelo Gin-Gonic devido à sua simplicidade, desempenho e facilidade de uso para criar rotas e gerenciar requisições. 

## Como Contribuir
1. Faça um fork do repositório.
2. Crie um branch para sua feature (`git checkout -b feature/nome-da-feature`).
3. Faça commit das suas alterações (`git commit -m 'Adiciona nova feature'`).
4. Faça push para o branch (`git push origin feature/nome-da-feature`).
5. Abra um Pull Request.

---

Projeto desenvolvido como parte do desafio técnico para Backend GoLang Jr na Taghos.

