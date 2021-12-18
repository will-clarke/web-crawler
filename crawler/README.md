# Crawler

The crawler will take a URL and go and find as many internal links as it can.

## Limitations

- I didn't get round to putting any limits on the crawler. If pointed at a big site like wikipedia, it'll take a *very* long time and probably run out of memory. A time and total-number-of-links limit would make sense. Maybe a `context` would be good here.

- This crawler potentially spawns a trillion goroutines. This may be okay if you've got a super-beefy server, but it could be an issue. Some sort of "worker pool" or maximum number of goroutines could help solve this. 

- My defenition of what an "internal" link is arbitrary and almost certainly wrong. I think it works in most cases, but I'd bet there are loads of edge cases I haven't thought about.

- Testing could be better; I don't think I properly tested the main `Crawl` function. I'd probably test this by either mocking out a fake httpClient or by creating a real local server with a couple of pages (the latter approach would be better IMO).

- I have no logging. Ideally I'd use a real logging library & have more {debug,info,warn,error} logs giving me context on what's going on.

- My errors aren't descriptive; when returning errors I'd probably normally wrap errrors with more details / context.

- It would be nice to configure more of the crawler (eg. configure the sleep so we don't hammer URLs too much).
