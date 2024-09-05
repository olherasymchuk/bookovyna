```
go get .
go run .
```
#### URLs to check:
- http://localhost:8080/api/1/authors
- http://localhost:8080/api/1/authors/2
- http://localhost:8080/api/1/authors/404
- http://localhost:8080/api/1/publishers
- http://localhost:8080/api/1/publishers/2
- http://localhost:8080/api/1/publishers/404
- http://localhost:8080/api/1/books
- http://localhost:8080/api/1/books/2
- http://localhost:8080/api/1/books/404

#### Adding new items:
```
curl --location 'http://localhost:8080/api/1/authors' \
--header 'Content-Type: application/json' \
--data '{
    "Name": "Григорій",
    "Surname": "Гавришко"
}'
```

```
curl --location 'http://localhost:8080/api/1/publishers' \
--header 'Content-Type: application/json' \
--data '{
    "Name": "Галичанська книга"
}'
```

```
curl --location 'http://localhost:8080/api/1/books' \
--header 'Content-Type: application/json' \
--data '{
    "Title":        "Асканія-Нова. Історія заповідника",
    "Author_ID":    1,
    "Price":        380.00,
    "Publisher_ID": 1,
    "Published":    2024,
    "ISBN":         "1111111111111"
}'
```