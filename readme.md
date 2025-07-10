# Мини пет проект по новым технологиям
[![Go Report Card](https://goreportcard.com/badge/github.com/eragon-mdi/ksu)](https://goreportcard.com/report/github.com/eragon-mdi/ksu)

## Запуск сервера
Порт `8080` - изменить в `.env`

1. Через `docker-compose` в корне проекта
```bash
# Запуск
docker compose build && docker compose up -d
# Завершение
docker compose down
```
2. Сборка с помощью `make`
```bash
make rebuild
make restart
```

### Конфиги в `config/config.yaml` 
```yaml
server:                     # Данные настройки не обязательны
  read_timeout: "5s"
  write_timeout: "10s"
  read_header_timeout: "2s"
  idle_timeout: "30s"       # Время простоя соединения, важно установить любое значение, иначе потекут горутины

logger:
  handler: "json"           # Формат логов: "json" или "text"
  level: "debug"            # debug, info, warn, error
  output_internal: false

app:
  semaphore: 16             # Кол-во одновременных задач (IO-bound)

clickhouse:                 # Чтобы писалиьс логи в clickhouse logger.output_internal = false
  address: clickhouse:9000
  username: ksu
  password: secret
  try_connection_perod: "3s"
  connection_attempts: 5
  batch_size: 10
  batch_inteval: "1m"

storage:    # Поддерживается in-memmory (fake) и sql хранилища, для добавления помимо postgres, добавить в фабрику storage
  type: postgres # internal | postgres
  host: db
  port: 5432
  user: app
  password: 123
  name: app_db
  ssl_mode: disable
  need_migrate: true
  migaret_src: file:///migrate
```

### Дополнительно можно использовать pprof для проверки течи горутин:
1. Склонировать ветку `dev` и работать из неё. В ветке `main` pprof не доступен
2. Перейти на http://localhost:6060/debug/pprof/goroutine?debug=1

## Как использовать. API
<details>
<summary>POST <code>/task - Старт новой задачи</code></summary>

```json
{
    "id": "607b6a8c-ca0e-4d2a-b48b-1f7dd8926e00",
    "created_at": "2025-06-23T09:21:18.261801666Z",
    "duration": 0
}
```
</details>

<details>
<summary>DELETE <code>/task/:id</code></summary>
Отмена задачи

```http
HTTP/1.1 204 No Content
```
</details>

<details>
<summary>GET <code>/task/:id/result</code></summary>
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
</details>

<details>
<summary>GET <code>/task/:id/status</code></summary>
Получение статуса задачи. Всего 4 статуса:
<ul>
    <li> created, but waiting to start; </li>
    <li> in process; </li>
    <li> completed; </li>
    <li> failed by io bound. </li>
</ul>

```json
{
    "status": "in process",
    "created_at": "2025-06-23T09:21:18.261801666Z",
    "duration": 6
}
```
> duration - динамическое время в секундах
</details>

<details>
<summary>GET <code>/task</code> - список всех задач, отладочная ручка</summary>
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
    },
	{
        ...
    }
]
```
</details>

## Архитектура
- `internal/handlers` Обработчики http запросов;
- `internal/service` Содержит сервис запросов, executer, задача которого контролировать работу I/O задач и task-state сервис;
- `internal/repository` Данный репозиторий реализован под fake хранилище, так что в случае замены на реальную БД, необходимо изменить логику;
- I/O задача-заглушка в `internal/service/executor/hardIOboundwork.go`. Задача выполняется строго 3-5 минут;
- Семафор ограничивает кол-во одновременно исполняемых задач.
- Логи либо во stdout контейнера, либо в clickhouse
- Хранилище либо in-memmory, либо sql (пока что только postgres, остальные через фабрику)
- Миграции в /migrate
