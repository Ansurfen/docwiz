// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type QTWalker struct {
	BaseWalker
}

func (*QTWalker) SubscribeExt() []string {
	return []string{".qrc"}
}

func (*QTWalker) ParseExt(fullpath string, file string, ctx *Context) error {
	ctx.Set("QT", upgradeBadge("QT", badge.ShieldQt))
	return nil
}
