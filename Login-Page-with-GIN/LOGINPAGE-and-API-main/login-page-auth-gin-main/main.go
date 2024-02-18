package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"main.go/handlers"
)

func main() {

	s := gin.Default()

	s.LoadHTMLGlob("site/pages/*.html")

	s.Static("/site/styles", "./site/styles")

	s.GET("/login", handlers.GetLogin)
	s.POST("/login", handlers.PostLogin)
	s.GET("/signup", handlers.GetSignup)
	s.POST("/signup", handlers.PostSignup)
	s.GET("/", handlers.GetHome)
	s.POST("/logout", handlers.Logout)

	s.Run(":5050") //start server

	fmt.Println("server started on http://Localhost:5050")

}
