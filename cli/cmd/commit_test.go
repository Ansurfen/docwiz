// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import "testing"

func TestCommit(t *testing.T) {
	emojiMap = map[string]string{
		"fix":  "🔧",
		"feat": "✨",
		"docs": "📄",
	}

	tests := []struct {
		message  string
		expected string
	}{
		{"fix bug in login", "🔧 fix bug in login"},
		{"feat: add dark mode", "✨ feat: add dark mode"},
		{"docs: update README", "📄 docs: update README"},
		{"fix and feat together", "🔧 ✨ fix and feat together"},
		{"no matching keyword", "no matching keyword"},
	}

	for _, tt := range tests {
		got := addGitEmoji(tt.message)
		if got != tt.expected {
			t.Errorf("addGitEmoji(%q) = %q; want %q", tt.message, got, tt.expected)
		}
	}
}
