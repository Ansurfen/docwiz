// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type CMakeWalker struct {
	BaseWalker
}

func (*CMakeWalker) SubscribeFile() []string {
	return []string{"CMakeLists.txt"}
}

func (*CMakeWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	ctx.Set("CMake", upgradeBadge("CMake", badge.ShieldCMake))
	return nil
}
