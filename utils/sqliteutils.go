package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Repoinfo struct {
	id       int
	metric   string
	reponame string
	repourl  string
	data     []byte
	uid      int
}

func Create(dbname string) {
	fmt.Println("打开数据")
	db, err := sql.Open("sqlite3", "./userDB.db") //若数据库没有在这个项目文件下，则需要写绝对路径
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("生成数据表")
	sql_table := `
	CREATE TABLE IF NOT EXISTS "repoinfo" (
	   "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
	   "reponame" VARCHAR(128) NULL,
	   "repourl" VARCHAR(128) NULL,
		"metric" VARCHAR(64) NULL,
	    "data" VARCHAR(512) NULL,
	   "created" TIMESTAMP default (datetime('now', 'localtime'))  
	);`
	_, err = db.Exec(sql_table) //执行数据表
	if err != nil {
		println("error create table")
		log.Fatal(err)
	}
}

func Insert(metric string, reponame string, repourl string, data []byte) {
	db, err := gorm.Open(sqlite.Open("userDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Repoinfo{})

	db.Create(&Repoinfo{uid: 0, metric: metric, reponame: reponame, repourl: repourl, data: data})
}

func Read(repoinfo *Repoinfo, metric string) {
	db, err := gorm.Open(sqlite.Open("cache.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.First(repoinfo, "metric = ?", metric)
}
