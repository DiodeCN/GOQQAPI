package main

import (
    "fmt"
    "log"
    "net/url"

    "github.com/gorilla/websocket"
)

func main() {
    // WebSocket服务器的URL
    url := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/websocket"}

    // 连接到WebSocket服务器
    conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
    if err != nil {
        log.Fatal("无法连接到WebSocket服务器：", err)
    }
    defer conn.Close()

    // 接收来自WebSocket服务器的消息
    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("读取消息出错：", err)
            return
        }
        fmt.Printf("收到消息：%s\n", message)
    }
}
