package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func OrbitWebsocet(conn *websocket.Conn, connections []*websocket.Conn, data map[string]interface{}) error {
	message, ok := data["message"].(string)
	if !ok {
		return fmt.Errorf("invalid name format")
	}

	name, ok := data["name"].(string)
	if !ok {
		return fmt.Errorf("invalid name format")
	}

	html := fmt.Sprintf(`
		<div  id="messages" hx-swap-oob="beforeend">
				<div class="message__name">%s</div>
				<div class="message__name">%s</div>
				
		</div>
		`, name, message)
	for _, c := range connections {
		err := c.WriteMessage(websocket.TextMessage, []byte(html))
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	return nil
}
