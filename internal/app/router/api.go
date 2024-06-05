package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"preject/internal/app/service"
)

type Service interface {
	Logic() int
}

type ServiceStruct struct {
	sert Service
}

var InvoService = ServiceStruct{
	sert: service.NewService(),
}

// @Tags CheckApi
// @Summary 检测测试
// @accept application/json
// @Produce application/json
// @Param data query request.Req true "查询渠道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /router/addr [get]
func Check(c *gin.Context) {
	//requestJson := request.Req{}
	//if err := c.BindJSON(&requestJson); err != nil {
	//	//log.Logger.Errorf("json get failed [%s]", err.Error())
	//	//response.ResponseRequestError(c, "", "json get failed")
	//	return
	//}

	logic := InvoService.sert.Logic()
	fmt.Println(logic)
}
