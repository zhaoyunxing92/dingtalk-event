package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhaoyunxing92/dingtalk/v2"
	"go.uber.org/zap"
)

var client = dingtalk.NewClient(1244553273, "dingkjy4w80esdwgjuyo",
	"bDKa_nfJg3zYRsFrj-wTohTuoJCtxTEHaGmybYF9vgaVAZJOz-mICsLGStB288nW", dingtalk.WithLevel(zap.DebugLevel))

type Encrypt struct {
	Text string `json:"encrypt"`
}

func main() {
	rout := gin.Default()
	rout.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	rout.POST("/dingtalk/event", func(ctx *gin.Context) {
		fmt.Println(ctx)
		sign := ctx.Query("msg_signature")
		timestamp := ctx.Query("timestamp")
		nonce := ctx.Query("nonce")
		fmt.Println(sign)
		fmt.Println(timestamp)
		fmt.Println(nonce)

		//buf := make([]byte, 4096)
		//n, _ := ctx.Request.Body.Read(buf)
		//fmt.Println(string(buf[0:n]))

		en := Encrypt{}
		ctx.BindJSON(&en)
		fmt.Println(en)
		/*post_gwid := c.PostForm("name")
		fmt.Println(post_gwid)*/
		//msg_signature=4af305f33739df19e07c705fbb0c6272072fbe77&timestamp=1640447283988&nonce=Cnlz0NPV

	})

	rout.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
