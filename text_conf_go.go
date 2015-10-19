package main

var appconfgo=`package conf

import (
	"github.com/Unknwon/goconfig"
	"fmt"
)
var Config *goconfig.ConfigFile
var (
	Appname  string
	Httpport int
	Runmode  string
)

func init() {
	var err error
	Config,err =goconfig.LoadConfigFile("conf/app.conf")
	if err!=nil{
		fmt.Println("加载配置文件失败"+err.Error())
	}
	Appname,_=Config.GetValue("","appname")
	Runmode,_=Config.GetValue("","runmode")
	Httpport=Config.MustInt(Runmode,"httpport",8080)
}
`