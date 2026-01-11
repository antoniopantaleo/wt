package ui

import (
	"fmt"
	"wt/internal/domain"
)

type FmtRenderer struct{}

func (r FmtRenderer) RenderWorktrees(worktrees []domain.Worktree) {
	for _, worktree := range worktrees {
		fmt.Printf("PATH\t\tBRANCH\t\tHEAD\n%v\t\t%v\t\t%v", worktree.Path, worktree.Branch, worktree.HeadSHA)
	}
}
