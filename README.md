HOW TO RUN

1. Create new mysql db named go_learning
2. Setting your db config at main.go
3. Deploy / run main in dev with below:
   go run main.go
4. Open postman
   **Get All Product**
    url     : localhost:1999/api/product
    method  : GET
    body    : -

   **Create Product**
    url     : localhost:1999/api/product
    method  : POST
    body    : raw
              {
                  "code" : "221",
                  "name" : "Book Blue",
                  "price" : 10,
              }
    
   **Update Product**
    url     : localhost:1999/api/product/[id]
    method  : PUT
    body    : {
                  "code" : "221",
                  "name" : "Book Blue",
                  "price" : 10,
              }

  **View Product**
      url     : localhost:1999/api/product/[id]
      method  : GET


  **Delete Product**
      url     : localhost:1999/api/product/[id]
      method  : DELETE
