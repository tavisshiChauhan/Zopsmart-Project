
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

## Testing
