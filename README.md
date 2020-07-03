# Bookstore REST API

The Bookstore project is a great example of a REST API with everything needed to deploy.

The entire API is documented with Swagger 2.0 using the <a href="https://www.github.com/swaggo/swag" target="_blank">swaggo/swag</a> package. The application also contains fully automated tests for unit testing (100% coverage) and a Dockerfile for the containerization.

### Settings

Everything is easily build and run with the `make` commands: `test`, `install`, `docker-build`, `docker-run`.
For simplicity, the API settings can be set in a single file: `settings.yaml`. Feel free to modify it as you need.

### API documentation

The documentation is located in the `docs`. The Swagger 2.0 description can be found in `docs/swagger.json` or `docs/swagger.yaml` files. Alternatively run the API and browse to <a href="https://localhost:8081/swagger/index.html" target="_blank">localhost:8081/swagger/index.html</a>.
The examples of the HTTP requests are in `example/requests.rest`

### Book model

<table>
   <tbody>
     <tr>
         <td>id</td>
         <td>integer</td>
      </tr>
     <tr>
         <td>sku</td>
         <td>string</td>
      </tr>
     <tr>
         <td>title</td>
         <td>string</td>
      </tr>
     <tr>
         <td>author</td>
         <td>string</td>
      </tr>
     <tr>
         <td>description</td>
         <td>string</td>
      </tr>
     <tr>
         <td>price</td>
         <td>integer</td>
      </tr>
   </tbody>
</table>

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
#### TODO
- log to file
