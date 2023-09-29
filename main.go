package main

import (
	"chatting/api"
	"chatting/api/ws"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	app := gin.Default()

	app.Static("/static", "./static")
	app.LoadHTMLGlob("templates/*")

	RegisterApiRoutes(app)
	RegisterWebSocketRoutes(app)

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout.html", nil)
	})

	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func RegisterApiRoutes(app *gin.Engine) {
	app.GET("/api/ping", api.Ping)
	app.GET("/api/increment", api.Increment)
	app.POST("/api/nameset", api.Nameset)
}

var connections []*websocket.Conn

func RegisterWebSocketRoutes(app *gin.Engine) {
	app.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			panic(err)
		}

		connections = append(connections, conn)
		ws.BroadcastUsers(connections)
		go handleWebSocket(conn)
	})
}

func handleWebSocket(conn *websocket.Conn) {
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			for i, c := range connections {
				if c == conn {
					connections = append(connections[:i], connections[i+1:]...)
					break
				}
			}
			ws.BroadcastUsers(connections)
			break
		}

		var data map[string]interface{}
		err = json.Unmarshal(message, &data)
		if err != nil {
			fmt.Println("Error parsing message:", err)
			continue
		}

		event, ok := data["event"].(string)
		if !ok {
			fmt.Println("Error parsing message:", err)
			continue
		}

		switch event {
		case "orbit":
			if ws.OrbitWebsocet(conn, connections, data) != nil {
				return
			}
			ws.BroadcastUsers(connections)
		}
	}
}
