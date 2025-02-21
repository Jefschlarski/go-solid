# Todo List API

Uma API para gerenciamento de tarefas (todo list) construída com Go, seguindo princípios SOLID e utilizando padrões de projeto.

## 🚀 Tecnologias

- Go 1.23+
- Echo Framework
- SQLite
- Swagger
- Docker

## 🐳 Docker

O projeto pode ser executado facilmente usando Docker e Docker Compose.

### Pré-requisitos
- Docker

### Executando com Docker

1. Construa e inicie os containers:
```bash
docker-compose up --build
```

2. Para executar em background:
```bash
docker-compose up -d
```

3. Para parar os containers:
```bash
docker-compose down
```

### Acessando a Aplicação
- API: http://localhost:8080
- Swagger UI: http://localhost:8080/swagger/index.html

### Estrutura Docker
- **Multi-stage build**: Otimiza o tamanho da imagem final
- **Swagger automático**: Geração automática da documentação
- **Persistência**: Volume Docker para dados do SQLite
- **Ambiente isolado**: Todas as dependências são containerizadas

### Variáveis de Ambiente
```env
DB_PATH=./todos.db
SERVER_PORT=8080
ENVIRONMENT=development
```

## 🔧 Instalação Local

É possível executar a aplicação localmente sem utilizar o Docker.

### Pré-requisitos
- Go 1.23 ou superior
- SQLite

1. Clone o repositório:
```bash
git clone https://github.com/Jefschlarski/go-solid.git
```

2. Instale as dependências:
```bash
go get -u github.com/labstack/echo/v4
go get -u github.com/swaggo/echo-swagger
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/mattn/go-sqlite3
go get -u github.com/joho/godotenv
```
Ou use o `go mod tidy` para instalar as dependências automaticamente.

3. Gere a documentação Swagger:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$PATH:$(go env GOPATH)/bin
swag init -g cmd/main.go
```

4. Execute a aplicação:
```bash
go run cmd/main.go
```

## 📖 Documentação da API

A API estará disponível em `http://localhost:8080` 

e a documentação Swagger em `http://localhost:8080/swagger/index.html`

### Endpoints

#### Criar Tarefa
```http
POST /api/todos

{
    "title": "Fazer compras",
    "description": "Comprar leite e pão"
}
```

#### Listar Tarefas
```http
GET /api/todos
```

#### Atualizar Status da Tarefa
```http
PATCH /api/todos/{id}/status

0: PENDING, 1: IN_PROGRESS, 2: PAUSED, 3: COMPLETED, 4: CANCELED
{
    "status": 1  
}
```

#### Adicionar Tempo Gasto
```http
PATCH /api/todos/{id}/time

{
    "minutes": 30
}
```

### Exemplos de Respostas

#### Tarefa Individual
```json
{
    "id": 1,
    "title": "Fazer compras",
    "description": "Comprar leite e pão",
    "status": "IN_PROGRESS",
    "time_spent": 30,
    "created_at": "15/03/2024 10:00",
    "updated_at": "15/03/2024 10:30",
    "completed_at": null
}
```

## 🎯 Padrões de Projeto Utilizados

### 1. State Pattern
Utilizado para gerenciar diferentes estados da tarefa.

```go
// Interface que define o contrato para diferentes estados
type ITodoState interface {
    ChangeStatus(todo *Todo, newStatus TodoStatus) error
    AddTimeSpent(todo *Todo, minutes int64) error
}

// Implementação para tarefas em progresso
type inProgressState struct{}

func (s *inProgressState) ChangeStatus(todo *Todo, newStatus TodoStatus) error {
    if newStatus != StatusCompleted && newStatus != StatusCanceled {
        return errors.New("tarefa em progresso só pode ser completada ou cancelada")
    }
    // ... lógica de mudança de status
}
```

### 2. Repository Pattern
Isola a lógica de acesso a dados do resto da aplicação.

```go
type ITodoRepository interface {
    Create(todo *model.Todo) error
    GetAll() ([]model.Todo, error)
    GetByID(id int) (*model.Todo, error)
    Update(todo *model.Todo) error
}
```

### 3. DTO (Data Transfer Object)
Separa a representação de dados externa da estrutura interna.

```go
// DTO para criação de tarefas
type CreateTodoDTO struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

// DTO para retorno de tarefas
type ReturnTodoDTO struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Status      string  `json:"status"`
    TimeSpent   int64   `json:"time_spent"`
    CreatedAt   string  `json:"created_at"`
    UpdatedAt   string  `json:"updated_at"`
    CompletedAt string  `json:"completed_at"`
}
```

## 🔄 SOLID

O projeto segue os princípios SOLID:

### S - Single Responsibility Principle (Princípio da Responsabilidade Única)
- **Implementação**: Cada componente tem uma única responsabilidade bem definida:
  - `todoRepository`: Responsável apenas pela persistência dos dados
  - `todoService`: Responsável pela lógica de negócios
  - `todoHandler`: Responsável pelo tratamento das requisições HTTP
  - `model.Todo`: Responsável pela estrutura de dados da tarefa

### O - Open/Closed Principle (Princípio Aberto/Fechado)
- **Implementação**: O padrão Strategy usado para gerenciar estados dos todos:
  ```go
  todoState := model.GetTodoState(todo.Status)
  if err := todoState.ChangeStatus(todo, newStatus); err != nil {
      return nil, err
  }
  ```
- Novas estratégias de estado podem ser adicionadas sem modificar o código existente
- Novos comportamentos podem ser adicionados criando novas implementações da interface

### L - Liskov Substitution Principle (Princípio da Substituição de Liskov)
- **Implementação**: As interfaces são respeitadas em suas implementações:
  - `ITodoRepository` pode ser implementada por qualquer estrutura que respeite seu contrato
  - `ITodoService` mantém a consistência de comportamento esperado
- As implementações concretas (`todoRepository` e `todoService`) podem ser substituídas por qualquer outra implementação das interfaces sem quebrar o funcionamento

### I - Interface Segregation Principle (Princípio da Segregação de Interface)
- **Implementação**: Interfaces coesas e específicas:
  ```go
  type ITodoRepository interface {
      Create(todo *model.Todo) error
      GetAll() ([]model.Todo, error)
      GetByID(id int) (*model.Todo, error)
      Update(todo *model.Todo) error
  }
  
  type ITodoService interface {
      CreateTodo(todo *model.Todo) error
      GetAllTodos() ([]model.Todo, error)
      UpdateTodoStatus(id int, newStatus model.TodoStatus) (*model.Todo, error)
      AddTimeSpent(id int, minutes int64) (*model.Todo, error)
  }
  ```
- Cada interface declara apenas os métodos necessários para seu propósito específico

### D - Dependency Inversion Principle (Princípio da Inversão de Dependência)
- **Implementação**: Dependências são injetadas e baseadas em abstrações:
  ```go
  func NewTodoService(repo repository.ITodoRepository) ITodoService {
      return &todoService{repo: repo}
  }
  ```
- O `main.go` atua como composição raiz:
  ```go
  repo := repository.NewTodoRepository(db)
  service := service.NewTodoService(repo)
  handler := handler.NewTodoHandler(service)
  ```
- As dependências são invertidas, com as implementações concretas dependendo de abstrações

## 🏗️ Arquitetura

O projeto segue uma arquitetura em camadas:

```
internal/
├── handler/     # Manipuladores HTTP
├── service/     # Lógica de negócios
├── repository/  # Acesso a dados
├── model/       # Entidades de domínio
└── dto/         # Objetos de transferência de dados
```

## 🔄 Fluxo de Estados

Uma tarefa pode ter os seguintes estados:
- `PENDING`: Estado inicial
- `IN_PROGRESS`: Tarefa em andamento
- `PAUSED`: Tarefa pausada
- `COMPLETED`: Tarefa concluída
- `CANCELED`: Tarefa cancelada

Transições permitidas:
- `PENDING` → `IN_PROGRESS` ou `CANCELED`
- `IN_PROGRESS` → `PAUSED` ou `COMPLETED` ou `CANCELED`
- `PAUSED` → `IN_PROGRESS` ou `COMPLETED` ou `CANCELED`
- `COMPLETED` → (estado final)
- `CANCELED` → (estado final)

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
