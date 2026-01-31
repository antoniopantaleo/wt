package ui

import (
	"fmt"
	"os"
	"text/tabwriter"
	"wt/internal/domain"
)

type FmtRenderer struct{}

func (r FmtRenderer) RenderWorktrees(worktrees []domain.Worktree) {
	if len(worktrees) == 0 {
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 15, 0, 3, ' ', 0)
	fmt.Fprintln(w, "PATH\tBRANCH\tHEAD")
	fmt.Fprintln(w, "----\t------\t----")
	for _, worktree := range worktrees {
		fmt.Fprintf(w, "%s\t%s\t%s\n", worktree.Path, worktree.Branch, worktree.HeadSHA)
	}
	w.Flush()
}

func (r FmtRenderer) RenderManagedPaths(paths []string) {
	for _, path := range paths {
		fmt.Println(path)
	}
}