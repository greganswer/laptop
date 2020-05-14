/*

References:
- https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
- https://zaiste.net/executing_external_commands_in_go/
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	brewfilePath, homeDir string
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		makeDirectories()
		downloadBrewfileToHomeDirectory()
		executeBrewBundle()
	},
}

func makeDirectories() {
	fmt.Println("Creating directories...")
	paths := []string{
		filepath.Join(homeDir, "go", "src"),
		filepath.Join(homeDir, "go", "bin"),
	}
	for _, p := range paths {
		os.MkdirAll(p, os.ModePerm)
	}
	finished()
}

func init() {
	// Set package variables.
	var err error
	homeDir, err = os.UserHomeDir()
	failIfError(err)
	brewfilePath = path.Join(homeDir, "Brewfile")

	// Cobra CLI setup code.
	rootCmd.AddCommand(setupCmd)
}

func executeBrewBundle() {
	fmt.Println("Executing brew bundle...")
	err := executeAndStream("brew", "bundle", "--file", brewfilePath)
	failIfError(err)
	finished()
}

func downloadBrewfileToHomeDirectory() {
	fmt.Println("Downloading Brewfile to home directory...")
	brewfileURL := "https://bit.ly/2xUtgmK"
	err := downloadFile(brewfileURL, brewfilePath)
	failOrOK(err)
}
