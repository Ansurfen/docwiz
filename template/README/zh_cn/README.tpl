<h1 align="center">欢迎来到 {{.ProjectName | default "<!-- projectName -->" | unescape}} 👋</h1>
<center>

{{.ProjectStack | default "<!-- projectStack -->" | unescape}}

</center>

---

<center>

{{.ProjectStatistics | default "<!-- projectStatistics -->" | unescape}}

</center>

[English]() | 简体中文

> {{.ProjectDescription | default "<!-- projectDescription -->" | unescape}}
{{ range $index, $section := .Sections}}
## {{$section.Title}}
{{$section.Description | unescape}}
{{end}}
## 🤝 贡献

欢迎提出贡献、问题和功能请求。<br />
如果你想参与贡献，请查看 [issues 页面](https://github.com/{{.ProjectOwner | default "<!-- projectOwner -->"| unescape}}/{{.ProjectName | default "<!-- projectOwner -->" | unescape}}/issues)。<br />
[查看贡献指南](./CONTRIBUTING.md)。<br />

## 📝 许可证

此软件采用 {{.License | default "<!-- license -->" | unescape}} 许可证，更多信息请参阅 [LICENSE](./LICENSE)。
