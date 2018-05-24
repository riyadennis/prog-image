package middleware

import (
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"context"
)

const (
	DefaultConfigPath = "config.yaml"
	ContextKey = "config"
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
func GetConfigFromContext(ctx context.Context) (*Config, error){
	config, ok := ctx.Value(ContextKey).(Config)

	if !ok {
		config, err := GetConfig(DefaultConfigPath)
		return config, err
	}
	return &config, nil
}
func GetConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil{
		return nil, err
	}
	fileReader := Reader{}
	config, err := fileReader.Read(file)
	if err != nil{
		return nil, err
	}
	return config, err
}
func ConfigMiddleWare(next http.Handler, config *Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newCtx := context.WithValue(r.Context(),ContextKey, config)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}
