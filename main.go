package main

import (
	"fmt"
	"github.com/olebedev/config"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "echo is called: %s", r.URL.Path[1:])
}

func startServer() {
	http.HandleFunc("/echo", handler)
	http.ListenAndServe(":8080", nil)
}

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

	startServer()
}
