# Bookstore REST API

The Bookstore REAT API project is a great example of a REST API with everything needed to deploy.

The entire API is documented with Swagger 2.0 using the <a href="https://www.github.com/swaggo/swag" target="blank">swaggo/swag</a> package. The application also contains fully automated tests for unit testing (100% coverage) and a Dockerfile for the containerization.

### Settings

Everything is easily build and run with the `make` commands: `test`, `install`, `docker-build`, `docker-run`.
For simplicity, the API settings can be set in the single file: `settings.yaml`, so feel free to modify it.

### API documentation

The documentation is located in the `docs`. The Swagger 2.0 description can be found in `docs/swagger.json` or `docs/swagger.yaml` files. Alternatively run the API and browse to <a href="https://localhost:8081/swagger/index.html">localhost:8081/swagger/index.html</a>.
The examples of the HTTP requests are in `example/requests.rest`

### Book model

### File tree

#### TODO
- log to file
