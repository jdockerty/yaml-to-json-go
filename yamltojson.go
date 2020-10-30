package main

import (
	"fmt"
	// "reflect"
	// "strings"

	"github.com/jdockerty/yaml-to-json-go/convert"

)





func main() {
	fmt.Println("Start")
	data := "test.yml"
	
	convert.FullYAMLToJSON(data)


}
