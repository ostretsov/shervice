package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os/exec"
	"log"
	"os"
	"strings"
)

type Config struct {
	Services [] struct {
		Url string
		Args [] struct {
			Name string
		}
		Command string
	}
}

func loadConfig(configFile string) Config {
	configData, err := ioutil.ReadFile(configFile)
	if nil != err {
		log.Fatalf("Can't open file %s", configFile)
	}

	c := Config{}
	err = yaml.Unmarshal(configData, &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
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
	c := loadConfig(configFile)

	fmt.Println(c.Services[0].Url)

	startServer()
}
