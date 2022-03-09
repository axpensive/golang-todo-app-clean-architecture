package main

import (
	"log"

	"github.com/axpensive/golang-todo-app-clean-architecture/infrastructure"
)

func main() {
	// サーバー接続呼び出し
	log.Println("Server start!!!")
	infrastructure.StartAppServer()
}
