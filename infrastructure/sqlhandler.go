package infrastructure

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"echo-clean-architecture/api/database"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	name := os.Getenv("MYSQL_DATABASE")
	conn, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, name))
	if err != nil { panic(err.Error) }

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Update(sql string, values string) *gorm.DB {
	return handler.Conn.Update(sql, values)
}


func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}
