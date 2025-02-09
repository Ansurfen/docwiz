# 🎉 作者与贡献者

[English]() | 简体中文

这个项目的实现离不开那些为其做出贡献的了不起的人们。以下是帮助使这个项目变得更加出色的所有人名单！

---

## 🛠️ 维护者

维护者负责项目的整体健康，包括定期发布、处理问题和管理拉取请求。

{{range .Maintainers}}
- **{{.Name}}** - 维护者，{{.Duty}}
{{- if gt (len .Profile) 0}}
_{{- .Profile -}}_
{{- end}}
{{- range $key, $value := .Others}}
  - {{$key}}: {{$value}}
{{- end}}
{{- end}}
---

## 💻 贡献者

这些是为这个项目贡献代码、功能或错误修复的了不起的人们。感谢你们的辛勤工作和奉献！

{{range .Contributors}}
- **{{.Name}}** - {{.Duty}}
{{- if gt (len .Profile) 0}}
_{{- .Profile -}}_
{{- end}}
{{- range $key, $value := .Others}}
  - {{$key}}: {{$value}}
{{- end}}
{{- end}}

---

## 🙏 特别感谢

特别感谢那些为本项目提供宝贵支持和资源的人们。你们的贡献不仅仅局限于代码！

{{range .SpecialContributors}}
- **{{.Name}}** - {{.Duty}}
{{- if gt (len .Profile) 0}}
_{{- .Profile -}}_
{{- end}}
{{- range $key, $value := .Others}}
  - {{$key}}: {{$value}}
{{- end}}
{{- end}}

{{if gt (len .License) 0}}
---

## 📜 许可证

此项目采用 [{{.License}}](LICENSE) 许可证，更多信息请参阅 [LICENSE](LICENSE)。
{{end}}
---

*如果你为这个项目做出了贡献，但没有出现在这个名单上，请提交拉取请求或联系我添加你的名字！*
