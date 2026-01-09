package git

import (
	"fmt"
	"os/exec"
	"strings"
	"wt/internal/domain"
)

type ExecGit struct{}

func (g ExecGit) GetWorktreesFromPath(path string) []domain.Worktree {
	var out strings.Builder
	cmd := exec.Command("git", "-C", path, "rev-parse", "--path-format=absolute", "--git-common-dir", "--show-toplevel")
	cmd.Stdout = &out
	_ = cmd.Run()
	fmt.Println("RESULT", out.String())
	return []domain.Worktree{}
}
