package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func OrbitWebsocet(conn *websocket.Conn, connections []*websocket.Conn, data map[string]interface{}) error {
	message, ok := data["message"].([]interface{})
	if !ok {
		return fmt.Errorf("invalid name format")
	}

	chatMessage := message[0].(string)
	imageSource := message[1].(string)

	name, ok := data["name"].(string)
	if !ok {
		return fmt.Errorf("invalid name format")
	}

	html := fmt.Sprintf(`
		<div id="messages" hx-swap-oob="beforeend">
				<div class="message__name">%s: %s</div>
				<img class="message__image" src="%s" />
		</div>
		`, name, chatMessage, imageSource)

	for _, c := range connections {
		err := c.WriteMessage(websocket.TextMessage, []byte(html))
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	return nil
}
