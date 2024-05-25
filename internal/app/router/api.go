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

func Check(c *gin.Context) {
	//requestJson := request.Req{}
	//if err := c.BindJSON(&requestJson); err != nil {
	//	//log.Logger.Errorf("json get failed [%s]", err.Error())
	//	//response.ResponseRequestError(c, "", "json get failed")
	//	return
	//}

	fmt.Println("2222222222222222222")
	logic := InvoService.sert.Logic()
	fmt.Println(logic)
}
