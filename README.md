# 🐹 Golang REST API with MongoDB

This is a simple RESTful API built using **Golang**, **MongoDB**, and **Julienschmidt's `httprouter`**. The API allows you to perform CRUD operations on a `users` collection stored in MongoDB.

---

## 🚀 Features

- 📦 Create a new user
- 🔍 Get a user by ID
- 🗑️ Delete a user by ID
- 📄 Get all users
- 🔌 Connects to MongoDB using official Go driver
- 🐳 Supports running MongoDB in Docker

---

## 🛠 Tech Stack

- Go `1.20+`
- MongoDB `6+` (tested with `8.0.11`)
- Docker (optional, for MongoDB)
- [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- [go.mongodb.org/mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)

---

## 📁 Project Structure

```
├── controllers
│   └── user.go
├── go.mod
├── go.sum
├── main.go
├── models
│   └── user.go
└── README.md
```


---

## ⚙️ Setup Instructions

### 1. 🚀 Clone the Repository

```bash
git clone https://github.com/your-username/mongodb-golang.git
cd mongodb-golang
```

### 2. 📦 Install Dependencies
``` 
go mod tidy

```

### 3. 🐳 Run MongoDB in Docker (Optional)

```
docker run -d \
  --name mongo \
  -p 27017:27017 \
  -e MONGO_INITDB_ROOT_USERNAME=root \
  -e MONGO_INITDB_ROOT_PASSWORD=root123 \
  mongo

```
### 4. 🧪 Run the App
```
go run main.go
```
#### The API will start on: ``` http://localhost:9001```

## 📬 API Endpoints

### ➕ Create a User

``` 
POST /user

curl -X POST http://localhost:9001/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","gender":"Female","age":22}'

```
### This Command Show Output like this. 

``` 
{"id":"686e557697cf284672870cac","name":"Alice","gender":"Female","age":22}
```

### 🔍 Get User by ID
```
GET /user/:id

curl -X GET http://localhost:9001/user/686e557697cf284672870cac | jq .

=== Output===
{
  "id": "686e557697cf284672870cac",
  "name": "Alice",
  "gender": "Female",
  "age": 22
}

```
### 🗑️ Delete User by ID
```
DELETE /user/:id

curl -X DELETE http://localhost:9001/user/686e557697cf284672870cac

=== Output === 
Deleted User ObjectID("686e557697cf284672870cac") 
                                                   
```
### 📄 Get All Users
```
GET /users

curl -X GET http://localhost:9001/user | jq .

```

## 📂 Sample MongoDB Connection (in Go)
```
uri := "mongodb://root:root123@localhost:27017/?authSource=admin"
client, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))
``` 

## ❗ Troubleshooting

    If you get a UTF-8 error in MongoDB shell:

        Use db.users.find({}, { _id: 1 }) to locate bad documents.

        Or run db.users.drop() to reset collection.

