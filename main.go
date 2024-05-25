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
	log.Infof("%s [version: %s commitId: %s] start", appName, Version, CommitId)
	viper.AutomaticEnv()
	config.InitConf()
	mysqlClients := db.InitMysql(config.Confs.SqlConf)
	db.SetClass3MysqlClient(mysqlClients)
}

func main() {
	initGlobal()
	cmd.Execute(Version, CommitId)
}
