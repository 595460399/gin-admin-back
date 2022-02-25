package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Admin struct {
	UserName string
	Password string
	Path     string
	Dbname   string
	Config   string
}

type Config struct {
	Admin Admin
}

var Dbconfig Config

func init() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("./config/dbconfig/")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
	})
	if err := v.Unmarshal(&Dbconfig); err != nil {
		fmt.Println(err)
	}
}
