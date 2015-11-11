package main

import (
	"fmt"
	"macaron-tool/txt"
	"os"
	path "path/filepath"
	"strings"
)

var cmdNew = &Command{
	UsageLine: "new [appname]",
	Short:     "create an application base on macaron framework",
	Long: `
create an application base on macaron framework,

which in the current path with folder named [appname].

The [appname] folder has following structure:

    |- [appname].go
    |- conf
        |-  app.conf
		|-  conf.go
    |- controllers
         |- default.go
		 |- page404.go
		 |- password.go
		 |- regist.go
	|- dal
		 |- createengine.go
    |- models
		 |- user.go
    |- routers
         |- router.go	
    |- tests
         |- default_test.go
	|- static
         |- js
		 	 |- jquery.js
			 |- bootstrap.js
         |- css
			 |- bootstrap.css
         |- img             
    |- views
        |- template
			|- subblock
				|- defaultfooter.html
				|- defaultheader.html
			|- main.html
		|- 404.html
		|- forgetpassword.html
		|- index.html
		|- regist.html
		|- registsucced.html

`,
}

func init() {
	cmdNew.Run = createApp
}

func createApp(cmd *Command, args []string) {
	curpath, _ := os.Getwd()
	if len(args) != 1 {
		ColorLog("[ERRO] Argument [appname] is missing\n")
		os.Exit(2)
	}

	gopath := os.Getenv("GOPATH")
	Debugf("gopath:%s", gopath)
	if gopath == "" {
		ColorLog("[ERRO] $GOPATH not found\n")
		ColorLog("[HINT] Set $GOPATH in your environment vairables\n")
		os.Exit(2)
	}
	haspath := false
	appsrcpath := ""

	wgopath := path.SplitList(gopath)
	for _, wg := range wgopath {
		wg, _ = path.EvalSymlinks(path.Join(wg, "src"))

		if strings.HasPrefix(strings.ToLower(curpath), strings.ToLower(wg)) {
			haspath = true
			appsrcpath = wg
			break
		}
	}

	if !haspath {
		ColorLog("[ERRO] Unable to create an application outside of $GOPATH(%s)\n", gopath)
		ColorLog("[HINT] Change your work directory by `cd ($GOPATH%ssrc)`\n", string(path.Separator))
		os.Exit(2)
	}

	apppath := path.Join(curpath, args[0])
	appname := args[0]

	if _, err := os.Stat(apppath); os.IsNotExist(err) == false {
		fmt.Printf("[ERRO] Path(%s) has alreay existed\n", apppath)
		os.Exit(2)
	}

	fmt.Println("[INFO] Creating application...")

	os.MkdirAll(apppath, 0755)
	fmt.Println(apppath + string(path.Separator))
	os.Mkdir(path.Join(apppath, "conf"), 0755)
	fmt.Println(path.Join(apppath, "conf") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "controllers"), 0755)
	fmt.Println(path.Join(apppath, "controllers") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "dal"), 0755)
	fmt.Println(path.Join(apppath, "dal") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "models"), 0755)
	fmt.Println(path.Join(apppath, "models") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "routers"), 0755)
	fmt.Println(path.Join(apppath, "routers") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static"), 0755)
	fmt.Println(path.Join(apppath, "static") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "js"), 0755)
	fmt.Println(path.Join(apppath, "static", "js") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "fonts"), 0755)
	fmt.Println(path.Join(apppath, "static", "fonts") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "css"), 0755)
	fmt.Println(path.Join(apppath, "static", "css") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "img"), 0755)
	fmt.Println(path.Join(apppath, "static", "img") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "tests"), 0755)
	fmt.Println(path.Join(apppath, "tests") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "views"), 0755)
	fmt.Println(path.Join(apppath, "views") + string(path.Separator))
	
	os.Mkdir(path.Join(apppath, "views","template"), 0755)
	fmt.Println(path.Join(apppath, "views","template") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "views","template","subblock"), 0755)
	fmt.Println(path.Join(apppath, "views","template","subblock") + string(path.Separator))
	
	
	
	
	
	

	fmt.Println(path.Join(apppath, "conf", "app.conf"))
	writetofile(path.Join(apppath, "conf", "app.conf"), strings.Replace(txt.Conf_app_conf, "{{.Appname}}", args[0], -1))

	fmt.Println(path.Join(apppath, "conf", "conf.go"))
	writetofile(path.Join(apppath, "conf", "conf.go"), txt.Conf_conf_go)

	fmt.Println(path.Join(apppath, "controllers", "default.go"))
	writetofile(path.Join(apppath, "controllers", "default.go"), txt.Controllers_default_go)

	writetofile(path.Join(apppath, "controllers", "page404.go"), txt.Controllers_page404_go)
	writetofile(path.Join(apppath, "controllers", "password.go"), txt.Controllers_password_go)
	writetofile(path.Join(apppath, "controllers", "regist.go"), txt.Controllers_regist_go)

	writetofile(path.Join(apppath, "dal", "createengine.go"), txt.Dal_createengine_go)

	writetofile(path.Join(apppath, "models", "user.go"), txt.Models_user_go)

	fmt.Println(path.Join(apppath, "routers", "router.go"))
	writetofile(path.Join(apppath, "routers", "router.go"), strings.Replace(txt.Routers_router_go, "{{.Appname}}", strings.Join(strings.Split(apppath[len(appsrcpath)+1:], string(path.Separator)), "/"), -1))

	writetofile(path.Join(apppath, "static", "css", "bootstrap.css"), txt.Static_css_bootstrap_css)
	writetofile(path.Join(apppath, "static", "js", "bootstrap.js"), txt.Static_js_bootstrap_js)
	writetofile(path.Join(apppath, "static", "js", "jquery.js"), txt.Static_js_jquery_js)

	fmt.Println(path.Join(apppath, "tests", "default_test.go"))
	writetofile(path.Join(apppath, "tests", "default_test.go"), strings.Replace(txt.Tests_default_test_go, "{{.Appname}}", strings.Join(strings.Split(apppath[len(appsrcpath)+1:], string(path.Separator)), "/"), -1))

	writetofile(path.Join(apppath, "views", "template", "subblock", "defaultheader.html"), txt.Views_template_subblock_defaultheader_html)
	writetofile(path.Join(apppath, "views", "template", "subblock", "defaultfooter.html"), txt.Views_template_subblock_defaultfooter_html)
	writetofile(path.Join(apppath, "views", "template", "main.html"), txt.Views_template_main_html)
	writetofile(path.Join(apppath, "views", "404.html"), txt.Views_404_html)
	writetofile(path.Join(apppath, "views", "forgetpassword.html"), txt.Views_forgetpassword_html)
	writetofile(path.Join(apppath, "views", "index.html"), txt.Views_index_html)
	writetofile(path.Join(apppath, "views", "regist.html"), txt.Views_regist_html)
	writetofile(path.Join(apppath, "views", "registsucced.html"), txt.Views_registsucced_html)

	fmt.Println(path.Join(apppath, appname+".go"))
	writetofile(path.Join(apppath, appname+".go"), strings.Replace(txt.Projectname, "{{.Appname}}", strings.Join(strings.Split(apppath[len(appsrcpath)+1:], string(path.Separator)), "/"), -1))

	ColorLog("[SUCC] New application successfully created!\n")
}

func writetofile(filename, content string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}
