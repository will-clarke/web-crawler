# Some curl requests you may wanna make:

``` sh
make run
curl -u user:pass localhost:8080/crawl -d '{"url":"https://wclarke.net/about.html"}'

curl -u user:pass localhost:8080/crawl -d '{"url":"https://wclarke.net/about.html"}' | jq ".id" | xargs -I {} sh -c 'for i in $(seq 5); do sleep 1; curl -u user:pass localhost:8080/crawl/{}; done'

curl localhost:8080/healthcheck
```


# Super important stuff we still need to do

- [x] crawl more than one url
- [ ] authentication
- [ ] web crawling limits. total request / time / something
- [ ] main-loop error handling. Do we return any errors? Just log them?
- [x] actually use goroutines
- [ ] use a worker pool


# ideas to extend this:
- could return a more useful result. 
  - Directed graph? Could even output a graphviz visualisation.
- We could enrich simplistic "link tracking" with loads more data. Instead of just returning the links, we could add the link count, how important the links were (were they within an `<h1>`?), etc..

# limitations
- basic auth is very basic
- Error handling's a bit ropey. Ideally I'd wrap errors a bit more to give some more context
- Logging could be better; we could have used a proper logging framework. Maybe backed by elasticsearch / Kibana to easily access & interact with these logs.
- Metrics / Stats would be super cool. Something like statsd / prometheus would be good with something like Grafana to visualise it.
- My defenition of what an "internal" link is arbitrary and almost certainly wrong. I think it works in most cases, but I'd bet there are loads of edge cases I haven't thought about.
- The "store" is a bit of a fraud and isn't totally ACID(ic?). But it helped demo shared state in memory. We'd probably want something more persistent.
