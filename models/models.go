package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gin-blog/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}
func (model *Model) BeforeCreate(tx *gorm.DB) error {
    tx.Statement.SetColumn("CreatedOn", time.Now().Unix())

    return nil
}

func (model *Model) BeforeUpdate(tx *gorm.DB) error {
    tx.Statement.SetColumn("ModifiedOn", time.Now().Unix())

    return nil
}


func init() {
	var (
		err error
		dbName, user, password, host,tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	// dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(
		mysql.Open( 
			fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName)),&gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   tablePrefix, // 表名前缀，`Tag` 的表名应该是 `blog_tag`
					SingularTable: true,    // 使用单数表名，启用该选项，此时，`Tag` 的表名应该是 `blog_tag`
				},

			})
		

	if err != nil {
		log.Println(err)
	}
	sqlDB,err := db.DB()
	if err != nil {
		log.Println(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)


	// db.SingularTable(true)
	// db.LogMode(true)
	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(100)
}

// func CloseDB() {
// sqlDB, err := db.DB()
// defer sqlDB.Close()
// if err != nil {
// fmt.Println(err)
// }
// }