package conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var localYAMLFile string

func init() {
	localYAMLFile = "../test-files/test.yml"
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

	jsonData, err := YAMLToJSONFull(localYAMLFile)
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
	// "age": 29,
	// "name": "Jimmy",
	// "relationships": {
	// "parents": [
	//	"Sally", 
	//	"Robert"
	//		]
	// 	}
	// }
}
