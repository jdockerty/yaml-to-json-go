package cmd

import (
	"fmt"

	"github.com/jdockerty/yaml-to-json-go/conversion"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a YAML or JSON file with this command.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runValidateCmd,
}

func runValidateCmd(cmd *cobra.Command, args []string) error {

	for _, file := range args {

		if fileType := fileExt(file); fileType == ".yaml" || fileType == ".yml" {
			data, err := conversion.UnmarshalYAMLFile(file)
			if err != nil {
				return err
			}

			if len(data) == 0 {
				fmt.Printf("%s is invalid YAML.\n", file)
				continue
			}

			fmt.Printf("%s is valid YAML.\n", file)

		} else if fileType == ".json" {
			data, err := conversion.UnmarshalJSONFile(file)
			if err != nil {
				return err
			}

			if len(data) == 0 {
				fmt.Printf("%s is invalid JSON.\n", file)
				continue
			}

			fmt.Printf("%s is valid JSON.\n", file)

		} else {
			return fmt.Errorf("%s is an invalid file, ensure the path is correct", file)

		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
