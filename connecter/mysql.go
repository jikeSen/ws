package connecter

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "go/ws/library"
)

var Db *gorm.DB
var err error

// 数据库连接
func MysqlConnect() error {
    args := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", library.MysqlSet.User, library.MysqlSet.Password, library.MysqlSet.Host, library.MysqlSet.Port, library.MysqlSet.Name)
    Db, err = gorm.Open("mysql", args)

    if err != nil {
        return err
    }

    Db.DB().SetMaxIdleConns(5)
    Db.DB().SetMaxOpenConns(100)
    Db.SingularTable(true)
    return nil
}
