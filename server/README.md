# Server

The server exposes three routes:

## GET /healthcheck
This request just checks that the server is up and running.

## POST /crawl
This request submits a "crawl" job to the server. It doesn't wait and it returns a `uuid` that we can later use to find the results of the queries.

### Request

``` json
{
    "url":"..."
}
```
where `url` is the url you want to parse.

### Response

``` json
{
    "id":"...",
    "error": "..."
}
```
where `id` is a `uuid` and
`error` is a nullable error message.

## GET /crawl/:id
where `:id` is the uuid received from the response of `POST /crawl`.
### Response
``` json
{
    "id":"...",
    "links": ["...", "..."]
}
```
where `id` is the `:id` from the url parameter and
`links` is a list of urls found either directly or indirectly from the `url` from the `POST /crawl` payload.


## Limitations
- Basic Auth is **SUPER BASIC**. We'd want to fix that & probablyn  use something fancy like JWTs in prod.
- We probably shouldn't be exposing our error messages in case they're a security flaw.
- Test coverage is pretty mediocre. I've covered the main paths, but I could improve tests by testing auth and by adding stubs to test that I'm actually calling the right methods, etc..
