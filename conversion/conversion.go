package conversion

import (
	"path/filepath"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// UnmarshalYAMLFile will return a YAML configuration file in it's raw Golang form.
func UnmarshalYAMLFile(filePath string) (map[interface{}]interface{}, error) {

	yamlData := make(map[interface{}]interface{})

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error when opening YAML file: %s", err.Error())
	}

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %s", err.Error())
	}

	err = yaml.Unmarshal(fileData, &yamlData)

	return yamlData, nil

}

// UnmarshalJSONFile will return a JSON file in it's raw Golang form.
func UnmarshalJSONFile(filePath string) (map[string]interface{}, error) {

	var jsonData map[string]interface{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error when opening JSON file: %s", err.Error())
	}

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error when opening JSON file: %s", err.Error())
	}

	json.Unmarshal(fileData, &jsonData)

	return jsonData, nil
}

// JSONToYAML will convert raw JSON into a byte array ready to be written to a file.
func JSONToYAML(jsonData map[string]interface{}) ([]byte, error) {

	output, err := yaml.Marshal(jsonData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling YAML: %s", err.Error())
	}

	return output, nil
}

// YAMLToJSON will convert raw YAML into a JSON encoded byte array, this is ready to be written to a file.
func YAMLToJSON(yamlData map[interface{}]interface{}) ([]byte, error) {

	cleanedYaml := cleanYaml(yamlData)

	output, err := json.MarshalIndent(cleanedYaml, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("error converting yaml to json: %s", err.Error())
	}

	return output, nil
}

// Allows YAML to be encoded into JSON format easily.
func cleanYaml(yamlData map[interface{}]interface{}) map[string]interface{} {

	cleanYamlMapping := make(map[string]interface{})

	for key, value := range yamlData {

		// Type assertion on the key within the yaml, key will be type of interface{}
		// so it must be asserted to ensure it is string.
		assertedKey := key.(string)
		cleanYamlMapping[assertedKey] = value

		assertedMapVal, isInterfaceMapType := value.(map[interface{}]interface{})
		assertedSliceVal, isInterfaceSliceType := value.([]interface{})

		// If the value is also another map, then you need to retreive that value, adding it into the outer map.
		if isInterfaceMapType {
			cleanInnerMap := cleanYaml(assertedMapVal)
			cleanYamlMapping[assertedKey] = cleanInnerMap
		}

		// If the item is a interface slice, we need to check whether it contains a map[interface{}]interface{} type, if so we can convert it.
		if isInterfaceSliceType {
			for _, item := range assertedSliceVal {

				itemAsserted, isInnerMap := item.(map[interface{}]interface{})

				if isInnerMap {
					cleanInnerMap := cleanYaml(itemAsserted)
					cleanYamlMapping[assertedKey] = cleanInnerMap
				}

			}

		}
	}

	return cleanYamlMapping
}

// YAMLToJSONFull is a wrapper function around the other underlying functions
// for ease of use. Simply, a file is specified and the conversion is handled internally.
func YAMLToJSONFull(filePath string) ([]byte, error) {

	yamlData, err := UnmarshalYAMLFile(filePath)
	if err != nil {
		return nil, err
	}

	jsonOutput, err := YAMLToJSON(yamlData)
	if err != nil {
		return nil, err
	}

	return jsonOutput, nil
}

// JSONToYAMLFull is a wrapper function around the other underlying functions
// for ease of use. Simply, a file is specified and the conversion is handled internally.
func JSONToYAMLFull(filePath string) ([]byte, error) {

	jsonData, err := UnmarshalJSONFile(filePath)
	if err != nil {
		return nil, err
	}

	yamlOutput, err := JSONToYAML(jsonData)
	if err != nil {
		return nil, err
	}

	return yamlOutput, nil
}

// IsJSONFile checks whether a specified file is JSON.
func IsJSONFile(filePath string) bool {

	fileType := filepath.Ext(filePath)

	if fileType == ".json" {
		return true
	}

	return false
}

// IsYAMLFile checks whether a specified file is YAML.
func IsYAMLFile(filePath string) bool {

	fileType := filepath.Ext(filePath)

	if fileType == ".yml" || fileType == ".yaml" {
		return true
	}

	return false
}