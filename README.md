# gobase

CLI em Go que gera a estrutura base de um serviço HTTP em segundos.

Ao rodar, ele pergunta o nome do app e o module path, e cria um projeto
pronto para compilar e subir, com as camadas essenciais separadas:
`main`, `server`, `middleware`, `handler` e `service`.

## O que ele gera

```
meu-servico/
├── cmd/api/main.go
├── internal/
│   ├── handler/handler.go
│   ├── service/service.go
│   ├── middleware/logging.go
│   └── server/server.go
└── go.mod
```

O projeto gerado expõe o endpoint `GET /hello?name=fulano` na porta `8080`.

## Requisitos

- Go 1.22 ou superior
- `make` (opcional, usado apenas no projeto gerado)

## Uso

```bash
gobase
```

```
App name? (app): meu-servico
Module path? (github.com/you/meu-servico):
```

- Enter em branco no nome usa `app`.
- Enter em branco no module path usa `github.com/you/<nome>`.
- Se a pasta já existir, o `gobase` aborta sem sobrescrever nada.

## Build

```bash
go build -o gobase .
```

Binário menor (sem símbolos de debug):

```bash
go build -ldflags="-s -w" -o gobase .
```

## Instalação no Linux

```bash
go build -o gobase .
sudo install -m 755 gobase /usr/local/bin/gobase
```

Depois, em qualquer pasta:

```bash
gobase
```

Os templates são embutidos no binário (`go:embed`), então não é preciso
manter a pasta `templates/` no sistema.

## Atualizar

Repita os comandos de instalação. O binário antigo é sobrescrito.

## Desinstalar

```bash
sudo rm /usr/local/bin/gobase
```

## Rodando o projeto gerado

```bash
cd meu-servico
make run          # ou: go run ./cmd/api
```

```bash
curl "localhost:8080/hello?name=Ana"
# hello, Ana
```