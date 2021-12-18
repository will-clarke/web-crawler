# Web Crawler

This web crawler asynchronously crawls URLs recursively and returns a list of internal links.

You interact with the via a (very limited) RESTFUL API.

I wrote this code in golang and there are three main components:

- [Crawler](crawler/README.md)
- [Server](server/README.md)
- [Store](store/README.md)

Each of these packages have their own readmes, so have a look in there if you're curious about how they work & fit together.

## Some commands to get you started

``` sh
# Start the server
make run

# Start the crawling process with an example URL
curl -u user:pass localhost:8080/crawl -d '{"url":"https://wclarke.net/about.html"}'

# See the links!!
curl -u user:pass localhost:8080/crawl/{id}

# Here's a one-liner for the lazy. This makes the request 
# and then, makes 5 requests over 5 seconds.
# You may be able to see the list of links growing over time.
curl -u user:pass localhost:8080/crawl -d '{"url":"https://wclarke.net/about.html"}' | jq ".id" | xargs -I {} sh -c 'for i in $(seq 5); do sleep 1; curl -u user:pass localhost:8080/crawl/{}; done'

# test that the server's actually working
curl localhost:8080/healthcheck
```

# Ideas to extend this:
- we could return a more useful result. 
  - Directed graph? Could even output a graphviz visualisation.
- We could enrich simplistic "link tracking" with loads more data. Instead of just returning the links, we could add the link count, how important the links were (were they within an `<h1>`?), etc..

# Limitations

(more limitations in the individual package readmes...)

- Error handling's a bit ropey. Ideally I'd wrap errors a bit more to give some more context
- Logging could be better; we could have used a proper logging framework. Maybe backed by elasticsearch / Kibana to easily access & interact with these logs.
- Metrics / Stats would be super cool. Something like statsd / prometheus would be good with something like Grafana to visualise it.
