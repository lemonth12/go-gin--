package config

import (
	"encoding/json"
	"github.com/spf13/viper"
)

var (
	Confs *Configs
)

type Configs struct {
	SqlConf []MysqlConf
}

type MysqlConf struct {
	SourceName string `json:"source_Name"`
	User       string `json:"user"`
	PassWord   string `json:"password"`
	Host       string `json:"host"`
	Database   string `json:"database"`
	Port       string `json:"port"`
}

func InitConf() error {
	viper.AutomaticEnv()
	SetDefaultForTest()
	Confs = &Configs{}

	//mysql
	mySqlJsonString := viper.GetString("MY_SQL")
	var sqlConf []MysqlConf
	err := json.Unmarshal([]byte(mySqlJsonString), &sqlConf)
	if err != nil {
		panic("mysql Environment variable parsing error")
	}
	Confs.SqlConf = sqlConf
	return nil
}

func SetDefaultForTest() {
	viper.AutomaticEnv()
	viper.SetDefault("MY_SQL", "[{\"source_name\":\"db\",\"user\":\"root\",\"password\":\"111111\",\"host\":\"localhost\",\"port\":\"3306\",\"database\":\"彩票\",\"maxidleconns\":\"\",\"maxopenconns\":\"\",\"maxIdleTime\":\"\",\"connMaxLifeTime\":\"\",\"connTimeOut\":\"\"}]\n")

}
