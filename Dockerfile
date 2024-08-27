FROM keinos/sqlite3 as db_builder

WORKDIR /tmp

COPY ./data/quotes.csv ./

RUN sqlite3 database.sqlite 'CREATE TABLE quotes(id INTEGER PRIMARY KEY, quote TEXT, author VARCHAR(255));' \
    && sqlite3 database.sqlite '.import quotes.csv quotes --csv'

FROM golang:1.23.0 as app_builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN  go build -a -installsuffix cgo ./cmd/daily_quote_api

FROM scratch

WORKDIR /app

COPY --from=db_builder /tmp/database.sqlite .
COPY --from=app_builder /app/daily_quote_api .

ENV GIN_MODE=release
ENV PORT=8080

EXPOSE 8080

CMD ["./daily_quote_api"]
