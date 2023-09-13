# Project Documentation

This document provides detailed information on how to use the REST API for the "Person" resource. Please refer to this documentation for setup instructions, request/response formats, sample API usage, and any known limitations or assumptions made during development.



## Table of Contents

- [Setup Instructions](#setup-instructions)
- [API Endpoints](#endpoints)
- [Request and Response Formats](#request-and-response-formats)
- [Sample API Usage](#sample-api-usage)
- [Assumptions](#assumptions)
- [Known Limitations](#known-limitations)

---

## Setup Instructions

Follow these steps to set up and run the API locally:

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/Clashkid155/HNGX_2S
    
    cd HNGX_2S
    ```

2. **Install Dependencies:**
   ```bash
    go mod download
   ```

3. **Run the API:**
   To run the API locally, you need to set the url of your Postgres db to `PDB_URL`:

    ```bash
    export PDB_URL="Postgres db url"
    go run .
    ```

    > The API will be available locally at `http://127.0.0.1:8080`.

---

## Endpoints

Endpoints available on the API:

- **Create a Person**:
    - **POST /api/**
    - Add a new person to the database.


- **Read a Person**:
    - **GET /api/{name|id}**
    - Retrieve details of a person by name or id.


- **Update a Person**:
    - **PUT /api/{name|id}**
    - Modify details of an existing person by name or id.


- **Delete a Person**:
    - **DELETE /api/{name|id}**
    - Remove a person from the database by name or id.

---

## Request and Response Formats

### Create a Person

#### Request
```http request
POST http://127.0.0.1:8080/api
```
#### Path Parameters

```http request
NONE
```
### Request Body
```json
{
  "name": "Jane"
}
```

**Response Format (Success - 200):**
```json
{
  "id": 10,
  "name": "Luke",
  "message": "Successfully created name",
  "status": "200"
}
```

**Response Format (Conflict - 409):**
```json
{
  "message": "Name: Luke, already exist",
  "status": "409"
}
```
---

### Read a Person

#### Request
```http request
GET http://127.0.0.1:8080/api/{name|id}
```

#### Path Parameters

| Parameter Name | Value   | Description                                  |        Additional         |
|:--------------:|---------|----------------------------------------------|:-------------------------:|
|      name      | string  | Person's name                                |         Required          |
|       id       | integer | Person's  id from returned from the database |  Use either name or this  |

#### Request Body
```text
NONE
```

**Response Format (Success- 200):**
```json
{
  "id": 19,
  "name": "Steam",
  "message": "Successful",
  "status": "200"
}
```

**Response Format (Not Found - 404):**
```json
{
  "message": "person not found",
  "status": "404"
}
```
---

### Update a Person

#### Request
```http request
PUT http://127.0.0.1:8080/api/{name|id}
```

#### Path Parameters

| Parameter Name | Value   | Description                                  |        Additional         |
|:--------------:|---------|----------------------------------------------|:-------------------------:|
|      name      | string  | Person's name                                |         Required          |
|       id       | integer | Person's  id from returned from the database |  Use either name or this  |

#### Request Body
```json
{
  "name": "John The Ripper"
}
```

**Response Format (Success - 200):**
```json
{
  "id": 18,
  "name": "John The Ripper",
  "message": "Successful",
  "status": "200"
}
```

**Response Format (Not Found - 404):**
```json
{
  "message": "person not updated. incorrect input parameter",
  "status": "404"
}
```
---

### Delete a Person

#### Request
```http request
DELETE http://127.0.0.1:8080/api/{name|id}
```

#### Path Parameters

| Parameter Name | Value   | Description                                  |        Additional         |
|:--------------:|---------|----------------------------------------------|:-------------------------:|
|      name      | string  | Person's name                                |         Required          |
|       id       | integer | Person's  id from returned from the database |  Use either name or this  |

#### Request Body
```text
NONE
```

**Response Format (Success - 200):**
```json
{
  "message": "Deleted Successfully",
  "status": "200"
}
```

**Response Format (Not Found - 404):**
```json
{
  "message": "person not deleted. incorrect input parameter",
  "status": "404"
}
```

---

## Sample API Usage

Here are some code examples:

1. **Create a Person:**
   ```go
    package main
    
    import (
    "fmt"
    "net/http"
    "encoding/json"
    )
    
    func main() {
    apiURL := "http://127.0.0.1:8080/api/"
    data := map[string]string{"name": "Amy"}

    resp, err := http.Post(apiURL, "application/json", json.NewEncoder(resp.Body).Encode(data))
    if err != nil {
        fmt.Println("Error making POST request:", err)
        return
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    decoder := json.NewDecoder(resp.Body)
    if err := decoder.Decode(&result); err != nil {
        fmt.Println("Error decoding JSON response:", err)
        return
    }
	   fmt.Println(result)
   }

   ```

2. **Get a Person:**
   ```go
   package main

    import (
    "fmt"
    "net/http"
    "encoding/json"
    )
    
    func main() {
    apiURL := "http://127.0.0.1:8080/api/Ali"
    
    resp, err := http.Get(apiURL)
    if err != nil {
        fmt.Println("Error making GET request:", err)
        return
    }
    defer resp.Body.Close()
    
    var result map[string]interface{}
    decoder := json.NewDecoder(resp.Body)
    if err := decoder.Decode(&result); err != nil {
        fmt.Println("Error decoding JSON response:", err)
        return
    }
    
    fmt.Println(result)
    }

   ```

3. **Delete a Person:**
   ```go
   package main
    
    import (
    "fmt"
    "net/http"
    )
    
    func main() {
    apiURL := "http://127.0.0.1:8080/api/Bob"

    client := &http.Client{}
    req, err := http.NewRequest(http.MethodDelete, apiURL, nil)
    if err != nil {
        fmt.Println("Error creating DELETE request:", err)
        return
    }

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making DELETE request:", err)
        return
    }
    defer resp.Body.Close()

    fmt.Println("HTTP Status Code:", resp.Status)
    }

   ```

---

## Assumptions
- Made name unique in the database definition

## Known Limitations
- Database connection always timeout.
- Input validation.
