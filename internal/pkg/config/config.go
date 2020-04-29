package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

//набор переменных для подключения
var (
	CONFIG = viper.New()
	DB     DataBaseConfig
)

type DataBaseConfig struct {
	Host string `json:"host"`
	Port int `json:"port"`
	Database string `json:"database"`
	User string `json:"user"`
	Password string `json:"password"`
}

func GetDBConfing() (DataBaseConfig, error){
	v := viper.New()
	c := DataBaseConfig{}

	v.SetConfigName("config")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v \n", err))
	}


	c.User = v.GetString("db.user")
	c.Password = v.GetString("db.password")
	c.Host = v.GetString("db.host")
	c.Database = v.GetString("db.database")
	c.Port = v.GetInt("db.port")

	err = v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("decoding problem: , %v", err)
	}

	return c, err
}



func GetTestConfing() (DataBaseConfig, error){
	v := viper.New()
	c := DataBaseConfig{}

	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AddConfigPath("../../../../.")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v \n", err))
	}


	c.User = v.GetString("test.user")
	c.Password = v.GetString("test.password")
	c.Host = v.GetString("test.host")
	c.Database = v.GetString("test.database")
	c.Port = v.GetInt("test.port")

	err = v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("decoding problem: , %v", err)
	}

	return c, err
}