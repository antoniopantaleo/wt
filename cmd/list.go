package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runList,
}

func runList(cmd *cobra.Command, args []string) {
	command := exec.Command("git", "worktree", "list", "--porcelain")
	out, err := command.Output()
	if err != nil {
		fmt.Println("Error executing git worktree list:", err)
		return
	}
	fmt.Println(string(out))
}

func init() {
	rootCmd.AddCommand(listCmd)
}
