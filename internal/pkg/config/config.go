package _020_1_Color_noise

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
	//TODO: заменить на строчку подключения, хранить имя драйвера
	ConnString string `json:"db.connect"`
	MaxConns   int    `json:"db.maxconns"`
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

	c.ConnString = v.GetString("db.connects")
	err = v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("decoding problem: , %v", err)
	}

	return c, err
}


