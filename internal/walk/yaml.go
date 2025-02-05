// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type YAMLWalker struct {
	BaseWalker
}

func (*YAMLWalker) SubscribeExt() []string {
	return []string{".yaml", ".yml"}
}

func (*YAMLWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("YAML", upgradeBadge("YAML", badge.ShieldYAML))
	return nil
}
