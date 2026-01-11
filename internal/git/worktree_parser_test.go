package git

import (
	"bufio"
	"bytes"
	"log"
	"strings"
	"testing"
	"wt/internal/domain"
)

func TestParsing(t *testing.T) {
	var sut = `worktree /Users/antonio/Documents/Development/XCode/Packages/APUtils
HEAD ca8577a1e8c2027bcf78823f04460628dd3b634e
branch refs/heads/develop

worktree /Users/antonio/Documents/Development/Worktrees/APUtils/develop
HEAD f704f185cba11dc9b6d47c06000b2fffa6f3b100
branch refs/heads/master

worktree /Users/antonio/Documents/Development/Worktrees/APUtils/locked
HEAD 4e04b2b0961c494fb643d91c8956813dbfcc799d
branch refs/heads/feature/locked
`
	sc := bufio.NewScanner(bytes.NewReader([]byte(sut)))
	var (
		worktrees []domain.Worktree
		current   *domain.Worktree
	)
	flush := func() {
		if current == nil {
			return
		}
		worktrees = append(worktrees, *current)
		current = nil
	}
	for sc.Scan() {
		line := sc.Text()
		log.Println("LINE: " + line)
		if line == "" {
			flush()
			continue
		}
		if path, found := strings.CutPrefix(line, "worktree "); found {
			current = &domain.Worktree{Path: path}
			log.Println("Added new worktree: " + path)
		}
		if head, found := strings.CutPrefix(line, "HEAD "); found {
			current.HeadSHA = head
		}
		if branch, found := strings.CutPrefix(line, "branch "); found {
			current.Branch = strings.TrimPrefix(branch, "refs/heads/")
		}
	}
	flush()
	log.Println("Saved worktrees: ", len(worktrees))
	for _, worktree := range worktrees {
		log.Printf("Path: %v\nHEAD: %v\nBranch: %v", worktree.Path, worktree.HeadSHA, worktree.Branch)
	}
}
