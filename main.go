package main

import (
	"fmt"
	"net/http"

	"Usecase/textgenerator"
	"Usecase/websocketserver"
)

func main() {
	qq := int64(12345678)
	text := textgenerator.GenerateText(qq)
	fmt.Println(text)
	http.HandleFunc("/", websocketserver.HandleWebSocket)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动Web服务失败：", err)
		return
	}
}
