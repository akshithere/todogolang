# TodoGolang

**Basic CRUD Todo App Backend** to learn the fundamentals of **Go (Golang)** — including HTTP routing, request handling, and slice manipulation.

---

## 🚀 Features

- ✅ Add a todo
- 📋 View all todos
- ✏️ Update a todo by index
- 🗑️ Delete a todo by index
- Simple code structure for learning

---

## Endpoints

| Method | Endpoint            | Description                         | Example                                     |
|--------|---------------------|-------------------------------------|---------------------------------------------|
| `GET`  | `/`                 | Health check                        | `/`                                         |
| `GET`  | `/todos`            | Get all todos                       | `/todos`                                    |
| `GET`  | `/addtodo`          | Add a new todo                      | `/addtodo?task=Buy+milk`                    |
| `GET`  | `/updatetodo`       | Update a todo by index              | `/updatetodo?index=0&task=Learn+Go+properly`|
| `GET`  | `/deletetodo`       | Delete a todo by index              | `/deletetodo?index=1`                       |

---

## 🛠️ How to Run

Make sure you have **Go installed**. Then:

```bash
go run main.go
