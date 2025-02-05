// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import (
	"docwiz/internal/git"
)

type GitWalker struct {
	BaseWalker
}

func (*GitWalker) SubscribeDir() []string {
	return []string{".git"}
}

func (*GitWalker) ParseDir(fullpath string, dir string, ctx *Context) error {
	repo := git.New(fullpath)

	ctx.ProjectName = repo.Name()
	ctx.ProjectOwner = repo.Owner()
	return nil
}
