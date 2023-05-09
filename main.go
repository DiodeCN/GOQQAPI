package main

import (
	"fmt"
	"net/http"

	"github.com/DiodeCN/GOQQAPI/Usecase/websocketserver"
)

func main() {
	http.HandleFunc("/", websocketserver.HandleWebSocket)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动Web服务失败：", err)
		return
	}
}
