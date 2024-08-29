# Daily Quote API

> API that provides your daily quotes, with a twist

Try it out the interactive documentation live at https://daily-quote-api.tuukka.net

## Summary

- By default, the [`/quote`](https://daily-quote-api.tuukka.net/quote) endpoint returns the daily quote that refreshes every midnight (UTC)
- In addition, the client can set the `of-the` query parameter, to request the quote of other time units
  - For example [`/quote?of-the=week`](https://daily-quote-api.tuukka.net/quote?of-the=week),  [`/quote?of-the=fortnight`](https://daily-quote-api.tuukka.net/quote?of-the=fortnight) or [`/quote?of-the=second`](https://daily-quote-api.tuukka.net/quote?of-the=second)
- See the [OpenAPI](https://swagger.io/specification/) documentation from the integrated [Swagger UI](https://swagger.io/tools/swagger-ui/) tool at https://daily-quote-api.tuukka.net
  - The API can be tested there with the list of all `of-the` parameter options

## Technical summary

- The API is written in [Go](https://go.dev/)
  - This is my first time using Go, so I've written some of my thoughts on the language in this README
- The quotes themselves are from this dataset: https://www.kaggle.com/datasets/abhishekvermasg1/goodreads-quotes
  - I removed likes and tags manually as they aren't needed
- The quotes are stored in an [SQLite](https://www.sqlite.org/) database
- [Swagger UI](https://swagger.io/tools/swagger-ui/) is built into to application for easy access
- All of this is built into a single [Docker image](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-an-image/) with the size of only around 30 MB!

### Server

#### API

- The API itself is built with [Gin](https://gin-gonic.com/); a popular Web framework for Go, that gives a similar API building experience as [express](https://expressjs.com/) from the [JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript) ecosystem
- I tried building the API with Go's built-in [net/http](https://pkg.go.dev/net/http) package, but I couldn't bear it. It seemed to lack some friendly tools for building simple APIs, as many things just had to be done manually (e.g. endpoint methods, JSON conversions and headers). I'm just saying I cut down the code by more than half in 15 minutes after switching to Gin. That might just be my inexperience with the language, but I'm still sticking with it from now on.
  - The reason I was so adamant in using the built-in library, was that it seemed to me like everyone on the internet was telling me to avoid using dependencies with Go. [Here's an article about that](https://medium.com/@joeybloggs/go-dependencies-are-the-devil-a2b60c25e5d9#.1c59f6goq). I'm glad I got over that phase though.

#### Database connection

- The database connection is created using [GORM](https://gorm.io/index.html) as the [ORM](https://www.freecodecamp.org/news/what-is-an-orm-the-meaning-of-object-relational-mapping-database-tools/)
  - I've used a few ORMs in other languages, and I think that this just might be the easiest to set up. In all fairness, the database for this project is **quote** simple (haha). I still recommend it.
- I used a [Pure Go SQLite driver](https://pkg.go.dev/github.com/glebarez/sqlite@v1.11.0)
  - Most SQLite drivers usually require [CGO](https://go.dev/wiki/cgo) to be enabled, as the original SQLite engine is built with C, and it's implemented into the app as a C library ([source](https://en.wikipedia.org/wiki/SQLite)) (yes the source is wikipedia, get off my back)
    - CGO requires [gcc](https://gcc.gnu.org/) and some other external dependencies to exist in the deployment environment, and I wanted to use an empty Docker base image, so I couldn't have that

#### Generating the quotes

- The basic idea behind the quote generation is basically this:
  1. Get the [UTC/Epoch](https://en.wikipedia.org/wiki/Unix_time) timestamp for the selected unit, e.g. today at 00:00 UTC for the quote of the day
  2. Set that as the seed of a random generator
  3. Generate a number between 0 and the amount of quotes in the database
  4. Pick a quote with that number as id (the quote id is just an incrementing int)

### Database

- I used [SQLite](https://www.sqlite.org/) as the database, as I just wanted to store this [CSV](https://en.wikipedia.org/wiki/Comma-separated_values) file somewhere other than memory
  - There is no need to modify the data, and the amount of quotes is quite small (around 3000), so setting up a completely separate database instance for this seemed too much

### Putting this all together

- The running environment is created using [Docker](https://www.docker.com/)
- The Dockerfile consists of four stages (3 build stages and the final stage)
  1. The database builder stage creates and populates the database file
  2. The app builder builds the Go application into a binary
  3. The Swagger UI builder takes the Swagger UI files from the base image, and modifies them a bit for this use case
  4. Finally, the app binary, database file and swagger files are copied to an [empty base image](https://hub.docker.com/_/scratch)

#### Built image

- The final image is only around 30 MB, of which the Swagger UI files take up around 10 MB and the populated database file takes up 0.5 MB
  - This is why I wanted the app to have no external dependencies such as gcc
  - The default golang base image which includes all CGO dependencies is 837 MB itself and the smaller alpine version, which doesn't even include the CGO dependencies, is still 245 MB. The difference seems crazy.

### Final thoughts

- It took me some time to get familiar with the language, everything seemed a bit off at first
- I was intrigued by the promise of really performant APIs, and I have to say I'm not disappointed. The size of the final image takes around 30 MB disc space and the memory usage while idle is only 5 MB. That's ~90% less disk space and ~95% less memory than my express APIs; just as a really bad comparison.
- I still think some things have to be done a bit too manually. For example I really miss more functional tools like [LINQ](https://learn.microsoft.com/en-us/dotnet/csharp/linq/) and [JavaScript Array methods](https://www.w3schools.com/js/js_array_methods.asp)
  - Although doing array manipulations manually might be more performant, and that's the preference with this language
- I have lately grown to love immutable programming, so I miss the const from JavaScript and readonly from C#. Only immutable variables here seem to be compile-time constants. Although again, mutable might be quite a bit more performant in some cases.
- I didn't get to check out the famous [goroutines](https://www.geeksforgeeks.org/goroutines-concurrency-in-golang/), but hope to some day
- I liked the pointers being there

#### TL;DR

- Liked the performance
- Didn't fall in love with the syntax at first, but it grew on me
- Felt the built-in libraries were a bit lacking
- Might try again sometime
