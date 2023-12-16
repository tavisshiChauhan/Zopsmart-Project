
# GoFr-Based Student Management API
The project implements a HTTP API using the GoFr framework of GoLang for managing student records.
The API facilitates CRUD operations for creating, retrieving, updating, and deleting student information.

## Features
**CRUD Operations:** Create, read, update, and delete operations for student records.

**DB Integration:** Integration with MongoDB for persistent storage.

**Real-World Scenario:** Implements a Student Management API for educational institutions.

## Getting Started
Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.
initialise your go module using:
```bash
  go mod init test-service

```
If you intend to push your code to github, it is recommended to name your module like this: 
```bash
  go mod init github.com/{USERNAME}/{REPO}

```

## Installation
Install project dependencies:

```bash
  go get ./...

```
    
## Usage
Start the API server:

```bash
  go run main.go
```

## Endpoints
Get Student by ID:
```bash
  GET /students/{ID}
```
Create Student:
```bash
  POST /students
```
Update Student:
```bash
  PUT /students/{ID}
```
Delete Student:
```bash
  DELETE /students/{ID}
```
## Sequence Diagram

<img width="503" alt="uml" src="https://github.com/tavisshiChauhan/Zopsmart-Project/assets/125811955/c4c44cde-471a-4e60-b64d-a6bc863e1ff8">



## Testing
### Manual Testing
#### Using Postman

**Create Operation:**
   - Endpoint:
   - `/api/create`
   - `http://localhost:9092/students`
     
     *Screenshot: Example of creating a new entry using Postman.*
    <img width="640" alt="Z1" src="https://github.com/tavisshiChauhan/Zopsmart-Project/assets/125811955/e26eaa47-acb0-4099-ae1a-d3546777e9c7">

   **Read Operation:**
   - Endpoint:
   - `/api/read/{id}`
   - `http://localhost:9092/students/7`

**Update Operation:**
   - Endpoint:
   - `/api/update/{id}`
   - `http://localhost:9092/students/7`

**Delete Operation:**
   - Endpoint:
   - `/api/delete/{id}`
   - `http://localhost:9092/students/7`


