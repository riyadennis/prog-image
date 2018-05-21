package middleware

import (
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"io"
)

const (
	DefaultConfigPath = "config.yaml"
	contextKey        = "config"
)

type ConfigReader interface {
	Read(r io.Reader) (*Config, error)
}

type Config struct {
	Prog Prog `yaml:"prog"`
}
type Prog struct {
	Port     int      `yaml:"port"`
	Folder   string   `yaml:"folder"`
	FileType []string `yaml:"file_types"`
	Db       Db
}
type Db struct {
	Source   string
	Type     string
	User     string
	Password string
}
type Reader struct {
}

func (fr Reader) Read(r io.Reader) (*Config, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	c := Config{}

	err = yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
func GetConfig() (*Config, error) {
	file, _ := os.Open(DefaultConfigPath)

	fileReader := Reader{}
	config, err := fileReader.Read(file)
	return config, err
}
