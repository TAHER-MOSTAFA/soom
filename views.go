package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func create(c *gin.Context) {
	room := createRoom()
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s", room.ID))
}

func room(c *gin.Context) {
	_, ex := Rooms[c.Param("room")]
	if !ex {
		c.Redirect(http.StatusFound, "/create")
	}
	c.HTML(http.StatusOK, "room.html", gin.H{})
}

func create_answer(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sdpChan := make(chan string)
	go core_handeler(sdpChan)
	sdpChan <- user.Offer

	c.JSON(http.StatusOK, gin.H{"answer": <-sdpChan})
}

func room_ws(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	fmt.Println("connected")
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		log.Printf("%d", mt)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
