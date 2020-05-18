package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
)

func outputScriptDuration(start time.Time) {
	end := time.Now()
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	fmt.Println(green("DONE: "), fmtDuration(end.Sub(start)))
}

// Reference: https://stackoverflow.com/a/47342272
func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	return fmt.Sprintf("%01dh:%02dm:%02ds", h, m, s)
}

func title(s string) {
	c := color.New(color.FgMagenta, color.Bold)
	c.Println(s)
}

// warnIfError displays a standardized warning message if an error occurred.
func warnIfError(err error) {
	if err != nil {
		warn(err.Error())
	}
}

func warn(s string) {
	yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
	os.Stderr.WriteString(fmt.Sprint(yellow("WARN: "), s, "\n"))
}

// failIfError exits the program with a standardized error message if an error occurred.
func failIfError(err error) {
	if err != nil {
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		os.Stderr.WriteString(fmt.Sprint(red("FAIL: "), err, "\n"))
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

	// Setup stdout and stderr.
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}

	// Start the command.
	if err := c.Start(); err != nil {
		return err
	}

	// Stream stdout and stderr.
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	scanner = bufio.NewScanner(stderr)
	for scanner.Scan() {
		warn(scanner.Text())
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
