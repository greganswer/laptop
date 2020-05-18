package cmd

import (
	"fmt"
	"os/exec"
	"path"
	"time"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update laptop dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		defer outputScriptDuration(time.Now())
		downloadBrewfileToHomeDirectory()
		updateBrews()
		updateZShell()
		pushRepoChanges()
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

func pushRepoChanges() {
	title("Pushing changes to laptop origin...")
	gitDir := path.Join(laptopRepoPath, ".git")
	err := executeAndStream("git", "--git-dir", gitDir, "commit", "-am", "'Update laptop settings'")
	failIfError(err)

	err = executeAndStream("git", "--git-dir", gitDir, "push")
	failIfError(err)

	finished()
}
