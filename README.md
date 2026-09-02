# 🚀 Synthetic Financial Data Generator in Go

A modular, CLI-driven synthetic data and event generator written in Go. Designed to simulate financial streams (accounts, transactions, and fraud detection events) for testing event-driven architectures, data pipelines, and analytics platforms.

---

## 📂 Project Structure

Following Domain-Driven Design (DDD):

- `cmd/generator/main.go` - Application entry point, CLI flag handling, and streaming/batch controller.
- `internal/data/account` - User account domain and event generator.
- `internal/data/transaction` - Financial transaction domain and event generator.
- `internal/data/fraud` - Fraud score and risk assessment domain and event generator.

---

## ⚙️ CLI Flags Reference

| Flag | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `-count` | `int` | `1` | Number of event sets to generate in batch mode. |
| `-interval` | `int` | `0` | Continuous mode (streaming): generate events every X milliseconds. |
| `-transaction` | `bool` | `true` | Enable/disable transaction events. |
| `-amount` | `float64` | `0` | Force a specific transaction amount (`0` = random). |
| `-currency` | `string` | `""` | Force currency (e.g., `PLN`, `USD`, `EUR`). |
| `-status` | `string` | `""` | Force transaction status (e.g., `PENDING`, `SUCCESS`, `FAILED`). |
| `-payment-type` | `string` | `""` | Force payment type (e.g., `CARD`, `CASH`, `TRANSFER`). |
| `-account` | `bool` | `true` | Enable/disable account events. |
| `-event-type` | `string` | `""` | Force account event type (e.g., `ACCOUNT_CREATED`, `LOGIN`). |
| `-fraud` | `bool` | `true` | Enable/disable fraud events. |
| `-fraud-score` | `int` | `-1` | Force fraud score 0-100 (`-1` = random). |
| `-country` | `string` | `""` | Force country code (e.g., `US`, `PL`, `EU`). |
| `-device` | `string` | `""` | Force device type (e.g., `MOBILE`, `DESKTOP`). |

---

## 📋 Usage Examples

### 1. Default Run (One event set from all domains)

```bash
go run cmd/generator/main.go
```

### 2. Batch Mode (Generate 5 sets of events)

```bash
go run cmd/generator/main.go -count 5
```

### 3. Continuous / Streaming Mode (Generate events every 1 second)

*(Press `Ctrl + C` to stop)*

```bash
go run cmd/generator/main.go -interval 1000
```

### 4. Single Domain Mode (Generate only accounts, disable others)

```bash
go run cmd/generator/main.go -transaction=false -fraud=false -account=true
```

### 5. Fully Overridden Mode (All flags combined in a single command)

```bash
go run cmd/generator/main.go -count 1 -transaction=true -amount 999.99 -currency USD -status SUCCESS -payment-type CARD -account=true -event-type ACCOUNT_CREATED -fraud=true -fraud-score 85 -country US -device MOBILE
```

---

## 📦 Sample JSON Output

The application uses Go's native `log/slog` package with a JSON handler, producing structured log lines ready for ingestion into ELK, Datadog, or Kafka:

```json
{"time":"2026-09-02T10:00:00.000000+02:00","level":"INFO","msg":"account event generated","data":{"account_id":"f47ac10b-58cc-4372-a567-0e02b2c3d479","user_id":"a1b2c3d4-e5f6-7890-abcd-ef0123456789","event_type":"ACCOUNT_CREATED","session_id":"98765432-1234-5678-abcd-ef0123456789","timestamp":"2026-09-02T08:00:00Z"}}
{"time":"2026-09-02T10:00:00.001000+02:00","level":"INFO","msg":"fraud event generated","data":{"event_id":"b2c3d4e5-f6a1-7892-bcde-f0123456789a","user_id":"c3d4e5f6-a1b2-7893-cdef-0123456789ab","fraud_score":85,"ip_address":"192.168.1.50","country_code":"US","device_type":"MOBILE","is_suspicious":true,"timestamp":"2026-09-02T08:00:00Z"}}
{"time":"2026-09-02T10:00:00.002000+02:00","level":"INFO","msg":"transaction event generated","data":{"transaction_id":"d4e5f6a1-b2c3-7894-def0-123456789abc","account_id":"f47ac10b-58cc-4372-a567-0e02b2c3d479","amount":999.99,"currency":"USD","status":"SUCCESS","payment_type":"CARD","timestamp":"2026-09-02T08:00:00Z"}}
```

---

## 🛠️ Requirements

- Go version **1.21+** (required for native `log/slog` support).
