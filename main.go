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
	Port string
	Services [] struct {
		Url string
		Args [] struct {
			Name string
		}
		Command string
	}
}

func newConfig() Config {
	c := Config{}
	c.Port = "8090"

	return c
}

func startServer(c Config) {
	for _, service := range c.Services {
		http.HandleFunc(service.Url, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, service.Url)

			cmd := exec.Command("/bin/echo", "wow")
			out, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprint(w, string(out))
		})
	}
	err := http.ListenAndServe(strings.Join([]string{":", c.Port}, ""), nil)
	if nil != err {
		log.Fatal(err)
	}
}

func loadConfig(configFile string) Config {
	configData, err := ioutil.ReadFile(configFile)
	if nil != err {
		log.Fatalf("Can't read file %s", configFile)
	}

	c := newConfig()
	err = yaml.Unmarshal(configData, &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func getConfigFilePath() string {
	path := os.Getenv("SHERVICE_CONFIG")
	if len(path) > 0 {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Fatalf("File %s does not exist!", path)
		}

		return path
	}

	dir, err := os.Getwd()
	if nil != err {
		log.Fatal("Can not reach current working directory")
	}

	path = strings.Join([]string{dir, "/shervice.yaml"}, "")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("File %s does not exist!", path)
	}

	return path
}

func main() {
	configFile := getConfigFilePath()
	log.Printf("Config file %s", configFile)
	c := loadConfig(configFile)
	startServer(c)
}
