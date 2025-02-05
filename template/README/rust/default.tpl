<h1 align="center">Welcome to {{.Name}} 👋</h1>
<p>

{{if .IsLib -}}
  {{- if notEmpty .HomePage -}}
    {{- $url := parseGitURL .HomePage -}}
    {{- $repo := cat $url.Owner "/" $url.Name -}}
    {{- $projectName := $url.Name | lower -}}
[![Crates.io](https://img.shields.io/crates/v/{{$projectName}}.svg)](https://crates.io/crates/{{$projectName}})
[![Crates.io](https://img.shields.io/crates/d/{{$projectName}})](https://crates.io/crates/{{$projectName}})
[![Docs.rs](https://docs.rs/{{$projectName}}/badge.svg)](https://docs.rs/{{$projectName}})
  {{- end }}
{{- end }}
{{- if notEmpty .License }}
![Software License](https://img.shields.io/badge/license-{{.License}}-brightgreen.svg?style=flat-square)
{{- end }}

</p>

{{ include "../default.body.tpl" . }}

{{ include "../default.foot.tpl" . }}