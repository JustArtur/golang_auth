## Запуск локально

### Создать и наполнить `.env.dev`
```bash
touch auth_app/.env.dev
```
#### Пример данных

```
mySigningKey=golang_auth_password
mySalt=golangAuthSalt

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=example
DB_NAME=golang_auth_delopment
```

### Запуск локального сервера
```bash
go install github.com/zzwx/fresh@latest
fresh
```

## Запуск в докере

### Создать `.env.docker`

```bash
touch auth_app/.env.docker
```

#### Пример переменных окружения
```
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=example
DB_NAME=golang_auth_development
POSTGRES_USER=postgres
POSTGRES_PASSWORD=example
POSTGRES_DB=golang_auth_development

mySigningKey=golang_auth_password
mySalt=golangAuthSalt
```

### Заупск контейнеров

```bash
docker-compose up --build
```