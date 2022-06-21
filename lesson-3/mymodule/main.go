package mymodule

import (
	"fmt"

	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
)

func printModulePath() {
	fmt.Println("github.com/slava-911/Go.Level-2.v2/lesson3/mymodule")
}
