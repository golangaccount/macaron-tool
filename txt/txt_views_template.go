package txt

var Views_template_main_html=`<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <title>{{ Title }}</title>
    {% block PageHead %}{% endblock %}
</head>
<body>
    {% block Header %}{% endblock %}
    {% block Body %}{% endblock %}
    {% block Footer %}{% endblock %}
</body>
</html>
`