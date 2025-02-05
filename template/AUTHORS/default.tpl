# 🎉 Authors & Contributors

This project wouldn't have been possible without the amazing people who have contributed. Here’s a list of those who have helped make this project great!

---

## 🛠️ Maintainers

The maintainers are responsible for the overall health of the project, including regular releases, handling issues, and managing pull requests.

{{range .Maintainers}}
- **{{.Name}}** - Maintainer, {{.Duty}}
{{- if gt (len .Profile) 0}}
_{{- .Profile -}}_
{{- end}}
{{- range $key, $value := .Data}}
  - {{$key}}: {{$value}}
{{- end}}
{{- end}}
---

## 💻 Contributors

These are the awesome people who have contributed code, features, or bug fixes to this project. Thank you for your hard work and dedication!

{{range .Contributors}}
- **{{.Name}}** - {{.Duty}}
{{- if gt (len .Profile) 0}}
_{{- .Profile -}}_
{{- end}}
{{- range $key, $value := .Data}}
  - {{$key}}: {{$value}}
{{- end}}
{{- end}}

---

## 🙏 Special Thanks

A big thank you to those who provided invaluable support and resources for this project. Your contributions go beyond just code!

{{range .SpecialContributors}}
- **{{.Name}}** - {{.Duty}}
{{- if gt (len .Profile) 0}}
_{{- .Profile -}}_
{{- end}}
{{- range $key, $value := .Data}}
  - {{$key}}: {{$value}}
{{- end}}
{{- end}}

{{if gt (len .License) 0}}
---

## 📜 License

This project is licensed under the [{{.License}}](LICENSE), see [LICENSE](LICENSE) for more information.
{{end}}
---

*If you have contributed to this project and don't see your name here, please submit a pull request or contact us to add you!*

_This AUTHORS was generated with 🔥 by [docwiz](https://github.com/ansurfen/docwiz)_