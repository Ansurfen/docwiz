<h1 align="center">Welcome to {{.Name}} 👋</h1>
<p>
{{if notEmpty .License}}
![Software License](https://img.shields.io/badge/license-{{.License}}-brightgreen.svg?style=flat-square)
{{end}}
</p>

{{ include "./default.body.tpl" . }}

{{ include "./default.foot.tpl" . }}
