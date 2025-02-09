// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package template

import (
	"os"
	"testing"
)

func TestInclude(t *testing.T) {
	tmpl, err := Default("./test/main.tpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, map[string]any{
		"Name": "docwiz",
	})
	if err != nil {
		panic(err)
	}
}
