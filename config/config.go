package config

import "github.com/aws/aws-sdk-go/aws"

// DB ...
type DB struct {
	MySQL struct {
		Protocol string
		Username string
		Password string
		DBName   string
	}
}

// NewMySQLDB ...
func NewMySQLDB() *DB {
	config := new(DB)

	config.MySQL.Protocol = "tcp(" + Env().DBHost() + ":3306)"
	config.MySQL.Username = Env().DBUserName()
	config.MySQL.Password = Env().DBPassword()
	config.MySQL.DBName = Env().DBName()

	return config
}

// NewAWSConfig ...
func NewAWSConfig() *aws.Config {
	return &aws.Config{
		Region: aws.String(Env().Region()),
	}
}
