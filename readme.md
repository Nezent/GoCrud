# 🚀 CRUD API Project with Golang, PostgreSQL, Gin, and GORM

![Golang](https://raw.githubusercontent.com/github/explore/main/topics/go/go.png) ![PostgreSQL](https://cdn.iconscout.com/icon/free/png-256/postgresql-11-1175122.png)

Welcome to the CRUD API project! This project demonstrates a robust and scalable backend service using Golang, PostgreSQL, Gin, and GORM. The service includes essential features like email verification, password hashing, and error handling.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Setup](#setup)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [To-Do](#to-do)

## Features

- **CRUD Operations**: Create, Read, Update, Delete operations for resources.
- **UUID as Primary Key**: Using Google UUID for unique resource identification.
- **Email Verification**: Ensure user authenticity with email verification.
- **Password Hashing**: Secure user passwords with hashing.
- **Error Handling**: Comprehensive error handling for better debugging.
- **Flexibility**: Utilizes Gin and GORM for a flexible and powerful API.

## Getting Started

### Prerequisites

- Go (1.16+)
- PostgreSQL
- Git

### Installation

1. **Clone the repository**
    ```bash
    git clone https://github.com/Nezent/GoCrud.git
    cd GoCrud
    ```

2. **Install dependencies**
    ```bash
    go mod tidy
    ```

3. **Set up PostgreSQL**

    Create a PostgreSQL database and update the `config.yml` file with your database credentials.


## Setup

### Configuration

Configure your database and other settings in the `docker-compose.yml` file:

```yaml
database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: postgres"
server:
  port: 8080
```

## Usage

### Running the Application

```bash
make run
```
The server will start on http://localhost:8080.

## Endpoints

- **Create Account**
    ```
    POST /
    ```

- **Read Account**
    ```
    GET /
    ```

- **Update Account**
    ```
    PUT /:id
    ```

- **Delete Account**
    ```
    DELETE /:id
    ```


## To-Do

- [x] Implement CRUD operations
- [x] Integrate UUID as Primary Key
- [ ] Set up email verification
- [ ] Add password hashing
- [ ] Include error handling
- [ ] **Implement JWT authentication** 🔐
- [ ] **Add rate limiting** 🚀
- [ ] **Improve documentation with detailed examples** 📚

