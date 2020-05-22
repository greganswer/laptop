package cmd

import (
	"fmt"
	"os/user"
	"path"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var currentUser *user.User

// File paths.
var (
	brewfilePath   string
	laptopRepoPath string
	zshellPath     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "laptop",
	Version: "0.3.1",
	Short:   "Laptop setup and update script for Mac users",
	Long: `Setup software development dependencies like 
settings, databases, dot files, and apps (CLI & native)
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	failIfError(rootCmd.Execute())
}

func init() {
	// Set package variables.
	var err error
	currentUser, err = user.Current()
	failIfError(err)

	brewfilePath = path.Join(currentUser.HomeDir, "Brewfile")
	laptopRepoPath = path.Join(currentUser.HomeDir, "go", "src", "github.com", "greganswer", "laptop")
	zshellPath = path.Join(currentUser.HomeDir, ".oh-my-zsh")

	// Cobra configs.
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.laptop.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".laptop" (without extension).
		viper.AddConfigPath(currentUser.HomeDir)
		viper.SetConfigName(".laptop")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
