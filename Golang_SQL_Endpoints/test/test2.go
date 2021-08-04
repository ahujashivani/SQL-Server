package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Data struct {
		Test string
	}
}

func main() {
	var conf Config
	reader, _ := os.Open("test.yml")
	buf, _ := ioutil.ReadAll(reader)
	yaml.Unmarshal(buf, &conf)
	fmt.Printf("%+v\n", conf)
}
