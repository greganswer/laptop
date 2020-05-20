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
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
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
		defer outputScriptDuration(time.Now())
		makeDirectories()
		downloadBrewfileToHomeDirectory()
		executeBrewBundle()
		cloneLaptopRepo()

		// NOTE: Laptop repo must be cloned in order for the rest of the setup to work.
		// TODO: Find a way to extract these files to a separate repo
		// Reference: https://github.com/greganswer/laptop/issues/10
		symlinkDotfiles()
		setupZShell()
		installRuby()
		configureVim()
		configureVSCode()
		setupLocalDatabases()
	},
}

func init() {
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
	brewfileURL := "https://gist.githubusercontent.com/greganswer/fc93b73085b171780d8c0cfd90e6ed25/raw"
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

func cloneLaptopRepo() {
	title("Cloning laptop repo...")
	err := executeAndStream("go", "get", "github.com/greganswer/laptop")
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
