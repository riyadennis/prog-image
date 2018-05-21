package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateImageWithInvalidDirectory(t *testing.T) {
	i := NewImage("dd", "testfile.png", "https://fyf.tac-cdn.net/images/products/large/BF116-11KM.jpg")
	created, error := i.CreateImage()

	assert.Error(t, error)
	assert.Equal(t, false, created)
}

func TestCreateImageWithInvalidFileType(t *testing.T) {
	i := NewImage("../images", "testfile", "https://fyf.tac-cdn.net/images/products/large/BF116-11KM.jpg")
	i.ImageType = "invalid"
	created, error := i.CreateImage()
	assert.Error(t, error)
	assert.Equal(t, false, created)
}

func TestCreateImageWithValidFileName(t *testing.T) {
	i := NewImage("../images", "testfile.jpg", "https://fyf.tac-cdn.net/images/products/large/BF116-11KM.jpg")
	created, error := i.CreateImage()

	assert.NoError(t, error)
	assert.Equal(t, true, created)
}
func TestCreateImageWithInvalidSource(t *testing.T) {
	i := NewImage("../images", "testfile.jpg", "http:116-11KM.jpg")
	created, error := i.CreateImage()

	assert.Error(t, error)
	assert.Equal(t, false, created)
}
