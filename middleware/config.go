package middleware

import (
	"net/http"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"context"
	"io"
	"github.com/pkg/errors"
)

const (
	defaultConfigPath = "config.yaml"
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
func GetConfigFromContext(ctx context.Context) (*Config, error) {
	config, ok := ctx.Value("config").(*Config)
	if !ok {
		return nil, errors.New("invalid context")
	}
	return config, nil
}
func ConfigMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, _ := os.Open(defaultConfigPath)

		fileReader := Reader{}
		config, _ := fileReader.Read(file)

		newCtx := context.WithValue(r.Context(), contextKey, config)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}
