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
	envs := NewEnvs()

	config.MySQL.Protocol = "tcp(" + envs.DBHost() + ":3306)"
	config.MySQL.Username = envs.DBUserName()
	config.MySQL.Password = envs.DBPassword()
	config.MySQL.DBName = envs.DBName()

	return config
}

// NewAWSConfig ...
func NewAWSConfig() *aws.Config {
	envs := NewEnvs()
	return &aws.Config{
		Region: aws.String(envs.Region()),
	}
}
