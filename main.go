package main

import (
	"flag"
	"github.com/spf13/viper"
	"preject/cmd"
	"preject/config"
	"preject/internal/app/db"
	"preject/pkg/log"
)

var Version = ""
var CommitId = ""

func initGlobal() {
	appName := "preject-name"
	flag.Parse()
	viper.AutomaticEnv()
	config.InitConf()
	log.Init("info")
	log.Logger.Infof("%s [version: %s commitId: %s] start", appName, Version, CommitId)
	mysqlClients := db.InitMysql(config.Confs.SqlConf)
	db.SetClass3MysqlClient(mysqlClients)
}

// @title Swagger Example API
// @version 1.0.0
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	initGlobal()
	cmd.Execute(Version, CommitId)
}
