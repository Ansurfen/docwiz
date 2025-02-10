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
