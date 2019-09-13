# itsallok

HTTP server that always returns 200 OK

# Run

```sh
docker run -p 8080:8080 nstogner/itsallok

# In another terminal...
curl localhost:8080/any-path-is-ok
```

# Configure

Use the environment variable `ADDR` (defaults to `:8080`).

