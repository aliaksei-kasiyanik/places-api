# High-loaded web-service
## *BSU 2016 Big Data Master*

### API

http://gotravel.today/places - API endpoint

https://aliaksei-kasiyanik.github.io - Swagger **!Не выполняет запросы, т.к. требуется поддержка https (политика Github Pages). Будет перенесено на сервер.**

**Изменения в API:**

1. /places GET - возвращает сокращенную версию Place entity, так как подразумевается, что даный ресурс будет использоваться для отображения точек на карте.  

2. /places/{id} GET - возвращает полную версию Place entity.

3. Мелкие изменения согласно рекомендациям после лабораторной по API, например, удалено поле id из респонса (избытачное).

### Architecture
https://github.com/aliaksei-kasiyanik/places-api/blob/master/docs/System_Architecture.md

### Infrastructure
Places API сервис развернут на сервере DigitalOcean, имеющий следующие параметры:

* 2 GB Memory / 40 GB Disk / AMS2 - Ubuntu 14.04.4 x64

На сервере установлен standalone instance MongoDB, две ноды Places API (REST-сервис на Golang) и Nginx, который балансирует нагрузку между нодами API. Вся инфраструктура разворачивалась с помощью Docker.

### Данные

Для тестирования сервиса были взяты данные с Foursquare (!в учебных целях), содержащие информацию о местах города Минска. Для этого был написан скрипт на Python, который делает запросы в Foursquare API, преобразует данные в формат Places API и сохраняет с помощью /places POST. На данный момент Places API содержит 8570 точек следующих категорий:

* "4d4b7104d754a06370d81259" Arts
*	"4d4b7105d754a06379d81259" Travel and transport (частично, так как был превышен лимит Foursquare API)
* "4d4b7105d754a06377d81259" Outdoor and recreation

Во избежание дубликации данных при импорте в коллекции places в MongoDB в документ было добавлено поле fsId (foursquareId), и создан unique partitial index по этому полю.

Также коллекция places имеет geospatial индекс для поиска точек по локации.

Пример документа в коллекции places:
```
{
    "_id" : ObjectId("586292862fcf72000646c606"),
    "loc" : {
        "type" : "Point",
        "coordinates" : [ 
            27.66547459333, 
            53.8816603283683
        ]
    },
    "name" : "Мемориал Тростенец",
    "cat" : [ 
        "Historic Site"
    ],
    "fsId" : "4f967fcfe4b0dc878d9d4df0",
    "lastModified" : ISODate("2016-12-27T16:10:46.806Z")
}
```
### Нагрузочное тестирование

TODO


## Useful Links
### Go Dev

**How to organize Go project and dev environment**

1. https://golang.org/doc/code.html
2. http://skife.org/golang/2013/03/24/go_dev_env.html
3. https://github.com/go-lang-plugin-org/go-lang-idea-plugin/wiki/Documentation

**Awesome Go**

A curated list of awesome Go frameworks, libraries and software.

https://awesome-go.com/

#### Web

http://golang-book.ru/

http://www.golang-book.com/public/pdf/gobook.0.pdf

http://openmymind.net/assets/go/go.pdf

https://www.gitbook.com/book/astaxie/build-web-application-with-golang/details

**Http implementation**

1. Non-default http implementation for Go. It is up to 10x faster than net/http.
 https://github.com/valyala/fasthttp

**Http Multiplexer**

1. Popular multiplexers benchmark  
   https://www.peterbe.com/plog/my-favorite-go-multiplexer
2. httprouter  
   https://github.com/julienschmidt/httprouter
   
   
#### MongoDB

1. Rich MongoDB driver for Go  
   https://godoc.org/labix.org/v2/mgo

2. Using MongoDB with the Go language  
   http://spf13.com/presentation/MongoDB-and-Go/
