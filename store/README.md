# Store

This is a slightly cheeky way of sharing state between concurrent crawler goroutines and then 'persisting' it for the server.

We'd really want a proper database to do this in prod.. or at least a proper cache, rather than this ephemeral in-memory store. But it should be able to use the same interface we've defined here.

## URL Store

Because of what we're trying to acheive in this project, the interface to the store isn't super simple. We need a map of maps. The first map holds the "uuid". And then one uuid has many "links".
This allows multiple clients to submit multiple urls for processing, and get them all back independently.
