// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import (
	"docwiz/internal/badge"
	"os"
	"regexp"
)

var (
	hasFastAPI = regexp.MustCompile(`(?i)\b(import\s+FastAPI|from\s+FastAPI)`)
	hasFlask   = regexp.MustCompile(`(?i)\b(import\s+flask|from\s+flask)`)
	hasDjango  = regexp.MustCompile(`(?i)\b(import\s+django|from\s+django)`)
	hasAioHttp = regexp.MustCompile(`(?i)\b(import\s+aiohttp|from\s+aiohttp)`)
	hasJinja2  = regexp.MustCompile(`(?i)\b(import\s+jinja2|from\s+jinja2)`)
)

type PythonWalker struct {
	BaseWalker
}

func (*PythonWalker) SubscribeExt() []string {
	return []string{".py", ".pyi", ".pyc", ".pyo"}
}

func (*PythonWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	switch ext {
	case ".py":
		data, err := os.ReadFile(fullpath)
		if err != nil {
			return err
		}
		script := string(data)
		if hasDjango.MatchString(script) {
			// DjangoREST
			ctx.Set("Django", upgradeBadge("Python", shieldDjango))
		} else if hasFlask.MatchString(script) {
			ctx.Set("Flask", upgradeBadge("Python", shieldFlask))
		} else if hasFastAPI.MatchString(script) {
			ctx.Set("FastAPI", upgradeBadge("Python", shieldFastAPI))
		} else if hasAioHttp.MatchString(script) {
			ctx.Set("AioHTTP", upgradeBadge("Python", badge.ShieldAiohttp))
		} else if hasJinja2.MatchString(script) {
			ctx.Set("Jinja", upgradeBadge("Python", badge.ShieldJinja))
		}
		fallthrough
	default:
		ctx.Set("Python", upgradeBadge("Python", badge.ShieldPython))
	}
	return nil
}

var (
	shieldDjango = &badge.ShieldBadge{
		ID:        "Django",
		Label:     "django",
		Color:     "#092E20",
		Style:     badge.ShieldStyleDefault,
		Logo:      "django",
		LogoColor: "white",
		Href:      "https://www.djangoproject.com/",
	}

	shieldFlask = &badge.ShieldBadge{
		ID:        "Flask",
		Label:     "flask",
		Color:     "#000000",
		Style:     badge.ShieldStyleDefault,
		Logo:      "flask",
		LogoColor: "white",
		Href:      "https://flask.palletsprojects.com/",
	}

	shieldFastAPI = &badge.ShieldBadge{
		ID:        "FastAPI",
		Label:     "FastAPI",
		Color:     "#005571",
		Style:     badge.ShieldStyleDefault,
		Logo:      "fastapi",
		LogoColor: "white",
		Href:      "https://fastapi.tiangolo.com/",
	}
)
