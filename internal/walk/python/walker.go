// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package pythonwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
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

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".py", ".pyi", ".pyc", ".pyo"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	switch ext {
	case ".py":
		data, err := os.ReadFile(fullpath)
		if err != nil {
			return err
		}
		script := string(data)
		if hasDjango.MatchString(script) {
			// DjangoREST
			ctx.Set("Django", walk.UpgradeBadge("Python", shieldDjango))
		} else if hasFlask.MatchString(script) {
			ctx.Set("Flask", walk.UpgradeBadge("Python", shieldFlask))
		} else if hasFastAPI.MatchString(script) {
			ctx.Set("FastAPI", walk.UpgradeBadge("Python", shieldFastAPI))
		} else if hasAioHttp.MatchString(script) {
			ctx.Set("AioHTTP", walk.UpgradeBadge("Python", badge.ShieldAiohttp))
		} else if hasJinja2.MatchString(script) {
			ctx.Set("Jinja", walk.UpgradeBadge("Python", badge.ShieldJinja))
		}
		fallthrough
	default:
		ctx.Set("Python", walk.UpgradeBadge("Python", badge.ShieldPython))
	}
	return nil
}
