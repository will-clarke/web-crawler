# Some curl requests you may wanna make:

``` sh
make run
curl localhost:8080/crawl -d '{"url":"https://wclarke.net"}'

curl localhost:8080/crawl -d '{"url":"https://wclarke.net"}' | jq ".id" | xargs -I {} curl localhost:8080/crawl/{}

curl localhost:8080/healthcheck
```
