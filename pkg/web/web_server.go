package web

import (
	"github.com/labstack/echo/v4"
)

// Created a new type Server
type Server struct {
	ec *echo.Echo
}

func CreateWebServer() *Server {
	// Created a new instance of Echo
	srv := new(Server)
	srv.ec = echo.New()

	// Requests to be handled on routes
	srv.ec.POST(knewposts, srv.AddPost)
	return srv
}
