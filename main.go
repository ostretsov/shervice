package main

import (
	"fmt"
	"github.com/olebedev/config"
	"io/ioutil"
	"net/http"
	"os/exec"
	"log"
	"os"
	"strings"
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

	return value, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("/bin/echo", "wow")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(out))
}

func startServer() {
	http.HandleFunc("/echo", handler)
	err := http.ListenAndServe(":8090", nil)
	if nil != err {
		fmt.Println(err)
	}
}

func getConfigFilePath() string {
	path := os.Getenv("SHERVICE_CONFIG")
	if len(path) > 0 {
		return path
	}

	dir, err := os.Getwd()
	if nil != err {
		log.Fatal("Can not reach current working directory")
	}

	p := []string{dir, "/shervice.yaml"}

	return strings.Join(p, "")
}

func main() {
	configFile := getConfigFilePath()
	log.Printf("Config file %s", configFile)
	value, err := loadConfig(configFile)
	if nil != err {
		fmt.Println(err)
	}

	fmt.Println(value)

	startServer()
}
