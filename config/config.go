package config

// DB ...
type DB struct {
	MySQL struct {
		Protocol string
		Username string
		Password string
		DBName   string
	}
}

// NewDB ...
func NewDB() *DB {
	config := new(DB)
	envs := NewEnvs()

	config.MySQL.Protocol = "tcp(" + envs.DBHost() + ":3306)"
	config.MySQL.Username = envs.DBUserName()
	config.MySQL.Password = envs.DBPassword()
	config.MySQL.DBName = envs.DBName()

	return config
}
