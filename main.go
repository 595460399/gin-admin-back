package main

import (
	"gin-admin-back/config"
	"gin-admin-back/initialize"
	"net/http"
	"time"
)

func main() {
	initialize.InitMysql(config.Dbconfig.Admin)
	defer initialize.DEFAULTDB.Close()
	initialize.InitRouter()
	s := &http.Server{
		Addr: "8888",
		Handler: initialize.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
