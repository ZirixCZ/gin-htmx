package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func BroadcastUsers(connections []*websocket.Conn) {
	html := fmt.Sprintf(`
		<div  id="users-list" hx-swap-oob="morphdom">
				<div class="message__name">%d</div>
				
		</div>
		`, len(connections))

	for _, c := range connections {
		err := c.WriteMessage(websocket.TextMessage, []byte(html))
		if err != nil {
			fmt.Println("Error broadcasting message:", err)
		}
	}
}
