// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk_test

import (
	"docwiz/internal/walk"
	"testing"
)

func TestWalk(t *testing.T) {
	walk.Walk(".", &walk.Context{
		Walkers: []walk.Walker{
			&walk.GoWalker{},
		},
	})
}
