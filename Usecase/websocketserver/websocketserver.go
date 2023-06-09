package websocketserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"gopkg.in/ini.v1"

	"github.com/DiodeCN/GOQQAPI/Usecase/textgenerator"
)

func sendResponse(conn *websocket.Conn, groupID int64, message string) {
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
			GroupID: groupID,
			Message: message,
		},
		Echo: 1,
	}
	if err := conn.WriteJSON(resp); err != nil {
		fmt.Println("发送响应出错：", err)
	}
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
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

		var msg struct {
			Message       string `json:"message"`
			GroupID       int64  `json:"group_id"`
			UserID        int64  `json:"user_id"`
			PostType      string `json:"post_type,omitempty"`
			MetaEventType string `json:"meta_event_type,omitempty"`
		}

		if err := json.Unmarshal(message, &msg); err != nil {
			fmt.Println("解析消息出错：", err)
			continue
		}

		if msg.MetaEventType == "heartbeat" {
			continue
		} else {
			fmt.Printf("收到消息：%s\n", message)
		}

		CC := fmt.Sprintf("%d", msg.UserID)

		titleIni, err := ini.Load("title.ini")
		if err != nil {
			log.Println("加载 title.ini 出错：", err)
			continue
		}

		section, err := titleIni.GetSection(CC)
		if err != nil {
			log.Println("获取 section 出错：", err)
			continue
		}

		titleKey, err := section.GetKey("title")
		if err != nil {
			log.Println("获取 titleKey 出错：", err)
			continue
		}

		AA := titleKey.String()

		if msg.Message == "好好好" {
			sendResponse(conn, msg.GroupID, "好好好")
		}
		if msg.Message == "佳乐能力" {
			qq := int64(msg.UserID)
			text := textgenerator.GenerateAbility(qq, AA)
			sendResponse(conn, msg.GroupID, text)
		}
		if msg.Message == "佳乐军事" {
			qq := int64(msg.UserID)
			text := textgenerator.GenerateMilitary(qq, AA)
			sendResponse(conn, msg.GroupID, text)
			log.Println(text)
		}
	}
}
