package main

const CSS = `
body {
    width: 95%;
	margin: auto;
	padding-top: 20px;

	font-weight: inherit;
	font-style: inherit;
	font-size: 100%;
	font-family: inherit;
    font: normal 100% sans-serif;
}

div.row {
	overflow: auto;
	margin-top: 0.875em;
	margin-bottom: 0.875em;
}

a { text-decoration: none; }

a:visited{ color: #6600CC; }

a:link{ color: #33CC33; }

div.name {
	float: left;
	width: 160px;

	background-color: #eee;
	border: 1px solid #aaa;
	border-right: 1px solid #eee;
	padding: 0.4em;
}

div.comic {
	margin-left: 170px;
	border-top: 1px solid #aaa;
	border-left: 1px solid #aaa;
	border-right: 1px solid #aaa;
	background-color: #eee;
	padding: 0.4em;
}

div.descr {
	font-weight: bold;
	font-size: 0.75em;
	background-color: #e3e2cf;
	border-right: 1px solid #aaa;
	border-left: 1px solid #aaa;
	border-bottom: 1px solid #aaa;
	padding-top: 0.875em;
	margin-left: 170px;
	padding: 10px;
	text-align: center;
}

img.comic {
	display: block;
	margin-left: auto;
	margin-right: auto;
}
`

const MAIN_TEMPLATE = `
<!DOCTYPE html>
<html>
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type">
    <link title="Default Comics" media="screen" href="comics.css" type="text/css" rel="stylesheet">
    <title>{{.Date.Year}}_{{.Date.Month}}_{{.Date.Day}}</title>
  </head>
  <body text="#000000" link="#ff00ff" bgcolor="#ffffff">
    <div><a href="log.txt" target="_blank">Logi błędów</a></div>
    {{range .Comics}}
      {{.HTML}}
    {{end}}
    <script>
      (function(){
        for (node of document.getElementsByClassName("js-toggle-image")){
          node.addEventListener("click", function(ev){
            var imgNode = document.getElementById(this.dataset["for"]);
            if (imgNode.style.display === "none"){ imgNode.style.display = "block"; }
            else { imgNode.style.display = "none"; }
          })
        }
      })()
    </script>
  </body>
</html>
`

const SEGMENT_TEMPLATE = `
<div class="row">
  <div class="name">
    <h1>
      <a href="{{.Comic.Url}}">{{.Comic.Name}}</a>
    </h1>
  </div>
  {{ if .ErrorMsg }}
    <div class="descr">{{.ErrorMsg}}</div>
  {{ else }}
    <div class="comic">
      <img id="comic-{{ .Comic.Name }}" src="{{.ImgSrc}}" class="comic{{ if .Comic.Nsfw }} nsfw" style="display:none{{ end }}">
      {{ if .Comic.Nsfw }}
        <a class="js-toggle-image" data-for="comic-{{ .Comic.Name }}">Pokaż / Ukryj</a>
      {{ end }}
    </div>
    {{ if .Title }}
      <div class="descr">{{.Title}}</div>
    {{ end }}
  {{ end }}
</div>
`
