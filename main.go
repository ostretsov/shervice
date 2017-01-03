package main

import (
	"fmt"
	"github.com/olebedev/config"
	"io/ioutil"
	"net/http"
)

func loadConfig(configFile string) (string, error) {
	configData, err := ioutil.ReadFile(configFile)
	if nil != err {
		return "", err
	}
	appConfig, err := config.ParseYaml(string(configData))
	if nil != err {
		return "", err
	}
	value, err := appConfig.String("services.0.url")
	if nil != err {
		return "", err
	}

	return value, err;
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "echo is called: %s", r.URL.Path[1:])
}

func startServer() {
	http.HandleFunc("/echo", handler)
	err := http.ListenAndServe(":8080", nil)
	if nil != err {
		fmt.Println(err)
	}
}

func main() {
	configFile := "./config/shervice.yaml"
	value, err := loadConfig(configFile)
	if nil != err {
		fmt.Println(err)
	}

	fmt.Println(value)

	startServer()
}
