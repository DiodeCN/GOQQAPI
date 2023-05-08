package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
    "encoding/json"
)


func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    // 将HTTP连接升级为WebSocket
    conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
    if err != nil {
        fmt.Println("升级WebSocket连接失败：", err)
        return
    }
    defer conn.Close()

    // 无限循环，读取传入的WebSocket消息
    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("读取消息出错：", err)
            return
        }
        fmt.Printf("收到消息：%s\n", message)

        // 检查消息内容
        var msg struct {
            Message string `json:"message"`
            GroupID int64  `json:"group_id"`
        }
        if err := json.Unmarshal(message, &msg); err != nil {
            fmt.Println("解析消息出错：", err)
            continue
        }
        if msg.Message == "好好好" {
            // 返回JSON响应
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
    }
}



func main() {
	http.HandleFunc("/", handleWebSocket)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动Web服务失败：", err)
		return
	}
}

// 今天都怪那个女人。