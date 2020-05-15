<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1,maximum-scale=1,user-scalable=no">
    <title>文档资料_微嵌</title>
    <meta name="description" content="公司主营业务软件无线电,认知无线电,通用软件无线电平台,USRP 设计、销售。" />
    <meta name="Keywords" content="USRP,GNURadio,软件无线电,SDR,Openbts,射频子板,收发子板">
    <meta name="renderer" content="webkit">
    <link rel="stylesheet" href="/static/css/markdown.css">
    <script>
    var _hmt = _hmt || [];
    (function() {
        var hm = document.createElement("script");
        hm.src = "https://hm.baidu.com/hm.js?0f0111c99240380ee020030f3be990f5";
        var s = document.getElementsByTagName("script")[0];
        s.parentNode.insertBefore(hm, s);
    })();
    </script>
</head>
<body>
<h1 style="font-weight:600;margin-bottom:0px;border:0px;">因上努力，果上随缘</h1>

<div id="left-sider">
    {{range .topics_l}}
    <ul>{{range .Topics}}
        <li>[{{.Time.Format "06-01-02"}}] <a href="{{$.domain}}/{{.TopicID}}.html">{{.Title}}</a></li>{{end}}
    </ul>
    {{end}}
</div>

<div id="right-sider">
    {{range .topics_r}}
    <ul>{{range .Topics}}
        <li>[{{.Time.Format "06-01-02"}}] <a href="{{$.domain}}/{{.TopicID}}.html">{{.Title}}</a></li>{{end}}
    </ul>
    {{end}}
</div>

<div id="footer">
    <ul>
        <li>
        @2013-{{.time.Format "2006"}} 老彭的博客&nbsp;[Hosted by <a href="https://pages.coding.me" style="font-weight: bold" target="_blank">Coding Pages</a>] <b>Github地址</b>：http://github.com/pengbotao/itopic.go</li>
    </ul>
</div>
</body>
</html>
