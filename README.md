
# 🧮 Pack Calculator API

A simple Go-based API that calculates the optimal combination of pack sizes to fulfill a requested number of items. Built with `gin`, and ready to use via Docker or locally.

---

## 🚀 Features

- Calculate least-overage and fewest-pack combinations
- Clean architecture (domain, usecase, package)
- Dockerized backend
- CORS-ready for frontend integration
- Unit tested

---

## 📦 Tech Stack

- **Go 1.21**
- **Gin Gonic** – HTTP framework
- **Docker & Docker Compose**
- **Makefile** – for basic automation

---

## ⚙️ Setup

### 1. Clone the repository
```bash
git clone https://github.com/abdelrhman-elsbagh/pack-calculator
cd pack-calculator
```

### 2. Run locally (requires Go installed)
```bash
make run
```

or just:
```bash
go run ./cmd
```

### 3. Run with Docker
```bash
make docker-run
```

Or directly:
```bash
docker-compose up --build
```

---

## 🌐 API Endpoint

### POST `/calculate`

**Request:**
```json
{
  "items": 251,
  "pack_sizes": [250, 500, 1000]
}
```

**Response:**
```json
{
  "packs": {
    "500": 1
  },
  "total_items": 500,
  "total_packs": 1
}
```

---

## 📄 Environment Variables

| Variable | Default | Description        |
|----------|---------|--------------------|
| `PORT`   | 8080    | API listening port |

Create a `.env` file or pass via terminal before running.

---

## 🔧 Makefile Commands

| Command        | Description                |
|----------------|----------------------------|
| `make build`   | Build the Go project       |
| `make run`     | Run locally using `go run` |
| `make test`    | Run unit tests             |
| `make docker-run` | Run via Docker Compose  |
| `make clean`   | Remove the built binary    |

---

## 📘 API Documentation

Once running, visit:

```
http://54.91.1.171:8080/calculate
```

To view Redoc UI (OpenAPI documentation)

---

## 🧪 Running Tests

```bash
make test
```

---

## 📂 Project Structure

```
internal/
├── configs/       # App configuration
├── domain/
│   ├── entity/    # Input structs
│   └── service/   # Core logic (calculator)
├── usecase/       # Business logic
├── delivery/
│   └── http/      # HTTP handlers & router
cmd/               # Application entry
test/              # Unit tests
```

---

## 🧑‍💻 Author

Made with 💙 by **Abdelrahman Tarek**

---
