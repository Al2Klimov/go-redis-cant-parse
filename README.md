# Rationale

This Dockerfile contains dummy applications designed to reproduce
[this issue](https://github.com/go-redis/redis/issues/808).

# Build

```bash
docker build -t go-redis-cant-parse --cache-from=go-redis-cant-parse .
```

# Run

```bash
docker run -it go-redis-cant-parse
```

# Watch out

... for an error message like this one:

```
2018-07-16 10:01:14,623 DEBG 'subscriber' stdout output:
redis: can't parse "ype\":\"PerfdataValue\",\"unit\":\"\",\"value\":0.0,\"warn\":null}],\"status\":{\"checkercomponent\":{\"checker\":{\"i"
```
