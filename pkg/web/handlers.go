package web

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (srv *Server) AddPost(c echo.Context) error {
	// Stored the form value entered by user into a new variable
	username := c.FormValue("username")
	post := c.FormValue("post")

	// Connected to mysql server
	db, err := sql.Open("mysql", "root:BM@mysql53@tcp(127.0.0.1:3306)/fora")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Inserted the entered form value into mysql table
	insert, err := db.Query("INSERT INTO `newposts` VALUES(?, ?)", username, post)
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	fmt.Println("Data inserted successfully")

	return c.String(http.StatusOK, "User added a post successfully")
}

// Made function to start server at specified address
func (srv *Server) StartListeningForRequests(addr string) error {
	if err := srv.ec.Start(addr); err != nil {
		return err
	}
	return nil
}
