package cmd

import (
	"errors"
	"fmt"
	"github.com/jdockerty/yaml-to-json-go/conversion"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

This assumes that you have the first file at the location specified, the second file will be created with the conversion in the current directory.

You can convert a full directory using the --directory or -d flag, although the output directory DOES NOT preserve the structure. This works best when the directory contains all of the same type, e.g. all JSON files.
	
	yamltojson convert --directory="source-dir,target-dir"
	yamltojson convert --directory="many-json-files-here/,/home/jimmy/work/configs/"

This converts all of the files within the directory to their corresponding counterpart, as such JSON will convert to YAML and YAML to JSON.
Remember to specify the source and target directory without a space, separated by a comma. `,
	RunE: runConvertCmd,
}

// PrintFlag is the --print or -p flag to print the output to the terminal, rather than writing to a file.
var PrintFlag bool

// DirectoryFlag is the --directory or -d flag. This is used to specify a directory containing files to convert and the output directory.
var DirectoryFlag []string

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().BoolVarP(&PrintFlag, "print", "p", false, "print the output to terminal, instead of writing to a target file.")
	convertCmd.Flags().StringSliceVarP(&DirectoryFlag, "directory", "d", []string{}, "specify a directory containing all YAML or all JSON files.")

}

func runConvertCmd(cmd *cobra.Command, args []string) error {

	if PrintFlag {

		if len(args) == 1 {
			return runConvertPrintFlag(args[0])
		}

		return fmt.Errorf("only a single file should be specific with the --print flag")

	} else if len(DirectoryFlag) != 0 {

		if len(DirectoryFlag) == 2 {
			return runConvertDirFlag(DirectoryFlag)
		}

		return fmt.Errorf("flag should be specified in the form: --directory='sourceDir, targetDir'")

	}

	if len(args) != 2 {
		return errors.New("you must only specify 2 files")
	}

	return runConvert(args)

}

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		isNotDirectory := !info.IsDir()
		if isNotDirectory {
			files = append(files, path)
		}

		return nil
	})

	return files, err

}

// exists returns whether the given file or directory exists
func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func runConvertDirFlag(args []string) error {
	sourceDir := args[0]
	destDir := args[1]

	directoryExists, _ := dirExists(destDir)

	// If the directory does NOT (!) exist, return an error.
	if !directoryExists {
		return fmt.Errorf("the target directory must exist in order to write the files")
	}

	files, err := filePathWalkDir(sourceDir)
	if err != nil {
		return err
	}

	for _, file := range files {

		isJSON := conversion.IsJSONFile(file)
		isYAML := conversion.IsYAMLFile(file)

		fullFileName := filepath.Base(file)
		fileName := strings.TrimSuffix(fullFileName, filepath.Ext(fullFileName))

		if isJSON {
			filePath := fmt.Sprintf("%s/%s.yml", destDir, fileName)

			yamlOutput, err := conversion.JSONToYAMLFull(file)
			if err != nil {
				return err
			}

			writeToFile(yamlOutput, filePath)
			fmt.Printf("Converted %s to YAML\n", file)

		} else if isYAML {
			filePath := fmt.Sprintf("%s/%s.json", destDir, fileName)
			jsonOutput, err := conversion.YAMLToJSONFull(file)
			if err != nil {
				return err
			}
			writeToFile(jsonOutput, filePath)
			fmt.Printf("Converted %s to JSON\n", file)
		}

	}

	return nil
}

func runConvertPrintFlag(file string) error {

	if conversion.IsYAMLFile(file) {

		jsonOutput, err := conversion.YAMLToJSONFull(file)
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", string(jsonOutput))
	} else if conversion.IsJSONFile(file) {

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

func runConvert(args []string) error {
	sourceFile, targetFile := args[0], args[1]

	err := createOutputFile(targetFile)
	if err != nil {
		return err
	}

	isYAML := conversion.IsYAMLFile(sourceFile)
	isJSON := conversion.IsJSONFile(sourceFile)
	if isYAML {

		jsonData, err := conversion.YAMLToJSONFull(sourceFile)
		if err != nil {
			return err
		}

		err = writeToFile(jsonData, targetFile)
		if err != nil {
			return err
		}

		fmt.Printf("Converting %s to %s\n", sourceFile, targetFile)

	} else if isJSON {

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
