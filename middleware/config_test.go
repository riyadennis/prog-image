package middleware

import (
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
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
