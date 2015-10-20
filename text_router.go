package main

var router = `//macaron路由执行图
//
//1.由go系统调用 macaron的ServeHTTP函数
//2.ServeHTTP函数内部先执行向macaron注册的beforeroute函数
//	函数参数为 rw http.ResponseWriter, req *http.Request
//	当注册的函数返回true值时，终止后续的beforeroute注册函数的执行
//3.路由的执行，先找到路由和路由对应的处理函数，注册的所有的处理函数都通过一个闭包函数进行封装
//
//关于路由的几个问题
//1.所有的注册函数均是先注册先执行
//2.使用macaron的use方法注册的函数会包含到所有的注册路由的处理函数队列中，
//	use注册的函数的执行顺序要大于使用get、post等方法注册的函数
//3.关于路由注册时回调函数的问题：
//	所有的注册的函数会使用一个闭包函数进行包裹
//	对函数进行回调时，使用类型映射的方式获取回调函数的参数集合在通过reflect方式进行调用
//	系统在macaron.context生成后就将ctx对象映射到ctx的参数映射队列中，因此在路由中执行回调函数时，macaron.context对象一定存在，因此一定可以获取到ctx
//  当我们在某一个处理函数中对封装了新的对象并通过ctx.map的方式添加则在后面的处理函数中可以将改类型作为参数取到该值,不要重复映射某一类型，否则值会被覆盖
//  ex:
//		m.map(func(ctx){
//			var obj=new objectA
//			ctx.map(obj)
// 		})
//		m.any("/xxx",func(objectA){ })
//	group函数的回调函数会添加到后续的所有的函数中，并且优于func函数注册回调函数
//
//
//关于路由执行队列的几个问题
//1.使用macaron.use的方法注册的处理函数会添加到每个路由的处理函数中 执行优先级处于第一队列
//2.使用macaron.group("/xxx",func(){macaron.any("/xxx",func(){})},handlers)注册的路由group的handler会添加到所有func注册的路由的处理函数中 执行优先级处于第二队列
//3.使用macaron.any注册的路由的处理函数 执行优先级处于第三队列
//

package routers

import (
	"macaron-tool/testnewcmd/controllers"
	"net/http"

	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/pongo2"
)

func Regist(m *macaron.Macaron) {
	RegistBeforeRouter(m)
	RegistGlobalRouter(m)
	RegistGroupRouter(m)
	RegistRouter(m)
	RegistAfterRouter(m)
}

//注册全局的路由处理函数
func RegistGlobalRouter(m *macaron.Macaron) {
	//模块的添加
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())

	//静态资源的添加
	m.Use(macaron.Static("static"))

	//处理模板
	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory:  "views",
		IndentJSON: macaron.Env != macaron.PROD,
		IndentXML:  macaron.Env != macaron.PROD,
	}))

	//其它的处理函数
}

//注册分组路由
func RegistGroupRouter(m *macaron.Macaron) {
	// exp：
	//	m.Group("xxx",func(){
	//		m.Any("/xxx",func(){})
	//	},handlers)
}

//注册单独的路由
func RegistRouter(m *macaron.Macaron) {
	m.Any("/", controllers.DefaultRouter)
}

func RegistBeforeRouter(m *macaron.Macaron) {
	m.Before(func(rw http.ResponseWriter, req *http.Request)bool {
		//在此处添加路由处理之前的逻辑
		return false
	})	
}

func RegistAfterRouter(m *macaron.Macaron){
	m.Use(func(ctx *macaron.Context){
		ctx.Next()
		
		//在此处添加路由完成后的处理逻辑
		//todo
	})
}
`