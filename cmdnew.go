package main

import (
	"fmt"
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
		|-  app.go
    |- controllers
         |- default.go
    |- models
    |- routers
         |- router.go	
    |- tests
         |- default_test.go
	|- static
         |- js
         |- css
         |- img             
    |- views
        index.tpl                   

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
	appname:=args[0]

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
	os.Mkdir(path.Join(apppath, "models"), 0755)
	fmt.Println(path.Join(apppath, "models") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "routers"), 0755)
	fmt.Println(path.Join(apppath, "routers") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "tests"), 0755)
	fmt.Println(path.Join(apppath, "tests") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static"), 0755)
	fmt.Println(path.Join(apppath, "static") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "js"), 0755)
	fmt.Println(path.Join(apppath, "static", "js") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "css"), 0755)
	fmt.Println(path.Join(apppath, "static", "css") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "img"), 0755)
	fmt.Println(path.Join(apppath, "static", "img") + string(path.Separator))
	fmt.Println(path.Join(apppath, "views") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "views"), 0755)
	
	fmt.Println(path.Join(apppath, "conf", "app.conf"))
	writetofile(path.Join(apppath, "conf", "app.conf"), strings.Replace(appconf, "{{.Appname}}", args[0], -1))
	
	fmt.Println(path.Join(apppath, "conf", "conf.go"))
	writetofile(path.Join(apppath, "conf", "conf.go"), appconfgo)

	fmt.Println(path.Join(apppath, "controllers", "default.go"))
	writetofile(path.Join(apppath, "controllers", "default.go"), controllers)

	fmt.Println(path.Join(apppath, "views", "index.html"))
	writetofile(path.Join(apppath, "views", "index.html"), indextpl)

	fmt.Println(path.Join(apppath, "routers", "router.go"))
	writetofile(path.Join(apppath, "routers", "router.go"), strings.Replace(router, "{{.Appname}}", strings.Join(strings.Split(apppath[len(appsrcpath)+1:], string(path.Separator)), "/"), -1))

	fmt.Println(path.Join(apppath, "tests", "default_test.go"))
	writetofile(path.Join(apppath, "tests", "default_test.go"), strings.Replace(test, "{{.Appname}}", strings.Join(strings.Split(apppath[len(appsrcpath)+1:], string(path.Separator)), "/"), -1))

	fmt.Println(path.Join(apppath, appname+ ".go"))
	writetofile(path.Join(apppath, appname+ ".go"), strings.Replace(maingo, "{{.Appname}}", strings.Join(strings.Split(apppath[len(appsrcpath)+1:], string(path.Separator)), "/"), -1))

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
