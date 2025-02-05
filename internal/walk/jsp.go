// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type JSPWalker struct {
	BaseWalker
}

func (*JSPWalker) SubscribeExt() []string {
	return []string{".jsp", ".jspx"}
}

func (*JSPWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("JSP", upgradeBadge("JSP", &badge.ShieldBadge{
		ID:        "JSP",
		Label:     "jsp",
		Color:     "#FF0000",
		Style:     badge.ShieldStyleDefault,
		LogoColor: "white",
		Href:      "https://www.oracle.com/java/technologies/jspt.html",
	}))
	return nil
}
