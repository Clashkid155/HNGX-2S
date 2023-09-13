# A dead simple, CRUD REST API
This Repo contains code which performs CRUD methods on a Person object.

> **LIVE API:** https://personapi-l9ah.onrender.com/api
>
> **Documentation:** [DOCUMENTATION.md](DOCUMENTATION.md)

## Table of Contents

- [Requirements](#requirements)
- [Getting Started](#getting-started)
    - [1. Clone the Repository](#1-clone-the-repository)
    - [2. Run the API Locally](#2-run-the-api-locally)
- [Sample API Usage](#sample-api-usage)

---

## Requirements

The following are required to run this repo:

- Go 1.18 or higher.
- Git (for cloning the repository).

---

## Getting Started

Follow these steps to set up and run the API locally.

### 1. Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/Clashkid155/HNGX_2S

cd HNGX_2S
```

### 2. Run the API Locally
Dependencies are installed on the first run.

To run the API locally, you need to set the url of your Postgres db to `PDB_URL`:

```bash
export PDB_URL="Postgres db url"
go run .
```

The API server will be available at `http://localhost:8080`.

---

## Sample API Usage
Using Python to create a new Person.
1. **Create a Person**:

   ```python
   import requests

   api_url = "http://127.0.0.1:8080/api/"
   
   data = {
       "name": "Jack",
   }

   response = requests.post(api_url, json=data)
   print(response.json())
   ```