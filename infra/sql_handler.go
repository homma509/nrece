package infra

import (
	"fmt"
	"os"

	"github.com/homma509/nrece/config"
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
	conn *gorm.DB
}

// NewSQLHandler ...
func NewSQLHandler() *SQLHandler {
	sqlHandler := &SQLHandler{}
	if err := sqlHandler.open(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}
	return sqlHandler
}

func (sqlHandler *SQLHandler) open() error {
	c := config.NewMySQLDB()
	PROTOCOL := c.MySQL.Protocol
	USER := c.MySQL.Username
	PASS := c.MySQL.Password
	DBNAME := c.MySQL.DBName

	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlHandler.conn = conn
	return nil
}

// Where ...
func (sqlHandler *SQLHandler) Where(out interface{}, query interface{}, args ...interface{}) error {
	if sqlHandler.conn == nil {
		if err := sqlHandler.open(); err != nil {
			return err
		}
	}
	return sqlHandler.conn.Where(query, args).Find(out).Error
}
