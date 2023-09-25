package main

import "go-gin-web/internal/pkg/core"

/*
*
程序入口
*/

func main() {
	engine, err := core.New()
	if err == nil {
		engine.Run()
	}

}
