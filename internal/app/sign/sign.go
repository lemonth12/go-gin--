package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"time"
)

func Verify(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		//log.Logger.Info("token is null")
		//response.ResponseSignError(c, "", "token is null")
		c.Abort()
		return
	}
	uuid := c.GetHeader("uuid")
	if uuid == "" {
		//log.Logger.Error("uuid is null")
		//response.ResponseSignError(c, "", "uuid is null")
		c.Abort()
		return
	}
	timeStamp := c.GetHeader("timeStamp")
	if timeStamp == "" {
		//log.Logger.Error("timeStamp is null")
		//response.ResponseSignError(c, "", "timeStamp is null")
		c.Abort()
		return
	}
	checkTime, err := strconv.ParseInt(timeStamp, 10, 64)
	if err != nil {
		//log.Logger.Error("time check fail")
		//response.ResponseSignError(c, "", "time check fail")
		c.Abort()
		return
	}
	if math.Abs(float64(time.Now().Unix()-checkTime)) > 5*60 {
		//log.Logger.Error("Signature timeout")
		//response.ResponseSignError(c, "", "Signature timeout")
		c.Abort()
		return
	}
	signToken, err := Sign(fmt.Sprintf("%s|%s", uuid, timeStamp))
	if err != nil {
		//log.Logger.Error("Signature generation failure")
		//response.ResponseSignError(c, "", "Signature generation failure")
		c.Abort()
		return
	}
	if signToken != token {
		//log.Logger.Error("Signature error")
		//response.ResponseSignError(c, "", "Signature error")
		c.Abort()
		return
	}
	c.Next()
}

//var Sql = db.NewClass3()

func Sign(signStr string) (sign string, err error) {
	//获取key
	//err, auto := Sql.GetAuthKey(config.Confs.CFG.Auto)
	//if err != nil {
	//	return
	//}
	//生产sign
	hash := hmac.New(sha256.New, []byte("key"))
	_, err = hash.Write([]byte(signStr))
	if err != nil {
		return
	}
	sign = fmt.Sprintf("%x", hash.Sum(nil))
	return
}
