**Задание**

Имплементация Cache в памяти.

Функционал: GET, SET, REMOVE, KEYS
TTL на каждый ключ
REST-api

**Зависимости и запуск**

go get github.com/gorilla/mux

go run main.go

Сервер доступен по адресу http://localhost:8080"

GET методы: /all; /cache/{key}; /keys

POST методы: /cache_create/{timeDuration}; /cache_create/{key}/{val}/{timeDuration}
 
DELETE методы: /cache/{key}

**Пример использования**

curl -i http://localhost:8080/cache_create/2m -XPOST -создаем кеш на 2 минуты
  
curl -i http://localhost:8080/cache_create/e/two/2m -XPOST - создаем key-value на 2 минуты

curl -i http://localhost:8080/all - получить все

curl -i http://localhost:8080/cache/e - получить value

curl -i http://localhost:8080/keys - получить keys

curl -i http://localhost:8080/cache/e -XDELETE - удалить по ключу