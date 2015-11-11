package txt

var Projectname=`package main

import (
	"fmt"
	"net/http"
	"{{.Appname}}/conf"
	"{{.Appname}}/routers"
	"github.com/Unknwon/macaron"

)

func newInstance() *macaron.Macaron {
	m := macaron.New()
	return m
}

func main() {
	m := newInstance()
	routers.Regist(m)
	listenAddr := fmt.Sprintf("0.0.0.0:%d", conf.Httpport)
	fmt.Println(listenAddr)
	fmt.Println("start")
	http.ListenAndServe(listenAddr, m)
}
`