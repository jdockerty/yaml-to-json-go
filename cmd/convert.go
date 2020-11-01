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

	yamltojson convert myConfig.yml output.json

This would convert the YAML file into JSON, since the YAML file is placed first. To do it the other way, swap them around

	yamltojson convert output.json myConfig.yml

This assumes that you have the first file specified at the location specified, the second file will be created with the conversion in the current directory.`,
	RunE: runConvertCmd,
}

func runConvertCmd(cmd *cobra.Command, args []string) error {

	if len(args) != 2 {
		return errors.New("you must only specify 2 files")
	}

	return runConvert(args)

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

func init() {
	rootCmd.AddCommand(convertCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
