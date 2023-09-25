package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-web/internal/color"
)

const _UI = `
 ██████╗  ██████╗        ██████╗ ██╗███╗   ██╗      ██╗    ██╗███████╗██████╗ 
██╔════╝ ██╔═══██╗      ██╔════╝ ██║████╗  ██║      ██║    ██║██╔════╝██╔══██╗
██║  ███╗██║   ██║█████╗██║  ███╗██║██╔██╗ ██║█████╗██║ █╗ ██║█████╗  ██████╔╝
██║   ██║██║   ██║╚════╝██║   ██║██║██║╚██╗██║╚════╝██║███╗██║██╔══╝  ██╔══██╗
╚██████╔╝╚██████╔╝      ╚██████╔╝██║██║ ╚████║      ╚███╔███╔╝███████╗██████╔╝
 ╚═════╝  ╚═════╝        ╚═════╝ ╚═╝╚═╝  ╚═══╝       ╚══╝╚══╝ ╚══════╝╚═════╝
`

// Mux http mux
type mux struct {
	engine *gin.Engine
}

func New() (*gin.Engine, error) {
	fmt.Println(color.Blue(_UI))
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r, nil
}
