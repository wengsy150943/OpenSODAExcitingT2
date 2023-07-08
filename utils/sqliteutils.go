package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
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
type Datestype []string

type CachedRepoInfo struct {
	gorm.Model
	Uid      int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Reponame string
	Repourl  string
	Metric   string
	Month    string
	Dates    Datestype
	Data     Datatype
}

func (d *Datatype) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, d)
}
func (d Datatype) Value() (driver.Value, error) {
	return json.Marshal(d)
}
func (a *Datestype) Scan(value interface{}) error {
	bytes, ok := value.(string)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan Array value:", value))
	}
	*a = strings.Split(string(bytes), ",")
	return nil
}
func (a Datestype) Value() (driver.Value, error) {
	if len(a) > 0 {
		var str string = a[0]
		for _, v := range a[1:] {
			str += "," + v
		}
		return str, nil
	} else {
		return "", nil
	}
}

/*
*
创建表
*/
func CreateTable() {
	db, err := gorm.Open(sqlite.Open("userDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&CachedRepoInfo{})
}

/*
*
数据表是否存在
*/
func TableExist(tablename string) bool {
	db, err := gorm.Open(sqlite.Open("userDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	exist := db.Migrator().HasTable(tablename)
	return exist
}

/*
*
更新单个行
注：这里参数一定要用Datestype和Datatype，直接使用map[]Gorm会因为反射报错。
*/
func UpdateSingleRow(reponame string, metric string, dates Datestype, data Datatype) error {
	db, err := gorm.Open(sqlite.Open("userDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	repoinfo := CachedRepoInfo{}
	res := db.Model(&repoinfo).Where("reponame = ? AND metric = ?", reponame, metric).Updates(map[string]interface{}{"dates": dates, "data": data})
	return res.Error
}

/*
*
插入查询结果
*/
func InsertSingleQuery(reponame string, repourl string, metric string, month string, dates []string, data map[string](map[string]interface{})) error {
	//暂时使用全局路径，后面改相对路径
	db, err := gorm.Open(sqlite.Open("userDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&CachedRepoInfo{})

	tx := db.Create(&CachedRepoInfo{Reponame: reponame, Repourl: repourl, Metric: metric, Month: month, Dates: dates, Data: data})
	if tx.Error != nil {
		println(tx.Error)
	}
	return tx.Error
}

/*
*
查询特定仓库的数据
*/
func ReadQuerySingleMetric(repoinfo *CachedRepoInfo, reponame string, metric string) error {
	db, err := gorm.Open(sqlite.Open("userDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	metric = strings.ToLower(metric)

	result := db.Where("reponame = ? AND metric = ?", reponame, metric).First(repoinfo)

	return result.Error
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
	println(result)
}
