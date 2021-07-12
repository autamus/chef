package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Packages []string `yaml:"packages"`
}

func Load(chefyaml string) Conf {
	c := Conf{}
	yamlFile, err := ioutil.ReadFile(chefyaml)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
