package main

import server "github.com/pu4mane/web-service-gin"

func main() {
	server := server.NewRecordServer()
	server.Run("localhost:8080")
}
