# 🚀 Simple Bank CRUD API Project with Golang, PostgreSQL, Gin, and GORM

<img src="https://user-images.githubusercontent.com/25181517/192149581-88194d20-1a37-4be8-8801-5dc0017ffbbe.png" alt="Golang" height="240"/> <img src="https://cdn.iconscout.com/icon/free/png-256/postgresql-11-1175122.png" alt="PostgreSQL" height="240"/>

Welcome to the Simple Bank CRUD API project! This project demonstrates a robust and scalable backend service using Golang, PostgreSQL, Gin, and GORM. The service includes essential features like email verification, password hashing, and error handling, designed to manage a simple banking system.

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

- Go (1.21.5)
- PostgreSQL (16.1)
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
    POST /api/v1/resources/
    ```

- **Read Account**
    ```
    GET /api/v1/resources/
    ```

- **Read Account By ID**
    ```
    GET /api/v1/resources/:id
    ```


- **Update Account**
    ```
    PUT /api/v1/resources/:id
    ```

- **Delete Account**
    ```
    DELETE /api/v1/resources/:id
    ```



## To-Do

- [x] Implement CRUD operations ✅
- [x] Integrate UUID as Primary Key 🆔
- [x] Set up email verification with Regular Expression 📧
- [x] Add password hashing 🔒
- [ ] Include error handling 🛠️
- [ ] Account balance addition 💰
- [ ] Transfer balance between accounts 🔄
- [ ] Implement JWT authentication 🔐
- [ ] Add rate limiting 🚀
- [ ] Improve documentation with detailed examples 📚


