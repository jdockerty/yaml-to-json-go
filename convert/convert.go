package convert

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// UnmarshalYAMLFile will return a YAML configuration file in it's raw form.
func UnmarshalYAMLFile(filePath string) (map[interface{}]interface{}, error) {

	yamlData := make(map[interface{}]interface{})

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("error when opening YAML file.\n%s", err.Error())
		return nil, err
	}

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("error reading file.\n%s", err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(fileData, &yamlData)

	return yamlData, nil

}

// YAMLToJSON will convert raw YAML into a JSON encoded byte array.
func YAMLToJSON(yamlData map[interface{}]interface{}) ([]byte, error) {

	cleanedYaml := cleanYaml(yamlData)

	output, err := json.MarshalIndent(cleanedYaml, "", "\t")
	if err != nil {
		log.Printf("error converting yaml to json.\n%s", err.Error())
		return nil, err
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

		assertedVal, isType := value.(map[interface{}]interface{})

		// If the value is also another map, then you need to retreive that value, adding it into the outer map.
		if isType {
			cleanInnerMap := cleanYaml(assertedVal)
			cleanYamlMapping[assertedKey] = cleanInnerMap
		}
	}

	return cleanYamlMapping
}

// FullYAMLToJSON is a wrapper function around the other underlying functions
// for ease of use. Simply, a file is specified and the conversion is handled internally.
func FullYAMLToJSON(filePath string) ([]byte, error) {

	yamlData, err := UnmarshalYAMLFile(filePath)
	if err != nil {
		return nil, err
	}

	jsonData, err := YAMLToJSON(yamlData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}