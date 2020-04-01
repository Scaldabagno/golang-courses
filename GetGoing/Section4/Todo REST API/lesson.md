# Todo REST API

## Request - Response

| Go Function  | HTTP Method/Codes | Endpoint       | Body Request                                   | Body Response                                       |
| ------------ | ----------------- | -------------- | ---------------------------------------------- | --------------------------------------------------- |
| ping()       | GET/{200}         | /ping          |                                                | {<br>"code": 200,<br> "body": "pong"<br>}           |
| ReadAll()    | GET/{200-500}     | /              |                                                | {<br>"name": "nameX",<br> "todo": "todoX"<br>}, ... |
| ReadByName() | GET/{200-500}     | /?name="$name" |                                                | {<br>"name": "$name",<br> "todo": "$todo"<br>}      |
| CreateTodo() | POST/{201-500}    | /              | {<br>"name": "$name",<br> "todo": "$todo"<br>} | {<br>"name": "$name",<br> "todo": "$todo"<br>}      |
| DeleteTodo() | DELETE/{200-500}  | /$name         |                                                | {"Item deleted"}                                    |


## MVC

### model

* connect.go - opens a connection to db and pings it to see if it's alive
* create.go - CreateTodo()/DeleteTodo() for inserting/deleting records from db
* reads.go - ReadByName()/ReadAll() for reading one/all record/s from db
  
### views

* structs.go - structures of response for ping and request for anything else

### controller

* crud.go - management of the REST API
* ping.go - to see if the API is up
* route.go - defines the endpoints
