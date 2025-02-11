package cfg

import (
	"bufio"
	"docwiz/internal/git"
	"os"
	"strings"
)

type DocWizIgnore struct {
	Git    *git.GitIgnore
	Badges map[string]struct{}
}

func LoadDocWizIgnore(filepath string) (*DocWizIgnore, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return &DocWizIgnore{Git: &git.GitIgnore{}, Badges: map[string]struct{}{}}, err
	}
	defer file.Close()

	ignoreLines := []string{}
	badges := map[string]struct{}{}
	status := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "# @docwiz-badge") {
			status = 1
			continue
		}
		switch status {
		case 0:
			ignoreLines = append(ignoreLines, line)
		case 1:
			badges[strings.TrimSpace(line)] = struct{}{}
		}
	}

	ignore := git.CompileIgnoreLines(ignoreLines...)

	return &DocWizIgnore{
		Git:    ignore,
		Badges: badges,
	}, nil
}
