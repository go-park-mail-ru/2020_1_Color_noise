package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

//набор переменных для подключения
var (
	CONFIG = viper.New()
	DB DataBaseConfig
)


type DataBaseConfig struct {
	//TODO: заменить на строчку подключения, хранить имя драйвера
	ConnString string `json:"db.connect"`
	MaxConns int `json:"db.maxconns"`
}

func configinit()  {

	CONFIG.SetConfigName("config")
	CONFIG.AddConfigPath(".")
	//TODO: добавить пути к конфигу

	err := CONFIG.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v \n", err))
	}

}

func parseDbConfig(v *viper.Viper, config *DataBaseConfig) {
	//обеспечивает
	v.SetDefault("db.connect", "user=postgres password=password dbname=pinterest sslmode=disable")
	v.SetDefault("db.maxconns", "20")

	config.ConnString = v.GetString("db.connect")

}

func Start() {
	CONFIG.SetConfigName("config")
	CONFIG.AddConfigPath(".")

	configinit()
	parseDbConfig(CONFIG, &DB)

	err := CONFIG.Unmarshal(&DB)
	if err != nil {
		log.Fatalf("decoding problem: , %v", err)
	}
}