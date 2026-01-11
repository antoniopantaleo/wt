package git

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
	"strings"
	"wt/internal/domain"
)

type ExecGit struct{}

func (g ExecGit) GetWorktreesFromPath(path string) []domain.Worktree {
	cmd := exec.Command("git", "-C", path, "rev-parse", "--path-format=absolute", "--git-common-dir", "--show-toplevel")
	// TODO: handle error
	data, _ := cmd.Output()
	sc := bufio.NewScanner(bytes.NewReader(data))
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
		if worktreePath, found := strings.CutPrefix(line, "worktree "); found {
			current = &domain.Worktree{Path: worktreePath, RepoPath: path}
		}
		if head, found := strings.CutPrefix(line, "HEAD "); found {
			current.HeadSHA = head
		}
		if branch, found := strings.CutPrefix(line, "branch "); found {
			current.Branch = strings.TrimPrefix(branch, "refs/heads/")
		}
	}
	flush()
	return worktrees
}
