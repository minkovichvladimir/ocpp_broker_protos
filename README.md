# OCPP Broker Protos

Этот модуль содержит определения Protocol Buffers и сгенерированный код для gRPC сервисов, используемых в проекте OCPP Broker.



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
    centralpb "ocpp_broker_protos.central_system.v1"
    connectionpb "ocpp_broker_protos.connection.v1"
    mappingpb "ocpp_broker_protos.mapping.v1"
    transactionpb "ocpp_broker_protos.transaction.v1"
)
```

## API

### CentralSystemService

Сервис для управления центральными системами:

```protobuf
service CentralSystemService {
  rpc CreateCentralSystem(CreateCentralSystemRequest) returns (CreateCentralSystemResponse) {}
  rpc GetCentralSystem(GetCentralSystemRequest) returns (GetCentralSystemResponse) {}
  rpc ListCentralSystems(ListCentralSystemsRequest) returns (ListCentralSystemsResponse) {}
  rpc UpdateCentralSystem(UpdateCentralSystemRequest) returns (UpdateCentralSystemResponse) {}
  rpc DeleteCentralSystem(DeleteCentralSystemRequest) returns (DeleteCentralSystemResponse) {}
}
```

#### Сообщения

- CentralSystem:
  - id (int32): Идентификатор центральной системы
  - name (string): Название центральной системы
  - websocket_url (string): URL для WebSocket соединения
  - description (string): Описание центральной системы

- CreateCentralSystemRequest/Response, GetCentralSystemRequest/Response, UpdateCentralSystemRequest/Response, DeleteCentralSystemRequest/Response:
  - Соответствующие запросы и ответы для операций CRUD

### ConnectionService

Сервис для управления соединениями станций с центральными системами:

```protobuf
service ConnectionService {
  rpc ListWorkers(ListWorkersRequest) returns (ListWorkersResponse);
  rpc GetWorker(GetWorkerRequest) returns (GetWorkerResponse);
  rpc GetWorkerConnections(GetWorkerConnectionsRequest) returns (GetWorkerConnectionsResponse);
  rpc CreateConnection(CreateConnectionRequest) returns (CreateConnectionResponse);
  rpc RecreateConnection(RecreateConnectionRequest) returns (RecreateConnectionResponse);
  rpc DisconnectStation(DisconnectStationRequest) returns (DisconnectStationResponse);
}
```

#### Сообщения

- WorkerInfo:
  - central_system_id (uint32): Идентификатор центральной системы
  - central_system_name (string): Название центральной системы
  - connections_count (int32): Количество соединений
  - connections (repeated ConnectionInfo): Список соединений

- ConnectionInfo:
  - station_id (string): Идентификатор зарядной станции
  - central_system_id (uint32): Идентификатор центральной системы
  - connected (bool): Статус соединения

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

### TransactionService

Сервис для управления транзакциями зарядных станций:

```protobuf
service TransactionService {
  rpc GetAllTransactions(GetAllTransactionsRequest) returns (GetAllTransactionsResponse) {}
}
```

#### Сообщения

- Transaction:
  - id (string): Уникальный идентификатор транзакции
  - station_id (string): Идентификатор зарядной станции
  - connector_id (string): Идентификатор коннектора
  - central_system_id (int32): Идентификатор центральной системы
  - started_at (google.protobuf.Timestamp): Время начала транзакции
  - ended_at (google.protobuf.Timestamp): Время окончания транзакции (может быть null)

- GetAllTransactionsRequest:
  - limit (int32): Ограничение количества записей
  - offset (int32): Смещение для пагинации

- GetAllTransactionsResponse:
  - transactions (repeated Transaction): Список транзакций

## Разработка

При внесении изменений в proto файлы:

1. Измените файлы в директории `proto/`
2. Запустите `make generate` для регенерации кода (файлы будут сгенерированы в `gen/go/`)
3. Обновите зависимости в сервисах, использующих этот модуль 