package main

import "fmt"
import "gopkg.in/yaml.v3"

type Source struct {
	subscriptions []Destination
}

type Destination struct {
	name string
	url  string
}

func main() {
    var yamlData map[string]interface;
    err := yaml.Unmarshal([]byte(data), &yamlData)
	fmt.Println("Hello")
}
