package convert

import (
	"github.com/stretchr/testify/assert"
	// "github.com/jdockerty/yaml-to-json-go/convert"
	"testing"
)

var localYAMLFile string


func init() {
	localYAMLFile = "test.yml"
	// This file contains
	//
	// name: Jimmy
	// age: 29
	// relationships:
	// 	parents:
	// 		- Sally
	// 		- Robert
}

func TestCanReadUnstructuredYamlFile(t *testing.T) {

	assert := assert.New(t)

	yamlData, err := UnmarshalYAMLFile(localYAMLFile)
	assert.Nil(err)

	expectedYaml := map[interface{}]interface{}{
		"age":  29,
		"name": "Jimmy",
		"relationships": map[interface{}]interface{}{
			"parents": []interface{}{"Sally", "Robert"},
		},
	}

	assert.Equal(yamlData, expectedYaml)
}

func TestOutputYamlIsCorrect(t *testing.T) {

	assert := assert.New(t)

	yamlData, err := UnmarshalYAMLFile(localYAMLFile)
	assert.Nil(err)

	incorrectYaml := map[interface{}]interface{}{
		"age":  105,
		"name": "Alex",
		"relationships": map[interface{}]interface{}{
			"parents": []interface{}{"Jim", "Jane"},
		},
	}

	assert.NotEqual(yamlData, incorrectYaml)

}

func TestFullConversionIsCorrect(t *testing.T) {
	assert := assert.New(t)

	jsonData, err := FullYAMLToJSON(localYAMLFile)
	assert.Nil(err)

	var expectedType []byte
	assert.IsType(expectedType, jsonData)
}

func TestConvertYamlToJSON(t *testing.T) {
	assert := assert.New(t)

	yamlData, err := UnmarshalYAMLFile(localYAMLFile)
	assert.Nil(err)

	jsonOutput, err := YAMLToJSON(yamlData)
	assert.Nil(err)

	t.Logf("%s\n", jsonOutput)
	// Output:{
	// "name": "Jimmy",
	// "age": 29,
	// "relationships": {
	// "parents": ["Sally", "Robert"]
	// }
	// }
}
