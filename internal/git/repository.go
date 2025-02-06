// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package git

import (
	"fmt"
	"io"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type RepoKind uint

const (
	RepoGitHub RepoKind = iota
	RepoGitLab
	RepoGitOthers
	RepoSVN
)

type Repository struct {
	repo     *git.Repository
	kind     RepoKind
	owner    string
	name     string
	url      string
	repoPath string
}

func New(repoPath string) (*Repository, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, err
	}

	r := &Repository{repo: repo, repoPath: repoPath}

	remotes, err := repo.Remotes()
	if err != nil {
		return nil, err
	}

	for _, remote := range remotes {
		if remote.Config().Name == "origin" {
			remoteURL := remote.Config().URLs[0]

			// git@github.com:username/repository.git
			if strings.HasPrefix(remoteURL, "git@") {
				remoteURL = strings.Replace(remoteURL, ":", "/", 1)
				remoteURL = strings.Replace(remoteURL, "git@", "https://", 1)
			}

			parsedURL, err := url.Parse(remoteURL)
			if err != nil {
				return nil, err
			}
			switch host := parsedURL.Hostname(); host {
			case "github":
				r.kind = RepoGitHub
			case "gitlab":
				r.kind = RepoGitLab
			default:
				if strings.Contains(host, "svn") {
					r.kind = RepoSVN
				} else {
					r.kind = RepoGitOthers
				}
			}

			r.url = fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Hostname())

			// ${r.url}/username/repository.git
			re := regexp.MustCompile(fmt.Sprintf(`^%s[:/](.*?)/(.*?)(?:\.git)?$`, regexp.QuoteMeta(r.url)))
			matches := re.FindStringSubmatch(remoteURL)

			if len(matches) > 2 {
				r.owner = matches[1]
				r.name = matches[2]
			} else {
				panic("fail to parse repository's name and owner")
			}
		}
	}

	return r, nil
}

func (r *Repository) Owner() string {
	return r.owner
}

func (r *Repository) Name() string {
	return r.name
}

func (r *Repository) GenerateContributors(w io.Writer) error {
	ref, err := r.repo.Head()
	if err != nil {
		return err
	}

	commitIter, err := r.repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return err
	}

	authorCommits := make(map[string]int)
	name2Email := make(map[string]string)

	err = commitIter.ForEach(func(commit *object.Commit) error {
		author := commit.Author.Name
		authorCommits[author]++
		name2Email[author] = commit.Author.Email
		return nil
	})
	if err != nil {
		return err
	}

	type authorStats struct {
		Author string
		Email  string
		Count  int
	}

	var stats []authorStats
	for author, count := range authorCommits {
		stats = append(stats, authorStats{Author: author, Email: name2Email[author], Count: count})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[j].Count // Sort by number of commit in descending order
	})

	fmt.Fprintln(w, "# Contributors")
	for _, s := range stats {
		fmt.Fprintf(w, "- [%s](%s/%s)\n", s.Author, r.url, extractGitUsername(s.Email))
	}

	return nil
}

// Extract the Git username from the email address
func extractGitUsername(email string) string {
	// Find the index of '@'
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		// If '@' is not found, return an empty string
		return ""
	}

	// Return the part before '@', which is the username
	return email[:atIndex]
}

func (r *Repository) GenerateChangelog(w io.Writer) error {
	tags := r.getSortedTags(r.repo)

	if len(tags) == 0 {
		ref, err := r.repo.Head()
		if err != nil {
			return err
		}

		commitIter, err := r.repo.Log(&git.LogOptions{From: ref.Hash()})
		if err != nil {
			return err
		}

		fmt.Fprintln(w, "# 📜 Changelog")
		err = commitIter.ForEach(func(commit *object.Commit) error {
			fmt.Fprintf(w, "- %s [%s]\n", r.formatCommitMessage(commit.Message), r.formatCommitHash(commit))
			return nil
		})

		if err != nil {
			return err
		}
	} else {
		// range commit log
		iter, err := r.repo.Log(&git.LogOptions{})
		if err != nil {
			return fmt.Errorf("failed to get commit log: %w", err)
		}
		defer iter.Close()

		tagIndex := 0
		tagMap := make(map[string][]*object.Commit)

		err = iter.ForEach(func(commit *object.Commit) error {
			// when a tag commit is encountered, switch to the next tag
			if tagIndex < len(tags) && commit.Hash == tags[tagIndex].Commit.Hash {
				tagIndex++
			}
			if tagIndex == 0 {
				// haven't entered any tags yet
				return nil
			}

			// categorized into the current tag
			tagMap[tags[tagIndex-1].Name] = append(tagMap[tags[tagIndex-1].Name], commit)
			return nil
		})

		if err != nil {
			panic(err)
		}

		fmt.Fprintln(w, "# 📜 Changelog")
		for _, tag := range tags {
			fmt.Fprintf(w, "\n## %s\n", tag.Name)
			for _, commit := range tagMap[tag.Name] {
				fmt.Fprintf(w, "- %s [%s]\n", r.formatCommitMessage(commit.Message), r.formatCommitHash(commit))
			}
		}
	}

	return nil
}

type TagInfo struct {
	Name   string
	Commit *object.Commit
}

func (r *Repository) GetTags() []TagInfo {
	return r.getSortedTags(r.repo)
}

// getSortedTags returns all tags and sort (by commit time)
func (r *Repository) getSortedTags(repo *git.Repository) []TagInfo {
	var tags []TagInfo

	iter, err := repo.Tags()
	if err != nil {
		panic(err)
	}

	err = iter.ForEach(func(ref *plumbing.Reference) error {
		tagObj, err := repo.TagObject(ref.Hash())
		var commit *object.Commit

		if err == nil {
			// Annotated tag，point at commit
			commit, err = repo.CommitObject(tagObj.Target)
		} else {
			// Lightweight tag，directly point at commit
			commit, err = repo.CommitObject(ref.Hash())
		}

		if err == nil {
			tags = append(tags, TagInfo{Name: ref.Name().Short(), Commit: commit})
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	// Sort by commit time (latest first)
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Commit.Committer.When.After(tags[j].Commit.Committer.When)
	})

	return tags
}

func (r *Repository) formatCommitHash(commit *object.Commit) string {
	return fmt.Sprintf("[%s](%s/%s/%s/commit/%s)",
		commit.Hash.String()[:7],
		r.url,
		r.owner,
		r.name,
		commit.Hash.String(),
	)
}

var (
	// hashRegex matches issue or pull request references in the form of "#23" or "GH-23".
	// Example: "#23" or "GH-23"
	hashRegex = regexp.MustCompile(`(?i)(?:GH-)?#(\d+)`)

	// mentionRegex matches user mentions in the form of "@username".
	// Example: "@john_doe" or "@alice123"
	mentionRegex = regexp.MustCompile(`@([a-zA-Z0-9_-]+)`)

	// mergeRequestRegex matches merge request references in the form of "!23".
	// Example: "!23"
	mergeRequestRegex = regexp.MustCompile(`!(\d+)`)

	// svnRegex matches Subversion (SVN) revision references in the form of "r1234".
	// Example: "r1234"
	svnRegex = regexp.MustCompile(`r(\d+)`)
)

// formatCommitMessage parse the repository's reference in the commit message
func (r *Repository) formatCommitMessage(message string) string {
	// parse Issue/PR（#23 / GH-23）
	message = hashRegex.ReplaceAllString(message, fmt.Sprintf("[#$1](%s/%s/%s/issues/$1)", r.url, r.owner, r.name))

	// parse @username
	message = mentionRegex.ReplaceAllString(message, fmt.Sprintf("[@$1](%s/$1)", r.url))

	if r.kind == RepoGitLab {
		// parse !23 format（GitLab Merge Request）
		message = mergeRequestRegex.ReplaceAllString(message, fmt.Sprintf("[GitLab Merge Request #$1](%s/%s/%s/merge_requests/$1)", r.url, r.owner, r.name))
	}

	if r.kind == RepoSVN {
		// parse r1234 format（Subversion）
		message = svnRegex.ReplaceAllString(message, fmt.Sprintf("[Subversion Revision r$1](%s/%s/%s/r$1)", r.url, r.owner, r.name))
	}

	return trimRightSpaceAndNewline(message)
}

func trimRightSpaceAndNewline(s string) string {
	return strings.TrimRightFunc(s, func(r rune) bool {
		return r == '\n' || unicode.IsSpace(r)
	})
}
