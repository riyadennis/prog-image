package middleware

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"context"
)
const (
	defaultConfigPath = "config.yaml"
	contextKey = "config"
)
type Config struct{
	Prog Prog `yaml:"prog"`
}
type Prog struct{
	Port int `yaml:"port"`
	Folder string `yaml:"folder"`
	FileType []string `yaml:"file_types"`
}
func ReadFromConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
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
	config := ctx.Value("config").(*Config)
	return config, nil
}
func ConfigMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("I am running from a middle ware")
		config, _ := ReadFromConfig(defaultConfigPath)

		newCtx := context.WithValue(r.Context(), contextKey, config)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}