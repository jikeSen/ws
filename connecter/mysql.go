package connecter

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

var Db *gorm.DB
var err error

// 数据库连接
func MysqlConnect() {

    mysqlHost := ""
    mysqlPort := 62332
    mysqlUser := "root"
    mysqlPwd := ""
    mysqlDataBase := "qiyu"

    args := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", mysqlUser, mysqlPwd, mysqlHost, mysqlPort, mysqlDataBase)
    Db, err = gorm.Open("mysql", args)

    if err != nil {
        panic(err)
    }

    Db.DB().SetMaxIdleConns(5)
    Db.DB().SetMaxOpenConns(100)
    Db.SingularTable(true)
}
