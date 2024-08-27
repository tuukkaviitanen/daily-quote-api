# Database builder stage for initializing and populating the quotes database
FROM keinos/sqlite3 as db_builder

WORKDIR /tmp

COPY ./data/quotes.csv ./

RUN sqlite3 database.sqlite 'CREATE TABLE quotes(id INTEGER PRIMARY KEY, quote TEXT, author VARCHAR(255));' \
    && sqlite3 database.sqlite '.import quotes.csv quotes --csv'

# App builder for building the API to a standalone binary
FROM golang:1.23.0 as app_builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN  go build -a -installsuffix cgo ./cmd/daily_quote_api

# Minimal final stage for running the application in a stripped down linux
FROM scratch

WORKDIR /app

# Only includes API and database files
COPY --from=db_builder /tmp/database.sqlite .
COPY --from=app_builder /app/daily_quote_api .

ENV GIN_MODE=release
ENV PORT=8080

EXPOSE 8080

CMD ["./daily_quote_api"]
