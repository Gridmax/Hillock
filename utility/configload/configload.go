package configload 

import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
)

type Conf struct {
  ServerPort string  `yaml:"server_port"`
  DatabaseType string  `yaml:"database"`
  DatabaseUrl string  `yaml:"db_url"`
  DatabaseUser string  `yaml:"db_user"`
  DatabasePw string  `yaml:"db_pw"`
  DatabaseName string  `yaml:"db_name"`
  DatabaseTable string  `yaml:"db_table"` 

}


func LoadConfig(filename string) (*Conf, error) {
	// Read the YAML file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Create a new Conf instance
	config := &Conf{}

	// Unmarshal the YAML data into the Conf struct
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
