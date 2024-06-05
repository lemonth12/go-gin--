package db

import (
	"gorm.io/gorm"
	"preject/config"
	"preject/pkg/conn/db"
	"preject/pkg/log"
)

const (
	Class3 = "db"
	pf     = "pf_db"
)

var (
	mysqlpfClients     *gorm.DB
	mysqlClass3Clients *gorm.DB
)

//func SetPfMysqlClient(client map[string]*gorm.DB) {
//	mysqlpfClients = client[pf]
//}
//
//// 对外暴露sql
//
//func MysqlPfClient() *gorm.DB {
//	return mysqlpfClients
//}

func SetClass3MysqlClient(client map[string]*gorm.DB) {
	mysqlClass3Clients = client[Class3]
}

// // 对外暴露sql
func MysqlClass3Client() *gorm.DB {
	//db, err := mysqlClass3Clients.DB()
	//if err != nil {
	//	log.Errorf("Database connection error：%v", err)
	//}
	//err = db.Ping()
	//if err != nil {
	//	log.Errorf("Database ping failure：%v", err)
	//	mysql := InitMysql(config.Confs.SqlConf)
	//	SetClass3MysqlClient(mysql)
	//}
	return mysqlClass3Clients
}

var Clients = make(map[string]*gorm.DB)

// 连接数据库
func InitMysql(confs []config.MysqlConf) map[string]*gorm.DB {

	for _, mysqlConf := range confs {
		err, mysql := db.GormMysql(mysqlConf)
		if err != nil {
			log.Logger.Errorf("%s conn fail:%v", mysqlConf.SourceName, err)
			panic(err)
		}
		//将创建的sql加入mao
		Clients[mysqlConf.SourceName] = mysql
		//给数据库创建特定表
		//if mysqlConf.SourceName == Class3 {
		//	err := mysql.AutoMigrate(
		//		system.TableName{},
		//	)
		//	if err != nil {
		//		log.Errorf("register table failed", err)
		//		panic(err)
		//	}
		//}
	}
	return Clients
}
