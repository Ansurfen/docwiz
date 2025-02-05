// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type ClojureWalker struct {
	BaseWalker
}

func (*ClojureWalker) SubscribeExt() []string {
	return []string{".clj", ".cljs", ".cljc", ".cljsc", ".edn"}
}

func (*ClojureWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Clojure", upgradeBadge("Clojure", badge.ShieldClojure))
	return nil
}
