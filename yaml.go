package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	var data struct {
		Version string `yaml:"version"`
		Redis   struct {
			Default struct {
				Host string `yaml:"host"`
				Port string `yaml:"port"`
			} `yaml:"default"`
		} `yaml:"redis"`
	}

	str, _ := ioutil.ReadFile("test.yaml")
	yaml.Unmarshal(str, &data)

	fmt.Println(data)
}
