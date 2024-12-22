# Personal Website API

This project is a personal website API built in Go, designed to manage blogs. It demonstrates a full CRUD implementation using modular and clean code practices.

---

## **Features**
- List all blogs.
- Retrieve a blog by its ID.
- Create a new blog with validations.
- Update an existing blog.
- Delete a blog by its ID.

---

## **Endpoints**

### **1. `GET /blogs`**
Retrieve a list of all blogs.

#### **Example Request**
```bash
curl -X GET http://localhost:8080/blogs
```

#### **Example Response**
```json
[
    {
        "id": 1,
        "title": "First Blog",
        "content": "This is the first blog post"
    }
]
```

---

### **2. `GET /blogs/{id}`**
Retrieve a specific blog by its ID.

#### **Example Request**
```bash
curl -X GET http://localhost:8080/blogs/1
```

#### **Example Response**
```json
{
    "id": 1,
    "title": "First Blog",
    "content": "This is the first blog post"
}
```

#### **Error Response (Blog Not Found)**
```bash
curl -X GET http://localhost:8080/blogs/99
```
```text
Blog not found
```

---

### **3. `POST /blogs`**
Create a new blog.

#### **Example Request**
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"id": 2, "title": "Second Blog", "content": "This is the second blog post"}' \
http://localhost:8080/blogs
```

#### **Example Response**
```json
{
    "id": 2,
    "title": "Second Blog",
    "content": "This is the second blog post"
}
```

#### **Error Responses**
- **Duplicated ID**:
  ```
  ID must be unique
  ```

- **Empty Title**:
  ```
  Title cannot be empty
  ```

- **Empty Content**:
  ```
  Content cannot be empty
  ```

---

### **4. `PUT /blogs/{id}`**
Update an existing blog by its ID.

#### **Example Request**
```bash
curl -X PUT -H "Content-Type: application/json" \
-d '{"title": "Updated Blog", "content": "Updated content"}' \
http://localhost:8080/blogs/1
```

#### **Example Response**
```json
{
    "id": 1,
    "title": "Updated Blog",
    "content": "Updated content"
}
```

#### **Error Responses**
- **Blog Not Found**:
  ```text
  Blog not found
  ```

- **Empty Title or Content**:
  ```
  Invalid input data
  ```

---

### **5. `DELETE /blogs/{id}`**
Delete a blog by its ID.

#### **Example Request**
```bash
curl -X DELETE http://localhost:8080/blogs/1
```

#### **Example Response**
```json
{
    "id": 1,
    "title": "First Blog",
    "content": "This is the first blog post"
}
```

#### **Error Response (Blog Not Found)**
```bash
curl -X DELETE http://localhost:8080/blogs/99
```
```text
Blog not found
```

---

## **Setup**

### **1. Prerequisites**
- Go installed (version 1.19 or higher).

### **2. Clone the repository**
```bash
git clone https://github.com/yourusername/yourrepository.git
cd yourrepository
```

### **3. Run the server**
```bash
go run cmd/main.go
```

### **4. Test the API**
Use tools like `curl`, Postman, or any HTTP client to test the API endpoints.

---

## **Future Improvements**
- Add authentication for admin-level access.
- Extend the API to include project management.
- Improve error handling with custom middleware.

---

## **Final notes**
This project is still in development. It's going to be a website where i'm going to be posting all of my projects and advances.
