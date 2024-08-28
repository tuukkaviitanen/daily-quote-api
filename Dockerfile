# Database builder stage for initializing and populating the quotes database
FROM keinos/sqlite3 as db-builder

WORKDIR /tmp

COPY ./data/quotes.csv ./

RUN sqlite3 database.sqlite 'CREATE TABLE quotes(id INTEGER PRIMARY KEY, quote TEXT, author VARCHAR(255));' \
    && sqlite3 database.sqlite '.import quotes.csv quotes --csv'

# App builder for building the API to a standalone binary
FROM golang:1.23.0 as app-builder

# Install taskfile for running tasks
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN  task build

# Swagger UI builder for fetching latest Swagger UI files
FROM swaggerapi/swagger-ui:v5.17.14 as swagger-builder

# Remove searchbar/topbar
RUN sed -i 's#SwaggerUIStandalonePreset#SwaggerUIStandalonePreset.slice(1)#' /usr/share/nginx/html/swagger-initializer.js
# Replace default doc with local doc
RUN sed -i 's#https://petstore.swagger.io/v2/swagger.json#/swagger.yaml#' /usr/share/nginx/html/swagger-initializer.js
RUN sed -i 's#Swagger UI#Daily quotes API#' /usr/share/nginx/html/index.html

# Minimal final stage for running the application in a stripped down linux
FROM scratch

WORKDIR /app

# Import app and database
COPY --from=db-builder /tmp/database.sqlite .
COPY --from=app-builder /app/daily-quote-api .

# Import Swagger UI and OpenAPI doc
COPY --from=swagger-builder /usr/share/nginx/html/ ./api/swagger-ui
COPY ./api/openapi.yaml ./api/openapi.yaml

ENV GIN_MODE=release
ENV PORT=8080

EXPOSE 8080

CMD ["./daily-quote-api"]
