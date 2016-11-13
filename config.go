package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func readConfig(configFile string, defaultUsername string, defaultPort uint) ([]Server, error) {
	fileData, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var servers []Server
	if err = yaml.Unmarshal(fileData, &servers); err != nil {
		return nil, err
	}

	for i := 0; i < len(servers); i++ {
		servers[i].normalize(defaultUsername, defaultPort)

		if err = servers[i].validate(); err != nil {
			return nil, err
		}
	}

	return servers, nil
}
