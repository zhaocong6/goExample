package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type dbType struct {
	Type      string `yaml:"type"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Database  string `yaml:"database"`
	Charset   string `yaml:"charset"`
	Collation string `yaml:"collation"`
}

var DB *dbType

var Config struct {
	DB *dbType `yaml:"db"`
}

//初始化模型配置
func init() {
	file, err := ioutil.ReadFile("conf.yml")
	if err != nil {
		log.Panicln(err)
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		log.Panicln(err)
	}

	DB = Config.DB
	initConnect()
}

//设置基础模型结构体
type Model struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

var db *gorm.DB

//初始化连接
func initConnect() {
	var err error
	db, err = gorm.Open(DB.Type,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=True&loc=Local",
			DB.User,
			DB.Password,
			DB.Host,
			DB.Port,
			DB.Database,
			DB.Charset,
			DB.Collation))
	if err != nil {
		log.Panicf("数据库连接失败:%s", err)
	}

	//默认使用单数
	db.SingularTable(true)
	//开启详情日志
	db.LogMode(true)
	//设置空闲中最大连接数
	db.DB().SetMaxIdleConns(10)
	//设置连接池最大数
	db.DB().SetMaxOpenConns(100)
	log.Println("数据库链接成功")
}

//关闭DB
func CloseDB() {
	if err := db.Close(); err != nil {
		log.Println(err)
	}
}

//设置save前置操作
func (m *Model) BeforeSave(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", Uint64ID())
	return err
}

//设置create前置操作
func (m *Model) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", Uint64ID())
	return err
}
