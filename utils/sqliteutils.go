package utils

import (
	"database/sql/driver"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
*
gorm.model包含

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
*/
type Searchhistory struct {
	gorm.Model
	Uid int
	Log string
}
type Datatype map[string](map[string]interface{})

type Repoinfo struct {
	gorm.Model
	Uid      int
	Reponame string   `json:"reponame"`
	Repourl  string   `json:"repourl"`
	Data     Datatype `json:"data"`
}

func (d *Datatype) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, d)
}
func (d Datatype) Value() (driver.Value, error) {
	return json.Marshal(d)
}

/*
*
插入查询结果
*/
func Insertsinglequery(reponame string, repourl string, data map[string](map[string]interface{})) error {
	db, err := gorm.Open(sqlite.Open("userDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Repoinfo{})

	tx := db.Create(&Repoinfo{Reponame: reponame, Repourl: repourl, Data: data})
	if tx.Error != nil {
		println(tx.Error)
	}
	return tx.Error
}

/*
*
插入命令行log
*/
func Insertlog(log string) error {
	db, err := gorm.Open(sqlite.Open("userDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Searchhistory{})

	tx := db.Create(&Searchhistory{Log: log})
	if tx.Error != nil {
		println(tx.Error)
	}
	return tx.Error
}
func Readlog(logs *[]Searchhistory) {
	db, err := gorm.Open(sqlite.Open("userDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	result := db.Find(&logs)
	println(result.Error)
}
func Readquery(repoinfo *Repoinfo, reponame string) {
	db, err := gorm.Open(sqlite.Open("userDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//db.First(repoinfo, 1)

	db.First(repoinfo, "reponame = ?", reponame)
	println(repoinfo.Reponame)
	println(repoinfo.Data["openrank"]["2020-08"].(float64))
}
