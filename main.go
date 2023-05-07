package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/", handleWebSocket)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动Web服务失败：", err)
		return
	}
}

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
	}
}
