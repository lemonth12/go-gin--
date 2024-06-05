package httpserv

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "preject/docs"
	"preject/internal/app/router"
	"time"
)

func Start(ctx context.Context, port uint) error {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, lemon")
	})
	//健康检查
	engine.GET("/liveness", func(c *gin.Context) {
		//router.LivenessHandler(&health.Check{
		//	CLASS3DB: db.MysqlClass3Client(),
		//	PFDB:     db.MysqlPfClient(),
		//	RDS:      redis_db.RedisClient(),
		//	CONN:     mq.GetMqConnClients(),
		//})
	})
	engine.GET("/readiness", func(c *gin.Context) {
		// 服务可用性监控
		// 当该接口返回非 2xx 3xx 时，k8s 会不再将用户的流量负载均衡至本容器，待接口恢复后，流量重新引入。
		// 研发通过代码检查服务是否负载超高，需要熔断暂停服务。
		c.Status(http.StatusNoContent)
	})

	//swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//路由组
	v1 := engine.Group("/router")
	{
		//签名验证
		//v1.Use(func(c *gin.Context) { sign.Verify(c) })
		v1.GET("/addr", router.Check)
	}

	//engine.POST("/avl_check", func(c *gin.Context) { sign.Verify(c) }, exapi.Check)

	return start(ctx, engine, port)
}

func start(ctx context.Context, r *gin.Engine, port uint) error {
	srv := &http.Server{
		ReadHeaderTimeout: time.Duration(60) * time.Second,
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           r,
	}

	go func() {
		<-ctx.Done()
		timeout := 10 * time.Second
		ctx, c := context.WithTimeout(context.Background(), timeout)
		defer c()
		_ = srv.Shutdown(ctx) //nolint
	}()

	//log.Logger.Infof("server start %d", port)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
