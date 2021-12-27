package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhaoyunxing92/dingtalk/v2"
	"github.com/zhaoyunxing92/dingtalk/v2/crypto"
	"go.uber.org/zap"
)

var client = dingtalk.NewClient(1244553273, "dingkjy4w80esdwgjuyo",
	"bDKa_nfJg3zYRsFrj-wTohTuoJCtxTEHaGmybYF9vgaVAZJOz-mICsLGStB288nW", dingtalk.WithLevel(zap.DebugLevel))

func main() {
	rout := gin.Default()

	rout.POST("/dingtalk/event", func(ctx *gin.Context) {
		fmt.Println(ctx)
		sign := ctx.Query("msg_signature")
		timestamp := ctx.Query("timestamp")
		nonce := ctx.Query("nonce")

		en := crypto.DingTalkEncrypt{}
		if err := ctx.BindJSON(&en); err != nil {
			return
		}

		fmt.Println(en.String())

		dingCrypto := client.GetDingTalkCrypto("mUBLh4IahPjvQKOP", "ge4elNG3RwMTwV6MNMgpdvtMMeMpg3CuHlifOB9OeA6")

		if msg, err := dingCrypto.Decrypt(en.Encrypt, sign, timestamp, nonce); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(msg)
		}

		encrypt, err := dingCrypto.Encrypt("success")
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.JSON(200, encrypt)
	})

	rout.Run()
}
