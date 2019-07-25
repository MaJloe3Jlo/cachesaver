# Cache saver by m3 v.0.0.1

## `EN` **Task**

App is implementation of Cache in ROM memory.

* Commands: GET, SET, REMOVE, KEYS
* TTL for every key.
* REST-API.

## **Requirements and start of app**
```
go get github.com/gorilla/mux
```
```
go run main.go
```

The server is available by URL: http://localhost:8080

* GET methods: /all; /cache/{key}; /keys
* POST methods: /cache_create/{timeDuration}; /cache_create/{key}/{val}/{timeDuration}
* DELETE methods: /cache/{key}

**Examples**
```
curl -i http://localhost:8080/cache_create/2m -XPOST - make an CACHE for 2 minutes.
  
curl -i http://localhost:8080/cache_create/e/two/2m -XPOST - make an key-value for 2 minutes.

curl -i http://localhost:8080/all - give all.

curl -i http://localhost:8080/cache/e - give an value.

curl -i http://localhost:8080/keys - give an keys.

curl -i http://localhost:8080/cache/e -XDELETE - delete by key.
```

# }{paHuTeJlb Cache от m3 вер.0.0.1

## `RU` **Задание**

Имплементация Cache в памяти.

* Функционал: GET, SET, REMOVE, KEYS
* TTL на каждый ключ
* REST-api

## **Зависимости и запуск**
```
go get github.com/gorilla/mux
```
```
go run main.go
```

Сервер доступен по адресу http://localhost:8080

* GET методы: /all; /cache/{key}; /keys
* POST методы: /cache_create/{timeDuration}; /cache_create/{key}/{val}/{timeDuration}
* DELETE методы: /cache/{key}

**Пример использования**
```
curl -i http://localhost:8080/cache_create/2m -XPOST -создаем кеш на 2 минуты.
  
curl -i http://localhost:8080/cache_create/e/two/2m -XPOST - создаем key-value на 2 минуты.

curl -i http://localhost:8080/all - получить все.

curl -i http://localhost:8080/cache/e - получить value.

curl -i http://localhost:8080/keys - получить keys.

curl -i http://localhost:8080/cache/e -XDELETE - удалить по ключу.
```
