package websocketserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/DiodeCN/GOQQAPI/Usecase/textgenerator"

)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		fmt.Println("升级WebSocket连接失败：", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("读取消息出错：", err)
			return
		}
		fmt.Printf("收到消息：%s\n", message)

		var msg struct {
			Message string `json:"message"`
			GroupID int64  `json:"group_id"`
			UserID  int64  `json:"user_id"`
		}
		if err := json.Unmarshal(message, &msg); err != nil {
			fmt.Println("解析消息出错：", err)
			continue
		}
		if msg.Message == "好好好" {
			resp := struct {
				Action string `json:"action"`
				Params struct {
					GroupID int64  `json:"group_id"`
					Message string `json:"message"`
				} `json:"params"`
				Echo int `json:"echo"`
			}{
				Action: "send_group_msg",
				Params: struct {
					GroupID int64  `json:"group_id"`
					Message string `json:"message"`
				}{
					GroupID: msg.GroupID,
					Message: "好好好",
				},
				Echo: 1,
			}
			if err := conn.WriteJSON(resp); err != nil {
				fmt.Println("发送响应出错：", err)
			}
		}
		if msg.Message == "佳乐能力" {
			qq := int64(msg.UserID)
			text := textgenerator.GenerateText(qq)
			fmt.Println(text)
		}
	}
}
