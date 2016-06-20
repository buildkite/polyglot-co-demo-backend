# Buildkite Polyglot Co. Demo Backend

## Developing

With Go 1.6+ installed:

```bash
go run main.go
open 'http://localhost:8000'
```

With docker-compose:

```bash
docker-compose up
open 'http://localhost:8000'
```

## Developing frontend

```bash
# In one tab
cd frontend && npm install && npm run start
# In another tab
env FRONTEND_DEV=true go run main.go
# And then...
open 'http://localhost:8000'
# Don't forget to check it in
cd frontend && npm run build && git commit
```
