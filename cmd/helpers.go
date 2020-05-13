package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

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
