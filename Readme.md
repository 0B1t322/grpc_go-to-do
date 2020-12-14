# GO-TO-DO simple service
---
# запуск:
    go run main.go

# api:
-----------
 - /api/tasker/CreateTask Method=POST
    - 
    - Request
```json
{
   "name": "some name",
   "description": "some desc",
   "token": "some token"
}
```
-
    -   Response
```json
 {
    "error": ""    
 }
 ```
 - /api/tasker/MarkTask Method=PATCH
    - 
    - Request
 ```json
 {
     "id": "2",
     "token": "token",
     "done": "true/false",
 }
 ```
 -
    -   Responce
```json
 {
     "error": "error"
 }
 ```
 
 - /api/tasker/ArchiveTask Method=PUT
    - 
    - Request
 ```json
 {
     "id": "2",
     "token": "token",
     "done": "true/false",
 }
 ```
 -
    -   Responce
```json
 {
     "error": "error"
 }
 ```
 
  - /api/tasker/GetTask Method=GET
    - 
    - Request
 ```json
 {
     "token": "token",
     "filter": "all:/done:true or false/period: [months] [day] - [months] [day]"
 }
 ```
 -
    -   Responce
```json
 {
     "task" {},
     "error": ""
 }
 ```
 
  - /api/tasker/GetAllTasks Method=GET
    - 
    - Request
 ```json
 {
     "token": "token",
 }
 ```
 -
    -   Responce
```json
 {
     "task" {},
     "error": ""
 }