package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestImageCreateImageWithInvalidDirectory(t *testing.T) {
	i := NewImage("dd", "testfile.png", "https://fyf.tac-cdn.net/images/products/large/BF116-11KM.jpg")
	created, error := i.CreateImage()

	assert.Error(t, error)
	assert.Equal(t, false, created)
}

func TestImageCreateImageWithInvalidFileName(t *testing.T) {
	i := NewImage("../images", "testfile", "https://fyf.tac-cdn.net/images/products/large/BF116-11KM.jpg")
	created, error := i.CreateImage()

	assert.Error(t, error)
	assert.Equal(t, false, created)
}

func TestImageCreateImageWithValidFileName(t *testing.T) {
	i := NewImage("../images", "testfile.jpg", "https://fyf.tac-cdn.net/images/products/large/BF116-11KM.jpg")
	created, error := i.CreateImage()

	assert.NoError(t, error)
	assert.Equal(t, true, created)
}
