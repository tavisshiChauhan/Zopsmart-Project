
# GoFr-Based Student Management API
The project implements a HTTP API using the GoFr framework of GoLang for managing student records.
The API facilitates CRUD operations for creating, retrieving, updating, and deleting student information.

## Features
**CRUD Operations:** Create, read, update, and delete operations for student records.

**DB Integration:** Integration with MongoDB for persistent storage.

**Real-World Scenario:** Implements a Student Management API for educational institutions.

## Getting Started
#### Prerequisites
Before running the project make sure you have mongoDB database created.


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
### Endpoints

- **Create Student:**
  - Endpoint: ```POST /students```
  

- **Read Student:**
  - Endpoint: ```GET /students/{ID}```
  

- **Update Student:**
  - Endpoint: ```PUT /students/{ID}```
  

- **Delete Student:**
  - Endpoint: ```DELETE /students/{ID}```
 


## Sequence Diagram




## Testing
### Manual Testing
#### Using Postman

**Create Operation:**
   - Endpoint:`POST /students`
   - URL: `http://localhost:9092/students`
   - Request Body:
```json
{
    "ID" : 7,
    "Name": "Tavisshi Chauhan",
    "Age": 22,
    "Class":"Final Year"
}
```
     
     *Screenshot: Example of creating a new entry using Postman.*


   **Read Operation:**
   - Endpoint:`/GET /students/{ID}`
   - URL: `http://localhost:9092/students/7`

**Update Operation:**
   - Endpoint: `PUT /students/{ID}`
   - URL`http://localhost:9092/students/7`

**Delete Operation:**
   - Endpoint:`DELETE /students/{ID}`
   - URL:`http://localhost:9092/students/7`


