# Simple SQL Database Controller
# Description
This project is a simple database controller which manipalation CRUD command to SQL Server.

In currently, we only implement MySQL, or you can change SQL driver for any SQL database you need.

# Dependency Library
We are using Gorilla mux to be HTTP request router.
See [Gorillla Mux](https://github.com/gorilla/mux).

And we using go-sql-driver for interface between database and project.
See [MySQL Driver for GO](https://github.com/go-sql-driver/mysql).

In authentication part, we using [JWT(JSON web token)](https://jwt.io/introduction/).
See [JWT Library](https://github.com/robbert229/jwt).

# API Introduction
- Authentiaction
  - We using JWT as authentiaction.
  - After first authorize client will get a token.
  - Client need send both of data and token to server every time.
  - How to get JWT?
    - Use POST mehtod.
    - The URL as ``` http://127.0.0.1:8080/auth ```
    - Payload is a Form data of request which composed by database username and password.
    
- GET
  - Use GET method to query specific data from database.
  - How to query data?
    - Use GET method.
    - The URL as ``` https://127.0.0.1:8080/admin/{table_name}/{constraint} ```
    - Client needs provide {table_name} and {constraint} in URL to query specific data.
  
- POST
  - Use PUT method to insert a new data into database.
  - How to insert data?
    - Use POST method.
    - URL: ``` https://127.0.0.1:8080/admin/{table_name}/```
    - Payload as {new_data}: ``` { "name":"Rex", "gender":"male", "title":"software engineer" } ```
    - Client needs provide {table_name} in URL and {new_data} in payload to insert new data.

- PUT
  - Use PUT method to update exist data in database.
  - How to update data?
    - Use PUT method.
    - URL: ``` https://127.0.0.1:8080/admin/{table_name}/{constraint}```
    - Payload as {update_data}: ``` { "title":"senior software engineer" } ```
    - Client needs provide {table_name} and {update_data} in payload to update specific data.


- DELETE
  - Use DELETE method to remove exist data in database
  - How to delete data?
      - Use DELETE method.
      - The URL as ``` https://127.0.0.1:8080/admin/{table_name}/{constraint} ```
      - Client needs provide {table_name} and {constraint} in URL to delete specific data.
 
# License
MIT
