# Тестовое задание в work-mate

## Запуск сервера
Порт `8080`

1. Через `docker-compose` в корне проекта
```bash
docker compose up -d
```
1. Сборка с помощью `make`
```bash
make rebuild
```
1. Запуск вручную
```bash
go run ./cmd/task/main.go
```

### Дополнительно можно включить pprof для проверки течи горутин:
1. Расскоментировать порт в `docker-compose.yaml`
2. Расскоментировать запуск pprof в go-коде `cmd/task/main.go`
3. Перейти на http://localhost:6060/debug/pprof/goroutine?debug=1

## Как использовать
- POST `/task`
Старт новой задачи
```json
{
    "id": "607b6a8c-ca0e-4d2a-b48b-1f7dd8926e00",
    "created_at": "2025-06-23T09:21:18.261801666Z",
    "duration": 0
}
```
- DELETE `/task/:id`
Отмена задачи
```json
status code 204
```
- GET `/task/:id/result`
Получение результата задачи
```json
{
    "result": "io work result = 0"
}
```
Или если задача на этапе выполнения
```json
{
    "problem": "task running or failed, no result, check status"
}
```
- GET `/task/:id/status`
Получение статуса задачи. Всего 4 статуса: `created, but waiting to start`, `in process`, `completed`, `failed by io bound`
```json
{
    "status": "in process",
    "created_at": "2025-06-23T09:21:18.261801666Z",
    "duration": 6	// время в секундах динамическое
}
```
- GET `/task`
Получение списка всех задач. По ТЗ не было, добавил для комфортной проверки работы сервиса
Возвращает массив задач
```json
[
    {
        "id": "13638b53-7021-4302-ae02-e89950ca44ae",
        "result": "io work result = 7",
        "status": "completed",
        "created_at": "2025-06-23T09:04:24.066769747Z",
        "duration": 201
    },
    {
        "id": "94989b6b-2013-49bf-919b-5c9564062eea",
        "result": null,
        "status": "failed by io bound",
        "created_at": "2025-06-23T09:04:30.121856784Z",
        "duration": 289
    }
	// ...
]
```

## Архитектура
- `internal/handlers` Обработчики http запросов;
- `internal/service` Содержит сервис запросов и executer, задача которого контролировать работу I/O задач;
- `internal/repository` Данный репозиторий реализован под fake хранилище, так что в случае замены на реальную БД, необходимо изменить логику;
- I/O задача-заглушка в `internal/service/executor/hardIOboundwork.go`. Задача выполняется строго 3-5 минут;
- Семафор ограничивает кол-во одновременно исполняемых задач.