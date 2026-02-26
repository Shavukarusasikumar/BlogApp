# 📝 BlogApp – Full Stack Blog Application (Go + JWT + HTML/CSS)

## 🚀 Project Overview

BlogApp is a full-stack blog application built using **Golang (Go)** for the backend and **HTML, CSS, and JavaScript** for the frontend.

This project demonstrates how to build a real-world REST API with authentication and connect it to a responsive frontend UI.

Users can:

- Register an account
- Login using JWT authentication
- Create posts
- Edit posts
- Delete posts
- View all posts
- View a single post
- Add comments
- View comments
- Use pagination and sorting

---

# 🧠 What I Learned in Go

## 1️⃣ Building REST APIs

- Creating HTTP servers using `net/http`
- Using `gorilla/mux` for routing
- Handling GET, POST, PUT, DELETE requests
- Parsing JSON request bodies
- Sending JSON responses

Example:

```go
r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
```

---

## 2️⃣ Clean Project Structure

```
BlogApp/
│
├── main.go
├── routes/
├── handlers/
├── middleware/
├── models/
├── tables/
└── static/ (Frontend files)
```

This helped me understand:

- Separation of concerns
- Scalable backend architecture
- Clean coding practices

---

## 3️⃣ JWT Authentication

Implemented:

- Login API
- JWT token generation
- Token validation middleware
- Protected routes

Example:

```go
r.Handle("/post",
    middleware.ValidateToken(http.HandlerFunc(handlers.CreatePost)),
).Methods("POST")
```

Learned:

- Token-based authentication
- Securing APIs
- Authorization header handling

---

## 4️⃣ Password Security

- Hashing passwords before storing
- Comparing hashed passwords during login
- Understanding why plain-text passwords are insecure

Security concepts applied:

- Password hashing
- JWT authentication

---

## 5️⃣ CRUD Operations

| Operation | Endpoint |
|-----------|----------|
| Create Post | POST /post |
| Get Posts | GET /posts |
| Get Post By ID | GET /post/{id} |
| Update Post | PUT /post/{id} |
| Delete Post | DELETE /post/{id} |

Learned:

- REST principles
- Resource-based routing
- Handling HTTP status codes

---

## 6️⃣ Pagination

Implemented:

- Page-based loading
- Limit and offset logic
- Next and previous navigation
- Edge case handling

Example:

```
/posts?page=1
```

---

## 7️⃣ Sorting

Used query parameters:

```
/posts?page=1&sortBy=created_at&sortOrder=desc
```

Learned:

- Dynamic query building
- Sorting results
- Handling URL query parameters

---

## 8️⃣ Middleware in Go

Created middleware to:

- Extract Authorization header
- Validate JWT token
- Protect routes

This helped me understand how real backend systems handle authentication.

---

# 🌐 What I Learned in Frontend

## 1️⃣ Connecting Frontend to Backend

Used `fetch()` API:

```javascript
fetch("/post", {
  method: "POST",
  headers: {
    "Authorization": "Bearer " + token
  }
})
```

Learned:

- Sending JSON requests
- Handling HTTP responses
- Error handling

---

## 2️⃣ LocalStorage for Authentication

Stored JWT token:

```javascript
localStorage.setItem("token", data.token);
```

Used it to access protected APIs.

---

## 3️⃣ DOM Manipulation

- Dynamically rendering posts
- Loading comments
- Handling pagination buttons
- Updating UI after API calls

---

## 4️⃣ UI/UX Design

Designed a clean:

- Apple-style black and white theme
- Responsive layout
- Modern minimal UI
- Smooth user experience

---

# 🔥 Features Implemented

- User Registration
- User Login
- JWT Authentication
- Create/Edit/Delete Posts
- Comments System
- Pagination
- Sorting
- Protected Routes
- Static file serving using Go

---

# ⚙️ Technologies Used

## Backend
- Golang
- Gorilla Mux
- JWT
- net/http
- JSON encoding/decoding

## Frontend
- HTML5
- CSS3
- JavaScript (Vanilla)
- Fetch API
- LocalStorage

---

# 🏗️ How I Applied My Learning

| Concept | Applied In |
|---------|------------|
| Structs | Models |
| JSON | API communication |
| Middleware | JWT validation |
| Hashing | Secure password storage |
| Query Params | Sorting |
| Limit & Offset | Pagination |
| File Serving | Static frontend hosting |

---

# 🎯 Key Takeaways

- Backend architecture matters.
- Authentication is essential in real applications.
- Middleware is powerful in Go.
- Clean project structure improves scalability.
- Frontend and backend integration is critical.
- Security should never be ignored.

---

# 🚀 Future Improvements

- Role-based access control
- Image upload for posts
- Search functionality
- Like system
- Docker support
- Cloud deployment (Render / AWS)

---

# 👨‍💻 Author

**Sasi Kumar**  
Backend Developer (Learning Go & Full Stack Development)
