package cmd

import (
	"github.com/jdockerty/yaml-to-json-go/conversion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateJSONFile(t *testing.T) {

	assert := assert.New(t)

	jsonFileToCreate := "../test-files/output.json"

	err := createOutputFile(jsonFileToCreate)
	assert.Nil(err)


	assert.FileExists(jsonFileToCreate)
	t.Logf("file created.\n")

}

func TestShouldCreateYAMLFile(t *testing.T) {

	assert := assert.New(t)

	yamlFileToCreate := "../test-files/output.yml"

	err := createOutputFile(yamlFileToCreate)
	assert.Nil(err)


	assert.FileExists(yamlFileToCreate)
	t.Logf("file created.\n")

}

func TestCorrectFileExtensions(t *testing.T) {
	yamlFile := "../test-files/test.yml"
	jsonFile := "../test-files/test.json"

	extensions := fileExts(yamlFile, jsonFile)

	assert.Equal(t, ".yml", extensions[0])
	assert.Equal(t, ".json", extensions[1])
}

func TestWriteDataToFile(t *testing.T) {

	assert := assert.New(t)

	fileToRead := "../test-files/test.yml"

	outputFile := "../test-files/dataOut.yaml"

	yamlData, err := conversion.YAMLToJSONFull(fileToRead)
	assert.Nil(err)

	err = writeToFile(yamlData, outputFile)

	assert.Nil(err)
	t.Logf("data written to %s", outputFile)
}

// func TestCanReadPassedFile(t *testing.T) {
// 	assert := assert.New(t)

// 	myYAMLFile := "../conversion/test.yml"

// 	data, err := readFileInArgs(myYAMLFile)
// 	assert.Nil(err)

// 	assert.IsType([]byte, data)


// }
