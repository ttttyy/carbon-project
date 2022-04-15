package common

import (
	"carbon/lib"
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	workDir, _ := os.Getwd()
	println(workDir)
	cfg, err := ini.Load(workDir + "/conf/app.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		cfg.Section("database").Key("User").String(), cfg.Section("database").Key("Password").String(),
		cfg.Section("database").Key("Host").String(), cfg.Section("database").Key("Port").String(),
		cfg.Section("database").Key("Name").String(), cfg.Section("database").Key("Charset").String())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "carbon_",
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Fail to open database: ", err)
	}
	db.AutoMigrate(&lib.User{})
	DB = db
	return DB
}
func GetDB() *gorm.DB {
	return DB
}
