// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// 
// This project uses code from [sabhiram/go-gitignore] which is licensed under the MIT License.
// For more information, please visit: [https://github.com/sabhiram/go-gitignore]
// 
// Implement tests, ported from https://github.com/kaelzhang/node-ignore.git
package git_test

import (
	"docwiz/internal/git"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(test *testing.T) {
	lines := []string{"foo"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "foo")
	shouldMatch(test, object, "foo/")
	shouldMatch(test, object, "/foo")
	shouldNotMatch(test, object, "fooo")
	shouldNotMatch(test, object, "ofoo")
}

func TestAnywhere(test *testing.T) {
	lines := []string{"**/foo"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "foo")
	shouldMatch(test, object, "foo/")
	shouldMatch(test, object, "/foo")
	shouldNotMatch(test, object, "fooo")
	shouldNotMatch(test, object, "ofoo")
}

func TestAnywhereFromRoot(test *testing.T) {
	lines := []string{"/**/foo"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "foo")
	shouldMatch(test, object, "foo/")
	shouldMatch(test, object, "/foo")
	shouldNotMatch(test, object, "fooo")
	shouldNotMatch(test, object, "ofoo")
}

func TestSimpleDir(test *testing.T) {
	lines := []string{"foo/"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "foo/")
	shouldMatch(test, object, "foo/a")
	shouldMatch(test, object, "/foo/")
	shouldNotMatch(test, object, "foo")
	shouldNotMatch(test, object, "/foo")
}

func TestRootExtensionOnly(test *testing.T) {
	lines := []string{"/.js"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, ".js")
	shouldMatch(test, object, ".js/")
	shouldMatch(test, object, ".js/a")
	// ???
	// shouldNotMatch(test, object, "/.js")
	shouldNotMatch(test, object, ".jsa")
}

func TestRootExtension(test *testing.T) {
	lines := []string{"/*.js"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, ".js")
	shouldMatch(test, object, ".js/")
	shouldMatch(test, object, ".js/a")
	shouldMatch(test, object, "a.js/a")
	shouldMatch(test, object, "a.js/a.js")
	// ???
	// shouldNotMatch(test, object, "/.js")
	shouldNotMatch(test, object, ".jsa")
}

func TestExtension(test *testing.T) {
	lines := []string{"*.js"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, ".js")
	shouldMatch(test, object, ".js/")
	shouldMatch(test, object, ".js/a")
	shouldMatch(test, object, "a.js/a")
	shouldMatch(test, object, "a.js/a.js")
	shouldMatch(test, object, "/.js")
	shouldNotMatch(test, object, ".jsa")
}

func TestStarExtension(test *testing.T) {
	lines := []string{".js*"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, ".js")
	shouldMatch(test, object, ".js/")
	shouldMatch(test, object, ".js/a")
	shouldNotMatch(test, object, "a.js/a")
	shouldNotMatch(test, object, "a.js/a.js")
	shouldMatch(test, object, "/.js")
	shouldMatch(test, object, ".jsa")
}

func TestDoubleStar(test *testing.T) {
	lines := []string{"foo/**/"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "foo/")
	shouldMatch(test, object, "foo/abc/")
	shouldMatch(test, object, "foo/x/y/z/")
	shouldNotMatch(test, object, "foo")
	shouldNotMatch(test, object, "/foo")
}

func TestStars(test *testing.T) {
	lines := []string{"foo/**/*.bar"}
	object := git.CompileIgnoreLines(lines...)

	shouldNotMatch(test, object, "foo/")
	shouldNotMatch(test, object, "abc.bar")
	shouldMatch(test, object, "foo/abc.bar")
	shouldMatch(test, object, "foo/abc.bar/")
	shouldMatch(test, object, "foo/x/y/z.bar")
	shouldMatch(test, object, "foo/x/y/z.bar/")
}

func TestCases_Comment(test *testing.T) {
	lines := []string{"#abc"}
	object := git.CompileIgnoreLines(lines...)

	shouldNotMatch(test, object, "#abc")
}

func TestCases_EscapedComment(test *testing.T) {
	lines := []string{`\#abc`}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "#abc")
}

func TestCases_CouldFilterPaths(test *testing.T) {
	lines := []string{"abc", "!abc/b"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "abc/a.js")
	shouldNotMatch(test, object, "abc/b/b.js")
}

func TestCases_IgnoreSelect(test *testing.T) {
	lines := []string{"abc", "!abc/b", "#e", `\#f`}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "abc/a.js")
	shouldNotMatch(test, object, "abc/b/b.js")
	shouldNotMatch(test, object, "#e")
	shouldMatch(test, object, "#f")
}

func TestCases_EscapeRegexMetacharacters(test *testing.T) {
	lines := []string{"*.js", `!\*.js`, "!a#b.js", "!?.js", "#abc", `\#abc`}
	object := git.CompileIgnoreLines(lines...)

	shouldNotMatch(test, object, "*.js")
	shouldMatch(test, object, "abc.js")
	shouldNotMatch(test, object, "a#b.js")
	shouldNotMatch(test, object, "abc")
	shouldMatch(test, object, "#abc")
	shouldNotMatch(test, object, "?.js")
}

func TestCases_QuestionMark(test *testing.T) {
	lines := []string{"/.project", "thumbs.db", "*.swp", ".sonar/*", ".*.sw?"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, ".project")
	shouldNotMatch(test, object, "abc/.project")
	shouldNotMatch(test, object, ".a.sw")
	shouldMatch(test, object, ".a.sw?")
	shouldMatch(test, object, "thumbs.db")
}

func TestCases_DirEndedWithStar(test *testing.T) {
	lines := []string{"abc/*"}
	object := git.CompileIgnoreLines(lines...)

	shouldNotMatch(test, object, "abc")
}

func TestCases_FileEndedWithStar(test *testing.T) {
	lines := []string{"abc.js*"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "abc.js/")
	shouldMatch(test, object, "abc.js/abc")
	shouldMatch(test, object, "abc.jsa/")
	shouldMatch(test, object, "abc.jsa/abc")
}

func TestCases_WildcardAsFilename(test *testing.T) {
	lines := []string{"*.b"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "b/a.b")
	shouldMatch(test, object, "b/.b")
	shouldNotMatch(test, object, "b/.ba")
	shouldMatch(test, object, "b/c/a.b")
}

func TestCases_SlashAtBeginningAndComeWithWildcard(test *testing.T) {
	lines := []string{"/*.c"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, ".c")
	shouldMatch(test, object, "c.c")
	shouldNotMatch(test, object, "c/c.c")
	shouldNotMatch(test, object, "c/d")
}

func TestCases_DotFile(test *testing.T) {
	lines := []string{".d"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, ".d")
	shouldNotMatch(test, object, ".dd")
	shouldNotMatch(test, object, "d.d")
	shouldMatch(test, object, "d/.d")
	shouldNotMatch(test, object, "d/d.d")
	shouldNotMatch(test, object, "d/e")
}

func TestCases_DotDir(test *testing.T) {
	lines := []string{".e"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, ".e/")
	shouldNotMatch(test, object, ".ee/")
	shouldNotMatch(test, object, "e.e/")
	shouldMatch(test, object, ".e/e")
	shouldMatch(test, object, "e/.e")
	shouldNotMatch(test, object, "e/e.e")
	shouldNotMatch(test, object, "e/f")
}

func TestCases_PatternOnce(test *testing.T) {
	lines := []string{"node_modules/"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "node_modules/gulp/node_modules/abc.md")
	shouldMatch(test, object, "node_modules/gulp/node_modules/abc.json")
}

func TestCases_PatternTwice(test *testing.T) {
	lines := []string{"node_modules/", "node_modules/"}
	object := git.CompileIgnoreLines(lines...)

	shouldMatch(test, object, "node_modules/gulp/node_modules/abc.md")
	shouldMatch(test, object, "node_modules/gulp/node_modules/abc.json")
}

////////////////////////////////////////////////////////////

func shouldMatch(test *testing.T, object *git.GitIgnore, path string) {
	assert.Equal(test, true, object.MatchesPath(path), path+" should match")
}

func shouldNotMatch(test *testing.T, object *git.GitIgnore, path string) {
	assert.Equal(test, false, object.MatchesPath(path), path+" should not match")
}

////////////////////////////////////////////////////////////
