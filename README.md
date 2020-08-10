# Simple MySQL Database Controller
# Description
This project is a simple database controller which manipalation CRUD command to MySQL.

# Dependency Library
We are using Gorilla mux to be HTTP request router.
See [Gorillla Mux](https://github.com/gorilla/mux).

And we using go-sql-driver for interface between database and project.
See [MySQL Driver for GO](https://github.com/go-sql-driver/mysql).

In authentication part, we using [JWT(JSON web token)](https://jwt.io/introduction/).
See [JWT Library](https://github.com/robbert229/jwt)

# API Introduction
- Authentiaction
  - We using JWT as authentiaction
  - After first authorize client will get a token.
  - Client need send both of data and token to server every time.
  - How to get JWT?
    - Use POST mehtod and the URL as ``` http://127.0.0.1:8080/auth ```
    - Payload is a Form data of request which composed by database username and password.
    
- GET
  - Use GET method to query specific data from database.
  - How to query data?
  
- POST
  - Use PUT method to insert a new data into database.
  - How to insert data?


- PUT
  - Use PUT method to update exist data in database.
  - How to update data?

- DELETE
  - Use DELETE method to remove exist data in database
  - How to delete data?
 
# License
MIT
