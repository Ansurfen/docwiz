// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type OCamlWalker struct {
	BaseWalker
}

func (*OCamlWalker) SubscribeExt() []string {
	return []string{".ml", ".mli"}
}

func (*OCamlWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("OCaml", upgradeBadge("OCaml", badge.ShieldOCaml))
	return nil
}
