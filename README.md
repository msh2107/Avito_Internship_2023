# Dynamic User Segmentation Service

Микросервис хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

Используемые технологии:
- PostgreSQL (в качестве хранилища данных)
- Docker (для запуска сервиса)
- Swagger (для документации API)
- Logrus (для логирования)
- cleanenv (для создания конфига)
- Gin (веб фреймворк)
- golang-migrate/migrate (для миграций БД)
- pgx (драйвер для работы с PostgreSQL)


Сервис был написан с использованием Clean Architecture
Также был реализован Graceful Shutdown для корректного завершения работы сервиса

# Usage

Запустить сервис можно с помощью команды `make compose-up`

Документацию после запуска сервиса можно посмотреть по адресу `http://localhost:8080/swagger/index.html`, если вы не меняли хост и порт.

## Examples

Некоторые примеры запросов
- [Создание](#create)
- [Удаление](#delete)
- [Изменение сегментов](#change-active-segments)
- [Получение сегментов](#get-active-segments)

### Создание <a name="create"></a>

Создание сегмента (принимает slug сегмента):
```curl
curl --location --request POST 'http://localhost:8080/v1/segment' \
--header 'Content-Type: application/json' \
--data-raw '{
    "slug": "AVITO_VOICE_MESSAGES"
}'
```
Пример ответа:
```json
{
  "ID": 1,
  "Slug": "AVITO_VOICE_MESSAGES"
}
```

### Удаление <a name="delete"></a>

Удаление сегмента (принимает slug сегмента):
```curl
curl --location --request DELETE 'http://localhost:8080/v1/segment' \
--header 'Content-Type: application/json' \
--data-raw '{
    "slug": "AVITO_VOICE_MESSAGES"
}'
```
Пример ответа:
```json
{
  "ID": 1,
  "Slug": "AVITO_VOICE_MESSAGES"
}
```


### Изменение сегментов <a name="change-active-segments"></a>

Изменение активных сегментов у пользователя (принимает ID пользователя и списки slug`ов сегментов, которые надо добавить и удалить):
```curl
curl --location --request PATCH 'http://localhost:8080/api/v1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "toAdd": ["AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS"]
    "toRemove": ["AVITO_DISCOUNT_30"]
}'
```
Пример ответа:
```json
{
  "id": 1
}
```

### Получение сегментов <a name="get-active-segments"></a>

Получение активных сегментов у пользователя (принимает ID пользователя):
```curl
curl --location --request GET 'http://localhost:8080/api/v1/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1
}'
```
Пример ответа:
```json
{
  {
    "ID": 1,
    "Slug": "AVITO_VOICE_MESSAGES"
  }
  {
    "ID": 2,
    "Slug": "AVITO_PERFORMANCE_VAS"
  }
}
```
