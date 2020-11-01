package conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var localYAMLFile string
var localJSONFile string

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

	localJSONFile = "../test-files/test.json"
	// This file contains
	//
	// {
	// 	"age": 29,
	// 	"name": "Jimmy",
	// 	"relationships": {
	// 		"parents": [
	// 			"Sally",
	// 			"Robert"
	// 		]
	// 	}
	// }
}

func TestCanReadUnstructuredYAMLFile(t *testing.T) {

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

func TestCanReadUnstructuredJSONFile(t *testing.T) {

	assert := assert.New(t)

	jsonData, err := UnmarshalJSONFile(localJSONFile)
	assert.Nil(err)

	expectedJSON := map[string]interface{}{
		"age":  float64(29),
		"name": "Jimmy",
		"relationships": map[string]interface{}{
			"parents": []interface{}{"Sally", "Robert"},
		},
	}
	assert.Equal(jsonData, expectedJSON)
}

func TestOutputYAMLIsCorrect(t *testing.T) {

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

func TestOutputJSONIsCorrect(t *testing.T) {
	assert := assert.New(t)

	jsonData, err := UnmarshalJSONFile(localJSONFile)
	assert.Nil(err)

	incorrectJSON := map[string]interface{}{
		"age":  float64(30),
		"name": "Ricky",
		"relationships": map[string]interface{}{
			"parents": []interface{}{"Mike", "Robert"},
		},
	}

	assert.NotEqual(jsonData, incorrectJSON)
}

func TestFullYAMLConversionIsCorrect(t *testing.T) {
	assert := assert.New(t)

	jsonData, err := YAMLToJSONFull(localYAMLFile)
	assert.Nil(err)

	var expectedType []byte
	assert.IsType(expectedType, jsonData)
}

func TestFullJSONConversionIsCorrect(t *testing.T) {
	assert := assert.New(t)

	yamlOutput, err := JSONToYAMLFull(localJSONFile)
	assert.Nil(err)

	var expectedType []byte
	assert.IsType(expectedType, yamlOutput)
}

func TestConvertYAMLToJSON(t *testing.T) {
	assert := assert.New(t)

	yamlData, err := UnmarshalYAMLFile(localYAMLFile)
	assert.Nil(err)

	jsonOutput, err := YAMLToJSON(yamlData)
	assert.Nil(err)

	var expectedType []byte
	assert.IsType(expectedType, jsonOutput)

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

func TestConvertJSONToYAML(t *testing.T) {
	assert := assert.New(t)

	jsonData, _ := UnmarshalJSONFile(localJSONFile)

	yamlOutput, err := JSONToYAML(jsonData)
	assert.Nil(err)

	t.Logf("\n%s", yamlOutput)
}
