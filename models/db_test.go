package models

import (
	"testing"
	"github.com/prog-image/middleware"
	"github.com/stretchr/testify/assert"
	_ "github.com/mattn/go-sqlite3"
	"os"
)
const(
	testdbName = "test.db"
	testDriverType = "sqlite3"
)
func TestInitDBWithWrongDriverType(t *testing.T) {
	db := middleware.Db{Source: testdbName, Type: "sqlitse3"}
	defer tearDown()
	_, err := InitDB(db)
	assert.Error(t, err)
}
func TestInitDBWithCorrectDriverType(t *testing.T) {
	db := middleware.Db{Source: testdbName, Type: testDriverType}
	defer tearDown()
	_, err := InitDB(db)
	assert.NoError(t, err)
}
func TestSaveImageWithoutTable(t *testing.T) {
	db := middleware.Db{Source: testdbName, Type: testDriverType}
	defer tearDown()
	err := SaveImage("test.jpg", "test url", db)
	assert.Error(t, err)
}

func TestSaveImageWithTable(t *testing.T) {
	db := setUpDb()
	defer tearDown()
	err := SaveImage("test.jpg", "test url", db)
	assert.NoError(t, err)
}
func TestGetImageWithoutTable(t *testing.T) {
	db := middleware.Db{Source: "test.db", Type: "sqlite3"}
	defer tearDown()
	img, err := GetImage("test.jpg", db)
	assert.Error(t, err)
	assert.Empty(t, img)
}
func TestGetImageWithTable(t *testing.T) {
	db := setUpDb()
	err := SaveImage("test.jpg", "test url", db)
	defer tearDown()
	img, err := GetImage("test.jpg", db)
	assert.NoError(t, err)
	assert.Equal(t, img.Id, "test.jpg")
}
func setUpDb() (middleware.Db){
	db := middleware.Db{Source: "test.db", Type: "sqlite3"}
	createTable(db)
	return db
}
func tearDown(){
	os.Remove(testdbName)
}
func createTable(db middleware.Db){
	dbConnec, _ := InitDB(db)
	statement, _ := dbConnec.Prepare("CREATE TABLE IF NOT EXISTS "+tableName+"(id varchar(100) NOT NULL PRIMARY KEY,source varchar(100),imageType varchar(200),InsertedDatetime DATETIME);")
	statement.Exec()
}