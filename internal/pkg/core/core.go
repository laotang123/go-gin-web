package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-web/internal/color"
	"math/rand"
	"strconv"
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
	r.GET("/user/ping", func(c *gin.Context) {
		data := "pong" + strconv.Itoa(rand.Int())
		c.JSON(200, gin.H{
			"message": "success",
			"code":    200,
			"data":    data,
		})
	})
	return r, nil
}
