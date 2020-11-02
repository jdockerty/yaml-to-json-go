package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yamltojson",
	Short: "A simple CLI tool to convert YAML to JSON or JSON to YAML",
	Long: `yamltojson lets you easily convert a YAML file to a JSON file or the other way. For example:

	yamltojson convert path/to/source_file path/to/write/output_file
	
	This lets you convert between JSON or YAML, depending on the file extension provided in the source file.

	yamltojson convert -p path/to/file.json | .yaml | .yml
	
	Allows you to print the output to the relevant corresponding source file, without writing to another.
	Specifying a .json file would convert print the YAML, whereas a source file of .yml or .yaml would print relevant JSON.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() int {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".yamltojson" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".yamltojson")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
