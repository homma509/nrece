package config

import "os"

var envs *Envs

// Envs 環境変数を表した構造体
type Envs struct {
	Cache map[string]string
}

// NewEnvs ...
func NewEnvs() *Envs {
	return &Envs{
		Cache: map[string]string{},
	}
}

// Env ...
func Env() *Envs {
	if envs == nil {
		envs = NewEnvs()
	}
	return envs
}

func (e *Envs) env(key string) string {
	return os.Getenv(key)
}

// Region ...
func (e *Envs) Region() string {
	return e.env("REGION")
}

// DBHost ...
func (e *Envs) DBHost() string {
	return e.env("DB_HOST")
}

// DBUserName ...
func (e *Envs) DBUserName() string {
	return e.env("DB_USERNAME")
}

// DBPassword ...
func (e *Envs) DBPassword() string {
	return e.env("DB_PASSWORD")
}

// DBName ...
func (e *Envs) DBName() string {
	return e.env("DB_NAME")
}
