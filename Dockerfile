# Database builder stage for initializing and populating the quotes database
FROM keinos/sqlite3:3.50.4 AS db-builder

WORKDIR /tmp

COPY ./data/quotes.csv ./

RUN sqlite3 database.sqlite 'CREATE TABLE quotes(id INTEGER PRIMARY KEY, quote TEXT, author VARCHAR(255));' \
    && sqlite3 database.sqlite '.import quotes.csv quotes --csv'

# App builder for building the API to a standalone binary
FROM rust:1.90.0 AS app-builder

WORKDIR /app

COPY ./Cargo.toml ./
COPY ./Cargo.lock ./
COPY ./src ./src

RUN cargo build

# Swagger UI builder for fetching latest Swagger UI files
FROM swaggerapi/swagger-ui:v5.17.14 AS swagger-builder

# Remove searchbar/topbar
RUN sed -i 's#SwaggerUIStandalonePreset#SwaggerUIStandalonePreset.slice(1)#' /usr/share/nginx/html/swagger-initializer.js
# Replace default doc with local doc
RUN sed -i 's#https://petstore.swagger.io/v2/swagger.json#/swagger.yaml#' /usr/share/nginx/html/swagger-initializer.js
RUN sed -i 's#Swagger UI#Daily Quote API#' /usr/share/nginx/html/index.html

# Minimal final stage for running the application in a stripped down linux
FROM cgr.dev/chainguard/glibc-dynamic:latest

WORKDIR /app

# Import app and database
COPY --from=db-builder /tmp/database.sqlite .
COPY --from=app-builder /app/target/debug/daily-quote-api .

# Import Swagger UI and OpenAPI doc
COPY --from=swagger-builder /usr/share/nginx/html/ ./api/swagger-ui
COPY ./api/openapi.yaml ./api/openapi.yaml

ENV DATABASE_URL=sqlite://./database.sqlite
ENV ROCKET_ADDRESS=0.0.0.0
ENV ROCKET_PORT=8080

EXPOSE 8080

CMD ["./daily-quote-api"]
