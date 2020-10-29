package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	// "reflect"
	// "strings"
)

// func inboundYamlHandler(s string) {

// 	if ( strings.Contains(s, ".json") || strings.Contains(s, ".yaml") || strings.Contains(s, ".yml") ) {
// 		fmt.Println("file.")
// 		unmarshalYamlFile(s)
// 	} else {

// 	}
// }

// Allows YAML to be encoded into JSON format easily.
func cleanYaml(inputYaml map[interface{}]interface{}) map[string]interface{} {

	cleanYamlMapping := make(map[string]interface{})

	for key, value := range inputYaml {

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

func unmarshalYamlFile(filePath string) (map[string]interface{}, error) {

	unstructuredYAML := make(map[interface{}]interface{})

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

	err = yaml.Unmarshal(fileData, &unstructuredYAML)

	cleanedYaml := cleanYaml(unstructuredYAML)

	return cleanedYaml, nil

}

func convertYamlToJSON(yaml map[string]interface{}) ([]byte, error) {

	output, err := json.MarshalIndent(yaml, "", "\t")
	if err != nil {
		log.Printf("error converting yaml to json.\n%s", err.Error())
		return nil, err
	}
	fmt.Println(string(output))

	return output, nil
}

func main() {
	fmt.Println("Start")
	y, _ := unmarshalYamlFile("test.yml")
	log.Printf("%v\n", y)

	_, _ = convertYamlToJSON(y)


	// for k, v := range y {
	// 	log.Printf("key:%s, value:%s, key-type: %T, value-type: %T", k, v, k, v)
	// }
	// fmt.Println(f["age"])

}
