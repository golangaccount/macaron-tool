package txt

var Views_404_html=`<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <title>404</title>
</head>
<body>

</body>
</html>
`

var Views_forgetpassword_html=`<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <title></title>
</head>
<body>

</body>
</html>
`

var Views_index_html=`{% extends "template/main.html" %}

{% block PageHead %}
<link href="../static/css/bootstrap.css" rel="stylesheet" />
<link href="./css/bootstrap.css" rel="stylesheet" />
<script src="../static/js/jquery-2.1.4.js"></script>
<script src="/js/jquery-2.1.4.js"></script>
<style>
    * {
        margin: 0;
        padding: 0;
    }

    html {
        width: 100%;
        height: 100%;
    }

    body {
        width: 100%;
        height: 100%;
        overflow: hidden;
    }

    .head {
        height: 60px;
        background-color: rgb(0,125,125);
    }

    .footer {
        height: 60px;
        background-color: rgb(200,200,200);
    }

    .body {
        height: calc(100% - 120px);
        width: 100%;
        background-color: rgb(125,125,0);
        display: table;
    }

    .errormsg {
        color: red;
    }
</style>
<script>
    function login() {
        var checkbox = document.getElementById("remember_me");
        var account = document.getElementById("account");
        var password = document.getElementById("password");
        //进行账号的验证验证不过返回false

        if (checkbox.checked) {
            //将账号密码进行存储
            localStorage.setItem("accountinfo", JSON.stringify({ Account: account.value, Password: password.value }));
        }
        return false;
    }
    
    $(document).ready(function(){
        if (localStorage.getItem("accountinfo")) {
            var info = JSON.parse(localStorage.getItem("accountinfo"));
            console.log("info");
            document.getElementById("remember_me").checked = "true";
            document.getElementById("account").value = info.Account;
            document.getElementById("password").value = info.Password;
        }
    });

</script>
{% endblock %}


{% block Header %}
{% include "template/subblock/defaultheader.html" %}
{% endblock %}

{% block Body %}
<div class="body">
    <div style="display: table-cell; text-align: center; vertical-align: middle;">
        <div style="-moz-border-radius: 5px; -webkit-border-radius: 5px; border-radius: 5px; padding-top: 10px; padding-bottom: 20px; padding-left: 20px; padding-right: 20px; background-color: rgba(150,150,150,0.7); display: block; width: 300px; margin: 0 auto;">
            <h4 style="text-align: left">账户登入</h4>
            {% if RequestMehod == "POST" %}
                <div class="errormsg">您输入的账号或密码错误</div>
            {% endif %}
                <form method="post" onsubmit="return login()">
                    <div class="form-group">
                        <div class="input-group">
                            <span class="input-group-addon"><span class="glyphicon glyphicon-user"></span></span>
                            <input id="account" name="account" type="text" class="form-control" placeholder="请输入账号">
                        </div>
                    </div>

                    <div class="form-group">
                        <div class="input-group">
                            <span class="input-group-addon"><span class="glyphicon glyphicon-lock"></span></span>
                            <input id="password" name="password" type="password" class="form-control" placeholder="请输入密码">
                        </div>
                    </div>

                    <div class="form-group">
                        <span style="float: left; position: relative; height: 27px; line-height: 27px;">
                            <input id="remember_me" type="checkbox" style="padding-top: 3px; position: relative; top: 2px;">
                            <label for="remember_me">记住密码</label>
                        </span>
                        <span style="float: right; height: 27px; line-height: 27px; font-weight: bold">
                            <a href="/forgetpassword">忘记密码</a>
                        </span>
                        <div style="clear: both;"></div>

                    </div>
                    <button class="btn btn-primary btn-block" type="submit">登录</button>
                </form>
        </div>

    </div>
</div>
{% endblock %}

{% block Footer %}
{% include "template/subblock/defaultfooter.html"%}
{% endblock %}`

var Views_regist_html=`<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title></title>
    <link href="../static/css/bootstrap.css" rel="stylesheet" />
    <link href="./css/bootstrap.css" rel="stylesheet" />
    <script src="../static/js/jquery-2.1.4.js"></script>
    <script src="/js/jquery-2.1.4.js"></script>

    <script>
        $(document).ready(function () {
            $("#useragreementlink").click(function () {
                setTimeout(function(){
                    //submit button enable

                    //设置用户协议已阅读的标记

                }, 5000);
            });
            
            $("#submit").click(function(){
                //数据提交验证
                
                //用户协议阅读验证
            })
        })
    </script>
</head>
<body>
    <div class="container">
        <div class="form-group">
            <div class="input-group">
                <span class="input-group-addon"><span class="glyphicon glyphicon-user"></span></span>
                <input id="account" name="account" type="text" class="form-control" placeholder="请输入账号" />
            </div>
        </div>
        <div>
            <input type="checkbox" id="useragreement" /><label for="useragreement">我已阅读并同意<a id="useragreementlink" href="" target="_blank">用户协议</a></label>
        </div>
        <div><button id="submit" class="btn btn-primary btn-block" disabled="disabled">添加</button></div>
    </div>
</body>
</html>
`

var Views_registsucced_html=`<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <title></title>
</head>
<body>

</body>
</html>
`