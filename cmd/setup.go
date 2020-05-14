/*

References:
- https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
- https://zaiste.net/executing_external_commands_in_go/
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Download the Brewfile to home directory.
		fmt.Println("Downloading Brewfile to home directory...")
		home, err := os.UserHomeDir()
		failIfError(err)

		brewfileURL := "https://bit.ly/2xUtgmK"
		brewfilePath := path.Join(home, "Brewfile")
		err = downloadFile(brewfileURL, brewfilePath)
		failOrOK(err)

		// Execute brew bundle command.
		fmt.Println("Executing brew bundle...")
		err = executeAndStream("brew", "bundle", "--file", brewfilePath)
		failIfError(err)
		finished()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}

// executeAndStream executes a shell command and streams the output to the terminal.
// Reference: https://stackoverflow.com/a/45957859
func executeAndStream(name string, arg ...string) error {
	c := exec.Command(name, arg...)
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	if err := c.Start(); err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return c.Wait()
}

// downloadFile copies a file from the internet to the specified file path.
// Reference: https://golangcode.com/download-a-file-from-a-url
func downloadFile(url, filePath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
