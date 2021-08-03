package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Info string
	Data struct {
		Source      string
		Destination string
	}
	Run []struct {
		Id     string
		Exe    string
		Output string
	}
}

func main() {
	var conf Config
	reader, _ := os.Open("test.yml")
	buf, _ := ioutil.ReadAll(reader)
	yaml.Unmarshal(buf, &conf)
	fmt.Printf("%+v\n", conf)
}
