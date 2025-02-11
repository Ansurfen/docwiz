// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package dartwalk_test

import (
	"docwiz/internal/git"
	"docwiz/internal/walk"
	dartwalk "docwiz/internal/walk/dart"
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	ctx := &walk.Context{
		Ignore: &git.GitIgnore{},
		Walkers: []walk.Walker{
			&dartwalk.Walker{},
		},
	}
	walk.Walk(".", ctx)
	fmt.Println(ctx)
}
