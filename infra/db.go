package infra

import (
	"log"

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
func NewSQLHandler() (*SQLHandler, error) {
	log.Println("[info] infra db new")
	sqlHandler := &SQLHandler{}
	err := sqlHandler.open()
	if err != nil {
		return nil, err
	}
	return sqlHandler, nil
}

func (sqlHandler *SQLHandler) open() error {
	c := config.NewMySQLDB()
	PROTOCOL := c.MySQL.Protocol
	USER := c.MySQL.Username
	PASS := c.MySQL.Password
	DBNAME := c.MySQL.DBName
	log.Println("[info] infra db open", c)

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
	log.Println("[info] infra db where", out, query, args)
	if sqlHandler.conn == nil {
		if err := sqlHandler.open(); err != nil {
			return err
		}
	}
	return sqlHandler.conn.Where(query, args).Find(out).Error
}
