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
	Long: `Quickly validate a list of YAML or JSON files. Be aware that there is no fancy implementation here, it is simply used to tell whether a file is considered valid enough to be parsed.

The general concept is

	yamltojson validate file1 file2 ...

For example:

	yamltojson validate myFile.yml anotherFile.yml
	yamltojson validate yamlConfigs/*

This also works with a mix of JSON and YAML files

	yamltojson validate myfiles/* imports/data.json config.yml`,
	RunE: runValidateCmd,
}

func runValidateCmd(cmd *cobra.Command, args []string) error {

	for _, file := range args {
		isYAML := conversion.IsYAMLFile(file)
		isJSON := conversion.IsJSONFile(file)
		if isYAML {
			data, err := conversion.UnmarshalYAMLFile(file)
			if err != nil {
				return err
			}

			if len(data) == 0 {
				fmt.Printf("%s is invalid YAML.\n", file)
				continue
			}

			fmt.Printf("%s is valid YAML.\n", file)

		} else if isJSON {
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
