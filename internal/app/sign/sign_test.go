package sign

import (
	"code.avlyun.org/l/class3/server/config"
	d "code.avlyun.org/l/class3/server/internal/app/db"
	"code.avlyun.org/l/class3/server/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	log.Init("info")
	config.InitConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&collation=utf8mb4_unicode_ci", "class_user_user", "test_2023", "10.251.17.101", "30943", "class3")
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db1, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		log.Logger.Error("mysql conn fail", err)
		panic(err)
	}
	dbMap := make(map[string]*gorm.DB)
	dbMap["class3_db"] = db1
	d.SetClass3MysqlClient(dbMap)
	Sql = d.NewClass3()

	m.Run()
}

func TestVerifyValidToken(t *testing.T) {
	// 创建一个模拟的 Gin 上下文
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	unix := time.Now().Unix()
	randomUUID := uuid.New().String()
	itoa := strconv.Itoa(int(unix))
	token, err := Sign(fmt.Sprintf("%s|%s", randomUUID, itoa))
	if err != nil {
		t.Log(err, "token获取错误")
	}
	t.Log("token", token, "uuid", randomUUID, "timeStamp", itoa)
	c.Request, _ = http.NewRequest("GET", "/path", nil)
	c.Request.Header.Set("token", token)
	c.Request.Header.Set("uuid", randomUUID)
	c.Request.Header.Set("timeStamp", itoa)

	// 调用 Verify 函数
	Verify(c)

	// 检查是否通过了验证
	assert.False(t, c.IsAborted())
}

func TestVerifyInvalidToken(t *testing.T) {
	// 创建一个模拟的 Gin 上下文
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	currentTime := time.Now()
	fiveMinutesAgo := currentTime.Add(-5 * time.Minute)
	unixTimestamp := fiveMinutesAgo.Unix()

	c.Request, _ = http.NewRequest("GET", "/path", nil)
	c.Request.Header.Set("token", "invalid_token")
	c.Request.Header.Set("uuid", "123")
	itoa := strconv.Itoa(int(unixTimestamp))
	t.Log(itoa)
	c.Request.Header.Set("timeStamp", itoa)
	// 调用 Verify 函数
	Verify(c)
	// 检查是否已中止请求
	assert.True(t, c.IsAborted())

}

func TestVerifyMissingTokenjust(t *testing.T) {
	// 创建一个模拟的 Gin 上下文
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	unix := time.Now().Unix()
	randomUUID := uuid.New().String()
	itoa := strconv.Itoa(int(unix))
	token, err := Sign(fmt.Sprintf("%s|%s", randomUUID, itoa))
	if err != nil {
		t.Log(err, "token获取错误")
	}
	token = token + "1"
	t.Log("token", token, "uuid", randomUUID, "timeStamp", itoa)

	c.Request, _ = http.NewRequest("GET", "/path", nil)
	c.Request.Header.Set("token", token)
	c.Request.Header.Set("uuid", randomUUID)
	c.Request.Header.Set("timeStamp", itoa)

	// 调用 Verify 函数
	Verify(c)
	// 检查是否已中止请求
	assert.True(t, c.IsAborted())
}

func TestVerifyMissingToken(t *testing.T) {
	// 创建一个模拟的 Gin 上下文
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	c.Request, _ = http.NewRequest("GET", "/path", nil)
	c.Request.Header.Set("token", "")
	c.Request.Header.Set("uuid", "123")
	c.Request.Header.Set("timeStamp", "123456789")
	// 调用 Verify 函数
	Verify(c)

	// 检查是否已中止请求
	assert.True(t, c.IsAborted())
}

func TestVerifyMissingTime(t *testing.T) {
	// 创建一个模拟的 Gin 上下文
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	c.Request, _ = http.NewRequest("GET", "/path", nil)
	c.Request.Header.Set("token", "sadf")
	c.Request.Header.Set("uuid", "")
	c.Request.Header.Set("timeStamp", "123456789")
	// 调用 Verify 函数
	Verify(c)

	// 检查是否已中止请求
	assert.True(t, c.IsAborted())
}
