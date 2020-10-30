package cmd

import (
	// "github.com/stretchr/testify/assert"
	"github.com/google/go-cmdtest"
	"testing"
)

func TestShouldCreateJSONFile(t *testing.T) {

}

func TestShouldCreateYAMLFile(t *testing.T) {

}

func TestCanOpenPassedFile(t *testing.T) {
	// assert := assert.New(t)

	// myYAMLFile := "../conversion/test.yml"

	// assert.FileExists(myYAMLFile)
	// testcli.Run("yamltojson convert ../conversion/test.yml output.json")
	testSuite, err := cmdtest.Read("cli-tests/convert")
	
}
