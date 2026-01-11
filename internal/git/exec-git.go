package git

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
	"wt/internal/domain"
)

type ExecGit struct{}

func (g ExecGit) GetWorktreesFromPath(path string) []domain.Worktree {
	cmd := exec.Command("git", "-C", path, "rev-parse", "--path-format=absolute", "--git-common-dir", "--show-toplevel")
	var out bytes.Buffer
	// TODO: permission error when using cmd.Path = path
	cmd.Stdout = &out
	cmd.Run()
	infos := strings.Split(out.String(), "\n")
	repoPath := strings.TrimSuffix(infos[0], "/.git")
	worktreePath := infos[1]
	cmd = exec.Command("git", "-C", worktreePath, "worktree", "list", "--porcelain")
	data, err := cmd.Output()
	if err != nil {
		return nil
	}
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
		if line == "" {
			flush()
			continue
		}
		if worktreePath, found := strings.CutPrefix(line, "worktree "); found {
			current = &domain.Worktree{Path: worktreePath, RepoPath: repoPath}
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
