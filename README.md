# Bookstore REST API example
The Bookstore project is a great example of a REST API with everything needed to deploy.

The entire API is documented with Swagger 2.0 using the <a href="https://www.github.com/swaggo/swag" target="_blank">swaggo/swag</a> package. The application also contains fully automated tests for unit testing (with 100% coverage) and a Dockerfile for the containerization.

### Settings
Everything is easily built and run with the `make` commands: `test`, `install`, `docker-build`, `docker-run`.
For simplicity, the API settings can be set in a single file: `settings.yaml`.

_(This API support postgres DB only)_

__Default:__
```
---
logging:
  destination: STD_OUT

server:
  addres: ':8081'
  read_timeout: 1500ms
  read_header_timeout: 800ms
  write_timeout: 3000ms
  idle_timeout: 6000ms
  max_header_bytes: 1048576

database:
  host: localhost
  port: 5432
  db_name: bookstore
  username: root
  password: 123456

debug_mode: false
```

### API documentation
The documentation is located in the `docs`. The Swagger 2.0 description can be found in `docs/swagger.json` or `docs/swagger.yaml` files. Alternatively run the API and browse to <a href="https://localhost:8081/swagger/index.html" target="_blank">localhost:8081/swagger/index.html</a>.
The examples of the HTTP requests are in `example/requests.rest`.

### Book model
<table>
   <tbody>
     <tr>
         <td>id</td>
         <td>integer</td>
         <td>PRIMARY KEY</td>
      </tr>
     <tr>
         <td>sku</td>
         <td>string</td>
         <td>UNIQUE, NOT NULL</td>
      </tr>
     <tr>
         <td>title</td>
         <td>string</td>
         <td>NOT NULL</td>
      </tr>
     <tr>
         <td>author</td>
         <td>string</td>
         <td>NOT NULL</td>
      </tr>
     <tr>
         <td>description</td>
         <td>string</td>
         <td></td>
      </tr>
     <tr>
         <td>price</td>
         <td>integer</td>
         <td>NOT NULL</td>
      </tr>
   </tbody>
</table>

### Test results
```
ok  	github.com/chutified/bookstore-api/app	(cached)	coverage: 100.0% of statements
ok  	github.com/chutified/bookstore-api/app/dbservices	(cached)	coverage: 100.0% of statements
ok  	github.com/chutified/bookstore-api/app/handlers	(cached)	coverage: 100.0% of statements
ok  	github.com/chutified/bookstore-api/app/middlewares	(cached)	coverage: 100.0% of statements
ok  	github.com/chutified/bookstore-api/config	(cached)	coverage: 100.0% of statements
```

### Directory tree
```
├── app
│   ├── app.go
│   ├── app_test.go
│   ├── dbservices
│   │   ├── conn.go
│   │   └── conn_test.go
│   ├── handlers
│   │   ├── books.go
│   │   ├── books_test.go
│   │   ├── common.go
│   │   ├── common_test.go
│   │   ├── routes.go
│   │   └── routes_test.go
│   ├── middlewares
│   │   ├── db.go
│   │   └── db_test.go
│   └── models
│       ├── book.go
│       └── error.go
├── config
│   ├── config.go
│   ├── config_test.go
│   ├── db.go
│   ├── log.go
│   ├── server.go
│   └── tests
│       ├── file-logs.yaml
│       ├── invalid.yaml
│       ├── 0_invalid.yaml
│       ├── 0_settings.yaml
│       ├── 1_invalid.yaml
│       ├── 1_settings.yaml
│       ├── 2_invalid.yaml
│       └── 3_invalid.yaml
├── Dockerfile
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── example
│   └── requests.rest
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── README.md
├── settings.yaml
└── temp
```
