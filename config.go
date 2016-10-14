package main

import (
	"encoding/json"
	"io/ioutil"
)

func readConfig(configFile string, defaultUsername string, defaultPort uint) ([]Server, error) {
	fileData, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var servers []Server
	if err = json.Unmarshal(fileData, &servers); err != nil {
		return nil, err
	}

	for i := 0; i < len(servers); i++ {
		servers[i].normalize(defaultUsername, defaultPort)
	}

	return servers, nil
}
