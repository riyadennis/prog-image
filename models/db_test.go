package models

import (
	"testing"
	"github.com/prog-image/middleware"
	"github.com/stretchr/testify/assert"
	_ "github.com/mattn/go-sqlite3"
)

func TestInitDBWithWrongDriverType(t *testing.T) {
	db := middleware.Db{Source: "test.db", Type:"sqlitse3"}
	_, err := InitDB(db)
	assert.Error(t, err)
}
func TestInitDBWithCorrectDriverType(t *testing.T) {
	db := middleware.Db{Source: "test.db", Type:"sqlite3"}
	_, err := InitDB(db)
	assert.NoError(t, err)
}