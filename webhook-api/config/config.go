package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

type config struct {
	ServiceName string `json:"service_name"`
	Loglevel    string `json:"log_level"`
	Port        string `json:"port"`
	Hostname    string `json:"hostname"`
	Database    struct {
		Filename string `json:"filename"`
	} `json:"database"`
}

type cfg struct {
	logger *logrus.Logger
}

func New(logger *logrus.Logger) *cfg {
	return &cfg{logger: logger}
}

func (c *cfg) ReadConfig(filepath string) (*config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		c.logger.Errorf("[ERROR READING CONFIG] %+v", err.Error())
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		c.logger.Errorf("[ERROR READING CONFIG] %+v", err.Error())
		return nil, err
	}

	var config = &config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		c.logger.Errorf("[ERROR READING CONFIG] %+v", err.Error())
		return nil, err
	}

	return config, nil
}
