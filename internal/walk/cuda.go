// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type CudaWalker struct {
	BaseWalker
}

func (*CudaWalker) SubscribeExt() []string {
	return []string{".cu"}
}

func (*CudaWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Cuda", upgradeBadge("Cuda", badge.ShieldCUDA))
	return nil
}
