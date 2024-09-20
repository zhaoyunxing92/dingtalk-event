package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhaoyunxing92/dingtalk-event/logger"
	"github.com/zhaoyunxing92/dingtalk/v2"
	"github.com/zhaoyunxing92/dingtalk/v2/crypto"
	"go.uber.org/zap"
)

var client = dingtalk.NewClient(23434, "dinggkluifooodnghss9",
	"1wptnXj8VFHQKayDdcCkgbD-",
	dingtalk.WithLevel(zap.DebugLevel))

func main() {
	router := gin.New()
	log := logger.InitLogger(zap.DebugLevel)

	router.Use(logger.GinLogger(log), logger.GinRecovery(log, true))

	//router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
	//	return fmt.Sprintf("%s - %s | %d | %s | %s | %s \n",
	//		params.TimeStamp.Format("2006-01-02 03:04:05"),
	//		params.ClientIP,
	//		params.StatusCode,
	//		params.Path,
	//		params.Method,
	//		params.Request.Proto,
	//	)
	//}))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	router.POST("/dingtalk/event", func(ctx *gin.Context) {
		sign := ctx.Query("msg_signature")
		timestamp := ctx.Query("timestamp")
		nonce := ctx.Query("nonce")

		en := crypto.DingTalkEncrypt{}
		if err := ctx.BindJSON(&en); err != nil {
			return
		}

		dingCrypto := client.GetDingTalkCrypto("mUBLh4IahPjvQKOP", "ge4elNG3RwMTwV6MNMgpdvtMMeMpg3CuHlifOB9OeA6")

		if msg, err := dingCrypto.Decrypt(en.Encrypt, sign, timestamp, nonce); err != nil {
			log.Error("ding event err:", zap.Error(err))
			return
		} else {
			log.Info("ding event:", zap.String("res", msg))
		}

		encrypt, err := dingCrypto.Encrypt("success")
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.JSON(200, encrypt)
	})

	router.Run()
}
