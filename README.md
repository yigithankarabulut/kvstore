# KVStore

Simple in-memory key/value store for training purpose. Project demonstrates
basic DDD approach (storage/service/http layer logic)

---

## Usage

To run server locally;

`cd` to project root;

```bash
go run cmd/server/main.go
```

Endpoints:

```http
GET    /healthz/live/
GET    /healthz/ready/

POST   /api/v1/set/
GET    /api/v1/get/?key={key}
PUT    /api/v1/update/
DELETE /api/v1/delete/?key={key}
GET    /api/v1/list/
```

---

## Development

### Requirements

- `go1.21.0`

You can create `.env` file inside of the project root for environment variables

Environment variables information:

| Variable Name | Description | Default Value |
|:--------------|:------------|:------------|
| `SERVER_ENV` | Server environment information for run-time | `local` |
| `LOG_LEVEL` | Logging level | `INFO` |
