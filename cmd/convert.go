package cmd

import (
	"errors"
	"fmt"
	"github.com/jdockerty/yaml-to-json-go/conversion"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts a YAML file into a JSON file or the other way.",
	Long: `Convert a YAML or JSON file to the opposing format, depending on which one is placed first. 

For example:

	yamltojson convert path/to/myConfig.yml path/to/place/output.json

This would convert the YAML file into JSON, since the YAML file is placed first. To do it the other way, swap them around

	yamltojson convert path/to/output.json path/to/place/myConfig.yml

This assumes that you have the first file at the location specified, the second file will be created with the conversion in the current directory.`,
	RunE: runConvertCmd,
}

// PrintFlag is the --print or -p flag to print the output to the terminal, rather than writing to a file.
var PrintFlag bool

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().BoolVarP(&PrintFlag, "print", "p", false, "print the output to terminal, instead of writing to a target file.")
}

func runConvertCmd(cmd *cobra.Command, args []string) error {

	if PrintFlag {

		if len(args) == 1 {
			return runConvertPrintFlag(args[0])
		}

		return fmt.Errorf("only a single file should be specific with the --print flag")

	}
	if len(args) != 2 {
		return errors.New("you must only specify 2 files")
	}

	return runConvert(args)

}

func runConvertPrintFlag(file string) error {

	if fileType := fileExt(file); fileType == ".yaml" || fileType == ".yml" {

		jsonOutput, err := conversion.YAMLToJSONFull(file)
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", string(jsonOutput))
	} else if fileType == ".json" {

		yamlOutput, err := conversion.JSONToYAMLFull(file)
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", string(yamlOutput))
	} else {
		return fmt.Errorf("only .yaml, .yml, or .json file extensions are supported")
	}

	return nil
}

func createOutputFile(f string) error {

	_, err := os.Create(f)
	if err != nil {
		return err
	}

	return nil
}

func fileExt(file string) string {
	return filepath.Ext(file)
}

func runConvert(args []string) error {
	sourceFile, targetFile := args[0], args[1]

	err := createOutputFile(targetFile)
	if err != nil {
		return err
	}

	if fileType := fileExt(sourceFile); fileType == ".yml" || fileType == ".yaml" {

		jsonData, err := conversion.YAMLToJSONFull(sourceFile)
		if err != nil {
			return err
		}

		err = writeToFile(jsonData, targetFile)
		if err != nil {
			return err
		}

		fmt.Printf("Converting %s to %s\n", sourceFile, targetFile)

	} else if fileType == ".json" {

		yamlData, err := conversion.JSONToYAMLFull(sourceFile)
		if err != nil {
			return err
		}

		err = writeToFile(yamlData, targetFile)
		if err != nil {
			return err
		}
		fmt.Printf("Converting %s to %s\n", sourceFile, targetFile)

	} else {
		return fmt.Errorf("only .yml, .yaml, or .json file extensions are supported")
	}

	return nil

}

func writeToFile(data []byte, file string) error {

	err := ioutil.WriteFile(file, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
