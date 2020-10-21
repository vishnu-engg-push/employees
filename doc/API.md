### List of API's exposed

#### Mongo CRUD API's

###### Get All Persons

- Endpoint: http://localhost:8080/mongo/persons
- Method: GET

- Response

```
{
    "data": [
        {
            "id": 1,
            "name": "P1",
            "active": true,
            "address": {
                "city": "Navi Mumbai",
                "pincode": 200091
            }
        },
        {
            "id": 2,
            "name": "P1",
            "active": true,
            "address": {
                "city": "New Delhi",
                "pincode": 200007
            }
        }
    ],
    "message": "All Persons",
    "status": 200
}
```

###### Get A Person By Id

- Endpoint: http://localhost:8080/mongo/persons/{id}
- Method: GET
- Response: 

```
{
    "data": {
        "id": 1,
        "name": "P1",
        "active": true,
        "address": {
            "city": "Navi Mumbai",
            "pincode": 200091
        }
    },
    "message": "Found the person",
    "status": 200
}
```

###### Add a Person

- Endpoint: http://localhost:8080/mongo/persons/
- Method: POST
- Request Body:

```
 {
    "id": 3,
    "name": "P3",
    "active": true,
    "address": {
        "city": "Bangalore",
        "pincode": 200009
     }
 }
```
- Response:

```
{
    "message": "Person added Successfully",
    "status": 201
}
```

###### Update a Person's Info

- Endpoint: http://localhost:8080/mongo/persons/{id}
- Method: PUT
- Request Body:
```
{
   "name": "P3",
   "active": true,
   "address": {
      "city": "Mumbai",
      "pincode": 200011
   }
}
```
- Response:
```
{
   "message": "Person Updated Successfully",
   "status": 200
}
```

###### Delete a Person

- Endpoint: http://localhost:8080/mongo/persons/{id}
- Method: DELETE
- Response:

```
{
   "message": "Person deleted successfully",
   "status": 200
}
```

#### Redis Crud API's

###### Get A Person By Id

- Endpoint: http://localhost:8080/redis/persons/{id}
- Method: GET
- Response: 

```
{
    "data": {
        "id": 1,
        "name": "P1",
        "active": true,
        "address": {
            "city": "Navi Mumbai",
            "pincode": 200091
        }
    },
    "message": "Found the person",
    "status": 200
}
```

###### Add a Person

- Endpoint: http://localhost:8080/redis/persons/
- Method: POST
- Request Body:

```
 {
    "id": 3,
    "name": "P3",
    "active": true,
    "address": {
        "city": "Bangalore",
        "pincode": 200009
     }
 }
```
- Response:

```
{
    "message": "Person added Successfully",
    "status": 201
}
```

###### Update a Person's Info

- Endpoint: http://localhost:8080/redis/persons/{id}
- Method: PUT
- Request Body:
```
{
   "name": "P3",
   "active": true,
   "address": {
      "city": "Mumbai",
      "pincode": 200011
   }
}
```
- Response:
```
{
   "message": "Person Updated Successfully",
   "status": 200
}
```

###### Delete a Person

- Endpoint: http://localhost:8080/redis/persons/{id}
- Method: DELETE
- Response:

```
{
   "message": "Person deleted successfully",
   "status": 200
}
```
