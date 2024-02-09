package main

import (
	"github.com/anjush-bhargavan/golib_api_gateway/config"
	"github.com/anjush-bhargavan/golib_api_gateway/pkg/admin"
	"github.com/anjush-bhargavan/golib_api_gateway/pkg/server"
	"github.com/anjush-bhargavan/golib_api_gateway/pkg/user"
)

func main() {
	server := server.Server()
	config.LoadConfig()
	server.R.LoadHTMLGlob("templates/*.html")
	user.NewUserRoute(server.R)
	admin.NewAdminRoute(server.R)
	server.StartServer()
}
