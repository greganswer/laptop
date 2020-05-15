package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func title(s string) {
	c := color.New(color.FgMagenta, color.Bold)
	c.Println(s)
}

// failIfError exits the program with a standardized error message if an error occurred.
func failIfError(err error) {
	if err != nil {
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		fmt.Println(red("FAIL:"), err)
		os.Exit(1)
	}
}

// failOrOK displays a failure message if there's an error otherwise displays "OK".
func failOrOK(err error) {
	failIfError(err)
	finished()
}

func finished() {
	c := color.New(color.FgGreen, color.Bold)
	c.Println("OK")
	fmt.Println()
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
// If the file exists, it is over written.
// Reference: https://golangcode.com/download-a-file-from-a-url
func downloadFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
