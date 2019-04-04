package config

// desc: 配置接口

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type cfg struct {
	Logger Logger

	Web Web

	Path Path
}

func newCfg(path string) (*cfg, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := &cfg{}
	if err := yaml.Unmarshal(data, conf); err != nil {
		return nil, fmt.Errorf("parse config err: %v", err)
	}

	return conf, nil
}
