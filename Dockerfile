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

RUN cargo build --release

# Minimal final stage for running the application in a stripped down linux
FROM cgr.dev/chainguard/glibc-dynamic:latest

WORKDIR /app

# Import app and database
COPY --from=db-builder /tmp/database.sqlite .
COPY --from=app-builder /app/target/release/daily-quote-api .

ENV DATABASE_URL=sqlite://./database.sqlite

EXPOSE 8080

CMD ["./daily-quote-api"]
