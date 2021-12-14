# Store

This is a slightly cheeky way of sharing state between concurrent crawlers goroutines and then 'persisting' it for the server.

We'd really want a proper database to do this in prod.. or at least a proper cache, rather than this ephemeral in-memory store. 
