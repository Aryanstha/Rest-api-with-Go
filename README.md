
# Golang REST API with Gin
This is a simple Golang REST API that uses the Gin web framework to create endpoints for managing a list of "info" objects.


## Project Structure
The Project consists of the following files:

`main.go`: The main entry of the application. It defines the routes, handlers, and the HTTP server.

## Dependencies
The project has the following dependencies:

[Gin](https://github.com/gin-gonic/gin): A HTTP web framework for building REST APIs.

## Run

To deploy this project run

```bash
  go run main.go
```

## Endpoints

The API has the following Endpoints:

## Get /infos

Retrives a list of all info objects.

#### Request
```bash
curl -X GET http://localhost:9090/infos

```
#### Response

```json
[
  {
    "id": "1",
    "name": "Aryan",
    "completed": false
  },
  {
    "id": "2",
    "name": "Ajay",
    "completed": false
  },
  {
    "id": "3",
    "name": "Molex",
    "completed": false
  },
  {
    "id": "4",
    "name": "Raj",
    "completed": false
  },
  {
    "id": "5",
    "name": "Rajesh",
    "completed": false
  }
]
```
## POST/infos
Creates a new info object.

#### Request
```bash
curl -X POST -H "Content-Type: application/json" -d 
'{"id":"6","name":"John","completed":false}' http://localhost:9090/infos
```
#### Response
```json
{
  "id": "6",
  "name": "John",
  "completed": false
}
```
## GET /infos/:id
Retrives a single info object by ID.
#### Request
```bash
curl -X GET http://localhost:9090/infos/1
```
#### Response
```json
{
  "id": "1",
  "name": "Aryan",
  "completed": false
}
```
## PATCH /infos/:id
Toggles the `completed` status of the info object with the specified ID.
#### Request
```bash
curl -X PATCH http://localhost:9090/infos/1
```
#### Response
```json
{
  "id": "1",
  "name": "Aryan",
  "completed": false
}
```
## Tech Stack

**Language**: GO Language

