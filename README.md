<h1 align="center">Welcome to docwiz 👋</h1>
<center>

[![BashScript](https://img.shields.io/badge/Bash%20Script-%23121011.svg?logo=gnu-bash&logoColor=white&style=for-the-badge)](https://www.gnu.org/software/bash/) [![Go](https://img.shields.io/badge/Go-1.23-%2300ADD8.svg?logo=go&logoColor=white&style=for-the-badge)](https://golang.org/) [![PowerShell](https://img.shields.io/badge/PowerShell-%235391FE.svg?logo=powershell&logoColor=white&style=for-the-badge)](https://learn.microsoft.com/en-us/powershell/)

</center>

---

<center>

<!-- statistics -->

</center>

English | [简体中文](docs/zh_cn/README.md)

> DocWiz is a versatile command-line tool that helps generate various types of project documentation like README, LICENSE, ROADMAP, CONTRIBUTORS, and more. It leverages templates and user inputs to create customized and professional documentation files.

## 📦 Install

### build by yourself
```bash
git clone git@github.com:Ansurfen/docwiz.git
cd ./docwiz

# method 1: using goreleaser
goreleaser release --snapshot --clean

# method 2: using goreleaser, aliax
aliax release
```

### download
You can download the binary version [here](https://github.com/Ansurfen/docwiz/releases).

## 🚀 Usage
> [!NOTE]
> The details are viewed using `docwiz -h`

### readme
Automatic scanning technology stack generation (✨RECOMMEND)
![readme_s](./docs/assets/readme_s.gif)

Based on TUI to generate
![readme_s](./docs/assets/readme.gif)

### changelog
```cmd
docwiz changelog
```

### contributor
```cmd
docwiz contributors
```

### gitignore
![gitignore](./docs/assets/gitignore.gif)

### license
![license](./docs/assets/license.gif)

### commit
![Commit](./docs/assets/commit.gif)

### copyright
![copyright](./docs/assets/copyright.gif)

### roadmap
```cmd
docwiz roadmap
```

## 🤝 Contributing

Contributions, issues and feature requests are welcome.<br />
Feel free to check [issues page](https://github.com/Ansurfen/docwiz/issues) if you want to contribute.<br />
[Check the contributing guide](./CONTRIBUTING.md).<br />

## 📝 License

This software is licensed under the MIT license, see [LICENSE](./LICENSE) for more information.

---

_This Markdown was generated with ❤️ by [docwiz](https://github.com/ansurfen/docwiz)_