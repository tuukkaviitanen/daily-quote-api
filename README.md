# Daily Quote API

> API that provides your daily quotes, with a twist

Try it out the interactive documentation live at https://daily-quote-api.tuukka.net

This is the new [Rust](https://rust-lang.org/) implementation. See the `go-implementation` branch for the original in [Go](https://go.dev/), which was my first Go project, so it has quite a lot of my thoughts documented in the README.

## Summary

- By default, the [`/quote`](https://daily-quote-api.tuukka.net/quote) endpoint returns the daily quote that refreshes every midnight (UTC)
- In addition, the client can set the `of_-_the` query parameter, to request the quote of other time units
  - For example [`/quote?of_the=week`](https://daily-quote-api.tuukka.net/quote?of_the=week),  [`/quote?of_the=fortnight`](https://daily-quote-api.tuukka.net/quote?of_the=fortnight) or [`/quote?of_the=second`](https://daily-quote-api.tuukka.net/quote?of_the=second)
- See the [OpenAPI](https://swagger.io/specification/) documentation from the integrated [Swagger UI](https://swagger.io/tools/swagger-ui/) tool at https://daily-quote-api.tuukka.net
  - The API can be tested there with the list of all `of_the` parameter options

## Technical summary

- The API is written in [Rust](https://rust-lang.org/)
- The quotes themselves are from this dataset: https://www.kaggle.com/datasets/abhishekvermasg1/goodreads-quotes
  - I removed likes and tags manually as they aren't needed
- The quotes are stored in an [SQLite](https://www.sqlite.org/) database
- [Swagger UI](https://swagger.io/tools/swagger-ui/) is built into to application for easy access
- All of this is built into a single [Docker image](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-an-image/) with the size of only around ~22 MB (while the Go implementation in a scratch base image was around 30 MB)
