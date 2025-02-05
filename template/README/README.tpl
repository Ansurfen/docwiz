<h1 align="center">Welcome to {{.ProjectName | default "<!-- projectName -->" | unescape}} 👋</h1>
<center>

{{.ProjectStack | default "<!-- projectStack -->" | unescape}}

</center>

---

<center>

{{.ProjectStatistics | default "<!-- projectStatistics -->" | unescape}}

</center>

> {{.ProjectDescription | default "<!-- projectDescription -->" | unescape}}
{{ range $index, $section := .Sections}}
## {{$section.Title}}
{{$section.Description | unescape}}
{{end}}
## 🤝 Contributing

Contributions, issues and feature requests are welcome.<br />
Feel free to check [issues page](https://github.com/{{.ProjectOwner | default "<!-- projectOwner -->"| unescape}}/{{.ProjectName | default "<!-- projectOwner -->" | unescape}}/issues) if you want to contribute.<br />
[Check the contributing guide](./CONTRIBUTING.md).<br />

## 📝 License

This software is licensed under the {{.License | default "<!-- license -->" | unescape}} license, see [LICENSE](./LICENSE) for more information.
