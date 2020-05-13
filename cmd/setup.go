/*

References:
- https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
- https://zaiste.net/executing_external_commands_in_go/
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
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
		c := exec.Command("brew", "bundle", "--file", brewfilePath)
		cmdOutput := &bytes.Buffer{}
		c.Stdout = cmdOutput
		failIfError(c.Run())

		fmt.Print(string(cmdOutput.Bytes()))
		finished()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}

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
