package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/jdockerty/yaml-to-json-go/convert"
	"testing"
)

func TestCanReadUnstructuredYamlFile(t *testing.T) {

	assert := assert.New(t)

	yamlFile := "test.yml"
	// This file contains
	//
	// name: Jimmy
	// age: 29
	// relationships:
	// 	parents:
	// 		- Sally
	// 		- Robert
	
	yamlData, err := convert.UnmarshalYAMLFile(yamlFile)
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

	yamlFile := "test.yml"
	// This file contains
	//
	// name: Jimmy
	// age: 29
	// relationships:
	// 	parents:
	// 		- Sally
	// 		- Robert

	yamlData, err := convert.UnmarshalYAMLFile(yamlFile)
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

func TestConvertYamlToJSON(t *testing.T) {
	assert := assert.New(t)

	yamlFile := "test.yml"
	// This file contains
	//
	// name: Jimmy
	// age: 29
	// relationships:
	// 	parents:
	// 		- Sally
	// 		- Robert

	yamlData, err := convert.UnmarshalYAMLFile(yamlFile)
	assert.Nil(err)

	jsonOutput, err := convert.YAMLToJSON(yamlData)
	assert.Nil(err)

	t.Logf("%s\n", jsonOutput)
	fmt.Println(jsonOutput)
	// Output:{
	// "name": "Jimmy",
	// "age": 29,
	// "relationships": {
	// "parents": ["Sally", "Robert"]
	// }
	// }
}
