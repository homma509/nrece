package infra

import (
	"fmt"
	"os"

	"github.com/homma509/9go/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB ...
type DB struct {
	Host     string
	Username string
	Password string
	DBName   string
	Connect  *gorm.DB
}

// SQLHandler ... SQL handler struct
type SQLHandler struct {
	Conn *gorm.DB
}

// NewSQLHandler ...
func NewSQLHandler() *SQLHandler {
	c := config.NewDB()

	PROTOCOL := c.MySQL.Protocol
	USER := c.MySQL.Username
	PASS := c.MySQL.Password
	DBNAME := c.MySQL.DBName

	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		// panic(err.Error())
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

// Where ...
func (sqlHandler *SQLHandler) Where(out interface{}, query interface{}, args ...interface{}) error {
	return sqlHandler.Conn.Where(query, args).Find(out).Error
}
