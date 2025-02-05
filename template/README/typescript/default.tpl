<h1 align="center">Welcome to {{.Name}} 👋</h1>
<p>

{{if .IsLib -}}
  {{- if notEmpty .HomePage -}}
    {{- $url := parseGitURL .HomePage -}}
    {{- $repo := cat $url.Owner "/" $url.Name -}}
    {{- $projectName := $url.Name | lower -}}
[![npm version](https://badge.fury.io/js/{{$projectName}}.svg)](https://www.npmjs.com/package/{{$projectName}})
[![Downloads](https://img.shields.io/npm/dm/{{$projectName}}.svg?color=blue)](https://img.shields.io/npm/dm/{{$projectName}}.svg?color=blue)
  {{- end }}
{{- end }}
{{- if notEmpty .License }}
![Software License](https://img.shields.io/badge/license-{{.License}}-brightgreen.svg?style=flat-square)
{{- end }}

</p>

{{ include "../default.body.tpl" . }}

{{ include "../default.foot.tpl" . }}