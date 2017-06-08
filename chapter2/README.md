# Chapter 2

- Interfaces as a behavior of an object
- Interface type
- HTTP server
- HTTP handler
- Routing
- Middlewares 

Download [Slides](https://www.slideshare.net/secret/LFnJtWAHX6kLxl).

## Execution

Build, run, and verify.

```bash
make

build/homework2 \
    --listen ":8888" \
    --key "GOOGLE API KEY"

curl -X GET "http://localhost:8888/cities-suggestions?q=berlin"

curl -X POST "http://localhost:8888/cities-info" -d '["ChIJS_tHEEfaFkYR6o_AO-5iJeg","ChIJd312ZkkNOUYRCAretD6gQp4"]'
```
