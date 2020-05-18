package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update laptop dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		downloadBrewfileToHomeDirectory()
		updateBrews()
		updateZShell()
	},
}

func init() {
	// Cobra CLI setup code.
	rootCmd.AddCommand(updateCmd)
}

// Reference: https://github.com/Homebrew/homebrew-bundle
func updateBrews() {
	title("Updating Brews...")
	out, err := exec.Command("brew", "bundle", "check", "--file", brewfilePath).CombinedOutput()
	if err != nil {
		err := executeAndStream("brew", "bundle", "--file", brewfilePath)
		failIfError(err)
	}
	fmt.Println(string(out))
	finished()
}

// Reference: https://github.com/ohmyzsh/ohmyzsh
func updateZShell() {
	title("Updating zShell...")
	err := executeAndStream("git", "-C", zshellPath, "pull", "--rebase", "--stat", "origin", "master")
	warnIfError(err)
	finished()
}