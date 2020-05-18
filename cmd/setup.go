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
	"os/user"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	brewfilePath, laptopRepoPath, zshellPath string
	currentUser                              *user.User
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the laptop for Software Development",
	Long: `Setup the following on your new Mac:

    - HomeBrew
    - Dot files
    - Native apps (Chrome, Spotify, Vs Code, etc.)
    - CLI apps (Git, Hub, Docker-compose, etc.)
    - Local Databases (Postgres, Redis, etc.)
    - Settings for your apps
`,
	Run: func(cmd *cobra.Command, args []string) {
		makeDirectories()
		downloadBrewfileToHomeDirectory()
		executeBrewBundle()
		symlinkDotfiles()
		setupZShell()
		installRuby()
		configureVim()
		configureVSCode()
		setupLocalDatabases()
	},
}

func init() {
	// Set package variables.
	var err error
	currentUser, err = user.Current()
	failIfError(err)

	brewfilePath = path.Join(currentUser.HomeDir, "Brewfile")
	laptopRepoPath = path.Join(currentUser.HomeDir, "go", "src", "github.com", "greganswer", "laptop")
	zshellPath = path.Join(currentUser.HomeDir, ".oh-my-zsh")

	// Cobra CLI setup code.
	rootCmd.AddCommand(setupCmd)
}

func makeDirectories() {
	title("Creating directories...")
	paths := []string{
		filepath.Join(currentUser.HomeDir, "go", "src"),
		filepath.Join(currentUser.HomeDir, "go", "bin"),
	}
	for _, p := range paths {
		fmt.Println(p)
		os.MkdirAll(p, os.ModePerm)
	}
	finished()
}

// Reference: https://brew.sh/
func downloadBrewfileToHomeDirectory() {
	title("Downloading Brewfile to home directory...")
	brewfileURL := "https://bit.ly/2xUtgmK"
	err := downloadFile(brewfileURL, brewfilePath)
	failOrOK(err)
}

// Reference: https://github.com/Homebrew/homebrew-bundle
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
		destination := filepath.Join(currentUser.HomeDir, fmt.Sprintf(".%s", f.Name()))
		os.Remove(destination)
		err := os.Symlink(source, destination)
		failIfError(err)
	}
	finished()
}

// Reference: https://github.com/ohmyzsh/ohmyzsh
func setupZShell() {
	title("Setting up zShell...")
	err := executeAndStream("git", "clone", "git://github.com/robbyrussell/oh-my-zsh.git", zshellPath)
	warnIfError(err)

	err = executeAndStream("chsh", "-s", "/bin/zsh")
	warnIfError(err)

	err = executeAndStream("zsh", "--version")
	warnIfError(err)

	finished()
}

func installRuby() {
	rubyVersion := "2.6.0"
	title(fmt.Sprintf("Installing Ruby %s...", rubyVersion))
	err := executeAndStream("rbenv", "install", rubyVersion)
	warnIfError(err)

	err = executeAndStream("rbenv", "global", rubyVersion)
	failIfError(err)

	err = executeAndStream("gem", "install", "bundler")
	failIfError(err)

	finished()
}

// Reference: https://github.com/VundleVim/Vundle.vim
func configureVim() {
	title("Configure VIM...")
	err := executeAndStream("git", "clone", "https://github.com/VundleVim/Vundle.vim.git", "~/.vim/bundle/Vundle.vim")
	warnIfError(err)

	err = executeAndStream("vim", "+PluginInstall", "+qall")
	warnIfError(err)

	finished()
}

func configureVSCode() {
	title("Configure VS Code...")
	source := filepath.Join(laptopRepoPath, "settings", "vscode.json")
	destination := filepath.Join(currentUser.HomeDir, "Library", "Application Support", "Code", "User", "settings.json")
	os.Remove(destination)
	err := os.Symlink(source, destination)
	failIfError(err)
	finished()
}

func setupLocalDatabases() {
	title("Setup local databases...")
	err := executeAndStream("brew", "services", "start", "postgresql")
	failIfError(err)

	err = executeAndStream("createdb", currentUser.Username)
	warnIfError(err)

	finished()
}
