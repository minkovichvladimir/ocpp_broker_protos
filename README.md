# OCPP Broker Protos

Этот модуль содержит определения Protocol Buffers и сгенерированный код для gRPC сервисов, используемых в проекте OCPP Broker.

## Структура проекта

```
ocpp_broker_protos/
├── proto/
│   ├── ocpp/
│   │   └── v1/
│   │       └── ocpp.proto         # Определения протокола OCPP
│   └── mapping/
│       └── v1/
│           └── mapping.proto      # Определения протокола для маппингов
├── gen/
│   └── go/
│       ├── ocpp/
│       │   └── v1/
│       │       ├── ocpp.pb.go         # Сгенерированный код для OCPP сообщений
│       │       └── ocpp_grpc.pb.go    # Сгенерированный код для OCPP gRPC сервиса
│       └── mapping/
│           └── v1/
│               ├── mapping.pb.go      # Сгенерированный код для маппингов
│               └── mapping_grpc.pb.go # Сгенерированный код для сервиса маппингов
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
make generate
```

## Использование

Этот модуль используется другими сервисами проекта. Для подключения к вашему сервису:

1. Добавьте зависимость в go.mod:
```go
require github.com/minkovichvladimir/ocpp_broker_protos v0.0.0
replace github.com/minkovichvladimir/ocpp_broker_protos => ../ocpp_broker_protos
```

2. Импортируйте сгенерированный код:
```go
import (
    ocpppb "github.com/minkovichvladimir/ocpp_broker_protos/gen/go/ocpp/v1"
    mappingpb "github.com/minkovichvladimir/ocpp_broker_protos/gen/go/mapping/v1"
)
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

### MappingService

Сервис для управления маппингами станций:

```protobuf
service MappingService {
  rpc GetAllMappings(GetAllMappingsRequest) returns (GetAllMappingsResponse) {}
  rpc CreateMapping(CreateMappingRequest) returns (CreateMappingResponse) {}
  rpc UpdateMapping(UpdateMappingRequest) returns (UpdateMappingResponse) {}
  rpc DeleteMapping(DeleteMappingRequest) returns (DeleteMappingResponse) {}
}
```

#### Сообщения

- Mapping:
  - station_id (string): Идентификатор зарядной станции
  - central_system_id (int32): Идентификатор центральной системы
  - weight (int32): Вес маппинга
  - is_master (bool): Флаг мастер-станции
  - is_enabled (bool): Флаг активности

## Разработка

При внесении изменений в proto файлы:

1. Измените файлы в директории `proto/`
2. Запустите `make generate` для регенерации кода (файлы будут сгенерированы в `gen/go/`)
3. Обновите зависимости в сервисах, использующих этот модуль 