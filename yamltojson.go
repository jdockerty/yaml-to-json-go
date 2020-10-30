package main

import (
	"fmt"
	// "reflect"
	// "strings"

	"github.com/jdockerty/yaml-to-json-go/convert"

)

// func inboundYamlHandler(s string) {

// 	if ( strings.Contains(s, ".json") || strings.Contains(s, ".yaml") || strings.Contains(s, ".yml") ) {
// 		fmt.Println("file.")
// 		unmarshalYamlFile(s)
// 	} else {

// 	}
// }






func main() {
	fmt.Println("Start")
	data := "test.yml"
	
	convert.FullYAMLToJSON(data)


}
