/*

References:
- https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
- https://zaiste.net/executing_external_commands_in_go/
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	brewfilePath, homeDir, laptopRepoPath, zshellPath string
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
		symlinkDotfiles()
		setupZShell()
	},
}

func init() {
	// Set package variables.
	var err error
	homeDir, err = os.UserHomeDir()
	failIfError(err)

	brewfilePath = path.Join(homeDir, "Brewfile")
	laptopRepoPath = path.Join(homeDir, "go", "src", "github.com", "greganswer", "laptop")
	zshellPath = path.Join(homeDir, ".oh-my-zsh")

	// Cobra CLI setup code.
	rootCmd.AddCommand(setupCmd)
}

func makeDirectories() {
	title("Creating directories...")
	paths := []string{
		filepath.Join(homeDir, "go", "src"),
		filepath.Join(homeDir, "go", "bin"),
	}
	for _, p := range paths {
		fmt.Println(p)
		os.MkdirAll(p, os.ModePerm)
	}
	finished()
}

func downloadBrewfileToHomeDirectory() {
	title("Downloading Brewfile to home directory...")
	brewfileURL := "https://bit.ly/2xUtgmK"
	err := downloadFile(brewfileURL, brewfilePath)
	failOrOK(err)
}

func executeBrewBundle() {
	title("Executing brew bundle...")
	err := executeAndStream("brew", "bundle", "--file", brewfilePath)
	failIfError(err)
	finished()
}

func symlinkDotfiles() {
	title("Symlinking dotfiles...")
	dotfilesPath := filepath.Join(laptopRepoPath, "dotfiles")
	files, err := ioutil.ReadDir(dotfilesPath)
	failIfError(err)

	for _, f := range files {
		fmt.Printf("Symlinking .%s\n", f.Name())
		source := filepath.Join(dotfilesPath, f.Name())
		destination := filepath.Join(homeDir, fmt.Sprintf(".%s", f.Name()))
		os.Remove(destination)
		err := os.Symlink(source, destination)
		failIfError(err)
	}
	finished()
}

func setupZShell() {
	title("Setting up zShell...")
	err := executeAndStream("git", "clone", "git://github.com/robbyrussell/oh-my-zsh.git", zshellPath)
	failIfError(err)
	finished()
}

// git clone git://github.com/robbyrussell/oh-my-zsh.git ~/.oh-my-zsh
