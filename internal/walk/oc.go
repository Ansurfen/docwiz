// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type ObjectiveCWalker struct {
	BaseWalker
}

func (*ObjectiveCWalker) SubscribeExt() []string {
	return []string{".m"}
}

func (*ObjectiveCWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Objective-C", upgradeBadge("Objective-C", badge.ShieldObjectiveC))
	return nil
}
