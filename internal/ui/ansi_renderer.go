package ui

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"wt/internal/domain"

	"golang.org/x/term"
)

type ANSIRenderer struct{}

const (
	ansiReset     = "\033[0m"
	ansiBoldCyan  = "\033[1;36m"
	ansiDimGray   = "\033[2;37m"
	ansiLightGray = "\033[0;37m"
)

func (r ANSIRenderer) RenderWorktrees(worktrees []domain.Worktree) {
	if len(worktrees) == 0 {
		fmt.Fprintln(os.Stdout, colorize("No worktrees found.", ansiDimGray))
		return
	}

	width := terminalWidth()

	if width >= 60 {
		pathW, branchW, headW := threeColWidths(width, worktrees)
		r.renderTable(worktrees, []string{"PATH", "BRANCH", "HEAD"}, []int{pathW, branchW, headW})
		return
	}

	if width >= 40 {
		pathW, branchW := twoColWidths(width, worktrees)
		r.renderTable(worktrees, []string{"PATH", "BRANCH"}, []int{pathW, branchW})
		return
	}

	r.renderTable(worktrees, []string{"PATH"}, []int{maxInt(8, width-2)})
}

func (r ANSIRenderer) RenderManagedPaths(paths []string) {
	if len(paths) == 0 {
		fmt.Fprintln(os.Stdout, colorize("No managed repositories configured.", ansiDimGray))
		return
	}

	for _, path := range paths {
		fmt.Fprintln(os.Stdout, colorize("• "+strings.TrimSpace(path), ansiLightGray))
	}
}

func (r ANSIRenderer) renderTable(worktrees []domain.Worktree, headers []string, widths []int) {
	fmt.Fprintln(os.Stdout, colorize(formatRow(headers, widths), ansiBoldCyan))
	fmt.Fprintln(os.Stdout, colorize(separatorRow(widths), ansiDimGray))

	for _, wt := range worktrees {
		switch len(headers) {
		case 3:
			fmt.Fprintln(os.Stdout, colorize(formatRow([]string{wt.Path, wt.Branch, wt.HeadSHA}, widths), ansiLightGray))
		case 2:
			fmt.Fprintln(os.Stdout, colorize(formatRow([]string{wt.Path, wt.Branch}, widths), ansiLightGray))
		default:
			fmt.Fprintln(os.Stdout, colorize(formatRow([]string{wt.Path}, widths), ansiLightGray))
		}
	}
}

func colorize(s, color string) string {
	if os.Getenv("NO_COLOR") != "" {
		return s
	}
	return color + s + ansiReset
}

func formatRow(values []string, widths []int) string {
	cells := make([]string, 0, len(values))
	for i := range values {
		cells = append(cells, fitCell(values[i], widths[i]))
	}
	return strings.Join(cells, "  ")
}

func separatorRow(widths []int) string {
	parts := make([]string, 0, len(widths))
	for _, width := range widths {
		parts = append(parts, strings.Repeat("-", maxInt(1, width)))
	}
	return strings.Join(parts, "  ")
}

func fitCell(value string, width int) string {
	value = strings.TrimSpace(value)
	if width <= 1 {
		return "…"
	}

	runes := []rune(value)
	if len(runes) > width {
		return string(runes[:width-1]) + "…"
	}
	if len(runes) < width {
		return value + strings.Repeat(" ", width-len(runes))
	}
	return value
}

func terminalWidth() int {
	if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil && w > 0 {
		return w
	}
	if columns := os.Getenv("COLUMNS"); columns != "" {
		if w, err := strconv.Atoi(columns); err == nil && w > 0 {
			return w
		}
	}
	return 80
}

func twoColWidths(totalWidth int, worktrees []domain.Worktree) (int, int) {
	available := maxInt(20, totalWidth-2)
	minPath, minBranch := 10, 12
	pathCap := maxInt(minPath, available/2)

	maxBranch := maxInt(len("BRANCH"), minInt(maxBranchLen(worktrees), 28))
	branch := clamp(maxBranch, minBranch, available-minPath)
	path := available - branch
	if path > pathCap {
		extra := path - pathCap
		path = pathCap
		branch = minInt(available-path, branch+extra)
	}
	return path, branch
}

func threeColWidths(totalWidth int, worktrees []domain.Worktree) (int, int, int) {
	available := maxInt(30, totalWidth-4)
	minPath, minBranch, minHead := 16, 12, 12
	pathCap := maxInt(minPath, available/2)

	path := minPath
	branch := minBranch
	head := minHead
	extra := available - (path + branch + head)
	if extra <= 0 {
		return path, branch, head
	}

	targetHead := maxInt(len("HEAD"), minInt(maxHeadLen(worktrees), 40))
	add := minInt(extra, maxInt(0, targetHead-head))
	head += add
	extra -= add

	targetBranch := maxInt(len("BRANCH"), minInt(maxBranchLen(worktrees), 28))
	add = minInt(extra, maxInt(0, targetBranch-branch))
	branch += add
	extra -= add

	targetPath := maxInt(len("PATH"), minInt(maxPathLen(worktrees), pathCap))
	add = minInt(extra, maxInt(0, targetPath-path))
	path += add
	extra -= add

	if extra > 0 {
		head += extra
	}
	return path, branch, head
}

func maxPathLen(worktrees []domain.Worktree) int {
	max := 0
	for _, wt := range worktrees {
		if l := len([]rune(strings.TrimSpace(wt.Path))); l > max {
			max = l
		}
	}
	return max
}

func maxBranchLen(worktrees []domain.Worktree) int {
	max := 0
	for _, wt := range worktrees {
		if l := len([]rune(strings.TrimSpace(wt.Branch))); l > max {
			max = l
		}
	}
	return max
}

func maxHeadLen(worktrees []domain.Worktree) int {
	max := 0
	for _, wt := range worktrees {
		if l := len([]rune(strings.TrimSpace(wt.HeadSHA))); l > max {
			max = l
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func clamp(v, low, high int) int {
	if v < low {
		return low
	}
	if v > high {
		return high
	}
	return v
}
