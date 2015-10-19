package main

var controllers = `package controllers

import (
	"github.com/Unknwon/macaron"
)

func DefaultRouter(ctx *macaron.Context) {
	ctx.HTML(200,"index")
}
`