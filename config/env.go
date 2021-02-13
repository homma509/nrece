package config

import (
	"os"
	"strings"
)

var envs *Envs

// Envs 環境変数を表した構造体
type Envs struct{}

// NewEnvs ...
func newEnvs() *Envs {
	return &Envs{}
}

// Env ...
func Env() *Envs {
	if envs == nil {
		envs = newEnvs()
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

// BucketName ...
func (e *Envs) BucketName() string {
	return e.env("BUCKET_NAME")
}

// BucketFolderName ...
func (e *Envs) BucketFolderName() string {
	return e.env("BUCKET_FOLDER_NAME")
}

// Cluster ...
func (e *Envs) Cluster() string {
	return e.env("ECS_CLUSTER")
}

// Container ...
func (e *Envs) Container() string {
	return e.env("ECS_CONTAINER")
}

// TaskDefinition ...
func (e *Envs) TaskDefinition() string {
	return e.env("ECS_TASK_DEFINITION")
}

// Subnets ...
func (e *Envs) Subnets() []string {
	return strings.Split(e.env("ECS_SUBNETS"), ",")
}
