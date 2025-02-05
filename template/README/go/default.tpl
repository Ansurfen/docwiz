<h1 align="center">Welcome to {{.Name}} 👋</h1>
<p>

{{if .GoPkg -}}
  {{- if notEmpty .HomePage -}}
    {{- $url := parseGitURL .HomePage -}}
    {{- $repo := cat $url.Owner "/" $url.Name -}}
[![Go Report Card](https://goreportcard.com/badge/{{- $repo -}})](https://goreportcard.com/report/{{- $repo -}})
[![GoDoc](https://godoc.org/{{- $repo -}}?status.svg)](https://pkg.go.dev/{{- $repo -}})
  {{- end }}
{{- end }}
{{- if notEmpty .License }}
![Software License](https://img.shields.io/badge/license-{{.License}}-brightgreen.svg?style=flat-square)
{{- end }}

</p>

{{ include "../default.body.tpl" . }}

{{ include "../default.foot.tpl" . }}