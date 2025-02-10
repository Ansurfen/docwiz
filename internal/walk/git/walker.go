// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package gitwalk

import (
	"docwiz/internal/git"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeDir() []string {
	return []string{".git"}
}

func (*Walker) ParseDir(fullpath string, dir string, ctx *walk.Context) error {
	repo, err := git.New(fullpath)
	if err != nil {
		return err
	}

	ctx.ProjectName = repo.Name()
	ctx.ProjectOwner = repo.Owner()
	return nil
}
