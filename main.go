package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	var fileName string

	flag.StringVar(&fileName, "conf", "config.toml", "Configuration file to start game")
	flag.Parse()
	glog.Errorln("Configuration is", fileName)

	// load config
	err := ParseToml(fileName)
	if err != nil {
		glog.Errorln("配置文件.toml出错")
		glog.Fatal(err)
	}
	InitAccount(GetWxAccount().Account)

	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/wechatcheck", check)

	e.Start(GetWxAccount().Port)
}

func check(c echo.Context) error {
	longUrl := c.QueryParam("url")
	return c.JSON(http.StatusOK, Check(longUrl))
}
