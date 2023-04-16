package main

import (
	_ "github.com/go-sql-driver/mysql"
	"test-sql/pkg/web"
)

func main() {
	// We created a server and started it on port 3000
	srv := web.CreateWebServer()
	if err := srv.StartListeningForRequests(":3000"); err != nil {
		panic(err)
	}
}
