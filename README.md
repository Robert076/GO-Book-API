**How to run the API**

1. go run main.go

2. Open another terminal

3. GET request: curl localhost:8080/books

4. POST request: curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"

5. GET by ID: curl localhost:8080/books/2
