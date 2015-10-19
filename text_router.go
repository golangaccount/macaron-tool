package main

var router = `package routers

import (
	"{{.Appname}}/controllers"
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/pongo2"
)

func Regist(m *macaron.Macaron) {
	RegistStatic(m)
	RegistRouter(m)
}

func RegistRouter(m *macaron.Macaron) {
	m.Any("/", controllers.DefaultRouter)
}

func RegistStatic(m *macaron.Macaron) {
	m.Use(macaron.Static("static"))
	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory:  "views",
		IndentJSON: macaron.Env != macaron.PROD,
		IndentXML:  macaron.Env != macaron.PROD,
	}))
}
`