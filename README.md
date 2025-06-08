# 🧠 Simple-CRM (Go + Fiber + GORM)

A lightweight CRM (Customer Relationship Management) REST API built with:

- ⚡ **Fiber** – Express-inspired fast HTTP framework in Go  
- 🧬 **GORM** – ORM for handling database operations  
- 💾 **SQLite** – Embedded file-based database  

---

## 🗂 Project Structure

```
Simple-CRM/
├── main.go           # Entry point
├── go.mod            # Go module config
├── /database
│   └── db.go         # DB initialization and connection
└── /lead
    └── lead.go       # Lead model and route handlers
```

---

## 📦 Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/OceanWhisperer/Simple-CRM.git
cd Simple-CRM
```

### 2. Install Go dependencies

```bash
go mod tidy
```

### 3. Run the app

```bash
go run main.go
```

The app will start on [http://localhost:3000](http://localhost:3000)

---

## 📡 API Endpoints

| Method | Endpoint           | Description       |
| ------ | ------------------ | ----------------- |
| GET    | `/api/v1/`         | Get all leads     |
| GET    | `/api/v1/lead/:id` | Get lead by ID    |
| POST   | `/api/v1/lead`     | Create new lead   |
| DELETE | `/api/v1/lead/:id` | Delete lead by ID |

---

## 🧱 Lead JSON Structure

```json
{
  "name": "Alice",
  "company": "CyberCorp",
  "email": "alice@cyber.com",
  "phone": 123456
}
```

---

## 🛠 Tech Stack

- **Go 1.21+**
- **Fiber v2**
- **GORM**
- **SQLite3**

---
