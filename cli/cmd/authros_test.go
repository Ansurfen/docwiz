// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	testCases := []struct {
		input    string
		expected map[string]string
	}{
		{
			input: `name=Alice,duty="Lead Developer",email="123456@gmail.com"`,
			expected: map[string]string{
				"name":  "Alice",
				"duty":  "Lead Developer",
				"email": "123456@gmail.com",
			},
		},
		{
			input: `name=Bob, duty="Contributor", homepage="git@github.com/bob.git"`,
			expected: map[string]string{
				"name":     "Bob",
				"duty":     "Contributor",
				"homepage": "git@github.com/bob.git",
			},
		},
	}

	for _, tc := range testCases {
		result, err := unmarshal(tc.input)
		assert.Equal(t, nil, err)
		assert.Equal(t, tc.expected, result, "Failed on input: %s", tc.input)
	}
}
