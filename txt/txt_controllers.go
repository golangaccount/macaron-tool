package txt

var Controllers_default_go=`package controllers

import (
	"github.com/Unknwon/macaron"
)

func DefaultRouter(ctx *macaron.Context) {
	ctx.HTML(200,"index")
}
`


var Controllers_page404_go=`package controllers

import(
	"github.com/Unknwon/macaron"
)


func Page404(ctx *macaron.Context){
	ctx.HTML(200,"404")
}`

var Controllers_password_go=`package controllers

import(
	"github.com/Unknwon/macaron"
)
//忘记密码
func ForgetPassword(ctx *macaron.Context){
	ctx.HTML(200,"forgetpassword")
}
//修改密码
func ChangePassword(ctx *macaron.Context){
	
}`

var Controllers_regist_go=`package controllers

import(
	"github.com/Unknwon/macaron"
)

func Regist(ctx *macaron.Context){
	ctx.HTML(200,"regist")
}

func RegistResult(ctx *macaron.Context){
	
}`
