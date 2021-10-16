package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("init")

}
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./templates/")

	r.GET("/", home)
	r.GET("/create/", create)
	r.GET("/:room/", room)

	r.GET("/:room/ws/", room_ws)

	r.Run()

}
