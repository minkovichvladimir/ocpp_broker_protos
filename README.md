# OCPP Broker Protos

Этот модуль содержит определения Protocol Buffers и сгенерированный код для gRPC сервисов, используемых в проекте OCPP Broker.

## Структура проекта

```
ocpp_broker_protos/
├── proto/
│   ├── ocpp.proto         # Определения протокола
│   ├── ocpp.pb.go         # Сгенерированный код для сообщений
│   └── ocpp_grpc.pb.go    # Сгенерированный код для gRPC сервиса
├── Makefile
├── go.mod
└── README.md
```

## Требования

- Protocol Buffers (protoc) версии 3
- Go плагины для protoc:
  - protoc-gen-go
  - protoc-gen-go-grpc

## Установка

1. Установите необходимые инструменты:
```bash
# Установка protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Установка protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. Сгенерируйте код:
```bash
make proto
```

## Использование

Этот модуль используется другими сервисами проекта. Для подключения к вашему сервису:

1. Добавьте зависимость в go.mod:
```go
require github.com/vladimirminkovich/rtk_volt/ocpp_broker_protos v0.0.0
replace github.com/vladimirminkovich/rtk_volt/ocpp_broker_protos => ../ocpp_broker_protos
```

2. Импортируйте сгенерированный код:
```go
import pb "github.com/vladimirminkovich/rtk_volt/ocpp_broker_protos/proto"
```

## API

### OCPPService

Сервис для обмена OCPP сообщениями:

```protobuf
service OCPPService {
  rpc SendOCPPMessage (OCPPMessageRequest) returns (OCPPMessageResponse) {}
}
```

#### Сообщения

- OCPPMessageRequest:
  - station_id (string): Идентификатор зарядной станции
  - message (bytes): Тело OCPP сообщения
  - message_type (string): Тип OCPP сообщения

- OCPPMessageResponse:
  - success (bool): Статус обработки сообщения

## Разработка

При внесении изменений в proto файлы:

1. Измените файлы в директории `proto/`
2. Запустите `make proto` для регенерации кода
3. Обновите зависимости в сервисах, использующих этот модуль 