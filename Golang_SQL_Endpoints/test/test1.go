package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Data struct {
		Test string
	}
}

func main() {
	filename, _ := filepath.Abs("./test.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("Value: %#v\n", config.Data.Test)
	fmt.Println(config)
}
