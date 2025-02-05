// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type DockerWalker struct {
	BaseWalker
}

func (*DockerWalker) SubscribeFile() []string {
	return []string{"dockerfile", "Dockerfile"}
}

func (*DockerWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	ctx.Set("Docker", upgradeBadge("Docker", badge.ShieldDocker))
	return nil
}
