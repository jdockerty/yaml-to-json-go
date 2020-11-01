/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
	"os"
	"io/ioutil"
	"github.com/jdockerty/yaml-to-json-go/conversion"
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

func fileExts(fileOne, fileTwo string) []string {
	var extensions []string

	extensions = append(extensions, filepath.Ext(fileOne), filepath.Ext(fileTwo))

	return extensions
}

func runConvert(args []string) error {
	sourceFile, targetFile := args[0], args[1]
	jsonData, _ := conversion.YAMLToJSONFull(sourceFile)
	createOutputFile(targetFile)
	writeToFile(jsonData, targetFile)
	fmt.Printf("Converting %s into %s\n", sourceFile, targetFile)
	fmt.Println(jsonData)

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
