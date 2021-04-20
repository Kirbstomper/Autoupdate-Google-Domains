package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Username string
	Password string
	Hostname string
}

func main() {
	//Read configuration file
	configfile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	config := unmarshallConfig(configfile)
	fmt.Println(config)
}

// Umarshalls the configuration byte slice passed to a config struct
// and returns it
func unmarshallConfig(yml []byte) config {
	c := config{}
	err := yaml.Unmarshal(yml, &c)
	if err != nil {
		log.Fatalf("error: %v , \n there was a problem parsing your configuration file", err)
	}
	return c
}
