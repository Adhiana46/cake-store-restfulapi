<h1 align="center">Cake Store RESTFul API</h1>
<p align="left">
    Cake Store RESTFul API is a simple api for testing as a Backend Engineer at Privy.
</p>

## API Spec
### Create Cake

Request :
- Method : POST
- Endpoint : `/api/v1/cakes`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "title" : "string",
    "description" : "string",
    "rating" : "float",
    "image" : "string"
}
```

- Response :

```json 
{
    "status" : "number",
    "message" : "string",
    "data" : {
         "id" : "int, unique, auto_increment",
         "title" : "string",
         "description" : "string",
         "rating" : "float",
         "image" : "string",
         "created_at" : "datetime",
         "updated_at" : "datetime"
     }
}
```

### Get Cake

Request :
- Method : GET
- Endpoint : `/api/v1/cakes/:id`
- Header :
    - Content-Type: application/json
    - Accept: application/json

- Response :

```json 
{
    "status" : "number",
    "message" : "string",
    "data" : {
         "id" : "int, unique, auto_increment",
         "title" : "string",
         "description" : "string",
         "rating" : "float",
         "image" : "string",
         "created_at" : "datetime",
         "updated_at" : "datetime"
     }
}
```

### Get List of Cakes

Request :
- Method : GET
- Endpoint : `/api/v1/cakes?limit=5&page=1&s=cheese&rating_min=5&rating_max=7&sort_by=rating.desc`
	- limit
	- page
	- s
	- rating_min
	- rating_max
	- sort_by
- Header :
    - Content-Type: application/json
    - Accept: application/json

- Response :

```json 
{
    "status" : "number",
    "message" : "string",
    "data" : {
         "id" : "int, unique, auto_increment",
         "title" : "string",
         "description" : "string",
         "rating" : "float",
         "image" : "string",
         "created_at" : "datetime",
         "updated_at" : "datetime"
     }
}
```

### Update Cake

Request :
- Method : PATCH
- Endpoint : `/api/v1/cakes/:id`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "title" : "string",
    "description" : "string",
    "rating" : "float",
    "image" : "string"
}
```

- Response :

```json 
{
    "status" : "number",
    "message" : "string",
    "data" : {
         "id" : "int, unique, auto_increment",
         "title" : "string",
         "description" : "string",
         "rating" : "float",
         "image" : "string",
         "created_at" : "datetime",
         "updated_at" : "datetime"
     }
}
```

### Delete Cake

Request :
- Method : DELETE
- Endpoint : `/api/v1/cakes/:id`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Response :

```json 
{
    "status" : "number",
    "message" : "string"
}
```

## Prerequisites
- Docker & Docker Compose Installed
- GO 1.8+ (running unit test only)
- Postman (for testing)

## Framework
- Web : Fiber
- Validator : go-playground/validator/v10
- Database : Mysql (Driver: go-sql-driver/mysql)
- Testing : Testify, go-sqlmock

## Architecture
Controller -> Service -> Repository
- Service : Validation & Business Logic
- Repository : Data Access Logic

## Running App (Linux)
```sh
docker compose up app -d
```

## Running App (Windows)
```sh
docker-compose up app -d
```

<p align="left">
	After running the up, next thing to do is run a database migration
</p>

## Run Database Migrations
```sh
docker compose run migrate up
```

## Run Unit Tests
```sh
go test ./... -v -cover
```

<p align="left">
And also don't forget to import postman collection `PRIVY - Pretest Cakes Store API.postman_collection.json`
</p>

---

_This README and Project was written with ❤️ by [Me](https://adhiana.me)_