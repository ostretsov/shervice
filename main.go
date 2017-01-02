package main

import (
	"fmt"
	"github.com/olebedev/config"
	"io/ioutil"
)

func main() {
	configFile := "./config/shervice.yaml"
	configData, err := ioutil.ReadFile(configFile)
	if nil != err {
		fmt.Printf("Could not open file %s", configFile)
	}
	appConfig, err := config.ParseYaml(string(configData))
	if nil != err {
		fmt.Printf(err.Error())
	}

	value, err := appConfig.String("services.0.url")
	if nil != err {
		fmt.Println("Could not find a service!")
	}

	fmt.Println(value)
}
