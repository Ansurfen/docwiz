// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package template

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/url"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/Masterminds/sprig/v3"
)

func DocwizFuncMap(tplPath string) template.FuncMap {
	docwizFuncs := sprig.TxtFuncMap()

	docwizFuncs["include"] = includeTemplateFunc(tplPath)
	docwizFuncs["notEmpty"] = notEmpty
	docwizFuncs["parseGitURL"] = parseGitURL
	docwizFuncs["cat"] = cat
	docwizFuncs["unescape"] = unescape

	// emoji
	docwizFuncs["emojilizePrefix"] = emojilizePrefix
	docwizFuncs["emojilizeSuffix"] = emojilizeSuffix
	docwizFuncs["registerEmoji"] = registerEmoji

	// version
	docwizFuncs["randomVersion"] = randomVersion
	docwizFuncs["newVersion"] = newVersion
	docwizFuncs["versionInc"] = versionInc
	docwizFuncs["versionIncMajor"] = versionIncMajor
	docwizFuncs["versionIncMinor"] = versionIncMinor
	docwizFuncs["versionIncPatch"] = versionIncPatch

	// time
	docwizFuncs["dateModify"] = dateModify
	docwizFuncs["nowQuarter"] = nowQuarter
	docwizFuncs["quarterModify"] = quarterModify

	return docwizFuncs
}

func includeTemplateFunc(templateDir string) func(string, interface{}) (string, error) {
	return func(templatePath string, data interface{}) (string, error) {
		// Parse the included template
		fullPath := filepath.Join(templateDir, templatePath)

		includedTpl, err := template.New(filepath.Base(templatePath)).
			Funcs(DocwizFuncMap(filepath.Dir(fullPath))).
			ParseFiles(fullPath)
		if err != nil {
			return "", err
		}

		// Execute the included template into a string
		var sb strings.Builder
		err = includedTpl.Execute(&sb, data)
		if err != nil {
			return "", err
		}

		// Return the rendered template content as a string
		return sb.String(), nil
	}
}

func unescape(s string) template.HTML {
	return template.HTML(s)
}

func notEmpty(given any) bool {
	g := reflect.ValueOf(given)
	if !g.IsValid() {
		return false
	}

	switch g.Kind() {
	default:
		return !g.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return g.Len() != 0
	case reflect.Bool:
		return g.Bool()
	case reflect.Complex64, reflect.Complex128:
		return g.Complex() != 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return g.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return g.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return g.Float() != 0
	case reflect.Struct:
		return true
	}
}

type GitUrl struct {
	url.URL
	Owner string
	Name  string
}

func parseGitURL(remoteURL string) GitUrl {
	u := GitUrl{}

	if strings.HasPrefix(remoteURL, "git@") {
		remoteURL = strings.Replace(remoteURL, "git@", "https://", 1)
		remoteURL = strings.Replace(remoteURL, ":", "/", 1)
	}

	parsedURL, err := url.Parse(remoteURL)
	if err != nil {
		panic(err)
	}

	host := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Hostname())

	// ${r.url}/username/repository.git
	re := regexp.MustCompile(fmt.Sprintf(`^%s[:/](.*?)/(.*?)(?:\.git)?$`, regexp.QuoteMeta(host)))
	matches := re.FindStringSubmatch(remoteURL)

	if len(matches) > 2 {
		u.Owner = matches[1]
		u.Name = matches[2]
	} else {
		panic("fail to parse")
	}
	return u
}

func cat(v ...interface{}) string {
	var sb strings.Builder
	for _, val := range v {
		sb.WriteString(fmt.Sprintf("%v", val))
	}
	return sb.String()
}

func nowQuarter() string {
	now := time.Now()
	return getQuarter(now)
}

func getQuarter(t time.Time) string {
	year, month := t.Year(), int(t.Month())
	quarter := (month-1)/3 + 1
	return fmt.Sprintf("Q%d %d", quarter, year)
}

func quarterModify(offset string, baseQuarter string) (string, error) {
	// parse baseQuarter, e.g. "Q1 2025"
	parts := strings.Split(baseQuarter, " ")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid quarter format: %s", baseQuarter)
	}

	quarter, err := strconv.Atoi(strings.TrimPrefix(parts[0], "Q"))
	if err != nil {
		return "", err
	}
	year, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}

	// parse offset，e.g. "+1" or "-1"
	modifier, err := strconv.Atoi(offset)
	if err != nil {
		return "", err
	}

	// calculate the new quarter and year
	newQuarter := quarter + modifier
	for newQuarter > 4 {
		newQuarter -= 4
		year++
	}
	for newQuarter < 1 {
		newQuarter += 4
		year--
	}

	return fmt.Sprintf("Q%d %d", newQuarter, year), nil
}

func addDurationToTime(t time.Time, durationStr string) time.Time {
	var years, months, days int

	// Regular expression to match the duration format like 4m, 1y0m7d, -1y-5m-10d
	re := regexp.MustCompile(`(?:(-?\d+)y)?(?:(-?\d+)m)?(?:(-?\d+)d)?`)

	// Find matches
	matches := re.FindStringSubmatch(durationStr)
	if len(matches) > 0 {
		if matches[1] != "" {
			fmt.Sscanf(matches[1], "%d", &years)
		}
		if matches[2] != "" {
			fmt.Sscanf(matches[2], "%d", &months)
		}
		if matches[3] != "" {
			fmt.Sscanf(matches[3], "%d", &days)
		}
	}

	// Add the parsed values to the time object (negative durations are handled)
	return t.AddDate(years, months, days)
}

func dateModify(format string, date time.Time) time.Time {
	return addDurationToTime(date, format)
}

func newVersion(v string) semver.Version {
	ver, err := semver.NewVersion(v)
	if err != nil {
		return *semver.New(0, 0, 1, "", "")
	}
	return *ver
}

func randomVersion() semver.Version {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	major := r.Intn(10)
	minor := r.Intn(10)
	patch := r.Intn(10)
	return *semver.New(uint64(major), uint64(minor), uint64(patch), "", "")
}

func versionInc(modifier string, left semver.Version) semver.Version {
	right, err := semver.NewVersion(modifier)
	if err != nil {
		return left
	}

	return *semver.New(
		left.Major()+right.Major(),
		left.Minor()+right.Minor(),
		left.Patch()+right.Patch(),
		left.Prerelease(),
		left.Metadata(),
	)
}

func versionIncMajor(v semver.Version) semver.Version {
	return v.IncMajor()
}

func versionIncMinor(v semver.Version) semver.Version {
	return v.IncMinor()
}

func versionIncPatch(v semver.Version) semver.Version {
	return v.IncPatch()
}

var emojiMap = map[string]string{
	// 🎨 Improve structure / format of the code.
	"format":   "🎨",
	"reformat": "🎨",

	// ⚡ Improve performance.
	"performance": "⚡",
	"optimize":    "⚡",
	"speed":       "⚡",

	// 🔥 Remove code or files.
	"remove":  "🔥",
	"delete":  "🔥",
	"cleanup": "🔥",

	// 🐛 Fix a bug.
	"fix":   "🐛",
	"bug":   "🐛",
	"error": "🐛",

	// 🚑 Critical hotfix.
	"hotfix": "🚑",
	"patch":  "🚑",

	// ✨ Introduce new features.
	"feature": "✨",
	"new":     "✨",
	"add":     "✨",

	// 📝 Add or update documentation.
	"docs":     "📝",
	"document": "📝",
	"readme":   "📝",

	// 🚀 Deploy stuff.
	"deploy":  "🚀",
	"release": "🚀",
	"launch":  "🚀",

	// 💄 Improve UI / UX.
	"ui":     "💄",
	"ux":     "💄",
	"design": "💄",
	"theme":  "💄",

	// 🎉 Initial commit.
	"init":    "🎉",
	"initial": "🎉",
	"first":   "🎉",

	// ✅ Add tests.
	"test":     "✅",
	"unittest": "✅",

	// 🔧 Configuration changes.
	"config":   "🔧",
	"settings": "🔧",
	"env":      "🔧",

	// 🔥 Remove dead code.
	"deadcode": "🔥",
	"unused":   "🔥",

	// ♻️ Refactoring code.
	"refactor": "♻️",
	"rewrite":  "♻️",
	"rework":   "♻️",

	// 🚚 Move / rename files.
	"rename": "🚚",
	"move":   "🚚",

	// 🔒 Fix security issues.
	"security": "🔒",
	"secure":   "🔒",

	// ⚰️ Remove dead code.
	"dead":   "⚰️",
	"legacy": "⚰️",

	// 🏗️ Work in progress.
	"wip":  "🏗️",
	"work": "🏗️",

	// 🚧 Work in progress.
	"inprogress": "🚧",
	"progress":   "🚧",

	// 💡 Add comments.
	"comment": "💡",
	"note":    "💡",

	// 📦 Add or update dependencies.
	"dependency": "📦",
	"deps":       "📦",
	"update":     "📦",

	// 🔊 Add or update logs.
	"log":    "🔊",
	"logger": "🔊",
	"debug":  "🔊",

	// 🎭 Mocking data.
	"mock": "🎭",
	"fake": "🎭",

	// 🚑 Critical hotfix.
	"urgent":    "🚑",
	"emergency": "🚑",

	// 🛠️ Maintenance work.
	"maintain":    "🛠️",
	"maintenance": "🛠️",
}

func emojilizePrefix(text string) string {
	lowerText := strings.ToLower(text)
	for key, emoji := range emojiMap {
		if strings.Contains(lowerText, key) {
			return emoji + " " + text
		}
	}
	return text
}

func emojilizeSuffix(text string) string {
	lowerText := strings.ToLower(text)
	for key, emoji := range emojiMap {
		if strings.Contains(lowerText, key) {
			return text + " " + emoji
		}
	}
	return text
}

func registerEmoji(keyword, emoji string) string {
	emojiMap[keyword] = emoji
	return ""
}
