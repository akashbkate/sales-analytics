Install Go, PostgreSQL

Run go run cmd/main.go

# 📊 Sales Analytics API (Golang + PostgreSQL)

This project allows you to ingest large CSV files with historical sales data, store them in a normalized PostgreSQL database, and perform revenue-based analytics through a REST API built in **Go (Golang)** using the **Gin** web framework.

---

## 📦 Features

- Load sales data from large CSV files
- Normalize and store in PostgreSQL
- Refresh data manually via API
- Revenue analytics API

---

## 🧰 Prerequisites

Before running the project, ensure you have the following installed:

- [Go (Golang)](https://go.dev/dl) – version 1.20 or higher
- [PostgreSQL](https://www.postgresql.org/download/) – running locally or in the cloud
- Git (to clone the repository)

---

## 🛠️ Setup & Run

### 1. Clone the Repository

```bash
git clone https://github.com/akash-kate/sales-analytics.git
cd sales-analytics

# Open PostgreSQL CLI or any DB GUI like pgAdmin
CREATE DATABASE sales_db;

3. Configure DB Connection
Update the DB credentials in config/db.go or use environment variables if supported.

4. Put your CSV file (e.g., sales.csv) inside the /data directory.

5. go run cmd/main.go
