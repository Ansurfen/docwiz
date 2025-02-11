// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

type Walker interface {
	SubscribeExt() []string
	SubscribeFile() []string
	SubscribeDir() []string

	ParseExt(fullpath, ext string, ctx *Context) error
	ParseFile(fullpath, file string, ctx *Context) error
	ParseDir(fullpath, dir string, ctx *Context) error
}

type parseHandler func(string, string, *Context) error

type BaseWalker struct{}

func (BaseWalker) SubscribeExt() []string  { return nil }
func (BaseWalker) SubscribeFile() []string { return nil }
func (BaseWalker) SubscribeDir() []string  { return nil }

func (BaseWalker) ParseExt(fullpath, ext string, ctx *Context) error   { return nil }
func (BaseWalker) ParseFile(fullpath, file string, ctx *Context) error { return nil }
func (BaseWalker) ParseDir(fullpath, dir string, ctx *Context) error   { return nil }

type Context struct {
	Ignore  *cfg.DocWizIgnore
	Walkers []Walker

	Output   string
	Template string

	ProjectName        string
	ProjectOwner       string
	ProjectDescription string
	ProjectStack       string

	stackKind BadgeKind
	stack     map[string]badge.SortableBadge

	statisticsKind BadgeKind
	statistics     map[string]badge.SortableBadge
	Sections       []Section
}

func (c *Context) StackBadgeKind() BadgeKind {
	return c.stackKind
}

type Section struct {
	Title       string
	Description string
}

type BadgeKind int

const (
	BadgeKindShield BadgeKind = iota
	BadgeKindBadgen
)

func (c *Context) Get(name string) badge.SortableBadge {
	return c.stack[name]
}

func (c *Context) Set(name string, b badge.SortableBadge) badge.SortableBadge {
	c.stack[name] = b
	return b
}

func (c *Context) generate() {
	c.Sections = append(c.Sections,
		Section{Title: "📦 Install", Description: "<!-- description -->"},
		Section{Title: "🚀 Usage", Description: "<!-- description -->"},
		Section{Title: "✅ Test", Description: "<!-- description -->"})

	var stackPairs []struct {
		tag   string
		badge badge.Badge
	}

	for _, b := range c.stack {
		stackPairs = append(stackPairs, struct {
			tag   string
			badge badge.Badge
		}{b.Tag, b.Badge})
	}

	sort.Slice(stackPairs, func(i, j int) bool {
		return stackPairs[i].tag < stackPairs[j].tag
	})

	badgeStr := []string{}
	for _, s := range stackPairs {
		if _, ok := c.Ignore.Badges[s.badge.Name()]; !ok {
			badgeStr = append(badgeStr, s.badge.Markdown())
		}
	}

	c.ProjectStack = strings.Join(badgeStr, " ")
}

func Walk(root string, ctx *Context) error {
	ctx.stack = make(map[string]badge.SortableBadge)

	extHandlers := map[string][]parseHandler{}
	dirHandlers := map[string][]parseHandler{}
	fileHandlers := map[string][]parseHandler{}

	for _, w := range ctx.Walkers {
		for _, ext := range w.SubscribeExt() {
			extHandlers[ext] = append(extHandlers[ext], w.ParseExt)
		}
		for _, dir := range w.SubscribeDir() {
			dirHandlers[dir] = append(dirHandlers[dir], w.ParseDir)
		}
		for _, file := range w.SubscribeFile() {
			fileHandlers[file] = append(fileHandlers[file], w.ParseFile)
		}
	}
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if ctx.Ignore.Git.MatchesPath(path) {
			return nil
		}

		fullpath, _ := filepath.Abs(path)

		if info.IsDir() {
			dir := filepath.Base(path)
			if handlers, ok := dirHandlers[dir]; ok {
				for _, handler := range handlers {
					handler(fullpath, dir, ctx)
				}
			}
			return nil
		}

		file := filepath.Base(path)
		if handlers, ok := fileHandlers[file]; ok {
			for _, handler := range handlers {
				handler(fullpath, file, ctx)
			}
		}

		ext := filepath.Ext(path)
		if handlers, ok := extHandlers[ext]; ok {
			for _, handler := range handlers {
				handler(fullpath, ext, ctx)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	ctx.generate()
	return nil
}

func UpgradeBadge(tag string, b badge.Badge) badge.SortableBadge {
	return badge.SortableBadge{Tag: tag, Badge: b}
}

type ExtraInfo int

type ExtendedBadge interface {
	Kind() ExtraInfo
	Unwrap() badge.Badge
}

const (
	ExtraInfoNone = iota
	ExtraInfoUseUseDependencyVersion
)

type DependencyVersionBadge struct {
	badge.Badge
}

func (DependencyVersionBadge) Kind() ExtraInfo {
	return ExtraInfoUseUseDependencyVersion
}

func (b DependencyVersionBadge) Unwrap() badge.Badge {
	return b.Badge
}

type SystemVersionBadge struct {
	badge.Badge
}

func (SystemVersionBadge) Kind() ExtraInfo {
	return ExtraInfoNone
}

func (b SystemVersionBadge) Unwrap() badge.Badge {
	return b.Badge
}
