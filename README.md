# 🧠 Simple-CRM (Go + Fiber + GORM)

A lightweight CRM (Customer Relationship Management) REST API built with:

- ⚡ [Fiber](https://github.com/gofiber/fiber) – Express-inspired fast HTTP framework in Go  
- 🧬 [GORM](https://gorm.io/) – ORM for handling database operations  
- 💾 SQLite – Embedded file-based database  

---

## 🗂 Project Structure

Simple-CRM/
├── main.go # Entry point
├── go.mod # Go module config
├── /database
│ └── db.go # DB initialization and connection
└── /lead
└── lead.go # Lead model and route handlers

yaml
Copy
Edit

---

## 📦 Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/OceanWhisperer/Simple-CRM.git
cd Simple-CRM
2. Install Go dependencies
bash
Copy
Edit
go mod tidy
3. Run the app
bash
Copy
Edit
go run main.go
App will start on http://localhost:3000

📡 API Endpoints
Method	Endpoint	Description
GET	/api/v1/	Get all leads
GET	/api/v1/lead/:id	Get lead by ID
POST	/api/v1/lead	Create new lead
DELETE	/api/v1/lead/:id	Delete lead by ID

🧱 Lead JSON Structure
json
Copy
Edit
{
  "name": "Alice",
  "company": "CyberCorp",
  "email": "alice@cyber.com",
  "phone": 123456
}
🛠 Tech Stack
Go 1.21+

Fiber v2

GORM

SQLite3

⚠️ Notes
✅ Commit go.mod and go.sum to version control.

❌ Don't commit .env files, API keys, or credentials.

Always check for errors in:

c.BodyParser(&struct)

db.First(...)

Fiber handlers must return error:

go
Copy
Edit
func Handler(c *fiber.Ctx) error
📄 License
MIT — free to use, modify, distribute.

yaml
Copy
Edit

---

Save this as a file named `README.md` in the root of your `Simple-CRM` project.

Let me know if you want it pre-written into a `.md` file ready to download.