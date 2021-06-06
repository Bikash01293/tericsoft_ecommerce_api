# tericsoft_ecommerce_api

## Summary

The tericsoft_ecommerce_api Golang API is used for retrieving content from MariaDB. The API is currently using Golang version 1.14, Go SQL driver 1.5.0 and MariaDB 10.4.17. 
To check the response I used Postman.

## Installation

1) Run main.go. 
2) Import `teric.sql` file in the mySQL to create the database and tables

## Endpoints for tericsoft_ecommerce_api

The endpoints work on `localhost:8000`.

### Register User

Use `POST` with endpoint `/api/register`. Using this, we can create Users.

## Payload for Creating a User
  ```
 {
    "name" : "abc",
    "email" : "abc.abc@gmail.com",
    "username" : "username",
    "password" : "password"
}
```

### Login User

Use `POST` with endpoint `/api/login`. Using this, we can login Users.

## Payload for Creating a User
  ```
 {
    "email" : "abc.abc@gmail.com",
    "password" : "password"
 }
  ```

### Logout User

Use `POST` with endpoint `/api/logout`. Using this, we can logout Users.


### Create Products

Use `POST` with endpoint `/api/products/create`. Using this, we can create products.

## Payload for Creating a product
  ```
  {
    "product_name": "Cabbage",
    "products_quantity": "38",
    "product_description": "Smooth-leafed, firm-headed green cabbages are the most common, with smooth-leafed purple cabbages and crinkle-leafed savoy cabbages of both colours being rarer",
    "product_price": 27,
    "sku": "IUTUYT",
    "category": [
        {
            "categoryId": 1
        }
    ]
}
```

### Retreive a Product with Categories

Use `GET` with endpoint `/api/products/show/{id}`. Using this, we can retreive a product with the ID.

### Retreive all products with Categories

Use `GET` with endpoint `/api/products/show`. Using this, we can retreive all products.

## Update a Product

Use `PUT` with endpoint `/api/products/update/{id}`. Using this, we can update a product with the products ID.

## Payload for Updating a product
  ```
  {
    
    "product_name": "lemon",
    "products_quantity": "87",
    "product_description": "Capsicum, is also known as red pepper or chili pepper, is an herb. The fruit of the capsicum plant is used to make medicine. Capsicum is most commonly used for rheumatoid arthritis (RA), osteoarthritis, and other painful conditions.",
    "product_price": 40
}
```

## Delete a product

Use `DELETE` with endpoint `/api/products/delete/{id}`. Using this, we can DELETE the products with products ID.


### Create Categories

Use `POST` with endpoint `/api/categories/create`. Using this, we can create Categories.

## Payload for Creating a Categories
  ```
  {
    "product_name": "Cabbage",
    "products_quantity": "38",
    "product_description": "Smooth-leafed, firm-headed green cabbages are the most common, with smooth-leafed purple cabbages and crinkle-leafed savoy cabbages of both colours being rarer",
    "product_price": 27,
    "sku": "IUTUYT",
    "category": [
        {
            "categoryId": 1
        }
    ]
}
```

### Retreive a Categories

Use `GET` with endpoint `/api/categories/show/{id}`. Using this, we can retreive an Categories with the ID.

### Retreive all Categories

Use `GET` with endpoint `/api/categories/show`. Using this, we can retreive all Categories.

## Update a Category

Use `PUT` with endpoint `/api/categories/update/{id}`. Using this, we can update categories with the Categories ID.

## Payload for Updating a Categories
  ```
  {
    
    "product_name": "lemon",
    "products_quantity": "87",
    "product_description": "Capsicum, is also known as red pepper or chili pepper, is an herb. The fruit of the capsicum plant is used to make medicine. Capsicum is most commonly used for rheumatoid arthritis (RA), osteoarthritis, and other painful conditions.",
    "product_price": 40
}
```

## Delete a Categories

Use `DELETE` with endpoint `/api/categories/delete/{id}`. Using this, we can delete with the Categories ID.
