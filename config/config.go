package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// LoadConfig 从 YAML 文件加载配置
func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
