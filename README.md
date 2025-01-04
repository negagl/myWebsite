
# Personal Website API

This project is a personal website API built in Go, designed to manage blogs and projects. It demonstrates a full CRUD implementation using modular and clean code practices.

---

## **Features**
- Manage blogs and projects with full CRUD operations.
- Modular and reusable code.
- Clean and well-structured architecture.

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

### **6. CRUD for Projects**

#### **Endpoints**

1. **`GET /projects`**
   Retrieve a list of all projects.
   ```bash
   curl -X GET http://localhost:8080/projects
   ```

2. **`GET /projects/{id}`**
   Retrieve a specific project by its ID.
   ```bash
   curl -X GET http://localhost:8080/projects/1
   ```

3. **`POST /projects`**
   Create a new project with a request body.
   ```bash
   curl -X POST -H "Content-Type: application/json" \
   -d '{"id": 1, "title": "Project Name", "description": "Project Description", "url": "https://example.com", "status": "in progress"}' \
   http://localhost:8080/projects
   ```

4. **`PUT /projects/{id}`**
   Update an existing project by its ID.
   ```bash
   curl -X PUT -H "Content-Type: application/json" \
   -d '{"title": "Updated Project Name", "description": "Updated Description", "url": "https://example.com", "status": "completed"}' \
   http://localhost:8080/projects/1
   ```

5. **`DELETE /projects/{id}`**
   Delete a project by its ID.
   ```bash
   curl -X DELETE http://localhost:8080/projects/1
   ```

---

## **Future Improvements**
- Add authentication for admin-level access.
- Improve error handling with custom middleware.

---

## **Final Notes**
This project is still in development. It's going to be a website where I'm going to be posting all of my projects and advances.
