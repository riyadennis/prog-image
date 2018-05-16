package middleware

import (
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
	"context"
	"github.com/julienschmidt/httprouter"
)

func TestFileReader_ReadInvalidData(t *testing.T) {
	r := Reader{}
	config, err := r.Read(strings.NewReader("hello"))
	assert.Error(t, err)
	assert.Nil(t, config)
}
func TestFileReader_ReadValidData(t *testing.T) {
	r := Reader{}
	data := `
prog:
  port: 8990
  folder: "../images"
  file_types: ["jpg", "png", "gif"]
`
	config, err := r.Read(strings.NewReader(data))
	assert.NoError(t, err)
	assert.Equal(t, config.Prog.Port, 8990)
}
func TestGetConfigFromInvalidContext(t *testing.T) {
	ctx := context.Background()
	_, err := GetConfigFromContext(ctx)
	assert.Error(t, err)
}
func TestGetConfigFromValidContext(t *testing.T) {
	config := &Config{
		Prog:
		Prog{
			Port: 8080,
		},
	}
	newCtx := context.WithValue(context.Background(), contextKey, config)
	_, err := GetConfigFromContext(newCtx)
	assert.NoError(t, err)
}
func TestConfigMiddleWare(t *testing.T) {
	route := httprouter.New()
	handler := ConfigMiddleWare(route)
	assert.NotNil(t, handler)
}