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
	t.Logf("JSON file created.\n")

}

func TestShouldCreateYAMLFile(t *testing.T) {

	assert := assert.New(t)

	yamlFileToCreate := "../test-files/output.yml"

	err := createOutputFile(yamlFileToCreate)
	assert.Nil(err)

	assert.FileExists(yamlFileToCreate)
	t.Logf("YAML file created.\n")

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

func TestDirectoryExists(t *testing.T) {

	assert := assert.New(t)

	nonExistentDir := "some/crazy/directory/in/space"
	dirDoesNotExist, err := dirExists(nonExistentDir)
	t.Logf("%s exists: %t\n", nonExistentDir, dirDoesNotExist)
	assert.Nil(err)
	assert.Equal(false, dirDoesNotExist)

	existingDir := "../conversion/"
	directoryExists, err := dirExists(existingDir)
	t.Logf("%s exists: %t\n", existingDir, directoryExists)
	assert.Nil(err)
	assert.Equal(true, directoryExists)
}
