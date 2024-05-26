use curl to make a request to your running web service.

responds with the list of all albums as JSON.
```
$ curl http://localhost:8080/albums
```
adds an album from JSON received in the request body.
```
$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "1","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

```
locates the album whose ID value matches the id.
```
$ curl http://localhost:8080/albums/1
```